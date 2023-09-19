package middleware

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func TimeoutMiddleware(timeout time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)
		defer cancel()

		c.Request = c.Request.WithContext(ctx)

		done := make(chan struct{})

		go func() {
			defer close(done)
			c.Next()
		}()

		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			fmt.Println("timeout")
			c.JSON(http.StatusRequestTimeout, gin.H{"error": "Request timed out"})
			c.AbortWithStatus(http.StatusRequestTimeout)
			return
		}

		select {
		case <-done:
		case <-ctx.Done():
			fmt.Println("canceled")
			c.AbortWithStatus(http.StatusRequestTimeout)
		}
	}
}
