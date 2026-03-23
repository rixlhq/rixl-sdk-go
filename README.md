# RIXL Go SDK

This repository contains the Go SDK split by service instead of one flat generated client.

## Layout

- `sdk/feeds`
- `sdk/videos`
- `sdk/images`

Each service folder contains its own generated Go package. The root `go.mod` keeps the module path stable for all service imports.

## Import Examples

```go
import feeds "github.com/qeeqez/rixl-sdk-go/sdk/feeds"
import videos "github.com/qeeqez/rixl-sdk-go/sdk/videos"
import images "github.com/qeeqez/rixl-sdk-go/sdk/images"
```

## Regenerate

Generate all services:

```sh
./scripts/generate.sh
```

Generate one service:

```sh
./scripts/generate.sh --service videos
```

Regenerate from a fresh OpenAPI file:

```sh
./scripts/generate.sh --spec /path/to/public.swagger.json --service images
```
