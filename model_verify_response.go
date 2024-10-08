/*
Fideo API

Fideo Intelligence offers an identity intelligence product that protects the public good. - [Fideo Privacy Policy](https://www.fideo.ai/privacy-policy/)

API version: 1.0.4
Contact: support@fideo.ai
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fideo

import (
	"encoding/json"
)

// checks if the VerifyResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VerifyResponse{}

// VerifyResponse struct for VerifyResponse
type VerifyResponse struct {
	AddressLine1 *string `json:"addressLine1,omitempty"`
	AddressLine2 *string `json:"addressLine2,omitempty"`
	City *string `json:"city,omitempty"`
	Region *string `json:"region,omitempty"`
	RegionCode *string `json:"regionCode,omitempty"`
	Country *string `json:"country,omitempty"`
	Continent *string `json:"continent,omitempty"`
	PostalCode *string `json:"postalCode,omitempty"`
	FamilyName *string `json:"familyName,omitempty"`
	GivenName *string `json:"givenName,omitempty"`
	FullName *string `json:"fullName,omitempty"`
	Phone *string `json:"phone,omitempty"`
	Email *string `json:"email,omitempty"`
	Maid *string `json:"maid,omitempty"`
	Social *string `json:"social,omitempty"`
	NonId *string `json:"nonId,omitempty"`
	PanoramaId *string `json:"panoramaId,omitempty"`
	IpAddress *string `json:"ipAddress,omitempty"`
	Birthday *string `json:"birthday,omitempty"`
	Title *string `json:"title,omitempty"`
	Organization *string `json:"organization,omitempty"`
	Risk *float64 `json:"risk,omitempty"`
	Evidence *Evidence `json:"evidence,omitempty"`
	RiskV2 *float64 `json:"riskV2,omitempty"`
	RiskV3 *float64 `json:"riskV3,omitempty"`
	ScoreDetails []ScoreDetails `json:"scoreDetails,omitempty"`
}

// NewVerifyResponse instantiates a new VerifyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVerifyResponse() *VerifyResponse {
	this := VerifyResponse{}
	return &this
}

// NewVerifyResponseWithDefaults instantiates a new VerifyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVerifyResponseWithDefaults() *VerifyResponse {
	this := VerifyResponse{}
	return &this
}

// GetAddressLine1 returns the AddressLine1 field value if set, zero value otherwise.
func (o *VerifyResponse) GetAddressLine1() string {
	if o == nil || IsNil(o.AddressLine1) {
		var ret string
		return ret
	}
	return *o.AddressLine1
}

// GetAddressLine1Ok returns a tuple with the AddressLine1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyResponse) GetAddressLine1Ok() (*string, bool) {
	if o == nil || IsNil(o.AddressLine1) {
		return nil, false
	}
	return o.AddressLine1, true
}

// HasAddressLine1 returns a boolean if a field has been set.
func (o *VerifyResponse) HasAddressLine1() bool {
	if o != nil && !IsNil(o.AddressLine1) {
		return true
	}

	return false
}

// SetAddressLine1 gets a reference to the given string and assigns it to the AddressLine1 field.
func (o *VerifyResponse) SetAddressLine1(v string) {
	o.AddressLine1 = &v
}

// GetAddressLine2 returns the AddressLine2 field value if set, zero value otherwise.
func (o *VerifyResponse) GetAddressLine2() string {
	if o == nil || IsNil(o.AddressLine2) {
		var ret string
		return ret
	}
	return *o.AddressLine2
}

// GetAddressLine2Ok returns a tuple with the AddressLine2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyResponse) GetAddressLine2Ok() (*string, bool) {
	if o == nil || IsNil(o.AddressLine2) {
		return nil, false
	}
	return o.AddressLine2, true
}

// HasAddressLine2 returns a boolean if a field has been set.
func (o *VerifyResponse) HasAddressLine2() bool {
	if o != nil && !IsNil(o.AddressLine2) {
		return true
	}

	return false
}

// SetAddressLine2 gets a reference to the given string and assigns it to the AddressLine2 field.
func (o *VerifyResponse) SetAddressLine2(v string) {
	o.AddressLine2 = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *VerifyResponse) GetCity() string {
	if o == nil || IsNil(o.City) {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyResponse) GetCityOk() (*string, bool) {
	if o == nil || IsNil(o.City) {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *VerifyResponse) HasCity() bool {
	if o != nil && !IsNil(o.City) {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *VerifyResponse) SetCity(v string) {
	o.City = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *VerifyResponse) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyResponse) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *VerifyResponse) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *VerifyResponse) SetRegion(v string) {
	o.Region = &v
}

// GetRegionCode returns the RegionCode field value if set, zero value otherwise.
func (o *VerifyResponse) GetRegionCode() string {
	if o == nil || IsNil(o.RegionCode) {
		var ret string
		return ret
	}
	return *o.RegionCode
}

// GetRegionCodeOk returns a tuple with the RegionCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyResponse) GetRegionCodeOk() (*string, bool) {
	if o == nil || IsNil(o.RegionCode) {
		return nil, false
	}
	return o.RegionCode, true
}

// HasRegionCode returns a boolean if a field has been set.
func (o *VerifyResponse) HasRegionCode() bool {
	if o != nil && !IsNil(o.RegionCode) {
		return true
	}

	return false
}

// SetRegionCode gets a reference to the given string and assigns it to the RegionCode field.
func (o *VerifyResponse) SetRegionCode(v string) {
	o.RegionCode = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *VerifyResponse) GetCountry() string {
	if o == nil || IsNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyResponse) GetCountryOk() (*string, bool) {
	if o == nil || IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *VerifyResponse) HasCountry() bool {
	if o != nil && !IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *VerifyResponse) SetCountry(v string) {
	o.Country = &v
}

// GetContinent returns the Continent field value if set, zero value otherwise.
func (o *VerifyResponse) GetContinent() string {
	if o == nil || IsNil(o.Continent) {
		var ret string
		return ret
	}
	return *o.Continent
}

// GetContinentOk returns a tuple with the Continent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyResponse) GetContinentOk() (*string, bool) {
	if o == nil || IsNil(o.Continent) {
		return nil, false
	}
	return o.Continent, true
}

// HasContinent returns a boolean if a field has been set.
func (o *VerifyResponse) HasContinent() bool {
	if o != nil && !IsNil(o.Continent) {
		return true
	}

	return false
}

// SetContinent gets a reference to the given string and assigns it to the Continent field.
func (o *VerifyResponse) SetContinent(v string) {
	o.Continent = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *VerifyResponse) GetPostalCode() string {
	if o == nil || IsNil(o.PostalCode) {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyResponse) GetPostalCodeOk() (*string, bool) {
	if o == nil || IsNil(o.PostalCode) {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *VerifyResponse) HasPostalCode() bool {
	if o != nil && !IsNil(o.PostalCode) {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *VerifyResponse) SetPostalCode(v string) {
	o.PostalCode = &v
}

// GetFamilyName returns the FamilyName field value if set, zero value otherwise.
func (o *VerifyResponse) GetFamilyName() string {
	if o == nil || IsNil(o.FamilyName) {
		var ret string
		return ret
	}
	return *o.FamilyName
}

// GetFamilyNameOk returns a tuple with the FamilyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyResponse) GetFamilyNameOk() (*string, bool) {
	if o == nil || IsNil(o.FamilyName) {
		return nil, false
	}
	return o.FamilyName, true
}

// HasFamilyName returns a boolean if a field has been set.
func (o *VerifyResponse) HasFamilyName() bool {
	if o != nil && !IsNil(o.FamilyName) {
		return true
	}

	return false
}

// SetFamilyName gets a reference to the given string and assigns it to the FamilyName field.
func (o *VerifyResponse) SetFamilyName(v string) {
	o.FamilyName = &v
}

// GetGivenName returns the GivenName field value if set, zero value otherwise.
func (o *VerifyResponse) GetGivenName() string {
	if o == nil || IsNil(o.GivenName) {
		var ret string
		return ret
	}
	return *o.GivenName
}

// GetGivenNameOk returns a tuple with the GivenName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyResponse) GetGivenNameOk() (*string, bool) {
	if o == nil || IsNil(o.GivenName) {
		return nil, false
	}
	return o.GivenName, true
}

// HasGivenName returns a boolean if a field has been set.
func (o *VerifyResponse) HasGivenName() bool {
	if o != nil && !IsNil(o.GivenName) {
		return true
	}

	return false
}

// SetGivenName gets a reference to the given string and assigns it to the GivenName field.
func (o *VerifyResponse) SetGivenName(v string) {
	o.GivenName = &v
}

// GetFullName returns the FullName field value if set, zero value otherwise.
func (o *VerifyResponse) GetFullName() string {
	if o == nil || IsNil(o.FullName) {
		var ret string
		return ret
	}
	return *o.FullName
}

// GetFullNameOk returns a tuple with the FullName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyResponse) GetFullNameOk() (*string, bool) {
	if o == nil || IsNil(o.FullName) {
		return nil, false
	}
	return o.FullName, true
}

// HasFullName returns a boolean if a field has been set.
func (o *VerifyResponse) HasFullName() bool {
	if o != nil && !IsNil(o.FullName) {
		return true
	}

	return false
}

// SetFullName gets a reference to the given string and assigns it to the FullName field.
func (o *VerifyResponse) SetFullName(v string) {
	o.FullName = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *VerifyResponse) GetPhone() string {
	if o == nil || IsNil(o.Phone) {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyResponse) GetPhoneOk() (*string, bool) {
	if o == nil || IsNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *VerifyResponse) HasPhone() bool {
	if o != nil && !IsNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *VerifyResponse) SetPhone(v string) {
	o.Phone = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *VerifyResponse) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyResponse) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *VerifyResponse) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *VerifyResponse) SetEmail(v string) {
	o.Email = &v
}

// GetMaid returns the Maid field value if set, zero value otherwise.
func (o *VerifyResponse) GetMaid() string {
	if o == nil || IsNil(o.Maid) {
		var ret string
		return ret
	}
	return *o.Maid
}

// GetMaidOk returns a tuple with the Maid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyResponse) GetMaidOk() (*string, bool) {
	if o == nil || IsNil(o.Maid) {
		return nil, false
	}
	return o.Maid, true
}

// HasMaid returns a boolean if a field has been set.
func (o *VerifyResponse) HasMaid() bool {
	if o != nil && !IsNil(o.Maid) {
		return true
	}

	return false
}

// SetMaid gets a reference to the given string and assigns it to the Maid field.
func (o *VerifyResponse) SetMaid(v string) {
	o.Maid = &v
}

// GetSocial returns the Social field value if set, zero value otherwise.
func (o *VerifyResponse) GetSocial() string {
	if o == nil || IsNil(o.Social) {
		var ret string
		return ret
	}
	return *o.Social
}

// GetSocialOk returns a tuple with the Social field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyResponse) GetSocialOk() (*string, bool) {
	if o == nil || IsNil(o.Social) {
		return nil, false
	}
	return o.Social, true
}

// HasSocial returns a boolean if a field has been set.
func (o *VerifyResponse) HasSocial() bool {
	if o != nil && !IsNil(o.Social) {
		return true
	}

	return false
}

// SetSocial gets a reference to the given string and assigns it to the Social field.
func (o *VerifyResponse) SetSocial(v string) {
	o.Social = &v
}

// GetNonId returns the NonId field value if set, zero value otherwise.
func (o *VerifyResponse) GetNonId() string {
	if o == nil || IsNil(o.NonId) {
		var ret string
		return ret
	}
	return *o.NonId
}

// GetNonIdOk returns a tuple with the NonId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyResponse) GetNonIdOk() (*string, bool) {
	if o == nil || IsNil(o.NonId) {
		return nil, false
	}
	return o.NonId, true
}

// HasNonId returns a boolean if a field has been set.
func (o *VerifyResponse) HasNonId() bool {
	if o != nil && !IsNil(o.NonId) {
		return true
	}

	return false
}

// SetNonId gets a reference to the given string and assigns it to the NonId field.
func (o *VerifyResponse) SetNonId(v string) {
	o.NonId = &v
}

// GetPanoramaId returns the PanoramaId field value if set, zero value otherwise.
func (o *VerifyResponse) GetPanoramaId() string {
	if o == nil || IsNil(o.PanoramaId) {
		var ret string
		return ret
	}
	return *o.PanoramaId
}

// GetPanoramaIdOk returns a tuple with the PanoramaId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyResponse) GetPanoramaIdOk() (*string, bool) {
	if o == nil || IsNil(o.PanoramaId) {
		return nil, false
	}
	return o.PanoramaId, true
}

// HasPanoramaId returns a boolean if a field has been set.
func (o *VerifyResponse) HasPanoramaId() bool {
	if o != nil && !IsNil(o.PanoramaId) {
		return true
	}

	return false
}

// SetPanoramaId gets a reference to the given string and assigns it to the PanoramaId field.
func (o *VerifyResponse) SetPanoramaId(v string) {
	o.PanoramaId = &v
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise.
func (o *VerifyResponse) GetIpAddress() string {
	if o == nil || IsNil(o.IpAddress) {
		var ret string
		return ret
	}
	return *o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyResponse) GetIpAddressOk() (*string, bool) {
	if o == nil || IsNil(o.IpAddress) {
		return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *VerifyResponse) HasIpAddress() bool {
	if o != nil && !IsNil(o.IpAddress) {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given string and assigns it to the IpAddress field.
func (o *VerifyResponse) SetIpAddress(v string) {
	o.IpAddress = &v
}

// GetBirthday returns the Birthday field value if set, zero value otherwise.
func (o *VerifyResponse) GetBirthday() string {
	if o == nil || IsNil(o.Birthday) {
		var ret string
		return ret
	}
	return *o.Birthday
}

// GetBirthdayOk returns a tuple with the Birthday field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyResponse) GetBirthdayOk() (*string, bool) {
	if o == nil || IsNil(o.Birthday) {
		return nil, false
	}
	return o.Birthday, true
}

// HasBirthday returns a boolean if a field has been set.
func (o *VerifyResponse) HasBirthday() bool {
	if o != nil && !IsNil(o.Birthday) {
		return true
	}

	return false
}

// SetBirthday gets a reference to the given string and assigns it to the Birthday field.
func (o *VerifyResponse) SetBirthday(v string) {
	o.Birthday = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *VerifyResponse) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyResponse) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *VerifyResponse) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *VerifyResponse) SetTitle(v string) {
	o.Title = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *VerifyResponse) GetOrganization() string {
	if o == nil || IsNil(o.Organization) {
		var ret string
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyResponse) GetOrganizationOk() (*string, bool) {
	if o == nil || IsNil(o.Organization) {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *VerifyResponse) HasOrganization() bool {
	if o != nil && !IsNil(o.Organization) {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given string and assigns it to the Organization field.
func (o *VerifyResponse) SetOrganization(v string) {
	o.Organization = &v
}

// GetRisk returns the Risk field value if set, zero value otherwise.
func (o *VerifyResponse) GetRisk() float64 {
	if o == nil || IsNil(o.Risk) {
		var ret float64
		return ret
	}
	return *o.Risk
}

// GetRiskOk returns a tuple with the Risk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyResponse) GetRiskOk() (*float64, bool) {
	if o == nil || IsNil(o.Risk) {
		return nil, false
	}
	return o.Risk, true
}

// HasRisk returns a boolean if a field has been set.
func (o *VerifyResponse) HasRisk() bool {
	if o != nil && !IsNil(o.Risk) {
		return true
	}

	return false
}

// SetRisk gets a reference to the given float64 and assigns it to the Risk field.
func (o *VerifyResponse) SetRisk(v float64) {
	o.Risk = &v
}

// GetEvidence returns the Evidence field value if set, zero value otherwise.
func (o *VerifyResponse) GetEvidence() Evidence {
	if o == nil || IsNil(o.Evidence) {
		var ret Evidence
		return ret
	}
	return *o.Evidence
}

// GetEvidenceOk returns a tuple with the Evidence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyResponse) GetEvidenceOk() (*Evidence, bool) {
	if o == nil || IsNil(o.Evidence) {
		return nil, false
	}
	return o.Evidence, true
}

// HasEvidence returns a boolean if a field has been set.
func (o *VerifyResponse) HasEvidence() bool {
	if o != nil && !IsNil(o.Evidence) {
		return true
	}

	return false
}

// SetEvidence gets a reference to the given Evidence and assigns it to the Evidence field.
func (o *VerifyResponse) SetEvidence(v Evidence) {
	o.Evidence = &v
}

// GetRiskV2 returns the RiskV2 field value if set, zero value otherwise.
func (o *VerifyResponse) GetRiskV2() float64 {
	if o == nil || IsNil(o.RiskV2) {
		var ret float64
		return ret
	}
	return *o.RiskV2
}

// GetRiskV2Ok returns a tuple with the RiskV2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyResponse) GetRiskV2Ok() (*float64, bool) {
	if o == nil || IsNil(o.RiskV2) {
		return nil, false
	}
	return o.RiskV2, true
}

// HasRiskV2 returns a boolean if a field has been set.
func (o *VerifyResponse) HasRiskV2() bool {
	if o != nil && !IsNil(o.RiskV2) {
		return true
	}

	return false
}

// SetRiskV2 gets a reference to the given float64 and assigns it to the RiskV2 field.
func (o *VerifyResponse) SetRiskV2(v float64) {
	o.RiskV2 = &v
}

// GetRiskV3 returns the RiskV3 field value if set, zero value otherwise.
func (o *VerifyResponse) GetRiskV3() float64 {
	if o == nil || IsNil(o.RiskV3) {
		var ret float64
		return ret
	}
	return *o.RiskV3
}

// GetRiskV3Ok returns a tuple with the RiskV3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyResponse) GetRiskV3Ok() (*float64, bool) {
	if o == nil || IsNil(o.RiskV3) {
		return nil, false
	}
	return o.RiskV3, true
}

// HasRiskV3 returns a boolean if a field has been set.
func (o *VerifyResponse) HasRiskV3() bool {
	if o != nil && !IsNil(o.RiskV3) {
		return true
	}

	return false
}

// SetRiskV3 gets a reference to the given float64 and assigns it to the RiskV3 field.
func (o *VerifyResponse) SetRiskV3(v float64) {
	o.RiskV3 = &v
}

// GetScoreDetails returns the ScoreDetails field value if set, zero value otherwise.
func (o *VerifyResponse) GetScoreDetails() []ScoreDetails {
	if o == nil || IsNil(o.ScoreDetails) {
		var ret []ScoreDetails
		return ret
	}
	return o.ScoreDetails
}

// GetScoreDetailsOk returns a tuple with the ScoreDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyResponse) GetScoreDetailsOk() ([]ScoreDetails, bool) {
	if o == nil || IsNil(o.ScoreDetails) {
		return nil, false
	}
	return o.ScoreDetails, true
}

// HasScoreDetails returns a boolean if a field has been set.
func (o *VerifyResponse) HasScoreDetails() bool {
	if o != nil && !IsNil(o.ScoreDetails) {
		return true
	}

	return false
}

// SetScoreDetails gets a reference to the given []ScoreDetails and assigns it to the ScoreDetails field.
func (o *VerifyResponse) SetScoreDetails(v []ScoreDetails) {
	o.ScoreDetails = v
}

func (o VerifyResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VerifyResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AddressLine1) {
		toSerialize["addressLine1"] = o.AddressLine1
	}
	if !IsNil(o.AddressLine2) {
		toSerialize["addressLine2"] = o.AddressLine2
	}
	if !IsNil(o.City) {
		toSerialize["city"] = o.City
	}
	if !IsNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	if !IsNil(o.RegionCode) {
		toSerialize["regionCode"] = o.RegionCode
	}
	if !IsNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	if !IsNil(o.Continent) {
		toSerialize["continent"] = o.Continent
	}
	if !IsNil(o.PostalCode) {
		toSerialize["postalCode"] = o.PostalCode
	}
	if !IsNil(o.FamilyName) {
		toSerialize["familyName"] = o.FamilyName
	}
	if !IsNil(o.GivenName) {
		toSerialize["givenName"] = o.GivenName
	}
	if !IsNil(o.FullName) {
		toSerialize["fullName"] = o.FullName
	}
	if !IsNil(o.Phone) {
		toSerialize["phone"] = o.Phone
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.Maid) {
		toSerialize["maid"] = o.Maid
	}
	if !IsNil(o.Social) {
		toSerialize["social"] = o.Social
	}
	if !IsNil(o.NonId) {
		toSerialize["nonId"] = o.NonId
	}
	if !IsNil(o.PanoramaId) {
		toSerialize["panoramaId"] = o.PanoramaId
	}
	if !IsNil(o.IpAddress) {
		toSerialize["ipAddress"] = o.IpAddress
	}
	if !IsNil(o.Birthday) {
		toSerialize["birthday"] = o.Birthday
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Organization) {
		toSerialize["organization"] = o.Organization
	}
	if !IsNil(o.Risk) {
		toSerialize["risk"] = o.Risk
	}
	if !IsNil(o.Evidence) {
		toSerialize["evidence"] = o.Evidence
	}
	if !IsNil(o.RiskV2) {
		toSerialize["riskV2"] = o.RiskV2
	}
	if !IsNil(o.RiskV3) {
		toSerialize["riskV3"] = o.RiskV3
	}
	if !IsNil(o.ScoreDetails) {
		toSerialize["scoreDetails"] = o.ScoreDetails
	}
	return toSerialize, nil
}

type NullableVerifyResponse struct {
	value *VerifyResponse
	isSet bool
}

func (v NullableVerifyResponse) Get() *VerifyResponse {
	return v.value
}

func (v *NullableVerifyResponse) Set(val *VerifyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableVerifyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableVerifyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVerifyResponse(val *VerifyResponse) *NullableVerifyResponse {
	return &NullableVerifyResponse{value: val, isSet: true}
}

func (v NullableVerifyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVerifyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


