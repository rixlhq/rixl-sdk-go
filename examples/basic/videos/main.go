// List videos in your project, optionally fetch one by ID.
package main

import (
	"context"
	"log"
	"os"
	"time"

	abs "github.com/microsoft/kiota-abstractions-go"
	kiotahttp "github.com/microsoft/kiota-http-go"

	"github.com/qeeqez/rixl-sdk-go/sdk"
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

	page, err := client.Videos().Get(ctx, nil)
	if err != nil {
		log.Fatalf("list: %v", err)
	}
	log.Printf("listed %d videos", len(page.GetData()))
	for _, v := range page.GetData() {
		if v.GetId() != nil {
			log.Printf("  - %s", *v.GetId())
		}
	}

	id := os.Getenv("VIDEO_ID")
	if id == "" {
		return
	}
	v, err := client.Videos().ByVideoId(id).Get(ctx, nil)
	if err != nil {
		log.Fatalf("get %s: %v", id, err)
	}
	log.Printf("video %s", *v.GetId())
}
