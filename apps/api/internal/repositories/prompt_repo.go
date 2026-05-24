package repositories

import (
	"github.com/ai-project-os/api/internal/database"
	"github.com/ai-project-os/api/internal/models"
	"gorm.io/gorm"
)

type PromptRepository struct {
	db *gorm.DB
}

func NewPromptRepository() *PromptRepository {
	return &PromptRepository{db: database.GetDB()}
}

func (r *PromptRepository) Create(p *models.Prompt) error {
	return r.db.Create(p).Error
}

func (r *PromptRepository) FindByProjectID(projectID string, promptType string) ([]models.Prompt, error) {
	var prompts []models.Prompt
	q := r.db.Where("project_id = ?", projectID)
	if promptType != "" {
		q = q.Where("type = ?", promptType)
	}
	if err := q.Order("updated_at DESC").Find(&prompts).Error; err != nil {
		return nil, err
	}
	return prompts, nil
}

func (r *PromptRepository) FindByProjectIDAndUserID(projectID, userID string, promptType string) ([]models.Prompt, error) {
	var prompts []models.Prompt
	q := r.db.Where("project_id = ? AND user_id = ?", projectID, userID)
	if promptType != "" {
		q = q.Where("type = ?", promptType)
	}
	if err := q.Order("updated_at DESC").Find(&prompts).Error; err != nil {
		return nil, err
	}
	return prompts, nil
}

func (r *PromptRepository) FindByID(id string) (*models.Prompt, error) {
	var p models.Prompt
	if err := r.db.First(&p, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *PromptRepository) FindByIDAndProjectIDAndUserID(id, projectID, userID string) (*models.Prompt, error) {
	var p models.Prompt
	if err := r.db.First(&p, "id = ? AND project_id = ? AND user_id = ?", id, projectID, userID).Error; err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *PromptRepository) Update(p *models.Prompt) error {
	return r.db.Save(p).Error
}

func (r *PromptRepository) Delete(id string) error {
	return r.db.Delete(&models.Prompt{}, "id = ?", id).Error
}

func (r *PromptRepository) DeleteByIDAndProjectIDAndUserID(id, projectID, userID string) error {
	result := r.db.Delete(&models.Prompt{}, "id = ? AND project_id = ? AND user_id = ?", id, projectID, userID)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *PromptRepository) CreateVersion(v *models.PromptVersion) error {
	return r.db.Create(v).Error
}

func (r *PromptRepository) FindVersions(promptID string) ([]models.PromptVersion, error) {
	var versions []models.PromptVersion
	if err := r.db.Where("prompt_id = ?", promptID).Order("version DESC").Find(&versions).Error; err != nil {
		return nil, err
	}
	return versions, nil
}

func (r *PromptRepository) FindVersionsByPromptIDAndProjectIDAndUserID(promptID, projectID, userID string) ([]models.PromptVersion, error) {
	var versions []models.PromptVersion
	if err := r.db.
		Joins("JOIN prompts ON prompts.id = prompt_versions.prompt_id").
		Where("prompt_versions.prompt_id = ? AND prompts.project_id = ? AND prompts.user_id = ?", promptID, projectID, userID).
		Order("prompt_versions.version DESC").
		Find(&versions).Error; err != nil {
		return nil, err
	}
	return versions, nil
}

func (r *PromptRepository) CountByUserID(userID string) (int64, error) {
	var count int64
	if err := r.db.Model(&models.Prompt{}).Where("user_id = ?", userID).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
