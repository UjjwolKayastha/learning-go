package bootstrap

import (
	"context"
	"github.com/ujjwolkayastha/learn-clean-gin/api/controllers"
	"github.com/ujjwolkayastha/learn-clean-gin/api/repositories"
	"github.com/ujjwolkayastha/learn-clean-gin/api/routes"
	"github.com/ujjwolkayastha/learn-clean-gin/api/services"
	"github.com/ujjwolkayastha/learn-clean-gin/lib"
	"go.uber.org/fx"
)

// Module exported for initializing application
var Module = fx.Options(
	controllers.Module,
	routes.Module,
	lib.Module,
	services.Module,
	repositories.Module,
	fx.Invoke(bootstrap),
)

func bootstrap(
	lifecycle fx.Lifecycle,
	handler lib.RequestHandler,
	routes routes.Routes,
	env lib.Env,
	logger lib.Logger,
	database lib.Database,
) {
	conn, _ := database.DB.DB()

	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			logger.Zap.Info("Starting Application...")
			conn.SetMaxOpenConns(10)
			go func() {
				routes.Setup()
				handler.Gin.Run(env.ServerPort)
			}()
			return nil
		},
		OnStop: func(context.Context) error {
			logger.Zap.Info("Stopping Application")
			conn.Close()
			return nil
		},
	})
}
