package services

import (
	"github.com/ai-project-os/api/internal/models"
	"github.com/ai-project-os/api/internal/repositories"
)

type TemplateService struct {
	repo *repositories.TemplateRepository
}

func NewTemplateService() *TemplateService {
	return &TemplateService{repo: repositories.NewTemplateRepository()}
}

type TemplateListResponse struct {
	Templates []models.Template `json:"templates"`
	Total     int64             `json:"total"`
}

func (s *TemplateService) List(projectType string, page, pageSize int) (*TemplateListResponse, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}
	offset := (page - 1) * pageSize
	templates, total, err := s.repo.FindAll(projectType, pageSize, offset)
	if err != nil {
		return nil, err
	}
	return &TemplateListResponse{Templates: templates, Total: total}, nil
}

func (s *TemplateService) GetByID(id string) (*models.Template, error) {
	return s.repo.FindByID(id)
}
