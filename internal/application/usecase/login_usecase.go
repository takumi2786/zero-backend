package usecase

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/takumi2786/zero-backend/internal/domain"
	"github.com/takumi2786/zero-backend/internal/util"
	"go.uber.org/zap"
	"golang.org/x/xerrors"
)

type ILoginUsecase interface {
	Login(identityType string, identifier string, credential string) (*string, error)
}

var (
	FailedToAuthorise        = xerrors.New("authentication failed")
	FailedToGenerateToken    = xerrors.New("failed to generate token")
	AuthExistButUserNotExist = xerrors.New("auth user exists but user does not exist")
)

type LoginUsecase struct {
	cfg                *util.Config
	logger             *zap.Logger
	userRepository     domain.UserRepository
	authUserRepository domain.AuthUserRepository
	tokenGenerator     ITokenGenerator
}

func NewLoginUsecase(
	cfg *util.Config,
	logger *zap.Logger,
	userRepository domain.UserRepository,
	authUserRepository domain.AuthUserRepository,
	tokenGenerator ITokenGenerator,
) ILoginUsecase {
	return &LoginUsecase{
		cfg:                cfg,
		logger:             logger,
		userRepository:     userRepository,
		authUserRepository: authUserRepository,
		tokenGenerator:     tokenGenerator,
	}
}

/*
指定した認証方法で認証を実行し、ユーザーを特定する。
*/
func (lu *LoginUsecase) authentication(
	identityType string,
	identifier string,
	credential string,
) (*domain.User, error) {
	if identityType != "email" {
		lu.logger.Error(
			"unsupported identity type was specified",
			zap.String("identityType", identityType),
		)
		return nil, nil
	}
	auth, err := lu.authUserRepository.GetByIdentifier(identityType, identifier)
	if err != nil {
		return nil, err
	}
	if auth == nil {
		return nil, nil
	}
	if auth.Credential != credential {
		return nil, nil
	}
	user, err := lu.userRepository.GetByUserId(auth.UserId)
	if err != nil {
		return nil, err
	}
	if user == nil {
		// データが不正なのでエラーログを出力
		error := AuthExistButUserNotExist
		lu.logger.Error(error.Error())
		return nil, error
	}
	return user, nil
}

/*
認証を実行したのちにトークンを生成する。
*/
func (lu *LoginUsecase) Login(identityType string, identifier string, credential string) (*string, error) {
	user, err := lu.authentication(identityType, identifier, credential)
	if err != nil {
		return nil, xerrors.Errorf("%v: %w", FailedToAuthorise, err)
	}
	if user == nil {
		return nil, FailedToAuthorise
	}
	// generate token
	token, err := lu.tokenGenerator.GenerateToken(user.UserId)
	if err != nil {
		lu.logger.Error("failed to generate token", zap.Error(err))
		return nil, xerrors.Errorf("%v: %w", FailedToGenerateToken, err)
	}
	return token, nil
}

/*
トークン生成処理
*/
type ITokenGenerator interface {
	GenerateToken(userId int64) (*string, error)
}

// JWTTokenGeneratorは、ITokenGeneratorを実装します。
type JWTTokenGenerator struct {
}

func NewJWTTokenGenerator() ITokenGenerator {
	return &JWTTokenGenerator{}
}

func (jtg *JWTTokenGenerator) GenerateToken(userId int64) (*string, error) {
	claims := jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(""))
	if err != nil {
		return nil, err
	}
	return &tokenString, nil
}
