package repositories

import (
	"github.com/ai-project-os/api/internal/database"
	"github.com/ai-project-os/api/internal/models"
	"gorm.io/gorm"
)

type TemplateRepository struct {
	db *gorm.DB
}

func NewTemplateRepository() *TemplateRepository {
	return &TemplateRepository{db: database.GetDB()}
}

func (r *TemplateRepository) FindAll(projectType string, limit, offset int) ([]models.Template, int64, error) {
	var templates []models.Template
	var total int64
	q := r.db.Model(&models.Template{})
	if projectType != "" {
		q = q.Where("project_type = ?", projectType)
	}
	q.Count(&total)
	if err := q.Order("created_at DESC").Limit(limit).Offset(offset).Find(&templates).Error; err != nil {
		return nil, 0, err
	}
	return templates, total, nil
}

func (r *TemplateRepository) FindByID(id string) (*models.Template, error) {
	var t models.Template
	if err := r.db.First(&t, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &t, nil
}

func (r *TemplateRepository) Create(t *models.Template) error {
	return r.db.Create(t).Error
}
