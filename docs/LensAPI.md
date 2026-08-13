# \LensAPI

All URIs are relative to *https://api.fideo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LensGraph**](LensAPI.md#LensGraph) | **Post** /v3/lens.graph | Query the Lens graph



## LensGraph

> LensGraphResponse LensGraph(ctx).LensGraphRequest(lensGraphRequest).Execute()

Query the Lens graph



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
	lensGraphRequest := *openapiclient.NewLensGraphRequest("Mode_example", "Query_example") // LensGraphRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LensAPI.LensGraph(context.Background()).LensGraphRequest(lensGraphRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LensAPI.LensGraph``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LensGraph`: LensGraphResponse
	fmt.Fprintf(os.Stdout, "Response from `LensAPI.LensGraph`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLensGraphRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lensGraphRequest** | [**LensGraphRequest**](LensGraphRequest.md) |  | 

### Return type

[**LensGraphResponse**](LensGraphResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

