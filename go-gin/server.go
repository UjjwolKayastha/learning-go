package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ujjwolkayastha/go-gin/controllers"
	"github.com/ujjwolkayastha/go-gin/services"
)

var (
	videoService    services.VideoService       = services.New()
	videoController controllers.VideoController = controllers.New(videoService)
)

func main() {
	server := gin.Default()

	server.GET("/health-check", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "API IS HEALTHY",
		})
	})

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.Save(ctx))
	})

	server.Run(":5000")
}
