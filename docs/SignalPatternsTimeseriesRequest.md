# SignalPatternsTimeseriesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Interval** | **string** | Time interval for timeseries data | 
**Count** | **int32** | Number of data points to return | 

## Methods

### NewSignalPatternsTimeseriesRequest

`func NewSignalPatternsTimeseriesRequest(interval string, count int32, ) *SignalPatternsTimeseriesRequest`

NewSignalPatternsTimeseriesRequest instantiates a new SignalPatternsTimeseriesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignalPatternsTimeseriesRequestWithDefaults

`func NewSignalPatternsTimeseriesRequestWithDefaults() *SignalPatternsTimeseriesRequest`

NewSignalPatternsTimeseriesRequestWithDefaults instantiates a new SignalPatternsTimeseriesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterval

`func (o *SignalPatternsTimeseriesRequest) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *SignalPatternsTimeseriesRequest) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *SignalPatternsTimeseriesRequest) SetInterval(v string)`

SetInterval sets Interval field to given value.


### GetCount

`func (o *SignalPatternsTimeseriesRequest) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *SignalPatternsTimeseriesRequest) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *SignalPatternsTimeseriesRequest) SetCount(v int32)`

SetCount sets Count field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


