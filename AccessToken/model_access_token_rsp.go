/*
NRF OAuth2

NRF OAuth2 Authorization. © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.2.0-alpha.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// AccessTokenRsp Contains information related to the access token response
type AccessTokenRsp struct {
	// JWS Compact Serialized representation of JWS signed JSON object (AccessTokenClaims)
	AccessToken string `json:"access_token"`
	TokenType string `json:"token_type"`
	ExpiresIn *int32 `json:"expires_in,omitempty"`
	Scope *string `json:"scope,omitempty"`
}

// NewAccessTokenRsp instantiates a new AccessTokenRsp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessTokenRsp(accessToken string, tokenType string) *AccessTokenRsp {
	this := AccessTokenRsp{}
	this.AccessToken = accessToken
	this.TokenType = tokenType
	return &this
}

// NewAccessTokenRspWithDefaults instantiates a new AccessTokenRsp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessTokenRspWithDefaults() *AccessTokenRsp {
	this := AccessTokenRsp{}
	return &this
}

// GetAccessToken returns the AccessToken field value
func (o *AccessTokenRsp) GetAccessToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value
// and a boolean to check if the value has been set.
func (o *AccessTokenRsp) GetAccessTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccessToken, true
}

// SetAccessToken sets field value
func (o *AccessTokenRsp) SetAccessToken(v string) {
	o.AccessToken = v
}

// GetTokenType returns the TokenType field value
func (o *AccessTokenRsp) GetTokenType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenType
}

// GetTokenTypeOk returns a tuple with the TokenType field value
// and a boolean to check if the value has been set.
func (o *AccessTokenRsp) GetTokenTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TokenType, true
}

// SetTokenType sets field value
func (o *AccessTokenRsp) SetTokenType(v string) {
	o.TokenType = v
}

// GetExpiresIn returns the ExpiresIn field value if set, zero value otherwise.
func (o *AccessTokenRsp) GetExpiresIn() int32 {
	if o == nil || o.ExpiresIn == nil {
		var ret int32
		return ret
	}
	return *o.ExpiresIn
}

// GetExpiresInOk returns a tuple with the ExpiresIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessTokenRsp) GetExpiresInOk() (*int32, bool) {
	if o == nil || o.ExpiresIn == nil {
		return nil, false
	}
	return o.ExpiresIn, true
}

// HasExpiresIn returns a boolean if a field has been set.
func (o *AccessTokenRsp) HasExpiresIn() bool {
	if o != nil && o.ExpiresIn != nil {
		return true
	}

	return false
}

// SetExpiresIn gets a reference to the given int32 and assigns it to the ExpiresIn field.
func (o *AccessTokenRsp) SetExpiresIn(v int32) {
	o.ExpiresIn = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *AccessTokenRsp) GetScope() string {
	if o == nil || o.Scope == nil {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessTokenRsp) GetScopeOk() (*string, bool) {
	if o == nil || o.Scope == nil {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *AccessTokenRsp) HasScope() bool {
	if o != nil && o.Scope != nil {
		return true
	}

	return false
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *AccessTokenRsp) SetScope(v string) {
	o.Scope = &v
}

func (o AccessTokenRsp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["access_token"] = o.AccessToken
	}
	if true {
		toSerialize["token_type"] = o.TokenType
	}
	if o.ExpiresIn != nil {
		toSerialize["expires_in"] = o.ExpiresIn
	}
	if o.Scope != nil {
		toSerialize["scope"] = o.Scope
	}
	return json.Marshal(toSerialize)
}

type NullableAccessTokenRsp struct {
	value *AccessTokenRsp
	isSet bool
}

func (v NullableAccessTokenRsp) Get() *AccessTokenRsp {
	return v.value
}

func (v *NullableAccessTokenRsp) Set(val *AccessTokenRsp) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessTokenRsp) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessTokenRsp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessTokenRsp(val *AccessTokenRsp) *NullableAccessTokenRsp {
	return &NullableAccessTokenRsp{value: val, isSet: true}
}

func (v NullableAccessTokenRsp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessTokenRsp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


