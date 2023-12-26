package service_auth

import (
	"context"
	"errors"
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/sukha-id/bee/internal/app_rest/repositories/repo_auth"
	"time"
)

func (a *authService) RefreshToken(ctx context.Context, rToken string) (*RefreshTokenResponse, error) {
	var (
		guid   = ctx.Value("request_id").(string)
		result RefreshTokenResponse
	)

	token, err := a.jwtAuth.ParseRefreshToken(rToken)
	if err != nil {
		a.logger.Error(guid, "err parse jwt refresh token", err)
		err = errors.New("unauthorized")
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims) //the token claims should conform to MapClaims
	if !ok && !token.Valid {
		a.logger.Error(guid, "err claim refresh token", err)
		err = errors.New("unauthorized")
		return nil, err
	}

	userID, ok := claims["UserID"].(string) //convert the interface to string
	if !ok {
		a.logger.Error(guid, "err get claims user id", err)
		err = errors.New("unauthorized")
		return nil, err
	}

	auth, err := a.repoAuth.FindOne(ctx, repo_auth.Auth{UserID: userID})
	if err != nil {
		a.logger.Error(guid, "error find auth by user id", err)
		return nil, err
	}

	if auth == nil {
		err = errors.New("user not found")
		return nil, err
	}

	newToken, newRefreshToken, _ := a.jwtAuth.GenerateAllTokens(
		userID,
	)

	accessToken := repo_auth.AccessToken{
		Token:        newToken,
		RefreshToken: newRefreshToken,
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

	result.Token = newToken
	result.RefreshToken = newRefreshToken

	return &result, nil
}
