# ForecastAcceptedResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobId** | **string** |  | 
**PollUrl** | **string** |  | 
**Workflow** | Pointer to **string** | Temporal workflow id or dev mock identifier; omitted in some degraded paths | [optional] 
**RunId** | Pointer to **string** | Temporal run id; omitted when Temporal is disabled or not yet assigned | [optional] 

## Methods

### NewForecastAcceptedResponse

`func NewForecastAcceptedResponse(jobId string, pollUrl string, ) *ForecastAcceptedResponse`

NewForecastAcceptedResponse instantiates a new ForecastAcceptedResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewForecastAcceptedResponseWithDefaults

`func NewForecastAcceptedResponseWithDefaults() *ForecastAcceptedResponse`

NewForecastAcceptedResponseWithDefaults instantiates a new ForecastAcceptedResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobId

`func (o *ForecastAcceptedResponse) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *ForecastAcceptedResponse) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *ForecastAcceptedResponse) SetJobId(v string)`

SetJobId sets JobId field to given value.


### GetPollUrl

`func (o *ForecastAcceptedResponse) GetPollUrl() string`

GetPollUrl returns the PollUrl field if non-nil, zero value otherwise.

### GetPollUrlOk

`func (o *ForecastAcceptedResponse) GetPollUrlOk() (*string, bool)`

GetPollUrlOk returns a tuple with the PollUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPollUrl

`func (o *ForecastAcceptedResponse) SetPollUrl(v string)`

SetPollUrl sets PollUrl field to given value.


### GetWorkflow

`func (o *ForecastAcceptedResponse) GetWorkflow() string`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *ForecastAcceptedResponse) GetWorkflowOk() (*string, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *ForecastAcceptedResponse) SetWorkflow(v string)`

SetWorkflow sets Workflow field to given value.

### HasWorkflow

`func (o *ForecastAcceptedResponse) HasWorkflow() bool`

HasWorkflow returns a boolean if a field has been set.

### GetRunId

`func (o *ForecastAcceptedResponse) GetRunId() string`

GetRunId returns the RunId field if non-nil, zero value otherwise.

### GetRunIdOk

`func (o *ForecastAcceptedResponse) GetRunIdOk() (*string, bool)`

GetRunIdOk returns a tuple with the RunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunId

`func (o *ForecastAcceptedResponse) SetRunId(v string)`

SetRunId sets RunId field to given value.

### HasRunId

`func (o *ForecastAcceptedResponse) HasRunId() bool`

HasRunId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


