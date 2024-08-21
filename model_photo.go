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

// checks if the Photo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Photo{}

// Photo struct for Photo
type Photo struct {
	Url *string `json:"url,omitempty"`
	Label *string `json:"label,omitempty"`
}

// NewPhoto instantiates a new Photo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPhoto() *Photo {
	this := Photo{}
	return &this
}

// NewPhotoWithDefaults instantiates a new Photo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPhotoWithDefaults() *Photo {
	this := Photo{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *Photo) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Photo) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *Photo) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *Photo) SetUrl(v string) {
	o.Url = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *Photo) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Photo) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *Photo) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *Photo) SetLabel(v string) {
	o.Label = &v
}

func (o Photo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Photo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	return toSerialize, nil
}

type NullablePhoto struct {
	value *Photo
	isSet bool
}

func (v NullablePhoto) Get() *Photo {
	return v.value
}

func (v *NullablePhoto) Set(val *Photo) {
	v.value = val
	v.isSet = true
}

func (v NullablePhoto) IsSet() bool {
	return v.isSet
}

func (v *NullablePhoto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePhoto(val *Photo) *NullablePhoto {
	return &NullablePhoto{value: val, isSet: true}
}

func (v NullablePhoto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePhoto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


