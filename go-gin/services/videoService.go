package services

import "github.com/ujjwolkayastha/go-gin/models"

// VideoService interface
type VideoService interface {
	Save(models.Video) models.Video
	FindAll() []models.Video
}

type videoService struct {
	videos []models.Video
}

//New is constructor function for video service
func New() VideoService {
	return &videoService{}
}

func (service *videoService) Save(video models.Video) models.Video {
	service.videos = append(service.videos, video)
	return video
}

func (service *videoService) FindAll() []models.Video {
	return service.videos
}
