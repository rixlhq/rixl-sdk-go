# \VideosAPI

All URIs are relative to *https://api.rixl.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteVideosVideoIdAudioTracks**](VideosAPI.md#DeleteVideosVideoIdAudioTracks) | **Delete** /videos/{videoId}/audio-tracks | Delete all audio tracks
[**DeleteVideosVideoIdAudioTracksLangCode**](VideosAPI.md#DeleteVideosVideoIdAudioTracksLangCode) | **Delete** /videos/{videoId}/audio-tracks/{lang_code} | Delete audio track by language
[**DeleteVideosVideoIdAudioTracksTrackId**](VideosAPI.md#DeleteVideosVideoIdAudioTracksTrackId) | **Delete** /videos/{videoId}/audio-tracks/{trackId} | Delete audio track
[**DeleteVideosVideoIdChapters**](VideosAPI.md#DeleteVideosVideoIdChapters) | **Delete** /videos/{videoId}/chapters | Delete video chapters
[**DeleteVideosVideoIdDelete**](VideosAPI.md#DeleteVideosVideoIdDelete) | **Delete** /videos/{videoId}/delete | Delete video
[**DeleteVideosVideoIdSubtitles**](VideosAPI.md#DeleteVideosVideoIdSubtitles) | **Delete** /videos/{videoId}/subtitles | Delete all subtitles
[**DeleteVideosVideoIdSubtitlesLangCode**](VideosAPI.md#DeleteVideosVideoIdSubtitlesLangCode) | **Delete** /videos/{videoId}/subtitles/{lang_code} | Delete subtitle by language
[**DeleteVideosVideoIdSubtitlesSubtitleId**](VideosAPI.md#DeleteVideosVideoIdSubtitlesSubtitleId) | **Delete** /videos/{videoId}/subtitles/{subtitleId} | Delete subtitle
[**GetVideos**](VideosAPI.md#GetVideos) | **Get** /videos | List videos for a project
[**GetVideosLanguages**](VideosAPI.md#GetVideosLanguages) | **Get** /videos/languages | List available subtitle languages
[**GetVideosVideoId**](VideosAPI.md#GetVideosVideoId) | **Get** /videos/{videoId} | Get a video
[**PostVideosUploadComplete**](VideosAPI.md#PostVideosUploadComplete) | **Post** /videos/upload/complete | Upload: Mark as complete
[**PostVideosUploadInit**](VideosAPI.md#PostVideosUploadInit) | **Post** /videos/upload/init | Upload: Init
[**PostVideosVideoIdAudioTracks**](VideosAPI.md#PostVideosVideoIdAudioTracks) | **Post** /videos/{videoId}/audio-tracks | Bulk upsert video audio tracks
[**PostVideosVideoIdSubtitles**](VideosAPI.md#PostVideosVideoIdSubtitles) | **Post** /videos/{videoId}/subtitles | Bulk upsert video subtitles
[**PutVideosVideoIdAudioTracksLangCode**](VideosAPI.md#PutVideosVideoIdAudioTracksLangCode) | **Put** /videos/{videoId}/audio-tracks/{lang_code} | Upsert video audio track
[**PutVideosVideoIdChapters**](VideosAPI.md#PutVideosVideoIdChapters) | **Put** /videos/{videoId}/chapters | Update video chapters
[**PutVideosVideoIdSubtitlesLangCode**](VideosAPI.md#PutVideosVideoIdSubtitlesLangCode) | **Put** /videos/{videoId}/subtitles/{lang_code} | Upsert video subtitle
[**PutVideosVideoIdThumbnail**](VideosAPI.md#PutVideosVideoIdThumbnail) | **Put** /videos/{videoId}/thumbnail | Update video thumbnail



## DeleteVideosVideoIdAudioTracks

> AudioTrackDelete DeleteVideosVideoIdAudioTracks(ctx, videoId).Execute()

Delete all audio tracks



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/rixl"
)

func main() {
	videoId := "videoId_example" // string | Video ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VideosAPI.DeleteVideosVideoIdAudioTracks(context.Background(), videoId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideosAPI.DeleteVideosVideoIdAudioTracks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteVideosVideoIdAudioTracks`: AudioTrackDelete
	fmt.Fprintf(os.Stdout, "Response from `VideosAPI.DeleteVideosVideoIdAudioTracks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**videoId** | **string** | Video ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVideosVideoIdAudioTracksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AudioTrackDelete**](AudioTrackDelete.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVideosVideoIdAudioTracksLangCode

> AudioTrackDelete DeleteVideosVideoIdAudioTracksLangCode(ctx, videoId, langCode).Execute()

Delete audio track by language



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/rixl"
)

func main() {
	videoId := "videoId_example" // string | Video ID
	langCode := "langCode_example" // string | Language Code (BCP 47)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VideosAPI.DeleteVideosVideoIdAudioTracksLangCode(context.Background(), videoId, langCode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideosAPI.DeleteVideosVideoIdAudioTracksLangCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteVideosVideoIdAudioTracksLangCode`: AudioTrackDelete
	fmt.Fprintf(os.Stdout, "Response from `VideosAPI.DeleteVideosVideoIdAudioTracksLangCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**videoId** | **string** | Video ID | 
**langCode** | **string** | Language Code (BCP 47) | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVideosVideoIdAudioTracksLangCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AudioTrackDelete**](AudioTrackDelete.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVideosVideoIdAudioTracksTrackId

> AudioTrackDelete DeleteVideosVideoIdAudioTracksTrackId(ctx, videoId, trackId).Execute()

Delete audio track



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/rixl"
)

func main() {
	videoId := "videoId_example" // string | Video ID
	trackId := "trackId_example" // string | Audio Track ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VideosAPI.DeleteVideosVideoIdAudioTracksTrackId(context.Background(), videoId, trackId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideosAPI.DeleteVideosVideoIdAudioTracksTrackId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteVideosVideoIdAudioTracksTrackId`: AudioTrackDelete
	fmt.Fprintf(os.Stdout, "Response from `VideosAPI.DeleteVideosVideoIdAudioTracksTrackId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**videoId** | **string** | Video ID | 
**trackId** | **string** | Audio Track ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVideosVideoIdAudioTracksTrackIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AudioTrackDelete**](AudioTrackDelete.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVideosVideoIdChapters

> UpdateChaptersResponse DeleteVideosVideoIdChapters(ctx, videoId).Execute()

Delete video chapters



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/rixl"
)

func main() {
	videoId := "videoId_example" // string | Video ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VideosAPI.DeleteVideosVideoIdChapters(context.Background(), videoId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideosAPI.DeleteVideosVideoIdChapters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteVideosVideoIdChapters`: UpdateChaptersResponse
	fmt.Fprintf(os.Stdout, "Response from `VideosAPI.DeleteVideosVideoIdChapters`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**videoId** | **string** | Video ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVideosVideoIdChaptersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UpdateChaptersResponse**](UpdateChaptersResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVideosVideoIdDelete

> DeleteVideosVideoIdDelete(ctx, videoId).Execute()

Delete video



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/rixl"
)

func main() {
	videoId := "videoId_example" // string | Video ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VideosAPI.DeleteVideosVideoIdDelete(context.Background(), videoId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideosAPI.DeleteVideosVideoIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**videoId** | **string** | Video ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVideosVideoIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVideosVideoIdSubtitles

> SubtitleDelete DeleteVideosVideoIdSubtitles(ctx, videoId).Execute()

Delete all subtitles



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/rixl"
)

func main() {
	videoId := "videoId_example" // string | Video ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VideosAPI.DeleteVideosVideoIdSubtitles(context.Background(), videoId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideosAPI.DeleteVideosVideoIdSubtitles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteVideosVideoIdSubtitles`: SubtitleDelete
	fmt.Fprintf(os.Stdout, "Response from `VideosAPI.DeleteVideosVideoIdSubtitles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**videoId** | **string** | Video ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVideosVideoIdSubtitlesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SubtitleDelete**](SubtitleDelete.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVideosVideoIdSubtitlesLangCode

> SubtitleDelete DeleteVideosVideoIdSubtitlesLangCode(ctx, videoId, langCode).Execute()

Delete subtitle by language



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/rixl"
)

func main() {
	videoId := "videoId_example" // string | Video ID
	langCode := "langCode_example" // string | Language Code (BCP 47)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VideosAPI.DeleteVideosVideoIdSubtitlesLangCode(context.Background(), videoId, langCode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideosAPI.DeleteVideosVideoIdSubtitlesLangCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteVideosVideoIdSubtitlesLangCode`: SubtitleDelete
	fmt.Fprintf(os.Stdout, "Response from `VideosAPI.DeleteVideosVideoIdSubtitlesLangCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**videoId** | **string** | Video ID | 
**langCode** | **string** | Language Code (BCP 47) | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVideosVideoIdSubtitlesLangCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SubtitleDelete**](SubtitleDelete.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVideosVideoIdSubtitlesSubtitleId

> SubtitleDelete DeleteVideosVideoIdSubtitlesSubtitleId(ctx, videoId, subtitleId).Execute()

Delete subtitle



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/rixl"
)

func main() {
	videoId := "videoId_example" // string | Video ID
	subtitleId := "subtitleId_example" // string | Subtitle ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VideosAPI.DeleteVideosVideoIdSubtitlesSubtitleId(context.Background(), videoId, subtitleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideosAPI.DeleteVideosVideoIdSubtitlesSubtitleId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteVideosVideoIdSubtitlesSubtitleId`: SubtitleDelete
	fmt.Fprintf(os.Stdout, "Response from `VideosAPI.DeleteVideosVideoIdSubtitlesSubtitleId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**videoId** | **string** | Video ID | 
**subtitleId** | **string** | Subtitle ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVideosVideoIdSubtitlesSubtitleIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SubtitleDelete**](SubtitleDelete.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVideos

> PaginationPaginatedResponseVideo GetVideos(ctx).Limit(limit).Offset(offset).Sort(sort).Order(order).Execute()

List videos for a project



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/rixl"
)

func main() {
	limit := int32(56) // int32 | Maximum number of items to return in a single request. <br> **Default:** `25` (optional) (default to 25)
	offset := int32(56) // int32 | Starting point of the result set. <br>To get page 2 with a limit of 25, set `offset` to `25`. <br> **Default:** `0` (optional) (default to 0)
	sort := "sort_example" // string | Field to sort by (created_at, name, size, updated_at, duration) (optional)
	order := "order_example" // string | Sort order (asc, desc) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VideosAPI.GetVideos(context.Background()).Limit(limit).Offset(offset).Sort(sort).Order(order).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideosAPI.GetVideos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVideos`: PaginationPaginatedResponseVideo
	fmt.Fprintf(os.Stdout, "Response from `VideosAPI.GetVideos`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVideosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Maximum number of items to return in a single request. &lt;br&gt; **Default:** &#x60;25&#x60; | [default to 25]
 **offset** | **int32** | Starting point of the result set. &lt;br&gt;To get page 2 with a limit of 25, set &#x60;offset&#x60; to &#x60;25&#x60;. &lt;br&gt; **Default:** &#x60;0&#x60; | [default to 0]
 **sort** | **string** | Field to sort by (created_at, name, size, updated_at, duration) | 
 **order** | **string** | Sort order (asc, desc) | 

### Return type

[**PaginationPaginatedResponseVideo**](PaginationPaginatedResponseVideo.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVideosLanguages

> []InternalVideosHandlerSubtitlesLanguageResponse GetVideosLanguages(ctx).Execute()

List available subtitle languages



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/rixl"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VideosAPI.GetVideosLanguages(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideosAPI.GetVideosLanguages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVideosLanguages`: []InternalVideosHandlerSubtitlesLanguageResponse
	fmt.Fprintf(os.Stdout, "Response from `VideosAPI.GetVideosLanguages`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVideosLanguagesRequest struct via the builder pattern


### Return type

[**[]InternalVideosHandlerSubtitlesLanguageResponse**](InternalVideosHandlerSubtitlesLanguageResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVideosVideoId

> Video GetVideosVideoId(ctx, videoId).Execute()

Get a video



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/rixl"
)

func main() {
	videoId := "videoId_example" // string | Video ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VideosAPI.GetVideosVideoId(context.Background(), videoId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideosAPI.GetVideosVideoId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVideosVideoId`: Video
	fmt.Fprintf(os.Stdout, "Response from `VideosAPI.GetVideosVideoId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**videoId** | **string** | Video ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVideosVideoIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Video**](Video.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostVideosUploadComplete

> Video PostVideosUploadComplete(ctx).Request(request).Execute()

Upload: Mark as complete



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/rixl"
)

func main() {
	request := *openapiclient.NewGithubComQeeqezApiInternalVideosHandlerUploadCompleteRequest() // GithubComQeeqezApiInternalVideosHandlerUploadCompleteRequest | Video upload completion request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VideosAPI.PostVideosUploadComplete(context.Background()).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideosAPI.PostVideosUploadComplete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostVideosUploadComplete`: Video
	fmt.Fprintf(os.Stdout, "Response from `VideosAPI.PostVideosUploadComplete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostVideosUploadCompleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**GithubComQeeqezApiInternalVideosHandlerUploadCompleteRequest**](GithubComQeeqezApiInternalVideosHandlerUploadCompleteRequest.md) | Video upload completion request | 

### Return type

[**Video**](Video.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostVideosUploadInit

> GithubComQeeqezApiInternalVideosHandlerUploadInitResponse PostVideosUploadInit(ctx).Request(request).Execute()

Upload: Init



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/rixl"
)

func main() {
	request := *openapiclient.NewVideoUploadInitRequest("my-video.mp4") // VideoUploadInitRequest | Video upload initialization request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VideosAPI.PostVideosUploadInit(context.Background()).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideosAPI.PostVideosUploadInit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostVideosUploadInit`: GithubComQeeqezApiInternalVideosHandlerUploadInitResponse
	fmt.Fprintf(os.Stdout, "Response from `VideosAPI.PostVideosUploadInit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostVideosUploadInitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**VideoUploadInitRequest**](VideoUploadInitRequest.md) | Video upload initialization request | 

### Return type

[**GithubComQeeqezApiInternalVideosHandlerUploadInitResponse**](GithubComQeeqezApiInternalVideosHandlerUploadInitResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostVideosVideoIdAudioTracks

> []AudioTrack PostVideosVideoIdAudioTracks(ctx, videoId).Files(files).LanguageCodes(languageCodes).Labels(labels).Execute()

Bulk upsert video audio tracks



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/rixl"
)

func main() {
	videoId := "videoId_example" // string | Video ID
	files := []*os.File{"TODO"} // []*os.File | Audio files (.mp3, .opus, .flac, .wav, .ac3, .m4a, .aac)
	languageCodes := "languageCodes_example" // string | Comma-separated language codes
	labels := "labels_example" // string | Comma-separated labels

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VideosAPI.PostVideosVideoIdAudioTracks(context.Background(), videoId).Files(files).LanguageCodes(languageCodes).Labels(labels).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideosAPI.PostVideosVideoIdAudioTracks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostVideosVideoIdAudioTracks`: []AudioTrack
	fmt.Fprintf(os.Stdout, "Response from `VideosAPI.PostVideosVideoIdAudioTracks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**videoId** | **string** | Video ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostVideosVideoIdAudioTracksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **files** | **[]*os.File** | Audio files (.mp3, .opus, .flac, .wav, .ac3, .m4a, .aac) | 
 **languageCodes** | **string** | Comma-separated language codes | 
 **labels** | **string** | Comma-separated labels | 

### Return type

[**[]AudioTrack**](AudioTrack.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostVideosVideoIdSubtitles

> []Subtitle PostVideosVideoIdSubtitles(ctx, videoId).Files(files).LanguageCodes(languageCodes).Labels(labels).Execute()

Bulk upsert video subtitles



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/rixl"
)

func main() {
	videoId := "videoId_example" // string | Video ID
	files := []*os.File{"TODO"} // []*os.File | Subtitle files (.srt or .vtt)
	languageCodes := "languageCodes_example" // string | Comma-separated language codes
	labels := "labels_example" // string | Comma-separated labels

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VideosAPI.PostVideosVideoIdSubtitles(context.Background(), videoId).Files(files).LanguageCodes(languageCodes).Labels(labels).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideosAPI.PostVideosVideoIdSubtitles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostVideosVideoIdSubtitles`: []Subtitle
	fmt.Fprintf(os.Stdout, "Response from `VideosAPI.PostVideosVideoIdSubtitles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**videoId** | **string** | Video ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostVideosVideoIdSubtitlesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **files** | **[]*os.File** | Subtitle files (.srt or .vtt) | 
 **languageCodes** | **string** | Comma-separated language codes | 
 **labels** | **string** | Comma-separated labels | 

### Return type

[**[]Subtitle**](Subtitle.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutVideosVideoIdAudioTracksLangCode

> AudioTrack PutVideosVideoIdAudioTracksLangCode(ctx, videoId, langCode).File(file).Label(label).Execute()

Upsert video audio track



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/rixl"
)

func main() {
	videoId := "videoId_example" // string | Video ID
	langCode := "langCode_example" // string | Language Code (BCP 47)
	file := os.NewFile(1234, "some_file") // *os.File | Audio file (.mp3, .opus, .flac, .wav, .ac3, .m4a, .aac)
	label := "label_example" // string | Label (e.g. English) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VideosAPI.PutVideosVideoIdAudioTracksLangCode(context.Background(), videoId, langCode).File(file).Label(label).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideosAPI.PutVideosVideoIdAudioTracksLangCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutVideosVideoIdAudioTracksLangCode`: AudioTrack
	fmt.Fprintf(os.Stdout, "Response from `VideosAPI.PutVideosVideoIdAudioTracksLangCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**videoId** | **string** | Video ID | 
**langCode** | **string** | Language Code (BCP 47) | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutVideosVideoIdAudioTracksLangCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **file** | ***os.File** | Audio file (.mp3, .opus, .flac, .wav, .ac3, .m4a, .aac) | 
 **label** | **string** | Label (e.g. English) | 

### Return type

[**AudioTrack**](AudioTrack.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutVideosVideoIdChapters

> UpdateChaptersResponse PutVideosVideoIdChapters(ctx, videoId).Request(request).Execute()

Update video chapters



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/rixl"
)

func main() {
	videoId := "videoId_example" // string | Video ID
	request := *openapiclient.NewUpdateChaptersRequest() // UpdateChaptersRequest | Chapters array

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VideosAPI.PutVideosVideoIdChapters(context.Background(), videoId).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideosAPI.PutVideosVideoIdChapters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutVideosVideoIdChapters`: UpdateChaptersResponse
	fmt.Fprintf(os.Stdout, "Response from `VideosAPI.PutVideosVideoIdChapters`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**videoId** | **string** | Video ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutVideosVideoIdChaptersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**UpdateChaptersRequest**](UpdateChaptersRequest.md) | Chapters array | 

### Return type

[**UpdateChaptersResponse**](UpdateChaptersResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutVideosVideoIdSubtitlesLangCode

> Subtitle PutVideosVideoIdSubtitlesLangCode(ctx, videoId, langCode).File(file).Label(label).Execute()

Upsert video subtitle



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/rixl"
)

func main() {
	videoId := "videoId_example" // string | Video ID
	langCode := "langCode_example" // string | Language Code (BCP 47)
	file := os.NewFile(1234, "some_file") // *os.File | Subtitle file (.srt or .vtt)
	label := "label_example" // string | Label (e.g. English) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VideosAPI.PutVideosVideoIdSubtitlesLangCode(context.Background(), videoId, langCode).File(file).Label(label).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideosAPI.PutVideosVideoIdSubtitlesLangCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutVideosVideoIdSubtitlesLangCode`: Subtitle
	fmt.Fprintf(os.Stdout, "Response from `VideosAPI.PutVideosVideoIdSubtitlesLangCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**videoId** | **string** | Video ID | 
**langCode** | **string** | Language Code (BCP 47) | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutVideosVideoIdSubtitlesLangCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **file** | ***os.File** | Subtitle file (.srt or .vtt) | 
 **label** | **string** | Label (e.g. English) | 

### Return type

[**Subtitle**](Subtitle.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutVideosVideoIdThumbnail

> Video PutVideosVideoIdThumbnail(ctx, videoId).Thumbnail(thumbnail).Execute()

Update video thumbnail



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/rixl"
)

func main() {
	videoId := "videoId_example" // string | Video ID
	thumbnail := os.NewFile(1234, "some_file") // *os.File | Thumbnail image file (max 5MB, image/_*)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VideosAPI.PutVideosVideoIdThumbnail(context.Background(), videoId).Thumbnail(thumbnail).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideosAPI.PutVideosVideoIdThumbnail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutVideosVideoIdThumbnail`: Video
	fmt.Fprintf(os.Stdout, "Response from `VideosAPI.PutVideosVideoIdThumbnail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**videoId** | **string** | Video ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutVideosVideoIdThumbnailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **thumbnail** | ***os.File** | Thumbnail image file (max 5MB, image/_*) | 

### Return type

[**Video**](Video.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

