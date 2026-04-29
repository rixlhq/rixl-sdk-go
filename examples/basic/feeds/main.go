// Read a feed and print the posts attached to it.
package main

import (
	"context"
	"log"
	"os"
	"time"

	abs "github.com/microsoft/kiota-abstractions-go"
	kiotahttp "github.com/microsoft/kiota-http-go"

	"github.com/rixlhq/rixl-go/sdk"
)

type apiKeyAuth struct{ key string }

func (a *apiKeyAuth) AuthenticateRequest(_ context.Context, req *abs.RequestInformation, _ map[string]any) error {
	req.Headers.Add("X-API-Key", a.key)
	return nil
}

func main() {
	apiKey := os.Getenv("RIXL_API_KEY")
	feedID := os.Getenv("RIXL_FEED_ID")
	if apiKey == "" || feedID == "" {
		log.Fatal("set RIXL_API_KEY and RIXL_FEED_ID")
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

	page, err := client.Feeds().ByFeedId(feedID).Get(ctx, nil)
	if err != nil {
		log.Fatalf("get feed %s: %v", feedID, err)
	}
	log.Printf("feed %s — %d posts", feedID, len(page.GetData()))
	for _, post := range page.GetData() {
		if post.GetId() != nil {
			log.Printf("  - %s", *post.GetId())
		}
	}
}
