package gptapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	BaseUrl = "https://api.openai.com/v1"
	ChatUrl = BaseUrl + "/chat/completions"
)

type Api struct {
	organizationId string
	key            string
	client         *http.Client
}

func NewApi(key string) *Api {
	return &Api{
		key:    key,
		client: http.DefaultClient,
	}
}

func (a *Api) WithClient(client *http.Client) *Api {
	a.client = client
	return a
}

func (a *Api) WithOrganizationId(organizationId string) *Api {
	a.organizationId = organizationId
	return a
}

func (a *Api) Chat(r *Request) (*Response, error) {
	body, err := json.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, ChatUrl, bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	response, err := a.do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	if response.StatusCode != http.StatusOK {
		apiError := &ApiError{
			StatusCode: response.StatusCode,
		}
		if err := json.NewDecoder(response.Body).Decode(apiError); err != nil {
			return nil, fmt.Errorf("request failed with status code %d", response.StatusCode)
		}

		return nil, apiError
	}

	result := &Response{}
	if err := json.NewDecoder(response.Body).Decode(result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return result, nil
}

func (a *Api) do(req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", "Bearer "+a.key)
	req.Header.Set("Content-Type", "application/json")
	if a.organizationId != "" {
		req.Header.Set("OpenAI-Organization", a.organizationId)
	}

	return a.client.Do(req)
}
