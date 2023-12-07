/*
NRF NFManagement Service

NRF NFManagement Service. © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.2.0-alpha.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// TransportProtocolAnyOf the model 'TransportProtocolAnyOf'
type TransportProtocolAnyOf string

// List of TransportProtocol_anyOf
const (
	TCP TransportProtocolAnyOf = "TCP"
)

var allowedTransportProtocolAnyOfEnumValues = []TransportProtocolAnyOf{
	"TCP",
}

func (v *TransportProtocolAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TransportProtocolAnyOf(value)
	for _, existing := range allowedTransportProtocolAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TransportProtocolAnyOf", value)
}

// NewTransportProtocolAnyOfFromValue returns a pointer to a valid TransportProtocolAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTransportProtocolAnyOfFromValue(v string) (*TransportProtocolAnyOf, error) {
	ev := TransportProtocolAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TransportProtocolAnyOf: valid values are %v", v, allowedTransportProtocolAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TransportProtocolAnyOf) IsValid() bool {
	for _, existing := range allowedTransportProtocolAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TransportProtocol_anyOf value
func (v TransportProtocolAnyOf) Ptr() *TransportProtocolAnyOf {
	return &v
}

type NullableTransportProtocolAnyOf struct {
	value *TransportProtocolAnyOf
	isSet bool
}

func (v NullableTransportProtocolAnyOf) Get() *TransportProtocolAnyOf {
	return v.value
}

func (v *NullableTransportProtocolAnyOf) Set(val *TransportProtocolAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTransportProtocolAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTransportProtocolAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransportProtocolAnyOf(val *TransportProtocolAnyOf) *NullableTransportProtocolAnyOf {
	return &NullableTransportProtocolAnyOf{value: val, isSet: true}
}

func (v NullableTransportProtocolAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransportProtocolAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

