# TsctsfInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SNssaiInfoList** | Pointer to [**map[string]SnssaiTsctsfInfoItem**](SnssaiTsctsfInfoItem.md) | A map (list of key-value pairs) where a valid JSON string serves as key | [optional] 
**ExternalGroupIdentifiersRanges** | Pointer to [**[]IdentityRange**](IdentityRange.md) |  | [optional] 

## Methods

### NewTsctsfInfo

`func NewTsctsfInfo() *TsctsfInfo`

NewTsctsfInfo instantiates a new TsctsfInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTsctsfInfoWithDefaults

`func NewTsctsfInfoWithDefaults() *TsctsfInfo`

NewTsctsfInfoWithDefaults instantiates a new TsctsfInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSNssaiInfoList

`func (o *TsctsfInfo) GetSNssaiInfoList() map[string]SnssaiTsctsfInfoItem`

GetSNssaiInfoList returns the SNssaiInfoList field if non-nil, zero value otherwise.

### GetSNssaiInfoListOk

`func (o *TsctsfInfo) GetSNssaiInfoListOk() (*map[string]SnssaiTsctsfInfoItem, bool)`

GetSNssaiInfoListOk returns a tuple with the SNssaiInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSNssaiInfoList

`func (o *TsctsfInfo) SetSNssaiInfoList(v map[string]SnssaiTsctsfInfoItem)`

SetSNssaiInfoList sets SNssaiInfoList field to given value.

### HasSNssaiInfoList

`func (o *TsctsfInfo) HasSNssaiInfoList() bool`

HasSNssaiInfoList returns a boolean if a field has been set.

### GetExternalGroupIdentifiersRanges

`func (o *TsctsfInfo) GetExternalGroupIdentifiersRanges() []IdentityRange`

GetExternalGroupIdentifiersRanges returns the ExternalGroupIdentifiersRanges field if non-nil, zero value otherwise.

### GetExternalGroupIdentifiersRangesOk

`func (o *TsctsfInfo) GetExternalGroupIdentifiersRangesOk() (*[]IdentityRange, bool)`

GetExternalGroupIdentifiersRangesOk returns a tuple with the ExternalGroupIdentifiersRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalGroupIdentifiersRanges

`func (o *TsctsfInfo) SetExternalGroupIdentifiersRanges(v []IdentityRange)`

SetExternalGroupIdentifiersRanges sets ExternalGroupIdentifiersRanges field to given value.

### HasExternalGroupIdentifiersRanges

`func (o *TsctsfInfo) HasExternalGroupIdentifiersRanges() bool`

HasExternalGroupIdentifiersRanges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


