# IpAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstSeenMs** | Pointer to **int64** |  | [optional] 
**LastSeenMs** | Pointer to **int64** |  | [optional] 
**Observations** | Pointer to **int32** |  | [optional] 
**Confidence** | Pointer to **float64** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 

## Methods

### NewIpAddress

`func NewIpAddress() *IpAddress`

NewIpAddress instantiates a new IpAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpAddressWithDefaults

`func NewIpAddressWithDefaults() *IpAddress`

NewIpAddressWithDefaults instantiates a new IpAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstSeenMs

`func (o *IpAddress) GetFirstSeenMs() int64`

GetFirstSeenMs returns the FirstSeenMs field if non-nil, zero value otherwise.

### GetFirstSeenMsOk

`func (o *IpAddress) GetFirstSeenMsOk() (*int64, bool)`

GetFirstSeenMsOk returns a tuple with the FirstSeenMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSeenMs

`func (o *IpAddress) SetFirstSeenMs(v int64)`

SetFirstSeenMs sets FirstSeenMs field to given value.

### HasFirstSeenMs

`func (o *IpAddress) HasFirstSeenMs() bool`

HasFirstSeenMs returns a boolean if a field has been set.

### GetLastSeenMs

`func (o *IpAddress) GetLastSeenMs() int64`

GetLastSeenMs returns the LastSeenMs field if non-nil, zero value otherwise.

### GetLastSeenMsOk

`func (o *IpAddress) GetLastSeenMsOk() (*int64, bool)`

GetLastSeenMsOk returns a tuple with the LastSeenMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeenMs

`func (o *IpAddress) SetLastSeenMs(v int64)`

SetLastSeenMs sets LastSeenMs field to given value.

### HasLastSeenMs

`func (o *IpAddress) HasLastSeenMs() bool`

HasLastSeenMs returns a boolean if a field has been set.

### GetObservations

`func (o *IpAddress) GetObservations() int32`

GetObservations returns the Observations field if non-nil, zero value otherwise.

### GetObservationsOk

`func (o *IpAddress) GetObservationsOk() (*int32, bool)`

GetObservationsOk returns a tuple with the Observations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservations

`func (o *IpAddress) SetObservations(v int32)`

SetObservations sets Observations field to given value.

### HasObservations

`func (o *IpAddress) HasObservations() bool`

HasObservations returns a boolean if a field has been set.

### GetConfidence

`func (o *IpAddress) GetConfidence() float64`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *IpAddress) GetConfidenceOk() (*float64, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *IpAddress) SetConfidence(v float64)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *IpAddress) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.

### GetId

`func (o *IpAddress) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IpAddress) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IpAddress) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IpAddress) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


