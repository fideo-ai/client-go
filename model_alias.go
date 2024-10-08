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

// checks if the Alias type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Alias{}

// Alias struct for Alias
type Alias struct {
	First *string `json:"first,omitempty"`
	Last *string `json:"last,omitempty"`
	Middle *string `json:"middle,omitempty"`
}

// NewAlias instantiates a new Alias object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlias() *Alias {
	this := Alias{}
	return &this
}

// NewAliasWithDefaults instantiates a new Alias object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAliasWithDefaults() *Alias {
	this := Alias{}
	return &this
}

// GetFirst returns the First field value if set, zero value otherwise.
func (o *Alias) GetFirst() string {
	if o == nil || IsNil(o.First) {
		var ret string
		return ret
	}
	return *o.First
}

// GetFirstOk returns a tuple with the First field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Alias) GetFirstOk() (*string, bool) {
	if o == nil || IsNil(o.First) {
		return nil, false
	}
	return o.First, true
}

// HasFirst returns a boolean if a field has been set.
func (o *Alias) HasFirst() bool {
	if o != nil && !IsNil(o.First) {
		return true
	}

	return false
}

// SetFirst gets a reference to the given string and assigns it to the First field.
func (o *Alias) SetFirst(v string) {
	o.First = &v
}

// GetLast returns the Last field value if set, zero value otherwise.
func (o *Alias) GetLast() string {
	if o == nil || IsNil(o.Last) {
		var ret string
		return ret
	}
	return *o.Last
}

// GetLastOk returns a tuple with the Last field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Alias) GetLastOk() (*string, bool) {
	if o == nil || IsNil(o.Last) {
		return nil, false
	}
	return o.Last, true
}

// HasLast returns a boolean if a field has been set.
func (o *Alias) HasLast() bool {
	if o != nil && !IsNil(o.Last) {
		return true
	}

	return false
}

// SetLast gets a reference to the given string and assigns it to the Last field.
func (o *Alias) SetLast(v string) {
	o.Last = &v
}

// GetMiddle returns the Middle field value if set, zero value otherwise.
func (o *Alias) GetMiddle() string {
	if o == nil || IsNil(o.Middle) {
		var ret string
		return ret
	}
	return *o.Middle
}

// GetMiddleOk returns a tuple with the Middle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Alias) GetMiddleOk() (*string, bool) {
	if o == nil || IsNil(o.Middle) {
		return nil, false
	}
	return o.Middle, true
}

// HasMiddle returns a boolean if a field has been set.
func (o *Alias) HasMiddle() bool {
	if o != nil && !IsNil(o.Middle) {
		return true
	}

	return false
}

// SetMiddle gets a reference to the given string and assigns it to the Middle field.
func (o *Alias) SetMiddle(v string) {
	o.Middle = &v
}

func (o Alias) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Alias) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.First) {
		toSerialize["first"] = o.First
	}
	if !IsNil(o.Last) {
		toSerialize["last"] = o.Last
	}
	if !IsNil(o.Middle) {
		toSerialize["middle"] = o.Middle
	}
	return toSerialize, nil
}

type NullableAlias struct {
	value *Alias
	isSet bool
}

func (v NullableAlias) Get() *Alias {
	return v.value
}

func (v *NullableAlias) Set(val *Alias) {
	v.value = val
	v.isSet = true
}

func (v NullableAlias) IsSet() bool {
	return v.isSet
}

func (v *NullableAlias) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlias(val *Alias) *NullableAlias {
	return &NullableAlias{value: val, isSet: true}
}

func (v NullableAlias) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlias) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


