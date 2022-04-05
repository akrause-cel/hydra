/*
 * Ory Hydra API
 *
 * Documentation for all of Ory Hydra's APIs.
 *
 * API version: 1.0.0
 * Contact: hi@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse503 struct for InlineResponse503
type InlineResponse503 struct {
	// Errors contains a list of errors that caused the not ready status.
	Errors map[string]string `json:"errors"`
}

// NewInlineResponse503 instantiates a new InlineResponse503 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse503(errors map[string]string) *InlineResponse503 {
	this := InlineResponse503{}
	this.Errors = errors
	return &this
}

// NewInlineResponse503WithDefaults instantiates a new InlineResponse503 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse503WithDefaults() *InlineResponse503 {
	this := InlineResponse503{}
	return &this
}

// GetErrors returns the Errors field value
func (o *InlineResponse503) GetErrors() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value
// and a boolean to check if the value has been set.
func (o *InlineResponse503) GetErrorsOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Errors, true
}

// SetErrors sets field value
func (o *InlineResponse503) SetErrors(v map[string]string) {
	o.Errors = v
}

func (o InlineResponse503) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["errors"] = o.Errors
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse503 struct {
	value *InlineResponse503
	isSet bool
}

func (v NullableInlineResponse503) Get() *InlineResponse503 {
	return v.value
}

func (v *NullableInlineResponse503) Set(val *InlineResponse503) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse503) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse503) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse503(val *InlineResponse503) *NullableInlineResponse503 {
	return &NullableInlineResponse503{value: val, isSet: true}
}

func (v NullableInlineResponse503) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse503) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}