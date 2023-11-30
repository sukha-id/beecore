package repo_auth

import (
	"time"
)

type Auth struct {
	UserID    string    `json:"user_id" bson:"user_id,omitempty"`
	Password  string    `json:"Password" bson:"password,omitempty" validate:"required,min=6"`
	Username  string    `json:"email" bson:"username,omitempty" validate:"email,required"`
	CreatedAt time.Time `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at,omitempty"`
}

type AccessToken struct {
	Token        string    `json:"token" bson:"token,omitempty"`
	RefreshToken string    `json:"refresh_token" bson:"refresh_token,omitempty"`
	Revoke       bool      `json:"revoke" bson:"revoke"`
	CreatedAt    time.Time `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAT    time.Time `json:"updated_at" bson:"updated_at,omitempty"`
	ExpiredAt    time.Time `json:"expired_at" bson:"expired_at,omitempty"`
	UserID       string    `json:"user_id" bson:"user_id,omitempty"`
}
