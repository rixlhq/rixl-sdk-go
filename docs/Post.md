# Post

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **string** |  | [optional] 
**CreatorId** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**FeedId** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Image** | Pointer to [**Image**](Image.md) |  | [optional] 
**PlanType** | Pointer to [**GithubComQeeqezApiDbSqlcPlanType**](GithubComQeeqezApiDbSqlcPlanType.md) |  | [optional] 
**Type** | Pointer to [**PostType**](PostType.md) |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**Video** | Pointer to [**GithubComQeeqezApiInternalVideosVideoResponse**](GithubComQeeqezApiInternalVideosVideoResponse.md) |  | [optional] 

## Methods

### NewPost

`func NewPost() *Post`

NewPost instantiates a new Post object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostWithDefaults

`func NewPostWithDefaults() *Post`

NewPostWithDefaults instantiates a new Post object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *Post) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Post) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Post) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Post) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatorId

`func (o *Post) GetCreatorId() string`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *Post) GetCreatorIdOk() (*string, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *Post) SetCreatorId(v string)`

SetCreatorId sets CreatorId field to given value.

### HasCreatorId

`func (o *Post) HasCreatorId() bool`

HasCreatorId returns a boolean if a field has been set.

### GetDescription

`func (o *Post) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Post) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Post) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Post) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFeedId

`func (o *Post) GetFeedId() string`

GetFeedId returns the FeedId field if non-nil, zero value otherwise.

### GetFeedIdOk

`func (o *Post) GetFeedIdOk() (*string, bool)`

GetFeedIdOk returns a tuple with the FeedId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedId

`func (o *Post) SetFeedId(v string)`

SetFeedId sets FeedId field to given value.

### HasFeedId

`func (o *Post) HasFeedId() bool`

HasFeedId returns a boolean if a field has been set.

### GetId

`func (o *Post) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Post) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Post) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Post) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImage

`func (o *Post) GetImage() Image`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *Post) GetImageOk() (*Image, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *Post) SetImage(v Image)`

SetImage sets Image field to given value.

### HasImage

`func (o *Post) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetPlanType

`func (o *Post) GetPlanType() GithubComQeeqezApiDbSqlcPlanType`

GetPlanType returns the PlanType field if non-nil, zero value otherwise.

### GetPlanTypeOk

`func (o *Post) GetPlanTypeOk() (*GithubComQeeqezApiDbSqlcPlanType, bool)`

GetPlanTypeOk returns a tuple with the PlanType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanType

`func (o *Post) SetPlanType(v GithubComQeeqezApiDbSqlcPlanType)`

SetPlanType sets PlanType field to given value.

### HasPlanType

`func (o *Post) HasPlanType() bool`

HasPlanType returns a boolean if a field has been set.

### GetType

`func (o *Post) GetType() PostType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Post) GetTypeOk() (*PostType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Post) SetType(v PostType)`

SetType sets Type field to given value.

### HasType

`func (o *Post) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Post) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Post) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Post) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Post) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetVideo

`func (o *Post) GetVideo() GithubComQeeqezApiInternalVideosVideoResponse`

GetVideo returns the Video field if non-nil, zero value otherwise.

### GetVideoOk

`func (o *Post) GetVideoOk() (*GithubComQeeqezApiInternalVideosVideoResponse, bool)`

GetVideoOk returns a tuple with the Video field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideo

`func (o *Post) SetVideo(v GithubComQeeqezApiInternalVideosVideoResponse)`

SetVideo sets Video field to given value.

### HasVideo

`func (o *Post) HasVideo() bool`

HasVideo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


