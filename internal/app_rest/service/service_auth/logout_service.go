package service_auth

import (
	"context"
	"errors"
	"github.com/sukha-id/bee/internal/app_rest/repositories/repo_auth"
)

func (a *authService) Logout(ctx context.Context, token string) error {
	var (
		guid = ctx.Value("request_id").(string)
	)

	err := a.repoAuth.UpdateOne(ctx, repo_auth.AccessToken{Token: token})
	if err != nil {
		a.logger.Error(guid, "error update by token", err)
		err = errors.New("username or password is incorrect")
		return err
	}

	return nil
}
