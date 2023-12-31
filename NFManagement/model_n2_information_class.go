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

// N2InformationClass Enumeration for N2 Information Class
type N2InformationClass struct {
	N2InformationClassAnyOf *N2InformationClassAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *N2InformationClass) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into N2InformationClassAnyOf
	err = json.Unmarshal(data, &dst.N2InformationClassAnyOf);
	if err == nil {
		jsonN2InformationClassAnyOf, _ := json.Marshal(dst.N2InformationClassAnyOf)
		if string(jsonN2InformationClassAnyOf) == "{}" { // empty struct
			dst.N2InformationClassAnyOf = nil
		} else {
			return nil // data stored in dst.N2InformationClassAnyOf, return on the first match
		}
	} else {
		dst.N2InformationClassAnyOf = nil
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

	return fmt.Errorf("Data failed to match schemas in anyOf(N2InformationClass)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *N2InformationClass) MarshalJSON() ([]byte, error) {
	if src.N2InformationClassAnyOf != nil {
		return json.Marshal(&src.N2InformationClassAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableN2InformationClass struct {
	value *N2InformationClass
	isSet bool
}

func (v NullableN2InformationClass) Get() *N2InformationClass {
	return v.value
}

func (v *NullableN2InformationClass) Set(val *N2InformationClass) {
	v.value = val
	v.isSet = true
}

func (v NullableN2InformationClass) IsSet() bool {
	return v.isSet
}

func (v *NullableN2InformationClass) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableN2InformationClass(val *N2InformationClass) *NullableN2InformationClass {
	return &NullableN2InformationClass{value: val, isSet: true}
}

func (v NullableN2InformationClass) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableN2InformationClass) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


