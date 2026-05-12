# ApiV1JobsGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Jobs** | [**[]JobSummary**](JobSummary.md) |  | 
**Pagination** | [**JobsPagination**](JobsPagination.md) |  | 

## Methods

### NewApiV1JobsGet200Response

`func NewApiV1JobsGet200Response(jobs []JobSummary, pagination JobsPagination, ) *ApiV1JobsGet200Response`

NewApiV1JobsGet200Response instantiates a new ApiV1JobsGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV1JobsGet200ResponseWithDefaults

`func NewApiV1JobsGet200ResponseWithDefaults() *ApiV1JobsGet200Response`

NewApiV1JobsGet200ResponseWithDefaults instantiates a new ApiV1JobsGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobs

`func (o *ApiV1JobsGet200Response) GetJobs() []JobSummary`

GetJobs returns the Jobs field if non-nil, zero value otherwise.

### GetJobsOk

`func (o *ApiV1JobsGet200Response) GetJobsOk() (*[]JobSummary, bool)`

GetJobsOk returns a tuple with the Jobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobs

`func (o *ApiV1JobsGet200Response) SetJobs(v []JobSummary)`

SetJobs sets Jobs field to given value.


### GetPagination

`func (o *ApiV1JobsGet200Response) GetPagination() JobsPagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ApiV1JobsGet200Response) GetPaginationOk() (*JobsPagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ApiV1JobsGet200Response) SetPagination(v JobsPagination)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


