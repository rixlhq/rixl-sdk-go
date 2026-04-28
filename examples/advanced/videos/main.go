// Upload a video end-to-end. Same shape as the image flow, but Init returns
// two presigned URLs (one for the video, one for the poster thumbnail) and
// we PUT to both.
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

	"github.com/rixlhq/sdk-go/sdk"
	"github.com/rixlhq/sdk-go/sdk/models"
	vidupload "github.com/rixlhq/sdk-go/sdk/models/github_com_rixlhq_api_internal_videos_handler_upload"
)

const (
	sampleVideoURL  = "https://download.samplelib.com/mp4/sample-5s.mp4"
	samplePosterURL = "https://picsum.photos/seed/rixl/800/600.jpg"
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

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	video, err := download(ctx, sampleVideoURL)
	if err != nil {
		log.Fatalf("download video: %v", err)
	}
	poster, err := download(ctx, samplePosterURL)
	if err != nil {
		log.Fatalf("download poster: %v", err)
	}
	log.Printf("downloaded video=%d poster=%d", len(video), len(poster))

	initReq := models.NewVideoUploadInitRequest()
	fileName, posterFormat := "sample.mp4", "jpeg"
	initReq.SetFileName(&fileName)
	initReq.SetImageFormat(&posterFormat)

	init, err := client.Videos().Upload().Init().Post(ctx, initReq, nil)
	if err != nil {
		log.Fatalf("init: %v", err)
	}
	log.Printf("init: video_id=%s poster_id=%s", *init.GetVideoId(), *init.GetPosterId())

	if err := putBytes(ctx, *init.GetVideoPresignedUrl(), video, "video/mp4"); err != nil {
		log.Fatalf("PUT video: %v", err)
	}
	if err := putBytes(ctx, *init.GetPosterPresignedUrl(), poster, "image/jpeg"); err != nil {
		log.Fatalf("PUT poster: %v", err)
	}

	completeReq := vidupload.NewCompleteRequest()
	completeReq.SetVideoId(init.GetVideoId())

	v, err := client.Videos().Upload().Complete().Post(ctx, completeReq, nil)
	if err != nil {
		log.Fatalf("complete: %v", err)
	}
	log.Printf("complete: id=%s", *v.GetId())
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
