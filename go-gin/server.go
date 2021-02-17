package main

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
	"github.com/ujjwolkayastha/go-gin/controllers"
	"github.com/ujjwolkayastha/go-gin/middlewares"
	"github.com/ujjwolkayastha/go-gin/services"
)

var (
	videoService    services.VideoService       = services.New()
	videoController controllers.VideoController = controllers.New(videoService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	setupLogOutput()
	server := gin.New()

	server.Use(gin.Recovery(), middlewares.Logger(), middlewares.Auth(), gindump.Dump())

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
