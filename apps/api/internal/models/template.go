package models

type Template struct {
	BaseModel
	Name           string `gorm:"size:200;not null" json:"name"`
	Description    string `gorm:"type:text" json:"description,omitempty"`
	ProjectType    string `gorm:"size:50;not null" json:"project_type"`
	FrontendStack  string `gorm:"size:200" json:"frontend_stack,omitempty"`
	BackendStack   string `gorm:"size:200" json:"backend_stack,omitempty"`
	DatabaseStack  string `gorm:"size:200" json:"database_stack,omitempty"`
	AIStack        string `gorm:"size:200" json:"ai_stack,omitempty"`
	UIStyle        string `gorm:"size:200" json:"ui_style,omitempty"`
	DefaultPrompts string `gorm:"type:text" json:"default_prompts,omitempty"`
	DefaultRules   string `gorm:"type:text" json:"default_rules,omitempty"`
}
