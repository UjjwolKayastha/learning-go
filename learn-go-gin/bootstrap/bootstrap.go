package bootstrap

import (
	"context"
	"hotels-api/api/controllers"
	"hotels-api/api/repositories"
	"hotels-api/api/routes"
	"hotels-api/api/services"
	"hotels-api/infrastructure"

	"go.uber.org/fx"
)

// Module exported for initializing application
var Module = fx.Options(
	controllers.Module,
	infrastructure.Module,
	routes.Module,
	services.Module,
	repositories.Module,

	fx.Invoke(bootstrap),
)

func bootstrap(
	lifecycle fx.Lifecycle,
	env infrastructure.Env,
	routes routes.Routes,
	logger infrastructure.Logger,
	database infrastructure.Database,
	handler infrastructure.Router,
) {
	conn, _ := database.DB.DB()

	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			logger.Info("Starting Application")
			logger.Info("---------------------------")
			logger.Info("------- HOTELS-API 🏨 -------")
			logger.Info("---------------------------")
			conn.SetMaxOpenConns(10)

			go func() {
				routes.Setup()
				if env.SERVER_PORT == "" {
					handler.Run()
				} else {
					handler.Run(env.SERVER_PORT)
				}
			}()
			return nil
		},
		OnStop: func(context.Context) error {
			logger.Info("Stopping Application")
			conn.Close()
			return nil
		},
	})
}
