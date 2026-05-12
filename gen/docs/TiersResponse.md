# TiersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentTier** | **int32** | The caller&#39;s current api_usage_tier. | 
**NextTier** | **NullableInt32** | Lowest level greater than current_tier in the ladder, or null when already at the top. | 
**Progress** | [**TiersResponseProgress**](TiersResponseProgress.md) |  | 
**Tiers** | [**[]TierConfig**](TierConfig.md) | Configured tier ladder, sorted ascending by level. | 

## Methods

### NewTiersResponse

`func NewTiersResponse(currentTier int32, nextTier NullableInt32, progress TiersResponseProgress, tiers []TierConfig, ) *TiersResponse`

NewTiersResponse instantiates a new TiersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTiersResponseWithDefaults

`func NewTiersResponseWithDefaults() *TiersResponse`

NewTiersResponseWithDefaults instantiates a new TiersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentTier

`func (o *TiersResponse) GetCurrentTier() int32`

GetCurrentTier returns the CurrentTier field if non-nil, zero value otherwise.

### GetCurrentTierOk

`func (o *TiersResponse) GetCurrentTierOk() (*int32, bool)`

GetCurrentTierOk returns a tuple with the CurrentTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentTier

`func (o *TiersResponse) SetCurrentTier(v int32)`

SetCurrentTier sets CurrentTier field to given value.


### GetNextTier

`func (o *TiersResponse) GetNextTier() int32`

GetNextTier returns the NextTier field if non-nil, zero value otherwise.

### GetNextTierOk

`func (o *TiersResponse) GetNextTierOk() (*int32, bool)`

GetNextTierOk returns a tuple with the NextTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextTier

`func (o *TiersResponse) SetNextTier(v int32)`

SetNextTier sets NextTier field to given value.


### SetNextTierNil

`func (o *TiersResponse) SetNextTierNil(b bool)`

 SetNextTierNil sets the value for NextTier to be an explicit nil

### UnsetNextTier
`func (o *TiersResponse) UnsetNextTier()`

UnsetNextTier ensures that no value is present for NextTier, not even an explicit nil
### GetProgress

`func (o *TiersResponse) GetProgress() TiersResponseProgress`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *TiersResponse) GetProgressOk() (*TiersResponseProgress, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *TiersResponse) SetProgress(v TiersResponseProgress)`

SetProgress sets Progress field to given value.


### GetTiers

`func (o *TiersResponse) GetTiers() []TierConfig`

GetTiers returns the Tiers field if non-nil, zero value otherwise.

### GetTiersOk

`func (o *TiersResponse) GetTiersOk() (*[]TierConfig, bool)`

GetTiersOk returns a tuple with the Tiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiers

`func (o *TiersResponse) SetTiers(v []TierConfig)`

SetTiers sets Tiers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


