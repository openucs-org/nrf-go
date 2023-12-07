/*
NRF NFDiscovery Service

NRF NFDiscovery Service. © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.2.0-alpha.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// GmlcInfo Information of a GMLC NF Instance
type GmlcInfo struct {
	ServingClientTypes *[]ExternalClientType `json:"servingClientTypes,omitempty"`
}

// NewGmlcInfo instantiates a new GmlcInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGmlcInfo() *GmlcInfo {
	this := GmlcInfo{}
	return &this
}

// NewGmlcInfoWithDefaults instantiates a new GmlcInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGmlcInfoWithDefaults() *GmlcInfo {
	this := GmlcInfo{}
	return &this
}

// GetServingClientTypes returns the ServingClientTypes field value if set, zero value otherwise.
func (o *GmlcInfo) GetServingClientTypes() []ExternalClientType {
	if o == nil || o.ServingClientTypes == nil {
		var ret []ExternalClientType
		return ret
	}
	return *o.ServingClientTypes
}

// GetServingClientTypesOk returns a tuple with the ServingClientTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GmlcInfo) GetServingClientTypesOk() (*[]ExternalClientType, bool) {
	if o == nil || o.ServingClientTypes == nil {
		return nil, false
	}
	return o.ServingClientTypes, true
}

// HasServingClientTypes returns a boolean if a field has been set.
func (o *GmlcInfo) HasServingClientTypes() bool {
	if o != nil && o.ServingClientTypes != nil {
		return true
	}

	return false
}

// SetServingClientTypes gets a reference to the given []ExternalClientType and assigns it to the ServingClientTypes field.
func (o *GmlcInfo) SetServingClientTypes(v []ExternalClientType) {
	o.ServingClientTypes = &v
}

func (o GmlcInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ServingClientTypes != nil {
		toSerialize["servingClientTypes"] = o.ServingClientTypes
	}
	return json.Marshal(toSerialize)
}

type NullableGmlcInfo struct {
	value *GmlcInfo
	isSet bool
}

func (v NullableGmlcInfo) Get() *GmlcInfo {
	return v.value
}

func (v *NullableGmlcInfo) Set(val *GmlcInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableGmlcInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableGmlcInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGmlcInfo(val *GmlcInfo) *NullableGmlcInfo {
	return &NullableGmlcInfo{value: val, isSet: true}
}

func (v NullableGmlcInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGmlcInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


