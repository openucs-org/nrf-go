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

// MbsSessionId struct for MbsSessionId
type MbsSessionId struct {
	Tmgi *Tmgi `json:"tmgi,omitempty"`
	Ssm *Ssm `json:"ssm,omitempty"`
	// This represents the Network Identifier, which together with a PLMN ID is used to identify an SNPN (see 3GPP TS 23.003 and 3GPP TS 23.501 clause 5.30.2.1).
	Nid *string `json:"nid,omitempty"`
}

// NewMbsSessionId instantiates a new MbsSessionId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMbsSessionId() *MbsSessionId {
	this := MbsSessionId{}
	return &this
}

// NewMbsSessionIdWithDefaults instantiates a new MbsSessionId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMbsSessionIdWithDefaults() *MbsSessionId {
	this := MbsSessionId{}
	return &this
}

// GetTmgi returns the Tmgi field value if set, zero value otherwise.
func (o *MbsSessionId) GetTmgi() Tmgi {
	if o == nil || o.Tmgi == nil {
		var ret Tmgi
		return ret
	}
	return *o.Tmgi
}

// GetTmgiOk returns a tuple with the Tmgi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbsSessionId) GetTmgiOk() (*Tmgi, bool) {
	if o == nil || o.Tmgi == nil {
		return nil, false
	}
	return o.Tmgi, true
}

// HasTmgi returns a boolean if a field has been set.
func (o *MbsSessionId) HasTmgi() bool {
	if o != nil && o.Tmgi != nil {
		return true
	}

	return false
}

// SetTmgi gets a reference to the given Tmgi and assigns it to the Tmgi field.
func (o *MbsSessionId) SetTmgi(v Tmgi) {
	o.Tmgi = &v
}

// GetSsm returns the Ssm field value if set, zero value otherwise.
func (o *MbsSessionId) GetSsm() Ssm {
	if o == nil || o.Ssm == nil {
		var ret Ssm
		return ret
	}
	return *o.Ssm
}

// GetSsmOk returns a tuple with the Ssm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbsSessionId) GetSsmOk() (*Ssm, bool) {
	if o == nil || o.Ssm == nil {
		return nil, false
	}
	return o.Ssm, true
}

// HasSsm returns a boolean if a field has been set.
func (o *MbsSessionId) HasSsm() bool {
	if o != nil && o.Ssm != nil {
		return true
	}

	return false
}

// SetSsm gets a reference to the given Ssm and assigns it to the Ssm field.
func (o *MbsSessionId) SetSsm(v Ssm) {
	o.Ssm = &v
}

// GetNid returns the Nid field value if set, zero value otherwise.
func (o *MbsSessionId) GetNid() string {
	if o == nil || o.Nid == nil {
		var ret string
		return ret
	}
	return *o.Nid
}

// GetNidOk returns a tuple with the Nid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbsSessionId) GetNidOk() (*string, bool) {
	if o == nil || o.Nid == nil {
		return nil, false
	}
	return o.Nid, true
}

// HasNid returns a boolean if a field has been set.
func (o *MbsSessionId) HasNid() bool {
	if o != nil && o.Nid != nil {
		return true
	}

	return false
}

// SetNid gets a reference to the given string and assigns it to the Nid field.
func (o *MbsSessionId) SetNid(v string) {
	o.Nid = &v
}

func (o MbsSessionId) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Tmgi != nil {
		toSerialize["tmgi"] = o.Tmgi
	}
	if o.Ssm != nil {
		toSerialize["ssm"] = o.Ssm
	}
	if o.Nid != nil {
		toSerialize["nid"] = o.Nid
	}
	return json.Marshal(toSerialize)
}

type NullableMbsSessionId struct {
	value *MbsSessionId
	isSet bool
}

func (v NullableMbsSessionId) Get() *MbsSessionId {
	return v.value
}

func (v *NullableMbsSessionId) Set(val *MbsSessionId) {
	v.value = val
	v.isSet = true
}

func (v NullableMbsSessionId) IsSet() bool {
	return v.isSet
}

func (v *NullableMbsSessionId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMbsSessionId(val *MbsSessionId) *NullableMbsSessionId {
	return &NullableMbsSessionId{value: val, isSet: true}
}

func (v NullableMbsSessionId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMbsSessionId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


