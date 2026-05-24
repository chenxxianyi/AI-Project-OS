package services

import (
	"errors"

	"github.com/ai-project-os/api/internal/models"
	"github.com/ai-project-os/api/internal/repositories"
	"go.uber.org/zap"
)

type PromptService struct {
	repo        *repositories.PromptRepository
	projectRepo *repositories.ProjectRepository
}

func NewPromptService() *PromptService {
	return &PromptService{
		repo:        repositories.NewPromptRepository(),
		projectRepo: repositories.NewProjectRepository(),
	}
}

type CreatePromptRequest struct {
	Type    string `json:"type"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type UpdatePromptRequest struct {
	Type    *string `json:"type,omitempty"`
	Title   *string `json:"title,omitempty"`
	Content *string `json:"content,omitempty"`
}

func (s *PromptService) Create(projectID, userID string, req CreatePromptRequest) (*models.Prompt, error) {
	if req.Title == "" || req.Content == "" {
		return nil, errors.New("title and content are required")
	}
	if _, err := s.projectRepo.FindByIDAndUserID(projectID, userID); err != nil {
		return nil, errors.New("project not found")
	}
	p := &models.Prompt{
		ProjectID: projectID,
		UserID:    userID,
		Type:      models.PromptType(req.Type),
		Title:     req.Title,
		Content:   req.Content,
		Version:   1,
	}
	if err := s.repo.Create(p); err != nil {
		return nil, err
	}
	// create initial version
	v := &models.PromptVersion{
		PromptID:   p.ID,
		Version:    1,
		Content:    req.Content,
		ChangeNote: "initial version",
	}
	if err := s.repo.CreateVersion(v); err != nil {
		zap.L().Warn("failed to create initial prompt version", zap.Error(err))
	}
	zap.L().Info("prompt created", zap.String("id", p.ID))
	return p, nil
}

func (s *PromptService) ListByProject(projectID string, promptType string) ([]models.Prompt, error) {
	return s.repo.FindByProjectID(projectID, promptType)
}

func (s *PromptService) GetByID(id string) (*models.Prompt, error) {
	return s.repo.FindByID(id)
}

func (s *PromptService) Update(id string, req UpdatePromptRequest) (*models.Prompt, error) {
	p, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	contentChanged := false
	if req.Type != nil {
		p.Type = models.PromptType(*req.Type)
	}
	if req.Title != nil {
		p.Title = *req.Title
	}
	if req.Content != nil && *req.Content != p.Content {
		p.Content = *req.Content
		p.Version++
		contentChanged = true
	}
	if err := s.repo.Update(p); err != nil {
		return nil, err
	}
	if contentChanged {
		v := &models.PromptVersion{
			PromptID: p.ID,
			Version:  p.Version,
			Content:  p.Content,
		}
		if err := s.repo.CreateVersion(v); err != nil {
			zap.L().Warn("failed to create prompt version", zap.Error(err))
		}
	}
	return p, nil
}

func (s *PromptService) Delete(id string) error {
	return s.repo.Delete(id)
}

func (s *PromptService) ListVersions(promptID string) ([]models.PromptVersion, error) {
	return s.repo.FindVersions(promptID)
}

func (s *PromptService) CountByUserID(userID string) (int64, error) {
	return s.repo.CountByUserID(userID)
}
