# LensGraphRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | **string** | Lens graph response mode. | 
**Query** | **string** | Comma-separated Lens graph seed query. | 
**Hops** | Pointer to **int32** | Requested graph hop count. Defaults to 2 when omitted. | [optional] [default to 2]

## Methods

### NewLensGraphRequest

`func NewLensGraphRequest(mode string, query string, ) *LensGraphRequest`

NewLensGraphRequest instantiates a new LensGraphRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLensGraphRequestWithDefaults

`func NewLensGraphRequestWithDefaults() *LensGraphRequest`

NewLensGraphRequestWithDefaults instantiates a new LensGraphRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *LensGraphRequest) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *LensGraphRequest) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *LensGraphRequest) SetMode(v string)`

SetMode sets Mode field to given value.


### GetQuery

`func (o *LensGraphRequest) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *LensGraphRequest) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *LensGraphRequest) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetHops

`func (o *LensGraphRequest) GetHops() int32`

GetHops returns the Hops field if non-nil, zero value otherwise.

### GetHopsOk

`func (o *LensGraphRequest) GetHopsOk() (*int32, bool)`

GetHopsOk returns a tuple with the Hops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHops

`func (o *LensGraphRequest) SetHops(v int32)`

SetHops sets Hops field to given value.

### HasHops

`func (o *LensGraphRequest) HasHops() bool`

HasHops returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


