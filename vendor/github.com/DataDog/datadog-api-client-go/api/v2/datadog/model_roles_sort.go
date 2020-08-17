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

// RolesSort Sorting options for roles.
type RolesSort string

// List of RolesSort
const (
	ROLESSORT_NAME_ASCENDING         RolesSort = "name"
	ROLESSORT_NAME_DESCENDING        RolesSort = "-name"
	ROLESSORT_MODIFIED_AT_ASCENDING  RolesSort = "modified_at"
	ROLESSORT_MODIFIED_AT_DESCENDING RolesSort = "-modified_at"
	ROLESSORT_USER_COUNT_ASCENDING   RolesSort = "user_count"
	ROLESSORT_USER_COUNT_DESCENDING  RolesSort = "-user_count"
)

func (v *RolesSort) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RolesSort(value)
	for _, existing := range []RolesSort{"name", "-name", "modified_at", "-modified_at", "user_count", "-user_count"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RolesSort", value)
}

// Ptr returns reference to RolesSort value
func (v RolesSort) Ptr() *RolesSort {
	return &v
}

type NullableRolesSort struct {
	value *RolesSort
	isSet bool
}

func (v NullableRolesSort) Get() *RolesSort {
	return v.value
}

func (v *NullableRolesSort) Set(val *RolesSort) {
	v.value = val
	v.isSet = true
}

func (v NullableRolesSort) IsSet() bool {
	return v.isSet
}

func (v *NullableRolesSort) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRolesSort(val *RolesSort) *NullableRolesSort {
	return &NullableRolesSort{value: val, isSet: true}
}

func (v NullableRolesSort) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRolesSort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
