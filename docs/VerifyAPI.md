# \VerifyAPI

All URIs are relative to *https://api.fideo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V3VerifyMatchPost**](VerifyAPI.md#V3VerifyMatchPost) | **Post** /v3/verify.match | 
[**VerifyPost**](VerifyAPI.md#VerifyPost) | **Post** /verify | 



## V3VerifyMatchPost

> MatchResponse V3VerifyMatchPost(ctx).MultiFieldReq(multiFieldReq).Execute()



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
	resp, r, err := apiClient.VerifyAPI.V3VerifyMatchPost(context.Background()).MultiFieldReq(multiFieldReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VerifyAPI.V3VerifyMatchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3VerifyMatchPost`: MatchResponse
	fmt.Fprintf(os.Stdout, "Response from `VerifyAPI.V3VerifyMatchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV3VerifyMatchPostRequest struct via the builder pattern


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


## VerifyPost

> MatchResponse VerifyPost(ctx).MultiFieldReq(multiFieldReq).Execute()



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
	// response from `VerifyPost`: MatchResponse
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

[**MatchResponse**](MatchResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

