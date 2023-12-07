# PatchItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Op** | **interface{}** | indicates the patch operation as defined in IETF RFC 6902 to be performed on the resource. | 
**Path** | **string** | contains a JSON pointer value (as defined in IETF RFC 6901) that references a location of a resource on which the patch operation shall be performed. | 
**From** | Pointer to **string** | indicates the path of the source JSON element (according to JSON Pointer syntax) being moved or copied to the location indicated by the \&quot;path\&quot; attribute. | [optional] 
**Value** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewPatchItem

`func NewPatchItem(op interface{}, path string, ) *PatchItem`

NewPatchItem instantiates a new PatchItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchItemWithDefaults

`func NewPatchItemWithDefaults() *PatchItem`

NewPatchItemWithDefaults instantiates a new PatchItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOp

`func (o *PatchItem) GetOp() interface{}`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *PatchItem) GetOpOk() (*interface{}, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *PatchItem) SetOp(v interface{})`

SetOp sets Op field to given value.


### SetOpNil

`func (o *PatchItem) SetOpNil(b bool)`

 SetOpNil sets the value for Op to be an explicit nil

### UnsetOp
`func (o *PatchItem) UnsetOp()`

UnsetOp ensures that no value is present for Op, not even an explicit nil
### GetPath

`func (o *PatchItem) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *PatchItem) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *PatchItem) SetPath(v string)`

SetPath sets Path field to given value.


### GetFrom

`func (o *PatchItem) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *PatchItem) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *PatchItem) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *PatchItem) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetValue

`func (o *PatchItem) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PatchItem) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PatchItem) SetValue(v interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *PatchItem) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *PatchItem) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *PatchItem) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


