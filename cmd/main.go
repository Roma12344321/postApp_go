package main

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
	"postApp"
	"postApp/pkg/handler"
	"postApp/pkg/repository"
	"postApp/pkg/service"
	"syscall"
)

func main() {
	sigChan := make(chan os.Signal, 1)
	ctx, cancel := context.WithCancel(context.Background())
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigChan
		cancel()
		fmt.Println("Received shutdown signal, shutting down...")
	}()
	initConfig()
	db := initDb()
	runServer(ctx, db)
}

func initConfig() {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
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
		log.Fatalf("Failed to initialize db: %s", err)
	}
	return db
}

func runServer(ctx context.Context, db *sqlx.DB) {
	server := new(postApp.Server)
	repositories := repository.NewRepository(db)
	services := service.NewService(repositories)
	handlers := handler.NewHandler(services)
	go func() {
		if err := server.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			log.Fatalf("Failed to run server: %s", err)
		}
	}()
	<-ctx.Done()
	fmt.Println("Shutting down server...")
	if err := server.ShutDown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}
	fmt.Println("Server exited properly")
}
