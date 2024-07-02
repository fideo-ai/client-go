# Phone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstSeenMs** | Pointer to **int64** |  | [optional] 
**LastSeenMs** | Pointer to **int64** |  | [optional] 
**Observations** | Pointer to **int32** |  | [optional] 
**Confidence** | Pointer to **float64** |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **string** |  | [optional] 

## Methods

### NewPhone

`func NewPhone() *Phone`

NewPhone instantiates a new Phone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhoneWithDefaults

`func NewPhoneWithDefaults() *Phone`

NewPhoneWithDefaults instantiates a new Phone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstSeenMs

`func (o *Phone) GetFirstSeenMs() int64`

GetFirstSeenMs returns the FirstSeenMs field if non-nil, zero value otherwise.

### GetFirstSeenMsOk

`func (o *Phone) GetFirstSeenMsOk() (*int64, bool)`

GetFirstSeenMsOk returns a tuple with the FirstSeenMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSeenMs

`func (o *Phone) SetFirstSeenMs(v int64)`

SetFirstSeenMs sets FirstSeenMs field to given value.

### HasFirstSeenMs

`func (o *Phone) HasFirstSeenMs() bool`

HasFirstSeenMs returns a boolean if a field has been set.

### GetLastSeenMs

`func (o *Phone) GetLastSeenMs() int64`

GetLastSeenMs returns the LastSeenMs field if non-nil, zero value otherwise.

### GetLastSeenMsOk

`func (o *Phone) GetLastSeenMsOk() (*int64, bool)`

GetLastSeenMsOk returns a tuple with the LastSeenMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeenMs

`func (o *Phone) SetLastSeenMs(v int64)`

SetLastSeenMs sets LastSeenMs field to given value.

### HasLastSeenMs

`func (o *Phone) HasLastSeenMs() bool`

HasLastSeenMs returns a boolean if a field has been set.

### GetObservations

`func (o *Phone) GetObservations() int32`

GetObservations returns the Observations field if non-nil, zero value otherwise.

### GetObservationsOk

`func (o *Phone) GetObservationsOk() (*int32, bool)`

GetObservationsOk returns a tuple with the Observations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservations

`func (o *Phone) SetObservations(v int32)`

SetObservations sets Observations field to given value.

### HasObservations

`func (o *Phone) HasObservations() bool`

HasObservations returns a boolean if a field has been set.

### GetConfidence

`func (o *Phone) GetConfidence() float64`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *Phone) GetConfidenceOk() (*float64, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *Phone) SetConfidence(v float64)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *Phone) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.

### GetLabel

`func (o *Phone) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *Phone) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *Phone) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *Phone) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetValue

`func (o *Phone) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Phone) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Phone) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *Phone) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


