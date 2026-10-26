package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/matisiekpl/similarcli/internal/dto"
)

type RdapRepository interface {
	Load(ctx context.Context, domain string) (dto.RdapDomain, error)
}

type rdapRepository struct{}

func NewRdapRepository() RdapRepository {
	return &rdapRepository{}
}

func (r *rdapRepository) Load(ctx context.Context, domain string) (dto.RdapDomain, error) {
	url := fmt.Sprintf("https://rdap.org/domain/%s", domain)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return dto.RdapDomain{}, fmt.Errorf("creating request: %w", err)
	}

	req.Header.Set("accept", "application/rdap+json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return dto.RdapDomain{}, fmt.Errorf("making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return dto.RdapDomain{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var result dto.RdapDomain
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return dto.RdapDomain{}, fmt.Errorf("decoding response: %w", err)
	}

	return result, nil
}
