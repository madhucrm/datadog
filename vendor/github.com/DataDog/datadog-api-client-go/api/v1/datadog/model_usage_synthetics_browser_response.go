/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
)

// UsageSyntheticsBrowserResponse Response containing the number of Synthetics Browser tests run for each hour for a given organization.
type UsageSyntheticsBrowserResponse struct {
	// Get hourly usage for Synthetics Browser tests.
	Usage *[]UsageSyntheticsBrowserHour `json:"usage,omitempty"`
}

// NewUsageSyntheticsBrowserResponse instantiates a new UsageSyntheticsBrowserResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsageSyntheticsBrowserResponse() *UsageSyntheticsBrowserResponse {
	this := UsageSyntheticsBrowserResponse{}
	return &this
}

// NewUsageSyntheticsBrowserResponseWithDefaults instantiates a new UsageSyntheticsBrowserResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsageSyntheticsBrowserResponseWithDefaults() *UsageSyntheticsBrowserResponse {
	this := UsageSyntheticsBrowserResponse{}
	return &this
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *UsageSyntheticsBrowserResponse) GetUsage() []UsageSyntheticsBrowserHour {
	if o == nil || o.Usage == nil {
		var ret []UsageSyntheticsBrowserHour
		return ret
	}
	return *o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSyntheticsBrowserResponse) GetUsageOk() (*[]UsageSyntheticsBrowserHour, bool) {
	if o == nil || o.Usage == nil {
		return nil, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *UsageSyntheticsBrowserResponse) HasUsage() bool {
	if o != nil && o.Usage != nil {
		return true
	}

	return false
}

// SetUsage gets a reference to the given []UsageSyntheticsBrowserHour and assigns it to the Usage field.
func (o *UsageSyntheticsBrowserResponse) SetUsage(v []UsageSyntheticsBrowserHour) {
	o.Usage = &v
}

func (o UsageSyntheticsBrowserResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Usage != nil {
		toSerialize["usage"] = o.Usage
	}
	return json.Marshal(toSerialize)
}

type NullableUsageSyntheticsBrowserResponse struct {
	value *UsageSyntheticsBrowserResponse
	isSet bool
}

func (v NullableUsageSyntheticsBrowserResponse) Get() *UsageSyntheticsBrowserResponse {
	return v.value
}

func (v *NullableUsageSyntheticsBrowserResponse) Set(val *UsageSyntheticsBrowserResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUsageSyntheticsBrowserResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUsageSyntheticsBrowserResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsageSyntheticsBrowserResponse(val *UsageSyntheticsBrowserResponse) *NullableUsageSyntheticsBrowserResponse {
	return &NullableUsageSyntheticsBrowserResponse{value: val, isSet: true}
}

func (v NullableUsageSyntheticsBrowserResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsageSyntheticsBrowserResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
