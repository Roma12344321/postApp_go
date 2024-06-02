package main

import (
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"log"
	"postApp"
	"postApp/pkg/handler"
	"postApp/pkg/repository"
	"postApp/pkg/service"
)

func main() {
	initConfig()
	db := initDb()
	runServer(db)
}

func initConfig() {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("config error")
	}
}

func initDb() *sqlx.DB {
	db, err := repository.NewPostgresDB(&repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: viper.GetString("db.password"),
	})
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}
	return db
}

func runServer(db *sqlx.DB) {
	server := new(postApp.Server)
	repositories := repository.NewRepository(db)
	services := service.NewService(repositories)
	handlers := handler.NewHandler(services)
	if err := server.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error")
	}
}
