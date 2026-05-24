package services

import (
	"github.com/ai-project-os/api/internal/database"
	"github.com/ai-project-os/api/internal/models"
)

type DashboardService struct{}

func NewDashboardService() *DashboardService {
	return &DashboardService{}
}

type DashboardStats struct {
	ProjectCount    int64 `json:"project_count"`
	PromptCount     int64 `json:"prompt_count"`
	GenerationCount int64 `json:"generation_count"`
}

func (s *DashboardService) GetStats(userID string) (*DashboardStats, error) {
	db := database.GetDB()
	var stats DashboardStats

	if err := db.Model(&models.Project{}).Where("user_id = ?", userID).Count(&stats.ProjectCount).Error; err != nil {
		return nil, err
	}
	if err := db.Model(&models.Prompt{}).Where("user_id = ?", userID).Count(&stats.PromptCount).Error; err != nil {
		return nil, err
	}
	if err := db.Model(&models.AIGeneration{}).Where("user_id = ?", userID).Count(&stats.GenerationCount).Error; err != nil {
		return nil, err
	}
	return &stats, nil
}
