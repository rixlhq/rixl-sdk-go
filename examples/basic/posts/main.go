// Fetch a single post inside a feed. Posts always live under a feed —
// there's no top-level posts collection.
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
	postID := os.Getenv("RIXL_POST_ID")
	if apiKey == "" || feedID == "" || postID == "" {
		log.Fatal("set RIXL_API_KEY, RIXL_FEED_ID, and RIXL_POST_ID")
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

	post, err := client.Feeds().ByFeedId(feedID).ByPostId(postID).Get(ctx, nil)
	if err != nil {
		log.Fatalf("get post: %v", err)
	}
	log.Printf("post %s", *post.GetId())
}
