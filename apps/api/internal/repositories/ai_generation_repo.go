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

type AdminGenerationRow struct {
	models.AIGeneration
	UserEmail   string `json:"user_email"`
	ProjectName string `json:"project_name"`
}

func (r *AIGenerationRepository) ListAll(status, genType string, limit, offset int) ([]AdminGenerationRow, int64, error) {
	var rows []AdminGenerationRow
	var total int64

	q := r.db.Model(&models.AIGeneration{})
	if status != "" {
		q = q.Where("status = ?", status)
	}
	if genType != "" {
		q = q.Where("generation_type = ?", genType)
	}
	q.Count(&total)

	qData := r.db.Table("ai_generations").
		Select("ai_generations.*, users.email AS user_email, projects.name AS project_name").
		Joins("LEFT JOIN users ON users.id = ai_generations.user_id AND users.deleted_at IS NULL").
		Joins("LEFT JOIN projects ON projects.id = ai_generations.project_id AND projects.deleted_at IS NULL")
	if status != "" {
		qData = qData.Where("ai_generations.status = ?", status)
	}
	if genType != "" {
		qData = qData.Where("ai_generations.generation_type = ?", genType)
	}
	err := qData.Order("ai_generations.created_at DESC").Limit(limit).Offset(offset).Scan(&rows).Error
	return rows, total, err
}

func (r *AIGenerationRepository) CountAll() (int64, error) {
	var count int64
	return count, r.db.Model(&models.AIGeneration{}).Count(&count).Error
}

func (r *AIGenerationRepository) CountByStatus(status string) (int64, error) {
	var count int64
	return count, r.db.Model(&models.AIGeneration{}).Where("status = ?", status).Count(&count).Error
}
