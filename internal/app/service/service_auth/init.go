package service_auth

import (
	"context"
	"github.com/sukha-id/bee/internal/app/middleware/jwtx"
	"github.com/sukha-id/bee/internal/app/repositories/repo_auth"
	"github.com/sukha-id/bee/pkg/logrusx"
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
	logger   *logrusx.LoggerEntry
}

func NewAuthService(logger *logrusx.LoggerEntry, repoAuth repo_auth.AuthRepositoryInterface, jwtAuth jwtx.AuthenticationInterface) AuthServiceInterface {
	return &authService{
		repoAuth: repoAuth,
		jwtAuth:  jwtAuth,
		logger:   logger,
	}
}
