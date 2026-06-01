# \BetaAPI

All URIs are relative to *https://api.fideo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BetaSignalsPatternsSummaryPost**](BetaAPI.md#BetaSignalsPatternsSummaryPost) | **Post** /beta/signals/patterns/summary | Get signal patterns summary
[**BetaSignalsPatternsTimeseriesPost**](BetaAPI.md#BetaSignalsPatternsTimeseriesPost) | **Post** /beta/signals/patterns/timeseries | Get signal patterns timeseries



## BetaSignalsPatternsSummaryPost

> SignalPatternRecencyResponse BetaSignalsPatternsSummaryPost(ctx).SignalPatternsRequest(signalPatternsRequest).Execute()

Get signal patterns summary



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	signalPatternsRequest := *openapiclient.NewSignalPatternsRequest("Email_example") // SignalPatternsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaAPI.BetaSignalsPatternsSummaryPost(context.Background()).SignalPatternsRequest(signalPatternsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaAPI.BetaSignalsPatternsSummaryPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BetaSignalsPatternsSummaryPost`: SignalPatternRecencyResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaAPI.BetaSignalsPatternsSummaryPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBetaSignalsPatternsSummaryPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **signalPatternsRequest** | [**SignalPatternsRequest**](SignalPatternsRequest.md) |  | 

### Return type

[**SignalPatternRecencyResponse**](SignalPatternRecencyResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BetaSignalsPatternsTimeseriesPost

> map[string]SignalPatternResponseUnit BetaSignalsPatternsTimeseriesPost(ctx).SignalPatternsRequest(signalPatternsRequest).Execute()

Get signal patterns timeseries



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	signalPatternsRequest := *openapiclient.NewSignalPatternsRequest("Email_example") // SignalPatternsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaAPI.BetaSignalsPatternsTimeseriesPost(context.Background()).SignalPatternsRequest(signalPatternsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaAPI.BetaSignalsPatternsTimeseriesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BetaSignalsPatternsTimeseriesPost`: map[string]SignalPatternResponseUnit
	fmt.Fprintf(os.Stdout, "Response from `BetaAPI.BetaSignalsPatternsTimeseriesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBetaSignalsPatternsTimeseriesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **signalPatternsRequest** | [**SignalPatternsRequest**](SignalPatternsRequest.md) |  | 

### Return type

[**map[string]SignalPatternResponseUnit**](SignalPatternResponseUnit.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

