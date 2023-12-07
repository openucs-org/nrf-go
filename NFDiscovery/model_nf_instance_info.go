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

// NfInstanceInfo Contains information on an NF profile matching a discovery request
type NfInstanceInfo struct {
	// String providing an URI formatted according to RFC 3986
	NrfDiscApiUri *string `json:"nrfDiscApiUri,omitempty"`
	PreferredSearch *PreferredSearch `json:"preferredSearch,omitempty"`
}

// NewNfInstanceInfo instantiates a new NfInstanceInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNfInstanceInfo() *NfInstanceInfo {
	this := NfInstanceInfo{}
	return &this
}

// NewNfInstanceInfoWithDefaults instantiates a new NfInstanceInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNfInstanceInfoWithDefaults() *NfInstanceInfo {
	this := NfInstanceInfo{}
	return &this
}

// GetNrfDiscApiUri returns the NrfDiscApiUri field value if set, zero value otherwise.
func (o *NfInstanceInfo) GetNrfDiscApiUri() string {
	if o == nil || o.NrfDiscApiUri == nil {
		var ret string
		return ret
	}
	return *o.NrfDiscApiUri
}

// GetNrfDiscApiUriOk returns a tuple with the NrfDiscApiUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NfInstanceInfo) GetNrfDiscApiUriOk() (*string, bool) {
	if o == nil || o.NrfDiscApiUri == nil {
		return nil, false
	}
	return o.NrfDiscApiUri, true
}

// HasNrfDiscApiUri returns a boolean if a field has been set.
func (o *NfInstanceInfo) HasNrfDiscApiUri() bool {
	if o != nil && o.NrfDiscApiUri != nil {
		return true
	}

	return false
}

// SetNrfDiscApiUri gets a reference to the given string and assigns it to the NrfDiscApiUri field.
func (o *NfInstanceInfo) SetNrfDiscApiUri(v string) {
	o.NrfDiscApiUri = &v
}

// GetPreferredSearch returns the PreferredSearch field value if set, zero value otherwise.
func (o *NfInstanceInfo) GetPreferredSearch() PreferredSearch {
	if o == nil || o.PreferredSearch == nil {
		var ret PreferredSearch
		return ret
	}
	return *o.PreferredSearch
}

// GetPreferredSearchOk returns a tuple with the PreferredSearch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NfInstanceInfo) GetPreferredSearchOk() (*PreferredSearch, bool) {
	if o == nil || o.PreferredSearch == nil {
		return nil, false
	}
	return o.PreferredSearch, true
}

// HasPreferredSearch returns a boolean if a field has been set.
func (o *NfInstanceInfo) HasPreferredSearch() bool {
	if o != nil && o.PreferredSearch != nil {
		return true
	}

	return false
}

// SetPreferredSearch gets a reference to the given PreferredSearch and assigns it to the PreferredSearch field.
func (o *NfInstanceInfo) SetPreferredSearch(v PreferredSearch) {
	o.PreferredSearch = &v
}

func (o NfInstanceInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NrfDiscApiUri != nil {
		toSerialize["nrfDiscApiUri"] = o.NrfDiscApiUri
	}
	if o.PreferredSearch != nil {
		toSerialize["preferredSearch"] = o.PreferredSearch
	}
	return json.Marshal(toSerialize)
}

type NullableNfInstanceInfo struct {
	value *NfInstanceInfo
	isSet bool
}

func (v NullableNfInstanceInfo) Get() *NfInstanceInfo {
	return v.value
}

func (v *NullableNfInstanceInfo) Set(val *NfInstanceInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableNfInstanceInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableNfInstanceInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNfInstanceInfo(val *NfInstanceInfo) *NullableNfInstanceInfo {
	return &NullableNfInstanceInfo{value: val, isSet: true}
}

func (v NullableNfInstanceInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNfInstanceInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

