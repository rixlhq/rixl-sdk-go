// Upload an image end-to-end:
//
//   1. Init     — tell the API you want to upload; it returns a presigned PUT URL.
//   2. PUT      — push the bytes straight to storage (the API never sees them).
//   3. Complete — tell the API the upload landed so it can finalize the record.
package main

import (
	"bytes"
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	abs "github.com/microsoft/kiota-abstractions-go"
	kiotahttp "github.com/microsoft/kiota-http-go"

	"github.com/rixlhq/rixl-go/sdk"
	imghandler "github.com/rixlhq/rixl-go/sdk/models/internal_images_handler"
)

const sampleImageURL = "https://picsum.photos/seed/rixl/800/600.jpg"

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

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	body, err := download(ctx, sampleImageURL)
	if err != nil {
		log.Fatalf("download: %v", err)
	}
	log.Printf("downloaded %d bytes", len(body))

	initReq := imghandler.NewUploadInitRequest()
	name, format := "sample.jpg", "jpeg"
	initReq.SetName(&name)
	initReq.SetFormat(&format)

	init, err := client.Images().Upload().Init().Post(ctx, initReq, nil)
	if err != nil {
		log.Fatalf("init: %v", err)
	}
	log.Printf("init: image_id=%s", *init.GetImageId())

	if err := putBytes(ctx, *init.GetPresignedUrl(), body, "image/jpeg"); err != nil {
		log.Fatalf("PUT: %v", err)
	}

	completeReq := imghandler.NewCompleteRequest()
	completeReq.SetImageId(init.GetImageId())
	notAttached := false
	completeReq.SetAttachedToVideo(&notAttached)

	img, err := client.Images().Upload().Complete().Post(ctx, completeReq, nil)
	if err != nil {
		log.Fatalf("complete: %v", err)
	}
	log.Printf("complete: id=%s %dx%d", *img.GetId(), *img.GetWidth(), *img.GetHeight())
}

func download(ctx context.Context, url string) ([]byte, error) {
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

func putBytes(ctx context.Context, url string, body []byte, contentType string) error {
	req, _ := http.NewRequestWithContext(ctx, http.MethodPut, url, bytes.NewReader(body))
	req.ContentLength = int64(len(body))
	req.Header.Set("Content-Type", contentType)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 300 {
		raw, _ := io.ReadAll(resp.Body)
		return &httpError{Status: resp.Status, Body: string(raw)}
	}
	return nil
}

type httpError struct{ Status, Body string }

func (e *httpError) Error() string { return e.Status + ": " + e.Body }
