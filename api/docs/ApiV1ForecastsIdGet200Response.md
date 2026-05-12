# ApiV1ForecastsIdGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Artifacts** | Pointer to [**[]ForecastArtifactMeta**](ForecastArtifactMeta.md) |  | [optional] 
**JobId** | Pointer to **string** |  | [optional] 
**PipelineError** | Pointer to **map[string]interface{}** | Present for some settled failed or canceled jobs when the pipeline wrote &#x60;error.json&#x60; under &#x60;forecasts/{id}/&#x60; or &#x60;forecasts/{id}/output/&#x60; in the artifact bucket. Shape is defined by the pipeline; omitted when missing or unreadable (bodies larger than 64 KiB are omitted). May be a JSON object or array depending on the pipeline.  | [optional] 
**Settled** | Pointer to **bool** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewApiV1ForecastsIdGet200Response

`func NewApiV1ForecastsIdGet200Response() *ApiV1ForecastsIdGet200Response`

NewApiV1ForecastsIdGet200Response instantiates a new ApiV1ForecastsIdGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV1ForecastsIdGet200ResponseWithDefaults

`func NewApiV1ForecastsIdGet200ResponseWithDefaults() *ApiV1ForecastsIdGet200Response`

NewApiV1ForecastsIdGet200ResponseWithDefaults instantiates a new ApiV1ForecastsIdGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArtifacts

`func (o *ApiV1ForecastsIdGet200Response) GetArtifacts() []ForecastArtifactMeta`

GetArtifacts returns the Artifacts field if non-nil, zero value otherwise.

### GetArtifactsOk

`func (o *ApiV1ForecastsIdGet200Response) GetArtifactsOk() (*[]ForecastArtifactMeta, bool)`

GetArtifactsOk returns a tuple with the Artifacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifacts

`func (o *ApiV1ForecastsIdGet200Response) SetArtifacts(v []ForecastArtifactMeta)`

SetArtifacts sets Artifacts field to given value.

### HasArtifacts

`func (o *ApiV1ForecastsIdGet200Response) HasArtifacts() bool`

HasArtifacts returns a boolean if a field has been set.

### GetJobId

`func (o *ApiV1ForecastsIdGet200Response) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *ApiV1ForecastsIdGet200Response) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *ApiV1ForecastsIdGet200Response) SetJobId(v string)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *ApiV1ForecastsIdGet200Response) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetPipelineError

`func (o *ApiV1ForecastsIdGet200Response) GetPipelineError() map[string]interface{}`

GetPipelineError returns the PipelineError field if non-nil, zero value otherwise.

### GetPipelineErrorOk

`func (o *ApiV1ForecastsIdGet200Response) GetPipelineErrorOk() (*map[string]interface{}, bool)`

GetPipelineErrorOk returns a tuple with the PipelineError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipelineError

`func (o *ApiV1ForecastsIdGet200Response) SetPipelineError(v map[string]interface{})`

SetPipelineError sets PipelineError field to given value.

### HasPipelineError

`func (o *ApiV1ForecastsIdGet200Response) HasPipelineError() bool`

HasPipelineError returns a boolean if a field has been set.

### SetPipelineErrorNil

`func (o *ApiV1ForecastsIdGet200Response) SetPipelineErrorNil(b bool)`

 SetPipelineErrorNil sets the value for PipelineError to be an explicit nil

### UnsetPipelineError
`func (o *ApiV1ForecastsIdGet200Response) UnsetPipelineError()`

UnsetPipelineError ensures that no value is present for PipelineError, not even an explicit nil
### GetSettled

`func (o *ApiV1ForecastsIdGet200Response) GetSettled() bool`

GetSettled returns the Settled field if non-nil, zero value otherwise.

### GetSettledOk

`func (o *ApiV1ForecastsIdGet200Response) GetSettledOk() (*bool, bool)`

GetSettledOk returns a tuple with the Settled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettled

`func (o *ApiV1ForecastsIdGet200Response) SetSettled(v bool)`

SetSettled sets Settled field to given value.

### HasSettled

`func (o *ApiV1ForecastsIdGet200Response) HasSettled() bool`

HasSettled returns a boolean if a field has been set.

### GetStatus

`func (o *ApiV1ForecastsIdGet200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApiV1ForecastsIdGet200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApiV1ForecastsIdGet200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApiV1ForecastsIdGet200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


