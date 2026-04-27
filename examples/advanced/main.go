package main

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	abs "github.com/microsoft/kiota-abstractions-go"
	kiotahttp "github.com/microsoft/kiota-http-go"
	"github.com/qeeqez/rixl-sdk-go/sdk"
	"github.com/qeeqez/rixl-sdk-go/sdk/models"
	apierr "github.com/qeeqez/rixl-sdk-go/sdk/models/github_com_qeeqez_api_internal_errors"
	vidupload "github.com/qeeqez/rixl-sdk-go/sdk/models/github_com_qeeqez_api_internal_videos_handler_upload"
	imghandler "github.com/qeeqez/rixl-sdk-go/sdk/models/internal_images_handler"
)

const (
	sampleImageURL = "https://picsum.photos/seed/rixl/800/600.jpg"
	sampleVideoURL = "https://download.samplelib.com/mp4/sample-5s.mp4"
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

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	uploadImage(ctx, client)
	uploadVideo(ctx, client)
}

func uploadImage(ctx context.Context, c *sdk.RixlClient) {
	fmt.Println("== Image upload ==")
	body, err := download(ctx, sampleImageURL)
	if err != nil {
		log.Fatalf("download image: %v", err)
	}

	initReq := imghandler.NewUploadInitRequest()
	name, format := "sample.jpg", "jpeg"
	initReq.SetName(&name)
	initReq.SetFormat(&format)

	initRes, err := c.Images().Upload().Init().Post(ctx, initReq, nil)
	if err != nil {
		log.Fatalf("image init: %s", explain(err))
	}
	fmt.Printf("init: image_id=%s\n", deref(initRes.GetImageId()))

	if err := putBytes(ctx, deref(initRes.GetPresignedUrl()), body, "image/jpeg"); err != nil {
		log.Fatalf("PUT image: %v", err)
	}

	completeReq := imghandler.NewCompleteRequest()
	completeReq.SetImageId(initRes.GetImageId())
	notAttached := false
	completeReq.SetAttachedToVideo(&notAttached)

	image, err := c.Images().Upload().Complete().Post(ctx, completeReq, nil)
	if err != nil {
		log.Fatalf("image complete: %s", explain(err))
	}
	fmt.Printf("complete: id=%s %dx%d\n\n",
		deref(image.GetId()), int32Or(image.GetWidth()), int32Or(image.GetHeight()))
}

func uploadVideo(ctx context.Context, c *sdk.RixlClient) {
	fmt.Println("== Video upload ==")
	video, err := download(ctx, sampleVideoURL)
	if err != nil {
		log.Fatalf("download video: %v", err)
	}
	poster, err := download(ctx, sampleImageURL)
	if err != nil {
		log.Fatalf("download poster: %v", err)
	}

	initReq := models.NewVideoUploadInitRequest()
	fileName, posterFormat := "sample.mp4", "jpeg"
	initReq.SetFileName(&fileName)
	initReq.SetImageFormat(&posterFormat)

	initRes, err := c.Videos().Upload().Init().Post(ctx, initReq, nil)
	if err != nil {
		log.Fatalf("video init: %s", explain(err))
	}
	fmt.Printf("init: video_id=%s poster_id=%s\n",
		deref(initRes.GetVideoId()), deref(initRes.GetPosterId()))

	if err := putBytes(ctx, deref(initRes.GetVideoPresignedUrl()), video, "video/mp4"); err != nil {
		log.Fatalf("PUT video: %v", err)
	}
	if err := putBytes(ctx, deref(initRes.GetPosterPresignedUrl()), poster, "image/jpeg"); err != nil {
		log.Fatalf("PUT poster: %v", err)
	}

	completeReq := vidupload.NewCompleteRequest()
	completeReq.SetVideoId(initRes.GetVideoId())

	finished, err := c.Videos().Upload().Complete().Post(ctx, completeReq, nil)
	if err != nil {
		log.Fatalf("video complete: %s", explain(err))
	}
	fmt.Printf("complete: id=%s\n", deref(finished.GetId()))
}

func download(ctx context.Context, url string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GET %s: %s", url, resp.Status)
	}
	return io.ReadAll(resp.Body)
}

func putBytes(ctx context.Context, url string, body []byte, contentType string) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Minute)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodPut, url, bytes.NewReader(body))
	if err != nil {
		return err
	}
	req.ContentLength = int64(len(body))
	req.Header.Set("Content-Type", contentType)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 300 {
		b, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("PUT %s: %s: %s", url, resp.Status, string(b))
	}
	return nil
}

func explain(err error) string {
	var apiErr *apierr.ErrorResponse
	if errors.As(err, &apiErr) {
		code, msg := int32(0), ""
		if apiErr.GetCode() != nil {
			code = *apiErr.GetCode()
		}
		if apiErr.GetErrorEscaped() != nil {
			msg = *apiErr.GetErrorEscaped()
		}
		return fmt.Sprintf("HTTP %d: %s", code, msg)
	}
	return err.Error()
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
