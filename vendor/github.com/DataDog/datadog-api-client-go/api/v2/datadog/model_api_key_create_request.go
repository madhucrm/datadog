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

// APIKeyCreateRequest Request used to create an API key.
type APIKeyCreateRequest struct {
	Data APIKeyCreateData `json:"data"`
}

// NewAPIKeyCreateRequest instantiates a new APIKeyCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAPIKeyCreateRequest(data APIKeyCreateData) *APIKeyCreateRequest {
	this := APIKeyCreateRequest{}
	this.Data = data
	return &this
}

// NewAPIKeyCreateRequestWithDefaults instantiates a new APIKeyCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAPIKeyCreateRequestWithDefaults() *APIKeyCreateRequest {
	this := APIKeyCreateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *APIKeyCreateRequest) GetData() APIKeyCreateData {
	if o == nil {
		var ret APIKeyCreateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *APIKeyCreateRequest) GetDataOk() (*APIKeyCreateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *APIKeyCreateRequest) SetData(v APIKeyCreateData) {
	o.Data = v
}

func (o APIKeyCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableAPIKeyCreateRequest struct {
	value *APIKeyCreateRequest
	isSet bool
}

func (v NullableAPIKeyCreateRequest) Get() *APIKeyCreateRequest {
	return v.value
}

func (v *NullableAPIKeyCreateRequest) Set(val *APIKeyCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAPIKeyCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAPIKeyCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAPIKeyCreateRequest(val *APIKeyCreateRequest) *NullableAPIKeyCreateRequest {
	return &NullableAPIKeyCreateRequest{value: val, isSet: true}
}

func (v NullableAPIKeyCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAPIKeyCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
