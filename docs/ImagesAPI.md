# \ImagesAPI

All URIs are relative to *https://api.rixl.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteImagesImageId**](ImagesAPI.md#DeleteImagesImageId) | **Delete** /images/{imageId} | Delete image
[**GetImages**](ImagesAPI.md#GetImages) | **Get** /images | List images for a project
[**GetImagesImageId**](ImagesAPI.md#GetImagesImageId) | **Get** /images/{imageId} | Get image
[**PostImagesUploadComplete**](ImagesAPI.md#PostImagesUploadComplete) | **Post** /images/upload/complete | Upload: Mark as complete
[**PostImagesUploadInit**](ImagesAPI.md#PostImagesUploadInit) | **Post** /images/upload/init | Upload: Init



## DeleteImagesImageId

> DeleteImagesImageId(ctx, imageId).Execute()

Delete image



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
	imageId := "imageId_example" // string | Image ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ImagesAPI.DeleteImagesImageId(context.Background(), imageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImagesAPI.DeleteImagesImageId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageId** | **string** | Image ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteImagesImageIdRequest struct via the builder pattern


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


## GetImages

> PaginationPaginatedResponseImage GetImages(ctx).Limit(limit).Offset(offset).Sort(sort).Order(order).Execute()

List images for a project



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
	sort := "sort_example" // string | Field to sort by (created_at, name, size, updated_at) (optional)
	order := "order_example" // string | Sort order (asc, desc) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImagesAPI.GetImages(context.Background()).Limit(limit).Offset(offset).Sort(sort).Order(order).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImagesAPI.GetImages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetImages`: PaginationPaginatedResponseImage
	fmt.Fprintf(os.Stdout, "Response from `ImagesAPI.GetImages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetImagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Maximum number of items to return in a single request. &lt;br&gt; **Default:** &#x60;25&#x60; | [default to 25]
 **offset** | **int32** | Starting point of the result set. &lt;br&gt;To get page 2 with a limit of 25, set &#x60;offset&#x60; to &#x60;25&#x60;. &lt;br&gt; **Default:** &#x60;0&#x60; | [default to 0]
 **sort** | **string** | Field to sort by (created_at, name, size, updated_at) | 
 **order** | **string** | Sort order (asc, desc) | 

### Return type

[**PaginationPaginatedResponseImage**](PaginationPaginatedResponseImage.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetImagesImageId

> Image GetImagesImageId(ctx, imageId).Execute()

Get image



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
	imageId := "imageId_example" // string | Image ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImagesAPI.GetImagesImageId(context.Background(), imageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImagesAPI.GetImagesImageId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetImagesImageId`: Image
	fmt.Fprintf(os.Stdout, "Response from `ImagesAPI.GetImagesImageId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageId** | **string** | Image ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetImagesImageIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Image**](Image.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostImagesUploadComplete

> Image PostImagesUploadComplete(ctx).Request(request).Execute()

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
	request := *openapiclient.NewInternalImagesHandlerCompleteRequest() // InternalImagesHandlerCompleteRequest | Upload completion request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImagesAPI.PostImagesUploadComplete(context.Background()).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImagesAPI.PostImagesUploadComplete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostImagesUploadComplete`: Image
	fmt.Fprintf(os.Stdout, "Response from `ImagesAPI.PostImagesUploadComplete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostImagesUploadCompleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**InternalImagesHandlerCompleteRequest**](InternalImagesHandlerCompleteRequest.md) | Upload completion request | 

### Return type

[**Image**](Image.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostImagesUploadInit

> InternalImagesHandlerInitResponse PostImagesUploadInit(ctx).Request(request).Execute()

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
	request := *openapiclient.NewInternalImagesHandlerUploadInitRequest() // InternalImagesHandlerUploadInitRequest | Upload initialization request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImagesAPI.PostImagesUploadInit(context.Background()).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImagesAPI.PostImagesUploadInit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostImagesUploadInit`: InternalImagesHandlerInitResponse
	fmt.Fprintf(os.Stdout, "Response from `ImagesAPI.PostImagesUploadInit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostImagesUploadInitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**InternalImagesHandlerUploadInitRequest**](InternalImagesHandlerUploadInitRequest.md) | Upload initialization request | 

### Return type

[**InternalImagesHandlerInitResponse**](InternalImagesHandlerInitResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

