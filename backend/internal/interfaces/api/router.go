package api

import (
	"github.com/labstack/echo/v4"
	echojwt "github.com/labstack/echo-jwt/v4"

	"github.com/hyorimitsu/knowledge-hub/backend/internal/infrastructure/persistence"
	"github.com/hyorimitsu/knowledge-hub/backend/internal/interfaces/api/handlers"
	"github.com/hyorimitsu/knowledge-hub/backend/internal/interfaces/api/middleware"
	"github.com/hyorimitsu/knowledge-hub/backend/internal/usecase/comment"
	"github.com/hyorimitsu/knowledge-hub/backend/internal/usecase/knowledge"
	"github.com/hyorimitsu/knowledge-hub/backend/internal/usecase/tag"
	"github.com/hyorimitsu/knowledge-hub/backend/internal/usecase/tenant"
	"github.com/hyorimitsu/knowledge-hub/backend/internal/usecase/user"
)

// Router handles the API routes
type Router struct {
	e           *echo.Echo
	repositories *persistence.Repositories
}

// NewRouter creates a new router
func NewRouter(e *echo.Echo, repositories *persistence.Repositories) *Router {
	return &Router{
		e:           e,
		repositories: repositories,
	}
}

// SetupRoutes sets up the API routes
func (r *Router) SetupRoutes() {
	// API group
	api := r.e.Group("/api")

	// Set repositories in context
	api.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("repositories", r.repositories)
			return next(c)
		}
	})

	// Public routes
	r.setupPublicRoutes(api)

	// Protected routes
	r.setupProtectedRoutes(api)
}

// setupPublicRoutes sets up the public API routes
func (r *Router) setupPublicRoutes(api *echo.Group) {
	// Health check
	api.GET("/health", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"status": "ok"})
	})

	// Auth handler (public routes)
	authHandler := handlers.NewAuthHandler(
		user.NewAuthenticateUserUseCase(r.repositories.User(), r.repositories.Tenant()),
		user.NewRegisterUserUseCase(r.repositories.User(), r.repositories.Tenant()),
	)
	authHandler.RegisterRoutes(api)

	// Tenant handler (public endpoints)
	tenantHandler := handlers.NewTenantHandler(
		tenant.NewCreateTenantUseCase(r.repositories.Tenant()),
		tenant.NewUpdateTenantSettingsUseCase(r.repositories.Tenant()),
		tenant.NewDeleteTenantUseCase(r.repositories.Tenant()),
	)
	api.GET("/tenants/domain/:domain", tenantHandler.GetByDomain)
}

// setupProtectedRoutes sets up the protected API routes
func (r *Router) setupProtectedRoutes(api *echo.Group) {
	// JWT middleware
	jwtMiddleware := echojwt.WithConfig(middleware.JWTConfig())
	protected := api.Group("", jwtMiddleware)

	// Auth handler (protected routes)
	authHandler := handlers.NewAuthHandler(
		user.NewAuthenticateUserUseCase(r.repositories.User(), r.repositories.Tenant()),
		user.NewRegisterUserUseCase(r.repositories.User(), r.repositories.Tenant()),
	)
	authHandler.RegisterProtectedRoutes(protected)

	// Tenant handler (protected endpoints)
	tenantHandler := handlers.NewTenantHandler(
		tenant.NewCreateTenantUseCase(r.repositories.Tenant()),
		tenant.NewUpdateTenantSettingsUseCase(r.repositories.Tenant()),
		tenant.NewDeleteTenantUseCase(r.repositories.Tenant()),
	)
	tenantGroup := protected.Group("/tenants")
	tenantGroup.POST("", tenantHandler.Create, middleware.RoleMiddleware("admin"))
	tenantGroup.GET("/:id", tenantHandler.Get)
	tenantGroup.PUT("/:id/settings", tenantHandler.UpdateSettings, middleware.RoleMiddleware("admin"))
	tenantGroup.DELETE("/:id", tenantHandler.Delete, middleware.RoleMiddleware("admin"))

	// Knowledge handler
	knowledgeHandler := handlers.NewKnowledgeHandler(
		knowledge.NewCreateKnowledgeUseCase(r.repositories.Knowledge(), r.repositories.User(), r.repositories.Tag(), r.repositories.Tenant()),
		knowledge.NewUpdateKnowledgeUseCase(r.repositories.Knowledge(), r.repositories.Tag(), r.repositories.Tenant()),
		knowledge.NewDeleteKnowledgeUseCase(r.repositories.Knowledge(), r.repositories.Tenant()),
		knowledge.NewSearchKnowledgeUseCase(r.repositories.Knowledge(), r.repositories.Tag(), r.repositories.Tenant()),
		r.repositories.Knowledge(),
	)
	knowledgeHandler.RegisterRoutes(protected)

	// Tag handler
	tagHandler := handlers.NewTagHandler(
		tag.NewCreateTagUseCase(r.repositories.Tag(), r.repositories.Tenant()),
		tag.NewUpdateTagUseCase(r.repositories.Tag(), r.repositories.Tenant()),
		tag.NewDeleteTagUseCase(r.repositories.Tag(), r.repositories.Tenant()),
	)
	tagHandler.RegisterRoutes(protected)

	// Comment handler
	commentHandler := handlers.NewCommentHandler(
		comment.NewCreateCommentUseCase(r.repositories.Comment(), r.repositories.Knowledge(), r.repositories.User(), r.repositories.Tenant()),
		comment.NewUpdateCommentUseCase(r.repositories.Comment(), r.repositories.Tenant()),
		comment.NewDeleteCommentUseCase(r.repositories.Comment(), r.repositories.Tenant()),
	)
	commentHandler.RegisterRoutes(protected)
}