# Email

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstSeenMs** | Pointer to **int64** |  | [optional] 
**LastSeenMs** | Pointer to **int64** |  | [optional] 
**Observations** | Pointer to **int32** |  | [optional] 
**Confidence** | Pointer to **float64** |  | [optional] 
**Value** | Pointer to **string** |  | [optional] 
**Md5** | Pointer to **string** |  | [optional] 
**Sha1** | Pointer to **string** |  | [optional] 
**Sha256** | Pointer to **string** |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**Activity** | Pointer to **float64** |  | [optional] 

## Methods

### NewEmail

`func NewEmail() *Email`

NewEmail instantiates a new Email object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailWithDefaults

`func NewEmailWithDefaults() *Email`

NewEmailWithDefaults instantiates a new Email object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstSeenMs

`func (o *Email) GetFirstSeenMs() int64`

GetFirstSeenMs returns the FirstSeenMs field if non-nil, zero value otherwise.

### GetFirstSeenMsOk

`func (o *Email) GetFirstSeenMsOk() (*int64, bool)`

GetFirstSeenMsOk returns a tuple with the FirstSeenMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSeenMs

`func (o *Email) SetFirstSeenMs(v int64)`

SetFirstSeenMs sets FirstSeenMs field to given value.

### HasFirstSeenMs

`func (o *Email) HasFirstSeenMs() bool`

HasFirstSeenMs returns a boolean if a field has been set.

### GetLastSeenMs

`func (o *Email) GetLastSeenMs() int64`

GetLastSeenMs returns the LastSeenMs field if non-nil, zero value otherwise.

### GetLastSeenMsOk

`func (o *Email) GetLastSeenMsOk() (*int64, bool)`

GetLastSeenMsOk returns a tuple with the LastSeenMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeenMs

`func (o *Email) SetLastSeenMs(v int64)`

SetLastSeenMs sets LastSeenMs field to given value.

### HasLastSeenMs

`func (o *Email) HasLastSeenMs() bool`

HasLastSeenMs returns a boolean if a field has been set.

### GetObservations

`func (o *Email) GetObservations() int32`

GetObservations returns the Observations field if non-nil, zero value otherwise.

### GetObservationsOk

`func (o *Email) GetObservationsOk() (*int32, bool)`

GetObservationsOk returns a tuple with the Observations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservations

`func (o *Email) SetObservations(v int32)`

SetObservations sets Observations field to given value.

### HasObservations

`func (o *Email) HasObservations() bool`

HasObservations returns a boolean if a field has been set.

### GetConfidence

`func (o *Email) GetConfidence() float64`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *Email) GetConfidenceOk() (*float64, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *Email) SetConfidence(v float64)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *Email) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.

### GetValue

`func (o *Email) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Email) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Email) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *Email) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetMd5

`func (o *Email) GetMd5() string`

GetMd5 returns the Md5 field if non-nil, zero value otherwise.

### GetMd5Ok

`func (o *Email) GetMd5Ok() (*string, bool)`

GetMd5Ok returns a tuple with the Md5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5

`func (o *Email) SetMd5(v string)`

SetMd5 sets Md5 field to given value.

### HasMd5

`func (o *Email) HasMd5() bool`

HasMd5 returns a boolean if a field has been set.

### GetSha1

`func (o *Email) GetSha1() string`

GetSha1 returns the Sha1 field if non-nil, zero value otherwise.

### GetSha1Ok

`func (o *Email) GetSha1Ok() (*string, bool)`

GetSha1Ok returns a tuple with the Sha1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha1

`func (o *Email) SetSha1(v string)`

SetSha1 sets Sha1 field to given value.

### HasSha1

`func (o *Email) HasSha1() bool`

HasSha1 returns a boolean if a field has been set.

### GetSha256

`func (o *Email) GetSha256() string`

GetSha256 returns the Sha256 field if non-nil, zero value otherwise.

### GetSha256Ok

`func (o *Email) GetSha256Ok() (*string, bool)`

GetSha256Ok returns a tuple with the Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256

`func (o *Email) SetSha256(v string)`

SetSha256 sets Sha256 field to given value.

### HasSha256

`func (o *Email) HasSha256() bool`

HasSha256 returns a boolean if a field has been set.

### GetLabel

`func (o *Email) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *Email) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *Email) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *Email) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetActivity

`func (o *Email) GetActivity() float64`

GetActivity returns the Activity field if non-nil, zero value otherwise.

### GetActivityOk

`func (o *Email) GetActivityOk() (*float64, bool)`

GetActivityOk returns a tuple with the Activity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivity

`func (o *Email) SetActivity(v float64)`

SetActivity sets Activity field to given value.

### HasActivity

`func (o *Email) HasActivity() bool`

HasActivity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


