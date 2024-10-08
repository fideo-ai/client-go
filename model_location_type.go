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
	"fmt"
)

// LocationType the model 'LocationType'
type LocationType string

// List of LocationType
const (
	WORK LocationType = "Work"
	PRIMARY LocationType = "Primary"
	SECONDARY LocationType = "Secondary"
)

// All allowed values of LocationType enum
var AllowedLocationTypeEnumValues = []LocationType{
	"Work",
	"Primary",
	"Secondary",
}

func (v *LocationType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LocationType(value)
	for _, existing := range AllowedLocationTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LocationType", value)
}

// NewLocationTypeFromValue returns a pointer to a valid LocationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLocationTypeFromValue(v string) (*LocationType, error) {
	ev := LocationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LocationType: valid values are %v", v, AllowedLocationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LocationType) IsValid() bool {
	for _, existing := range AllowedLocationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LocationType value
func (v LocationType) Ptr() *LocationType {
	return &v
}

type NullableLocationType struct {
	value *LocationType
	isSet bool
}

func (v NullableLocationType) Get() *LocationType {
	return v.value
}

func (v *NullableLocationType) Set(val *LocationType) {
	v.value = val
	v.isSet = true
}

func (v NullableLocationType) IsSet() bool {
	return v.isSet
}

func (v *NullableLocationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocationType(val *LocationType) *NullableLocationType {
	return &NullableLocationType{value: val, isSet: true}
}

func (v NullableLocationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

