/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
	"fmt"
)

// WidgetMonitorSummaryDisplayFormat What to display on the widget.
type WidgetMonitorSummaryDisplayFormat string

// List of WidgetMonitorSummaryDisplayFormat
const (
	WIDGETMONITORSUMMARYDISPLAYFORMAT_COUNTS          WidgetMonitorSummaryDisplayFormat = "counts"
	WIDGETMONITORSUMMARYDISPLAYFORMAT_COUNTS_AND_LIST WidgetMonitorSummaryDisplayFormat = "countsAndList"
	WIDGETMONITORSUMMARYDISPLAYFORMAT_LIST            WidgetMonitorSummaryDisplayFormat = "list"
)

func (v *WidgetMonitorSummaryDisplayFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := WidgetMonitorSummaryDisplayFormat(value)
	for _, existing := range []WidgetMonitorSummaryDisplayFormat{"counts", "countsAndList", "list"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid WidgetMonitorSummaryDisplayFormat", *v)
}

// Ptr returns reference to WidgetMonitorSummaryDisplayFormat value
func (v WidgetMonitorSummaryDisplayFormat) Ptr() *WidgetMonitorSummaryDisplayFormat {
	return &v
}

type NullableWidgetMonitorSummaryDisplayFormat struct {
	value *WidgetMonitorSummaryDisplayFormat
	isSet bool
}

func (v NullableWidgetMonitorSummaryDisplayFormat) Get() *WidgetMonitorSummaryDisplayFormat {
	return v.value
}

func (v *NullableWidgetMonitorSummaryDisplayFormat) Set(val *WidgetMonitorSummaryDisplayFormat) {
	v.value = val
	v.isSet = true
}

func (v NullableWidgetMonitorSummaryDisplayFormat) IsSet() bool {
	return v.isSet
}

func (v *NullableWidgetMonitorSummaryDisplayFormat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWidgetMonitorSummaryDisplayFormat(val *WidgetMonitorSummaryDisplayFormat) *NullableWidgetMonitorSummaryDisplayFormat {
	return &NullableWidgetMonitorSummaryDisplayFormat{value: val, isSet: true}
}

func (v NullableWidgetMonitorSummaryDisplayFormat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWidgetMonitorSummaryDisplayFormat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
