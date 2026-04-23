# RIXL Go SDK

The official Go client for the [RIXL](https://rixl.com) API.

[![Go Reference](https://pkg.go.dev/badge/github.com/qeeqez/rixl-sdk-go.svg)](https://pkg.go.dev/github.com/qeeqez/rixl-sdk-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/qeeqez/rixl-sdk-go)](https://goreportcard.com/report/github.com/qeeqez/rixl-sdk-go)

[Installation](#installation) • [Quick start](#quick-start) • [Authentication](#authentication) • [Resources](#resources) • [Pagination](#pagination) • [Errors](#errors)

## Features

- Typed fluent API generated from the RIXL OpenAPI spec
- `context.Context` on every request — deadlines and cancellation built in
- Pre-mapped error responses for 400, 401, 403, 404, and 500
- Pluggable `RequestAdapter` and authentication providers
- Support for JSON, form, multipart, and plain-text payloads

## Requirements

- Go 1.22+
- A RIXL API key

## Installation

```bash
go get github.com/qeeqez/rixl-sdk-go
go get github.com/microsoft/kiota-http-go
go get github.com/microsoft/kiota-abstractions-go
```

## Quick start

```go
package main

import (
    "context"
    "fmt"

    "github.com/microsoft/kiota-abstractions-go/authentication"
    kiotahttp "github.com/microsoft/kiota-http-go"
    "github.com/qeeqez/rixl-sdk-go/sdk"
)

func main() {
    auth, _ := authentication.NewApiKeyAuthenticationProvider(
        "YOUR_RIXL_API_KEY", "X-API-Key", authentication.HEADER_KEYLOCATION,
    )
    adapter, _ := kiotahttp.NewNetHttpRequestAdapter(auth)
    client := sdk.NewRixlClient(adapter)

    image, err := client.Images().ByImageId("PS5IMKoFLm").Get(context.Background(), nil)
    if err != nil {
        panic(err)
    }
    fmt.Println(*image.GetId(), *image.GetWidth(), *image.GetHeight())
}
```

Base URL defaults to `https://api.rixl.com`. Override with `adapter.SetBaseUrl("...")`.

## Authentication

```go
import "github.com/microsoft/kiota-abstractions-go/authentication"

// API key in a header
auth, _ := authentication.NewApiKeyAuthenticationProvider(
    "YOUR_RIXL_API_KEY", "X-API-Key", authentication.HEADER_KEYLOCATION,
)

// Bearer token
// Implement authentication.AccessTokenProvider, then:
// auth, _ := authentication.NewBaseBearerTokenAuthenticationProvider(tokenProvider)
```

## Resources

### Feeds

```go
posts, err := client.Feeds().ByFeedId("FD4y3QB38S").Get(ctx, nil)
for _, post := range posts.GetData() {
    fmt.Println(*post.GetId())
}
```

### Images

```go
// List
page, err := client.Images().Get(ctx, nil)

// Get
image, err := client.Images().ByImageId("PS5IMKoFLm").Get(ctx, nil)

// Delete
err = client.Images().ByImageId("PS5IMKoFLm").Delete(ctx, nil)

// Presigned upload
import imghandler "github.com/qeeqez/rixl-sdk-go/sdk/models/internal_images_handler"

req := imghandler.NewUploadInitRequest()
name, format := "photo.jpg", "jpeg"
req.SetName(&name)
req.SetFormat(&format)

upload, err := client.Images().Upload().Init().Post(ctx, req, nil)
fmt.Println(*upload.GetPresignedUrl())
```

### Videos

```go
// List
videos, err := client.Videos().Get(ctx, nil)

// Get
video, err := client.Videos().ByVideoId("VI9VXQxWXQ").Get(ctx, nil)

// Subtitle tracks
tracks, err := client.Videos().ByVideoId("VI9VXQxWXQ").Subtitles().Get(ctx, nil)
```

## Pagination

List endpoints accept `limit`, `offset`, `sort`, and `order`:

```go
import (
    kabs "github.com/microsoft/kiota-abstractions-go"
    "github.com/qeeqez/rixl-sdk-go/sdk/images"
)

limit, offset := int32(50), int32(0)
for {
    config := &kabs.RequestConfiguration[images.ImagesRequestBuilderGetQueryParameters]{
        QueryParameters: &images.ImagesRequestBuilderGetQueryParameters{
            Limit: &limit, Offset: &offset,
        },
    }
    page, err := client.Images().Get(ctx, config)
    if err != nil {
        return err
    }
    for _, img := range page.GetData() {
        // ...
    }
    if offset+int32(len(page.GetData())) >= *page.GetPagination().GetTotal() {
        break
    }
    offset += limit
}
```

## Errors

API errors (400, 401, 403, 404, 500) are returned as `*ErrorResponse`:

```go
import (
    "errors"
    apierr "github.com/qeeqez/rixl-sdk-go/sdk/models/github_com_qeeqez_api_internal_errors"
)

image, err := client.Images().ByImageId("PS5IMKoFLm").Get(ctx, nil)
if err != nil {
    var apiErr *apierr.ErrorResponse
    if errors.As(err, &apiErr) {
        fmt.Printf("HTTP %d: %s\n", *apiErr.GetCode(), *apiErr.GetError())
    }
    return err
}
```

## Models

Generated types live under `github.com/qeeqez/rixl-sdk-go/sdk/models/`:

| Package | Contents |
|---------|----------|
| `models` | `Imageable`, `Videoable`, `Postable`, `Fileable` |
| `models/pagination` | `PaginatedResponseImageable`, `PaginatedResponseVideoable`, `PaginatedResponsePostable` |
| `models/internal_images_handler` | Upload request and response payloads for images |
| `models/internal_videos_handler` | Upload request and response payloads for videos |
| `models/github_com_qeeqez_api_internal_errors` | `ErrorResponse` |

Getters return pointers; dereference before use.

## Context

Every request takes `context.Context` as its first argument:

```go
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

image, err := client.Images().ByImageId("PS5IMKoFLm").Get(ctx, nil)
```

## Development

Regenerate from the OpenAPI spec (run from the monorepo root):

```bash
brew install kiota
bash sdk-manager/generate.sh rixl-sdk-go
```

Generation uses `--clean-output`; do not hand-edit files under `sdk/`.

## Support

Open an issue at [github.com/qeeqez/rixl-sdk-go](https://github.com/qeeqez/rixl-sdk-go/issues).
