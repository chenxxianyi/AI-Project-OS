package services

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/ai-project-os/api/internal/ai"
	"github.com/ai-project-os/api/internal/config"
	"github.com/ai-project-os/api/internal/models"
	"github.com/ai-project-os/api/internal/repositories"
	"go.uber.org/zap"
)

type AIGenerationService struct {
	provider    ai.Provider
	repo        *repositories.AIGenerationRepository
	projectRepo *repositories.ProjectRepository
	cfg         *config.Config
}

func NewAIGenerationService(provider ai.Provider, cfg *config.Config) *AIGenerationService {
	return &AIGenerationService{
		provider:    provider,
		repo:        repositories.NewAIGenerationRepository(),
		projectRepo: repositories.NewProjectRepository(),
		cfg:         cfg,
	}
}

type GeneratePromptRequest struct {
	ProjectID      string `json:"project_id"`
	GenerationType string `json:"generation_type"`
	InputPayload   string `json:"input_payload"`
}

func (s *AIGenerationService) Generate(userID string, req GeneratePromptRequest) (*models.AIGeneration, error) {
	if req.ProjectID == "" || req.GenerationType == "" || req.InputPayload == "" {
		return nil, errors.New("project_id, generation_type and input_payload are required")
	}
	if _, err := s.projectRepo.FindByIDAndUserID(req.ProjectID, userID); err != nil {
		return nil, errors.New("project not found")
	}

	aiReq := ai.GenerateRequest{
		ProjectID:      req.ProjectID,
		GenerationType: req.GenerationType,
		InputPayload:   req.InputPayload,
	}

	result, err := s.provider.Generate(context.Background(), aiReq)
	if err != nil {
		return nil, err
	}

	inputJSON, _ := json.Marshal(req)

	gen := &models.AIGeneration{
		UserID:         userID,
		ProjectID:      req.ProjectID,
		GenerationType: models.GenerationType(req.GenerationType),
		InputPayload:   string(inputJSON),
		OutputContent:  result.Content,
		ModelProvider:  result.Provider,
		ModelName:      result.Model,
		Status:         models.GenerationStatusSuccess,
		CreatedAt:      time.Now().Unix(),
	}
	if err := s.repo.Create(gen); err != nil {
		zap.L().Error("failed to save ai generation record", zap.Error(err))
	}

	zap.L().Info("ai generation completed",
		zap.String("type", req.GenerationType),
		zap.String("provider", result.Provider),
	)
	return gen, nil
}

func (s *AIGenerationService) ListByUser(userID string, page, pageSize int) ([]models.AIGeneration, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}
	offset := (page - 1) * pageSize
	return s.repo.FindByUserID(userID, pageSize, offset)
}
