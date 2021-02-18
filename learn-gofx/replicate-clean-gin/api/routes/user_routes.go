package routes

import (
	"github.com/ujjwolkayastha/learn-clean-gin/api/controllers"
	"github.com/ujjwolkayastha/learn-clean-gin/lib"
)

// UserRoutes struct
type UserRoutes struct {
	logger         lib.Logger
	handler        lib.RequestHandler
	userController controllers.UserController
}

// Setup user routes
func (s UserRoutes) Setup() {
	s.logger.Zap.Info("Setting up routes")
	api := s.handler.Gin.Group("/api")
	{
		api.GET("/users", s.userController.GetUsers)
		api.GET("/user/:id", s.userController.GetOneUser)
		api.POST("/user", s.userController.SaveUser)
		api.POST("/user/:id", s.userController.UpdateUser)
		api.DELETE("/user/:id", s.userController.DeleteUser)
	}
}

// NewUserRoutes creates new user controller
func NewUserRoutes(
	logger lib.Logger,
	handler lib.RequestHandler,
	userController controllers.UserController,
) UserRoutes {
	return UserRoutes{
		handler:        handler,
		logger:         logger,
		userController: userController,
	}
}
