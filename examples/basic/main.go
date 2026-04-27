package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	abs "github.com/microsoft/kiota-abstractions-go"
	kiotahttp "github.com/microsoft/kiota-http-go"
	"github.com/qeeqez/rixl-sdk-go/sdk"
)

// Kiota's stock ApiKeyAuthenticationProvider rejects non-HTTPS URLs.
type apiKeyHeaderAuth struct{ key string }

func (a *apiKeyHeaderAuth) AuthenticateRequest(_ context.Context, req *abs.RequestInformation, _ map[string]any) error {
	req.Headers.Add("X-API-Key", a.key)
	return nil
}

func main() {
	apiKey := mustEnv("RIXL_API_KEY")
	baseURL := envOr("RIXL_BASE_URL", "http://localhost:8081")

	adapter, err := kiotahttp.NewNetHttpRequestAdapter(&apiKeyHeaderAuth{key: apiKey})
	if err != nil {
		log.Fatalf("adapter: %v", err)
	}
	adapter.SetBaseUrl(baseURL)
	client := sdk.NewRixlClient(adapter)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	page, err := client.Images().Get(ctx, nil)
	if err != nil {
		log.Fatalf("list images: %v", err)
	}
	fmt.Printf("Listed %d images\n", len(page.GetData()))
	for _, img := range page.GetData() {
		if img.GetId() != nil {
			fmt.Printf("  - %s\n", *img.GetId())
		}
	}

	if id := os.Getenv("IMAGE_ID"); id != "" {
		image, err := client.Images().ByImageId(id).Get(ctx, nil)
		if err != nil {
			log.Fatalf("get image %s: %v", id, err)
		}
		fmt.Printf("Image %s: %dx%d\n",
			deref(image.GetId()), int32Or(image.GetWidth()), int32Or(image.GetHeight()))
	}
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

func deref(p *string) string {
	if p == nil {
		return ""
	}
	return *p
}

func int32Or(p *int32) int32 {
	if p == nil {
		return 0
	}
	return *p
}
