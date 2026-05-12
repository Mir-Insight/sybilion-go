# ForecastJobResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobId** | **string** |  | 
**Status** | **string** |  | 
**PipelineType** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**SettledAt** | Pointer to **NullableTime** |  | [optional] 
**EurCentsFinal** | Pointer to **NullableInt64** | Final charge in EUR cents when settled; null until then | [optional] 
**Settled** | **bool** |  | 
**WorkflowId** | Pointer to **string** | Present when the job has a Temporal workflow id assigned | [optional] 
**RunId** | Pointer to **string** | Present when the job has a Temporal run id assigned | [optional] 
**TerminalReason** | Pointer to **NullableString** |  | [optional] 
**Artifacts** | Pointer to [**[]ForecastArtifactMeta**](ForecastArtifactMeta.md) | Present only for completed settled jobs with discoverable outputs | [optional] 
**PipelineError** | Pointer to **map[string]interface{}** | Present for some settled failed or canceled jobs when &#x60;error.json&#x60; exists under &#x60;forecasts/{id}/&#x60; or &#x60;forecasts/{id}/output/&#x60; in the artifact bucket. Omitted when missing or unreadable (bodies larger than 64 KiB are skipped). The pipeline may emit a JSON object or array; clients should treat this as opaque diagnostic JSON.  | [optional] 

## Methods

### NewForecastJobResponse

`func NewForecastJobResponse(jobId string, status string, pipelineType string, createdAt time.Time, settled bool, ) *ForecastJobResponse`

NewForecastJobResponse instantiates a new ForecastJobResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewForecastJobResponseWithDefaults

`func NewForecastJobResponseWithDefaults() *ForecastJobResponse`

NewForecastJobResponseWithDefaults instantiates a new ForecastJobResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobId

`func (o *ForecastJobResponse) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *ForecastJobResponse) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *ForecastJobResponse) SetJobId(v string)`

SetJobId sets JobId field to given value.


### GetStatus

`func (o *ForecastJobResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ForecastJobResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ForecastJobResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetPipelineType

`func (o *ForecastJobResponse) GetPipelineType() string`

GetPipelineType returns the PipelineType field if non-nil, zero value otherwise.

### GetPipelineTypeOk

`func (o *ForecastJobResponse) GetPipelineTypeOk() (*string, bool)`

GetPipelineTypeOk returns a tuple with the PipelineType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipelineType

`func (o *ForecastJobResponse) SetPipelineType(v string)`

SetPipelineType sets PipelineType field to given value.


### GetCreatedAt

`func (o *ForecastJobResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ForecastJobResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ForecastJobResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetSettledAt

`func (o *ForecastJobResponse) GetSettledAt() time.Time`

GetSettledAt returns the SettledAt field if non-nil, zero value otherwise.

### GetSettledAtOk

`func (o *ForecastJobResponse) GetSettledAtOk() (*time.Time, bool)`

GetSettledAtOk returns a tuple with the SettledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettledAt

`func (o *ForecastJobResponse) SetSettledAt(v time.Time)`

SetSettledAt sets SettledAt field to given value.

### HasSettledAt

`func (o *ForecastJobResponse) HasSettledAt() bool`

HasSettledAt returns a boolean if a field has been set.

### SetSettledAtNil

`func (o *ForecastJobResponse) SetSettledAtNil(b bool)`

 SetSettledAtNil sets the value for SettledAt to be an explicit nil

### UnsetSettledAt
`func (o *ForecastJobResponse) UnsetSettledAt()`

UnsetSettledAt ensures that no value is present for SettledAt, not even an explicit nil
### GetEurCentsFinal

`func (o *ForecastJobResponse) GetEurCentsFinal() int64`

GetEurCentsFinal returns the EurCentsFinal field if non-nil, zero value otherwise.

### GetEurCentsFinalOk

`func (o *ForecastJobResponse) GetEurCentsFinalOk() (*int64, bool)`

GetEurCentsFinalOk returns a tuple with the EurCentsFinal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEurCentsFinal

`func (o *ForecastJobResponse) SetEurCentsFinal(v int64)`

SetEurCentsFinal sets EurCentsFinal field to given value.

### HasEurCentsFinal

`func (o *ForecastJobResponse) HasEurCentsFinal() bool`

HasEurCentsFinal returns a boolean if a field has been set.

### SetEurCentsFinalNil

`func (o *ForecastJobResponse) SetEurCentsFinalNil(b bool)`

 SetEurCentsFinalNil sets the value for EurCentsFinal to be an explicit nil

### UnsetEurCentsFinal
`func (o *ForecastJobResponse) UnsetEurCentsFinal()`

UnsetEurCentsFinal ensures that no value is present for EurCentsFinal, not even an explicit nil
### GetSettled

`func (o *ForecastJobResponse) GetSettled() bool`

GetSettled returns the Settled field if non-nil, zero value otherwise.

### GetSettledOk

`func (o *ForecastJobResponse) GetSettledOk() (*bool, bool)`

GetSettledOk returns a tuple with the Settled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettled

`func (o *ForecastJobResponse) SetSettled(v bool)`

SetSettled sets Settled field to given value.


### GetWorkflowId

`func (o *ForecastJobResponse) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *ForecastJobResponse) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *ForecastJobResponse) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.

### HasWorkflowId

`func (o *ForecastJobResponse) HasWorkflowId() bool`

HasWorkflowId returns a boolean if a field has been set.

### GetRunId

`func (o *ForecastJobResponse) GetRunId() string`

GetRunId returns the RunId field if non-nil, zero value otherwise.

### GetRunIdOk

`func (o *ForecastJobResponse) GetRunIdOk() (*string, bool)`

GetRunIdOk returns a tuple with the RunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunId

`func (o *ForecastJobResponse) SetRunId(v string)`

SetRunId sets RunId field to given value.

### HasRunId

`func (o *ForecastJobResponse) HasRunId() bool`

HasRunId returns a boolean if a field has been set.

### GetTerminalReason

`func (o *ForecastJobResponse) GetTerminalReason() string`

GetTerminalReason returns the TerminalReason field if non-nil, zero value otherwise.

### GetTerminalReasonOk

`func (o *ForecastJobResponse) GetTerminalReasonOk() (*string, bool)`

GetTerminalReasonOk returns a tuple with the TerminalReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalReason

`func (o *ForecastJobResponse) SetTerminalReason(v string)`

SetTerminalReason sets TerminalReason field to given value.

### HasTerminalReason

`func (o *ForecastJobResponse) HasTerminalReason() bool`

HasTerminalReason returns a boolean if a field has been set.

### SetTerminalReasonNil

`func (o *ForecastJobResponse) SetTerminalReasonNil(b bool)`

 SetTerminalReasonNil sets the value for TerminalReason to be an explicit nil

### UnsetTerminalReason
`func (o *ForecastJobResponse) UnsetTerminalReason()`

UnsetTerminalReason ensures that no value is present for TerminalReason, not even an explicit nil
### GetArtifacts

`func (o *ForecastJobResponse) GetArtifacts() []ForecastArtifactMeta`

GetArtifacts returns the Artifacts field if non-nil, zero value otherwise.

### GetArtifactsOk

`func (o *ForecastJobResponse) GetArtifactsOk() (*[]ForecastArtifactMeta, bool)`

GetArtifactsOk returns a tuple with the Artifacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifacts

`func (o *ForecastJobResponse) SetArtifacts(v []ForecastArtifactMeta)`

SetArtifacts sets Artifacts field to given value.

### HasArtifacts

`func (o *ForecastJobResponse) HasArtifacts() bool`

HasArtifacts returns a boolean if a field has been set.

### GetPipelineError

`func (o *ForecastJobResponse) GetPipelineError() map[string]interface{}`

GetPipelineError returns the PipelineError field if non-nil, zero value otherwise.

### GetPipelineErrorOk

`func (o *ForecastJobResponse) GetPipelineErrorOk() (*map[string]interface{}, bool)`

GetPipelineErrorOk returns a tuple with the PipelineError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipelineError

`func (o *ForecastJobResponse) SetPipelineError(v map[string]interface{})`

SetPipelineError sets PipelineError field to given value.

### HasPipelineError

`func (o *ForecastJobResponse) HasPipelineError() bool`

HasPipelineError returns a boolean if a field has been set.

### SetPipelineErrorNil

`func (o *ForecastJobResponse) SetPipelineErrorNil(b bool)`

 SetPipelineErrorNil sets the value for PipelineError to be an explicit nil

### UnsetPipelineError
`func (o *ForecastJobResponse) UnsetPipelineError()`

UnsetPipelineError ensures that no value is present for PipelineError, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


