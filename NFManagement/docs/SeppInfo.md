# SeppInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SeppPorts** | Pointer to **map[string]int32** | Port numbers for HTTP and HTTPS. The key of the map shall be \&quot;http\&quot; or \&quot;https\&quot; | [optional] 
**RemotePlmnList** | Pointer to [**[]PlmnId**](PlmnId.md) |  | [optional] 

## Methods

### NewSeppInfo

`func NewSeppInfo() *SeppInfo`

NewSeppInfo instantiates a new SeppInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSeppInfoWithDefaults

`func NewSeppInfoWithDefaults() *SeppInfo`

NewSeppInfoWithDefaults instantiates a new SeppInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSeppPorts

`func (o *SeppInfo) GetSeppPorts() map[string]int32`

GetSeppPorts returns the SeppPorts field if non-nil, zero value otherwise.

### GetSeppPortsOk

`func (o *SeppInfo) GetSeppPortsOk() (*map[string]int32, bool)`

GetSeppPortsOk returns a tuple with the SeppPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeppPorts

`func (o *SeppInfo) SetSeppPorts(v map[string]int32)`

SetSeppPorts sets SeppPorts field to given value.

### HasSeppPorts

`func (o *SeppInfo) HasSeppPorts() bool`

HasSeppPorts returns a boolean if a field has been set.

### GetRemotePlmnList

`func (o *SeppInfo) GetRemotePlmnList() []PlmnId`

GetRemotePlmnList returns the RemotePlmnList field if non-nil, zero value otherwise.

### GetRemotePlmnListOk

`func (o *SeppInfo) GetRemotePlmnListOk() (*[]PlmnId, bool)`

GetRemotePlmnListOk returns a tuple with the RemotePlmnList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotePlmnList

`func (o *SeppInfo) SetRemotePlmnList(v []PlmnId)`

SetRemotePlmnList sets RemotePlmnList field to given value.

### HasRemotePlmnList

`func (o *SeppInfo) HasRemotePlmnList() bool`

HasRemotePlmnList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


