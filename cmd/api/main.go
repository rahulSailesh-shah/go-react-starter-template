package main

import (
	"context"

	"github.com/joho/godotenv"
	"github.com/rahulSailesh-shah/converSense/internal/app"
	"github.com/rahulSailesh-shah/converSense/internal/server"
	"github.com/rahulSailesh-shah/converSense/pkg/config"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}

	config, err := config.LoadConfig()
	if  err != nil {
		panic("Failed to load config: " + err.Error())
	}

	ctx := context.Background()
	app, err := app.NewApp(ctx, config)
	if err != nil {
		panic("Failed to initialize app: " + err.Error())
	}

	server, err := server.NewServer(ctx, app)
	if err != nil {
		panic("Failed to initialize server: " + err.Error())
	}

	if err := server.Run(); err != nil {
		panic("Failed to run server: " + err.Error())
	}
}
