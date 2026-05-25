package main

import (
	"fmt"
	"log"

	"github.com/ai-project-os/api/internal/ai"
	"github.com/ai-project-os/api/internal/config"
	"github.com/ai-project-os/api/internal/database"
	"github.com/ai-project-os/api/internal/handlers"
	"github.com/ai-project-os/api/internal/logger"
	"github.com/ai-project-os/api/internal/middleware"
	"github.com/ai-project-os/api/internal/routes"
	"github.com/ai-project-os/api/internal/services"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Load config
	cfg := config.Load()

	// Init logger
	logger.Init(cfg.LogLevel)

	// Init database
	if err := database.Init(cfg); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Init AI provider
	var provider ai.Provider
	switch cfg.AIProvider {
	case "mock", "":
		provider = ai.NewMockProvider(cfg.AIModel)
	default:
		// All OpenAI-compatible providers (deepseek, openai, groq, moonshot,
		// siliconflow, zhipu, together, ollama, mistral, perplexity, custom...)
		provider = ai.NewOpenAICompatProvider(cfg.AIProvider, cfg.AIBaseURL, cfg.AIAPIKey, cfg.AIModel)
	}

	// Init services
	authSvc := services.NewAuthService(cfg)
	projectSvc := services.NewProjectService()
	promptSvc := services.NewPromptService()
	projectDocSvc := services.NewProjectDocService()
	templateSvc := services.NewTemplateService()
	aiGenSvc := services.NewAIGenerationService(provider, cfg)
	dashboardSvc := services.NewDashboardService()
	adminSvc := services.NewAdminService()

	// Init handlers
	h := &routes.Handlers{
		Health:       handlers.NewHealthHandler(),
		Auth:         handlers.NewAuthHandler(authSvc),
		Project:      handlers.NewProjectHandler(projectSvc),
		Prompt:       handlers.NewPromptHandler(promptSvc),
		ProjectDoc:   handlers.NewProjectDocHandler(projectDocSvc),
		Template:     handlers.NewTemplateHandler(templateSvc),
		AIGeneration: handlers.NewAIGenerationHandler(aiGenSvc),
		Dashboard:    handlers.NewDashboardHandler(dashboardSvc),
		Admin:        handlers.NewAdminHandler(adminSvc),
	}

	// Init Fiber app
	app := fiber.New(fiber.Config{
		AppName: "AI Project OS API",
	})

	// Register global middleware
	middleware.RegisterGlobal(app)

	// Register routes
	routes.Register(app, cfg, h)

	// Start server
	addr := fmt.Sprintf(":%s", cfg.ServerPort)
	fmt.Printf("🚀 AI Project OS API server starting on %s\n", addr)
	if err := app.Listen(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
