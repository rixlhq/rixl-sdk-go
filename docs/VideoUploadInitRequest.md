# VideoUploadInitRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileName** | **string** |  | 
**ImageFormat** | Pointer to **string** |  | [optional] 
**VideoQuality** | Pointer to [**GithubComQeeqezApiDbSqlcVideoQuality**](GithubComQeeqezApiDbSqlcVideoQuality.md) |  | [optional] 

## Methods

### NewVideoUploadInitRequest

`func NewVideoUploadInitRequest(fileName string, ) *VideoUploadInitRequest`

NewVideoUploadInitRequest instantiates a new VideoUploadInitRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideoUploadInitRequestWithDefaults

`func NewVideoUploadInitRequestWithDefaults() *VideoUploadInitRequest`

NewVideoUploadInitRequestWithDefaults instantiates a new VideoUploadInitRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileName

`func (o *VideoUploadInitRequest) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *VideoUploadInitRequest) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *VideoUploadInitRequest) SetFileName(v string)`

SetFileName sets FileName field to given value.


### GetImageFormat

`func (o *VideoUploadInitRequest) GetImageFormat() string`

GetImageFormat returns the ImageFormat field if non-nil, zero value otherwise.

### GetImageFormatOk

`func (o *VideoUploadInitRequest) GetImageFormatOk() (*string, bool)`

GetImageFormatOk returns a tuple with the ImageFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageFormat

`func (o *VideoUploadInitRequest) SetImageFormat(v string)`

SetImageFormat sets ImageFormat field to given value.

### HasImageFormat

`func (o *VideoUploadInitRequest) HasImageFormat() bool`

HasImageFormat returns a boolean if a field has been set.

### GetVideoQuality

`func (o *VideoUploadInitRequest) GetVideoQuality() GithubComQeeqezApiDbSqlcVideoQuality`

GetVideoQuality returns the VideoQuality field if non-nil, zero value otherwise.

### GetVideoQualityOk

`func (o *VideoUploadInitRequest) GetVideoQualityOk() (*GithubComQeeqezApiDbSqlcVideoQuality, bool)`

GetVideoQualityOk returns a tuple with the VideoQuality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoQuality

`func (o *VideoUploadInitRequest) SetVideoQuality(v GithubComQeeqezApiDbSqlcVideoQuality)`

SetVideoQuality sets VideoQuality field to given value.

### HasVideoQuality

`func (o *VideoUploadInitRequest) HasVideoQuality() bool`

HasVideoQuality returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


