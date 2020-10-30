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

// IncidentTeamUpdateAttributes The incident team's attributes for an update request.
type IncidentTeamUpdateAttributes struct {
	// Name of the incident team.
	Name string `json:"name"`
}

// NewIncidentTeamUpdateAttributes instantiates a new IncidentTeamUpdateAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncidentTeamUpdateAttributes(name string) *IncidentTeamUpdateAttributes {
	this := IncidentTeamUpdateAttributes{}
	this.Name = name
	return &this
}

// NewIncidentTeamUpdateAttributesWithDefaults instantiates a new IncidentTeamUpdateAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncidentTeamUpdateAttributesWithDefaults() *IncidentTeamUpdateAttributes {
	this := IncidentTeamUpdateAttributes{}
	return &this
}

// GetName returns the Name field value
func (o *IncidentTeamUpdateAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *IncidentTeamUpdateAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *IncidentTeamUpdateAttributes) SetName(v string) {
	o.Name = v
}

func (o IncidentTeamUpdateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableIncidentTeamUpdateAttributes struct {
	value *IncidentTeamUpdateAttributes
	isSet bool
}

func (v NullableIncidentTeamUpdateAttributes) Get() *IncidentTeamUpdateAttributes {
	return v.value
}

func (v *NullableIncidentTeamUpdateAttributes) Set(val *IncidentTeamUpdateAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableIncidentTeamUpdateAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableIncidentTeamUpdateAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncidentTeamUpdateAttributes(val *IncidentTeamUpdateAttributes) *NullableIncidentTeamUpdateAttributes {
	return &NullableIncidentTeamUpdateAttributes{value: val, isSet: true}
}

func (v NullableIncidentTeamUpdateAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncidentTeamUpdateAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
