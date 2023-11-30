package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"time"
)

func TimeoutMiddleware(timeout time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)
		defer cancel()

		c.Request = c.Request.WithContext(ctx)
		c.Set("request_id", uuid.New().String())

		c.Next()
	}
}
