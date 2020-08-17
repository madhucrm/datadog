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

// WidgetLayoutType Layout type of the group.
type WidgetLayoutType string

// List of WidgetLayoutType
const (
	WIDGETLAYOUTTYPE_ORDERED WidgetLayoutType = "ordered"
)

func (v *WidgetLayoutType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := WidgetLayoutType(value)
	for _, existing := range []WidgetLayoutType{"ordered"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid WidgetLayoutType", value)
}

// Ptr returns reference to WidgetLayoutType value
func (v WidgetLayoutType) Ptr() *WidgetLayoutType {
	return &v
}

type NullableWidgetLayoutType struct {
	value *WidgetLayoutType
	isSet bool
}

func (v NullableWidgetLayoutType) Get() *WidgetLayoutType {
	return v.value
}

func (v *NullableWidgetLayoutType) Set(val *WidgetLayoutType) {
	v.value = val
	v.isSet = true
}

func (v NullableWidgetLayoutType) IsSet() bool {
	return v.isSet
}

func (v *NullableWidgetLayoutType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWidgetLayoutType(val *WidgetLayoutType) *NullableWidgetLayoutType {
	return &NullableWidgetLayoutType{value: val, isSet: true}
}

func (v NullableWidgetLayoutType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWidgetLayoutType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
