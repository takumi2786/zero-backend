package usecase

import (
	"context"

	"github.com/takumi2786/zero-backend/internal/domain"
	"go.uber.org/zap"
	"golang.org/x/xerrors"
)

var (
	FailedToAuthorise        = xerrors.New("authentication failed")
	FailedToGenerateToken    = xerrors.New("failed to generate token")
	AuthExistButUserNotExist = xerrors.New("auth user exists but user does not exist")
)

type LoginInteractor struct {
	logger             *zap.Logger
	userRepository     domain.UserRepository
	authUserRepository domain.AuthUserRepository
	tokenGenerator     TokenGenerator
}

func NewLoginInteractor(
	logger *zap.Logger,
	userRepository domain.UserRepository,
	authUserRepository domain.AuthUserRepository,
	tokenGenerator TokenGenerator,
) LoginUsecase {
	return &LoginInteractor{
		logger:             logger,
		userRepository:     userRepository,
		authUserRepository: authUserRepository,
		tokenGenerator:     tokenGenerator,
	}
}

/*
指定した認証方法で認証を実行し、ユーザーを特定する。
*/
func (lu *LoginInteractor) authentication(
	ctx context.Context,
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
	auth, err := lu.authUserRepository.GetByIdentifier(ctx, identityType, identifier)
	if err != nil {
		return nil, err
	}
	if auth == nil {
		return nil, nil
	}
	if auth.Credential != credential {
		return nil, nil
	}
	user, err := lu.userRepository.GetByUserId(ctx, auth.UserId)
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
func (lu *LoginInteractor) Login(ctx context.Context, identityType string, identifier string, credential string) (*string, error) {
	user, err := lu.authentication(ctx, identityType, identifier, credential)
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
