package service_auth

import (
	"context"
	"errors"
	"github.com/sukha-id/bee/internal/app/repositories/repo_auth"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func (a *authService) Login(ctx context.Context, l LoginPayload) (*LoginResponse, error) {
	var (
		guid   = ctx.Value("request_id").(string)
		result LoginResponse
	)
	auth, err := a.repoAuth.FindOne(ctx, repo_auth.Auth{Username: l.Username})
	if err != nil {
		a.logger.Error(guid, "error find auth by username", err)
		err = errors.New("username or password is incorrect")
		return nil, err
	}

	passwordIsValid, err := VerifyPassword(l.Password, auth.Password)
	if passwordIsValid != true || err != nil {
		a.logger.Error(guid, "error verify password", err)
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
		a.logger.Error(guid, "error find auth by user id", err)
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
		a.logger.Error(guid, "err", err)

		return nil, err
	}

	result.Token = token
	result.RefreshToken = refreshToken

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
