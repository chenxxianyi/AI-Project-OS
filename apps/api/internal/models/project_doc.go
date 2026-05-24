package models

// uuid fields stored as char(36) strings

type DocType string

const (
	DocTypeArchitecture    DocType = "architecture"
	DocTypeAPIDocs         DocType = "api_docs"
	DocTypeDatabaseSchema  DocType = "database_schema"
	DocTypeUIGuidelines    DocType = "ui_guidelines"
	DocTypeCodingStandards DocType = "coding_standards"
	DocTypeProductReqs     DocType = "product_requirements"
	DocTypeChangelog       DocType = "changelog"
	DocTypeDecisions       DocType = "decisions"
)

type ProjectDoc struct {
	BaseModel
	ProjectID string  `gorm:"type:char(36);index;not null" json:"project_id"`
	UserID    string  `gorm:"type:char(36);index;not null" json:"user_id"`
	DocType   DocType `gorm:"size:50;index;not null" json:"doc_type"`
	Title     string  `gorm:"size:300;not null" json:"title"`
	Content   string  `gorm:"type:text;not null" json:"content"`
}
