# RIXL Go SDK

Welcome to the official RIXL Go SDK. This package provides a service-oriented interface for interacting with the RIXL API, organized into discrete modules for Feeds, Videos, and Images.

## Packages

The SDK is structured into service-specific packages to minimize dependency overhead:

- **Feeds**: `github.com/qeeqez/rixl-sdk-go/sdk/feeds` — Access and manage community feeds.
- **Videos**: `github.com/qeeqez/rixl-sdk-go/sdk/videos` — High-performance video processing and retrieval.
- **Images**: `github.com/qeeqez/rixl-sdk-go/sdk/images` — Image upload and transformation services.

## Installation

Install the desired service packages using `go get`:

```bash
go get github.com/qeeqez/rixl-sdk-go/sdk/feeds
go get github.com/qeeqez/rixl-sdk-go/sdk/videos
go get github.com/qeeqez/rixl-sdk-go/sdk/images
```

## Basic Usage

Individual services can be imported and instantiated as needed:

```go
import (
    "github.com/qeeqez/rixl-sdk-go/sdk/feeds"
    "github.com/qeeqez/rixl-sdk-go/sdk/videos"
)

func main() {
    // feedsClient := feeds.NewAPIClient(feeds.NewConfiguration())
    // videosClient := videos.NewAPIClient(videos.NewConfiguration())
}
```

## Support

For issues, bug reports, or feature requests, please open an issue in the official GitHub repository.
