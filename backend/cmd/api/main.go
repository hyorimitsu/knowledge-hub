package main

import (
	"fmt"
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/hyorimitsu/knowledge-hub/backend/docs/openapi"
	"github.com/hyorimitsu/knowledge-hub/backend/internal/infrastructure/persistence"
)

// @title Knowledge Hub API
// @version 1.0
// @description Knowledge Hub API documentation
// @host localhost:8080
// @BasePath /api
func main() {
	// Initialize Echo
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Database connection
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// Initialize repositories
	repos := persistence.NewRepositories(db)
	_ = repos // Temporary to avoid unused variable error until we implement handlers

	// API routes
	api := e.Group("/api")
	{
		// Health check
		api.GET("/health", func(c echo.Context) error {
			return c.JSON(200, map[string]string{"status": "ok"})
		})
	}

	// Swagger
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	e.Logger.Fatal(e.Start(":" + port))
}
