package repositories

import (
	"github.com/ai-project-os/api/internal/database"
	"github.com/ai-project-os/api/internal/models"
	"gorm.io/gorm"
)

type ProjectRepository struct {
	db *gorm.DB
}

func NewProjectRepository() *ProjectRepository {
	return &ProjectRepository{db: database.GetDB()}
}

func (r *ProjectRepository) Create(project *models.Project) error {
	return r.db.Create(project).Error
}

func (r *ProjectRepository) FindByUserID(userID string, projectType, search string, limit, offset int) ([]models.Project, int64, error) {
	var projects []models.Project
	var total int64
	q := r.db.Where("user_id = ?", userID)
	if projectType != "" {
		q = q.Where("project_type = ?", projectType)
	}
	if search != "" {
		q = q.Where("name LIKE ?", "%"+search+"%")
	}
	q.Model(&models.Project{}).Count(&total)
	if err := q.Order("updated_at DESC").Limit(limit).Offset(offset).Find(&projects).Error; err != nil {
		return nil, 0, err
	}
	return projects, total, nil
}

func (r *ProjectRepository) FindByID(id string) (*models.Project, error) {
	var p models.Project
	if err := r.db.First(&p, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *ProjectRepository) FindByIDAndUserID(id, userID string) (*models.Project, error) {
	var p models.Project
	if err := r.db.First(&p, "id = ? AND user_id = ?", id, userID).Error; err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *ProjectRepository) Update(p *models.Project) error {
	return r.db.Save(p).Error
}

func (r *ProjectRepository) Delete(id string) error {
	return r.db.Delete(&models.Project{}, "id = ?", id).Error
}

func (r *ProjectRepository) DeleteByIDAndUserID(id, userID string) error {
	result := r.db.Delete(&models.Project{}, "id = ? AND user_id = ?", id, userID)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

type AdminProjectRow struct {
	models.Project
	OwnerName string `json:"owner_name"`
}

func (r *ProjectRepository) ListAll(search, status string, limit, offset int) ([]AdminProjectRow, int64, error) {
	var rows []AdminProjectRow
	var total int64

	q := r.db.Model(&models.Project{})
	if search != "" {
		q = q.Where("name LIKE ?", "%"+search+"%")
	}
	if status != "" {
		q = q.Where("status = ?", status)
	}
	q.Count(&total)

	qData := r.db.Table("projects").
		Select("projects.*, users.username AS owner_name").
		Joins("LEFT JOIN users ON users.id = projects.user_id AND users.deleted_at IS NULL").
		Where("projects.deleted_at IS NULL")
	if search != "" {
		qData = qData.Where("projects.name LIKE ?", "%"+search+"%")
	}
	if status != "" {
		qData = qData.Where("projects.status = ?", status)
	}
	err := qData.Order("projects.updated_at DESC").Limit(limit).Offset(offset).Scan(&rows).Error
	return rows, total, err
}

func (r *ProjectRepository) CountAll() (int64, error) {
	var count int64
	return count, r.db.Model(&models.Project{}).Count(&count).Error
}
