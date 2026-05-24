package models

// uuid fields stored as char(36) strings

type ProjectType string
type ProjectStatus string

const (
	ProjectTypeSaaS        ProjectType = "saas"
	ProjectTypeAIChatApp   ProjectType = "ai_chat_app"
	ProjectTypeAIAgentApp  ProjectType = "ai_agent_app"
	ProjectTypeAdminDash   ProjectType = "admin_dashboard"
	ProjectTypeEcommerce   ProjectType = "ecommerce"
	ProjectTypeBlogCMS     ProjectType = "blog_cms"
	ProjectTypeLandingPage ProjectType = "landing_page"
	ProjectTypeMobileApp   ProjectType = "mobile_app"
	ProjectTypeDevTool     ProjectType = "developer_tool"
	ProjectTypeInternalSys ProjectType = "internal_system"

	ProjectStatusActive   ProjectStatus = "active"
	ProjectStatusArchived ProjectStatus = "archived"
)

type Project struct {
	BaseModel
	UserID        string        `gorm:"type:char(36);index;not null" json:"user_id"`
	Name          string        `gorm:"size:200;not null" json:"name"`
	Description   string        `gorm:"type:text" json:"description,omitempty"`
	ProjectType   ProjectType   `gorm:"size:50;index;not null" json:"project_type"`
	FrontendStack string        `gorm:"size:200" json:"frontend_stack,omitempty"`
	BackendStack  string        `gorm:"size:200" json:"backend_stack,omitempty"`
	DatabaseStack string        `gorm:"size:200" json:"database_stack,omitempty"`
	AIStack       string        `gorm:"size:200" json:"ai_stack,omitempty"`
	UIStyle       string        `gorm:"size:200" json:"ui_style,omitempty"`
	TargetUser    string        `gorm:"size:200" json:"target_user,omitempty"`
	ProductGoal   string        `gorm:"type:text" json:"product_goal,omitempty"`
	Status        ProjectStatus `gorm:"size:30;default:'active';not null" json:"status"`
}
