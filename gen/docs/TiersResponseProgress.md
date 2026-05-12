# TiersResponseProgress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LifetimePaidCents** | **int64** |  | 
**PaymentCount** | **int32** |  | 
**MaxSinglePaymentCents** | **int64** | Largest single Stripe payment amount (in cents) observed for the caller; 0 if none. | 

## Methods

### NewTiersResponseProgress

`func NewTiersResponseProgress(lifetimePaidCents int64, paymentCount int32, maxSinglePaymentCents int64, ) *TiersResponseProgress`

NewTiersResponseProgress instantiates a new TiersResponseProgress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTiersResponseProgressWithDefaults

`func NewTiersResponseProgressWithDefaults() *TiersResponseProgress`

NewTiersResponseProgressWithDefaults instantiates a new TiersResponseProgress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLifetimePaidCents

`func (o *TiersResponseProgress) GetLifetimePaidCents() int64`

GetLifetimePaidCents returns the LifetimePaidCents field if non-nil, zero value otherwise.

### GetLifetimePaidCentsOk

`func (o *TiersResponseProgress) GetLifetimePaidCentsOk() (*int64, bool)`

GetLifetimePaidCentsOk returns a tuple with the LifetimePaidCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifetimePaidCents

`func (o *TiersResponseProgress) SetLifetimePaidCents(v int64)`

SetLifetimePaidCents sets LifetimePaidCents field to given value.


### GetPaymentCount

`func (o *TiersResponseProgress) GetPaymentCount() int32`

GetPaymentCount returns the PaymentCount field if non-nil, zero value otherwise.

### GetPaymentCountOk

`func (o *TiersResponseProgress) GetPaymentCountOk() (*int32, bool)`

GetPaymentCountOk returns a tuple with the PaymentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentCount

`func (o *TiersResponseProgress) SetPaymentCount(v int32)`

SetPaymentCount sets PaymentCount field to given value.


### GetMaxSinglePaymentCents

`func (o *TiersResponseProgress) GetMaxSinglePaymentCents() int64`

GetMaxSinglePaymentCents returns the MaxSinglePaymentCents field if non-nil, zero value otherwise.

### GetMaxSinglePaymentCentsOk

`func (o *TiersResponseProgress) GetMaxSinglePaymentCentsOk() (*int64, bool)`

GetMaxSinglePaymentCentsOk returns a tuple with the MaxSinglePaymentCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSinglePaymentCents

`func (o *TiersResponseProgress) SetMaxSinglePaymentCents(v int64)`

SetMaxSinglePaymentCents sets MaxSinglePaymentCents field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


