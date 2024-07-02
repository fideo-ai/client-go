# SignalsResponseV20240424

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to [**NameWithAlias**](NameWithAlias.md) |  | [optional] 
**Demographics** | Pointer to [**Demographics**](Demographics.md) |  | [optional] 
**Locations** | Pointer to [**[]Location**](Location.md) |  | [optional] 
**Emails** | Pointer to [**[]Email**](Email.md) |  | [optional] 
**Phones** | Pointer to [**[]Phone**](Phone.md) |  | [optional] 
**PersonIds** | Pointer to **[]string** |  | [optional] 
**IpAddresses** | Pointer to [**[]IpAddress**](IpAddress.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**SocialProfiles** | Pointer to [**[]SocialProfileDetails**](SocialProfileDetails.md) |  | [optional] 
**Employment** | Pointer to [**[]Employment**](Employment.md) |  | [optional] 
**Education** | Pointer to [**[]Education**](Education.md) |  | [optional] 
**Photos** | Pointer to [**[]Photo**](Photo.md) |  | [optional] 
**Economic** | Pointer to [**Economic**](Economic.md) |  | [optional] 

## Methods

### NewSignalsResponseV20240424

`func NewSignalsResponseV20240424() *SignalsResponseV20240424`

NewSignalsResponseV20240424 instantiates a new SignalsResponseV20240424 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignalsResponseV20240424WithDefaults

`func NewSignalsResponseV20240424WithDefaults() *SignalsResponseV20240424`

NewSignalsResponseV20240424WithDefaults instantiates a new SignalsResponseV20240424 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SignalsResponseV20240424) GetName() NameWithAlias`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SignalsResponseV20240424) GetNameOk() (*NameWithAlias, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SignalsResponseV20240424) SetName(v NameWithAlias)`

SetName sets Name field to given value.

### HasName

`func (o *SignalsResponseV20240424) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDemographics

`func (o *SignalsResponseV20240424) GetDemographics() Demographics`

GetDemographics returns the Demographics field if non-nil, zero value otherwise.

### GetDemographicsOk

`func (o *SignalsResponseV20240424) GetDemographicsOk() (*Demographics, bool)`

GetDemographicsOk returns a tuple with the Demographics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDemographics

`func (o *SignalsResponseV20240424) SetDemographics(v Demographics)`

SetDemographics sets Demographics field to given value.

### HasDemographics

`func (o *SignalsResponseV20240424) HasDemographics() bool`

HasDemographics returns a boolean if a field has been set.

### GetLocations

`func (o *SignalsResponseV20240424) GetLocations() []Location`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *SignalsResponseV20240424) GetLocationsOk() (*[]Location, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *SignalsResponseV20240424) SetLocations(v []Location)`

SetLocations sets Locations field to given value.

### HasLocations

`func (o *SignalsResponseV20240424) HasLocations() bool`

HasLocations returns a boolean if a field has been set.

### GetEmails

`func (o *SignalsResponseV20240424) GetEmails() []Email`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *SignalsResponseV20240424) GetEmailsOk() (*[]Email, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *SignalsResponseV20240424) SetEmails(v []Email)`

SetEmails sets Emails field to given value.

### HasEmails

`func (o *SignalsResponseV20240424) HasEmails() bool`

HasEmails returns a boolean if a field has been set.

### GetPhones

`func (o *SignalsResponseV20240424) GetPhones() []Phone`

GetPhones returns the Phones field if non-nil, zero value otherwise.

### GetPhonesOk

`func (o *SignalsResponseV20240424) GetPhonesOk() (*[]Phone, bool)`

GetPhonesOk returns a tuple with the Phones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhones

`func (o *SignalsResponseV20240424) SetPhones(v []Phone)`

SetPhones sets Phones field to given value.

### HasPhones

`func (o *SignalsResponseV20240424) HasPhones() bool`

HasPhones returns a boolean if a field has been set.

### GetPersonIds

`func (o *SignalsResponseV20240424) GetPersonIds() []string`

GetPersonIds returns the PersonIds field if non-nil, zero value otherwise.

### GetPersonIdsOk

`func (o *SignalsResponseV20240424) GetPersonIdsOk() (*[]string, bool)`

GetPersonIdsOk returns a tuple with the PersonIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonIds

`func (o *SignalsResponseV20240424) SetPersonIds(v []string)`

SetPersonIds sets PersonIds field to given value.

### HasPersonIds

`func (o *SignalsResponseV20240424) HasPersonIds() bool`

HasPersonIds returns a boolean if a field has been set.

### GetIpAddresses

`func (o *SignalsResponseV20240424) GetIpAddresses() []IpAddress`

GetIpAddresses returns the IpAddresses field if non-nil, zero value otherwise.

### GetIpAddressesOk

`func (o *SignalsResponseV20240424) GetIpAddressesOk() (*[]IpAddress, bool)`

GetIpAddressesOk returns a tuple with the IpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddresses

`func (o *SignalsResponseV20240424) SetIpAddresses(v []IpAddress)`

SetIpAddresses sets IpAddresses field to given value.

### HasIpAddresses

`func (o *SignalsResponseV20240424) HasIpAddresses() bool`

HasIpAddresses returns a boolean if a field has been set.

### GetMessage

`func (o *SignalsResponseV20240424) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SignalsResponseV20240424) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SignalsResponseV20240424) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *SignalsResponseV20240424) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetSocialProfiles

`func (o *SignalsResponseV20240424) GetSocialProfiles() []SocialProfileDetails`

GetSocialProfiles returns the SocialProfiles field if non-nil, zero value otherwise.

### GetSocialProfilesOk

`func (o *SignalsResponseV20240424) GetSocialProfilesOk() (*[]SocialProfileDetails, bool)`

GetSocialProfilesOk returns a tuple with the SocialProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialProfiles

`func (o *SignalsResponseV20240424) SetSocialProfiles(v []SocialProfileDetails)`

SetSocialProfiles sets SocialProfiles field to given value.

### HasSocialProfiles

`func (o *SignalsResponseV20240424) HasSocialProfiles() bool`

HasSocialProfiles returns a boolean if a field has been set.

### GetEmployment

`func (o *SignalsResponseV20240424) GetEmployment() []Employment`

GetEmployment returns the Employment field if non-nil, zero value otherwise.

### GetEmploymentOk

`func (o *SignalsResponseV20240424) GetEmploymentOk() (*[]Employment, bool)`

GetEmploymentOk returns a tuple with the Employment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployment

`func (o *SignalsResponseV20240424) SetEmployment(v []Employment)`

SetEmployment sets Employment field to given value.

### HasEmployment

`func (o *SignalsResponseV20240424) HasEmployment() bool`

HasEmployment returns a boolean if a field has been set.

### GetEducation

`func (o *SignalsResponseV20240424) GetEducation() []Education`

GetEducation returns the Education field if non-nil, zero value otherwise.

### GetEducationOk

`func (o *SignalsResponseV20240424) GetEducationOk() (*[]Education, bool)`

GetEducationOk returns a tuple with the Education field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEducation

`func (o *SignalsResponseV20240424) SetEducation(v []Education)`

SetEducation sets Education field to given value.

### HasEducation

`func (o *SignalsResponseV20240424) HasEducation() bool`

HasEducation returns a boolean if a field has been set.

### GetPhotos

`func (o *SignalsResponseV20240424) GetPhotos() []Photo`

GetPhotos returns the Photos field if non-nil, zero value otherwise.

### GetPhotosOk

`func (o *SignalsResponseV20240424) GetPhotosOk() (*[]Photo, bool)`

GetPhotosOk returns a tuple with the Photos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhotos

`func (o *SignalsResponseV20240424) SetPhotos(v []Photo)`

SetPhotos sets Photos field to given value.

### HasPhotos

`func (o *SignalsResponseV20240424) HasPhotos() bool`

HasPhotos returns a boolean if a field has been set.

### GetEconomic

`func (o *SignalsResponseV20240424) GetEconomic() Economic`

GetEconomic returns the Economic field if non-nil, zero value otherwise.

### GetEconomicOk

`func (o *SignalsResponseV20240424) GetEconomicOk() (*Economic, bool)`

GetEconomicOk returns a tuple with the Economic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEconomic

`func (o *SignalsResponseV20240424) SetEconomic(v Economic)`

SetEconomic sets Economic field to given value.

### HasEconomic

`func (o *SignalsResponseV20240424) HasEconomic() bool`

HasEconomic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


