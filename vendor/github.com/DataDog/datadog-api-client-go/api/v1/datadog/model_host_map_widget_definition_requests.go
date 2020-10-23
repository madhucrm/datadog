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

// HostMapWidgetDefinitionRequests List of definitions.
type HostMapWidgetDefinitionRequests struct {
	Fill *HostMapRequest `json:"fill,omitempty"`
	Size *HostMapRequest `json:"size,omitempty"`
}

// NewHostMapWidgetDefinitionRequests instantiates a new HostMapWidgetDefinitionRequests object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHostMapWidgetDefinitionRequests() *HostMapWidgetDefinitionRequests {
	this := HostMapWidgetDefinitionRequests{}
	return &this
}

// NewHostMapWidgetDefinitionRequestsWithDefaults instantiates a new HostMapWidgetDefinitionRequests object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHostMapWidgetDefinitionRequestsWithDefaults() *HostMapWidgetDefinitionRequests {
	this := HostMapWidgetDefinitionRequests{}
	return &this
}

// GetFill returns the Fill field value if set, zero value otherwise.
func (o *HostMapWidgetDefinitionRequests) GetFill() HostMapRequest {
	if o == nil || o.Fill == nil {
		var ret HostMapRequest
		return ret
	}
	return *o.Fill
}

// GetFillOk returns a tuple with the Fill field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostMapWidgetDefinitionRequests) GetFillOk() (*HostMapRequest, bool) {
	if o == nil || o.Fill == nil {
		return nil, false
	}
	return o.Fill, true
}

// HasFill returns a boolean if a field has been set.
func (o *HostMapWidgetDefinitionRequests) HasFill() bool {
	if o != nil && o.Fill != nil {
		return true
	}

	return false
}

// SetFill gets a reference to the given HostMapRequest and assigns it to the Fill field.
func (o *HostMapWidgetDefinitionRequests) SetFill(v HostMapRequest) {
	o.Fill = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *HostMapWidgetDefinitionRequests) GetSize() HostMapRequest {
	if o == nil || o.Size == nil {
		var ret HostMapRequest
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostMapWidgetDefinitionRequests) GetSizeOk() (*HostMapRequest, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *HostMapWidgetDefinitionRequests) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given HostMapRequest and assigns it to the Size field.
func (o *HostMapWidgetDefinitionRequests) SetSize(v HostMapRequest) {
	o.Size = &v
}

func (o HostMapWidgetDefinitionRequests) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Fill != nil {
		toSerialize["fill"] = o.Fill
	}
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}
	return json.Marshal(toSerialize)
}

type NullableHostMapWidgetDefinitionRequests struct {
	value *HostMapWidgetDefinitionRequests
	isSet bool
}

func (v NullableHostMapWidgetDefinitionRequests) Get() *HostMapWidgetDefinitionRequests {
	return v.value
}

func (v *NullableHostMapWidgetDefinitionRequests) Set(val *HostMapWidgetDefinitionRequests) {
	v.value = val
	v.isSet = true
}

func (v NullableHostMapWidgetDefinitionRequests) IsSet() bool {
	return v.isSet
}

func (v *NullableHostMapWidgetDefinitionRequests) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHostMapWidgetDefinitionRequests(val *HostMapWidgetDefinitionRequests) *NullableHostMapWidgetDefinitionRequests {
	return &NullableHostMapWidgetDefinitionRequests{value: val, isSet: true}
}

func (v NullableHostMapWidgetDefinitionRequests) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHostMapWidgetDefinitionRequests) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
