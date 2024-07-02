# SignalsResponseV0

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to [**Name**](Name.md) |  | [optional] 
**Demographics** | Pointer to [**Demographics**](Demographics.md) |  | [optional] 
**Locations** | Pointer to [**[]Location**](Location.md) |  | [optional] 
**Emails** | Pointer to [**[]Email**](Email.md) |  | [optional] 
**Phones** | Pointer to [**[]Phone**](Phone.md) |  | [optional] 
**PersonIds** | Pointer to **[]string** |  | [optional] 
**IpAddresses** | Pointer to [**[]IpAddress**](IpAddress.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**SocialProfiles** | Pointer to [**SocialProfileUrls**](SocialProfileUrls.md) |  | [optional] 
**Employment** | Pointer to [**Employment**](Employment.md) |  | [optional] 

## Methods

### NewSignalsResponseV0

`func NewSignalsResponseV0() *SignalsResponseV0`

NewSignalsResponseV0 instantiates a new SignalsResponseV0 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignalsResponseV0WithDefaults

`func NewSignalsResponseV0WithDefaults() *SignalsResponseV0`

NewSignalsResponseV0WithDefaults instantiates a new SignalsResponseV0 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SignalsResponseV0) GetName() Name`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SignalsResponseV0) GetNameOk() (*Name, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SignalsResponseV0) SetName(v Name)`

SetName sets Name field to given value.

### HasName

`func (o *SignalsResponseV0) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDemographics

`func (o *SignalsResponseV0) GetDemographics() Demographics`

GetDemographics returns the Demographics field if non-nil, zero value otherwise.

### GetDemographicsOk

`func (o *SignalsResponseV0) GetDemographicsOk() (*Demographics, bool)`

GetDemographicsOk returns a tuple with the Demographics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDemographics

`func (o *SignalsResponseV0) SetDemographics(v Demographics)`

SetDemographics sets Demographics field to given value.

### HasDemographics

`func (o *SignalsResponseV0) HasDemographics() bool`

HasDemographics returns a boolean if a field has been set.

### GetLocations

`func (o *SignalsResponseV0) GetLocations() []Location`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *SignalsResponseV0) GetLocationsOk() (*[]Location, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *SignalsResponseV0) SetLocations(v []Location)`

SetLocations sets Locations field to given value.

### HasLocations

`func (o *SignalsResponseV0) HasLocations() bool`

HasLocations returns a boolean if a field has been set.

### GetEmails

`func (o *SignalsResponseV0) GetEmails() []Email`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *SignalsResponseV0) GetEmailsOk() (*[]Email, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *SignalsResponseV0) SetEmails(v []Email)`

SetEmails sets Emails field to given value.

### HasEmails

`func (o *SignalsResponseV0) HasEmails() bool`

HasEmails returns a boolean if a field has been set.

### GetPhones

`func (o *SignalsResponseV0) GetPhones() []Phone`

GetPhones returns the Phones field if non-nil, zero value otherwise.

### GetPhonesOk

`func (o *SignalsResponseV0) GetPhonesOk() (*[]Phone, bool)`

GetPhonesOk returns a tuple with the Phones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhones

`func (o *SignalsResponseV0) SetPhones(v []Phone)`

SetPhones sets Phones field to given value.

### HasPhones

`func (o *SignalsResponseV0) HasPhones() bool`

HasPhones returns a boolean if a field has been set.

### GetPersonIds

`func (o *SignalsResponseV0) GetPersonIds() []string`

GetPersonIds returns the PersonIds field if non-nil, zero value otherwise.

### GetPersonIdsOk

`func (o *SignalsResponseV0) GetPersonIdsOk() (*[]string, bool)`

GetPersonIdsOk returns a tuple with the PersonIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonIds

`func (o *SignalsResponseV0) SetPersonIds(v []string)`

SetPersonIds sets PersonIds field to given value.

### HasPersonIds

`func (o *SignalsResponseV0) HasPersonIds() bool`

HasPersonIds returns a boolean if a field has been set.

### GetIpAddresses

`func (o *SignalsResponseV0) GetIpAddresses() []IpAddress`

GetIpAddresses returns the IpAddresses field if non-nil, zero value otherwise.

### GetIpAddressesOk

`func (o *SignalsResponseV0) GetIpAddressesOk() (*[]IpAddress, bool)`

GetIpAddressesOk returns a tuple with the IpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddresses

`func (o *SignalsResponseV0) SetIpAddresses(v []IpAddress)`

SetIpAddresses sets IpAddresses field to given value.

### HasIpAddresses

`func (o *SignalsResponseV0) HasIpAddresses() bool`

HasIpAddresses returns a boolean if a field has been set.

### GetMessage

`func (o *SignalsResponseV0) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SignalsResponseV0) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SignalsResponseV0) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *SignalsResponseV0) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetSocialProfiles

`func (o *SignalsResponseV0) GetSocialProfiles() SocialProfileUrls`

GetSocialProfiles returns the SocialProfiles field if non-nil, zero value otherwise.

### GetSocialProfilesOk

`func (o *SignalsResponseV0) GetSocialProfilesOk() (*SocialProfileUrls, bool)`

GetSocialProfilesOk returns a tuple with the SocialProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialProfiles

`func (o *SignalsResponseV0) SetSocialProfiles(v SocialProfileUrls)`

SetSocialProfiles sets SocialProfiles field to given value.

### HasSocialProfiles

`func (o *SignalsResponseV0) HasSocialProfiles() bool`

HasSocialProfiles returns a boolean if a field has been set.

### GetEmployment

`func (o *SignalsResponseV0) GetEmployment() Employment`

GetEmployment returns the Employment field if non-nil, zero value otherwise.

### GetEmploymentOk

`func (o *SignalsResponseV0) GetEmploymentOk() (*Employment, bool)`

GetEmploymentOk returns a tuple with the Employment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployment

`func (o *SignalsResponseV0) SetEmployment(v Employment)`

SetEmployment sets Employment field to given value.

### HasEmployment

`func (o *SignalsResponseV0) HasEmployment() bool`

HasEmployment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


