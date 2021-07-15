package infrastructure

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	LOG_OUTPUT  string `mapstructure:"LOG_OUTPUT"`
	SERVER_PORT string `mapstructure:"SERVER_PORT"`
	ENVIRONMENT string `mapstructure:"ENVIRONMENT"`
	DB_USERNAME string `mapstructure:"DB_USERNAME"`
	DB_PASSWORD string `mapstructure:"DB_PASSWORD"`
	DB_HOST     string `mapstructure:"DB_HOST"`
	DB_PORT     string `mapstructure:"DB_PORT"`
	DB_NAME     string `mapstructure:"DB_NAME"`
}

// NewEnv creates a new environment
func NewEnv() Env {
	env := Env{}
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("cannot read cofiguration")
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("environment cant be loaded: ", err)
	}

	log.Printf("%#v \n", env)
	return env
}
