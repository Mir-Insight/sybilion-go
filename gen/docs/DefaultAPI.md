# \DefaultAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1ApiKeysGet**](DefaultAPI.md#ApiV1ApiKeysGet) | **Get** /api/v1/api-keys | List API keys (metadata only)
[**ApiV1ApiKeysIdDelete**](DefaultAPI.md#ApiV1ApiKeysIdDelete) | **Delete** /api/v1/api-keys/{id} | Revoke an API key
[**ApiV1ApiKeysPost**](DefaultAPI.md#ApiV1ApiKeysPost) | **Post** /api/v1/api-keys | Create API key
[**ApiV1BillingCheckoutPost**](DefaultAPI.md#ApiV1BillingCheckoutPost) | **Post** /api/v1/billing/checkout | Stripe Checkout URL (variable top-up, min €5)
[**ApiV1BillingPortalPost**](DefaultAPI.md#ApiV1BillingPortalPost) | **Post** /api/v1/billing/portal | Stripe Customer Portal URL (billing history and receipts as configured in Stripe Dashboard)
[**ApiV1DriversPost**](DefaultAPI.md#ApiV1DriversPost) | **Post** /api/v1/drivers | Recommend ranked driver-dataset candidates (synchronous proxy)
[**ApiV1ForecastsIdArtifactsNameGet**](DefaultAPI.md#ApiV1ForecastsIdArtifactsNameGet) | **Get** /api/v1/forecasts/{id}/artifacts/{name} | Stream a single output file (proxied from internal artifact store; no gs:// URLs)
[**ApiV1ForecastsIdGet**](DefaultAPI.md#ApiV1ForecastsIdGet) | **Get** /api/v1/forecasts/{id} | Poll forecast job status and artifact metadata (downloads via artifacts sub-path)
[**ApiV1ForecastsPost**](DefaultAPI.md#ApiV1ForecastsPost) | **Post** /api/v1/forecasts | Start async forecast job (Temporal + artifact store)
[**ApiV1JobsGet**](DefaultAPI.md#ApiV1JobsGet) | **Get** /api/v1/jobs | Paginated list of the caller&#39;s async jobs
[**ApiV1MeAutoRechargePatch**](DefaultAPI.md#ApiV1MeAutoRechargePatch) | **Patch** /api/v1/me/auto-recharge | Update auto-recharge preferences (threshold, target, monthly cap)
[**ApiV1MeGet**](DefaultAPI.md#ApiV1MeGet) | **Get** /api/v1/me | Current user, balances, and usage tier
[**ApiV1TiersGet**](DefaultAPI.md#ApiV1TiersGet) | **Get** /api/v1/tiers | List the configured tier ladder, current tier, and progress
[**ApiV1UsageGet**](DefaultAPI.md#ApiV1UsageGet) | **Get** /api/v1/usage | Paginated usage (charged tasks)
[**HealthGet**](DefaultAPI.md#HealthGet) | **Get** /health | Health check



## ApiV1ApiKeysGet

> ApiV1ApiKeysGet200Response ApiV1ApiKeysGet(ctx).Execute()

List API keys (metadata only)

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Mir-Insight/developers-portal-api-sdk-go/gen"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ApiV1ApiKeysGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ApiV1ApiKeysGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1ApiKeysGet`: ApiV1ApiKeysGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ApiV1ApiKeysGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1ApiKeysGetRequest struct via the builder pattern


### Return type

[**ApiV1ApiKeysGet200Response**](ApiV1ApiKeysGet200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1ApiKeysIdDelete

> ApiV1ApiKeysIdDelete(ctx, id).Execute()

Revoke an API key

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Mir-Insight/developers-portal-api-sdk-go/gen"
)

func main() {
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.ApiV1ApiKeysIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ApiV1ApiKeysIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1ApiKeysIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1ApiKeysPost

> CreateAPIKeyResponse ApiV1ApiKeysPost(ctx).ApiV1ApiKeysPostRequest(apiV1ApiKeysPostRequest).Execute()

Create API key

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Mir-Insight/developers-portal-api-sdk-go/gen"
)

func main() {
	apiV1ApiKeysPostRequest := *openapiclient.NewApiV1ApiKeysPostRequest("Name_example") // ApiV1ApiKeysPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ApiV1ApiKeysPost(context.Background()).ApiV1ApiKeysPostRequest(apiV1ApiKeysPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ApiV1ApiKeysPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1ApiKeysPost`: CreateAPIKeyResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ApiV1ApiKeysPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1ApiKeysPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiV1ApiKeysPostRequest** | [**ApiV1ApiKeysPostRequest**](ApiV1ApiKeysPostRequest.md) |  | 

### Return type

[**CreateAPIKeyResponse**](CreateAPIKeyResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1BillingCheckoutPost

> ApiV1BillingCheckoutPost(ctx).CheckoutRequest(checkoutRequest).Execute()

Stripe Checkout URL (variable top-up, min €5)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Mir-Insight/developers-portal-api-sdk-go/gen"
)

func main() {
	checkoutRequest := *openapiclient.NewCheckoutRequest() // CheckoutRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.ApiV1BillingCheckoutPost(context.Background()).CheckoutRequest(checkoutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ApiV1BillingCheckoutPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1BillingCheckoutPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **checkoutRequest** | [**CheckoutRequest**](CheckoutRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1BillingPortalPost

> ApiV1BillingPortalPost200Response ApiV1BillingPortalPost(ctx).Execute()

Stripe Customer Portal URL (billing history and receipts as configured in Stripe Dashboard)

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Mir-Insight/developers-portal-api-sdk-go/gen"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ApiV1BillingPortalPost(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ApiV1BillingPortalPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1BillingPortalPost`: ApiV1BillingPortalPost200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ApiV1BillingPortalPost`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1BillingPortalPostRequest struct via the builder pattern


### Return type

[**ApiV1BillingPortalPost200Response**](ApiV1BillingPortalPost200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1DriversPost

> ApiV1DriversPost(ctx).RecommendRequestV1(recommendRequestV1).Execute()

Recommend ranked driver-dataset candidates (synchronous proxy)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Mir-Insight/developers-portal-api-sdk-go/gen"
)

func main() {
	recommendRequestV1 := *openapiclient.NewRecommendRequestV1("Version_example", float64(123), *openapiclient.NewTimeseriesMetadata("Title_example")) // RecommendRequestV1 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.ApiV1DriversPost(context.Background()).RecommendRequestV1(recommendRequestV1).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ApiV1DriversPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1DriversPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **recommendRequestV1** | [**RecommendRequestV1**](RecommendRequestV1.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1ForecastsIdArtifactsNameGet

> *os.File ApiV1ForecastsIdArtifactsNameGet(ctx, id, name).Execute()

Stream a single output file (proxied from internal artifact store; no gs:// URLs)

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Mir-Insight/developers-portal-api-sdk-go/gen"
)

func main() {
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ApiV1ForecastsIdArtifactsNameGet(context.Background(), id, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ApiV1ForecastsIdArtifactsNameGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1ForecastsIdArtifactsNameGet`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ApiV1ForecastsIdArtifactsNameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1ForecastsIdArtifactsNameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[***os.File**](*os.File.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1ForecastsIdGet

> ForecastJobResponse ApiV1ForecastsIdGet(ctx, id).Execute()

Poll forecast job status and artifact metadata (downloads via artifacts sub-path)

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Mir-Insight/developers-portal-api-sdk-go/gen"
)

func main() {
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ApiV1ForecastsIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ApiV1ForecastsIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1ForecastsIdGet`: ForecastJobResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ApiV1ForecastsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1ForecastsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ForecastJobResponse**](ForecastJobResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1ForecastsPost

> ForecastAcceptedResponse ApiV1ForecastsPost(ctx).ForecastRequestV1(forecastRequestV1).Execute()

Start async forecast job (Temporal + artifact store)

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Mir-Insight/developers-portal-api-sdk-go/gen"
)

func main() {
	forecastRequestV1 := *openapiclient.NewForecastRequestV1("PipelineVersion_example", int32(123), "Frequency_example", float64(123), *openapiclient.NewTimeseriesMetadata("Title_example"), map[string]float32{"key": float32(123)}) // ForecastRequestV1 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ApiV1ForecastsPost(context.Background()).ForecastRequestV1(forecastRequestV1).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ApiV1ForecastsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1ForecastsPost`: ForecastAcceptedResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ApiV1ForecastsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1ForecastsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **forecastRequestV1** | [**ForecastRequestV1**](ForecastRequestV1.md) |  | 

### Return type

[**ForecastAcceptedResponse**](ForecastAcceptedResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1JobsGet

> ApiV1JobsGet200Response ApiV1JobsGet(ctx).Page(page).Limit(limit).Sort(sort).Order(order).Status(status).PipelineType(pipelineType).Execute()

Paginated list of the caller's async jobs



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Mir-Insight/developers-portal-api-sdk-go/gen"
)

func main() {
	page := int32(56) // int32 |  (optional) (default to 1)
	limit := int32(56) // int32 | Page size. Capped at 200. (optional) (default to 50)
	sort := "sort_example" // string |  (optional) (default to "created_at")
	order := "order_example" // string |  (optional) (default to "desc")
	status := "status_example" // string | Optional status filter. (optional)
	pipelineType := "pipelineType_example" // string | Optional pipeline type filter (today only \"forecast\" is emitted). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ApiV1JobsGet(context.Background()).Page(page).Limit(limit).Sort(sort).Order(order).Status(status).PipelineType(pipelineType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ApiV1JobsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1JobsGet`: ApiV1JobsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ApiV1JobsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1JobsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 1]
 **limit** | **int32** | Page size. Capped at 200. | [default to 50]
 **sort** | **string** |  | [default to &quot;created_at&quot;]
 **order** | **string** |  | [default to &quot;desc&quot;]
 **status** | **string** | Optional status filter. | 
 **pipelineType** | **string** | Optional pipeline type filter (today only \&quot;forecast\&quot; is emitted). | 

### Return type

[**ApiV1JobsGet200Response**](ApiV1JobsGet200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1MeAutoRechargePatch

> ApiV1MeAutoRechargePatch(ctx).AutoRechargePatch(autoRechargePatch).Execute()

Update auto-recharge preferences (threshold, target, monthly cap)

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Mir-Insight/developers-portal-api-sdk-go/gen"
)

func main() {
	autoRechargePatch := *openapiclient.NewAutoRechargePatch(false, int64(123), int64(123), int64(123)) // AutoRechargePatch | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.ApiV1MeAutoRechargePatch(context.Background()).AutoRechargePatch(autoRechargePatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ApiV1MeAutoRechargePatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1MeAutoRechargePatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **autoRechargePatch** | [**AutoRechargePatch**](AutoRechargePatch.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1MeGet

> MeResponse ApiV1MeGet(ctx).Execute()

Current user, balances, and usage tier

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Mir-Insight/developers-portal-api-sdk-go/gen"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ApiV1MeGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ApiV1MeGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1MeGet`: MeResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ApiV1MeGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1MeGetRequest struct via the builder pattern


### Return type

[**MeResponse**](MeResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1TiersGet

> TiersResponse ApiV1TiersGet(ctx).Execute()

List the configured tier ladder, current tier, and progress



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Mir-Insight/developers-portal-api-sdk-go/gen"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ApiV1TiersGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ApiV1TiersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1TiersGet`: TiersResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ApiV1TiersGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1TiersGetRequest struct via the builder pattern


### Return type

[**TiersResponse**](TiersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1UsageGet

> ApiV1UsageGet200Response ApiV1UsageGet(ctx).Page(page).Limit(limit).Sort(sort).Order(order).Execute()

Paginated usage (charged tasks)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Mir-Insight/developers-portal-api-sdk-go/gen"
)

func main() {
	page := int32(56) // int32 | 1-indexed page number. (optional) (default to 1)
	limit := int32(56) // int32 | Page size. Capped at 200. (optional) (default to 50)
	sort := "sort_example" // string | Column to sort by. (optional) (default to "id")
	order := "order_example" // string | Sort direction. (optional) (default to "desc")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ApiV1UsageGet(context.Background()).Page(page).Limit(limit).Sort(sort).Order(order).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ApiV1UsageGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1UsageGet`: ApiV1UsageGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ApiV1UsageGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1UsageGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | 1-indexed page number. | [default to 1]
 **limit** | **int32** | Page size. Capped at 200. | [default to 50]
 **sort** | **string** | Column to sort by. | [default to &quot;id&quot;]
 **order** | **string** | Sort direction. | [default to &quot;desc&quot;]

### Return type

[**ApiV1UsageGet200Response**](ApiV1UsageGet200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HealthGet

> map[string]interface{} HealthGet(ctx).Execute()

Health check

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Mir-Insight/developers-portal-api-sdk-go/gen"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.HealthGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.HealthGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HealthGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.HealthGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiHealthGetRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

