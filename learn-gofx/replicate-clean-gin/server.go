package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/ujjwolkayastha/learn-clean-gin/bootstrap"
	"go.uber.org/fx"
)

func main() {
	godotenv.Load()
	fx.New(bootstrap.Module).Run()
}
