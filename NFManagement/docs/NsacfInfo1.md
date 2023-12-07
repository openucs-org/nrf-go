# NsacfInfo1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NsacfCapability** | [**NsacfCapability**](NsacfCapability.md) |  | 
**TaiList** | Pointer to [**[]Tai**](Tai.md) |  | [optional] 
**TaiRangeList** | Pointer to [**[]TaiRange**](TaiRange.md) |  | [optional] 
**OtherServingAreas** | Pointer to **[]string** |  | [optional] 

## Methods

### NewNsacfInfo1

`func NewNsacfInfo1(nsacfCapability NsacfCapability, ) *NsacfInfo1`

NewNsacfInfo1 instantiates a new NsacfInfo1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNsacfInfo1WithDefaults

`func NewNsacfInfo1WithDefaults() *NsacfInfo1`

NewNsacfInfo1WithDefaults instantiates a new NsacfInfo1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNsacfCapability

`func (o *NsacfInfo1) GetNsacfCapability() NsacfCapability`

GetNsacfCapability returns the NsacfCapability field if non-nil, zero value otherwise.

### GetNsacfCapabilityOk

`func (o *NsacfInfo1) GetNsacfCapabilityOk() (*NsacfCapability, bool)`

GetNsacfCapabilityOk returns a tuple with the NsacfCapability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsacfCapability

`func (o *NsacfInfo1) SetNsacfCapability(v NsacfCapability)`

SetNsacfCapability sets NsacfCapability field to given value.


### GetTaiList

`func (o *NsacfInfo1) GetTaiList() []Tai`

GetTaiList returns the TaiList field if non-nil, zero value otherwise.

### GetTaiListOk

`func (o *NsacfInfo1) GetTaiListOk() (*[]Tai, bool)`

GetTaiListOk returns a tuple with the TaiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaiList

`func (o *NsacfInfo1) SetTaiList(v []Tai)`

SetTaiList sets TaiList field to given value.

### HasTaiList

`func (o *NsacfInfo1) HasTaiList() bool`

HasTaiList returns a boolean if a field has been set.

### GetTaiRangeList

`func (o *NsacfInfo1) GetTaiRangeList() []TaiRange`

GetTaiRangeList returns the TaiRangeList field if non-nil, zero value otherwise.

### GetTaiRangeListOk

`func (o *NsacfInfo1) GetTaiRangeListOk() (*[]TaiRange, bool)`

GetTaiRangeListOk returns a tuple with the TaiRangeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaiRangeList

`func (o *NsacfInfo1) SetTaiRangeList(v []TaiRange)`

SetTaiRangeList sets TaiRangeList field to given value.

### HasTaiRangeList

`func (o *NsacfInfo1) HasTaiRangeList() bool`

HasTaiRangeList returns a boolean if a field has been set.

### GetOtherServingAreas

`func (o *NsacfInfo1) GetOtherServingAreas() []string`

GetOtherServingAreas returns the OtherServingAreas field if non-nil, zero value otherwise.

### GetOtherServingAreasOk

`func (o *NsacfInfo1) GetOtherServingAreasOk() (*[]string, bool)`

GetOtherServingAreasOk returns a tuple with the OtherServingAreas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherServingAreas

`func (o *NsacfInfo1) SetOtherServingAreas(v []string)`

SetOtherServingAreas sets OtherServingAreas field to given value.

### HasOtherServingAreas

`func (o *NsacfInfo1) HasOtherServingAreas() bool`

HasOtherServingAreas returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


