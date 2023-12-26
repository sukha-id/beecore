package repo_auth

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

func (a *authRepository) UpdateOne(ctx context.Context, p AccessToken) (err error) {
	update := bson.D{{"$set", bson.D{{"revoke", true}}}}
	filter := bson.D{{"token", p.Token}}

	_, err = a.mongoDB.
		Database("stock_collector").
		Collection("access_token").UpdateOne(ctx, filter, update)
	if err != nil {
		return
	}

	return
}
