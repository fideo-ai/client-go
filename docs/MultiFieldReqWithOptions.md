# MultiFieldReqWithOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Twitter** | Pointer to **string** |  | [optional] 
**Linkedin** | Pointer to **string** |  | [optional] 
**RecordId** | Pointer to **string** |  | [optional] 
**PersonId** | Pointer to **string** |  | [optional] 
**PartnerId** | Pointer to **string** |  | [optional] 
**Location** | Pointer to [**LocationReq**](LocationReq.md) |  | [optional] 
**Avatar** | Pointer to **string** |  | [optional] 
**Website** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Organization** | Pointer to **string** |  | [optional] 
**Emails** | Pointer to **[]string** |  | [optional] 
**Phones** | Pointer to **[]string** |  | [optional] 
**Profiles** | Pointer to [**[]SocialProfileReq**](SocialProfileReq.md) |  | [optional] 
**Maids** | Pointer to **[]string** |  | [optional] 
**Name** | Pointer to [**PersonNameReq**](PersonNameReq.md) |  | [optional] 
**PartnerKeys** | Pointer to **map[string]string** |  | [optional] 
**LiNonid** | Pointer to **string** |  | [optional] 
**PanoramaId** | Pointer to **string** |  | [optional] 
**Placekey** | Pointer to **string** |  | [optional] 
**GeneratePid** | Pointer to **bool** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Phone** | Pointer to **string** |  | [optional] 
**Profile** | Pointer to [**SocialProfileReq**](SocialProfileReq.md) |  | [optional] 
**Maid** | Pointer to **string** |  | [optional] 
**Infer** | Pointer to **bool** |  | [optional] 
**Confidence** | Pointer to **string** |  | [optional] 
**Birthday** | Pointer to **string** |  | [optional] 
**IpAddress** | Pointer to **string** |  | [optional] 
**Countries** | Pointer to **[]string** |  | [optional] 
**ExcludeCountries** | Pointer to **[]string** |  | [optional] 

## Methods

### NewMultiFieldReqWithOptions

`func NewMultiFieldReqWithOptions() *MultiFieldReqWithOptions`

NewMultiFieldReqWithOptions instantiates a new MultiFieldReqWithOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultiFieldReqWithOptionsWithDefaults

`func NewMultiFieldReqWithOptionsWithDefaults() *MultiFieldReqWithOptions`

NewMultiFieldReqWithOptionsWithDefaults instantiates a new MultiFieldReqWithOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTwitter

`func (o *MultiFieldReqWithOptions) GetTwitter() string`

GetTwitter returns the Twitter field if non-nil, zero value otherwise.

### GetTwitterOk

`func (o *MultiFieldReqWithOptions) GetTwitterOk() (*string, bool)`

GetTwitterOk returns a tuple with the Twitter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitter

`func (o *MultiFieldReqWithOptions) SetTwitter(v string)`

SetTwitter sets Twitter field to given value.

### HasTwitter

`func (o *MultiFieldReqWithOptions) HasTwitter() bool`

HasTwitter returns a boolean if a field has been set.

### GetLinkedin

`func (o *MultiFieldReqWithOptions) GetLinkedin() string`

GetLinkedin returns the Linkedin field if non-nil, zero value otherwise.

### GetLinkedinOk

`func (o *MultiFieldReqWithOptions) GetLinkedinOk() (*string, bool)`

GetLinkedinOk returns a tuple with the Linkedin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedin

`func (o *MultiFieldReqWithOptions) SetLinkedin(v string)`

SetLinkedin sets Linkedin field to given value.

### HasLinkedin

`func (o *MultiFieldReqWithOptions) HasLinkedin() bool`

HasLinkedin returns a boolean if a field has been set.

### GetRecordId

`func (o *MultiFieldReqWithOptions) GetRecordId() string`

GetRecordId returns the RecordId field if non-nil, zero value otherwise.

### GetRecordIdOk

`func (o *MultiFieldReqWithOptions) GetRecordIdOk() (*string, bool)`

GetRecordIdOk returns a tuple with the RecordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordId

`func (o *MultiFieldReqWithOptions) SetRecordId(v string)`

SetRecordId sets RecordId field to given value.

### HasRecordId

`func (o *MultiFieldReqWithOptions) HasRecordId() bool`

HasRecordId returns a boolean if a field has been set.

### GetPersonId

`func (o *MultiFieldReqWithOptions) GetPersonId() string`

GetPersonId returns the PersonId field if non-nil, zero value otherwise.

### GetPersonIdOk

`func (o *MultiFieldReqWithOptions) GetPersonIdOk() (*string, bool)`

GetPersonIdOk returns a tuple with the PersonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonId

`func (o *MultiFieldReqWithOptions) SetPersonId(v string)`

SetPersonId sets PersonId field to given value.

### HasPersonId

`func (o *MultiFieldReqWithOptions) HasPersonId() bool`

HasPersonId returns a boolean if a field has been set.

### GetPartnerId

`func (o *MultiFieldReqWithOptions) GetPartnerId() string`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *MultiFieldReqWithOptions) GetPartnerIdOk() (*string, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *MultiFieldReqWithOptions) SetPartnerId(v string)`

SetPartnerId sets PartnerId field to given value.

### HasPartnerId

`func (o *MultiFieldReqWithOptions) HasPartnerId() bool`

HasPartnerId returns a boolean if a field has been set.

### GetLocation

`func (o *MultiFieldReqWithOptions) GetLocation() LocationReq`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *MultiFieldReqWithOptions) GetLocationOk() (*LocationReq, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *MultiFieldReqWithOptions) SetLocation(v LocationReq)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *MultiFieldReqWithOptions) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetAvatar

`func (o *MultiFieldReqWithOptions) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *MultiFieldReqWithOptions) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *MultiFieldReqWithOptions) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *MultiFieldReqWithOptions) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### GetWebsite

`func (o *MultiFieldReqWithOptions) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *MultiFieldReqWithOptions) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *MultiFieldReqWithOptions) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *MultiFieldReqWithOptions) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.

### GetTitle

`func (o *MultiFieldReqWithOptions) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *MultiFieldReqWithOptions) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *MultiFieldReqWithOptions) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *MultiFieldReqWithOptions) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetOrganization

`func (o *MultiFieldReqWithOptions) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *MultiFieldReqWithOptions) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *MultiFieldReqWithOptions) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *MultiFieldReqWithOptions) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetEmails

`func (o *MultiFieldReqWithOptions) GetEmails() []string`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *MultiFieldReqWithOptions) GetEmailsOk() (*[]string, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *MultiFieldReqWithOptions) SetEmails(v []string)`

SetEmails sets Emails field to given value.

### HasEmails

`func (o *MultiFieldReqWithOptions) HasEmails() bool`

HasEmails returns a boolean if a field has been set.

### GetPhones

`func (o *MultiFieldReqWithOptions) GetPhones() []string`

GetPhones returns the Phones field if non-nil, zero value otherwise.

### GetPhonesOk

`func (o *MultiFieldReqWithOptions) GetPhonesOk() (*[]string, bool)`

GetPhonesOk returns a tuple with the Phones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhones

`func (o *MultiFieldReqWithOptions) SetPhones(v []string)`

SetPhones sets Phones field to given value.

### HasPhones

`func (o *MultiFieldReqWithOptions) HasPhones() bool`

HasPhones returns a boolean if a field has been set.

### GetProfiles

`func (o *MultiFieldReqWithOptions) GetProfiles() []SocialProfileReq`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *MultiFieldReqWithOptions) GetProfilesOk() (*[]SocialProfileReq, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *MultiFieldReqWithOptions) SetProfiles(v []SocialProfileReq)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *MultiFieldReqWithOptions) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.

### GetMaids

`func (o *MultiFieldReqWithOptions) GetMaids() []string`

GetMaids returns the Maids field if non-nil, zero value otherwise.

### GetMaidsOk

`func (o *MultiFieldReqWithOptions) GetMaidsOk() (*[]string, bool)`

GetMaidsOk returns a tuple with the Maids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaids

`func (o *MultiFieldReqWithOptions) SetMaids(v []string)`

SetMaids sets Maids field to given value.

### HasMaids

`func (o *MultiFieldReqWithOptions) HasMaids() bool`

HasMaids returns a boolean if a field has been set.

### GetName

`func (o *MultiFieldReqWithOptions) GetName() PersonNameReq`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MultiFieldReqWithOptions) GetNameOk() (*PersonNameReq, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MultiFieldReqWithOptions) SetName(v PersonNameReq)`

SetName sets Name field to given value.

### HasName

`func (o *MultiFieldReqWithOptions) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPartnerKeys

`func (o *MultiFieldReqWithOptions) GetPartnerKeys() map[string]string`

GetPartnerKeys returns the PartnerKeys field if non-nil, zero value otherwise.

### GetPartnerKeysOk

`func (o *MultiFieldReqWithOptions) GetPartnerKeysOk() (*map[string]string, bool)`

GetPartnerKeysOk returns a tuple with the PartnerKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerKeys

`func (o *MultiFieldReqWithOptions) SetPartnerKeys(v map[string]string)`

SetPartnerKeys sets PartnerKeys field to given value.

### HasPartnerKeys

`func (o *MultiFieldReqWithOptions) HasPartnerKeys() bool`

HasPartnerKeys returns a boolean if a field has been set.

### GetLiNonid

`func (o *MultiFieldReqWithOptions) GetLiNonid() string`

GetLiNonid returns the LiNonid field if non-nil, zero value otherwise.

### GetLiNonidOk

`func (o *MultiFieldReqWithOptions) GetLiNonidOk() (*string, bool)`

GetLiNonidOk returns a tuple with the LiNonid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiNonid

`func (o *MultiFieldReqWithOptions) SetLiNonid(v string)`

SetLiNonid sets LiNonid field to given value.

### HasLiNonid

`func (o *MultiFieldReqWithOptions) HasLiNonid() bool`

HasLiNonid returns a boolean if a field has been set.

### GetPanoramaId

`func (o *MultiFieldReqWithOptions) GetPanoramaId() string`

GetPanoramaId returns the PanoramaId field if non-nil, zero value otherwise.

### GetPanoramaIdOk

`func (o *MultiFieldReqWithOptions) GetPanoramaIdOk() (*string, bool)`

GetPanoramaIdOk returns a tuple with the PanoramaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPanoramaId

`func (o *MultiFieldReqWithOptions) SetPanoramaId(v string)`

SetPanoramaId sets PanoramaId field to given value.

### HasPanoramaId

`func (o *MultiFieldReqWithOptions) HasPanoramaId() bool`

HasPanoramaId returns a boolean if a field has been set.

### GetPlacekey

`func (o *MultiFieldReqWithOptions) GetPlacekey() string`

GetPlacekey returns the Placekey field if non-nil, zero value otherwise.

### GetPlacekeyOk

`func (o *MultiFieldReqWithOptions) GetPlacekeyOk() (*string, bool)`

GetPlacekeyOk returns a tuple with the Placekey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacekey

`func (o *MultiFieldReqWithOptions) SetPlacekey(v string)`

SetPlacekey sets Placekey field to given value.

### HasPlacekey

`func (o *MultiFieldReqWithOptions) HasPlacekey() bool`

HasPlacekey returns a boolean if a field has been set.

### GetGeneratePid

`func (o *MultiFieldReqWithOptions) GetGeneratePid() bool`

GetGeneratePid returns the GeneratePid field if non-nil, zero value otherwise.

### GetGeneratePidOk

`func (o *MultiFieldReqWithOptions) GetGeneratePidOk() (*bool, bool)`

GetGeneratePidOk returns a tuple with the GeneratePid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratePid

`func (o *MultiFieldReqWithOptions) SetGeneratePid(v bool)`

SetGeneratePid sets GeneratePid field to given value.

### HasGeneratePid

`func (o *MultiFieldReqWithOptions) HasGeneratePid() bool`

HasGeneratePid returns a boolean if a field has been set.

### GetEmail

`func (o *MultiFieldReqWithOptions) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *MultiFieldReqWithOptions) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *MultiFieldReqWithOptions) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *MultiFieldReqWithOptions) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPhone

`func (o *MultiFieldReqWithOptions) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *MultiFieldReqWithOptions) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *MultiFieldReqWithOptions) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *MultiFieldReqWithOptions) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetProfile

`func (o *MultiFieldReqWithOptions) GetProfile() SocialProfileReq`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *MultiFieldReqWithOptions) GetProfileOk() (*SocialProfileReq, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *MultiFieldReqWithOptions) SetProfile(v SocialProfileReq)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *MultiFieldReqWithOptions) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetMaid

`func (o *MultiFieldReqWithOptions) GetMaid() string`

GetMaid returns the Maid field if non-nil, zero value otherwise.

### GetMaidOk

`func (o *MultiFieldReqWithOptions) GetMaidOk() (*string, bool)`

GetMaidOk returns a tuple with the Maid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaid

`func (o *MultiFieldReqWithOptions) SetMaid(v string)`

SetMaid sets Maid field to given value.

### HasMaid

`func (o *MultiFieldReqWithOptions) HasMaid() bool`

HasMaid returns a boolean if a field has been set.

### GetInfer

`func (o *MultiFieldReqWithOptions) GetInfer() bool`

GetInfer returns the Infer field if non-nil, zero value otherwise.

### GetInferOk

`func (o *MultiFieldReqWithOptions) GetInferOk() (*bool, bool)`

GetInferOk returns a tuple with the Infer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfer

`func (o *MultiFieldReqWithOptions) SetInfer(v bool)`

SetInfer sets Infer field to given value.

### HasInfer

`func (o *MultiFieldReqWithOptions) HasInfer() bool`

HasInfer returns a boolean if a field has been set.

### GetConfidence

`func (o *MultiFieldReqWithOptions) GetConfidence() string`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *MultiFieldReqWithOptions) GetConfidenceOk() (*string, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *MultiFieldReqWithOptions) SetConfidence(v string)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *MultiFieldReqWithOptions) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.

### GetBirthday

`func (o *MultiFieldReqWithOptions) GetBirthday() string`

GetBirthday returns the Birthday field if non-nil, zero value otherwise.

### GetBirthdayOk

`func (o *MultiFieldReqWithOptions) GetBirthdayOk() (*string, bool)`

GetBirthdayOk returns a tuple with the Birthday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthday

`func (o *MultiFieldReqWithOptions) SetBirthday(v string)`

SetBirthday sets Birthday field to given value.

### HasBirthday

`func (o *MultiFieldReqWithOptions) HasBirthday() bool`

HasBirthday returns a boolean if a field has been set.

### GetIpAddress

`func (o *MultiFieldReqWithOptions) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *MultiFieldReqWithOptions) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *MultiFieldReqWithOptions) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *MultiFieldReqWithOptions) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetCountries

`func (o *MultiFieldReqWithOptions) GetCountries() []string`

GetCountries returns the Countries field if non-nil, zero value otherwise.

### GetCountriesOk

`func (o *MultiFieldReqWithOptions) GetCountriesOk() (*[]string, bool)`

GetCountriesOk returns a tuple with the Countries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountries

`func (o *MultiFieldReqWithOptions) SetCountries(v []string)`

SetCountries sets Countries field to given value.

### HasCountries

`func (o *MultiFieldReqWithOptions) HasCountries() bool`

HasCountries returns a boolean if a field has been set.

### GetExcludeCountries

`func (o *MultiFieldReqWithOptions) GetExcludeCountries() []string`

GetExcludeCountries returns the ExcludeCountries field if non-nil, zero value otherwise.

### GetExcludeCountriesOk

`func (o *MultiFieldReqWithOptions) GetExcludeCountriesOk() (*[]string, bool)`

GetExcludeCountriesOk returns a tuple with the ExcludeCountries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeCountries

`func (o *MultiFieldReqWithOptions) SetExcludeCountries(v []string)`

SetExcludeCountries sets ExcludeCountries field to given value.

### HasExcludeCountries

`func (o *MultiFieldReqWithOptions) HasExcludeCountries() bool`

HasExcludeCountries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


