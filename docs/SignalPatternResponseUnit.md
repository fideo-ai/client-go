# SignalPatternResponseUnit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Observations** | Pointer to **int64** | Number of observations | [optional] 
**Sources** | Pointer to **int64** | Number of sources | [optional] 
**Events** | Pointer to [**[]SignalPatternResponseUnitCount**](SignalPatternResponseUnitCount.md) |  | [optional] 

## Methods

### NewSignalPatternResponseUnit

`func NewSignalPatternResponseUnit() *SignalPatternResponseUnit`

NewSignalPatternResponseUnit instantiates a new SignalPatternResponseUnit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignalPatternResponseUnitWithDefaults

`func NewSignalPatternResponseUnitWithDefaults() *SignalPatternResponseUnit`

NewSignalPatternResponseUnitWithDefaults instantiates a new SignalPatternResponseUnit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObservations

`func (o *SignalPatternResponseUnit) GetObservations() int64`

GetObservations returns the Observations field if non-nil, zero value otherwise.

### GetObservationsOk

`func (o *SignalPatternResponseUnit) GetObservationsOk() (*int64, bool)`

GetObservationsOk returns a tuple with the Observations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservations

`func (o *SignalPatternResponseUnit) SetObservations(v int64)`

SetObservations sets Observations field to given value.

### HasObservations

`func (o *SignalPatternResponseUnit) HasObservations() bool`

HasObservations returns a boolean if a field has been set.

### GetSources

`func (o *SignalPatternResponseUnit) GetSources() int64`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *SignalPatternResponseUnit) GetSourcesOk() (*int64, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *SignalPatternResponseUnit) SetSources(v int64)`

SetSources sets Sources field to given value.

### HasSources

`func (o *SignalPatternResponseUnit) HasSources() bool`

HasSources returns a boolean if a field has been set.

### GetEvents

`func (o *SignalPatternResponseUnit) GetEvents() []SignalPatternResponseUnitCount`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *SignalPatternResponseUnit) GetEventsOk() (*[]SignalPatternResponseUnitCount, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *SignalPatternResponseUnit) SetEvents(v []SignalPatternResponseUnitCount)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *SignalPatternResponseUnit) HasEvents() bool`

HasEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


