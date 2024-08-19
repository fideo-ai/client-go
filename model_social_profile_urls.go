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

// checks if the SocialProfileUrls type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SocialProfileUrls{}

// SocialProfileUrls struct for SocialProfileUrls
type SocialProfileUrls struct {
	TwitterUrl *string `json:"twitterUrl,omitempty"`
	LinkedInUrl *string `json:"linkedInUrl,omitempty"`
}

// NewSocialProfileUrls instantiates a new SocialProfileUrls object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSocialProfileUrls() *SocialProfileUrls {
	this := SocialProfileUrls{}
	return &this
}

// NewSocialProfileUrlsWithDefaults instantiates a new SocialProfileUrls object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSocialProfileUrlsWithDefaults() *SocialProfileUrls {
	this := SocialProfileUrls{}
	return &this
}

// GetTwitterUrl returns the TwitterUrl field value if set, zero value otherwise.
func (o *SocialProfileUrls) GetTwitterUrl() string {
	if o == nil || IsNil(o.TwitterUrl) {
		var ret string
		return ret
	}
	return *o.TwitterUrl
}

// GetTwitterUrlOk returns a tuple with the TwitterUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SocialProfileUrls) GetTwitterUrlOk() (*string, bool) {
	if o == nil || IsNil(o.TwitterUrl) {
		return nil, false
	}
	return o.TwitterUrl, true
}

// HasTwitterUrl returns a boolean if a field has been set.
func (o *SocialProfileUrls) HasTwitterUrl() bool {
	if o != nil && !IsNil(o.TwitterUrl) {
		return true
	}

	return false
}

// SetTwitterUrl gets a reference to the given string and assigns it to the TwitterUrl field.
func (o *SocialProfileUrls) SetTwitterUrl(v string) {
	o.TwitterUrl = &v
}

// GetLinkedInUrl returns the LinkedInUrl field value if set, zero value otherwise.
func (o *SocialProfileUrls) GetLinkedInUrl() string {
	if o == nil || IsNil(o.LinkedInUrl) {
		var ret string
		return ret
	}
	return *o.LinkedInUrl
}

// GetLinkedInUrlOk returns a tuple with the LinkedInUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SocialProfileUrls) GetLinkedInUrlOk() (*string, bool) {
	if o == nil || IsNil(o.LinkedInUrl) {
		return nil, false
	}
	return o.LinkedInUrl, true
}

// HasLinkedInUrl returns a boolean if a field has been set.
func (o *SocialProfileUrls) HasLinkedInUrl() bool {
	if o != nil && !IsNil(o.LinkedInUrl) {
		return true
	}

	return false
}

// SetLinkedInUrl gets a reference to the given string and assigns it to the LinkedInUrl field.
func (o *SocialProfileUrls) SetLinkedInUrl(v string) {
	o.LinkedInUrl = &v
}

func (o SocialProfileUrls) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SocialProfileUrls) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TwitterUrl) {
		toSerialize["twitterUrl"] = o.TwitterUrl
	}
	if !IsNil(o.LinkedInUrl) {
		toSerialize["linkedInUrl"] = o.LinkedInUrl
	}
	return toSerialize, nil
}

type NullableSocialProfileUrls struct {
	value *SocialProfileUrls
	isSet bool
}

func (v NullableSocialProfileUrls) Get() *SocialProfileUrls {
	return v.value
}

func (v *NullableSocialProfileUrls) Set(val *SocialProfileUrls) {
	v.value = val
	v.isSet = true
}

func (v NullableSocialProfileUrls) IsSet() bool {
	return v.isSet
}

func (v *NullableSocialProfileUrls) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSocialProfileUrls(val *SocialProfileUrls) *NullableSocialProfileUrls {
	return &NullableSocialProfileUrls{value: val, isSet: true}
}

func (v NullableSocialProfileUrls) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSocialProfileUrls) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


