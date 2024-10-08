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

// checks if the SignalsResponseV0 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SignalsResponseV0{}

// SignalsResponseV0 struct for SignalsResponseV0
type SignalsResponseV0 struct {
	Name *Name `json:"name,omitempty"`
	Demographics *Demographics `json:"demographics,omitempty"`
	Locations []Location `json:"locations,omitempty"`
	Emails []Email `json:"emails,omitempty"`
	Phones []Phone `json:"phones,omitempty"`
	PersonIds []string `json:"personIds,omitempty"`
	IpAddresses []IpAddress `json:"ipAddresses,omitempty"`
	Message *string `json:"message,omitempty"`
	SocialProfiles *SocialProfileUrls `json:"socialProfiles,omitempty"`
	Employment *Employment `json:"employment,omitempty"`
}

// NewSignalsResponseV0 instantiates a new SignalsResponseV0 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignalsResponseV0() *SignalsResponseV0 {
	this := SignalsResponseV0{}
	return &this
}

// NewSignalsResponseV0WithDefaults instantiates a new SignalsResponseV0 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignalsResponseV0WithDefaults() *SignalsResponseV0 {
	this := SignalsResponseV0{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SignalsResponseV0) GetName() Name {
	if o == nil || IsNil(o.Name) {
		var ret Name
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignalsResponseV0) GetNameOk() (*Name, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SignalsResponseV0) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given Name and assigns it to the Name field.
func (o *SignalsResponseV0) SetName(v Name) {
	o.Name = &v
}

// GetDemographics returns the Demographics field value if set, zero value otherwise.
func (o *SignalsResponseV0) GetDemographics() Demographics {
	if o == nil || IsNil(o.Demographics) {
		var ret Demographics
		return ret
	}
	return *o.Demographics
}

// GetDemographicsOk returns a tuple with the Demographics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignalsResponseV0) GetDemographicsOk() (*Demographics, bool) {
	if o == nil || IsNil(o.Demographics) {
		return nil, false
	}
	return o.Demographics, true
}

// HasDemographics returns a boolean if a field has been set.
func (o *SignalsResponseV0) HasDemographics() bool {
	if o != nil && !IsNil(o.Demographics) {
		return true
	}

	return false
}

// SetDemographics gets a reference to the given Demographics and assigns it to the Demographics field.
func (o *SignalsResponseV0) SetDemographics(v Demographics) {
	o.Demographics = &v
}

// GetLocations returns the Locations field value if set, zero value otherwise.
func (o *SignalsResponseV0) GetLocations() []Location {
	if o == nil || IsNil(o.Locations) {
		var ret []Location
		return ret
	}
	return o.Locations
}

// GetLocationsOk returns a tuple with the Locations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignalsResponseV0) GetLocationsOk() ([]Location, bool) {
	if o == nil || IsNil(o.Locations) {
		return nil, false
	}
	return o.Locations, true
}

// HasLocations returns a boolean if a field has been set.
func (o *SignalsResponseV0) HasLocations() bool {
	if o != nil && !IsNil(o.Locations) {
		return true
	}

	return false
}

// SetLocations gets a reference to the given []Location and assigns it to the Locations field.
func (o *SignalsResponseV0) SetLocations(v []Location) {
	o.Locations = v
}

// GetEmails returns the Emails field value if set, zero value otherwise.
func (o *SignalsResponseV0) GetEmails() []Email {
	if o == nil || IsNil(o.Emails) {
		var ret []Email
		return ret
	}
	return o.Emails
}

// GetEmailsOk returns a tuple with the Emails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignalsResponseV0) GetEmailsOk() ([]Email, bool) {
	if o == nil || IsNil(o.Emails) {
		return nil, false
	}
	return o.Emails, true
}

// HasEmails returns a boolean if a field has been set.
func (o *SignalsResponseV0) HasEmails() bool {
	if o != nil && !IsNil(o.Emails) {
		return true
	}

	return false
}

// SetEmails gets a reference to the given []Email and assigns it to the Emails field.
func (o *SignalsResponseV0) SetEmails(v []Email) {
	o.Emails = v
}

// GetPhones returns the Phones field value if set, zero value otherwise.
func (o *SignalsResponseV0) GetPhones() []Phone {
	if o == nil || IsNil(o.Phones) {
		var ret []Phone
		return ret
	}
	return o.Phones
}

// GetPhonesOk returns a tuple with the Phones field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignalsResponseV0) GetPhonesOk() ([]Phone, bool) {
	if o == nil || IsNil(o.Phones) {
		return nil, false
	}
	return o.Phones, true
}

// HasPhones returns a boolean if a field has been set.
func (o *SignalsResponseV0) HasPhones() bool {
	if o != nil && !IsNil(o.Phones) {
		return true
	}

	return false
}

// SetPhones gets a reference to the given []Phone and assigns it to the Phones field.
func (o *SignalsResponseV0) SetPhones(v []Phone) {
	o.Phones = v
}

// GetPersonIds returns the PersonIds field value if set, zero value otherwise.
func (o *SignalsResponseV0) GetPersonIds() []string {
	if o == nil || IsNil(o.PersonIds) {
		var ret []string
		return ret
	}
	return o.PersonIds
}

// GetPersonIdsOk returns a tuple with the PersonIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignalsResponseV0) GetPersonIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.PersonIds) {
		return nil, false
	}
	return o.PersonIds, true
}

// HasPersonIds returns a boolean if a field has been set.
func (o *SignalsResponseV0) HasPersonIds() bool {
	if o != nil && !IsNil(o.PersonIds) {
		return true
	}

	return false
}

// SetPersonIds gets a reference to the given []string and assigns it to the PersonIds field.
func (o *SignalsResponseV0) SetPersonIds(v []string) {
	o.PersonIds = v
}

// GetIpAddresses returns the IpAddresses field value if set, zero value otherwise.
func (o *SignalsResponseV0) GetIpAddresses() []IpAddress {
	if o == nil || IsNil(o.IpAddresses) {
		var ret []IpAddress
		return ret
	}
	return o.IpAddresses
}

// GetIpAddressesOk returns a tuple with the IpAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignalsResponseV0) GetIpAddressesOk() ([]IpAddress, bool) {
	if o == nil || IsNil(o.IpAddresses) {
		return nil, false
	}
	return o.IpAddresses, true
}

// HasIpAddresses returns a boolean if a field has been set.
func (o *SignalsResponseV0) HasIpAddresses() bool {
	if o != nil && !IsNil(o.IpAddresses) {
		return true
	}

	return false
}

// SetIpAddresses gets a reference to the given []IpAddress and assigns it to the IpAddresses field.
func (o *SignalsResponseV0) SetIpAddresses(v []IpAddress) {
	o.IpAddresses = v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *SignalsResponseV0) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignalsResponseV0) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *SignalsResponseV0) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *SignalsResponseV0) SetMessage(v string) {
	o.Message = &v
}

// GetSocialProfiles returns the SocialProfiles field value if set, zero value otherwise.
func (o *SignalsResponseV0) GetSocialProfiles() SocialProfileUrls {
	if o == nil || IsNil(o.SocialProfiles) {
		var ret SocialProfileUrls
		return ret
	}
	return *o.SocialProfiles
}

// GetSocialProfilesOk returns a tuple with the SocialProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignalsResponseV0) GetSocialProfilesOk() (*SocialProfileUrls, bool) {
	if o == nil || IsNil(o.SocialProfiles) {
		return nil, false
	}
	return o.SocialProfiles, true
}

// HasSocialProfiles returns a boolean if a field has been set.
func (o *SignalsResponseV0) HasSocialProfiles() bool {
	if o != nil && !IsNil(o.SocialProfiles) {
		return true
	}

	return false
}

// SetSocialProfiles gets a reference to the given SocialProfileUrls and assigns it to the SocialProfiles field.
func (o *SignalsResponseV0) SetSocialProfiles(v SocialProfileUrls) {
	o.SocialProfiles = &v
}

// GetEmployment returns the Employment field value if set, zero value otherwise.
func (o *SignalsResponseV0) GetEmployment() Employment {
	if o == nil || IsNil(o.Employment) {
		var ret Employment
		return ret
	}
	return *o.Employment
}

// GetEmploymentOk returns a tuple with the Employment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignalsResponseV0) GetEmploymentOk() (*Employment, bool) {
	if o == nil || IsNil(o.Employment) {
		return nil, false
	}
	return o.Employment, true
}

// HasEmployment returns a boolean if a field has been set.
func (o *SignalsResponseV0) HasEmployment() bool {
	if o != nil && !IsNil(o.Employment) {
		return true
	}

	return false
}

// SetEmployment gets a reference to the given Employment and assigns it to the Employment field.
func (o *SignalsResponseV0) SetEmployment(v Employment) {
	o.Employment = &v
}

func (o SignalsResponseV0) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SignalsResponseV0) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Demographics) {
		toSerialize["demographics"] = o.Demographics
	}
	if !IsNil(o.Locations) {
		toSerialize["locations"] = o.Locations
	}
	if !IsNil(o.Emails) {
		toSerialize["emails"] = o.Emails
	}
	if !IsNil(o.Phones) {
		toSerialize["phones"] = o.Phones
	}
	if !IsNil(o.PersonIds) {
		toSerialize["personIds"] = o.PersonIds
	}
	if !IsNil(o.IpAddresses) {
		toSerialize["ipAddresses"] = o.IpAddresses
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.SocialProfiles) {
		toSerialize["socialProfiles"] = o.SocialProfiles
	}
	if !IsNil(o.Employment) {
		toSerialize["employment"] = o.Employment
	}
	return toSerialize, nil
}

type NullableSignalsResponseV0 struct {
	value *SignalsResponseV0
	isSet bool
}

func (v NullableSignalsResponseV0) Get() *SignalsResponseV0 {
	return v.value
}

func (v *NullableSignalsResponseV0) Set(val *SignalsResponseV0) {
	v.value = val
	v.isSet = true
}

func (v NullableSignalsResponseV0) IsSet() bool {
	return v.isSet
}

func (v *NullableSignalsResponseV0) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignalsResponseV0(val *SignalsResponseV0) *NullableSignalsResponseV0 {
	return &NullableSignalsResponseV0{value: val, isSet: true}
}

func (v NullableSignalsResponseV0) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignalsResponseV0) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


