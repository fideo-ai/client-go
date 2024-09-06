# \VerifyAPI

All URIs are relative to *https://api.fideo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VerifyPost**](VerifyAPI.md#VerifyPost) | **Post** /verify | 



## VerifyPost

> VerifyResponse VerifyPost(ctx).MultiFieldReq(multiFieldReq).Execute()



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
	multiFieldReq := *openapiclient.NewMultiFieldReq() // MultiFieldReq |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VerifyAPI.VerifyPost(context.Background()).MultiFieldReq(multiFieldReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VerifyAPI.VerifyPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VerifyPost`: VerifyResponse
	fmt.Fprintf(os.Stdout, "Response from `VerifyAPI.VerifyPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVerifyPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **multiFieldReq** | [**MultiFieldReq**](MultiFieldReq.md) |  | 

### Return type

[**VerifyResponse**](VerifyResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

