# CreateAPIKeyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKey** | **string** | Full secret key; shown only on create | 
**Id** | **string** |  | 
**Name** | **string** |  | 
**Prefix** | **string** | Public prefix of the key (same semantics as key_prefix on list) | 
**Warning** | **string** |  | 

## Methods

### NewCreateAPIKeyResponse

`func NewCreateAPIKeyResponse(apiKey string, id string, name string, prefix string, warning string, ) *CreateAPIKeyResponse`

NewCreateAPIKeyResponse instantiates a new CreateAPIKeyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAPIKeyResponseWithDefaults

`func NewCreateAPIKeyResponseWithDefaults() *CreateAPIKeyResponse`

NewCreateAPIKeyResponseWithDefaults instantiates a new CreateAPIKeyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKey

`func (o *CreateAPIKeyResponse) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *CreateAPIKeyResponse) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *CreateAPIKeyResponse) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.


### GetId

`func (o *CreateAPIKeyResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateAPIKeyResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateAPIKeyResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *CreateAPIKeyResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateAPIKeyResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateAPIKeyResponse) SetName(v string)`

SetName sets Name field to given value.


### GetPrefix

`func (o *CreateAPIKeyResponse) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *CreateAPIKeyResponse) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *CreateAPIKeyResponse) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.


### GetWarning

`func (o *CreateAPIKeyResponse) GetWarning() string`

GetWarning returns the Warning field if non-nil, zero value otherwise.

### GetWarningOk

`func (o *CreateAPIKeyResponse) GetWarningOk() (*string, bool)`

GetWarningOk returns a tuple with the Warning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarning

`func (o *CreateAPIKeyResponse) SetWarning(v string)`

SetWarning sets Warning field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


