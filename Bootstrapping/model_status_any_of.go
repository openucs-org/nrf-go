/*
NRF Bootstrapping

NRF Bootstrapping. © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.1.0-alpha.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// StatusAnyOf the model 'StatusAnyOf'
type StatusAnyOf string

// List of Status_anyOf
const (
	OPERATIVE StatusAnyOf = "OPERATIVE"
	NON_OPERATIVE StatusAnyOf = "NON_OPERATIVE"
)

var allowedStatusAnyOfEnumValues = []StatusAnyOf{
	"OPERATIVE",
	"NON_OPERATIVE",
}

func (v *StatusAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := StatusAnyOf(value)
	for _, existing := range allowedStatusAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid StatusAnyOf", value)
}

// NewStatusAnyOfFromValue returns a pointer to a valid StatusAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStatusAnyOfFromValue(v string) (*StatusAnyOf, error) {
	ev := StatusAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for StatusAnyOf: valid values are %v", v, allowedStatusAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StatusAnyOf) IsValid() bool {
	for _, existing := range allowedStatusAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Status_anyOf value
func (v StatusAnyOf) Ptr() *StatusAnyOf {
	return &v
}

type NullableStatusAnyOf struct {
	value *StatusAnyOf
	isSet bool
}

func (v NullableStatusAnyOf) Get() *StatusAnyOf {
	return v.value
}

func (v *NullableStatusAnyOf) Set(val *StatusAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStatusAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStatusAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatusAnyOf(val *StatusAnyOf) *NullableStatusAnyOf {
	return &NullableStatusAnyOf{value: val, isSet: true}
}

func (v NullableStatusAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatusAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

