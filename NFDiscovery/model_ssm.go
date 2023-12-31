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

// Ssm struct for Ssm
type Ssm struct {
	SourceIpAddr IpAddr `json:"sourceIpAddr"`
	DestIpAddr IpAddr `json:"destIpAddr"`
}

// NewSsm instantiates a new Ssm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSsm(sourceIpAddr IpAddr, destIpAddr IpAddr) *Ssm {
	this := Ssm{}
	this.SourceIpAddr = sourceIpAddr
	this.DestIpAddr = destIpAddr
	return &this
}

// NewSsmWithDefaults instantiates a new Ssm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSsmWithDefaults() *Ssm {
	this := Ssm{}
	return &this
}

// GetSourceIpAddr returns the SourceIpAddr field value
func (o *Ssm) GetSourceIpAddr() IpAddr {
	if o == nil {
		var ret IpAddr
		return ret
	}

	return o.SourceIpAddr
}

// GetSourceIpAddrOk returns a tuple with the SourceIpAddr field value
// and a boolean to check if the value has been set.
func (o *Ssm) GetSourceIpAddrOk() (*IpAddr, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SourceIpAddr, true
}

// SetSourceIpAddr sets field value
func (o *Ssm) SetSourceIpAddr(v IpAddr) {
	o.SourceIpAddr = v
}

// GetDestIpAddr returns the DestIpAddr field value
func (o *Ssm) GetDestIpAddr() IpAddr {
	if o == nil {
		var ret IpAddr
		return ret
	}

	return o.DestIpAddr
}

// GetDestIpAddrOk returns a tuple with the DestIpAddr field value
// and a boolean to check if the value has been set.
func (o *Ssm) GetDestIpAddrOk() (*IpAddr, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DestIpAddr, true
}

// SetDestIpAddr sets field value
func (o *Ssm) SetDestIpAddr(v IpAddr) {
	o.DestIpAddr = v
}

func (o Ssm) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["sourceIpAddr"] = o.SourceIpAddr
	}
	if true {
		toSerialize["destIpAddr"] = o.DestIpAddr
	}
	return json.Marshal(toSerialize)
}

type NullableSsm struct {
	value *Ssm
	isSet bool
}

func (v NullableSsm) Get() *Ssm {
	return v.value
}

func (v *NullableSsm) Set(val *Ssm) {
	v.value = val
	v.isSet = true
}

func (v NullableSsm) IsSet() bool {
	return v.isSet
}

func (v *NullableSsm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSsm(val *Ssm) *NullableSsm {
	return &NullableSsm{value: val, isSet: true}
}

func (v NullableSsm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSsm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


