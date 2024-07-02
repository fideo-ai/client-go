# Education

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Degree** | Pointer to **string** |  | [optional] 
**End** | Pointer to [**EducationDate**](EducationDate.md) |  | [optional] 

## Methods

### NewEducation

`func NewEducation() *Education`

NewEducation instantiates a new Education object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEducationWithDefaults

`func NewEducationWithDefaults() *Education`

NewEducationWithDefaults instantiates a new Education object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Education) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Education) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Education) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Education) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDegree

`func (o *Education) GetDegree() string`

GetDegree returns the Degree field if non-nil, zero value otherwise.

### GetDegreeOk

`func (o *Education) GetDegreeOk() (*string, bool)`

GetDegreeOk returns a tuple with the Degree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDegree

`func (o *Education) SetDegree(v string)`

SetDegree sets Degree field to given value.

### HasDegree

`func (o *Education) HasDegree() bool`

HasDegree returns a boolean if a field has been set.

### GetEnd

`func (o *Education) GetEnd() EducationDate`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *Education) GetEndOk() (*EducationDate, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *Education) SetEnd(v EducationDate)`

SetEnd sets End field to given value.

### HasEnd

`func (o *Education) HasEnd() bool`

HasEnd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


