# VerifyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressLine1** | Pointer to **string** |  | [optional] 
**AddressLine2** | Pointer to **string** |  | [optional] 
**City** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**RegionCode** | Pointer to **string** |  | [optional] 
**Country** | Pointer to **string** |  | [optional] 
**Continent** | Pointer to **string** |  | [optional] 
**PostalCode** | Pointer to **string** |  | [optional] 
**FamilyName** | Pointer to **string** |  | [optional] 
**GivenName** | Pointer to **string** |  | [optional] 
**FullName** | Pointer to **string** |  | [optional] 
**Phone** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Maid** | Pointer to **string** |  | [optional] 
**Social** | Pointer to **string** |  | [optional] 
**NonId** | Pointer to **string** |  | [optional] 
**PanoramaId** | Pointer to **string** |  | [optional] 
**IpAddress** | Pointer to **string** |  | [optional] 
**Birthday** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Organization** | Pointer to **string** |  | [optional] 
**Risk** | Pointer to **float64** |  | [optional] 
**Evidence** | Pointer to [**Evidence**](Evidence.md) |  | [optional] 
**RiskV2** | Pointer to **float64** |  | [optional] 
**RiskV3** | Pointer to **float64** |  | [optional] 
**ScoreDetails** | Pointer to [**[]ScoreDetails**](ScoreDetails.md) |  | [optional] 

## Methods

### NewVerifyResponse

`func NewVerifyResponse() *VerifyResponse`

NewVerifyResponse instantiates a new VerifyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerifyResponseWithDefaults

`func NewVerifyResponseWithDefaults() *VerifyResponse`

NewVerifyResponseWithDefaults instantiates a new VerifyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressLine1

`func (o *VerifyResponse) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *VerifyResponse) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *VerifyResponse) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.

### HasAddressLine1

`func (o *VerifyResponse) HasAddressLine1() bool`

HasAddressLine1 returns a boolean if a field has been set.

### GetAddressLine2

`func (o *VerifyResponse) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *VerifyResponse) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *VerifyResponse) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *VerifyResponse) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### GetCity

`func (o *VerifyResponse) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *VerifyResponse) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *VerifyResponse) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *VerifyResponse) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetRegion

`func (o *VerifyResponse) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *VerifyResponse) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *VerifyResponse) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *VerifyResponse) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetRegionCode

`func (o *VerifyResponse) GetRegionCode() string`

GetRegionCode returns the RegionCode field if non-nil, zero value otherwise.

### GetRegionCodeOk

`func (o *VerifyResponse) GetRegionCodeOk() (*string, bool)`

GetRegionCodeOk returns a tuple with the RegionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionCode

`func (o *VerifyResponse) SetRegionCode(v string)`

SetRegionCode sets RegionCode field to given value.

### HasRegionCode

`func (o *VerifyResponse) HasRegionCode() bool`

HasRegionCode returns a boolean if a field has been set.

### GetCountry

`func (o *VerifyResponse) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *VerifyResponse) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *VerifyResponse) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *VerifyResponse) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetContinent

`func (o *VerifyResponse) GetContinent() string`

GetContinent returns the Continent field if non-nil, zero value otherwise.

### GetContinentOk

`func (o *VerifyResponse) GetContinentOk() (*string, bool)`

GetContinentOk returns a tuple with the Continent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinent

`func (o *VerifyResponse) SetContinent(v string)`

SetContinent sets Continent field to given value.

### HasContinent

`func (o *VerifyResponse) HasContinent() bool`

HasContinent returns a boolean if a field has been set.

### GetPostalCode

`func (o *VerifyResponse) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *VerifyResponse) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *VerifyResponse) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *VerifyResponse) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetFamilyName

`func (o *VerifyResponse) GetFamilyName() string`

GetFamilyName returns the FamilyName field if non-nil, zero value otherwise.

### GetFamilyNameOk

`func (o *VerifyResponse) GetFamilyNameOk() (*string, bool)`

GetFamilyNameOk returns a tuple with the FamilyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamilyName

`func (o *VerifyResponse) SetFamilyName(v string)`

SetFamilyName sets FamilyName field to given value.

### HasFamilyName

`func (o *VerifyResponse) HasFamilyName() bool`

HasFamilyName returns a boolean if a field has been set.

### GetGivenName

`func (o *VerifyResponse) GetGivenName() string`

GetGivenName returns the GivenName field if non-nil, zero value otherwise.

### GetGivenNameOk

`func (o *VerifyResponse) GetGivenNameOk() (*string, bool)`

GetGivenNameOk returns a tuple with the GivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGivenName

`func (o *VerifyResponse) SetGivenName(v string)`

SetGivenName sets GivenName field to given value.

### HasGivenName

`func (o *VerifyResponse) HasGivenName() bool`

HasGivenName returns a boolean if a field has been set.

### GetFullName

`func (o *VerifyResponse) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *VerifyResponse) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *VerifyResponse) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *VerifyResponse) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### GetPhone

`func (o *VerifyResponse) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *VerifyResponse) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *VerifyResponse) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *VerifyResponse) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetEmail

`func (o *VerifyResponse) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *VerifyResponse) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *VerifyResponse) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *VerifyResponse) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetMaid

`func (o *VerifyResponse) GetMaid() string`

GetMaid returns the Maid field if non-nil, zero value otherwise.

### GetMaidOk

`func (o *VerifyResponse) GetMaidOk() (*string, bool)`

GetMaidOk returns a tuple with the Maid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaid

`func (o *VerifyResponse) SetMaid(v string)`

SetMaid sets Maid field to given value.

### HasMaid

`func (o *VerifyResponse) HasMaid() bool`

HasMaid returns a boolean if a field has been set.

### GetSocial

`func (o *VerifyResponse) GetSocial() string`

GetSocial returns the Social field if non-nil, zero value otherwise.

### GetSocialOk

`func (o *VerifyResponse) GetSocialOk() (*string, bool)`

GetSocialOk returns a tuple with the Social field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocial

`func (o *VerifyResponse) SetSocial(v string)`

SetSocial sets Social field to given value.

### HasSocial

`func (o *VerifyResponse) HasSocial() bool`

HasSocial returns a boolean if a field has been set.

### GetNonId

`func (o *VerifyResponse) GetNonId() string`

GetNonId returns the NonId field if non-nil, zero value otherwise.

### GetNonIdOk

`func (o *VerifyResponse) GetNonIdOk() (*string, bool)`

GetNonIdOk returns a tuple with the NonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonId

`func (o *VerifyResponse) SetNonId(v string)`

SetNonId sets NonId field to given value.

### HasNonId

`func (o *VerifyResponse) HasNonId() bool`

HasNonId returns a boolean if a field has been set.

### GetPanoramaId

`func (o *VerifyResponse) GetPanoramaId() string`

GetPanoramaId returns the PanoramaId field if non-nil, zero value otherwise.

### GetPanoramaIdOk

`func (o *VerifyResponse) GetPanoramaIdOk() (*string, bool)`

GetPanoramaIdOk returns a tuple with the PanoramaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPanoramaId

`func (o *VerifyResponse) SetPanoramaId(v string)`

SetPanoramaId sets PanoramaId field to given value.

### HasPanoramaId

`func (o *VerifyResponse) HasPanoramaId() bool`

HasPanoramaId returns a boolean if a field has been set.

### GetIpAddress

`func (o *VerifyResponse) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *VerifyResponse) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *VerifyResponse) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *VerifyResponse) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetBirthday

`func (o *VerifyResponse) GetBirthday() string`

GetBirthday returns the Birthday field if non-nil, zero value otherwise.

### GetBirthdayOk

`func (o *VerifyResponse) GetBirthdayOk() (*string, bool)`

GetBirthdayOk returns a tuple with the Birthday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthday

`func (o *VerifyResponse) SetBirthday(v string)`

SetBirthday sets Birthday field to given value.

### HasBirthday

`func (o *VerifyResponse) HasBirthday() bool`

HasBirthday returns a boolean if a field has been set.

### GetTitle

`func (o *VerifyResponse) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *VerifyResponse) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *VerifyResponse) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *VerifyResponse) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetOrganization

`func (o *VerifyResponse) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *VerifyResponse) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *VerifyResponse) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *VerifyResponse) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetRisk

`func (o *VerifyResponse) GetRisk() float64`

GetRisk returns the Risk field if non-nil, zero value otherwise.

### GetRiskOk

`func (o *VerifyResponse) GetRiskOk() (*float64, bool)`

GetRiskOk returns a tuple with the Risk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRisk

`func (o *VerifyResponse) SetRisk(v float64)`

SetRisk sets Risk field to given value.

### HasRisk

`func (o *VerifyResponse) HasRisk() bool`

HasRisk returns a boolean if a field has been set.

### GetEvidence

`func (o *VerifyResponse) GetEvidence() Evidence`

GetEvidence returns the Evidence field if non-nil, zero value otherwise.

### GetEvidenceOk

`func (o *VerifyResponse) GetEvidenceOk() (*Evidence, bool)`

GetEvidenceOk returns a tuple with the Evidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidence

`func (o *VerifyResponse) SetEvidence(v Evidence)`

SetEvidence sets Evidence field to given value.

### HasEvidence

`func (o *VerifyResponse) HasEvidence() bool`

HasEvidence returns a boolean if a field has been set.

### GetRiskV2

`func (o *VerifyResponse) GetRiskV2() float64`

GetRiskV2 returns the RiskV2 field if non-nil, zero value otherwise.

### GetRiskV2Ok

`func (o *VerifyResponse) GetRiskV2Ok() (*float64, bool)`

GetRiskV2Ok returns a tuple with the RiskV2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskV2

`func (o *VerifyResponse) SetRiskV2(v float64)`

SetRiskV2 sets RiskV2 field to given value.

### HasRiskV2

`func (o *VerifyResponse) HasRiskV2() bool`

HasRiskV2 returns a boolean if a field has been set.

### GetRiskV3

`func (o *VerifyResponse) GetRiskV3() float64`

GetRiskV3 returns the RiskV3 field if non-nil, zero value otherwise.

### GetRiskV3Ok

`func (o *VerifyResponse) GetRiskV3Ok() (*float64, bool)`

GetRiskV3Ok returns a tuple with the RiskV3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskV3

`func (o *VerifyResponse) SetRiskV3(v float64)`

SetRiskV3 sets RiskV3 field to given value.

### HasRiskV3

`func (o *VerifyResponse) HasRiskV3() bool`

HasRiskV3 returns a boolean if a field has been set.

### GetScoreDetails

`func (o *VerifyResponse) GetScoreDetails() []ScoreDetails`

GetScoreDetails returns the ScoreDetails field if non-nil, zero value otherwise.

### GetScoreDetailsOk

`func (o *VerifyResponse) GetScoreDetailsOk() (*[]ScoreDetails, bool)`

GetScoreDetailsOk returns a tuple with the ScoreDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScoreDetails

`func (o *VerifyResponse) SetScoreDetails(v []ScoreDetails)`

SetScoreDetails sets ScoreDetails field to given value.

### HasScoreDetails

`func (o *VerifyResponse) HasScoreDetails() bool`

HasScoreDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


