package services

import (
	"errors"

	"github.com/ai-project-os/api/internal/models"
	"github.com/ai-project-os/api/internal/repositories"
)

type ProjectDocService struct {
	repo *repositories.ProjectDocRepository
}

func NewProjectDocService() *ProjectDocService {
	return &ProjectDocService{repo: repositories.NewProjectDocRepository()}
}

type CreateDocRequest struct {
	DocType string `json:"doc_type"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type UpdateDocRequest struct {
	DocType *string `json:"doc_type,omitempty"`
	Title   *string `json:"title,omitempty"`
	Content *string `json:"content,omitempty"`
}

func (s *ProjectDocService) Create(projectID, userID string, req CreateDocRequest) (*models.ProjectDoc, error) {
	if req.Title == "" || req.Content == "" {
		return nil, errors.New("title and content are required")
	}
	doc := &models.ProjectDoc{
		ProjectID: projectID,
		UserID:    userID,
		DocType:   models.DocType(req.DocType),
		Title:     req.Title,
		Content:   req.Content,
	}
	if err := s.repo.Create(doc); err != nil {
		return nil, err
	}
	return doc, nil
}

func (s *ProjectDocService) ListByProject(projectID string, docType string) ([]models.ProjectDoc, error) {
	return s.repo.FindByProjectID(projectID, docType)
}

func (s *ProjectDocService) GetByID(id string) (*models.ProjectDoc, error) {
	return s.repo.FindByID(id)
}

func (s *ProjectDocService) Update(id string, req UpdateDocRequest) (*models.ProjectDoc, error) {
	doc, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if req.DocType != nil {
		doc.DocType = models.DocType(*req.DocType)
	}
	if req.Title != nil {
		doc.Title = *req.Title
	}
	if req.Content != nil {
		doc.Content = *req.Content
	}
	if err := s.repo.Update(doc); err != nil {
		return nil, err
	}
	return doc, nil
}

func (s *ProjectDocService) Delete(id string) error {
	return s.repo.Delete(id)
}
