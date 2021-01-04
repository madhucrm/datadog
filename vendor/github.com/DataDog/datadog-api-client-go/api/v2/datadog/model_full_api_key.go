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

// FullAPIKey Datadog API key.
type FullAPIKey struct {
	Attributes *FullAPIKeyAttributes `json:"attributes,omitempty"`
	// ID of the API key.
	Id            *string              `json:"id,omitempty"`
	Relationships *APIKeyRelationships `json:"relationships,omitempty"`
	Type          *APIKeysType         `json:"type,omitempty"`
}

// NewFullAPIKey instantiates a new FullAPIKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFullAPIKey() *FullAPIKey {
	this := FullAPIKey{}
	var type_ APIKeysType = "api_keys"
	this.Type = &type_
	return &this
}

// NewFullAPIKeyWithDefaults instantiates a new FullAPIKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFullAPIKeyWithDefaults() *FullAPIKey {
	this := FullAPIKey{}
	var type_ APIKeysType = "api_keys"
	this.Type = &type_
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *FullAPIKey) GetAttributes() FullAPIKeyAttributes {
	if o == nil || o.Attributes == nil {
		var ret FullAPIKeyAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullAPIKey) GetAttributesOk() (*FullAPIKeyAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *FullAPIKey) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given FullAPIKeyAttributes and assigns it to the Attributes field.
func (o *FullAPIKey) SetAttributes(v FullAPIKeyAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FullAPIKey) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullAPIKey) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FullAPIKey) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *FullAPIKey) SetId(v string) {
	o.Id = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *FullAPIKey) GetRelationships() APIKeyRelationships {
	if o == nil || o.Relationships == nil {
		var ret APIKeyRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullAPIKey) GetRelationshipsOk() (*APIKeyRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *FullAPIKey) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given APIKeyRelationships and assigns it to the Relationships field.
func (o *FullAPIKey) SetRelationships(v APIKeyRelationships) {
	o.Relationships = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *FullAPIKey) GetType() APIKeysType {
	if o == nil || o.Type == nil {
		var ret APIKeysType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullAPIKey) GetTypeOk() (*APIKeysType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *FullAPIKey) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given APIKeysType and assigns it to the Type field.
func (o *FullAPIKey) SetType(v APIKeysType) {
	o.Type = &v
}

func (o FullAPIKey) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableFullAPIKey struct {
	value *FullAPIKey
	isSet bool
}

func (v NullableFullAPIKey) Get() *FullAPIKey {
	return v.value
}

func (v *NullableFullAPIKey) Set(val *FullAPIKey) {
	v.value = val
	v.isSet = true
}

func (v NullableFullAPIKey) IsSet() bool {
	return v.isSet
}

func (v *NullableFullAPIKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFullAPIKey(val *FullAPIKey) *NullableFullAPIKey {
	return &NullableFullAPIKey{value: val, isSet: true}
}

func (v NullableFullAPIKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFullAPIKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
