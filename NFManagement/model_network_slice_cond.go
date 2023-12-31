/*
NRF NFManagement Service

NRF NFManagement Service. © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.2.0-alpha.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// NetworkSliceCond Subscription to a set of NFs, based on the slices (S-NSSAI and NSI) they support
type NetworkSliceCond struct {
	SnssaiList []Snssai `json:"snssaiList"`
	NsiList *[]string `json:"nsiList,omitempty"`
}

// NewNetworkSliceCond instantiates a new NetworkSliceCond object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkSliceCond(snssaiList []Snssai) *NetworkSliceCond {
	this := NetworkSliceCond{}
	this.SnssaiList = snssaiList
	return &this
}

// NewNetworkSliceCondWithDefaults instantiates a new NetworkSliceCond object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkSliceCondWithDefaults() *NetworkSliceCond {
	this := NetworkSliceCond{}
	return &this
}

// GetSnssaiList returns the SnssaiList field value
func (o *NetworkSliceCond) GetSnssaiList() []Snssai {
	if o == nil {
		var ret []Snssai
		return ret
	}

	return o.SnssaiList
}

// GetSnssaiListOk returns a tuple with the SnssaiList field value
// and a boolean to check if the value has been set.
func (o *NetworkSliceCond) GetSnssaiListOk() (*[]Snssai, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SnssaiList, true
}

// SetSnssaiList sets field value
func (o *NetworkSliceCond) SetSnssaiList(v []Snssai) {
	o.SnssaiList = v
}

// GetNsiList returns the NsiList field value if set, zero value otherwise.
func (o *NetworkSliceCond) GetNsiList() []string {
	if o == nil || o.NsiList == nil {
		var ret []string
		return ret
	}
	return *o.NsiList
}

// GetNsiListOk returns a tuple with the NsiList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkSliceCond) GetNsiListOk() (*[]string, bool) {
	if o == nil || o.NsiList == nil {
		return nil, false
	}
	return o.NsiList, true
}

// HasNsiList returns a boolean if a field has been set.
func (o *NetworkSliceCond) HasNsiList() bool {
	if o != nil && o.NsiList != nil {
		return true
	}

	return false
}

// SetNsiList gets a reference to the given []string and assigns it to the NsiList field.
func (o *NetworkSliceCond) SetNsiList(v []string) {
	o.NsiList = &v
}

func (o NetworkSliceCond) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["snssaiList"] = o.SnssaiList
	}
	if o.NsiList != nil {
		toSerialize["nsiList"] = o.NsiList
	}
	return json.Marshal(toSerialize)
}

type NullableNetworkSliceCond struct {
	value *NetworkSliceCond
	isSet bool
}

func (v NullableNetworkSliceCond) Get() *NetworkSliceCond {
	return v.value
}

func (v *NullableNetworkSliceCond) Set(val *NetworkSliceCond) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkSliceCond) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkSliceCond) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkSliceCond(val *NetworkSliceCond) *NullableNetworkSliceCond {
	return &NullableNetworkSliceCond{value: val, isSet: true}
}

func (v NullableNetworkSliceCond) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkSliceCond) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


