package ai

import "context"

type GenerateRequest struct {
	ProjectID      string `json:"project_id"`
	GenerationType string `json:"generation_type"`
	InputPayload   string `json:"input_payload"`
}

type GenerateResult struct {
	Content      string `json:"content"`
	Provider     string `json:"provider"`
	Model        string `json:"model"`
	TokensUsed   int    `json:"tokens_used,omitempty"`
}

type Provider interface {
	Name() string
	Generate(ctx context.Context, req GenerateRequest) (*GenerateResult, error)
}
