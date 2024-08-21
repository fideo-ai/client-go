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

// checks if the Education type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Education{}

// Education struct for Education
type Education struct {
	Name *string `json:"name,omitempty"`
	Degree *string `json:"degree,omitempty"`
	End *EducationDate `json:"end,omitempty"`
}

// NewEducation instantiates a new Education object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEducation() *Education {
	this := Education{}
	return &this
}

// NewEducationWithDefaults instantiates a new Education object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEducationWithDefaults() *Education {
	this := Education{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Education) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Education) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Education) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Education) SetName(v string) {
	o.Name = &v
}

// GetDegree returns the Degree field value if set, zero value otherwise.
func (o *Education) GetDegree() string {
	if o == nil || IsNil(o.Degree) {
		var ret string
		return ret
	}
	return *o.Degree
}

// GetDegreeOk returns a tuple with the Degree field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Education) GetDegreeOk() (*string, bool) {
	if o == nil || IsNil(o.Degree) {
		return nil, false
	}
	return o.Degree, true
}

// HasDegree returns a boolean if a field has been set.
func (o *Education) HasDegree() bool {
	if o != nil && !IsNil(o.Degree) {
		return true
	}

	return false
}

// SetDegree gets a reference to the given string and assigns it to the Degree field.
func (o *Education) SetDegree(v string) {
	o.Degree = &v
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *Education) GetEnd() EducationDate {
	if o == nil || IsNil(o.End) {
		var ret EducationDate
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Education) GetEndOk() (*EducationDate, bool) {
	if o == nil || IsNil(o.End) {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *Education) HasEnd() bool {
	if o != nil && !IsNil(o.End) {
		return true
	}

	return false
}

// SetEnd gets a reference to the given EducationDate and assigns it to the End field.
func (o *Education) SetEnd(v EducationDate) {
	o.End = &v
}

func (o Education) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Education) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Degree) {
		toSerialize["degree"] = o.Degree
	}
	if !IsNil(o.End) {
		toSerialize["end"] = o.End
	}
	return toSerialize, nil
}

type NullableEducation struct {
	value *Education
	isSet bool
}

func (v NullableEducation) Get() *Education {
	return v.value
}

func (v *NullableEducation) Set(val *Education) {
	v.value = val
	v.isSet = true
}

func (v NullableEducation) IsSet() bool {
	return v.isSet
}

func (v *NullableEducation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEducation(val *Education) *NullableEducation {
	return &NullableEducation{value: val, isSet: true}
}

func (v NullableEducation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEducation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


