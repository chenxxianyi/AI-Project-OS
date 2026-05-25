package repositories

import (
	"github.com/ai-project-os/api/internal/database"
	"github.com/ai-project-os/api/internal/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{db: database.GetDB()}
}

func (r *UserRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) FindByID(id string) (*models.User, error) {
	var user models.User
	if err := r.db.First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

type AdminUserRow struct {
	models.User
	ProjectCount    int64 `json:"project_count"`
	GenerationCount int64 `json:"generation_count"`
}

func (r *UserRepository) ListAll(search, role string, limit, offset int) ([]AdminUserRow, int64, error) {
	var rows []AdminUserRow
	var total int64

	countQ := r.db.Model(&models.User{})
	if search != "" {
		like := "%" + search + "%"
		countQ = countQ.Where("email LIKE ? OR username LIKE ?", like, like)
	}
	if role != "" {
		countQ = countQ.Where("role = ?", role)
	}
	countQ.Count(&total)

	dataQ := r.db.Table("users").
		Select(`users.*, 
			COALESCE((SELECT COUNT(*) FROM projects WHERE projects.user_id = users.id AND projects.deleted_at IS NULL), 0) AS project_count,
			COALESCE((SELECT COUNT(*) FROM ai_generations WHERE ai_generations.user_id = users.id), 0) AS generation_count`).
		Where("users.deleted_at IS NULL")
	if search != "" {
		like := "%" + search + "%"
		dataQ = dataQ.Where("users.email LIKE ? OR users.username LIKE ?", like, like)
	}
	if role != "" {
		dataQ = dataQ.Where("users.role = ?", role)
	}
	err := dataQ.Order("users.created_at DESC").Limit(limit).Offset(offset).Scan(&rows).Error
	return rows, total, err
}

func (r *UserRepository) CountAll() (int64, error) {
	var count int64
	return count, r.db.Model(&models.User{}).Count(&count).Error
}

func (r *UserRepository) CountByStatus(status models.UserStatus) (int64, error) {
	var count int64
	return count, r.db.Model(&models.User{}).Where("status = ?", status).Count(&count).Error
}

func (r *UserRepository) Update(user *models.User) error {
	return r.db.Save(user).Error
}

func (r *UserRepository) Delete(id string) error {
	return r.db.Delete(&models.User{}, "id = ?", id).Error
}
