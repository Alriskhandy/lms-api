package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const RequestIDKey = "X-Request-ID"

func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		reqID := c.GetHeader(RequestIDKey)
		if reqID == "" {
			reqID = uuid.New().String()
		}

		// Set ke context Gin
		c.Set(RequestIDKey, reqID)

		// Juga set ke response header
		c.Writer.Header().Set(RequestIDKey, reqID)

		c.Next()
	}
}
