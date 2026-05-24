package repositories

import (
	"github.com/ai-project-os/api/internal/database"
	"github.com/ai-project-os/api/internal/models"
	"gorm.io/gorm"
)

type ProjectDocRepository struct {
	db *gorm.DB
}

func NewProjectDocRepository() *ProjectDocRepository {
	return &ProjectDocRepository{db: database.GetDB()}
}

func (r *ProjectDocRepository) Create(doc *models.ProjectDoc) error {
	return r.db.Create(doc).Error
}

func (r *ProjectDocRepository) FindByProjectID(projectID string, docType string) ([]models.ProjectDoc, error) {
	var docs []models.ProjectDoc
	q := r.db.Where("project_id = ?", projectID)
	if docType != "" {
		q = q.Where("doc_type = ?", docType)
	}
	if err := q.Order("updated_at DESC").Find(&docs).Error; err != nil {
		return nil, err
	}
	return docs, nil
}

func (r *ProjectDocRepository) FindByProjectIDAndUserID(projectID, userID string, docType string) ([]models.ProjectDoc, error) {
	var docs []models.ProjectDoc
	q := r.db.Where("project_id = ? AND user_id = ?", projectID, userID)
	if docType != "" {
		q = q.Where("doc_type = ?", docType)
	}
	if err := q.Order("updated_at DESC").Find(&docs).Error; err != nil {
		return nil, err
	}
	return docs, nil
}

func (r *ProjectDocRepository) FindByID(id string) (*models.ProjectDoc, error) {
	var doc models.ProjectDoc
	if err := r.db.First(&doc, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &doc, nil
}

func (r *ProjectDocRepository) FindByIDAndProjectIDAndUserID(id, projectID, userID string) (*models.ProjectDoc, error) {
	var doc models.ProjectDoc
	if err := r.db.First(&doc, "id = ? AND project_id = ? AND user_id = ?", id, projectID, userID).Error; err != nil {
		return nil, err
	}
	return &doc, nil
}

func (r *ProjectDocRepository) Update(doc *models.ProjectDoc) error {
	return r.db.Save(doc).Error
}

func (r *ProjectDocRepository) Delete(id string) error {
	return r.db.Delete(&models.ProjectDoc{}, "id = ?", id).Error
}

func (r *ProjectDocRepository) DeleteByIDAndProjectIDAndUserID(id, projectID, userID string) error {
	result := r.db.Delete(&models.ProjectDoc{}, "id = ? AND project_id = ? AND user_id = ?", id, projectID, userID)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
