package repo_auth

import (
	"context"
	"github.com/sukha-id/bee/pkg/logrusx"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthRepositoryInterface interface {
	FindOne(ctx context.Context, p Auth) (result *Auth, err error)
	FindOneAccessToken(ctx context.Context, p AccessToken) (result *AccessToken, err error)
	StoreAuth(ctx context.Context, p Auth) (err error)
	StoreAccessToken(ctx context.Context, p AccessToken) (err error)
	UpdateOne(ctx context.Context, p AccessToken) (err error)
}

type authRepository struct {
	mongoDB           *mongo.Client
	MongoDBCollection *mongo.Collection
	logger            *logrusx.LoggerEntry
}

func NewAuthRepository(mongoDB *mongo.Client, logger *logrusx.LoggerEntry) AuthRepositoryInterface {
	return &authRepository{
		mongoDB:           mongoDB,
		MongoDBCollection: mongoDB.Database("stock_collector").Collection("auth"),
		logger:            logger,
	}
}
