# \VerifyAPI

All URIs are relative to *https://api.fullcontact.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VerifyMatchPost**](VerifyAPI.md#VerifyMatchPost) | **Post** /verify.match | 



## VerifyMatchPost

> MatchResponse VerifyMatchPost(ctx).MultiFieldReq(multiFieldReq).Execute()



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
	resp, r, err := apiClient.VerifyAPI.VerifyMatchPost(context.Background()).MultiFieldReq(multiFieldReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VerifyAPI.VerifyMatchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VerifyMatchPost`: MatchResponse
	fmt.Fprintf(os.Stdout, "Response from `VerifyAPI.VerifyMatchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVerifyMatchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **multiFieldReq** | [**MultiFieldReq**](MultiFieldReq.md) |  | 

### Return type

[**MatchResponse**](MatchResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

