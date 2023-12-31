/*
NRF NFManagement Service

NRF NFManagement Service. © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.2.0-alpha.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
	"fmt"
)

// NFProfile Information of an NF Instance registered in the NRF
type NFProfile struct {
	interface{} *interface{}
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *NFProfile) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into interface{}
	err = json.Unmarshal(data, &dst.interface{});
	if err == nil {
		jsoninterface{}, _ := json.Marshal(dst.interface{})
		if string(jsoninterface{}) == "{}" { // empty struct
			dst.interface{} = nil
		} else {
			return nil // data stored in dst.interface{}, return on the first match
		}
	} else {
		dst.interface{} = nil
	}

	return fmt.Errorf("Data failed to match schemas in anyOf(NFProfile)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *NFProfile) MarshalJSON() ([]byte, error) {
	if src.interface{} != nil {
		return json.Marshal(&src.interface{})
	}

	return nil, nil // no data in anyOf schemas
}

type NullableNFProfile struct {
	value *NFProfile
	isSet bool
}

func (v NullableNFProfile) Get() *NFProfile {
	return v.value
}

func (v *NullableNFProfile) Set(val *NFProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableNFProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableNFProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNFProfile(val *NFProfile) *NullableNFProfile {
	return &NullableNFProfile{value: val, isSet: true}
}

func (v NullableNFProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNFProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


