# PaginationPagination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | Pointer to **int32** | Maximum number of items to return in a single request. | [optional] 
**Offset** | Pointer to **int32** | Starting point of the result set. | [optional] 
**Total** | Pointer to **int32** | The total number of available items in the full list. | [optional] 

## Methods

### NewPaginationPagination

`func NewPaginationPagination() *PaginationPagination`

NewPaginationPagination instantiates a new PaginationPagination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginationPaginationWithDefaults

`func NewPaginationPaginationWithDefaults() *PaginationPagination`

NewPaginationPaginationWithDefaults instantiates a new PaginationPagination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *PaginationPagination) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *PaginationPagination) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *PaginationPagination) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *PaginationPagination) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *PaginationPagination) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *PaginationPagination) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *PaginationPagination) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *PaginationPagination) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetTotal

`func (o *PaginationPagination) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *PaginationPagination) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *PaginationPagination) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *PaginationPagination) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


