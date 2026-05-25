package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// uuid fields stored as char(36) strings

type PromptType string

const (
	PromptTypeBackend     PromptType = "backend_prompt"
	PromptTypeFrontend    PromptType = "frontend_prompt"
	PromptTypeFullstack   PromptType = "fullstack_prompt"
	PromptTypeUIDesign    PromptType = "ui_design_prompt"
	PromptTypeDatabase    PromptType = "database_prompt"
	PromptTypeAPIDesign   PromptType = "api_design_prompt"
	PromptTypeTesting     PromptType = "testing_prompt"
	PromptTypeDeployment  PromptType = "deployment_prompt"
	PromptTypeCursorRules PromptType = "cursor_rules"
	PromptTypeClaudeMD    PromptType = "claude_md"
	PromptTypeAgentsMD    PromptType = "agents_md"
)

type Prompt struct {
	BaseModel
	ProjectID string     `gorm:"type:char(36);index;not null" json:"project_id"`
	UserID    string     `gorm:"type:char(36);index;not null" json:"user_id"`
	Type      PromptType `gorm:"size:50;index;not null" json:"type"`
	Title     string     `gorm:"size:300;not null" json:"title"`
	Content   string     `gorm:"type:text;not null" json:"content"`
	Version   int        `gorm:"default:1;not null" json:"version"`
}

type PromptVersion struct {
	ID         string `gorm:"type:char(36);primaryKey" json:"id"`
	PromptID   string `gorm:"type:char(36);index;not null" json:"prompt_id"`
	Version    int    `gorm:"not null" json:"version"`
	Content    string `gorm:"type:text;not null" json:"content"`
	ChangeNote string `gorm:"size:500" json:"change_note,omitempty"`
	CreatedAt  int64  `json:"created_at"`
}

func (v *PromptVersion) BeforeCreate(tx *gorm.DB) error {
	if v.ID == "" {
		v.ID = uuid.New().String()
	}
	return nil
}
