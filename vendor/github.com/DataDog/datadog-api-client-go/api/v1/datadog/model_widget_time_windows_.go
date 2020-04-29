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

// WidgetTimeWindows Define a time window.
type WidgetTimeWindows string

// List of WidgetTimeWindows
const (
	WIDGETTIMEWINDOWS_SEVEN_DAYS     WidgetTimeWindows = "7d"
	WIDGETTIMEWINDOWS_THIRTY_DAYS    WidgetTimeWindows = "30d"
	WIDGETTIMEWINDOWS_NINETY_DAYS    WidgetTimeWindows = "90d"
	WIDGETTIMEWINDOWS_WEEK_TO_DATE   WidgetTimeWindows = "week_to_date"
	WIDGETTIMEWINDOWS_PREVIOUS_WEEK  WidgetTimeWindows = "previous_week"
	WIDGETTIMEWINDOWS_MONTH_TO_DATE  WidgetTimeWindows = "month_to_date"
	WIDGETTIMEWINDOWS_PREVIOUS_MONTH WidgetTimeWindows = "previous_month"
)

// Ptr returns reference to WidgetTimeWindows value
func (v WidgetTimeWindows) Ptr() *WidgetTimeWindows {
	return &v
}

type NullableWidgetTimeWindows struct {
	value *WidgetTimeWindows
	isSet bool
}

func (v NullableWidgetTimeWindows) Get() *WidgetTimeWindows {
	return v.value
}

func (v *NullableWidgetTimeWindows) Set(val *WidgetTimeWindows) {
	v.value = val
	v.isSet = true
}

func (v NullableWidgetTimeWindows) IsSet() bool {
	return v.isSet
}

func (v *NullableWidgetTimeWindows) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWidgetTimeWindows(val *WidgetTimeWindows) *NullableWidgetTimeWindows {
	return &NullableWidgetTimeWindows{value: val, isSet: true}
}

func (v NullableWidgetTimeWindows) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWidgetTimeWindows) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
