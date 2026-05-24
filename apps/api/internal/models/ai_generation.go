package models

// uuid fields stored as char(36) strings

type GenerationType string
type GenerationStatus string

const (
	GenerationTypePrompt GenerationType = "prompt"
	GenerationTypeRules  GenerationType = "rules"

	GenerationStatusSuccess GenerationStatus = "success"
	GenerationStatusFailed  GenerationStatus = "failed"
)

type AIGeneration struct {
	ID             string           `gorm:"type:char(36);primaryKey" json:"id"`
	UserID         string           `gorm:"type:char(36);index;not null" json:"user_id"`
	ProjectID      string           `gorm:"type:char(36);index;not null" json:"project_id"`
	GenerationType GenerationType   `gorm:"size:30;index;not null" json:"generation_type"`
	InputPayload   string           `gorm:"type:text" json:"input_payload,omitempty"`
	OutputContent  string           `gorm:"type:text" json:"output_content,omitempty"`
	ModelProvider  string           `gorm:"size:50" json:"model_provider,omitempty"`
	ModelName      string           `gorm:"size:100" json:"model_name,omitempty"`
	Status         GenerationStatus `gorm:"size:20;not null" json:"status"`
	ErrorMessage   string           `gorm:"type:text" json:"error_message,omitempty"`
	CreatedAt      int64            `json:"created_at"`
}
