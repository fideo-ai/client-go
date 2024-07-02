/*
Fideo API

This is a representation of the Fideo API based on the OpenAPI 3.0 specification.  You can find out more about Swagger at [https://swagger.io](https://swagger.io). Some useful links: - [Fideo Privacy Policy](https://www.fideo.ai/privacy-policy/)

API version: 1.0.0
Contact: support@fideo.ai
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fideo

import (
	"encoding/json"
)

// checks if the NameWithAlias type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NameWithAlias{}

// NameWithAlias struct for NameWithAlias
type NameWithAlias struct {
	First *string `json:"first,omitempty"`
	Last *string `json:"last,omitempty"`
	Middle *string `json:"middle,omitempty"`
	GivenName *string `json:"givenName,omitempty"`
	FamilyName *string `json:"familyName,omitempty"`
	Aliases []Alias `json:"aliases,omitempty"`
}

// NewNameWithAlias instantiates a new NameWithAlias object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNameWithAlias() *NameWithAlias {
	this := NameWithAlias{}
	return &this
}

// NewNameWithAliasWithDefaults instantiates a new NameWithAlias object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNameWithAliasWithDefaults() *NameWithAlias {
	this := NameWithAlias{}
	return &this
}

// GetFirst returns the First field value if set, zero value otherwise.
func (o *NameWithAlias) GetFirst() string {
	if o == nil || IsNil(o.First) {
		var ret string
		return ret
	}
	return *o.First
}

// GetFirstOk returns a tuple with the First field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NameWithAlias) GetFirstOk() (*string, bool) {
	if o == nil || IsNil(o.First) {
		return nil, false
	}
	return o.First, true
}

// HasFirst returns a boolean if a field has been set.
func (o *NameWithAlias) HasFirst() bool {
	if o != nil && !IsNil(o.First) {
		return true
	}

	return false
}

// SetFirst gets a reference to the given string and assigns it to the First field.
func (o *NameWithAlias) SetFirst(v string) {
	o.First = &v
}

// GetLast returns the Last field value if set, zero value otherwise.
func (o *NameWithAlias) GetLast() string {
	if o == nil || IsNil(o.Last) {
		var ret string
		return ret
	}
	return *o.Last
}

// GetLastOk returns a tuple with the Last field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NameWithAlias) GetLastOk() (*string, bool) {
	if o == nil || IsNil(o.Last) {
		return nil, false
	}
	return o.Last, true
}

// HasLast returns a boolean if a field has been set.
func (o *NameWithAlias) HasLast() bool {
	if o != nil && !IsNil(o.Last) {
		return true
	}

	return false
}

// SetLast gets a reference to the given string and assigns it to the Last field.
func (o *NameWithAlias) SetLast(v string) {
	o.Last = &v
}

// GetMiddle returns the Middle field value if set, zero value otherwise.
func (o *NameWithAlias) GetMiddle() string {
	if o == nil || IsNil(o.Middle) {
		var ret string
		return ret
	}
	return *o.Middle
}

// GetMiddleOk returns a tuple with the Middle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NameWithAlias) GetMiddleOk() (*string, bool) {
	if o == nil || IsNil(o.Middle) {
		return nil, false
	}
	return o.Middle, true
}

// HasMiddle returns a boolean if a field has been set.
func (o *NameWithAlias) HasMiddle() bool {
	if o != nil && !IsNil(o.Middle) {
		return true
	}

	return false
}

// SetMiddle gets a reference to the given string and assigns it to the Middle field.
func (o *NameWithAlias) SetMiddle(v string) {
	o.Middle = &v
}

// GetGivenName returns the GivenName field value if set, zero value otherwise.
func (o *NameWithAlias) GetGivenName() string {
	if o == nil || IsNil(o.GivenName) {
		var ret string
		return ret
	}
	return *o.GivenName
}

// GetGivenNameOk returns a tuple with the GivenName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NameWithAlias) GetGivenNameOk() (*string, bool) {
	if o == nil || IsNil(o.GivenName) {
		return nil, false
	}
	return o.GivenName, true
}

// HasGivenName returns a boolean if a field has been set.
func (o *NameWithAlias) HasGivenName() bool {
	if o != nil && !IsNil(o.GivenName) {
		return true
	}

	return false
}

// SetGivenName gets a reference to the given string and assigns it to the GivenName field.
func (o *NameWithAlias) SetGivenName(v string) {
	o.GivenName = &v
}

// GetFamilyName returns the FamilyName field value if set, zero value otherwise.
func (o *NameWithAlias) GetFamilyName() string {
	if o == nil || IsNil(o.FamilyName) {
		var ret string
		return ret
	}
	return *o.FamilyName
}

// GetFamilyNameOk returns a tuple with the FamilyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NameWithAlias) GetFamilyNameOk() (*string, bool) {
	if o == nil || IsNil(o.FamilyName) {
		return nil, false
	}
	return o.FamilyName, true
}

// HasFamilyName returns a boolean if a field has been set.
func (o *NameWithAlias) HasFamilyName() bool {
	if o != nil && !IsNil(o.FamilyName) {
		return true
	}

	return false
}

// SetFamilyName gets a reference to the given string and assigns it to the FamilyName field.
func (o *NameWithAlias) SetFamilyName(v string) {
	o.FamilyName = &v
}

// GetAliases returns the Aliases field value if set, zero value otherwise.
func (o *NameWithAlias) GetAliases() []Alias {
	if o == nil || IsNil(o.Aliases) {
		var ret []Alias
		return ret
	}
	return o.Aliases
}

// GetAliasesOk returns a tuple with the Aliases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NameWithAlias) GetAliasesOk() ([]Alias, bool) {
	if o == nil || IsNil(o.Aliases) {
		return nil, false
	}
	return o.Aliases, true
}

// HasAliases returns a boolean if a field has been set.
func (o *NameWithAlias) HasAliases() bool {
	if o != nil && !IsNil(o.Aliases) {
		return true
	}

	return false
}

// SetAliases gets a reference to the given []Alias and assigns it to the Aliases field.
func (o *NameWithAlias) SetAliases(v []Alias) {
	o.Aliases = v
}

func (o NameWithAlias) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NameWithAlias) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.GivenName) {
		toSerialize["givenName"] = o.GivenName
	}
	if !IsNil(o.FamilyName) {
		toSerialize["familyName"] = o.FamilyName
	}
	if !IsNil(o.Aliases) {
		toSerialize["aliases"] = o.Aliases
	}
	return toSerialize, nil
}

type NullableNameWithAlias struct {
	value *NameWithAlias
	isSet bool
}

func (v NullableNameWithAlias) Get() *NameWithAlias {
	return v.value
}

func (v *NullableNameWithAlias) Set(val *NameWithAlias) {
	v.value = val
	v.isSet = true
}

func (v NullableNameWithAlias) IsSet() bool {
	return v.isSet
}

func (v *NullableNameWithAlias) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNameWithAlias(val *NameWithAlias) *NullableNameWithAlias {
	return &NullableNameWithAlias{value: val, isSet: true}
}

func (v NullableNameWithAlias) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNameWithAlias) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


