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

// checks if the Phone type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Phone{}

// Phone struct for Phone
type Phone struct {
	FirstSeenMs *int64 `json:"firstSeenMs,omitempty"`
	LastSeenMs *int64 `json:"lastSeenMs,omitempty"`
	Observations *int32 `json:"observations,omitempty"`
	Confidence *float64 `json:"confidence,omitempty"`
	Label *string `json:"label,omitempty"`
	Value *string `json:"value,omitempty"`
}

// NewPhone instantiates a new Phone object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPhone() *Phone {
	this := Phone{}
	return &this
}

// NewPhoneWithDefaults instantiates a new Phone object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPhoneWithDefaults() *Phone {
	this := Phone{}
	return &this
}

// GetFirstSeenMs returns the FirstSeenMs field value if set, zero value otherwise.
func (o *Phone) GetFirstSeenMs() int64 {
	if o == nil || IsNil(o.FirstSeenMs) {
		var ret int64
		return ret
	}
	return *o.FirstSeenMs
}

// GetFirstSeenMsOk returns a tuple with the FirstSeenMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Phone) GetFirstSeenMsOk() (*int64, bool) {
	if o == nil || IsNil(o.FirstSeenMs) {
		return nil, false
	}
	return o.FirstSeenMs, true
}

// HasFirstSeenMs returns a boolean if a field has been set.
func (o *Phone) HasFirstSeenMs() bool {
	if o != nil && !IsNil(o.FirstSeenMs) {
		return true
	}

	return false
}

// SetFirstSeenMs gets a reference to the given int64 and assigns it to the FirstSeenMs field.
func (o *Phone) SetFirstSeenMs(v int64) {
	o.FirstSeenMs = &v
}

// GetLastSeenMs returns the LastSeenMs field value if set, zero value otherwise.
func (o *Phone) GetLastSeenMs() int64 {
	if o == nil || IsNil(o.LastSeenMs) {
		var ret int64
		return ret
	}
	return *o.LastSeenMs
}

// GetLastSeenMsOk returns a tuple with the LastSeenMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Phone) GetLastSeenMsOk() (*int64, bool) {
	if o == nil || IsNil(o.LastSeenMs) {
		return nil, false
	}
	return o.LastSeenMs, true
}

// HasLastSeenMs returns a boolean if a field has been set.
func (o *Phone) HasLastSeenMs() bool {
	if o != nil && !IsNil(o.LastSeenMs) {
		return true
	}

	return false
}

// SetLastSeenMs gets a reference to the given int64 and assigns it to the LastSeenMs field.
func (o *Phone) SetLastSeenMs(v int64) {
	o.LastSeenMs = &v
}

// GetObservations returns the Observations field value if set, zero value otherwise.
func (o *Phone) GetObservations() int32 {
	if o == nil || IsNil(o.Observations) {
		var ret int32
		return ret
	}
	return *o.Observations
}

// GetObservationsOk returns a tuple with the Observations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Phone) GetObservationsOk() (*int32, bool) {
	if o == nil || IsNil(o.Observations) {
		return nil, false
	}
	return o.Observations, true
}

// HasObservations returns a boolean if a field has been set.
func (o *Phone) HasObservations() bool {
	if o != nil && !IsNil(o.Observations) {
		return true
	}

	return false
}

// SetObservations gets a reference to the given int32 and assigns it to the Observations field.
func (o *Phone) SetObservations(v int32) {
	o.Observations = &v
}

// GetConfidence returns the Confidence field value if set, zero value otherwise.
func (o *Phone) GetConfidence() float64 {
	if o == nil || IsNil(o.Confidence) {
		var ret float64
		return ret
	}
	return *o.Confidence
}

// GetConfidenceOk returns a tuple with the Confidence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Phone) GetConfidenceOk() (*float64, bool) {
	if o == nil || IsNil(o.Confidence) {
		return nil, false
	}
	return o.Confidence, true
}

// HasConfidence returns a boolean if a field has been set.
func (o *Phone) HasConfidence() bool {
	if o != nil && !IsNil(o.Confidence) {
		return true
	}

	return false
}

// SetConfidence gets a reference to the given float64 and assigns it to the Confidence field.
func (o *Phone) SetConfidence(v float64) {
	o.Confidence = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *Phone) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Phone) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *Phone) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *Phone) SetLabel(v string) {
	o.Label = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *Phone) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Phone) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *Phone) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *Phone) SetValue(v string) {
	o.Value = &v
}

func (o Phone) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Phone) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FirstSeenMs) {
		toSerialize["firstSeenMs"] = o.FirstSeenMs
	}
	if !IsNil(o.LastSeenMs) {
		toSerialize["lastSeenMs"] = o.LastSeenMs
	}
	if !IsNil(o.Observations) {
		toSerialize["observations"] = o.Observations
	}
	if !IsNil(o.Confidence) {
		toSerialize["confidence"] = o.Confidence
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullablePhone struct {
	value *Phone
	isSet bool
}

func (v NullablePhone) Get() *Phone {
	return v.value
}

func (v *NullablePhone) Set(val *Phone) {
	v.value = val
	v.isSet = true
}

func (v NullablePhone) IsSet() bool {
	return v.isSet
}

func (v *NullablePhone) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePhone(val *Phone) *NullablePhone {
	return &NullablePhone{value: val, isSet: true}
}

func (v NullablePhone) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePhone) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


