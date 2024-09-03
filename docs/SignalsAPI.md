# \SignalsAPI

All URIs are relative to *https://api.fideo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SignalsPost**](SignalsAPI.md#SignalsPost) | **Post** /signals | 



## SignalsPost

> SignalsPost200Response SignalsPost(ctx).V(v).MultiFieldReqWithOptions(multiFieldReqWithOptions).Execute()



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
	v := "v_example" // string |  (optional)
	multiFieldReqWithOptions := *openapiclient.NewMultiFieldReqWithOptions() // MultiFieldReqWithOptions |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SignalsAPI.SignalsPost(context.Background()).V(v).MultiFieldReqWithOptions(multiFieldReqWithOptions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SignalsAPI.SignalsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SignalsPost`: SignalsPost200Response
	fmt.Fprintf(os.Stdout, "Response from `SignalsAPI.SignalsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSignalsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v** | **string** |  | 
 **multiFieldReqWithOptions** | [**MultiFieldReqWithOptions**](MultiFieldReqWithOptions.md) |  | 

### Return type

[**SignalsPost200Response**](SignalsPost200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

