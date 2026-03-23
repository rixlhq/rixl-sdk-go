# \FeedsAPI

All URIs are relative to *https://api.rixl.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFeedsFeedId**](FeedsAPI.md#GetFeedsFeedId) | **Get** /feeds/{feedId} | List posts in a feed
[**GetFeedsFeedIdCreatorsCreatorId**](FeedsAPI.md#GetFeedsFeedIdCreatorsCreatorId) | **Get** /feeds/{feedId}/creators/{creatorId} | List posts by creator
[**GetFeedsFeedIdPostId**](FeedsAPI.md#GetFeedsFeedIdPostId) | **Get** /feeds/{feedId}/{postId} | Get a post



## GetFeedsFeedId

> PaginationPaginatedResponsePost GetFeedsFeedId(ctx, feedId).Limit(limit).Offset(offset).Execute()

List posts in a feed



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
	feedId := "feedId_example" // string | Feed ID
	limit := int32(56) // int32 | Maximum number of items to return in a single request. <br> **Default:** `25` (optional) (default to 25)
	offset := int32(56) // int32 | Starting point of the result set. <br>To get page 2 with a limit of 25, set `offset` to `25`. <br> **Default:** `0` (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FeedsAPI.GetFeedsFeedId(context.Background(), feedId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeedsAPI.GetFeedsFeedId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFeedsFeedId`: PaginationPaginatedResponsePost
	fmt.Fprintf(os.Stdout, "Response from `FeedsAPI.GetFeedsFeedId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**feedId** | **string** | Feed ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFeedsFeedIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Maximum number of items to return in a single request. &lt;br&gt; **Default:** &#x60;25&#x60; | [default to 25]
 **offset** | **int32** | Starting point of the result set. &lt;br&gt;To get page 2 with a limit of 25, set &#x60;offset&#x60; to &#x60;25&#x60;. &lt;br&gt; **Default:** &#x60;0&#x60; | [default to 0]

### Return type

[**PaginationPaginatedResponsePost**](PaginationPaginatedResponsePost.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFeedsFeedIdCreatorsCreatorId

> PaginationPaginatedResponsePost GetFeedsFeedIdCreatorsCreatorId(ctx, feedId, creatorId).Limit(limit).Offset(offset).Execute()

List posts by creator



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
	feedId := "feedId_example" // string | Feed ID
	creatorId := "creatorId_example" // string | Creator ID
	limit := int32(56) // int32 | Maximum number of items to return in a single request. <br> **Default:** `25` (optional) (default to 25)
	offset := int32(56) // int32 | Starting point of the result set. <br>To get page 2 with a limit of 25, set `offset` to `25`. <br> **Default:** `0` (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FeedsAPI.GetFeedsFeedIdCreatorsCreatorId(context.Background(), feedId, creatorId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeedsAPI.GetFeedsFeedIdCreatorsCreatorId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFeedsFeedIdCreatorsCreatorId`: PaginationPaginatedResponsePost
	fmt.Fprintf(os.Stdout, "Response from `FeedsAPI.GetFeedsFeedIdCreatorsCreatorId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**feedId** | **string** | Feed ID | 
**creatorId** | **string** | Creator ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFeedsFeedIdCreatorsCreatorIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** | Maximum number of items to return in a single request. &lt;br&gt; **Default:** &#x60;25&#x60; | [default to 25]
 **offset** | **int32** | Starting point of the result set. &lt;br&gt;To get page 2 with a limit of 25, set &#x60;offset&#x60; to &#x60;25&#x60;. &lt;br&gt; **Default:** &#x60;0&#x60; | [default to 0]

### Return type

[**PaginationPaginatedResponsePost**](PaginationPaginatedResponsePost.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFeedsFeedIdPostId

> Post GetFeedsFeedIdPostId(ctx, feedId, postId).Execute()

Get a post



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
	feedId := "feedId_example" // string | Feed ID
	postId := "postId_example" // string | Post ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FeedsAPI.GetFeedsFeedIdPostId(context.Background(), feedId, postId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeedsAPI.GetFeedsFeedIdPostId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFeedsFeedIdPostId`: Post
	fmt.Fprintf(os.Stdout, "Response from `FeedsAPI.GetFeedsFeedIdPostId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**feedId** | **string** | Feed ID | 
**postId** | **string** | Post ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFeedsFeedIdPostIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Post**](Post.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

