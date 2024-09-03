# V3VerifySignalsPost200Response

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

### NewV3VerifySignalsPost200Response

`func NewV3VerifySignalsPost200Response() *V3VerifySignalsPost200Response`

NewV3VerifySignalsPost200Response instantiates a new V3VerifySignalsPost200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV3VerifySignalsPost200ResponseWithDefaults

`func NewV3VerifySignalsPost200ResponseWithDefaults() *V3VerifySignalsPost200Response`

NewV3VerifySignalsPost200ResponseWithDefaults instantiates a new V3VerifySignalsPost200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *V3VerifySignalsPost200Response) GetName() NameWithAlias`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V3VerifySignalsPost200Response) GetNameOk() (*NameWithAlias, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V3VerifySignalsPost200Response) SetName(v NameWithAlias)`

SetName sets Name field to given value.

### HasName

`func (o *V3VerifySignalsPost200Response) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDemographics

`func (o *V3VerifySignalsPost200Response) GetDemographics() Demographics`

GetDemographics returns the Demographics field if non-nil, zero value otherwise.

### GetDemographicsOk

`func (o *V3VerifySignalsPost200Response) GetDemographicsOk() (*Demographics, bool)`

GetDemographicsOk returns a tuple with the Demographics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDemographics

`func (o *V3VerifySignalsPost200Response) SetDemographics(v Demographics)`

SetDemographics sets Demographics field to given value.

### HasDemographics

`func (o *V3VerifySignalsPost200Response) HasDemographics() bool`

HasDemographics returns a boolean if a field has been set.

### GetLocations

`func (o *V3VerifySignalsPost200Response) GetLocations() []Location`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *V3VerifySignalsPost200Response) GetLocationsOk() (*[]Location, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *V3VerifySignalsPost200Response) SetLocations(v []Location)`

SetLocations sets Locations field to given value.

### HasLocations

`func (o *V3VerifySignalsPost200Response) HasLocations() bool`

HasLocations returns a boolean if a field has been set.

### GetEmails

`func (o *V3VerifySignalsPost200Response) GetEmails() []Email`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *V3VerifySignalsPost200Response) GetEmailsOk() (*[]Email, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *V3VerifySignalsPost200Response) SetEmails(v []Email)`

SetEmails sets Emails field to given value.

### HasEmails

`func (o *V3VerifySignalsPost200Response) HasEmails() bool`

HasEmails returns a boolean if a field has been set.

### GetPhones

`func (o *V3VerifySignalsPost200Response) GetPhones() []Phone`

GetPhones returns the Phones field if non-nil, zero value otherwise.

### GetPhonesOk

`func (o *V3VerifySignalsPost200Response) GetPhonesOk() (*[]Phone, bool)`

GetPhonesOk returns a tuple with the Phones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhones

`func (o *V3VerifySignalsPost200Response) SetPhones(v []Phone)`

SetPhones sets Phones field to given value.

### HasPhones

`func (o *V3VerifySignalsPost200Response) HasPhones() bool`

HasPhones returns a boolean if a field has been set.

### GetPersonIds

`func (o *V3VerifySignalsPost200Response) GetPersonIds() []string`

GetPersonIds returns the PersonIds field if non-nil, zero value otherwise.

### GetPersonIdsOk

`func (o *V3VerifySignalsPost200Response) GetPersonIdsOk() (*[]string, bool)`

GetPersonIdsOk returns a tuple with the PersonIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonIds

`func (o *V3VerifySignalsPost200Response) SetPersonIds(v []string)`

SetPersonIds sets PersonIds field to given value.

### HasPersonIds

`func (o *V3VerifySignalsPost200Response) HasPersonIds() bool`

HasPersonIds returns a boolean if a field has been set.

### GetIpAddresses

`func (o *V3VerifySignalsPost200Response) GetIpAddresses() []IpAddress`

GetIpAddresses returns the IpAddresses field if non-nil, zero value otherwise.

### GetIpAddressesOk

`func (o *V3VerifySignalsPost200Response) GetIpAddressesOk() (*[]IpAddress, bool)`

GetIpAddressesOk returns a tuple with the IpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddresses

`func (o *V3VerifySignalsPost200Response) SetIpAddresses(v []IpAddress)`

SetIpAddresses sets IpAddresses field to given value.

### HasIpAddresses

`func (o *V3VerifySignalsPost200Response) HasIpAddresses() bool`

HasIpAddresses returns a boolean if a field has been set.

### GetMessage

`func (o *V3VerifySignalsPost200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *V3VerifySignalsPost200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *V3VerifySignalsPost200Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *V3VerifySignalsPost200Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetSocialProfiles

`func (o *V3VerifySignalsPost200Response) GetSocialProfiles() []SocialProfileDetails`

GetSocialProfiles returns the SocialProfiles field if non-nil, zero value otherwise.

### GetSocialProfilesOk

`func (o *V3VerifySignalsPost200Response) GetSocialProfilesOk() (*[]SocialProfileDetails, bool)`

GetSocialProfilesOk returns a tuple with the SocialProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialProfiles

`func (o *V3VerifySignalsPost200Response) SetSocialProfiles(v []SocialProfileDetails)`

SetSocialProfiles sets SocialProfiles field to given value.

### HasSocialProfiles

`func (o *V3VerifySignalsPost200Response) HasSocialProfiles() bool`

HasSocialProfiles returns a boolean if a field has been set.

### GetEmployment

`func (o *V3VerifySignalsPost200Response) GetEmployment() []Employment`

GetEmployment returns the Employment field if non-nil, zero value otherwise.

### GetEmploymentOk

`func (o *V3VerifySignalsPost200Response) GetEmploymentOk() (*[]Employment, bool)`

GetEmploymentOk returns a tuple with the Employment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployment

`func (o *V3VerifySignalsPost200Response) SetEmployment(v []Employment)`

SetEmployment sets Employment field to given value.

### HasEmployment

`func (o *V3VerifySignalsPost200Response) HasEmployment() bool`

HasEmployment returns a boolean if a field has been set.

### GetEducation

`func (o *V3VerifySignalsPost200Response) GetEducation() []Education`

GetEducation returns the Education field if non-nil, zero value otherwise.

### GetEducationOk

`func (o *V3VerifySignalsPost200Response) GetEducationOk() (*[]Education, bool)`

GetEducationOk returns a tuple with the Education field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEducation

`func (o *V3VerifySignalsPost200Response) SetEducation(v []Education)`

SetEducation sets Education field to given value.

### HasEducation

`func (o *V3VerifySignalsPost200Response) HasEducation() bool`

HasEducation returns a boolean if a field has been set.

### GetPhotos

`func (o *V3VerifySignalsPost200Response) GetPhotos() []Photo`

GetPhotos returns the Photos field if non-nil, zero value otherwise.

### GetPhotosOk

`func (o *V3VerifySignalsPost200Response) GetPhotosOk() (*[]Photo, bool)`

GetPhotosOk returns a tuple with the Photos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhotos

`func (o *V3VerifySignalsPost200Response) SetPhotos(v []Photo)`

SetPhotos sets Photos field to given value.

### HasPhotos

`func (o *V3VerifySignalsPost200Response) HasPhotos() bool`

HasPhotos returns a boolean if a field has been set.

### GetEconomic

`func (o *V3VerifySignalsPost200Response) GetEconomic() Economic`

GetEconomic returns the Economic field if non-nil, zero value otherwise.

### GetEconomicOk

`func (o *V3VerifySignalsPost200Response) GetEconomicOk() (*Economic, bool)`

GetEconomicOk returns a tuple with the Economic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEconomic

`func (o *V3VerifySignalsPost200Response) SetEconomic(v Economic)`

SetEconomic sets Economic field to given value.

### HasEconomic

`func (o *V3VerifySignalsPost200Response) HasEconomic() bool`

HasEconomic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


