# PrefillResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SessionId** | **string** | UUIDv7 session identifier to send with reviewed identity data. | 
**Status** | **string** |  | 
**Individual** | Pointer to [**MultiFieldReqWithOptions**](MultiFieldReqWithOptions.md) | Resolved identity in the same multifield shape accepted by the follow-up Prefill request. The caller may review or edit these values and resubmit them with the returned sessionId. | [optional] 
**Evaluation** | Pointer to [**PrefillEvaluation**](PrefillEvaluation.md) |  | [optional] 

## Methods

### NewPrefillResponse

`func NewPrefillResponse(sessionId string, status string, ) *PrefillResponse`

NewPrefillResponse instantiates a new PrefillResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrefillResponseWithDefaults

`func NewPrefillResponseWithDefaults() *PrefillResponse`

NewPrefillResponseWithDefaults instantiates a new PrefillResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSessionId

`func (o *PrefillResponse) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *PrefillResponse) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *PrefillResponse) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.


### GetStatus

`func (o *PrefillResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PrefillResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PrefillResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetIndividual

`func (o *PrefillResponse) GetIndividual() MultiFieldReqWithOptions`

GetIndividual returns the Individual field if non-nil, zero value otherwise.

### GetIndividualOk

`func (o *PrefillResponse) GetIndividualOk() (*MultiFieldReqWithOptions, bool)`

GetIndividualOk returns a tuple with the Individual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividual

`func (o *PrefillResponse) SetIndividual(v MultiFieldReqWithOptions)`

SetIndividual sets Individual field to given value.

### HasIndividual

`func (o *PrefillResponse) HasIndividual() bool`

HasIndividual returns a boolean if a field has been set.

### GetEvaluation

`func (o *PrefillResponse) GetEvaluation() PrefillEvaluation`

GetEvaluation returns the Evaluation field if non-nil, zero value otherwise.

### GetEvaluationOk

`func (o *PrefillResponse) GetEvaluationOk() (*PrefillEvaluation, bool)`

GetEvaluationOk returns a tuple with the Evaluation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluation

`func (o *PrefillResponse) SetEvaluation(v PrefillEvaluation)`

SetEvaluation sets Evaluation field to given value.

### HasEvaluation

`func (o *PrefillResponse) HasEvaluation() bool`

HasEvaluation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


