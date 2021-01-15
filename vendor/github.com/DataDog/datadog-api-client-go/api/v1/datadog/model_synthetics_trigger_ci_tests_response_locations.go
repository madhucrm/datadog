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

// SyntheticsTriggerCITestsResponseLocations Synthetics location.
type SyntheticsTriggerCITestsResponseLocations struct {
	// Unique identifier of the location.
	Id *int64 `json:"id,omitempty"`
	// Name of the location.
	Name *string `json:"name,omitempty"`
}

// NewSyntheticsTriggerCITestsResponseLocations instantiates a new SyntheticsTriggerCITestsResponseLocations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyntheticsTriggerCITestsResponseLocations() *SyntheticsTriggerCITestsResponseLocations {
	this := SyntheticsTriggerCITestsResponseLocations{}
	return &this
}

// NewSyntheticsTriggerCITestsResponseLocationsWithDefaults instantiates a new SyntheticsTriggerCITestsResponseLocations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyntheticsTriggerCITestsResponseLocationsWithDefaults() *SyntheticsTriggerCITestsResponseLocations {
	this := SyntheticsTriggerCITestsResponseLocations{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SyntheticsTriggerCITestsResponseLocations) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTriggerCITestsResponseLocations) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SyntheticsTriggerCITestsResponseLocations) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *SyntheticsTriggerCITestsResponseLocations) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SyntheticsTriggerCITestsResponseLocations) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTriggerCITestsResponseLocations) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SyntheticsTriggerCITestsResponseLocations) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SyntheticsTriggerCITestsResponseLocations) SetName(v string) {
	o.Name = &v
}

func (o SyntheticsTriggerCITestsResponseLocations) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableSyntheticsTriggerCITestsResponseLocations struct {
	value *SyntheticsTriggerCITestsResponseLocations
	isSet bool
}

func (v NullableSyntheticsTriggerCITestsResponseLocations) Get() *SyntheticsTriggerCITestsResponseLocations {
	return v.value
}

func (v *NullableSyntheticsTriggerCITestsResponseLocations) Set(val *SyntheticsTriggerCITestsResponseLocations) {
	v.value = val
	v.isSet = true
}

func (v NullableSyntheticsTriggerCITestsResponseLocations) IsSet() bool {
	return v.isSet
}

func (v *NullableSyntheticsTriggerCITestsResponseLocations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyntheticsTriggerCITestsResponseLocations(val *SyntheticsTriggerCITestsResponseLocations) *NullableSyntheticsTriggerCITestsResponseLocations {
	return &NullableSyntheticsTriggerCITestsResponseLocations{value: val, isSet: true}
}

func (v NullableSyntheticsTriggerCITestsResponseLocations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyntheticsTriggerCITestsResponseLocations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
