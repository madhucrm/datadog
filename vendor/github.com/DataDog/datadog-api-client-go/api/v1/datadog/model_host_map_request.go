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

// HostMapRequest Updated host map.
type HostMapRequest struct {
	ApmQuery     *LogQueryDefinition     `json:"apm_query,omitempty"`
	EventQuery   *EventQueryDefinition   `json:"event_query,omitempty"`
	LogQuery     *LogQueryDefinition     `json:"log_query,omitempty"`
	NetworkQuery *LogQueryDefinition     `json:"network_query,omitempty"`
	ProcessQuery *ProcessQueryDefinition `json:"process_query,omitempty"`
	// Query definition.
	Q        *string             `json:"q,omitempty"`
	RumQuery *LogQueryDefinition `json:"rum_query,omitempty"`
}

// NewHostMapRequest instantiates a new HostMapRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHostMapRequest() *HostMapRequest {
	this := HostMapRequest{}
	return &this
}

// NewHostMapRequestWithDefaults instantiates a new HostMapRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHostMapRequestWithDefaults() *HostMapRequest {
	this := HostMapRequest{}
	return &this
}

// GetApmQuery returns the ApmQuery field value if set, zero value otherwise.
func (o *HostMapRequest) GetApmQuery() LogQueryDefinition {
	if o == nil || o.ApmQuery == nil {
		var ret LogQueryDefinition
		return ret
	}
	return *o.ApmQuery
}

// GetApmQueryOk returns a tuple with the ApmQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostMapRequest) GetApmQueryOk() (*LogQueryDefinition, bool) {
	if o == nil || o.ApmQuery == nil {
		return nil, false
	}
	return o.ApmQuery, true
}

// HasApmQuery returns a boolean if a field has been set.
func (o *HostMapRequest) HasApmQuery() bool {
	if o != nil && o.ApmQuery != nil {
		return true
	}

	return false
}

// SetApmQuery gets a reference to the given LogQueryDefinition and assigns it to the ApmQuery field.
func (o *HostMapRequest) SetApmQuery(v LogQueryDefinition) {
	o.ApmQuery = &v
}

// GetEventQuery returns the EventQuery field value if set, zero value otherwise.
func (o *HostMapRequest) GetEventQuery() EventQueryDefinition {
	if o == nil || o.EventQuery == nil {
		var ret EventQueryDefinition
		return ret
	}
	return *o.EventQuery
}

// GetEventQueryOk returns a tuple with the EventQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostMapRequest) GetEventQueryOk() (*EventQueryDefinition, bool) {
	if o == nil || o.EventQuery == nil {
		return nil, false
	}
	return o.EventQuery, true
}

// HasEventQuery returns a boolean if a field has been set.
func (o *HostMapRequest) HasEventQuery() bool {
	if o != nil && o.EventQuery != nil {
		return true
	}

	return false
}

// SetEventQuery gets a reference to the given EventQueryDefinition and assigns it to the EventQuery field.
func (o *HostMapRequest) SetEventQuery(v EventQueryDefinition) {
	o.EventQuery = &v
}

// GetLogQuery returns the LogQuery field value if set, zero value otherwise.
func (o *HostMapRequest) GetLogQuery() LogQueryDefinition {
	if o == nil || o.LogQuery == nil {
		var ret LogQueryDefinition
		return ret
	}
	return *o.LogQuery
}

// GetLogQueryOk returns a tuple with the LogQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostMapRequest) GetLogQueryOk() (*LogQueryDefinition, bool) {
	if o == nil || o.LogQuery == nil {
		return nil, false
	}
	return o.LogQuery, true
}

// HasLogQuery returns a boolean if a field has been set.
func (o *HostMapRequest) HasLogQuery() bool {
	if o != nil && o.LogQuery != nil {
		return true
	}

	return false
}

// SetLogQuery gets a reference to the given LogQueryDefinition and assigns it to the LogQuery field.
func (o *HostMapRequest) SetLogQuery(v LogQueryDefinition) {
	o.LogQuery = &v
}

// GetNetworkQuery returns the NetworkQuery field value if set, zero value otherwise.
func (o *HostMapRequest) GetNetworkQuery() LogQueryDefinition {
	if o == nil || o.NetworkQuery == nil {
		var ret LogQueryDefinition
		return ret
	}
	return *o.NetworkQuery
}

// GetNetworkQueryOk returns a tuple with the NetworkQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostMapRequest) GetNetworkQueryOk() (*LogQueryDefinition, bool) {
	if o == nil || o.NetworkQuery == nil {
		return nil, false
	}
	return o.NetworkQuery, true
}

// HasNetworkQuery returns a boolean if a field has been set.
func (o *HostMapRequest) HasNetworkQuery() bool {
	if o != nil && o.NetworkQuery != nil {
		return true
	}

	return false
}

// SetNetworkQuery gets a reference to the given LogQueryDefinition and assigns it to the NetworkQuery field.
func (o *HostMapRequest) SetNetworkQuery(v LogQueryDefinition) {
	o.NetworkQuery = &v
}

// GetProcessQuery returns the ProcessQuery field value if set, zero value otherwise.
func (o *HostMapRequest) GetProcessQuery() ProcessQueryDefinition {
	if o == nil || o.ProcessQuery == nil {
		var ret ProcessQueryDefinition
		return ret
	}
	return *o.ProcessQuery
}

// GetProcessQueryOk returns a tuple with the ProcessQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostMapRequest) GetProcessQueryOk() (*ProcessQueryDefinition, bool) {
	if o == nil || o.ProcessQuery == nil {
		return nil, false
	}
	return o.ProcessQuery, true
}

// HasProcessQuery returns a boolean if a field has been set.
func (o *HostMapRequest) HasProcessQuery() bool {
	if o != nil && o.ProcessQuery != nil {
		return true
	}

	return false
}

// SetProcessQuery gets a reference to the given ProcessQueryDefinition and assigns it to the ProcessQuery field.
func (o *HostMapRequest) SetProcessQuery(v ProcessQueryDefinition) {
	o.ProcessQuery = &v
}

// GetQ returns the Q field value if set, zero value otherwise.
func (o *HostMapRequest) GetQ() string {
	if o == nil || o.Q == nil {
		var ret string
		return ret
	}
	return *o.Q
}

// GetQOk returns a tuple with the Q field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostMapRequest) GetQOk() (*string, bool) {
	if o == nil || o.Q == nil {
		return nil, false
	}
	return o.Q, true
}

// HasQ returns a boolean if a field has been set.
func (o *HostMapRequest) HasQ() bool {
	if o != nil && o.Q != nil {
		return true
	}

	return false
}

// SetQ gets a reference to the given string and assigns it to the Q field.
func (o *HostMapRequest) SetQ(v string) {
	o.Q = &v
}

// GetRumQuery returns the RumQuery field value if set, zero value otherwise.
func (o *HostMapRequest) GetRumQuery() LogQueryDefinition {
	if o == nil || o.RumQuery == nil {
		var ret LogQueryDefinition
		return ret
	}
	return *o.RumQuery
}

// GetRumQueryOk returns a tuple with the RumQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostMapRequest) GetRumQueryOk() (*LogQueryDefinition, bool) {
	if o == nil || o.RumQuery == nil {
		return nil, false
	}
	return o.RumQuery, true
}

// HasRumQuery returns a boolean if a field has been set.
func (o *HostMapRequest) HasRumQuery() bool {
	if o != nil && o.RumQuery != nil {
		return true
	}

	return false
}

// SetRumQuery gets a reference to the given LogQueryDefinition and assigns it to the RumQuery field.
func (o *HostMapRequest) SetRumQuery(v LogQueryDefinition) {
	o.RumQuery = &v
}

func (o HostMapRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApmQuery != nil {
		toSerialize["apm_query"] = o.ApmQuery
	}
	if o.EventQuery != nil {
		toSerialize["event_query"] = o.EventQuery
	}
	if o.LogQuery != nil {
		toSerialize["log_query"] = o.LogQuery
	}
	if o.NetworkQuery != nil {
		toSerialize["network_query"] = o.NetworkQuery
	}
	if o.ProcessQuery != nil {
		toSerialize["process_query"] = o.ProcessQuery
	}
	if o.Q != nil {
		toSerialize["q"] = o.Q
	}
	if o.RumQuery != nil {
		toSerialize["rum_query"] = o.RumQuery
	}
	return json.Marshal(toSerialize)
}

type NullableHostMapRequest struct {
	value *HostMapRequest
	isSet bool
}

func (v NullableHostMapRequest) Get() *HostMapRequest {
	return v.value
}

func (v *NullableHostMapRequest) Set(val *HostMapRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableHostMapRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableHostMapRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHostMapRequest(val *HostMapRequest) *NullableHostMapRequest {
	return &NullableHostMapRequest{value: val, isSet: true}
}

func (v NullableHostMapRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHostMapRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
