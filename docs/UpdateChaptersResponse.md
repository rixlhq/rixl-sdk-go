# UpdateChaptersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Chapters** | Pointer to [**[]Chapter**](Chapter.md) |  | [optional] 
**VideoId** | Pointer to **string** |  | [optional] 

## Methods

### NewUpdateChaptersResponse

`func NewUpdateChaptersResponse() *UpdateChaptersResponse`

NewUpdateChaptersResponse instantiates a new UpdateChaptersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateChaptersResponseWithDefaults

`func NewUpdateChaptersResponseWithDefaults() *UpdateChaptersResponse`

NewUpdateChaptersResponseWithDefaults instantiates a new UpdateChaptersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChapters

`func (o *UpdateChaptersResponse) GetChapters() []Chapter`

GetChapters returns the Chapters field if non-nil, zero value otherwise.

### GetChaptersOk

`func (o *UpdateChaptersResponse) GetChaptersOk() (*[]Chapter, bool)`

GetChaptersOk returns a tuple with the Chapters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChapters

`func (o *UpdateChaptersResponse) SetChapters(v []Chapter)`

SetChapters sets Chapters field to given value.

### HasChapters

`func (o *UpdateChaptersResponse) HasChapters() bool`

HasChapters returns a boolean if a field has been set.

### GetVideoId

`func (o *UpdateChaptersResponse) GetVideoId() string`

GetVideoId returns the VideoId field if non-nil, zero value otherwise.

### GetVideoIdOk

`func (o *UpdateChaptersResponse) GetVideoIdOk() (*string, bool)`

GetVideoIdOk returns a tuple with the VideoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoId

`func (o *UpdateChaptersResponse) SetVideoId(v string)`

SetVideoId sets VideoId field to given value.

### HasVideoId

`func (o *UpdateChaptersResponse) HasVideoId() bool`

HasVideoId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


