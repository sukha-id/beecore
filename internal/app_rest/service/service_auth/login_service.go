package service_auth

import (
	"context"
	"errors"
	"github.com/sukha-id/bee/internal/app_rest/repositories/repo_auth"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func (a *authService) Login(ctx context.Context, l LoginPayload) (*LoginResponse, error) {
	var (
		guid   = ctx.Value("request_id").(string)
		result LoginResponse
	)
	cLogger := zap.L().With(
		zap.String("layer", "service.login"),
		zap.String("request_id", guid),
	)

	auth, err := a.repoAuth.FindOne(ctx, repo_auth.Auth{Username: l.Username})
	if err != nil {
		cLogger.Error("error find auth by username", zap.Error(err))
		err = errors.New("username or password is incorrect")
		return nil, err
	}

	passwordIsValid, err := VerifyPassword(l.Password, auth.Password)
	if passwordIsValid != true || err != nil {
		cLogger.Error("error verify password", zap.Error(err))
		err = errors.New("username or password is incorrect")
		return nil, err
	}

	if auth == nil {
		return nil, err
	}
	token, refreshToken, _ := a.jwtAuth.GenerateAllTokens(
		auth.UserID,
	)

	auth, err = a.repoAuth.FindOne(ctx, repo_auth.Auth{UserID: auth.UserID})
	if err != nil {
		cLogger.Error("error find auth by user id", zap.Error(err))
		return nil, err
	}

	accessToken := repo_auth.AccessToken{
		Token:        token,
		RefreshToken: refreshToken,
		Revoke:       false,
		CreatedAt:    time.Now(),
		UpdatedAT:    time.Now(),
		ExpiredAt:    time.Now().Local().Add(time.Hour * time.Duration(168)),
		UserID:       auth.UserID,
	}

	err = a.repoAuth.StoreAccessToken(ctx, accessToken)
	if err != nil {
		cLogger.Error("error", zap.Error(err))
		return nil, err
	}

	result.Token = token
	result.RefreshToken = refreshToken

	cLogger.Info("success service login")
	return &result, nil
}

func VerifyPassword(userPassword string, hashedPassword string) (bool, error) {
	check := true
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(userPassword))

	if err != nil {
		check = false
		return false, err
	}

	return check, err
}
