package repositories

import (
	"github.com/ai-project-os/api/internal/database"
	"github.com/ai-project-os/api/internal/models"
	"gorm.io/gorm"
)

type AIGenerationRepository struct {
	db *gorm.DB
}

func NewAIGenerationRepository() *AIGenerationRepository {
	return &AIGenerationRepository{db: database.GetDB()}
}

func (r *AIGenerationRepository) Create(g *models.AIGeneration) error {
	return r.db.Create(g).Error
}

func (r *AIGenerationRepository) FindByUserID(userID string, limit, offset int) ([]models.AIGeneration, error) {
	var gens []models.AIGeneration
	if err := r.db.Where("user_id = ?", userID).Order("created_at DESC").Limit(limit).Offset(offset).Find(&gens).Error; err != nil {
		return nil, err
	}
	return gens, nil
}

func (r *AIGenerationRepository) CountByUserID(userID string) (int64, error) {
	var count int64
	if err := r.db.Model(&models.AIGeneration{}).Where("user_id = ?", userID).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
