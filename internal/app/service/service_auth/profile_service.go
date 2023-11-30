package service_auth

import (
	"context"
	"errors"
	"github.com/sukha-id/bee/internal/app/repositories/repo_auth"
)

func (a *authService) Profile(ctx context.Context, userID string, username string) (*ProfileReturn, error) {
	var (
		guid   = ctx.Value("request_id").(string)
		result ProfileReturn
	)

	auth, err := a.repoAuth.FindOne(ctx, repo_auth.Auth{Username: username, UserID: userID})
	if err != nil {
		a.logger.Error(guid, "error find auth by username", err)
		err = errors.New("username or password is incorrect")
		return nil, err
	}

	result.Username = auth.Username

	return &result, nil
}
