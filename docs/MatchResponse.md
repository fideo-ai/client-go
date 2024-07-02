# MatchResponse

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

### NewMatchResponse

`func NewMatchResponse() *MatchResponse`

NewMatchResponse instantiates a new MatchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMatchResponseWithDefaults

`func NewMatchResponseWithDefaults() *MatchResponse`

NewMatchResponseWithDefaults instantiates a new MatchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressLine1

`func (o *MatchResponse) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *MatchResponse) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *MatchResponse) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.

### HasAddressLine1

`func (o *MatchResponse) HasAddressLine1() bool`

HasAddressLine1 returns a boolean if a field has been set.

### GetAddressLine2

`func (o *MatchResponse) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *MatchResponse) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *MatchResponse) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *MatchResponse) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### GetCity

`func (o *MatchResponse) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *MatchResponse) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *MatchResponse) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *MatchResponse) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetRegion

`func (o *MatchResponse) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *MatchResponse) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *MatchResponse) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *MatchResponse) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetRegionCode

`func (o *MatchResponse) GetRegionCode() string`

GetRegionCode returns the RegionCode field if non-nil, zero value otherwise.

### GetRegionCodeOk

`func (o *MatchResponse) GetRegionCodeOk() (*string, bool)`

GetRegionCodeOk returns a tuple with the RegionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionCode

`func (o *MatchResponse) SetRegionCode(v string)`

SetRegionCode sets RegionCode field to given value.

### HasRegionCode

`func (o *MatchResponse) HasRegionCode() bool`

HasRegionCode returns a boolean if a field has been set.

### GetCountry

`func (o *MatchResponse) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *MatchResponse) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *MatchResponse) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *MatchResponse) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetContinent

`func (o *MatchResponse) GetContinent() string`

GetContinent returns the Continent field if non-nil, zero value otherwise.

### GetContinentOk

`func (o *MatchResponse) GetContinentOk() (*string, bool)`

GetContinentOk returns a tuple with the Continent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinent

`func (o *MatchResponse) SetContinent(v string)`

SetContinent sets Continent field to given value.

### HasContinent

`func (o *MatchResponse) HasContinent() bool`

HasContinent returns a boolean if a field has been set.

### GetPostalCode

`func (o *MatchResponse) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *MatchResponse) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *MatchResponse) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *MatchResponse) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetFamilyName

`func (o *MatchResponse) GetFamilyName() string`

GetFamilyName returns the FamilyName field if non-nil, zero value otherwise.

### GetFamilyNameOk

`func (o *MatchResponse) GetFamilyNameOk() (*string, bool)`

GetFamilyNameOk returns a tuple with the FamilyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamilyName

`func (o *MatchResponse) SetFamilyName(v string)`

SetFamilyName sets FamilyName field to given value.

### HasFamilyName

`func (o *MatchResponse) HasFamilyName() bool`

HasFamilyName returns a boolean if a field has been set.

### GetGivenName

`func (o *MatchResponse) GetGivenName() string`

GetGivenName returns the GivenName field if non-nil, zero value otherwise.

### GetGivenNameOk

`func (o *MatchResponse) GetGivenNameOk() (*string, bool)`

GetGivenNameOk returns a tuple with the GivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGivenName

`func (o *MatchResponse) SetGivenName(v string)`

SetGivenName sets GivenName field to given value.

### HasGivenName

`func (o *MatchResponse) HasGivenName() bool`

HasGivenName returns a boolean if a field has been set.

### GetFullName

`func (o *MatchResponse) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *MatchResponse) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *MatchResponse) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *MatchResponse) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### GetPhone

`func (o *MatchResponse) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *MatchResponse) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *MatchResponse) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *MatchResponse) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetEmail

`func (o *MatchResponse) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *MatchResponse) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *MatchResponse) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *MatchResponse) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetMaid

`func (o *MatchResponse) GetMaid() string`

GetMaid returns the Maid field if non-nil, zero value otherwise.

### GetMaidOk

`func (o *MatchResponse) GetMaidOk() (*string, bool)`

GetMaidOk returns a tuple with the Maid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaid

`func (o *MatchResponse) SetMaid(v string)`

SetMaid sets Maid field to given value.

### HasMaid

`func (o *MatchResponse) HasMaid() bool`

HasMaid returns a boolean if a field has been set.

### GetSocial

`func (o *MatchResponse) GetSocial() string`

GetSocial returns the Social field if non-nil, zero value otherwise.

### GetSocialOk

`func (o *MatchResponse) GetSocialOk() (*string, bool)`

GetSocialOk returns a tuple with the Social field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocial

`func (o *MatchResponse) SetSocial(v string)`

SetSocial sets Social field to given value.

### HasSocial

`func (o *MatchResponse) HasSocial() bool`

HasSocial returns a boolean if a field has been set.

### GetNonId

`func (o *MatchResponse) GetNonId() string`

GetNonId returns the NonId field if non-nil, zero value otherwise.

### GetNonIdOk

`func (o *MatchResponse) GetNonIdOk() (*string, bool)`

GetNonIdOk returns a tuple with the NonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonId

`func (o *MatchResponse) SetNonId(v string)`

SetNonId sets NonId field to given value.

### HasNonId

`func (o *MatchResponse) HasNonId() bool`

HasNonId returns a boolean if a field has been set.

### GetPanoramaId

`func (o *MatchResponse) GetPanoramaId() string`

GetPanoramaId returns the PanoramaId field if non-nil, zero value otherwise.

### GetPanoramaIdOk

`func (o *MatchResponse) GetPanoramaIdOk() (*string, bool)`

GetPanoramaIdOk returns a tuple with the PanoramaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPanoramaId

`func (o *MatchResponse) SetPanoramaId(v string)`

SetPanoramaId sets PanoramaId field to given value.

### HasPanoramaId

`func (o *MatchResponse) HasPanoramaId() bool`

HasPanoramaId returns a boolean if a field has been set.

### GetIpAddress

`func (o *MatchResponse) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *MatchResponse) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *MatchResponse) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *MatchResponse) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetBirthday

`func (o *MatchResponse) GetBirthday() string`

GetBirthday returns the Birthday field if non-nil, zero value otherwise.

### GetBirthdayOk

`func (o *MatchResponse) GetBirthdayOk() (*string, bool)`

GetBirthdayOk returns a tuple with the Birthday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthday

`func (o *MatchResponse) SetBirthday(v string)`

SetBirthday sets Birthday field to given value.

### HasBirthday

`func (o *MatchResponse) HasBirthday() bool`

HasBirthday returns a boolean if a field has been set.

### GetTitle

`func (o *MatchResponse) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *MatchResponse) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *MatchResponse) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *MatchResponse) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetOrganization

`func (o *MatchResponse) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *MatchResponse) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *MatchResponse) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *MatchResponse) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetRisk

`func (o *MatchResponse) GetRisk() float64`

GetRisk returns the Risk field if non-nil, zero value otherwise.

### GetRiskOk

`func (o *MatchResponse) GetRiskOk() (*float64, bool)`

GetRiskOk returns a tuple with the Risk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRisk

`func (o *MatchResponse) SetRisk(v float64)`

SetRisk sets Risk field to given value.

### HasRisk

`func (o *MatchResponse) HasRisk() bool`

HasRisk returns a boolean if a field has been set.

### GetEvidence

`func (o *MatchResponse) GetEvidence() Evidence`

GetEvidence returns the Evidence field if non-nil, zero value otherwise.

### GetEvidenceOk

`func (o *MatchResponse) GetEvidenceOk() (*Evidence, bool)`

GetEvidenceOk returns a tuple with the Evidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidence

`func (o *MatchResponse) SetEvidence(v Evidence)`

SetEvidence sets Evidence field to given value.

### HasEvidence

`func (o *MatchResponse) HasEvidence() bool`

HasEvidence returns a boolean if a field has been set.

### GetRiskV2

`func (o *MatchResponse) GetRiskV2() float64`

GetRiskV2 returns the RiskV2 field if non-nil, zero value otherwise.

### GetRiskV2Ok

`func (o *MatchResponse) GetRiskV2Ok() (*float64, bool)`

GetRiskV2Ok returns a tuple with the RiskV2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskV2

`func (o *MatchResponse) SetRiskV2(v float64)`

SetRiskV2 sets RiskV2 field to given value.

### HasRiskV2

`func (o *MatchResponse) HasRiskV2() bool`

HasRiskV2 returns a boolean if a field has been set.

### GetRiskV3

`func (o *MatchResponse) GetRiskV3() float64`

GetRiskV3 returns the RiskV3 field if non-nil, zero value otherwise.

### GetRiskV3Ok

`func (o *MatchResponse) GetRiskV3Ok() (*float64, bool)`

GetRiskV3Ok returns a tuple with the RiskV3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskV3

`func (o *MatchResponse) SetRiskV3(v float64)`

SetRiskV3 sets RiskV3 field to given value.

### HasRiskV3

`func (o *MatchResponse) HasRiskV3() bool`

HasRiskV3 returns a boolean if a field has been set.

### GetScoreDetails

`func (o *MatchResponse) GetScoreDetails() []ScoreDetails`

GetScoreDetails returns the ScoreDetails field if non-nil, zero value otherwise.

### GetScoreDetailsOk

`func (o *MatchResponse) GetScoreDetailsOk() (*[]ScoreDetails, bool)`

GetScoreDetailsOk returns a tuple with the ScoreDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScoreDetails

`func (o *MatchResponse) SetScoreDetails(v []ScoreDetails)`

SetScoreDetails sets ScoreDetails field to given value.

### HasScoreDetails

`func (o *MatchResponse) HasScoreDetails() bool`

HasScoreDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


