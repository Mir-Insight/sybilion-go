# ForecastRequestV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PipelineVersion** | **string** | Pipeline version. Closed set; no aliases or \&quot;latest\&quot; resolution. Only v1 is supported today. | 
**Horizon** | **int32** | Number of forecast steps to produce. | 
**Frequency** | **string** | Series cadence. Only \&quot;monthly\&quot; is currently supported; \&quot;daily\&quot; and \&quot;weekly\&quot; are reserved. | 
**Backtest** | Pointer to **bool** | When true, run a backtest evaluation alongside the forecast. | [optional] 
**RecencyFactor** | **float64** |  | 
**StrictlyPositive** | Pointer to **bool** | When true, every value in &#x60;timeseries&#x60; must be &#x60;&gt;&#x3D; 0&#x60; (zero is allowed); a single negative observation rejects the request with 422. The downstream forecasting pipeline (PPL) also clamps the produced forecast at zero so no output point can be negative. Defaults to false, in which case no positivity constraint is applied to inputs or outputs and negative values are returned unchanged. Optional.  | [optional] [default to false]
**TimeseriesMetadata** | [**TimeseriesMetadata**](TimeseriesMetadata.md) |  | 
**Filters** | Pointer to [**Filters**](Filters.md) |  | [optional] 
**Timeseries** | **map[string]float32** | Map of YYYY-MM-DD date keys to numeric observation values. Must contain at least 60 points (5 years of monthly history) and be aligned to the declared frequency. | 

## Methods

### NewForecastRequestV1

`func NewForecastRequestV1(pipelineVersion string, horizon int32, frequency string, recencyFactor float64, timeseriesMetadata TimeseriesMetadata, timeseries map[string]float32, ) *ForecastRequestV1`

NewForecastRequestV1 instantiates a new ForecastRequestV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewForecastRequestV1WithDefaults

`func NewForecastRequestV1WithDefaults() *ForecastRequestV1`

NewForecastRequestV1WithDefaults instantiates a new ForecastRequestV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPipelineVersion

`func (o *ForecastRequestV1) GetPipelineVersion() string`

GetPipelineVersion returns the PipelineVersion field if non-nil, zero value otherwise.

### GetPipelineVersionOk

`func (o *ForecastRequestV1) GetPipelineVersionOk() (*string, bool)`

GetPipelineVersionOk returns a tuple with the PipelineVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipelineVersion

`func (o *ForecastRequestV1) SetPipelineVersion(v string)`

SetPipelineVersion sets PipelineVersion field to given value.


### GetHorizon

`func (o *ForecastRequestV1) GetHorizon() int32`

GetHorizon returns the Horizon field if non-nil, zero value otherwise.

### GetHorizonOk

`func (o *ForecastRequestV1) GetHorizonOk() (*int32, bool)`

GetHorizonOk returns a tuple with the Horizon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHorizon

`func (o *ForecastRequestV1) SetHorizon(v int32)`

SetHorizon sets Horizon field to given value.


### GetFrequency

`func (o *ForecastRequestV1) GetFrequency() string`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *ForecastRequestV1) GetFrequencyOk() (*string, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *ForecastRequestV1) SetFrequency(v string)`

SetFrequency sets Frequency field to given value.


### GetBacktest

`func (o *ForecastRequestV1) GetBacktest() bool`

GetBacktest returns the Backtest field if non-nil, zero value otherwise.

### GetBacktestOk

`func (o *ForecastRequestV1) GetBacktestOk() (*bool, bool)`

GetBacktestOk returns a tuple with the Backtest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBacktest

`func (o *ForecastRequestV1) SetBacktest(v bool)`

SetBacktest sets Backtest field to given value.

### HasBacktest

`func (o *ForecastRequestV1) HasBacktest() bool`

HasBacktest returns a boolean if a field has been set.

### GetRecencyFactor

`func (o *ForecastRequestV1) GetRecencyFactor() float64`

GetRecencyFactor returns the RecencyFactor field if non-nil, zero value otherwise.

### GetRecencyFactorOk

`func (o *ForecastRequestV1) GetRecencyFactorOk() (*float64, bool)`

GetRecencyFactorOk returns a tuple with the RecencyFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecencyFactor

`func (o *ForecastRequestV1) SetRecencyFactor(v float64)`

SetRecencyFactor sets RecencyFactor field to given value.


### GetStrictlyPositive

`func (o *ForecastRequestV1) GetStrictlyPositive() bool`

GetStrictlyPositive returns the StrictlyPositive field if non-nil, zero value otherwise.

### GetStrictlyPositiveOk

`func (o *ForecastRequestV1) GetStrictlyPositiveOk() (*bool, bool)`

GetStrictlyPositiveOk returns a tuple with the StrictlyPositive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrictlyPositive

`func (o *ForecastRequestV1) SetStrictlyPositive(v bool)`

SetStrictlyPositive sets StrictlyPositive field to given value.

### HasStrictlyPositive

`func (o *ForecastRequestV1) HasStrictlyPositive() bool`

HasStrictlyPositive returns a boolean if a field has been set.

### GetTimeseriesMetadata

`func (o *ForecastRequestV1) GetTimeseriesMetadata() TimeseriesMetadata`

GetTimeseriesMetadata returns the TimeseriesMetadata field if non-nil, zero value otherwise.

### GetTimeseriesMetadataOk

`func (o *ForecastRequestV1) GetTimeseriesMetadataOk() (*TimeseriesMetadata, bool)`

GetTimeseriesMetadataOk returns a tuple with the TimeseriesMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeseriesMetadata

`func (o *ForecastRequestV1) SetTimeseriesMetadata(v TimeseriesMetadata)`

SetTimeseriesMetadata sets TimeseriesMetadata field to given value.


### GetFilters

`func (o *ForecastRequestV1) GetFilters() Filters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *ForecastRequestV1) GetFiltersOk() (*Filters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *ForecastRequestV1) SetFilters(v Filters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *ForecastRequestV1) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetTimeseries

`func (o *ForecastRequestV1) GetTimeseries() map[string]float32`

GetTimeseries returns the Timeseries field if non-nil, zero value otherwise.

### GetTimeseriesOk

`func (o *ForecastRequestV1) GetTimeseriesOk() (*map[string]float32, bool)`

GetTimeseriesOk returns a tuple with the Timeseries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeseries

`func (o *ForecastRequestV1) SetTimeseries(v map[string]float32)`

SetTimeseries sets Timeseries field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


