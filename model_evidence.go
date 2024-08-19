/*
Fideo API

Fideo Intelligence offers an identity intelligence product that protects the public good. - [Fideo Privacy Policy](https://www.fideo.ai/privacy-policy/)

API version: 1.0.2
Contact: support@fideo.ai
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fideo

import (
	"encoding/json"
)

// checks if the Evidence type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Evidence{}

// Evidence struct for Evidence
type Evidence struct {
	IpTor *bool `json:"ipTor,omitempty"`
	IpCountry *IPCountry `json:"ipCountry,omitempty"`
	CountryOfIp *string `json:"countryOfIp,omitempty"`
}

// NewEvidence instantiates a new Evidence object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEvidence() *Evidence {
	this := Evidence{}
	return &this
}

// NewEvidenceWithDefaults instantiates a new Evidence object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEvidenceWithDefaults() *Evidence {
	this := Evidence{}
	return &this
}

// GetIpTor returns the IpTor field value if set, zero value otherwise.
func (o *Evidence) GetIpTor() bool {
	if o == nil || IsNil(o.IpTor) {
		var ret bool
		return ret
	}
	return *o.IpTor
}

// GetIpTorOk returns a tuple with the IpTor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Evidence) GetIpTorOk() (*bool, bool) {
	if o == nil || IsNil(o.IpTor) {
		return nil, false
	}
	return o.IpTor, true
}

// HasIpTor returns a boolean if a field has been set.
func (o *Evidence) HasIpTor() bool {
	if o != nil && !IsNil(o.IpTor) {
		return true
	}

	return false
}

// SetIpTor gets a reference to the given bool and assigns it to the IpTor field.
func (o *Evidence) SetIpTor(v bool) {
	o.IpTor = &v
}

// GetIpCountry returns the IpCountry field value if set, zero value otherwise.
func (o *Evidence) GetIpCountry() IPCountry {
	if o == nil || IsNil(o.IpCountry) {
		var ret IPCountry
		return ret
	}
	return *o.IpCountry
}

// GetIpCountryOk returns a tuple with the IpCountry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Evidence) GetIpCountryOk() (*IPCountry, bool) {
	if o == nil || IsNil(o.IpCountry) {
		return nil, false
	}
	return o.IpCountry, true
}

// HasIpCountry returns a boolean if a field has been set.
func (o *Evidence) HasIpCountry() bool {
	if o != nil && !IsNil(o.IpCountry) {
		return true
	}

	return false
}

// SetIpCountry gets a reference to the given IPCountry and assigns it to the IpCountry field.
func (o *Evidence) SetIpCountry(v IPCountry) {
	o.IpCountry = &v
}

// GetCountryOfIp returns the CountryOfIp field value if set, zero value otherwise.
func (o *Evidence) GetCountryOfIp() string {
	if o == nil || IsNil(o.CountryOfIp) {
		var ret string
		return ret
	}
	return *o.CountryOfIp
}

// GetCountryOfIpOk returns a tuple with the CountryOfIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Evidence) GetCountryOfIpOk() (*string, bool) {
	if o == nil || IsNil(o.CountryOfIp) {
		return nil, false
	}
	return o.CountryOfIp, true
}

// HasCountryOfIp returns a boolean if a field has been set.
func (o *Evidence) HasCountryOfIp() bool {
	if o != nil && !IsNil(o.CountryOfIp) {
		return true
	}

	return false
}

// SetCountryOfIp gets a reference to the given string and assigns it to the CountryOfIp field.
func (o *Evidence) SetCountryOfIp(v string) {
	o.CountryOfIp = &v
}

func (o Evidence) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Evidence) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IpTor) {
		toSerialize["ipTor"] = o.IpTor
	}
	if !IsNil(o.IpCountry) {
		toSerialize["ipCountry"] = o.IpCountry
	}
	if !IsNil(o.CountryOfIp) {
		toSerialize["countryOfIp"] = o.CountryOfIp
	}
	return toSerialize, nil
}

type NullableEvidence struct {
	value *Evidence
	isSet bool
}

func (v NullableEvidence) Get() *Evidence {
	return v.value
}

func (v *NullableEvidence) Set(val *Evidence) {
	v.value = val
	v.isSet = true
}

func (v NullableEvidence) IsSet() bool {
	return v.isSet
}

func (v *NullableEvidence) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEvidence(val *Evidence) *NullableEvidence {
	return &NullableEvidence{value: val, isSet: true}
}

func (v NullableEvidence) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEvidence) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


