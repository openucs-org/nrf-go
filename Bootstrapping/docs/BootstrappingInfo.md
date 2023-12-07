# BootstrappingInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**Status**](Status.md) |  | [optional] 
**Links** | [**map[string]LinksValueSchema**](LinksValueSchema.md) | Map of link objects where the keys are the link relations defined in 3GPP TS 29.510 clause 6.4.6.3.3 | 
**NrfFeatures** | Pointer to **map[string]string** | Map of features supported by the NRF, where the keys are the NRF services as defined in 3GPP TS 29.510 clause 6.1.6.3.11 | [optional] 

## Methods

### NewBootstrappingInfo

`func NewBootstrappingInfo(links map[string]LinksValueSchema, ) *BootstrappingInfo`

NewBootstrappingInfo instantiates a new BootstrappingInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBootstrappingInfoWithDefaults

`func NewBootstrappingInfoWithDefaults() *BootstrappingInfo`

NewBootstrappingInfoWithDefaults instantiates a new BootstrappingInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *BootstrappingInfo) GetStatus() Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BootstrappingInfo) GetStatusOk() (*Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BootstrappingInfo) SetStatus(v Status)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BootstrappingInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLinks

`func (o *BootstrappingInfo) GetLinks() map[string]LinksValueSchema`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BootstrappingInfo) GetLinksOk() (*map[string]LinksValueSchema, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BootstrappingInfo) SetLinks(v map[string]LinksValueSchema)`

SetLinks sets Links field to given value.


### GetNrfFeatures

`func (o *BootstrappingInfo) GetNrfFeatures() map[string]string`

GetNrfFeatures returns the NrfFeatures field if non-nil, zero value otherwise.

### GetNrfFeaturesOk

`func (o *BootstrappingInfo) GetNrfFeaturesOk() (*map[string]string, bool)`

GetNrfFeaturesOk returns a tuple with the NrfFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNrfFeatures

`func (o *BootstrappingInfo) SetNrfFeatures(v map[string]string)`

SetNrfFeatures sets NrfFeatures field to given value.

### HasNrfFeatures

`func (o *BootstrappingInfo) HasNrfFeatures() bool`

HasNrfFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


