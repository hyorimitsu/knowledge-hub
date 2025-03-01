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
	appErrors "github.com/hyorimitsu/knowledge-hub/backend/internal/infrastructure/errors"
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
	
	// Configure custom logger
	logger := appErrors.NewLogger(appErrors.DefaultLogConfig)
	logger.SetEcho(e)
	
	// Register custom validator
	appErrors.RegisterEchoValidator(e)
	
	// Set custom HTTP error handler
	appErrors.RegisterErrorHandlers(e)

	// Middleware
	e.Use(middleware.RequestID())
	e.Use(middleware.Logger())
	e.Use(appErrors.RecoverWithConfig()) // Use our custom recover middleware
	e.Use(middleware.CORS())
	e.Use(appErrors.ErrorHandler()) // Use our custom error handler middleware
	e.Use(appErrors.ValidationMiddleware()) // Use our validation middleware

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
			return appErrors.SendOK(c, map[string]string{"status": "ok"})
		})

		// Test error handling
		api.GET("/test-error/:type", func(c echo.Context) error {
			errorType := c.Param("type")
			
			switch errorType {
			case "validation":
				fieldErrors := map[string]string{
					"username": "Username is required",
					"email":    "Email must be valid",
				}
				return appErrors.NewValidationError("Validation failed", fieldErrors, nil)
			case "not-found":
				return appErrors.NotFound("Resource not found", nil)
			case "unauthorized":
				return appErrors.Unauthorized("Authentication required", nil)
			case "forbidden":
				return appErrors.Forbidden("Access denied", nil)
			case "conflict":
				return appErrors.Conflict("Resource already exists", nil)
			case "internal":
				return appErrors.InternalServerError("Something went wrong", nil)
			case "domain":
				return appErrors.NewDomainError(appErrors.ErrTenantNotFound, "Tenant not found", nil)
			case "panic":
				panic("Test panic")
			default:
				return appErrors.SendOK(c, map[string]string{"message": "No error"})
			}
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
