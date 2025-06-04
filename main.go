package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/sunnysingha911/admin-service/config"
	"github.com/sunnysingha911/admin-service/grpc"
	"github.com/sunnysingha911/admin-service/handlers"
	"github.com/sunnysingha911/admin-service/routes/v1"
)

func main() {

	if err := config.LoadEnv(); err != nil {
		log.Fatalf("Failed to load env: %v", err)
	}

	app := fiber.New()

	// You can get this from env/config instead of hardcoding

	userClient := grpc.NewUserClient(config.GRPCUSER)

	defer userClient.Conn.Close()

	v1 := app.Group("/api/v1")
	userHandler := handlers.NewUserHandler(userClient)
	routes.RegisterUserRoutes(v1, userHandler)

	// Run server in goroutine so we can listen for shutdown signals
	go func() {
		if err := app.Listen(":5000"); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Graceful shutdown on Ctrl+C or SIGTERM
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")
	if err := app.Shutdown(); err != nil {
		log.Fatalf("Error during server shutdown: %v", err)
	}
}
