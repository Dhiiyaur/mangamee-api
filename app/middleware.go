package app

import (
	"context"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func requestId() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := uuid.New()
		c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), "request_id", userID.String()))
		c.Next()
	}
}

func MiddlewareSetup(r *gin.Engine) {

	r.Use(cors.Default())
	r.Use(requestId())
	r.Use(gin.Recovery())
	r.Use(gin.Logger())
}
