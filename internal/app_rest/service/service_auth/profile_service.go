package service_auth

import (
	"context"
	"errors"
	"github.com/sukha-id/bee/internal/app_rest/repositories/repo_auth"
	"go.uber.org/zap"
)

func (a *authService) Profile(ctx context.Context, userID string, username string) (*ProfileReturn, error) {
	var (
		guid   = ctx.Value("request_id").(string)
		result ProfileReturn
	)

	cLogger := zap.L().With(
		zap.String("layer", "service.profile"),
		zap.String("request_id", guid),
	)

	auth, err := a.repoAuth.FindOne(ctx, repo_auth.Auth{Username: username, UserID: userID})
	if err != nil {
		cLogger.Error("error find auth by username", zap.Error(err))
		err = errors.New("username or password is incorrect")
		return nil, err
	}

	result.Username = auth.Username

	cLogger.Info("success service profile")
	return &result, nil
}
