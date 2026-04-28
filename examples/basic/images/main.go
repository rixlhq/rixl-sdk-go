// List images in your project, optionally fetch one by ID.
package main

import (
	"context"
	"log"
	"os"
	"time"

	abs "github.com/microsoft/kiota-abstractions-go"
	kiotahttp "github.com/microsoft/kiota-http-go"

	"github.com/rixlhq/sdk-go/sdk"
)

type apiKeyAuth struct{ key string }

func (a *apiKeyAuth) AuthenticateRequest(_ context.Context, req *abs.RequestInformation, _ map[string]any) error {
	req.Headers.Add("X-API-Key", a.key)
	return nil
}

func main() {
	apiKey := os.Getenv("RIXL_API_KEY")
	if apiKey == "" {
		log.Fatal("missing RIXL_API_KEY")
	}
	baseURL := os.Getenv("RIXL_BASE_URL")
	if baseURL == "" {
		baseURL = "http://localhost:8081"
	}

	adapter, err := kiotahttp.NewNetHttpRequestAdapter(&apiKeyAuth{key: apiKey})
	if err != nil {
		log.Fatalf("adapter: %v", err)
	}
	adapter.SetBaseUrl(baseURL)
	client := sdk.NewRixlClient(adapter)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	page, err := client.Images().Get(ctx, nil)
	if err != nil {
		log.Fatalf("list: %v", err)
	}
	log.Printf("listed %d images", len(page.GetData()))
	for _, img := range page.GetData() {
		if img.GetId() != nil {
			log.Printf("  - %s", *img.GetId())
		}
	}

	id := os.Getenv("IMAGE_ID")
	if id == "" {
		return
	}
	img, err := client.Images().ByImageId(id).Get(ctx, nil)
	if err != nil {
		log.Fatalf("get %s: %v", id, err)
	}
	log.Printf("image %s: %dv x %dv", *img.GetId(), *img.GetWidth(), *img.GetHeight())
}
