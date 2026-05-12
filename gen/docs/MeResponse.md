# MeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **string** |  | 
**BalanceEurCents** | **int64** | Ledger sum exposed as EUR cents. | 
**AvailableEurCents** | **int64** | balance_eur_cents minus active holds for in-flight async jobs. | 
**ApiUsageTier** | **int32** |  | 
**LifetimePaidCents** | **int64** |  | 
**PaymentCount** | **int32** |  | 
**HasEverPaid** | **bool** |  | 
**EuroTranches** | [**[]EuroTranche**](EuroTranche.md) | Active grants (non-empty, unexpired). Consumed in expires_at ascending order. | 
**SignupTrial** | Pointer to [**NullableMeResponseSignupTrial**](MeResponseSignupTrial.md) |  | [optional] 
**AutoRecharge** | Pointer to [**AutoRechargeState**](AutoRechargeState.md) |  | [optional] 

## Methods

### NewMeResponse

`func NewMeResponse(userId string, balanceEurCents int64, availableEurCents int64, apiUsageTier int32, lifetimePaidCents int64, paymentCount int32, hasEverPaid bool, euroTranches []EuroTranche, ) *MeResponse`

NewMeResponse instantiates a new MeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMeResponseWithDefaults

`func NewMeResponseWithDefaults() *MeResponse`

NewMeResponseWithDefaults instantiates a new MeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *MeResponse) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *MeResponse) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *MeResponse) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetBalanceEurCents

`func (o *MeResponse) GetBalanceEurCents() int64`

GetBalanceEurCents returns the BalanceEurCents field if non-nil, zero value otherwise.

### GetBalanceEurCentsOk

`func (o *MeResponse) GetBalanceEurCentsOk() (*int64, bool)`

GetBalanceEurCentsOk returns a tuple with the BalanceEurCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceEurCents

`func (o *MeResponse) SetBalanceEurCents(v int64)`

SetBalanceEurCents sets BalanceEurCents field to given value.


### GetAvailableEurCents

`func (o *MeResponse) GetAvailableEurCents() int64`

GetAvailableEurCents returns the AvailableEurCents field if non-nil, zero value otherwise.

### GetAvailableEurCentsOk

`func (o *MeResponse) GetAvailableEurCentsOk() (*int64, bool)`

GetAvailableEurCentsOk returns a tuple with the AvailableEurCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableEurCents

`func (o *MeResponse) SetAvailableEurCents(v int64)`

SetAvailableEurCents sets AvailableEurCents field to given value.


### GetApiUsageTier

`func (o *MeResponse) GetApiUsageTier() int32`

GetApiUsageTier returns the ApiUsageTier field if non-nil, zero value otherwise.

### GetApiUsageTierOk

`func (o *MeResponse) GetApiUsageTierOk() (*int32, bool)`

GetApiUsageTierOk returns a tuple with the ApiUsageTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiUsageTier

`func (o *MeResponse) SetApiUsageTier(v int32)`

SetApiUsageTier sets ApiUsageTier field to given value.


### GetLifetimePaidCents

`func (o *MeResponse) GetLifetimePaidCents() int64`

GetLifetimePaidCents returns the LifetimePaidCents field if non-nil, zero value otherwise.

### GetLifetimePaidCentsOk

`func (o *MeResponse) GetLifetimePaidCentsOk() (*int64, bool)`

GetLifetimePaidCentsOk returns a tuple with the LifetimePaidCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifetimePaidCents

`func (o *MeResponse) SetLifetimePaidCents(v int64)`

SetLifetimePaidCents sets LifetimePaidCents field to given value.


### GetPaymentCount

`func (o *MeResponse) GetPaymentCount() int32`

GetPaymentCount returns the PaymentCount field if non-nil, zero value otherwise.

### GetPaymentCountOk

`func (o *MeResponse) GetPaymentCountOk() (*int32, bool)`

GetPaymentCountOk returns a tuple with the PaymentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentCount

`func (o *MeResponse) SetPaymentCount(v int32)`

SetPaymentCount sets PaymentCount field to given value.


### GetHasEverPaid

`func (o *MeResponse) GetHasEverPaid() bool`

GetHasEverPaid returns the HasEverPaid field if non-nil, zero value otherwise.

### GetHasEverPaidOk

`func (o *MeResponse) GetHasEverPaidOk() (*bool, bool)`

GetHasEverPaidOk returns a tuple with the HasEverPaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasEverPaid

`func (o *MeResponse) SetHasEverPaid(v bool)`

SetHasEverPaid sets HasEverPaid field to given value.


### GetEuroTranches

`func (o *MeResponse) GetEuroTranches() []EuroTranche`

GetEuroTranches returns the EuroTranches field if non-nil, zero value otherwise.

### GetEuroTranchesOk

`func (o *MeResponse) GetEuroTranchesOk() (*[]EuroTranche, bool)`

GetEuroTranchesOk returns a tuple with the EuroTranches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEuroTranches

`func (o *MeResponse) SetEuroTranches(v []EuroTranche)`

SetEuroTranches sets EuroTranches field to given value.


### GetSignupTrial

`func (o *MeResponse) GetSignupTrial() MeResponseSignupTrial`

GetSignupTrial returns the SignupTrial field if non-nil, zero value otherwise.

### GetSignupTrialOk

`func (o *MeResponse) GetSignupTrialOk() (*MeResponseSignupTrial, bool)`

GetSignupTrialOk returns a tuple with the SignupTrial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignupTrial

`func (o *MeResponse) SetSignupTrial(v MeResponseSignupTrial)`

SetSignupTrial sets SignupTrial field to given value.

### HasSignupTrial

`func (o *MeResponse) HasSignupTrial() bool`

HasSignupTrial returns a boolean if a field has been set.

### SetSignupTrialNil

`func (o *MeResponse) SetSignupTrialNil(b bool)`

 SetSignupTrialNil sets the value for SignupTrial to be an explicit nil

### UnsetSignupTrial
`func (o *MeResponse) UnsetSignupTrial()`

UnsetSignupTrial ensures that no value is present for SignupTrial, not even an explicit nil
### GetAutoRecharge

`func (o *MeResponse) GetAutoRecharge() AutoRechargeState`

GetAutoRecharge returns the AutoRecharge field if non-nil, zero value otherwise.

### GetAutoRechargeOk

`func (o *MeResponse) GetAutoRechargeOk() (*AutoRechargeState, bool)`

GetAutoRechargeOk returns a tuple with the AutoRecharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRecharge

`func (o *MeResponse) SetAutoRecharge(v AutoRechargeState)`

SetAutoRecharge sets AutoRecharge field to given value.

### HasAutoRecharge

`func (o *MeResponse) HasAutoRecharge() bool`

HasAutoRecharge returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


