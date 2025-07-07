package clients

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ajohnston1219/eatme/api/models"
)

type MLClient struct {
	http *http.Client
	host string
}

func NewMLClient(host string) *MLClient {
	return &MLClient{
		http: &http.Client{},
		host: host,
	}
}

func (c *MLClient) Chat(ctx context.Context, req *models.InternalChatRequest) (*models.ChatResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("marshal req: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, c.host+"/chat", bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("new req: %w", err)
	}
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := c.http.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("ml call: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("ml bad status: %s", resp.Status)
	}

	var mlResp models.ChatResponse
	if err := json.NewDecoder(resp.Body).Decode(&mlResp); err != nil {
		return nil, fmt.Errorf("decode ml resp: %w", err)
	}
	return &mlResp, nil
}
