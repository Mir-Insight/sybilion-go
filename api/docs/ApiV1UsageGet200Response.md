# ApiV1UsageGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | [**Pagination**](Pagination.md) |  | 
**UsageEvents** | [**[]UsageEvent**](UsageEvent.md) |  | 

## Methods

### NewApiV1UsageGet200Response

`func NewApiV1UsageGet200Response(pagination Pagination, usageEvents []UsageEvent, ) *ApiV1UsageGet200Response`

NewApiV1UsageGet200Response instantiates a new ApiV1UsageGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV1UsageGet200ResponseWithDefaults

`func NewApiV1UsageGet200ResponseWithDefaults() *ApiV1UsageGet200Response`

NewApiV1UsageGet200ResponseWithDefaults instantiates a new ApiV1UsageGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *ApiV1UsageGet200Response) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ApiV1UsageGet200Response) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ApiV1UsageGet200Response) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.


### GetUsageEvents

`func (o *ApiV1UsageGet200Response) GetUsageEvents() []UsageEvent`

GetUsageEvents returns the UsageEvents field if non-nil, zero value otherwise.

### GetUsageEventsOk

`func (o *ApiV1UsageGet200Response) GetUsageEventsOk() (*[]UsageEvent, bool)`

GetUsageEventsOk returns a tuple with the UsageEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageEvents

`func (o *ApiV1UsageGet200Response) SetUsageEvents(v []UsageEvent)`

SetUsageEvents sets UsageEvents field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


