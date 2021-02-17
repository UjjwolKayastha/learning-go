package main

import (
	"al.com/bundle"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		bundle.Module,
	).Run()
}
