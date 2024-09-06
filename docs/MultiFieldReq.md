# MultiFieldReq

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

## Methods

### NewMultiFieldReq

`func NewMultiFieldReq() *MultiFieldReq`

NewMultiFieldReq instantiates a new MultiFieldReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultiFieldReqWithDefaults

`func NewMultiFieldReqWithDefaults() *MultiFieldReq`

NewMultiFieldReqWithDefaults instantiates a new MultiFieldReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTwitter

`func (o *MultiFieldReq) GetTwitter() string`

GetTwitter returns the Twitter field if non-nil, zero value otherwise.

### GetTwitterOk

`func (o *MultiFieldReq) GetTwitterOk() (*string, bool)`

GetTwitterOk returns a tuple with the Twitter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitter

`func (o *MultiFieldReq) SetTwitter(v string)`

SetTwitter sets Twitter field to given value.

### HasTwitter

`func (o *MultiFieldReq) HasTwitter() bool`

HasTwitter returns a boolean if a field has been set.

### GetLinkedin

`func (o *MultiFieldReq) GetLinkedin() string`

GetLinkedin returns the Linkedin field if non-nil, zero value otherwise.

### GetLinkedinOk

`func (o *MultiFieldReq) GetLinkedinOk() (*string, bool)`

GetLinkedinOk returns a tuple with the Linkedin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedin

`func (o *MultiFieldReq) SetLinkedin(v string)`

SetLinkedin sets Linkedin field to given value.

### HasLinkedin

`func (o *MultiFieldReq) HasLinkedin() bool`

HasLinkedin returns a boolean if a field has been set.

### GetRecordId

`func (o *MultiFieldReq) GetRecordId() string`

GetRecordId returns the RecordId field if non-nil, zero value otherwise.

### GetRecordIdOk

`func (o *MultiFieldReq) GetRecordIdOk() (*string, bool)`

GetRecordIdOk returns a tuple with the RecordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordId

`func (o *MultiFieldReq) SetRecordId(v string)`

SetRecordId sets RecordId field to given value.

### HasRecordId

`func (o *MultiFieldReq) HasRecordId() bool`

HasRecordId returns a boolean if a field has been set.

### GetPersonId

`func (o *MultiFieldReq) GetPersonId() string`

GetPersonId returns the PersonId field if non-nil, zero value otherwise.

### GetPersonIdOk

`func (o *MultiFieldReq) GetPersonIdOk() (*string, bool)`

GetPersonIdOk returns a tuple with the PersonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonId

`func (o *MultiFieldReq) SetPersonId(v string)`

SetPersonId sets PersonId field to given value.

### HasPersonId

`func (o *MultiFieldReq) HasPersonId() bool`

HasPersonId returns a boolean if a field has been set.

### GetPartnerId

`func (o *MultiFieldReq) GetPartnerId() string`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *MultiFieldReq) GetPartnerIdOk() (*string, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *MultiFieldReq) SetPartnerId(v string)`

SetPartnerId sets PartnerId field to given value.

### HasPartnerId

`func (o *MultiFieldReq) HasPartnerId() bool`

HasPartnerId returns a boolean if a field has been set.

### GetLocation

`func (o *MultiFieldReq) GetLocation() LocationReq`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *MultiFieldReq) GetLocationOk() (*LocationReq, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *MultiFieldReq) SetLocation(v LocationReq)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *MultiFieldReq) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetAvatar

`func (o *MultiFieldReq) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *MultiFieldReq) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *MultiFieldReq) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *MultiFieldReq) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### GetWebsite

`func (o *MultiFieldReq) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *MultiFieldReq) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *MultiFieldReq) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *MultiFieldReq) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.

### GetTitle

`func (o *MultiFieldReq) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *MultiFieldReq) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *MultiFieldReq) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *MultiFieldReq) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetOrganization

`func (o *MultiFieldReq) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *MultiFieldReq) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *MultiFieldReq) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *MultiFieldReq) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetEmails

`func (o *MultiFieldReq) GetEmails() []string`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *MultiFieldReq) GetEmailsOk() (*[]string, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *MultiFieldReq) SetEmails(v []string)`

SetEmails sets Emails field to given value.

### HasEmails

`func (o *MultiFieldReq) HasEmails() bool`

HasEmails returns a boolean if a field has been set.

### GetPhones

`func (o *MultiFieldReq) GetPhones() []string`

GetPhones returns the Phones field if non-nil, zero value otherwise.

### GetPhonesOk

`func (o *MultiFieldReq) GetPhonesOk() (*[]string, bool)`

GetPhonesOk returns a tuple with the Phones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhones

`func (o *MultiFieldReq) SetPhones(v []string)`

SetPhones sets Phones field to given value.

### HasPhones

`func (o *MultiFieldReq) HasPhones() bool`

HasPhones returns a boolean if a field has been set.

### GetProfiles

`func (o *MultiFieldReq) GetProfiles() []SocialProfileReq`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *MultiFieldReq) GetProfilesOk() (*[]SocialProfileReq, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *MultiFieldReq) SetProfiles(v []SocialProfileReq)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *MultiFieldReq) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.

### GetMaids

`func (o *MultiFieldReq) GetMaids() []string`

GetMaids returns the Maids field if non-nil, zero value otherwise.

### GetMaidsOk

`func (o *MultiFieldReq) GetMaidsOk() (*[]string, bool)`

GetMaidsOk returns a tuple with the Maids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaids

`func (o *MultiFieldReq) SetMaids(v []string)`

SetMaids sets Maids field to given value.

### HasMaids

`func (o *MultiFieldReq) HasMaids() bool`

HasMaids returns a boolean if a field has been set.

### GetName

`func (o *MultiFieldReq) GetName() PersonNameReq`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MultiFieldReq) GetNameOk() (*PersonNameReq, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MultiFieldReq) SetName(v PersonNameReq)`

SetName sets Name field to given value.

### HasName

`func (o *MultiFieldReq) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPartnerKeys

`func (o *MultiFieldReq) GetPartnerKeys() map[string]string`

GetPartnerKeys returns the PartnerKeys field if non-nil, zero value otherwise.

### GetPartnerKeysOk

`func (o *MultiFieldReq) GetPartnerKeysOk() (*map[string]string, bool)`

GetPartnerKeysOk returns a tuple with the PartnerKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerKeys

`func (o *MultiFieldReq) SetPartnerKeys(v map[string]string)`

SetPartnerKeys sets PartnerKeys field to given value.

### HasPartnerKeys

`func (o *MultiFieldReq) HasPartnerKeys() bool`

HasPartnerKeys returns a boolean if a field has been set.

### GetLiNonid

`func (o *MultiFieldReq) GetLiNonid() string`

GetLiNonid returns the LiNonid field if non-nil, zero value otherwise.

### GetLiNonidOk

`func (o *MultiFieldReq) GetLiNonidOk() (*string, bool)`

GetLiNonidOk returns a tuple with the LiNonid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiNonid

`func (o *MultiFieldReq) SetLiNonid(v string)`

SetLiNonid sets LiNonid field to given value.

### HasLiNonid

`func (o *MultiFieldReq) HasLiNonid() bool`

HasLiNonid returns a boolean if a field has been set.

### GetPanoramaId

`func (o *MultiFieldReq) GetPanoramaId() string`

GetPanoramaId returns the PanoramaId field if non-nil, zero value otherwise.

### GetPanoramaIdOk

`func (o *MultiFieldReq) GetPanoramaIdOk() (*string, bool)`

GetPanoramaIdOk returns a tuple with the PanoramaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPanoramaId

`func (o *MultiFieldReq) SetPanoramaId(v string)`

SetPanoramaId sets PanoramaId field to given value.

### HasPanoramaId

`func (o *MultiFieldReq) HasPanoramaId() bool`

HasPanoramaId returns a boolean if a field has been set.

### GetPlacekey

`func (o *MultiFieldReq) GetPlacekey() string`

GetPlacekey returns the Placekey field if non-nil, zero value otherwise.

### GetPlacekeyOk

`func (o *MultiFieldReq) GetPlacekeyOk() (*string, bool)`

GetPlacekeyOk returns a tuple with the Placekey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacekey

`func (o *MultiFieldReq) SetPlacekey(v string)`

SetPlacekey sets Placekey field to given value.

### HasPlacekey

`func (o *MultiFieldReq) HasPlacekey() bool`

HasPlacekey returns a boolean if a field has been set.

### GetGeneratePid

`func (o *MultiFieldReq) GetGeneratePid() bool`

GetGeneratePid returns the GeneratePid field if non-nil, zero value otherwise.

### GetGeneratePidOk

`func (o *MultiFieldReq) GetGeneratePidOk() (*bool, bool)`

GetGeneratePidOk returns a tuple with the GeneratePid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratePid

`func (o *MultiFieldReq) SetGeneratePid(v bool)`

SetGeneratePid sets GeneratePid field to given value.

### HasGeneratePid

`func (o *MultiFieldReq) HasGeneratePid() bool`

HasGeneratePid returns a boolean if a field has been set.

### GetEmail

`func (o *MultiFieldReq) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *MultiFieldReq) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *MultiFieldReq) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *MultiFieldReq) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPhone

`func (o *MultiFieldReq) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *MultiFieldReq) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *MultiFieldReq) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *MultiFieldReq) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetProfile

`func (o *MultiFieldReq) GetProfile() SocialProfileReq`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *MultiFieldReq) GetProfileOk() (*SocialProfileReq, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *MultiFieldReq) SetProfile(v SocialProfileReq)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *MultiFieldReq) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetMaid

`func (o *MultiFieldReq) GetMaid() string`

GetMaid returns the Maid field if non-nil, zero value otherwise.

### GetMaidOk

`func (o *MultiFieldReq) GetMaidOk() (*string, bool)`

GetMaidOk returns a tuple with the Maid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaid

`func (o *MultiFieldReq) SetMaid(v string)`

SetMaid sets Maid field to given value.

### HasMaid

`func (o *MultiFieldReq) HasMaid() bool`

HasMaid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


