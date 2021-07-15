package infrastructure

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Router -> Gin Router
type Router struct {
	*gin.Engine
}

//NewRouter : all the routes are defined here
func NewRouter(env Env) Router {

	httpRouter := gin.Default()

	httpRouter.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))

	httpRouter.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "HOTELS API is Up and Running 🚀"})
	})

	return Router{
		Engine: httpRouter,
	}
}
