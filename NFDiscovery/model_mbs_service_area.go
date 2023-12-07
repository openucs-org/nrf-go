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

// MbsServiceArea struct for MbsServiceArea
type MbsServiceArea struct {
	// List of NR cell Ids
	NcgiList *[]Ncgi `json:"ncgiList,omitempty"`
	// List of tracking area Ids
	TaiList *[]Tai `json:"taiList,omitempty"`
}

// NewMbsServiceArea instantiates a new MbsServiceArea object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMbsServiceArea() *MbsServiceArea {
	this := MbsServiceArea{}
	return &this
}

// NewMbsServiceAreaWithDefaults instantiates a new MbsServiceArea object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMbsServiceAreaWithDefaults() *MbsServiceArea {
	this := MbsServiceArea{}
	return &this
}

// GetNcgiList returns the NcgiList field value if set, zero value otherwise.
func (o *MbsServiceArea) GetNcgiList() []Ncgi {
	if o == nil || o.NcgiList == nil {
		var ret []Ncgi
		return ret
	}
	return *o.NcgiList
}

// GetNcgiListOk returns a tuple with the NcgiList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbsServiceArea) GetNcgiListOk() (*[]Ncgi, bool) {
	if o == nil || o.NcgiList == nil {
		return nil, false
	}
	return o.NcgiList, true
}

// HasNcgiList returns a boolean if a field has been set.
func (o *MbsServiceArea) HasNcgiList() bool {
	if o != nil && o.NcgiList != nil {
		return true
	}

	return false
}

// SetNcgiList gets a reference to the given []Ncgi and assigns it to the NcgiList field.
func (o *MbsServiceArea) SetNcgiList(v []Ncgi) {
	o.NcgiList = &v
}

// GetTaiList returns the TaiList field value if set, zero value otherwise.
func (o *MbsServiceArea) GetTaiList() []Tai {
	if o == nil || o.TaiList == nil {
		var ret []Tai
		return ret
	}
	return *o.TaiList
}

// GetTaiListOk returns a tuple with the TaiList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbsServiceArea) GetTaiListOk() (*[]Tai, bool) {
	if o == nil || o.TaiList == nil {
		return nil, false
	}
	return o.TaiList, true
}

// HasTaiList returns a boolean if a field has been set.
func (o *MbsServiceArea) HasTaiList() bool {
	if o != nil && o.TaiList != nil {
		return true
	}

	return false
}

// SetTaiList gets a reference to the given []Tai and assigns it to the TaiList field.
func (o *MbsServiceArea) SetTaiList(v []Tai) {
	o.TaiList = &v
}

func (o MbsServiceArea) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NcgiList != nil {
		toSerialize["ncgiList"] = o.NcgiList
	}
	if o.TaiList != nil {
		toSerialize["taiList"] = o.TaiList
	}
	return json.Marshal(toSerialize)
}

type NullableMbsServiceArea struct {
	value *MbsServiceArea
	isSet bool
}

func (v NullableMbsServiceArea) Get() *MbsServiceArea {
	return v.value
}

func (v *NullableMbsServiceArea) Set(val *MbsServiceArea) {
	v.value = val
	v.isSet = true
}

func (v NullableMbsServiceArea) IsSet() bool {
	return v.isSet
}

func (v *NullableMbsServiceArea) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMbsServiceArea(val *MbsServiceArea) *NullableMbsServiceArea {
	return &NullableMbsServiceArea{value: val, isSet: true}
}

func (v NullableMbsServiceArea) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMbsServiceArea) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


