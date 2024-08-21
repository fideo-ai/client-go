/*
Fideo API

Fideo Intelligence offers an identity intelligence product that protects the public good. - [Fideo Privacy Policy](https://www.fideo.ai/privacy-policy/)

API version: 1.0.3
Contact: support@fideo.ai
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fideo

import (
	"encoding/json"
)

// checks if the LocationReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LocationReq{}

// LocationReq struct for LocationReq
type LocationReq struct {
	AddressLine1 *string `json:"addressLine1,omitempty"`
	AddressLine2 *string `json:"addressLine2,omitempty"`
	City *string `json:"city,omitempty"`
	Region *string `json:"region,omitempty"`
	RegionCode *string `json:"regionCode,omitempty"`
	PostalCode *string `json:"postalCode,omitempty"`
	Country *string `json:"country,omitempty"`
	CountryCode *string `json:"countryCode,omitempty"`
	Latitude *float64 `json:"latitude,omitempty"`
	Longitude *float64 `json:"longitude,omitempty"`
	TimeZone *string `json:"timeZone,omitempty"`
	Formatted *string `json:"formatted,omitempty"`
	Type *LocationType `json:"type,omitempty"`
}

// NewLocationReq instantiates a new LocationReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocationReq() *LocationReq {
	this := LocationReq{}
	return &this
}

// NewLocationReqWithDefaults instantiates a new LocationReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocationReqWithDefaults() *LocationReq {
	this := LocationReq{}
	return &this
}

// GetAddressLine1 returns the AddressLine1 field value if set, zero value otherwise.
func (o *LocationReq) GetAddressLine1() string {
	if o == nil || IsNil(o.AddressLine1) {
		var ret string
		return ret
	}
	return *o.AddressLine1
}

// GetAddressLine1Ok returns a tuple with the AddressLine1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationReq) GetAddressLine1Ok() (*string, bool) {
	if o == nil || IsNil(o.AddressLine1) {
		return nil, false
	}
	return o.AddressLine1, true
}

// HasAddressLine1 returns a boolean if a field has been set.
func (o *LocationReq) HasAddressLine1() bool {
	if o != nil && !IsNil(o.AddressLine1) {
		return true
	}

	return false
}

// SetAddressLine1 gets a reference to the given string and assigns it to the AddressLine1 field.
func (o *LocationReq) SetAddressLine1(v string) {
	o.AddressLine1 = &v
}

// GetAddressLine2 returns the AddressLine2 field value if set, zero value otherwise.
func (o *LocationReq) GetAddressLine2() string {
	if o == nil || IsNil(o.AddressLine2) {
		var ret string
		return ret
	}
	return *o.AddressLine2
}

// GetAddressLine2Ok returns a tuple with the AddressLine2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationReq) GetAddressLine2Ok() (*string, bool) {
	if o == nil || IsNil(o.AddressLine2) {
		return nil, false
	}
	return o.AddressLine2, true
}

// HasAddressLine2 returns a boolean if a field has been set.
func (o *LocationReq) HasAddressLine2() bool {
	if o != nil && !IsNil(o.AddressLine2) {
		return true
	}

	return false
}

// SetAddressLine2 gets a reference to the given string and assigns it to the AddressLine2 field.
func (o *LocationReq) SetAddressLine2(v string) {
	o.AddressLine2 = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *LocationReq) GetCity() string {
	if o == nil || IsNil(o.City) {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationReq) GetCityOk() (*string, bool) {
	if o == nil || IsNil(o.City) {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *LocationReq) HasCity() bool {
	if o != nil && !IsNil(o.City) {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *LocationReq) SetCity(v string) {
	o.City = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *LocationReq) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationReq) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *LocationReq) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *LocationReq) SetRegion(v string) {
	o.Region = &v
}

// GetRegionCode returns the RegionCode field value if set, zero value otherwise.
func (o *LocationReq) GetRegionCode() string {
	if o == nil || IsNil(o.RegionCode) {
		var ret string
		return ret
	}
	return *o.RegionCode
}

// GetRegionCodeOk returns a tuple with the RegionCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationReq) GetRegionCodeOk() (*string, bool) {
	if o == nil || IsNil(o.RegionCode) {
		return nil, false
	}
	return o.RegionCode, true
}

// HasRegionCode returns a boolean if a field has been set.
func (o *LocationReq) HasRegionCode() bool {
	if o != nil && !IsNil(o.RegionCode) {
		return true
	}

	return false
}

// SetRegionCode gets a reference to the given string and assigns it to the RegionCode field.
func (o *LocationReq) SetRegionCode(v string) {
	o.RegionCode = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *LocationReq) GetPostalCode() string {
	if o == nil || IsNil(o.PostalCode) {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationReq) GetPostalCodeOk() (*string, bool) {
	if o == nil || IsNil(o.PostalCode) {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *LocationReq) HasPostalCode() bool {
	if o != nil && !IsNil(o.PostalCode) {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *LocationReq) SetPostalCode(v string) {
	o.PostalCode = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *LocationReq) GetCountry() string {
	if o == nil || IsNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationReq) GetCountryOk() (*string, bool) {
	if o == nil || IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *LocationReq) HasCountry() bool {
	if o != nil && !IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *LocationReq) SetCountry(v string) {
	o.Country = &v
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *LocationReq) GetCountryCode() string {
	if o == nil || IsNil(o.CountryCode) {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationReq) GetCountryCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CountryCode) {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *LocationReq) HasCountryCode() bool {
	if o != nil && !IsNil(o.CountryCode) {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *LocationReq) SetCountryCode(v string) {
	o.CountryCode = &v
}

// GetLatitude returns the Latitude field value if set, zero value otherwise.
func (o *LocationReq) GetLatitude() float64 {
	if o == nil || IsNil(o.Latitude) {
		var ret float64
		return ret
	}
	return *o.Latitude
}

// GetLatitudeOk returns a tuple with the Latitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationReq) GetLatitudeOk() (*float64, bool) {
	if o == nil || IsNil(o.Latitude) {
		return nil, false
	}
	return o.Latitude, true
}

// HasLatitude returns a boolean if a field has been set.
func (o *LocationReq) HasLatitude() bool {
	if o != nil && !IsNil(o.Latitude) {
		return true
	}

	return false
}

// SetLatitude gets a reference to the given float64 and assigns it to the Latitude field.
func (o *LocationReq) SetLatitude(v float64) {
	o.Latitude = &v
}

// GetLongitude returns the Longitude field value if set, zero value otherwise.
func (o *LocationReq) GetLongitude() float64 {
	if o == nil || IsNil(o.Longitude) {
		var ret float64
		return ret
	}
	return *o.Longitude
}

// GetLongitudeOk returns a tuple with the Longitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationReq) GetLongitudeOk() (*float64, bool) {
	if o == nil || IsNil(o.Longitude) {
		return nil, false
	}
	return o.Longitude, true
}

// HasLongitude returns a boolean if a field has been set.
func (o *LocationReq) HasLongitude() bool {
	if o != nil && !IsNil(o.Longitude) {
		return true
	}

	return false
}

// SetLongitude gets a reference to the given float64 and assigns it to the Longitude field.
func (o *LocationReq) SetLongitude(v float64) {
	o.Longitude = &v
}

// GetTimeZone returns the TimeZone field value if set, zero value otherwise.
func (o *LocationReq) GetTimeZone() string {
	if o == nil || IsNil(o.TimeZone) {
		var ret string
		return ret
	}
	return *o.TimeZone
}

// GetTimeZoneOk returns a tuple with the TimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationReq) GetTimeZoneOk() (*string, bool) {
	if o == nil || IsNil(o.TimeZone) {
		return nil, false
	}
	return o.TimeZone, true
}

// HasTimeZone returns a boolean if a field has been set.
func (o *LocationReq) HasTimeZone() bool {
	if o != nil && !IsNil(o.TimeZone) {
		return true
	}

	return false
}

// SetTimeZone gets a reference to the given string and assigns it to the TimeZone field.
func (o *LocationReq) SetTimeZone(v string) {
	o.TimeZone = &v
}

// GetFormatted returns the Formatted field value if set, zero value otherwise.
func (o *LocationReq) GetFormatted() string {
	if o == nil || IsNil(o.Formatted) {
		var ret string
		return ret
	}
	return *o.Formatted
}

// GetFormattedOk returns a tuple with the Formatted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationReq) GetFormattedOk() (*string, bool) {
	if o == nil || IsNil(o.Formatted) {
		return nil, false
	}
	return o.Formatted, true
}

// HasFormatted returns a boolean if a field has been set.
func (o *LocationReq) HasFormatted() bool {
	if o != nil && !IsNil(o.Formatted) {
		return true
	}

	return false
}

// SetFormatted gets a reference to the given string and assigns it to the Formatted field.
func (o *LocationReq) SetFormatted(v string) {
	o.Formatted = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *LocationReq) GetType() LocationType {
	if o == nil || IsNil(o.Type) {
		var ret LocationType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationReq) GetTypeOk() (*LocationType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *LocationReq) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given LocationType and assigns it to the Type field.
func (o *LocationReq) SetType(v LocationType) {
	o.Type = &v
}

func (o LocationReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LocationReq) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.PostalCode) {
		toSerialize["postalCode"] = o.PostalCode
	}
	if !IsNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	if !IsNil(o.CountryCode) {
		toSerialize["countryCode"] = o.CountryCode
	}
	if !IsNil(o.Latitude) {
		toSerialize["latitude"] = o.Latitude
	}
	if !IsNil(o.Longitude) {
		toSerialize["longitude"] = o.Longitude
	}
	if !IsNil(o.TimeZone) {
		toSerialize["timeZone"] = o.TimeZone
	}
	if !IsNil(o.Formatted) {
		toSerialize["formatted"] = o.Formatted
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableLocationReq struct {
	value *LocationReq
	isSet bool
}

func (v NullableLocationReq) Get() *LocationReq {
	return v.value
}

func (v *NullableLocationReq) Set(val *LocationReq) {
	v.value = val
	v.isSet = true
}

func (v NullableLocationReq) IsSet() bool {
	return v.isSet
}

func (v *NullableLocationReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocationReq(val *LocationReq) *NullableLocationReq {
	return &NullableLocationReq{value: val, isSet: true}
}

func (v NullableLocationReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocationReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


