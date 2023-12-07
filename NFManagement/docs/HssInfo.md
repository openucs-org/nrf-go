# HssInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | Pointer to **string** | Identifier of a group of NFs. | [optional] 
**ImsiRanges** | Pointer to [**[]ImsiRange**](ImsiRange.md) |  | [optional] 
**ImsPrivateIdentityRanges** | Pointer to [**[]IdentityRange**](IdentityRange.md) |  | [optional] 
**ImsPublicIdentityRanges** | Pointer to [**[]IdentityRange**](IdentityRange.md) |  | [optional] 
**MsisdnRanges** | Pointer to [**[]IdentityRange**](IdentityRange.md) |  | [optional] 

## Methods

### NewHssInfo

`func NewHssInfo() *HssInfo`

NewHssInfo instantiates a new HssInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHssInfoWithDefaults

`func NewHssInfoWithDefaults() *HssInfo`

NewHssInfoWithDefaults instantiates a new HssInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *HssInfo) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *HssInfo) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *HssInfo) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *HssInfo) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetImsiRanges

`func (o *HssInfo) GetImsiRanges() []ImsiRange`

GetImsiRanges returns the ImsiRanges field if non-nil, zero value otherwise.

### GetImsiRangesOk

`func (o *HssInfo) GetImsiRangesOk() (*[]ImsiRange, bool)`

GetImsiRangesOk returns a tuple with the ImsiRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImsiRanges

`func (o *HssInfo) SetImsiRanges(v []ImsiRange)`

SetImsiRanges sets ImsiRanges field to given value.

### HasImsiRanges

`func (o *HssInfo) HasImsiRanges() bool`

HasImsiRanges returns a boolean if a field has been set.

### GetImsPrivateIdentityRanges

`func (o *HssInfo) GetImsPrivateIdentityRanges() []IdentityRange`

GetImsPrivateIdentityRanges returns the ImsPrivateIdentityRanges field if non-nil, zero value otherwise.

### GetImsPrivateIdentityRangesOk

`func (o *HssInfo) GetImsPrivateIdentityRangesOk() (*[]IdentityRange, bool)`

GetImsPrivateIdentityRangesOk returns a tuple with the ImsPrivateIdentityRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImsPrivateIdentityRanges

`func (o *HssInfo) SetImsPrivateIdentityRanges(v []IdentityRange)`

SetImsPrivateIdentityRanges sets ImsPrivateIdentityRanges field to given value.

### HasImsPrivateIdentityRanges

`func (o *HssInfo) HasImsPrivateIdentityRanges() bool`

HasImsPrivateIdentityRanges returns a boolean if a field has been set.

### GetImsPublicIdentityRanges

`func (o *HssInfo) GetImsPublicIdentityRanges() []IdentityRange`

GetImsPublicIdentityRanges returns the ImsPublicIdentityRanges field if non-nil, zero value otherwise.

### GetImsPublicIdentityRangesOk

`func (o *HssInfo) GetImsPublicIdentityRangesOk() (*[]IdentityRange, bool)`

GetImsPublicIdentityRangesOk returns a tuple with the ImsPublicIdentityRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImsPublicIdentityRanges

`func (o *HssInfo) SetImsPublicIdentityRanges(v []IdentityRange)`

SetImsPublicIdentityRanges sets ImsPublicIdentityRanges field to given value.

### HasImsPublicIdentityRanges

`func (o *HssInfo) HasImsPublicIdentityRanges() bool`

HasImsPublicIdentityRanges returns a boolean if a field has been set.

### GetMsisdnRanges

`func (o *HssInfo) GetMsisdnRanges() []IdentityRange`

GetMsisdnRanges returns the MsisdnRanges field if non-nil, zero value otherwise.

### GetMsisdnRangesOk

`func (o *HssInfo) GetMsisdnRangesOk() (*[]IdentityRange, bool)`

GetMsisdnRangesOk returns a tuple with the MsisdnRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsisdnRanges

`func (o *HssInfo) SetMsisdnRanges(v []IdentityRange)`

SetMsisdnRanges sets MsisdnRanges field to given value.

### HasMsisdnRanges

`func (o *HssInfo) HasMsisdnRanges() bool`

HasMsisdnRanges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


