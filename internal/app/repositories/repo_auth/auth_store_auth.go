package repo_auth

import "context"

func (a *authRepository) StoreAuth(ctx context.Context, p Auth) (err error) {
	_, err = a.MongoDBCollection.InsertOne(ctx, p)
	if err != nil {
		return
	}

	return
}
