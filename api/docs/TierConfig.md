# TierConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GeneralRpm** | **int32** | Per-minute cap on general authenticated /api/v1/_* requests at this tier. | 
**Level** | **int32** | Tier level identifier (0 is the baseline). | 
**MaxConcurrentJobs** | **int32** | Maximum in-flight async forecast jobs (queued + running) at this tier. | 
**MinLifetimePaidCents** | **int64** | Minimum lifetime payment in EUR cents to reach this tier. | 
**MinPaymentCount** | **int32** | Minimum number of successful payments to reach this tier. | 
**MinSinglePaymentCents** | **int64** | Minimum single payment in EUR cents to reach this tier; 0 means no constraint. | 
**SyncBilledRpm** | **int32** | Per-minute cap on synchronous billed endpoints (e.g. POST /drivers) at this tier. | 

## Methods

### NewTierConfig

`func NewTierConfig(generalRpm int32, level int32, maxConcurrentJobs int32, minLifetimePaidCents int64, minPaymentCount int32, minSinglePaymentCents int64, syncBilledRpm int32, ) *TierConfig`

NewTierConfig instantiates a new TierConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTierConfigWithDefaults

`func NewTierConfigWithDefaults() *TierConfig`

NewTierConfigWithDefaults instantiates a new TierConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGeneralRpm

`func (o *TierConfig) GetGeneralRpm() int32`

GetGeneralRpm returns the GeneralRpm field if non-nil, zero value otherwise.

### GetGeneralRpmOk

`func (o *TierConfig) GetGeneralRpmOk() (*int32, bool)`

GetGeneralRpmOk returns a tuple with the GeneralRpm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneralRpm

`func (o *TierConfig) SetGeneralRpm(v int32)`

SetGeneralRpm sets GeneralRpm field to given value.


### GetLevel

`func (o *TierConfig) GetLevel() int32`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *TierConfig) GetLevelOk() (*int32, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *TierConfig) SetLevel(v int32)`

SetLevel sets Level field to given value.


### GetMaxConcurrentJobs

`func (o *TierConfig) GetMaxConcurrentJobs() int32`

GetMaxConcurrentJobs returns the MaxConcurrentJobs field if non-nil, zero value otherwise.

### GetMaxConcurrentJobsOk

`func (o *TierConfig) GetMaxConcurrentJobsOk() (*int32, bool)`

GetMaxConcurrentJobsOk returns a tuple with the MaxConcurrentJobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConcurrentJobs

`func (o *TierConfig) SetMaxConcurrentJobs(v int32)`

SetMaxConcurrentJobs sets MaxConcurrentJobs field to given value.


### GetMinLifetimePaidCents

`func (o *TierConfig) GetMinLifetimePaidCents() int64`

GetMinLifetimePaidCents returns the MinLifetimePaidCents field if non-nil, zero value otherwise.

### GetMinLifetimePaidCentsOk

`func (o *TierConfig) GetMinLifetimePaidCentsOk() (*int64, bool)`

GetMinLifetimePaidCentsOk returns a tuple with the MinLifetimePaidCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinLifetimePaidCents

`func (o *TierConfig) SetMinLifetimePaidCents(v int64)`

SetMinLifetimePaidCents sets MinLifetimePaidCents field to given value.


### GetMinPaymentCount

`func (o *TierConfig) GetMinPaymentCount() int32`

GetMinPaymentCount returns the MinPaymentCount field if non-nil, zero value otherwise.

### GetMinPaymentCountOk

`func (o *TierConfig) GetMinPaymentCountOk() (*int32, bool)`

GetMinPaymentCountOk returns a tuple with the MinPaymentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinPaymentCount

`func (o *TierConfig) SetMinPaymentCount(v int32)`

SetMinPaymentCount sets MinPaymentCount field to given value.


### GetMinSinglePaymentCents

`func (o *TierConfig) GetMinSinglePaymentCents() int64`

GetMinSinglePaymentCents returns the MinSinglePaymentCents field if non-nil, zero value otherwise.

### GetMinSinglePaymentCentsOk

`func (o *TierConfig) GetMinSinglePaymentCentsOk() (*int64, bool)`

GetMinSinglePaymentCentsOk returns a tuple with the MinSinglePaymentCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinSinglePaymentCents

`func (o *TierConfig) SetMinSinglePaymentCents(v int64)`

SetMinSinglePaymentCents sets MinSinglePaymentCents field to given value.


### GetSyncBilledRpm

`func (o *TierConfig) GetSyncBilledRpm() int32`

GetSyncBilledRpm returns the SyncBilledRpm field if non-nil, zero value otherwise.

### GetSyncBilledRpmOk

`func (o *TierConfig) GetSyncBilledRpmOk() (*int32, bool)`

GetSyncBilledRpmOk returns a tuple with the SyncBilledRpm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncBilledRpm

`func (o *TierConfig) SetSyncBilledRpm(v int32)`

SetSyncBilledRpm sets SyncBilledRpm field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


