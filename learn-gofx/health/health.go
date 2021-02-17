package health

import (
	"al.com/handler"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func registerRoutes(handler *handler.Handler) {
	handler.Gin.GET("/health-check", func(c *gin.Context) {
		c.JSON(200, gin.H{"response": "HEALTH OKAY"})
	})
}

// Module for registering routes
var Module = fx.Options(fx.Invoke(registerRoutes))
