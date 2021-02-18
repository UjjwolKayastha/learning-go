package middlewares

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewMiddlewares),
)

type IMiddleware interface {
	Setup()
}

type Middlewares []IMiddleware

func NewMiddlewares() Middlewares {
	return Middlewares{}
}

func (m Middlewares) Setup() {
	for _, middleware := range m {
		middleware.Setup()
	}
}
