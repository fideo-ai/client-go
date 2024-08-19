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
	"fmt"
)

// IPCountry the model 'IPCountry'
type IPCountry string

// List of IPCountry
const (
	DOMESTIC IPCountry = "DOMESTIC"
	FOREIGN IPCountry = "FOREIGN"
	UNKNOWN IPCountry = "UNKNOWN"
)

// All allowed values of IPCountry enum
var AllowedIPCountryEnumValues = []IPCountry{
	"DOMESTIC",
	"FOREIGN",
	"UNKNOWN",
}

func (v *IPCountry) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := IPCountry(value)
	for _, existing := range AllowedIPCountryEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid IPCountry", value)
}

// NewIPCountryFromValue returns a pointer to a valid IPCountry
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewIPCountryFromValue(v string) (*IPCountry, error) {
	ev := IPCountry(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for IPCountry: valid values are %v", v, AllowedIPCountryEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v IPCountry) IsValid() bool {
	for _, existing := range AllowedIPCountryEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IPCountry value
func (v IPCountry) Ptr() *IPCountry {
	return &v
}

type NullableIPCountry struct {
	value *IPCountry
	isSet bool
}

func (v NullableIPCountry) Get() *IPCountry {
	return v.value
}

func (v *NullableIPCountry) Set(val *IPCountry) {
	v.value = val
	v.isSet = true
}

func (v NullableIPCountry) IsSet() bool {
	return v.isSet
}

func (v *NullableIPCountry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIPCountry(val *IPCountry) *NullableIPCountry {
	return &NullableIPCountry{value: val, isSet: true}
}

func (v NullableIPCountry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIPCountry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

