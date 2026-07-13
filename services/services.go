package services

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"codeReviewer/models"

	"github.com/sashabaranov/go-openai"
)

type AIService struct {
	client *openai.Client
}

func New(apiKey string) *AIService {
	config := openai.DefaultConfig(apiKey)
	config.BaseURL = "https://api.groq.com/openai/v1"

	return &AIService{
		client: openai.NewClientWithConfig(config),
	}
}

func (s *AIService) AnalyzeCode(ctx context.Context, req models.ReviewRequest) (*models.ReviewResponse, error) {
	prompt := fmt.Sprintf(`You are an expert %s code reviewer. Review the following code and return ONLY a raw JSON object (no markdown block backticks) matching this exact structure:
{
  "score": 85,
  "grade": "B+",
  "summary": "Good overall, but...",
  "issues": [
    {"severity": "low/medium/high/critical", "line": 12, "message": "issue description"}
  ],
  "suggestions": ["suggestion 1", "suggestion 2"]
}

Code to review:
%s`, req.Language, req.Code)

	modelName := os.Getenv("GROQ_MODEL")
	if modelName == "" {
		modelName = "llama-3.1-8b-instant" // A supported Groq model
	}

	resp, err := s.client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model: modelName,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: prompt,
			},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to call ai: %v", err)
	}

	content := resp.Choices[0].Message.Content
	// Strip markdown if it was added
	content = strings.TrimPrefix(content, "```json")
	content = strings.TrimPrefix(content, "```")
	content = strings.TrimSuffix(content, "```")
	content = strings.TrimSpace(content)

	var review models.ReviewResponse
	if err := json.Unmarshal([]byte(content), &review); err != nil {
		return nil, fmt.Errorf("failed to parse json response: %v\nResponse was: %s", err, content)
	}

	return &review, nil
}
