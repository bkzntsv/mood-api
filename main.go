package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	config := GetConfig()
	service := NewService(config)
	handler := NewHandler(service)

	router := setupRouter(handler)
	server := setupServer(router, config.Port)

	// Start server in goroutine
	go startServer(server)

	// Handle graceful shutdown
	gracefulShutdown(server)
}

func setupRouter(handler *Handler) *gin.Engine {
	// Set gin mode based on environment
	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	// Add basic middleware
	router.Use(gin.Recovery())
	router.Use(gin.Logger())

	// Group API routes
	api := router.Group("/api/v1")
	{
		api.GET("/health", handler.HealthCheck)
		api.POST("/analyze", handler.AnalyzeSentiment)
		api.POST("/model", handler.CallModelEndpoint)
	}

	return router
}

func setupServer(router *gin.Engine, port string) *http.Server {
	if port == "" {
		port = "8080"
	}

	return &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}
}

func startServer(srv *http.Server) {
	log.Printf("Starting server on port %s", srv.Addr[1:])
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func gracefulShutdown(srv *http.Server) {
	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	// Gracefully shutdown with 5 second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exited properly")
}
