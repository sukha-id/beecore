package service_auth

type LoginPayload struct {
	Username string `json:"username" bson:"username,omitempty"`
	Password string `json:"password" bson:"password,omitempty"`
}

type LoginResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}
