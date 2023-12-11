# MbsAreaSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AreaSessionId** | **int32** | Integer where the allowed values correspond to the value range of an unsigned 16-bit integer. | 
**MbsServiceArea** | [**MbsServiceArea**](MbsServiceArea.md) |  | 

## Methods

### NewMbsAreaSession

`func NewMbsAreaSession(areaSessionId int32, mbsServiceArea MbsServiceArea, ) *MbsAreaSession`

NewMbsAreaSession instantiates a new MbsAreaSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMbsAreaSessionWithDefaults

`func NewMbsAreaSessionWithDefaults() *MbsAreaSession`

NewMbsAreaSessionWithDefaults instantiates a new MbsAreaSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAreaSessionId

`func (o *MbsAreaSession) GetAreaSessionId() int32`

GetAreaSessionId returns the AreaSessionId field if non-nil, zero value otherwise.

### GetAreaSessionIdOk

`func (o *MbsAreaSession) GetAreaSessionIdOk() (*int32, bool)`

GetAreaSessionIdOk returns a tuple with the AreaSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreaSessionId

`func (o *MbsAreaSession) SetAreaSessionId(v int32)`

SetAreaSessionId sets AreaSessionId field to given value.


### GetMbsServiceArea

`func (o *MbsAreaSession) GetMbsServiceArea() MbsServiceArea`

GetMbsServiceArea returns the MbsServiceArea field if non-nil, zero value otherwise.

### GetMbsServiceAreaOk

`func (o *MbsAreaSession) GetMbsServiceAreaOk() (*MbsServiceArea, bool)`

GetMbsServiceAreaOk returns a tuple with the MbsServiceArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbsServiceArea

`func (o *MbsAreaSession) SetMbsServiceArea(v MbsServiceArea)`

SetMbsServiceArea sets MbsServiceArea field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


