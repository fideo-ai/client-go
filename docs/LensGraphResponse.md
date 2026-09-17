# LensGraphResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Graph** | Pointer to **[]interface{}** |  | [optional] 
**SeedGraphIds** | Pointer to **map[string][]string** |  | [optional] 

## Methods

### NewLensGraphResponse

`func NewLensGraphResponse() *LensGraphResponse`

NewLensGraphResponse instantiates a new LensGraphResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLensGraphResponseWithDefaults

`func NewLensGraphResponseWithDefaults() *LensGraphResponse`

NewLensGraphResponseWithDefaults instantiates a new LensGraphResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGraph

`func (o *LensGraphResponse) GetGraph() []interface{}`

GetGraph returns the Graph field if non-nil, zero value otherwise.

### GetGraphOk

`func (o *LensGraphResponse) GetGraphOk() (*[]interface{}, bool)`

GetGraphOk returns a tuple with the Graph field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraph

`func (o *LensGraphResponse) SetGraph(v []interface{})`

SetGraph sets Graph field to given value.

### HasGraph

`func (o *LensGraphResponse) HasGraph() bool`

HasGraph returns a boolean if a field has been set.

### GetSeedGraphIds

`func (o *LensGraphResponse) GetSeedGraphIds() map[string][]string`

GetSeedGraphIds returns the SeedGraphIds field if non-nil, zero value otherwise.

### GetSeedGraphIdsOk

`func (o *LensGraphResponse) GetSeedGraphIdsOk() (*map[string][]string, bool)`

GetSeedGraphIdsOk returns a tuple with the SeedGraphIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeedGraphIds

`func (o *LensGraphResponse) SetSeedGraphIds(v map[string][]string)`

SetSeedGraphIds sets SeedGraphIds field to given value.

### HasSeedGraphIds

`func (o *LensGraphResponse) HasSeedGraphIds() bool`

HasSeedGraphIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


