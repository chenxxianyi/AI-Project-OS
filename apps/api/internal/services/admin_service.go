package services

import (
	"errors"
	"fmt"
	"runtime"
	"sync"
	"time"

	"github.com/ai-project-os/api/internal/database"
	"github.com/ai-project-os/api/internal/models"
	"github.com/ai-project-os/api/internal/repositories"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// ─── DTOs ────────────────────────────────────────────────────────────────────

type AdminSystemStats struct {
	TotalUsers       int64  `json:"total_users"`
	ActiveUsers      int64  `json:"active_users"`
	TotalProjects    int64  `json:"total_projects"`
	TotalPrompts     int64  `json:"total_prompts"`
	TotalGenerations int64  `json:"total_generations"`
	TotalTokensUsed  int64  `json:"total_tokens_used"`
	FailedGenerations int64 `json:"failed_generations"`
	Uptime           string `json:"uptime"`
	Version          string `json:"version"`
}

type AdminSystemHealth struct {
	Database    string `json:"database"`
	Redis       string `json:"redis"`
	AIProvider  string `json:"ai_provider"`
	DiskUsage   string `json:"disk_usage"`
	MemoryUsage string `json:"memory_usage"`
}

type AdminSystemSettings struct {
	AllowRegistration    bool   `json:"allow_registration"`
	DefaultUserRole      string `json:"default_user_role"`
	MaxProjectsPerUser   int    `json:"max_projects_per_user"`
	AIRateLimit          int    `json:"ai_rate_limit"`
	MaintenanceMode      bool   `json:"maintenance_mode"`
}

type AdminUserItem struct {
	ID              string `json:"id"`
	Email           string `json:"email"`
	Username        string `json:"username"`
	Avatar          string `json:"avatar,omitempty"`
	Role            string `json:"role"`
	Status          string `json:"status"`
	ProjectCount    int64  `json:"project_count"`
	GenerationCount int64  `json:"generation_count"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
}

type AdminUserListResponse struct {
	Users []AdminUserItem `json:"users"`
	Total int64           `json:"total"`
}

type AdminProjectItem struct {
	ID            string `json:"id"`
	UserID        string `json:"user_id"`
	Name          string `json:"name"`
	Description   string `json:"description,omitempty"`
	ProjectType   string `json:"project_type"`
	FrontendStack string `json:"frontend_stack,omitempty"`
	BackendStack  string `json:"backend_stack,omitempty"`
	Status        string `json:"status"`
	OwnerName     string `json:"owner_name"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}

type AdminProjectListResponse struct {
	Projects []AdminProjectItem `json:"projects"`
	Total    int64              `json:"total"`
}

type AdminGenerationItem struct {
	ID             string `json:"id"`
	UserID         string `json:"user_id"`
	ProjectID      string `json:"project_id"`
	GenerationType string `json:"generation_type"`
	ModelProvider  string `json:"model_provider"`
	ModelName      string `json:"model_name"`
	Status         string `json:"status"`
	ErrorMessage   string `json:"error_message,omitempty"`
	UserEmail      string `json:"user_email"`
	ProjectName    string `json:"project_name"`
	CreatedAt      string `json:"created_at"`
}

type AdminGenerationListResponse struct {
	Generations []AdminGenerationItem `json:"generations"`
	Total       int64                 `json:"total"`
}

type UpdateUserRequest struct {
	Role   *string `json:"role,omitempty"`
	Status *string `json:"status,omitempty"`
}

// ─── Settings store (in-memory for MVP) ──────────────────────────────────────

var (
	settingsMu   sync.RWMutex
	globalSettings = AdminSystemSettings{
		AllowRegistration:  true,
		DefaultUserRole:    "user",
		MaxProjectsPerUser: 10,
		AIRateLimit:        100,
		MaintenanceMode:    false,
	}
	startTime = time.Now()
)

// ─── Service ─────────────────────────────────────────────────────────────────

type AdminService struct {
	userRepo    *repositories.UserRepository
	projectRepo *repositories.ProjectRepository
	genRepo     *repositories.AIGenerationRepository
}

func NewAdminService() *AdminService {
	return &AdminService{
		userRepo:    repositories.NewUserRepository(),
		projectRepo: repositories.NewProjectRepository(),
		genRepo:     repositories.NewAIGenerationRepository(),
	}
}

func (s *AdminService) GetSystemStats() (*AdminSystemStats, error) {
	db := database.GetDB()
	stats := &AdminSystemStats{
		Uptime:  formatDuration(time.Since(startTime)),
		Version: "1.0.0",
	}

	if err := db.Model(&models.User{}).Count(&stats.TotalUsers).Error; err != nil {
		return nil, err
	}
	if err := db.Model(&models.User{}).Where("status = ?", models.UserStatusActive).Count(&stats.ActiveUsers).Error; err != nil {
		return nil, err
	}
	if err := db.Model(&models.Project{}).Count(&stats.TotalProjects).Error; err != nil {
		return nil, err
	}
	if err := db.Model(&models.Prompt{}).Count(&stats.TotalPrompts).Error; err != nil {
		return nil, err
	}
	if err := db.Model(&models.AIGeneration{}).Count(&stats.TotalGenerations).Error; err != nil {
		return nil, err
	}
	if err := db.Model(&models.AIGeneration{}).Where("status = ?", models.GenerationStatusFailed).Count(&stats.FailedGenerations).Error; err != nil {
		return nil, err
	}
	return stats, nil
}

func (s *AdminService) GetSystemHealth() *AdminSystemHealth {
	health := &AdminSystemHealth{
		Database:   "unknown",
		Redis:      "healthy",
		AIProvider: "healthy",
	}

	db := database.GetDB()
	sqlDB, err := db.DB()
	if err == nil && sqlDB.Ping() == nil {
		health.Database = "healthy"
	} else {
		health.Database = "down"
	}

	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	usedMB := memStats.Alloc / 1024 / 1024
	health.MemoryUsage = fmt.Sprintf("%dMB", usedMB)
	health.DiskUsage = "N/A"

	return health
}

func (s *AdminService) GetSettings() *AdminSystemSettings {
	settingsMu.RLock()
	defer settingsMu.RUnlock()
	cp := globalSettings
	return &cp
}

func (s *AdminService) UpdateSettings(req AdminSystemSettings) *AdminSystemSettings {
	settingsMu.Lock()
	defer settingsMu.Unlock()
	globalSettings = req
	zap.L().Info("admin settings updated")
	return &req
}

// ─── Users ────────────────────────────────────────────────────────────────────

func (s *AdminService) ListUsers(search, role string, page, pageSize int) (*AdminUserListResponse, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}
	offset := (page - 1) * pageSize
	rows, total, err := s.userRepo.ListAll(search, role, pageSize, offset)
	if err != nil {
		return nil, err
	}

	items := make([]AdminUserItem, len(rows))
	for i, r := range rows {
		items[i] = AdminUserItem{
			ID:              r.User.ID,
			Email:           r.User.Email,
			Username:        r.User.Username,
			Avatar:          r.User.Avatar,
			Role:            string(r.User.Role),
			Status:          string(r.User.Status),
			ProjectCount:    r.ProjectCount,
			GenerationCount: r.GenerationCount,
			CreatedAt:       r.User.CreatedAt.Format(time.RFC3339),
			UpdatedAt:       r.User.UpdatedAt.Format(time.RFC3339),
		}
	}
	return &AdminUserListResponse{Users: items, Total: total}, nil
}

func (s *AdminService) UpdateUser(id string, req UpdateUserRequest) (*AdminUserItem, error) {
	user, err := s.userRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	if req.Role != nil {
		user.Role = models.UserRole(*req.Role)
	}
	if req.Status != nil {
		user.Status = models.UserStatus(*req.Status)
	}
	if err := s.userRepo.Update(user); err != nil {
		return nil, err
	}
	zap.L().Info("admin updated user", zap.String("id", id))
	item := &AdminUserItem{
		ID:       user.ID,
		Email:    user.Email,
		Username: user.Username,
		Role:     string(user.Role),
		Status:   string(user.Status),
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
		UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
	}
	return item, nil
}

func (s *AdminService) DeleteUser(id string) error {
	if _, err := s.userRepo.FindByID(id); err != nil {
		return errors.New("user not found")
	}
	zap.L().Info("admin deleted user", zap.String("id", id))
	return s.userRepo.Delete(id)
}

// ─── Projects ─────────────────────────────────────────────────────────────────

func (s *AdminService) ListProjects(search, status string, page, pageSize int) (*AdminProjectListResponse, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}
	offset := (page - 1) * pageSize
	rows, total, err := s.projectRepo.ListAll(search, status, pageSize, offset)
	if err != nil {
		return nil, err
	}

	items := make([]AdminProjectItem, len(rows))
	for i, r := range rows {
		items[i] = AdminProjectItem{
			ID:            r.Project.ID,
			UserID:        r.Project.UserID,
			Name:          r.Project.Name,
			Description:   r.Project.Description,
			ProjectType:   string(r.Project.ProjectType),
			FrontendStack: r.Project.FrontendStack,
			BackendStack:  r.Project.BackendStack,
			Status:        string(r.Project.Status),
			OwnerName:     r.OwnerName,
			CreatedAt:     r.Project.CreatedAt.Format(time.RFC3339),
			UpdatedAt:     r.Project.UpdatedAt.Format(time.RFC3339),
		}
	}
	return &AdminProjectListResponse{Projects: items, Total: total}, nil
}

func (s *AdminService) DeleteProject(id string) error {
	if _, err := s.projectRepo.FindByID(id); err != nil {
		return errors.New("project not found")
	}
	zap.L().Info("admin deleted project", zap.String("id", id))
	return s.projectRepo.Delete(id)
}

// ─── Generations ──────────────────────────────────────────────────────────────

func (s *AdminService) ListGenerations(status, genType string, page, pageSize int) (*AdminGenerationListResponse, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}
	offset := (page - 1) * pageSize
	rows, total, err := s.genRepo.ListAll(status, genType, pageSize, offset)
	if err != nil {
		return nil, err
	}

	items := make([]AdminGenerationItem, len(rows))
	for i, r := range rows {
		items[i] = AdminGenerationItem{
			ID:             r.AIGeneration.ID,
			UserID:         r.AIGeneration.UserID,
			ProjectID:      r.AIGeneration.ProjectID,
			GenerationType: string(r.AIGeneration.GenerationType),
			ModelProvider:  r.AIGeneration.ModelProvider,
			ModelName:      r.AIGeneration.ModelName,
			Status:         string(r.AIGeneration.Status),
			ErrorMessage:   r.AIGeneration.ErrorMessage,
			UserEmail:      r.UserEmail,
			ProjectName:    r.ProjectName,
			CreatedAt:      time.Unix(r.AIGeneration.CreatedAt, 0).Format(time.RFC3339),
		}
	}
	return &AdminGenerationListResponse{Generations: items, Total: total}, nil
}

// ─── Helpers ──────────────────────────────────────────────────────────────────

func formatDuration(d time.Duration) string {
	days := int(d.Hours()) / 24
	hours := int(d.Hours()) % 24
	minutes := int(d.Minutes()) % 60
	if days > 0 {
		return fmt.Sprintf("%dd %dh %dm", days, hours, minutes)
	}
	if hours > 0 {
		return fmt.Sprintf("%dh %dm", hours, minutes)
	}
	return fmt.Sprintf("%dm", minutes)
}
