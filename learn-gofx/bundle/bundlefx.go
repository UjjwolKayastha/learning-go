package bundle

import (
	"context"
	"fmt"

	"al.com/handler"
	"al.com/health"
	"go.uber.org/fx"
)

// Module exports for fx
var Module = fx.Options(
	handler.Module,
	health.Module,
	fx.Invoke(registerHooks),
)

func registerHooks(lifecycle fx.Lifecycle, h *handler.Handler) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				fmt.Println("Application is running in port :5000")
				go h.Gin.Run(":5000")
				return nil
			},
			OnStop: func(context.Context) error {
				fmt.Println("Stopping application")
				return nil
			},
		},
	)
}
