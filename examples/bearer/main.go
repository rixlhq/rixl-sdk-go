package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	abs "github.com/microsoft/kiota-abstractions-go"
	kiotahttp "github.com/microsoft/kiota-http-go"
	"github.com/qeeqez/rixl-sdk-go/sdk"
)

// Kiota's stock providers reject non-HTTPS URLs.
type bearerAuth struct{ token string }

func (b *bearerAuth) AuthenticateRequest(_ context.Context, req *abs.RequestInformation, _ map[string]any) error {
	req.Headers.Add("Authorization", "Bearer "+b.token)
	return nil
}

type tokenRequest struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Subject      string `json:"subject"`
	ProjectID    string `json:"project_id,omitempty"`
	TTLMinutes   *int   `json:"ttl_minutes,omitempty"`
}

type tokenResponse struct {
	AccessToken string    `json:"access_token"`
	TokenType   string    `json:"token_type"`
	ExpiresIn   int64     `json:"expires_in"`
	ExpiresAt   time.Time `json:"expires_at"`
}

func mintToken(ctx context.Context, baseURL string, body tokenRequest) (*tokenResponse, error) {
	payload, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, baseURL+"/clientauth/token", bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		raw, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("mint token: %s: %s", resp.Status, string(raw))
	}
	var out tokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func main() {
	clientID := mustEnv("RIXL_CLIENT_ID")
	clientSecret := mustEnv("RIXL_CLIENT_SECRET")
	projectID := mustEnv("RIXL_PROJECT_ID")
	subject := mustEnv("RIXL_SUBJECT")
	baseURL := envOr("RIXL_BASE_URL", "http://localhost:8081")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	tok, err := mintToken(ctx, baseURL, tokenRequest{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Subject:      subject,
		ProjectID:    projectID,
	})
	if err != nil {
		log.Fatalf("mint: %v", err)
	}
	fmt.Printf("minted token (expires_in=%ds, type=%s)\n", tok.ExpiresIn, tok.TokenType)

	adapter, err := kiotahttp.NewNetHttpRequestAdapter(&bearerAuth{token: tok.AccessToken})
	if err != nil {
		log.Fatalf("adapter: %v", err)
	}
	adapter.SetBaseUrl(baseURL)
	client := sdk.NewRixlClient(adapter)

	page, err := client.Images().Get(ctx, nil)
	if err != nil {
		log.Fatalf("list images: %v", err)
	}
	fmt.Printf("Listed %d images\n", len(page.GetData()))
}

func mustEnv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Fatalf("missing %s", k)
	}
	return v
}

func envOr(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}
