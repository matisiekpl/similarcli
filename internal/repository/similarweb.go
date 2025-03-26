package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/matisiekpl/similarcli/internal/dto"
)

type SimilarWebRepository interface {
	Load(ctx context.Context, address string) (dto.Website, error)
}

type similarWebRepository struct{}

func NewSimilarWebRepository() SimilarWebRepository {
	return &similarWebRepository{}
}

func (r *similarWebRepository) Load(ctx context.Context, address string) (dto.Website, error) {
	url := fmt.Sprintf("https://data.similarweb.com/api/v1/data?domain=%s", address)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return dto.Website{}, fmt.Errorf("creating request: %w", err)
	}

	req.Header.Set("accept", "*/*")
	req.Header.Set("accept-language", "en-EN")
	req.Header.Set("content-type", "application/json")
	req.Header.Set("priority", "u=1, i")
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "none")
	req.Header.Set("sec-fetch-storage-access", "active")
	req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/134.0.0.0 Safari/537.36")
	req.Header.Set("x-extension-version", "6.12.6")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return dto.Website{}, fmt.Errorf("making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return dto.Website{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var result dto.Website
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return dto.Website{}, fmt.Errorf("decoding response: %w", err)
	}

	return result, nil
}
