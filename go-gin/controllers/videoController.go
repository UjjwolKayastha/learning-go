package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ujjwolkayastha/go-gin/models"
	"github.com/ujjwolkayastha/go-gin/services"
)

//VideoController interface
type VideoController interface {
	FindAll() []models.Video
	Save(ctx *gin.Context) models.Video
}

type controller struct {
	service services.VideoService
}

//New is constructor controller function
func New(service services.VideoService) VideoController {
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []models.Video {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) models.Video {
	var video models.Video
	ctx.BindJSON(&video)
	c.service.Save(video)

	return video
}
