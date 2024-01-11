package service_auth

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/sukha-id/bee/internal/app_rest/repositories/repo_auth"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func (a *authService) SignUp(ctx context.Context, l SignUpPayload) (*SignUpResponse, error) {
	var (
		guid   = ctx.Value("request_id").(string)
		result SignUpResponse
	)

	cLogger := zap.L().With(
		zap.String("layer", "service.logout"),
		zap.String("request_id", guid),
	)

	existingAuth, err := a.repoAuth.FindOne(ctx, repo_auth.Auth{Username: l.Username})
	if err != nil {
		cLogger.Error("err find one auth", zap.Error(err))
		return nil, err
	}

	if existingAuth != nil {
		return nil, errors.New("this username already exists")
	}

	password, err := HashPassword(l.Password)
	if err != nil {
		cLogger.Error("err hash password", zap.Error(err))
		return nil, errors.New("err your password")
	}

	auth := repo_auth.Auth{
		UserID:    uuid.New().String(),
		Password:  password,
		Username:  l.Username,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	auth.Password = password

	token, refreshToken, _ := a.jwtAuth.GenerateAllTokens(
		auth.UserID)

	accessToken := repo_auth.AccessToken{
		Token:        token,
		RefreshToken: refreshToken,
		Revoke:       false,
		CreatedAt:    time.Now(),
		UpdatedAT:    time.Now(),
		ExpiredAt:    time.Now().Local().Add(time.Hour * time.Duration(168)),
		UserID:       auth.UserID,
	}

	err = a.repoAuth.StoreAuth(ctx, auth)
	if err != nil {
		cLogger.Error("err store auth", zap.Error(err))
		return nil, err
	}

	err = a.repoAuth.StoreAccessToken(ctx, accessToken)
	if err != nil {
		cLogger.Error("err store auth", zap.Error(err))
		return nil, err
	}

	result.Token = accessToken.Token
	result.RefreshToken = accessToken.RefreshToken

	cLogger.Error("success service signup")
	return &result, err
}

// HashPassword is used to encrypt the password before it is stored in the DB
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", nil
	}

	return string(bytes), nil
}
