package repo_auth

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (a *authRepository) FindOne(ctx context.Context, p Auth) (*Auth, error) {
	var result Auth
	bsonBytes, err := bson.Marshal(&p)
	if err != nil {
		return nil, err
	}

	err = a.mongoDB.
		Database("stock_collector").
		Collection("auth").
		FindOne(ctx, bsonBytes).
		Decode(&result)

	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, err
	}

	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return nil, err
	}

	return &result, nil
}
