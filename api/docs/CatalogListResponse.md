# CatalogListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | **[]map[string]interface{}** | All regions or categories, sorted by numeric &#x60;id&#x60; ascending. Each element is an object with at least &#x60;id&#x60; (integer) plus the fields available for that dimension.  | 

## Methods

### NewCatalogListResponse

`func NewCatalogListResponse(items []map[string]interface{}, ) *CatalogListResponse`

NewCatalogListResponse instantiates a new CatalogListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogListResponseWithDefaults

`func NewCatalogListResponseWithDefaults() *CatalogListResponse`

NewCatalogListResponseWithDefaults instantiates a new CatalogListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *CatalogListResponse) GetItems() []map[string]interface{}`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *CatalogListResponse) GetItemsOk() (*[]map[string]interface{}, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *CatalogListResponse) SetItems(v []map[string]interface{})`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


