# PrefillEvaluation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Risk** | Pointer to **float64** |  | [optional] 
**Checks** | Pointer to [**[]CheckResult**](CheckResult.md) |  | [optional] 

## Methods

### NewPrefillEvaluation

`func NewPrefillEvaluation() *PrefillEvaluation`

NewPrefillEvaluation instantiates a new PrefillEvaluation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrefillEvaluationWithDefaults

`func NewPrefillEvaluationWithDefaults() *PrefillEvaluation`

NewPrefillEvaluationWithDefaults instantiates a new PrefillEvaluation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRisk

`func (o *PrefillEvaluation) GetRisk() float64`

GetRisk returns the Risk field if non-nil, zero value otherwise.

### GetRiskOk

`func (o *PrefillEvaluation) GetRiskOk() (*float64, bool)`

GetRiskOk returns a tuple with the Risk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRisk

`func (o *PrefillEvaluation) SetRisk(v float64)`

SetRisk sets Risk field to given value.

### HasRisk

`func (o *PrefillEvaluation) HasRisk() bool`

HasRisk returns a boolean if a field has been set.

### GetChecks

`func (o *PrefillEvaluation) GetChecks() []CheckResult`

GetChecks returns the Checks field if non-nil, zero value otherwise.

### GetChecksOk

`func (o *PrefillEvaluation) GetChecksOk() (*[]CheckResult, bool)`

GetChecksOk returns a tuple with the Checks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecks

`func (o *PrefillEvaluation) SetChecks(v []CheckResult)`

SetChecks sets Checks field to given value.

### HasChecks

`func (o *PrefillEvaluation) HasChecks() bool`

HasChecks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


