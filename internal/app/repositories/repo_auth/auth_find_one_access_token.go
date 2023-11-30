package repo_auth

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (a *authRepository) FindOneAccessToken(ctx context.Context, p AccessToken) (*AccessToken, error) {
	var result AccessToken
	bsonBytes, err := bson.Marshal(&p)
	if err != nil {
		return nil, err
	}

	err = a.mongoDB.
		Database("stock_collector").
		Collection("access_token").
		FindOne(ctx, bsonBytes).
		Decode(&result)

	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return nil, err
	}

	return &result, nil
}
