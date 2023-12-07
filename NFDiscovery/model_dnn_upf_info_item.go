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

// DnnUpfInfoItem Set of parameters supported by UPF for a given DNN
type DnnUpfInfoItem struct {
	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003; it shall contain either a DNN Network Identifier, or a full DNN with both the Network Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots (e.g. \"Label1.Label2.Label3\").
	Dnn string `json:"dnn"`
	DnaiList *[]string `json:"dnaiList,omitempty"`
	PduSessionTypes *[]PduSessionType `json:"pduSessionTypes,omitempty"`
	Ipv4AddressRanges *[]Ipv4AddressRange `json:"ipv4AddressRanges,omitempty"`
	Ipv6PrefixRanges *[]Ipv6PrefixRange `json:"ipv6PrefixRanges,omitempty"`
	// Map of network instance per DNAI for the DNN, where the key of the map is the DNAI. When present, the value of each entry of the map shall contain a N6 network instance that is configured for the DNAI indicated by the key.
	DnaiNwInstanceList *map[string]string `json:"dnaiNwInstanceList,omitempty"`
}

// NewDnnUpfInfoItem instantiates a new DnnUpfInfoItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnnUpfInfoItem(dnn string) *DnnUpfInfoItem {
	this := DnnUpfInfoItem{}
	this.Dnn = dnn
	return &this
}

// NewDnnUpfInfoItemWithDefaults instantiates a new DnnUpfInfoItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnnUpfInfoItemWithDefaults() *DnnUpfInfoItem {
	this := DnnUpfInfoItem{}
	return &this
}

// GetDnn returns the Dnn field value
func (o *DnnUpfInfoItem) GetDnn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Dnn
}

// GetDnnOk returns a tuple with the Dnn field value
// and a boolean to check if the value has been set.
func (o *DnnUpfInfoItem) GetDnnOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Dnn, true
}

// SetDnn sets field value
func (o *DnnUpfInfoItem) SetDnn(v string) {
	o.Dnn = v
}

// GetDnaiList returns the DnaiList field value if set, zero value otherwise.
func (o *DnnUpfInfoItem) GetDnaiList() []string {
	if o == nil || o.DnaiList == nil {
		var ret []string
		return ret
	}
	return *o.DnaiList
}

// GetDnaiListOk returns a tuple with the DnaiList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnnUpfInfoItem) GetDnaiListOk() (*[]string, bool) {
	if o == nil || o.DnaiList == nil {
		return nil, false
	}
	return o.DnaiList, true
}

// HasDnaiList returns a boolean if a field has been set.
func (o *DnnUpfInfoItem) HasDnaiList() bool {
	if o != nil && o.DnaiList != nil {
		return true
	}

	return false
}

// SetDnaiList gets a reference to the given []string and assigns it to the DnaiList field.
func (o *DnnUpfInfoItem) SetDnaiList(v []string) {
	o.DnaiList = &v
}

// GetPduSessionTypes returns the PduSessionTypes field value if set, zero value otherwise.
func (o *DnnUpfInfoItem) GetPduSessionTypes() []PduSessionType {
	if o == nil || o.PduSessionTypes == nil {
		var ret []PduSessionType
		return ret
	}
	return *o.PduSessionTypes
}

// GetPduSessionTypesOk returns a tuple with the PduSessionTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnnUpfInfoItem) GetPduSessionTypesOk() (*[]PduSessionType, bool) {
	if o == nil || o.PduSessionTypes == nil {
		return nil, false
	}
	return o.PduSessionTypes, true
}

// HasPduSessionTypes returns a boolean if a field has been set.
func (o *DnnUpfInfoItem) HasPduSessionTypes() bool {
	if o != nil && o.PduSessionTypes != nil {
		return true
	}

	return false
}

// SetPduSessionTypes gets a reference to the given []PduSessionType and assigns it to the PduSessionTypes field.
func (o *DnnUpfInfoItem) SetPduSessionTypes(v []PduSessionType) {
	o.PduSessionTypes = &v
}

// GetIpv4AddressRanges returns the Ipv4AddressRanges field value if set, zero value otherwise.
func (o *DnnUpfInfoItem) GetIpv4AddressRanges() []Ipv4AddressRange {
	if o == nil || o.Ipv4AddressRanges == nil {
		var ret []Ipv4AddressRange
		return ret
	}
	return *o.Ipv4AddressRanges
}

// GetIpv4AddressRangesOk returns a tuple with the Ipv4AddressRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnnUpfInfoItem) GetIpv4AddressRangesOk() (*[]Ipv4AddressRange, bool) {
	if o == nil || o.Ipv4AddressRanges == nil {
		return nil, false
	}
	return o.Ipv4AddressRanges, true
}

// HasIpv4AddressRanges returns a boolean if a field has been set.
func (o *DnnUpfInfoItem) HasIpv4AddressRanges() bool {
	if o != nil && o.Ipv4AddressRanges != nil {
		return true
	}

	return false
}

// SetIpv4AddressRanges gets a reference to the given []Ipv4AddressRange and assigns it to the Ipv4AddressRanges field.
func (o *DnnUpfInfoItem) SetIpv4AddressRanges(v []Ipv4AddressRange) {
	o.Ipv4AddressRanges = &v
}

// GetIpv6PrefixRanges returns the Ipv6PrefixRanges field value if set, zero value otherwise.
func (o *DnnUpfInfoItem) GetIpv6PrefixRanges() []Ipv6PrefixRange {
	if o == nil || o.Ipv6PrefixRanges == nil {
		var ret []Ipv6PrefixRange
		return ret
	}
	return *o.Ipv6PrefixRanges
}

// GetIpv6PrefixRangesOk returns a tuple with the Ipv6PrefixRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnnUpfInfoItem) GetIpv6PrefixRangesOk() (*[]Ipv6PrefixRange, bool) {
	if o == nil || o.Ipv6PrefixRanges == nil {
		return nil, false
	}
	return o.Ipv6PrefixRanges, true
}

// HasIpv6PrefixRanges returns a boolean if a field has been set.
func (o *DnnUpfInfoItem) HasIpv6PrefixRanges() bool {
	if o != nil && o.Ipv6PrefixRanges != nil {
		return true
	}

	return false
}

// SetIpv6PrefixRanges gets a reference to the given []Ipv6PrefixRange and assigns it to the Ipv6PrefixRanges field.
func (o *DnnUpfInfoItem) SetIpv6PrefixRanges(v []Ipv6PrefixRange) {
	o.Ipv6PrefixRanges = &v
}

// GetDnaiNwInstanceList returns the DnaiNwInstanceList field value if set, zero value otherwise.
func (o *DnnUpfInfoItem) GetDnaiNwInstanceList() map[string]string {
	if o == nil || o.DnaiNwInstanceList == nil {
		var ret map[string]string
		return ret
	}
	return *o.DnaiNwInstanceList
}

// GetDnaiNwInstanceListOk returns a tuple with the DnaiNwInstanceList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnnUpfInfoItem) GetDnaiNwInstanceListOk() (*map[string]string, bool) {
	if o == nil || o.DnaiNwInstanceList == nil {
		return nil, false
	}
	return o.DnaiNwInstanceList, true
}

// HasDnaiNwInstanceList returns a boolean if a field has been set.
func (o *DnnUpfInfoItem) HasDnaiNwInstanceList() bool {
	if o != nil && o.DnaiNwInstanceList != nil {
		return true
	}

	return false
}

// SetDnaiNwInstanceList gets a reference to the given map[string]string and assigns it to the DnaiNwInstanceList field.
func (o *DnnUpfInfoItem) SetDnaiNwInstanceList(v map[string]string) {
	o.DnaiNwInstanceList = &v
}

func (o DnnUpfInfoItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["dnn"] = o.Dnn
	}
	if o.DnaiList != nil {
		toSerialize["dnaiList"] = o.DnaiList
	}
	if o.PduSessionTypes != nil {
		toSerialize["pduSessionTypes"] = o.PduSessionTypes
	}
	if o.Ipv4AddressRanges != nil {
		toSerialize["ipv4AddressRanges"] = o.Ipv4AddressRanges
	}
	if o.Ipv6PrefixRanges != nil {
		toSerialize["ipv6PrefixRanges"] = o.Ipv6PrefixRanges
	}
	if o.DnaiNwInstanceList != nil {
		toSerialize["dnaiNwInstanceList"] = o.DnaiNwInstanceList
	}
	return json.Marshal(toSerialize)
}

type NullableDnnUpfInfoItem struct {
	value *DnnUpfInfoItem
	isSet bool
}

func (v NullableDnnUpfInfoItem) Get() *DnnUpfInfoItem {
	return v.value
}

func (v *NullableDnnUpfInfoItem) Set(val *DnnUpfInfoItem) {
	v.value = val
	v.isSet = true
}

func (v NullableDnnUpfInfoItem) IsSet() bool {
	return v.isSet
}

func (v *NullableDnnUpfInfoItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnnUpfInfoItem(val *DnnUpfInfoItem) *NullableDnnUpfInfoItem {
	return &NullableDnnUpfInfoItem{value: val, isSet: true}
}

func (v NullableDnnUpfInfoItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnnUpfInfoItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

