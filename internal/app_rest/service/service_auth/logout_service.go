package service_auth

import (
	"context"
	"errors"
	"github.com/sukha-id/bee/internal/app_rest/repositories/repo_auth"
	"go.uber.org/zap"
)

func (a *authService) Logout(ctx context.Context, token string) error {
	var (
		guid = ctx.Value("request_id").(string)
	)

	cLogger := zap.L().With(
		zap.String("layer", "service.logout"),
		zap.String("request_id", guid),
	)

	err := a.repoAuth.UpdateOne(ctx, repo_auth.AccessToken{Token: token})
	if err != nil {
		cLogger.Error("error update by token", zap.Error(err))
		err = errors.New("username or password is incorrect")
		return err
	}

	cLogger.Info("success logout service")
	return nil
}
