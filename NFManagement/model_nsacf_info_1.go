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

// NsacfInfo1 Information of a NSACF NF Instance
type NsacfInfo1 struct {
	NsacfCapability NsacfCapability `json:"nsacfCapability"`
	TaiList *[]Tai `json:"taiList,omitempty"`
	TaiRangeList *[]TaiRange `json:"taiRangeList,omitempty"`
	OtherServingAreas *[]string `json:"otherServingAreas,omitempty"`
}

// NewNsacfInfo1 instantiates a new NsacfInfo1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNsacfInfo1(nsacfCapability NsacfCapability) *NsacfInfo1 {
	this := NsacfInfo1{}
	this.NsacfCapability = nsacfCapability
	return &this
}

// NewNsacfInfo1WithDefaults instantiates a new NsacfInfo1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNsacfInfo1WithDefaults() *NsacfInfo1 {
	this := NsacfInfo1{}
	return &this
}

// GetNsacfCapability returns the NsacfCapability field value
func (o *NsacfInfo1) GetNsacfCapability() NsacfCapability {
	if o == nil {
		var ret NsacfCapability
		return ret
	}

	return o.NsacfCapability
}

// GetNsacfCapabilityOk returns a tuple with the NsacfCapability field value
// and a boolean to check if the value has been set.
func (o *NsacfInfo1) GetNsacfCapabilityOk() (*NsacfCapability, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NsacfCapability, true
}

// SetNsacfCapability sets field value
func (o *NsacfInfo1) SetNsacfCapability(v NsacfCapability) {
	o.NsacfCapability = v
}

// GetTaiList returns the TaiList field value if set, zero value otherwise.
func (o *NsacfInfo1) GetTaiList() []Tai {
	if o == nil || o.TaiList == nil {
		var ret []Tai
		return ret
	}
	return *o.TaiList
}

// GetTaiListOk returns a tuple with the TaiList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NsacfInfo1) GetTaiListOk() (*[]Tai, bool) {
	if o == nil || o.TaiList == nil {
		return nil, false
	}
	return o.TaiList, true
}

// HasTaiList returns a boolean if a field has been set.
func (o *NsacfInfo1) HasTaiList() bool {
	if o != nil && o.TaiList != nil {
		return true
	}

	return false
}

// SetTaiList gets a reference to the given []Tai and assigns it to the TaiList field.
func (o *NsacfInfo1) SetTaiList(v []Tai) {
	o.TaiList = &v
}

// GetTaiRangeList returns the TaiRangeList field value if set, zero value otherwise.
func (o *NsacfInfo1) GetTaiRangeList() []TaiRange {
	if o == nil || o.TaiRangeList == nil {
		var ret []TaiRange
		return ret
	}
	return *o.TaiRangeList
}

// GetTaiRangeListOk returns a tuple with the TaiRangeList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NsacfInfo1) GetTaiRangeListOk() (*[]TaiRange, bool) {
	if o == nil || o.TaiRangeList == nil {
		return nil, false
	}
	return o.TaiRangeList, true
}

// HasTaiRangeList returns a boolean if a field has been set.
func (o *NsacfInfo1) HasTaiRangeList() bool {
	if o != nil && o.TaiRangeList != nil {
		return true
	}

	return false
}

// SetTaiRangeList gets a reference to the given []TaiRange and assigns it to the TaiRangeList field.
func (o *NsacfInfo1) SetTaiRangeList(v []TaiRange) {
	o.TaiRangeList = &v
}

// GetOtherServingAreas returns the OtherServingAreas field value if set, zero value otherwise.
func (o *NsacfInfo1) GetOtherServingAreas() []string {
	if o == nil || o.OtherServingAreas == nil {
		var ret []string
		return ret
	}
	return *o.OtherServingAreas
}

// GetOtherServingAreasOk returns a tuple with the OtherServingAreas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NsacfInfo1) GetOtherServingAreasOk() (*[]string, bool) {
	if o == nil || o.OtherServingAreas == nil {
		return nil, false
	}
	return o.OtherServingAreas, true
}

// HasOtherServingAreas returns a boolean if a field has been set.
func (o *NsacfInfo1) HasOtherServingAreas() bool {
	if o != nil && o.OtherServingAreas != nil {
		return true
	}

	return false
}

// SetOtherServingAreas gets a reference to the given []string and assigns it to the OtherServingAreas field.
func (o *NsacfInfo1) SetOtherServingAreas(v []string) {
	o.OtherServingAreas = &v
}

func (o NsacfInfo1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["nsacfCapability"] = o.NsacfCapability
	}
	if o.TaiList != nil {
		toSerialize["taiList"] = o.TaiList
	}
	if o.TaiRangeList != nil {
		toSerialize["taiRangeList"] = o.TaiRangeList
	}
	if o.OtherServingAreas != nil {
		toSerialize["otherServingAreas"] = o.OtherServingAreas
	}
	return json.Marshal(toSerialize)
}

type NullableNsacfInfo1 struct {
	value *NsacfInfo1
	isSet bool
}

func (v NullableNsacfInfo1) Get() *NsacfInfo1 {
	return v.value
}

func (v *NullableNsacfInfo1) Set(val *NsacfInfo1) {
	v.value = val
	v.isSet = true
}

func (v NullableNsacfInfo1) IsSet() bool {
	return v.isSet
}

func (v *NullableNsacfInfo1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNsacfInfo1(val *NsacfInfo1) *NullableNsacfInfo1 {
	return &NullableNsacfInfo1{value: val, isSet: true}
}

func (v NullableNsacfInfo1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNsacfInfo1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


