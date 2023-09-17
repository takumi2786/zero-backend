package usecase

import (
	"context"

	"github.com/takumi2786/zero-backend/internal/domain"
	"go.uber.org/zap"
)

type LoginUsecase struct {
	logger             *zap.Logger
	userRepository     domain.UserRepository
	authUserRepository domain.AuthUserRepository
}

func NewLoginUsecase(userRepository domain.UserRepository, authUserRepository domain.AuthUserRepository) *LoginUsecase {
	return &LoginUsecase{userRepository: userRepository, authUserRepository: authUserRepository}
}

type AuthenticationOutput struct {
	UserId int64
	Name   string
}

/*
指定した認証方法で認証を実行し、ユーザーを特定する。
*/
func (lu *LoginUsecase) authentication(ctx context.Context, identityType string, identity string, credential string) (*AuthenticationOutput, error) {
	if identityType != "email" {
		lu.logger.Error(
			"unsupported identity type was specified",
			zap.String("identityType", identityType),
		)
		return nil, nil
	}
	auth, err := lu.authUserRepository.GetByIdentity(ctx, identity, identityType)
	if err != nil {
		return nil, err
	}
	if auth == nil {
		return nil, nil
	}
	if auth.Credential != credential {
		return nil, nil
	}
	// 認証情報が正しいのでユーザー情報を取得
	user, err := lu.userRepository.GetByUserId(ctx, auth.UserId)
	if err != nil {
		return nil, err
	}
	if user == nil {
		// データが不正なのでエラーログを出力
		lu.logger.Error(
			"auth user exists but user does not exist",
			zap.Int64("userId", auth.UserId),
		)
		return nil, nil
	}
	// ログイン成功
	return &AuthenticationOutput{UserId: user.UserId, Name: user.Name}, nil
}

// /*
// 認証を実行したのちにトークンを生成する。
// */
// func (lu *LoginUsecase) Login(ctx context.Context, identityType string, identity string, credential string) (*string, error) {
// 	output, err := lu.authentication(ctx, identityType, identity, credential)
// 	if err != nil {
// 		return nil, err
// 	}
// 	// generate token
// 	return nil, nil
// }
