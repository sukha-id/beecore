package repo_auth

import "context"

func (a *authRepository) StoreAccessToken(ctx context.Context, p AccessToken) (err error) {
	_, err = a.mongoDB.Database("stock_collector").Collection("access_token").InsertOne(ctx, p)
	if err != nil {
		return
	}

	return
}
