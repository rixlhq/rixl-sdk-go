# PaginationPaginatedResponsePost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]Post**](Post.md) | Data contains the slice of items for the current request. | [optional] 
**Pagination** | Pointer to [**PaginationPagination**](PaginationPagination.md) | Pagination data for the request. | [optional] 

## Methods

### NewPaginationPaginatedResponsePost

`func NewPaginationPaginatedResponsePost() *PaginationPaginatedResponsePost`

NewPaginationPaginatedResponsePost instantiates a new PaginationPaginatedResponsePost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginationPaginatedResponsePostWithDefaults

`func NewPaginationPaginatedResponsePostWithDefaults() *PaginationPaginatedResponsePost`

NewPaginationPaginatedResponsePostWithDefaults instantiates a new PaginationPaginatedResponsePost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *PaginationPaginatedResponsePost) GetData() []Post`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PaginationPaginatedResponsePost) GetDataOk() (*[]Post, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PaginationPaginatedResponsePost) SetData(v []Post)`

SetData sets Data field to given value.

### HasData

`func (o *PaginationPaginatedResponsePost) HasData() bool`

HasData returns a boolean if a field has been set.

### GetPagination

`func (o *PaginationPaginatedResponsePost) GetPagination() PaginationPagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *PaginationPaginatedResponsePost) GetPaginationOk() (*PaginationPagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *PaginationPaginatedResponsePost) SetPagination(v PaginationPagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *PaginationPaginatedResponsePost) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


