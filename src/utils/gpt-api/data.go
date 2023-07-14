package gptapi

import "fmt"

type Model string

type Request struct {
	Model            Model      `json:"model"`
	Messages         []*Message `json:"messages"`
	Temperature      float64    `json:"temperature,omitempty"`
	TopP             float64    `json:"top_p,omitempty"`
	N                int        `json:"n,omitempty"`
	Stop             []string   `json:"stop,omitempty"`
	MaxTokens        int        `json:"max_tokens,omitempty"`
	PresencePenalty  float64    `json:"presence_penalty,omitempty"`
	FrequencyPenalty float64    `json:"frequency_penalty,omitempty"`
	User             string     `json:"user,omitempty"`
}

type Response struct {
	ID       string    `json:"id"`
	Object   string    `json:"object"`
	Created  int64     `json:"created"`
	Choices  []*Choice `json:"choices"`
	Usage    *Usage    `json:"usage"`
	ThreadId string    `json:"-"`
}

type Choice struct {
	Index        int      `json:"index"`
	Message      *Message `json:"message"`
	FinishReason string   `json:"finish_reason"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type ApiError struct {
	StatusCode   int           `json:"-"`
	ErrorDetails *ErrorDetails `json:"error"`
}

func (e *ApiError) Error() string {
	return fmt.Sprintf("ChatGPT API Error: %s", e.ErrorDetails.Message)
}

type ErrorDetails struct {
	Message string `json:"message"`
	Type    string `json:"type"`
	Code    string `json:"code"`
}
