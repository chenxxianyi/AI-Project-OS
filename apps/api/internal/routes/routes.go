package routes

import (
	"github.com/ai-project-os/api/internal/config"
	"github.com/ai-project-os/api/internal/handlers"
	"github.com/ai-project-os/api/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func Register(app *fiber.App, cfg *config.Config, handlers *Handlers) {
	// Public routes
	app.Get("/health", handlers.Health.Check)

	api := app.Group("/api/v1")

	// Auth (public)
	auth := api.Group("/auth")
	auth.Post("/register", handlers.Auth.Register)
	auth.Post("/login", handlers.Auth.Login)
	auth.Post("/logout", handlers.Auth.Logout)

	// Protected routes
	protected := api.Use(middleware.JWTAuth(cfg))

	// User
	protected.Get("/auth/me", handlers.Auth.Me)

	// Projects
	projects := protected.Group("/projects")
	projects.Post("/", handlers.Project.Create)
	projects.Get("/", handlers.Project.List)
	projects.Get("/:id", handlers.Project.Get)
	projects.Put("/:id", handlers.Project.Update)
	projects.Delete("/:id", handlers.Project.Delete)

	// Prompts (nested under projects)
	prompts := projects.Group("/:projectId/prompts")
	prompts.Post("/", handlers.Prompt.Create)
	prompts.Get("/", handlers.Prompt.List)
	prompts.Get("/:id", handlers.Prompt.Get)
	prompts.Put("/:id", handlers.Prompt.Update)
	prompts.Delete("/:id", handlers.Prompt.Delete)
	prompts.Get("/:id/versions", handlers.Prompt.ListVersions)

	// Project Brain (docs)
	docs := projects.Group("/:projectId/docs")
	docs.Post("/", handlers.ProjectDoc.Create)
	docs.Get("/", handlers.ProjectDoc.List)
	docs.Get("/:docId", handlers.ProjectDoc.Get)
	docs.Put("/:docId", handlers.ProjectDoc.Update)
	docs.Delete("/:docId", handlers.ProjectDoc.Delete)

	// Templates (public read)
	templates := api.Group("/templates")
	templates.Get("/", handlers.Template.List)
	templates.Get("/:id", handlers.Template.Get)

	// AI Generation
	generations := protected.Group("/generations")
	generations.Post("/", handlers.AIGeneration.Generate)
	generations.Get("/", handlers.AIGeneration.List)

	// Dashboard
	protected.Get("/dashboard/stats", handlers.Dashboard.Stats)
}

// Handlers holds all handler instances
type Handlers struct {
	Health       *handlers.HealthHandler
	Auth         *handlers.AuthHandler
	Project      *handlers.ProjectHandler
	Prompt       *handlers.PromptHandler
	ProjectDoc   *handlers.ProjectDocHandler
	Template     *handlers.TemplateHandler
	AIGeneration *handlers.AIGenerationHandler
	Dashboard    *handlers.DashboardHandler
}
