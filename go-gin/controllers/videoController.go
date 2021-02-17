package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/ujjwolkayastha/go-gin/models"
	"github.com/ujjwolkayastha/go-gin/services"
	"github.com/ujjwolkayastha/go-gin/validators"
)

//VideoController interface
type VideoController interface {
	FindAll() []models.Video
	Save(ctx *gin.Context) error
}

type controller struct {
	service services.VideoService
}

var coolValidation *validator.Validate

//New is constructor controller function
func New(service services.VideoService) VideoController {
	coolValidation = validator.New()
	coolValidation.RegisterValidation("is-cool", validators.ValidateCoolTitle)
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []models.Video {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) error {
	var video models.Video
	err := ctx.ShouldBindJSON(&video)

	if err != nil {
		return err
	}
	err = coolValidation.Struct(video)
	if err != nil {
		return err
	}
	c.service.Save(video)
	return nil
}
