package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/myorg/lms-backend/internal/logger"
	"go.uber.org/zap"
)

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// Proses request
		c.Next()

		latency := time.Since(start)
		status := c.Writer.Status()
		method := c.Request.Method
		path := c.Request.URL.Path
		clientIP := c.ClientIP()
		reqID := c.GetString(RequestIDKey)

		logger.Log.Info("incoming request",
			// field JSON
			// contoh: { "method": "GET", "status": 200, ... }
			zap.String("request_id", reqID),
			zap.Int("status", status),
			zap.String("method", method),
			zap.String("path", path),
			zap.String("client_ip", clientIP),
			zap.Duration("latency", latency),
		)
	}
}
