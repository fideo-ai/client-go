# \SignalsAPI

All URIs are relative to *https://api.fullcontact.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VerifySignalsPost**](SignalsAPI.md#VerifySignalsPost) | **Post** /verify.signals | 



## VerifySignalsPost

> VerifySignalsPost200Response VerifySignalsPost(ctx).V(v).MultiFieldReq(multiFieldReq).Execute()



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
	multiFieldReq := *openapiclient.NewMultiFieldReq() // MultiFieldReq |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SignalsAPI.VerifySignalsPost(context.Background()).V(v).MultiFieldReq(multiFieldReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SignalsAPI.VerifySignalsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VerifySignalsPost`: VerifySignalsPost200Response
	fmt.Fprintf(os.Stdout, "Response from `SignalsAPI.VerifySignalsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVerifySignalsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v** | **string** |  | 
 **multiFieldReq** | [**MultiFieldReq**](MultiFieldReq.md) |  | 

### Return type

[**VerifySignalsPost200Response**](VerifySignalsPost200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

