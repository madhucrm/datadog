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

// LogsAggregateBucket A bucket values
type LogsAggregateBucket struct {
	// The key, value pairs for each group by
	By *map[string]string `json:"by,omitempty"`
	// A map of the metric name -> value for regular compute or list of values for a timeseries
	Computes *map[string]LogsAggregateBucketValue `json:"computes,omitempty"`
}

// NewLogsAggregateBucket instantiates a new LogsAggregateBucket object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogsAggregateBucket() *LogsAggregateBucket {
	this := LogsAggregateBucket{}
	return &this
}

// NewLogsAggregateBucketWithDefaults instantiates a new LogsAggregateBucket object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogsAggregateBucketWithDefaults() *LogsAggregateBucket {
	this := LogsAggregateBucket{}
	return &this
}

// GetBy returns the By field value if set, zero value otherwise.
func (o *LogsAggregateBucket) GetBy() map[string]string {
	if o == nil || o.By == nil {
		var ret map[string]string
		return ret
	}
	return *o.By
}

// GetByOk returns a tuple with the By field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsAggregateBucket) GetByOk() (*map[string]string, bool) {
	if o == nil || o.By == nil {
		return nil, false
	}
	return o.By, true
}

// HasBy returns a boolean if a field has been set.
func (o *LogsAggregateBucket) HasBy() bool {
	if o != nil && o.By != nil {
		return true
	}

	return false
}

// SetBy gets a reference to the given map[string]string and assigns it to the By field.
func (o *LogsAggregateBucket) SetBy(v map[string]string) {
	o.By = &v
}

// GetComputes returns the Computes field value if set, zero value otherwise.
func (o *LogsAggregateBucket) GetComputes() map[string]LogsAggregateBucketValue {
	if o == nil || o.Computes == nil {
		var ret map[string]LogsAggregateBucketValue
		return ret
	}
	return *o.Computes
}

// GetComputesOk returns a tuple with the Computes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsAggregateBucket) GetComputesOk() (*map[string]LogsAggregateBucketValue, bool) {
	if o == nil || o.Computes == nil {
		return nil, false
	}
	return o.Computes, true
}

// HasComputes returns a boolean if a field has been set.
func (o *LogsAggregateBucket) HasComputes() bool {
	if o != nil && o.Computes != nil {
		return true
	}

	return false
}

// SetComputes gets a reference to the given map[string]LogsAggregateBucketValue and assigns it to the Computes field.
func (o *LogsAggregateBucket) SetComputes(v map[string]LogsAggregateBucketValue) {
	o.Computes = &v
}

func (o LogsAggregateBucket) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.By != nil {
		toSerialize["by"] = o.By
	}
	if o.Computes != nil {
		toSerialize["computes"] = o.Computes
	}
	return json.Marshal(toSerialize)
}

type NullableLogsAggregateBucket struct {
	value *LogsAggregateBucket
	isSet bool
}

func (v NullableLogsAggregateBucket) Get() *LogsAggregateBucket {
	return v.value
}

func (v *NullableLogsAggregateBucket) Set(val *LogsAggregateBucket) {
	v.value = val
	v.isSet = true
}

func (v NullableLogsAggregateBucket) IsSet() bool {
	return v.isSet
}

func (v *NullableLogsAggregateBucket) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogsAggregateBucket(val *LogsAggregateBucket) *NullableLogsAggregateBucket {
	return &NullableLogsAggregateBucket{value: val, isSet: true}
}

func (v NullableLogsAggregateBucket) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogsAggregateBucket) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
