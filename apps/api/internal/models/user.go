package models

type UserRole string

const (
	RoleAdmin UserRole = "admin"
	RoleUser  UserRole = "user"
)

type User struct {
	BaseModel
	Email        string   `gorm:"uniqueIndex;size:255;not null" json:"email"`
	Username     string   `gorm:"size:100;not null" json:"username"`
	PasswordHash string   `gorm:"size:255;not null" json:"-"`
	Avatar       string   `gorm:"size:500" json:"avatar,omitempty"`
	Role         UserRole `gorm:"size:20;default:'user';not null" json:"role"`
}
