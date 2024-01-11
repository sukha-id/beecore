package service_auth

import (
	"context"
	"github.com/sukha-id/bee/internal/app_rest/middleware/jwtx"
	"github.com/sukha-id/bee/internal/app_rest/repositories/repo_auth"
)

type AuthServiceInterface interface {
	Login(ctx context.Context, l LoginPayload) (result *LoginResponse, err error)
	SignUp(ctx context.Context, l SignUpPayload) (result *SignUpResponse, err error)
	Profile(ctx context.Context, userID string, username string) (*ProfileReturn, error)
	Logout(ctx context.Context, token string) error
	RefreshToken(ctx context.Context, refreshToken string) (*RefreshTokenResponse, error)
}

type authService struct {
	repoAuth repo_auth.AuthRepositoryInterface
	jwtAuth  jwtx.AuthenticationInterface
}

func NewAuthService(repoAuth repo_auth.AuthRepositoryInterface, jwtAuth jwtx.AuthenticationInterface) AuthServiceInterface {
	return &authService{
		repoAuth: repoAuth,
		jwtAuth:  jwtAuth,
	}
}
