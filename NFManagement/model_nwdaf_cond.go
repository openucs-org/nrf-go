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

// NwdafCond Subscription to a set of NF Instances (NWDAFs), identified by Analytics ID(s), S-NSSAI(s) or NWDAF Serving Area information, i.e. list of TAIs for which the NWDAF can provide analytics.
type NwdafCond struct {
	ConditionType string `json:"conditionType"`
	AnalyticsIds *[]string `json:"analyticsIds,omitempty"`
	SnssaiList *[]Snssai `json:"snssaiList,omitempty"`
	TaiList *[]Tai `json:"taiList,omitempty"`
	TaiRangeList *[]TaiRange `json:"taiRangeList,omitempty"`
	ServingNfTypeList *[]NFType `json:"servingNfTypeList,omitempty"`
	ServingNfSetIdList *[]string `json:"servingNfSetIdList,omitempty"`
	MlAnalyticsList *[]MlAnalyticsInfo `json:"mlAnalyticsList,omitempty"`
}

// NewNwdafCond instantiates a new NwdafCond object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNwdafCond(conditionType string) *NwdafCond {
	this := NwdafCond{}
	this.ConditionType = conditionType
	return &this
}

// NewNwdafCondWithDefaults instantiates a new NwdafCond object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNwdafCondWithDefaults() *NwdafCond {
	this := NwdafCond{}
	return &this
}

// GetConditionType returns the ConditionType field value
func (o *NwdafCond) GetConditionType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConditionType
}

// GetConditionTypeOk returns a tuple with the ConditionType field value
// and a boolean to check if the value has been set.
func (o *NwdafCond) GetConditionTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ConditionType, true
}

// SetConditionType sets field value
func (o *NwdafCond) SetConditionType(v string) {
	o.ConditionType = v
}

// GetAnalyticsIds returns the AnalyticsIds field value if set, zero value otherwise.
func (o *NwdafCond) GetAnalyticsIds() []string {
	if o == nil || o.AnalyticsIds == nil {
		var ret []string
		return ret
	}
	return *o.AnalyticsIds
}

// GetAnalyticsIdsOk returns a tuple with the AnalyticsIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NwdafCond) GetAnalyticsIdsOk() (*[]string, bool) {
	if o == nil || o.AnalyticsIds == nil {
		return nil, false
	}
	return o.AnalyticsIds, true
}

// HasAnalyticsIds returns a boolean if a field has been set.
func (o *NwdafCond) HasAnalyticsIds() bool {
	if o != nil && o.AnalyticsIds != nil {
		return true
	}

	return false
}

// SetAnalyticsIds gets a reference to the given []string and assigns it to the AnalyticsIds field.
func (o *NwdafCond) SetAnalyticsIds(v []string) {
	o.AnalyticsIds = &v
}

// GetSnssaiList returns the SnssaiList field value if set, zero value otherwise.
func (o *NwdafCond) GetSnssaiList() []Snssai {
	if o == nil || o.SnssaiList == nil {
		var ret []Snssai
		return ret
	}
	return *o.SnssaiList
}

// GetSnssaiListOk returns a tuple with the SnssaiList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NwdafCond) GetSnssaiListOk() (*[]Snssai, bool) {
	if o == nil || o.SnssaiList == nil {
		return nil, false
	}
	return o.SnssaiList, true
}

// HasSnssaiList returns a boolean if a field has been set.
func (o *NwdafCond) HasSnssaiList() bool {
	if o != nil && o.SnssaiList != nil {
		return true
	}

	return false
}

// SetSnssaiList gets a reference to the given []Snssai and assigns it to the SnssaiList field.
func (o *NwdafCond) SetSnssaiList(v []Snssai) {
	o.SnssaiList = &v
}

// GetTaiList returns the TaiList field value if set, zero value otherwise.
func (o *NwdafCond) GetTaiList() []Tai {
	if o == nil || o.TaiList == nil {
		var ret []Tai
		return ret
	}
	return *o.TaiList
}

// GetTaiListOk returns a tuple with the TaiList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NwdafCond) GetTaiListOk() (*[]Tai, bool) {
	if o == nil || o.TaiList == nil {
		return nil, false
	}
	return o.TaiList, true
}

// HasTaiList returns a boolean if a field has been set.
func (o *NwdafCond) HasTaiList() bool {
	if o != nil && o.TaiList != nil {
		return true
	}

	return false
}

// SetTaiList gets a reference to the given []Tai and assigns it to the TaiList field.
func (o *NwdafCond) SetTaiList(v []Tai) {
	o.TaiList = &v
}

// GetTaiRangeList returns the TaiRangeList field value if set, zero value otherwise.
func (o *NwdafCond) GetTaiRangeList() []TaiRange {
	if o == nil || o.TaiRangeList == nil {
		var ret []TaiRange
		return ret
	}
	return *o.TaiRangeList
}

// GetTaiRangeListOk returns a tuple with the TaiRangeList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NwdafCond) GetTaiRangeListOk() (*[]TaiRange, bool) {
	if o == nil || o.TaiRangeList == nil {
		return nil, false
	}
	return o.TaiRangeList, true
}

// HasTaiRangeList returns a boolean if a field has been set.
func (o *NwdafCond) HasTaiRangeList() bool {
	if o != nil && o.TaiRangeList != nil {
		return true
	}

	return false
}

// SetTaiRangeList gets a reference to the given []TaiRange and assigns it to the TaiRangeList field.
func (o *NwdafCond) SetTaiRangeList(v []TaiRange) {
	o.TaiRangeList = &v
}

// GetServingNfTypeList returns the ServingNfTypeList field value if set, zero value otherwise.
func (o *NwdafCond) GetServingNfTypeList() []NFType {
	if o == nil || o.ServingNfTypeList == nil {
		var ret []NFType
		return ret
	}
	return *o.ServingNfTypeList
}

// GetServingNfTypeListOk returns a tuple with the ServingNfTypeList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NwdafCond) GetServingNfTypeListOk() (*[]NFType, bool) {
	if o == nil || o.ServingNfTypeList == nil {
		return nil, false
	}
	return o.ServingNfTypeList, true
}

// HasServingNfTypeList returns a boolean if a field has been set.
func (o *NwdafCond) HasServingNfTypeList() bool {
	if o != nil && o.ServingNfTypeList != nil {
		return true
	}

	return false
}

// SetServingNfTypeList gets a reference to the given []NFType and assigns it to the ServingNfTypeList field.
func (o *NwdafCond) SetServingNfTypeList(v []NFType) {
	o.ServingNfTypeList = &v
}

// GetServingNfSetIdList returns the ServingNfSetIdList field value if set, zero value otherwise.
func (o *NwdafCond) GetServingNfSetIdList() []string {
	if o == nil || o.ServingNfSetIdList == nil {
		var ret []string
		return ret
	}
	return *o.ServingNfSetIdList
}

// GetServingNfSetIdListOk returns a tuple with the ServingNfSetIdList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NwdafCond) GetServingNfSetIdListOk() (*[]string, bool) {
	if o == nil || o.ServingNfSetIdList == nil {
		return nil, false
	}
	return o.ServingNfSetIdList, true
}

// HasServingNfSetIdList returns a boolean if a field has been set.
func (o *NwdafCond) HasServingNfSetIdList() bool {
	if o != nil && o.ServingNfSetIdList != nil {
		return true
	}

	return false
}

// SetServingNfSetIdList gets a reference to the given []string and assigns it to the ServingNfSetIdList field.
func (o *NwdafCond) SetServingNfSetIdList(v []string) {
	o.ServingNfSetIdList = &v
}

// GetMlAnalyticsList returns the MlAnalyticsList field value if set, zero value otherwise.
func (o *NwdafCond) GetMlAnalyticsList() []MlAnalyticsInfo {
	if o == nil || o.MlAnalyticsList == nil {
		var ret []MlAnalyticsInfo
		return ret
	}
	return *o.MlAnalyticsList
}

// GetMlAnalyticsListOk returns a tuple with the MlAnalyticsList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NwdafCond) GetMlAnalyticsListOk() (*[]MlAnalyticsInfo, bool) {
	if o == nil || o.MlAnalyticsList == nil {
		return nil, false
	}
	return o.MlAnalyticsList, true
}

// HasMlAnalyticsList returns a boolean if a field has been set.
func (o *NwdafCond) HasMlAnalyticsList() bool {
	if o != nil && o.MlAnalyticsList != nil {
		return true
	}

	return false
}

// SetMlAnalyticsList gets a reference to the given []MlAnalyticsInfo and assigns it to the MlAnalyticsList field.
func (o *NwdafCond) SetMlAnalyticsList(v []MlAnalyticsInfo) {
	o.MlAnalyticsList = &v
}

func (o NwdafCond) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["conditionType"] = o.ConditionType
	}
	if o.AnalyticsIds != nil {
		toSerialize["analyticsIds"] = o.AnalyticsIds
	}
	if o.SnssaiList != nil {
		toSerialize["snssaiList"] = o.SnssaiList
	}
	if o.TaiList != nil {
		toSerialize["taiList"] = o.TaiList
	}
	if o.TaiRangeList != nil {
		toSerialize["taiRangeList"] = o.TaiRangeList
	}
	if o.ServingNfTypeList != nil {
		toSerialize["servingNfTypeList"] = o.ServingNfTypeList
	}
	if o.ServingNfSetIdList != nil {
		toSerialize["servingNfSetIdList"] = o.ServingNfSetIdList
	}
	if o.MlAnalyticsList != nil {
		toSerialize["mlAnalyticsList"] = o.MlAnalyticsList
	}
	return json.Marshal(toSerialize)
}

type NullableNwdafCond struct {
	value *NwdafCond
	isSet bool
}

func (v NullableNwdafCond) Get() *NwdafCond {
	return v.value
}

func (v *NullableNwdafCond) Set(val *NwdafCond) {
	v.value = val
	v.isSet = true
}

func (v NullableNwdafCond) IsSet() bool {
	return v.isSet
}

func (v *NullableNwdafCond) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNwdafCond(val *NwdafCond) *NullableNwdafCond {
	return &NullableNwdafCond{value: val, isSet: true}
}

func (v NullableNwdafCond) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNwdafCond) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


