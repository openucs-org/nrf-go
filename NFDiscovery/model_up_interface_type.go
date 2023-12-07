/*
NRF NFDiscovery Service

NRF NFDiscovery Service. © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.2.0-alpha.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// UPInterfaceType Types of User-Plane interfaces of the UPF
type UPInterfaceType struct {
	UPInterfaceTypeAnyOf *UPInterfaceTypeAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *UPInterfaceType) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into UPInterfaceTypeAnyOf
	err = json.Unmarshal(data, &dst.UPInterfaceTypeAnyOf);
	if err == nil {
		jsonUPInterfaceTypeAnyOf, _ := json.Marshal(dst.UPInterfaceTypeAnyOf)
		if string(jsonUPInterfaceTypeAnyOf) == "{}" { // empty struct
			dst.UPInterfaceTypeAnyOf = nil
		} else {
			return nil // data stored in dst.UPInterfaceTypeAnyOf, return on the first match
		}
	} else {
		dst.UPInterfaceTypeAnyOf = nil
	}

	// try to unmarshal JSON data into string
	err = json.Unmarshal(data, &dst.string);
	if err == nil {
		jsonstring, _ := json.Marshal(dst.string)
		if string(jsonstring) == "{}" { // empty struct
			dst.string = nil
		} else {
			return nil // data stored in dst.string, return on the first match
		}
	} else {
		dst.string = nil
	}

	return fmt.Errorf("Data failed to match schemas in anyOf(UPInterfaceType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *UPInterfaceType) MarshalJSON() ([]byte, error) {
	if src.UPInterfaceTypeAnyOf != nil {
		return json.Marshal(&src.UPInterfaceTypeAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableUPInterfaceType struct {
	value *UPInterfaceType
	isSet bool
}

func (v NullableUPInterfaceType) Get() *UPInterfaceType {
	return v.value
}

func (v *NullableUPInterfaceType) Set(val *UPInterfaceType) {
	v.value = val
	v.isSet = true
}

func (v NullableUPInterfaceType) IsSet() bool {
	return v.isSet
}

func (v *NullableUPInterfaceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUPInterfaceType(val *UPInterfaceType) *NullableUPInterfaceType {
	return &NullableUPInterfaceType{value: val, isSet: true}
}

func (v NullableUPInterfaceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUPInterfaceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


