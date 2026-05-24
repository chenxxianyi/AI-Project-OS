package services

import (
	"errors"

	"github.com/ai-project-os/api/internal/models"
	"github.com/ai-project-os/api/internal/repositories"
	"go.uber.org/zap"
)

type ProjectService struct {
	repo *repositories.ProjectRepository
}

func NewProjectService() *ProjectService {
	return &ProjectService{repo: repositories.NewProjectRepository()}
}

type CreateProjectRequest struct {
	Name          string `json:"name"`
	Description   string `json:"description,omitempty"`
	ProjectType   string `json:"project_type"`
	FrontendStack string `json:"frontend_stack,omitempty"`
	BackendStack  string `json:"backend_stack,omitempty"`
	DatabaseStack string `json:"database_stack,omitempty"`
	AIStack       string `json:"ai_stack,omitempty"`
	UIStyle       string `json:"ui_style,omitempty"`
	TargetUser    string `json:"target_user,omitempty"`
	ProductGoal   string `json:"product_goal,omitempty"`
}

type UpdateProjectRequest struct {
	Name          *string `json:"name,omitempty"`
	Description   *string `json:"description,omitempty"`
	FrontendStack *string `json:"frontend_stack,omitempty"`
	BackendStack  *string `json:"backend_stack,omitempty"`
	DatabaseStack *string `json:"database_stack,omitempty"`
	AIStack       *string `json:"ai_stack,omitempty"`
	UIStyle       *string `json:"ui_style,omitempty"`
	TargetUser    *string `json:"target_user,omitempty"`
	ProductGoal   *string `json:"product_goal,omitempty"`
	Status        *string `json:"status,omitempty"`
}

type ProjectListResponse struct {
	Projects []models.Project `json:"projects"`
	Total    int64            `json:"total"`
}

func (s *ProjectService) Create(userID string, req CreateProjectRequest) (*models.Project, error) {
	if req.Name == "" || req.ProjectType == "" {
		return nil, errors.New("name and project_type are required")
	}
	p := &models.Project{
		UserID:        userID,
		Name:          req.Name,
		Description:   req.Description,
		ProjectType:   models.ProjectType(req.ProjectType),
		FrontendStack: req.FrontendStack,
		BackendStack:  req.BackendStack,
		DatabaseStack: req.DatabaseStack,
		AIStack:       req.AIStack,
		UIStyle:       req.UIStyle,
		TargetUser:    req.TargetUser,
		ProductGoal:   req.ProductGoal,
		Status:        models.ProjectStatusActive,
	}
	if err := s.repo.Create(p); err != nil {
		return nil, err
	}
	zap.L().Info("project created", zap.String("id", p.ID), zap.String("name", p.Name))
	return p, nil
}

func (s *ProjectService) List(userID string, projectType, search string, page, pageSize int) (*ProjectListResponse, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}
	offset := (page - 1) * pageSize
	projects, total, err := s.repo.FindByUserID(userID, projectType, search, pageSize, offset)
	if err != nil {
		return nil, err
	}
	return &ProjectListResponse{Projects: projects, Total: total}, nil
}

func (s *ProjectService) GetByID(id string) (*models.Project, error) {
	return s.repo.FindByID(id)
}

func (s *ProjectService) Update(id string, req UpdateProjectRequest) (*models.Project, error) {
	p, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if req.Name != nil {
		p.Name = *req.Name
	}
	if req.Description != nil {
		p.Description = *req.Description
	}
	if req.FrontendStack != nil {
		p.FrontendStack = *req.FrontendStack
	}
	if req.BackendStack != nil {
		p.BackendStack = *req.BackendStack
	}
	if req.DatabaseStack != nil {
		p.DatabaseStack = *req.DatabaseStack
	}
	if req.AIStack != nil {
		p.AIStack = *req.AIStack
	}
	if req.UIStyle != nil {
		p.UIStyle = *req.UIStyle
	}
	if req.TargetUser != nil {
		p.TargetUser = *req.TargetUser
	}
	if req.ProductGoal != nil {
		p.ProductGoal = *req.ProductGoal
	}
	if req.Status != nil {
		p.Status = models.ProjectStatus(*req.Status)
	}
	if err := s.repo.Update(p); err != nil {
		return nil, err
	}
	return p, nil
}

func (s *ProjectService) Delete(id string) error {
	return s.repo.Delete(id)
}
