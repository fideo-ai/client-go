# Demographics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Age** | Pointer to **int32** |  | [optional] 
**AgeRange** | Pointer to **string** |  | [optional] 
**Gender** | Pointer to **string** |  | [optional] 
**LocationGeneral** | Pointer to **string** |  | [optional] 

## Methods

### NewDemographics

`func NewDemographics() *Demographics`

NewDemographics instantiates a new Demographics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDemographicsWithDefaults

`func NewDemographicsWithDefaults() *Demographics`

NewDemographicsWithDefaults instantiates a new Demographics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAge

`func (o *Demographics) GetAge() int32`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *Demographics) GetAgeOk() (*int32, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *Demographics) SetAge(v int32)`

SetAge sets Age field to given value.

### HasAge

`func (o *Demographics) HasAge() bool`

HasAge returns a boolean if a field has been set.

### GetAgeRange

`func (o *Demographics) GetAgeRange() string`

GetAgeRange returns the AgeRange field if non-nil, zero value otherwise.

### GetAgeRangeOk

`func (o *Demographics) GetAgeRangeOk() (*string, bool)`

GetAgeRangeOk returns a tuple with the AgeRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgeRange

`func (o *Demographics) SetAgeRange(v string)`

SetAgeRange sets AgeRange field to given value.

### HasAgeRange

`func (o *Demographics) HasAgeRange() bool`

HasAgeRange returns a boolean if a field has been set.

### GetGender

`func (o *Demographics) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *Demographics) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *Demographics) SetGender(v string)`

SetGender sets Gender field to given value.

### HasGender

`func (o *Demographics) HasGender() bool`

HasGender returns a boolean if a field has been set.

### GetLocationGeneral

`func (o *Demographics) GetLocationGeneral() string`

GetLocationGeneral returns the LocationGeneral field if non-nil, zero value otherwise.

### GetLocationGeneralOk

`func (o *Demographics) GetLocationGeneralOk() (*string, bool)`

GetLocationGeneralOk returns a tuple with the LocationGeneral field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationGeneral

`func (o *Demographics) SetLocationGeneral(v string)`

SetLocationGeneral sets LocationGeneral field to given value.

### HasLocationGeneral

`func (o *Demographics) HasLocationGeneral() bool`

HasLocationGeneral returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


