# Employment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Current** | Pointer to **bool** |  | [optional] 
**Company** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Domain** | Pointer to **string** |  | [optional] 
**Start** | Pointer to [**EmploymentDate**](EmploymentDate.md) |  | [optional] 
**End** | Pointer to [**EmploymentDate**](EmploymentDate.md) |  | [optional] 

## Methods

### NewEmployment

`func NewEmployment() *Employment`

NewEmployment instantiates a new Employment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmploymentWithDefaults

`func NewEmploymentWithDefaults() *Employment`

NewEmploymentWithDefaults instantiates a new Employment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrent

`func (o *Employment) GetCurrent() bool`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *Employment) GetCurrentOk() (*bool, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *Employment) SetCurrent(v bool)`

SetCurrent sets Current field to given value.

### HasCurrent

`func (o *Employment) HasCurrent() bool`

HasCurrent returns a boolean if a field has been set.

### GetCompany

`func (o *Employment) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *Employment) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *Employment) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *Employment) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetTitle

`func (o *Employment) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Employment) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Employment) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Employment) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDomain

`func (o *Employment) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *Employment) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *Employment) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *Employment) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetStart

`func (o *Employment) GetStart() EmploymentDate`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *Employment) GetStartOk() (*EmploymentDate, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *Employment) SetStart(v EmploymentDate)`

SetStart sets Start field to given value.

### HasStart

`func (o *Employment) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetEnd

`func (o *Employment) GetEnd() EmploymentDate`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *Employment) GetEndOk() (*EmploymentDate, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *Employment) SetEnd(v EmploymentDate)`

SetEnd sets End field to given value.

### HasEnd

`func (o *Employment) HasEnd() bool`

HasEnd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


