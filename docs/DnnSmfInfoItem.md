# DnnSmfInfoItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dnn** | [**AnyOfstringstring**](anyOf&lt;string,string&gt;.md) |  | 
**DnaiList** | Pointer to [**[]AnyOfstringstring**](AnyOfstringstring.md) |  | [optional] 

## Methods

### NewDnnSmfInfoItem

`func NewDnnSmfInfoItem(dnn AnyOfstringstring, ) *DnnSmfInfoItem`

NewDnnSmfInfoItem instantiates a new DnnSmfInfoItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnnSmfInfoItemWithDefaults

`func NewDnnSmfInfoItemWithDefaults() *DnnSmfInfoItem`

NewDnnSmfInfoItemWithDefaults instantiates a new DnnSmfInfoItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnn

`func (o *DnnSmfInfoItem) GetDnn() AnyOfstringstring`

GetDnn returns the Dnn field if non-nil, zero value otherwise.

### GetDnnOk

`func (o *DnnSmfInfoItem) GetDnnOk() (*AnyOfstringstring, bool)`

GetDnnOk returns a tuple with the Dnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnn

`func (o *DnnSmfInfoItem) SetDnn(v AnyOfstringstring)`

SetDnn sets Dnn field to given value.


### GetDnaiList

`func (o *DnnSmfInfoItem) GetDnaiList() []AnyOfstringstring`

GetDnaiList returns the DnaiList field if non-nil, zero value otherwise.

### GetDnaiListOk

`func (o *DnnSmfInfoItem) GetDnaiListOk() (*[]AnyOfstringstring, bool)`

GetDnaiListOk returns a tuple with the DnaiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnaiList

`func (o *DnnSmfInfoItem) SetDnaiList(v []AnyOfstringstring)`

SetDnaiList sets DnaiList field to given value.

### HasDnaiList

`func (o *DnnSmfInfoItem) HasDnaiList() bool`

HasDnaiList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


