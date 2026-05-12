# AutoRechargePatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BelowEurCents** | **int64** | Trigger threshold in EUR cents (&gt;&#x3D; 1 when enabled). | 
**Enabled** | **bool** |  | 
**MonthlyCapCents** | **int64** |  | 
**TargetEurCents** | **int64** | Target balance in EUR cents (&gt; below_eur_cents when enabled). | 

## Methods

### NewAutoRechargePatch

`func NewAutoRechargePatch(belowEurCents int64, enabled bool, monthlyCapCents int64, targetEurCents int64, ) *AutoRechargePatch`

NewAutoRechargePatch instantiates a new AutoRechargePatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoRechargePatchWithDefaults

`func NewAutoRechargePatchWithDefaults() *AutoRechargePatch`

NewAutoRechargePatchWithDefaults instantiates a new AutoRechargePatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBelowEurCents

`func (o *AutoRechargePatch) GetBelowEurCents() int64`

GetBelowEurCents returns the BelowEurCents field if non-nil, zero value otherwise.

### GetBelowEurCentsOk

`func (o *AutoRechargePatch) GetBelowEurCentsOk() (*int64, bool)`

GetBelowEurCentsOk returns a tuple with the BelowEurCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBelowEurCents

`func (o *AutoRechargePatch) SetBelowEurCents(v int64)`

SetBelowEurCents sets BelowEurCents field to given value.


### GetEnabled

`func (o *AutoRechargePatch) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AutoRechargePatch) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AutoRechargePatch) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetMonthlyCapCents

`func (o *AutoRechargePatch) GetMonthlyCapCents() int64`

GetMonthlyCapCents returns the MonthlyCapCents field if non-nil, zero value otherwise.

### GetMonthlyCapCentsOk

`func (o *AutoRechargePatch) GetMonthlyCapCentsOk() (*int64, bool)`

GetMonthlyCapCentsOk returns a tuple with the MonthlyCapCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyCapCents

`func (o *AutoRechargePatch) SetMonthlyCapCents(v int64)`

SetMonthlyCapCents sets MonthlyCapCents field to given value.


### GetTargetEurCents

`func (o *AutoRechargePatch) GetTargetEurCents() int64`

GetTargetEurCents returns the TargetEurCents field if non-nil, zero value otherwise.

### GetTargetEurCentsOk

`func (o *AutoRechargePatch) GetTargetEurCentsOk() (*int64, bool)`

GetTargetEurCentsOk returns a tuple with the TargetEurCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetEurCents

`func (o *AutoRechargePatch) SetTargetEurCents(v int64)`

SetTargetEurCents sets TargetEurCents field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


