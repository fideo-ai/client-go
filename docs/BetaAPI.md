# \BetaAPI

All URIs are relative to *https://api.fideo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSignalPatternsSummary**](BetaAPI.md#GetSignalPatternsSummary) | **Post** /beta/signals/patterns/summary | Get signal patterns summary
[**GetSignalPatternsTimeseries**](BetaAPI.md#GetSignalPatternsTimeseries) | **Post** /beta/signals/patterns/timeseries | Get signal patterns timeseries



## GetSignalPatternsSummary

> SignalPatternRecencyResponse GetSignalPatternsSummary(ctx).SignalPatternsRequest(signalPatternsRequest).Execute()

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
	resp, r, err := apiClient.BetaAPI.GetSignalPatternsSummary(context.Background()).SignalPatternsRequest(signalPatternsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaAPI.GetSignalPatternsSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSignalPatternsSummary`: SignalPatternRecencyResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaAPI.GetSignalPatternsSummary`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSignalPatternsSummaryRequest struct via the builder pattern


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


## GetSignalPatternsTimeseries

> map[string]SignalPatternResponseUnit GetSignalPatternsTimeseries(ctx).SignalPatternsTimeseriesRequest(signalPatternsTimeseriesRequest).Execute()

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
	signalPatternsTimeseriesRequest := *openapiclient.NewSignalPatternsTimeseriesRequest("Interval_example", int32(123), "Email_example") // SignalPatternsTimeseriesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaAPI.GetSignalPatternsTimeseries(context.Background()).SignalPatternsTimeseriesRequest(signalPatternsTimeseriesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaAPI.GetSignalPatternsTimeseries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSignalPatternsTimeseries`: map[string]SignalPatternResponseUnit
	fmt.Fprintf(os.Stdout, "Response from `BetaAPI.GetSignalPatternsTimeseries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSignalPatternsTimeseriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **signalPatternsTimeseriesRequest** | [**SignalPatternsTimeseriesRequest**](SignalPatternsTimeseriesRequest.md) |  | 

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

