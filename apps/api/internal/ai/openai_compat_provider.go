package ai

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// ProviderPresets maps provider name → base URL for OpenAI-compatible APIs.
// Add any new provider here — no code changes needed elsewhere.
var ProviderPresets = map[string]string{
	"openai":      "https://api.openai.com/v1",
	"deepseek":    "https://api.deepseek.com/v1",
	"groq":        "https://api.groq.com/openai/v1",
	"moonshot":    "https://api.moonshot.cn/v1",
	"siliconflow": "https://api.siliconflow.cn/v1",
	"zhipu":       "https://open.bigmodel.cn/api/paas/v4",
	"together":    "https://api.together.xyz/v1",
	"ollama":      "http://localhost:11434/v1",
	"mistral":     "https://api.mistral.ai/v1",
	"perplexity":  "https://api.perplexity.ai",
}

// OpenAICompatProvider works with any OpenAI-compatible chat completions endpoint.
type OpenAICompatProvider struct {
	name    string
	baseURL string
	apiKey  string
	model   string
	client  *http.Client
}

// NewOpenAICompatProvider creates a provider using a named preset.
// If the name is not in ProviderPresets, falls back to customBaseURL.
func NewOpenAICompatProvider(name, customBaseURL, apiKey, model string) *OpenAICompatProvider {
	baseURL := customBaseURL
	if preset, ok := ProviderPresets[strings.ToLower(name)]; ok {
		baseURL = preset
	}
	return &OpenAICompatProvider{
		name:    name,
		baseURL: strings.TrimRight(baseURL, "/"),
		apiKey:  apiKey,
		model:   model,
		client:  &http.Client{Timeout: 90 * time.Second},
	}
}

func (p *OpenAICompatProvider) Name() string { return p.name }

// ─── Internal request/response types ─────────────────────────────────────────

type oaiMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type oaiRequest struct {
	Model    string       `json:"model"`
	Messages []oaiMessage `json:"messages"`
}

type oaiResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
	Usage struct {
		TotalTokens int `json:"total_tokens"`
	} `json:"usage"`
	Error *struct {
		Message string `json:"message"`
		Type    string `json:"type,omitempty"`
	} `json:"error,omitempty"`
}

// ─── Generate ─────────────────────────────────────────────────────────────────

func (p *OpenAICompatProvider) Generate(ctx context.Context, req GenerateRequest) (*GenerateResult, error) {
	if p.apiKey == "" && p.name != "ollama" {
		return nil, fmt.Errorf("AI_API_KEY is not set for provider %q", p.name)
	}
	if p.baseURL == "" {
		return nil, fmt.Errorf("no base URL configured for provider %q", p.name)
	}

	payload := oaiRequest{
		Model: p.model,
		Messages: []oaiMessage{
			{Role: "system", Content: systemPromptFor(req.GenerationType)},
			{Role: "user", Content: fmt.Sprintf("Project ID: %s\n\n%s", req.ProjectID, req.InputPayload)},
		},
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	endpoint := p.baseURL + "/chat/completions"
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("Content-Type", "application/json")
	if p.apiKey != "" {
		httpReq.Header.Set("Authorization", "Bearer "+p.apiKey)
	}

	resp, err := p.client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("request to %s failed: %w", p.name, err)
	}
	defer resp.Body.Close()

	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var oaiResp oaiResponse
	if err := json.Unmarshal(raw, &oaiResp); err != nil {
		return nil, fmt.Errorf("failed to parse response from %s: %w", p.name, err)
	}
	if oaiResp.Error != nil {
		return nil, fmt.Errorf("%s API error: %s", p.name, oaiResp.Error.Message)
	}
	if len(oaiResp.Choices) == 0 {
		return nil, errors.New("AI provider returned empty response")
	}

	return &GenerateResult{
		Content:    oaiResp.Choices[0].Message.Content,
		Provider:   p.name,
		Model:      p.model,
		TokensUsed: oaiResp.Usage.TotalTokens,
	}, nil
}

// ─── System prompts ───────────────────────────────────────────────────────────

func systemPromptFor(genType string) string {
	switch genType {
	case "prompt":
		return "你是一位资深 AI 工程师，专精于 Prompt Engineering。根据用户提供的项目信息，生成高质量的系统提示词（System Prompt）。要求：结构清晰、职责明确、包含约束条件、使用中文输出。"
	case "rules":
		return "你是一位技术架构师，专精于制定开发规范。根据用户提供的项目信息，生成 AI 编程规则文件（如 .cursorrules / CLAUDE.md）。要求：涵盖代码风格、架构约定、禁止事项，使用中文输出，Markdown 格式。"
	case "code":
		return "你是一位全栈工程师。根据用户的需求生成高质量、可直接运行的代码。要求：包含必要的注释、错误处理、遵循项目技术栈规范。"
	case "architecture":
		return "你是一位系统架构师。根据用户提供的项目信息，生成详细的技术架构文档。要求：包含系统设计、模块划分、数据流、技术选型理由，使用 Markdown 格式输出。"
	default:
		return "你是一位专业的 AI 助手，请根据用户需求提供专业、准确的回答。"
	}
}
