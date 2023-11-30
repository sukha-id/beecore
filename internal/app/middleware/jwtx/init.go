package jwtx

import (
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gin-gonic/gin"
	"github.com/sukha-id/bee/config"
	"github.com/sukha-id/bee/internal/app/repositories/repo_auth"
	"github.com/sukha-id/bee/pkg/logrusx"
)

type AuthenticationJWT struct {
	cfg      *config.ConfigApp
	repoAuth repo_auth.AuthRepositoryInterface
	logger   *logrusx.LoggerEntry
}
type AuthenticationInterface interface {
	Authentication() gin.HandlerFunc
	GenerateAllTokens(userID string) (signedToken string, signedRefreshToken string, err error)
	ValidateToken(signedToken string) (claims *SignedDetails, msg string)
	ParseRefreshToken(rToken string) (*jwt.Token, error)
}

func NewJWTAuthentication(
	cfg *config.ConfigApp,
	repoAuth repo_auth.AuthRepositoryInterface,
	logger *logrusx.LoggerEntry) AuthenticationInterface {
	return &AuthenticationJWT{
		cfg:      cfg,
		repoAuth: repoAuth,
		logger:   logger,
	}
}
