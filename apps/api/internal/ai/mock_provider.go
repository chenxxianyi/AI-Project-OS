package ai

import (
	"context"
	"fmt"
)

type MockProvider struct {
	model string
}

func NewMockProvider(model string) *MockProvider {
	return &MockProvider{model: model}
}

func (p *MockProvider) Name() string {
	return "mock"
}

func (p *MockProvider) Generate(ctx context.Context, req GenerateRequest) (*GenerateResult, error) {
	var content string
	switch req.GenerationType {
	case "prompt":
		content = fmt.Sprintf("# AI Generated Prompt\n\n## Project Context\n%s\n\n## Generated Prompt\n\nYou are an expert developer. Based on the following project requirements:\n\n%s\n\nPlease generate high-quality code that follows best practices, includes proper error handling, and is well-documented.\n\n### Guidelines:\n- Follow the project's tech stack conventions\n- Write clean, maintainable code\n- Include appropriate tests\n- Document public APIs", req.ProjectID, req.InputPayload)
	case "rules":
		content = fmt.Sprintf("# Project Rules\n\n## General Rules\n- Follow the project's coding standards\n- Write clear commit messages\n- Keep functions focused and small\n\n## Context\n%s\n\n## AI Behavior Rules\n- Always ask for clarification when requirements are ambiguous\n- Prefer existing patterns in the codebase\n- Suggest improvements when you spot potential issues\n- Document any assumptions made", req.InputPayload)
	default:
		content = fmt.Sprintf("# Generated Content\n\nInput: %s", req.InputPayload)
	}

	return &GenerateResult{
		Content:    content,
		Provider:   p.Name(),
		Model:      p.model,
		TokensUsed: len(content) / 4,
	}, nil
}
