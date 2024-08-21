/*
Fideo API

Fideo Intelligence offers an identity intelligence product that protects the public good. - [Fideo Privacy Policy](https://www.fideo.ai/privacy-policy/)

API version: 1.0.3
Contact: support@fideo.ai
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fideo

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// SignalsAPIService SignalsAPI service
type SignalsAPIService service

type ApiVerifySignalsPostRequest struct {
	ctx context.Context
	ApiService *SignalsAPIService
	v *string
	multiFieldReqWithOptions *MultiFieldReqWithOptions
}

func (r ApiVerifySignalsPostRequest) V(v string) ApiVerifySignalsPostRequest {
	r.v = &v
	return r
}

func (r ApiVerifySignalsPostRequest) MultiFieldReqWithOptions(multiFieldReqWithOptions MultiFieldReqWithOptions) ApiVerifySignalsPostRequest {
	r.multiFieldReqWithOptions = &multiFieldReqWithOptions
	return r
}

func (r ApiVerifySignalsPostRequest) Execute() (*VerifySignalsPost200Response, *http.Response, error) {
	return r.ApiService.VerifySignalsPostExecute(r)
}

/*
VerifySignalsPost Method for VerifySignalsPost

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiVerifySignalsPostRequest
*/
func (a *SignalsAPIService) VerifySignalsPost(ctx context.Context) ApiVerifySignalsPostRequest {
	return ApiVerifySignalsPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return VerifySignalsPost200Response
func (a *SignalsAPIService) VerifySignalsPostExecute(r ApiVerifySignalsPostRequest) (*VerifySignalsPost200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *VerifySignalsPost200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SignalsAPIService.VerifySignalsPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/verify.signals"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.v != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "v", r.v, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.multiFieldReqWithOptions
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
