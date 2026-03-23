# Video

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bitrate** | Pointer to **int32** |  | [optional] 
**Chapters** | Pointer to [**[]Chapter**](Chapter.md) |  | [optional] 
**Codec** | Pointer to **string** |  | [optional] 
**Duration** | Pointer to **float32** |  | [optional] 
**File** | Pointer to [**File**](File.md) |  | [optional] 
**Framerate** | Pointer to **string** |  | [optional] 
**Hdr** | Pointer to **bool** |  | [optional] 
**Height** | Pointer to **int32** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**PlanType** | Pointer to [**GithubComQeeqezApiDbSqlcPlanType**](GithubComQeeqezApiDbSqlcPlanType.md) |  | [optional] 
**Poster** | Pointer to [**Image**](Image.md) |  | [optional] 
**Width** | Pointer to **int32** |  | [optional] 

## Methods

### NewVideo

`func NewVideo() *Video`

NewVideo instantiates a new Video object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideoWithDefaults

`func NewVideoWithDefaults() *Video`

NewVideoWithDefaults instantiates a new Video object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBitrate

`func (o *Video) GetBitrate() int32`

GetBitrate returns the Bitrate field if non-nil, zero value otherwise.

### GetBitrateOk

`func (o *Video) GetBitrateOk() (*int32, bool)`

GetBitrateOk returns a tuple with the Bitrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBitrate

`func (o *Video) SetBitrate(v int32)`

SetBitrate sets Bitrate field to given value.

### HasBitrate

`func (o *Video) HasBitrate() bool`

HasBitrate returns a boolean if a field has been set.

### GetChapters

`func (o *Video) GetChapters() []Chapter`

GetChapters returns the Chapters field if non-nil, zero value otherwise.

### GetChaptersOk

`func (o *Video) GetChaptersOk() (*[]Chapter, bool)`

GetChaptersOk returns a tuple with the Chapters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChapters

`func (o *Video) SetChapters(v []Chapter)`

SetChapters sets Chapters field to given value.

### HasChapters

`func (o *Video) HasChapters() bool`

HasChapters returns a boolean if a field has been set.

### GetCodec

`func (o *Video) GetCodec() string`

GetCodec returns the Codec field if non-nil, zero value otherwise.

### GetCodecOk

`func (o *Video) GetCodecOk() (*string, bool)`

GetCodecOk returns a tuple with the Codec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodec

`func (o *Video) SetCodec(v string)`

SetCodec sets Codec field to given value.

### HasCodec

`func (o *Video) HasCodec() bool`

HasCodec returns a boolean if a field has been set.

### GetDuration

`func (o *Video) GetDuration() float32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *Video) GetDurationOk() (*float32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *Video) SetDuration(v float32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *Video) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetFile

`func (o *Video) GetFile() File`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *Video) GetFileOk() (*File, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *Video) SetFile(v File)`

SetFile sets File field to given value.

### HasFile

`func (o *Video) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetFramerate

`func (o *Video) GetFramerate() string`

GetFramerate returns the Framerate field if non-nil, zero value otherwise.

### GetFramerateOk

`func (o *Video) GetFramerateOk() (*string, bool)`

GetFramerateOk returns a tuple with the Framerate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFramerate

`func (o *Video) SetFramerate(v string)`

SetFramerate sets Framerate field to given value.

### HasFramerate

`func (o *Video) HasFramerate() bool`

HasFramerate returns a boolean if a field has been set.

### GetHdr

`func (o *Video) GetHdr() bool`

GetHdr returns the Hdr field if non-nil, zero value otherwise.

### GetHdrOk

`func (o *Video) GetHdrOk() (*bool, bool)`

GetHdrOk returns a tuple with the Hdr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHdr

`func (o *Video) SetHdr(v bool)`

SetHdr sets Hdr field to given value.

### HasHdr

`func (o *Video) HasHdr() bool`

HasHdr returns a boolean if a field has been set.

### GetHeight

`func (o *Video) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *Video) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *Video) SetHeight(v int32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *Video) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetId

`func (o *Video) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Video) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Video) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Video) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPlanType

`func (o *Video) GetPlanType() GithubComQeeqezApiDbSqlcPlanType`

GetPlanType returns the PlanType field if non-nil, zero value otherwise.

### GetPlanTypeOk

`func (o *Video) GetPlanTypeOk() (*GithubComQeeqezApiDbSqlcPlanType, bool)`

GetPlanTypeOk returns a tuple with the PlanType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanType

`func (o *Video) SetPlanType(v GithubComQeeqezApiDbSqlcPlanType)`

SetPlanType sets PlanType field to given value.

### HasPlanType

`func (o *Video) HasPlanType() bool`

HasPlanType returns a boolean if a field has been set.

### GetPoster

`func (o *Video) GetPoster() Image`

GetPoster returns the Poster field if non-nil, zero value otherwise.

### GetPosterOk

`func (o *Video) GetPosterOk() (*Image, bool)`

GetPosterOk returns a tuple with the Poster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoster

`func (o *Video) SetPoster(v Image)`

SetPoster sets Poster field to given value.

### HasPoster

`func (o *Video) HasPoster() bool`

HasPoster returns a boolean if a field has been set.

### GetWidth

`func (o *Video) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *Video) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *Video) SetWidth(v int32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *Video) HasWidth() bool`

HasWidth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


