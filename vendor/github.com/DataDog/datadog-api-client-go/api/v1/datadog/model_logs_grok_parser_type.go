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

// LogsGrokParserType Type of logs grok parser.
type LogsGrokParserType string

// List of LogsGrokParserType
const (
	LOGSGROKPARSERTYPE_GROK_PARSER LogsGrokParserType = "grok-parser"
)

func (v *LogsGrokParserType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LogsGrokParserType(value)
	for _, existing := range []LogsGrokParserType{"grok-parser"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LogsGrokParserType", value)
}

// Ptr returns reference to LogsGrokParserType value
func (v LogsGrokParserType) Ptr() *LogsGrokParserType {
	return &v
}

type NullableLogsGrokParserType struct {
	value *LogsGrokParserType
	isSet bool
}

func (v NullableLogsGrokParserType) Get() *LogsGrokParserType {
	return v.value
}

func (v *NullableLogsGrokParserType) Set(val *LogsGrokParserType) {
	v.value = val
	v.isSet = true
}

func (v NullableLogsGrokParserType) IsSet() bool {
	return v.isSet
}

func (v *NullableLogsGrokParserType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogsGrokParserType(val *LogsGrokParserType) *NullableLogsGrokParserType {
	return &NullableLogsGrokParserType{value: val, isSet: true}
}

func (v NullableLogsGrokParserType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogsGrokParserType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
