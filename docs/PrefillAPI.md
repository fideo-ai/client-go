# \PrefillAPI

All URIs are relative to *https://api.fideo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Prefill**](PrefillAPI.md#Prefill) | **Post** /prefill | Resolve or evaluate onboarding identity fields



## Prefill

> PrefillResponse Prefill(ctx).MultiFieldReqWithOptions(multiFieldReqWithOptions).Execute()

Resolve or evaluate onboarding identity fields



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
	multiFieldReqWithOptions := *openapiclient.NewMultiFieldReqWithOptions() // MultiFieldReqWithOptions | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrefillAPI.Prefill(context.Background()).MultiFieldReqWithOptions(multiFieldReqWithOptions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrefillAPI.Prefill``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Prefill`: PrefillResponse
	fmt.Fprintf(os.Stdout, "Response from `PrefillAPI.Prefill`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPrefillRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **multiFieldReqWithOptions** | [**MultiFieldReqWithOptions**](MultiFieldReqWithOptions.md) |  | 

### Return type

[**PrefillResponse**](PrefillResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

