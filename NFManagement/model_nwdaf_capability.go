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

// NwdafCapability Indicates the capability supported by the NWDAF
type NwdafCapability struct {
	AnalyticsAggregation *bool `json:"analyticsAggregation,omitempty"`
	AnalyticsMetadataProvisioning *bool `json:"analyticsMetadataProvisioning,omitempty"`
}

// NewNwdafCapability instantiates a new NwdafCapability object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNwdafCapability() *NwdafCapability {
	this := NwdafCapability{}
	var analyticsAggregation bool = false
	this.AnalyticsAggregation = &analyticsAggregation
	var analyticsMetadataProvisioning bool = false
	this.AnalyticsMetadataProvisioning = &analyticsMetadataProvisioning
	return &this
}

// NewNwdafCapabilityWithDefaults instantiates a new NwdafCapability object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNwdafCapabilityWithDefaults() *NwdafCapability {
	this := NwdafCapability{}
	var analyticsAggregation bool = false
	this.AnalyticsAggregation = &analyticsAggregation
	var analyticsMetadataProvisioning bool = false
	this.AnalyticsMetadataProvisioning = &analyticsMetadataProvisioning
	return &this
}

// GetAnalyticsAggregation returns the AnalyticsAggregation field value if set, zero value otherwise.
func (o *NwdafCapability) GetAnalyticsAggregation() bool {
	if o == nil || o.AnalyticsAggregation == nil {
		var ret bool
		return ret
	}
	return *o.AnalyticsAggregation
}

// GetAnalyticsAggregationOk returns a tuple with the AnalyticsAggregation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NwdafCapability) GetAnalyticsAggregationOk() (*bool, bool) {
	if o == nil || o.AnalyticsAggregation == nil {
		return nil, false
	}
	return o.AnalyticsAggregation, true
}

// HasAnalyticsAggregation returns a boolean if a field has been set.
func (o *NwdafCapability) HasAnalyticsAggregation() bool {
	if o != nil && o.AnalyticsAggregation != nil {
		return true
	}

	return false
}

// SetAnalyticsAggregation gets a reference to the given bool and assigns it to the AnalyticsAggregation field.
func (o *NwdafCapability) SetAnalyticsAggregation(v bool) {
	o.AnalyticsAggregation = &v
}

// GetAnalyticsMetadataProvisioning returns the AnalyticsMetadataProvisioning field value if set, zero value otherwise.
func (o *NwdafCapability) GetAnalyticsMetadataProvisioning() bool {
	if o == nil || o.AnalyticsMetadataProvisioning == nil {
		var ret bool
		return ret
	}
	return *o.AnalyticsMetadataProvisioning
}

// GetAnalyticsMetadataProvisioningOk returns a tuple with the AnalyticsMetadataProvisioning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NwdafCapability) GetAnalyticsMetadataProvisioningOk() (*bool, bool) {
	if o == nil || o.AnalyticsMetadataProvisioning == nil {
		return nil, false
	}
	return o.AnalyticsMetadataProvisioning, true
}

// HasAnalyticsMetadataProvisioning returns a boolean if a field has been set.
func (o *NwdafCapability) HasAnalyticsMetadataProvisioning() bool {
	if o != nil && o.AnalyticsMetadataProvisioning != nil {
		return true
	}

	return false
}

// SetAnalyticsMetadataProvisioning gets a reference to the given bool and assigns it to the AnalyticsMetadataProvisioning field.
func (o *NwdafCapability) SetAnalyticsMetadataProvisioning(v bool) {
	o.AnalyticsMetadataProvisioning = &v
}

func (o NwdafCapability) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AnalyticsAggregation != nil {
		toSerialize["analyticsAggregation"] = o.AnalyticsAggregation
	}
	if o.AnalyticsMetadataProvisioning != nil {
		toSerialize["analyticsMetadataProvisioning"] = o.AnalyticsMetadataProvisioning
	}
	return json.Marshal(toSerialize)
}

type NullableNwdafCapability struct {
	value *NwdafCapability
	isSet bool
}

func (v NullableNwdafCapability) Get() *NwdafCapability {
	return v.value
}

func (v *NullableNwdafCapability) Set(val *NwdafCapability) {
	v.value = val
	v.isSet = true
}

func (v NullableNwdafCapability) IsSet() bool {
	return v.isSet
}

func (v *NullableNwdafCapability) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNwdafCapability(val *NwdafCapability) *NullableNwdafCapability {
	return &NullableNwdafCapability{value: val, isSet: true}
}

func (v NullableNwdafCapability) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNwdafCapability) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


