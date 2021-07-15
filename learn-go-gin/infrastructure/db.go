package infrastructure

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Database modal
type Database struct {
	*gorm.DB
}

// NewDatabase creates a new database instance
func NewDatabase(env Env, zapLogger Logger) Database {

	username := env.DB_USERNAME
	password := env.DB_PASSWORD
	host := env.DB_HOST
	port := env.DB_PORT
	dbname := env.DB_NAME

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, dbname)

	if env.ENVIRONMENT != "local" {
		url = fmt.Sprintf(
			"%s:%s@unix(/cloudsql/%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			env.DB_USERNAME,
			env.DB_PASSWORD,
			env.DB_HOST,
			env.DB_NAME,
		)
	}

	db, err := gorm.Open(mysql.Open(url), &gorm.Config{
		Logger: newLogger,
	})
	_ = db.Exec("CREATE DATABASE IF NOT EXISTS " + env.DB_NAME + ";")

	if err != nil {
		zapLogger.Info("Url: ", url)
		zapLogger.Panic(err)
	}

	zapLogger.Info("Database connected.")

	return Database{
		DB: db,
	}
}
