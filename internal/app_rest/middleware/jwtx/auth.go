package jwtx

import (
	"github.com/gin-gonic/gin"
	"github.com/sukha-id/bee/internal/app_rest/repositories/repo_auth"
	"github.com/sukha-id/bee/pkg/ginx"
	"go.uber.org/zap"
	"net/http"
	"strings"
)

func (j *AuthenticationJWT) Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			guid = ctx.Value("request_id").(string)
		)

		cLogger := zap.L().With(
			zap.String("layer", "service.login"),
			zap.String("request_id", guid),
		)

		authorization := ctx.Request.Header.Get("Authorization")
		clientToken := strings.TrimPrefix(authorization, "Bearer ")
		if clientToken == "" {
			ginx.RespondWithJSON(
				ctx,
				http.StatusUnauthorized,
				"No Authorization header provided",
				nil,
			)
			return
		}

		existingToken, err := j.repoAuth.FindOneAccessToken(ctx, repo_auth.AccessToken{Token: clientToken, Revoke: true})
		if err != nil {
			cLogger.Error("err find one access token", zap.Error(err))
			ginx.RespondWithJSON(
				ctx,
				http.StatusInternalServerError,
				http.StatusText(http.StatusInternalServerError),
				nil,
			)
			return
		}

		if existingToken != nil && existingToken.Revoke {
			ginx.RespondWithJSON(
				ctx,
				http.StatusUnauthorized,
				"Something wrong with authorization",
				nil,
			)
			return
		}

		claims, msg := j.ValidateToken(clientToken)
		if msg != "" {
			ginx.RespondWithJSON(
				ctx,
				http.StatusUnauthorized,
				"Invalid Authorization",
				nil,
			)
			return
		}

		ctx.Set("username", claims.Username)
		ctx.Set("token", clientToken)
		ctx.Set("uid", claims.UserID)

		ctx.Next()
	}
}
