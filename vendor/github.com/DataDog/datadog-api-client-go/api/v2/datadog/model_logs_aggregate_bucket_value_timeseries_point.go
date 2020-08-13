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

// LogsAggregateBucketValueTimeseriesPoint A timeseries point
type LogsAggregateBucketValueTimeseriesPoint struct {
	// The time value for this point
	Time *string `json:"time,omitempty"`
	// The value for this point
	Value *float64 `json:"value,omitempty"`
}

// NewLogsAggregateBucketValueTimeseriesPoint instantiates a new LogsAggregateBucketValueTimeseriesPoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogsAggregateBucketValueTimeseriesPoint() *LogsAggregateBucketValueTimeseriesPoint {
	this := LogsAggregateBucketValueTimeseriesPoint{}
	return &this
}

// NewLogsAggregateBucketValueTimeseriesPointWithDefaults instantiates a new LogsAggregateBucketValueTimeseriesPoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogsAggregateBucketValueTimeseriesPointWithDefaults() *LogsAggregateBucketValueTimeseriesPoint {
	this := LogsAggregateBucketValueTimeseriesPoint{}
	return &this
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *LogsAggregateBucketValueTimeseriesPoint) GetTime() string {
	if o == nil || o.Time == nil {
		var ret string
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsAggregateBucketValueTimeseriesPoint) GetTimeOk() (*string, bool) {
	if o == nil || o.Time == nil {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *LogsAggregateBucketValueTimeseriesPoint) HasTime() bool {
	if o != nil && o.Time != nil {
		return true
	}

	return false
}

// SetTime gets a reference to the given string and assigns it to the Time field.
func (o *LogsAggregateBucketValueTimeseriesPoint) SetTime(v string) {
	o.Time = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *LogsAggregateBucketValueTimeseriesPoint) GetValue() float64 {
	if o == nil || o.Value == nil {
		var ret float64
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsAggregateBucketValueTimeseriesPoint) GetValueOk() (*float64, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *LogsAggregateBucketValueTimeseriesPoint) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given float64 and assigns it to the Value field.
func (o *LogsAggregateBucketValueTimeseriesPoint) SetValue(v float64) {
	o.Value = &v
}

func (o LogsAggregateBucketValueTimeseriesPoint) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Time != nil {
		toSerialize["time"] = o.Time
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableLogsAggregateBucketValueTimeseriesPoint struct {
	value *LogsAggregateBucketValueTimeseriesPoint
	isSet bool
}

func (v NullableLogsAggregateBucketValueTimeseriesPoint) Get() *LogsAggregateBucketValueTimeseriesPoint {
	return v.value
}

func (v *NullableLogsAggregateBucketValueTimeseriesPoint) Set(val *LogsAggregateBucketValueTimeseriesPoint) {
	v.value = val
	v.isSet = true
}

func (v NullableLogsAggregateBucketValueTimeseriesPoint) IsSet() bool {
	return v.isSet
}

func (v *NullableLogsAggregateBucketValueTimeseriesPoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogsAggregateBucketValueTimeseriesPoint(val *LogsAggregateBucketValueTimeseriesPoint) *NullableLogsAggregateBucketValueTimeseriesPoint {
	return &NullableLogsAggregateBucketValueTimeseriesPoint{value: val, isSet: true}
}

func (v NullableLogsAggregateBucketValueTimeseriesPoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogsAggregateBucketValueTimeseriesPoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
