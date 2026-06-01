# SignalPatternsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | Email address to get signal patterns for | 
**Interval** | Pointer to **string** | Time interval for timeseries data (used only for timeseries endpoint) | [optional] 
**Count** | Pointer to **int32** | Number of data points to return (used only for timeseries endpoint) | [optional] 

## Methods

### NewSignalPatternsRequest

`func NewSignalPatternsRequest(email string, ) *SignalPatternsRequest`

NewSignalPatternsRequest instantiates a new SignalPatternsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignalPatternsRequestWithDefaults

`func NewSignalPatternsRequestWithDefaults() *SignalPatternsRequest`

NewSignalPatternsRequestWithDefaults instantiates a new SignalPatternsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *SignalPatternsRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *SignalPatternsRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *SignalPatternsRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetInterval

`func (o *SignalPatternsRequest) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *SignalPatternsRequest) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *SignalPatternsRequest) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *SignalPatternsRequest) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetCount

`func (o *SignalPatternsRequest) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *SignalPatternsRequest) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *SignalPatternsRequest) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *SignalPatternsRequest) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


