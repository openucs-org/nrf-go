# BsfInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DnnList** | Pointer to **[]string** |  | [optional] 
**IpDomainList** | Pointer to **[]string** |  | [optional] 
**Ipv4AddressRanges** | Pointer to [**[]Ipv4AddressRange**](Ipv4AddressRange.md) |  | [optional] 
**Ipv6PrefixRanges** | Pointer to [**[]Ipv6PrefixRange**](Ipv6PrefixRange.md) |  | [optional] 
**RxDiamHost** | Pointer to **string** | string containing an FQDN or realm as defined in RFC 6733. | [optional] 
**RxDiamRealm** | Pointer to **string** | string containing an FQDN or realm as defined in RFC 6733. | [optional] 

## Methods

### NewBsfInfo

`func NewBsfInfo() *BsfInfo`

NewBsfInfo instantiates a new BsfInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBsfInfoWithDefaults

`func NewBsfInfoWithDefaults() *BsfInfo`

NewBsfInfoWithDefaults instantiates a new BsfInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnnList

`func (o *BsfInfo) GetDnnList() []string`

GetDnnList returns the DnnList field if non-nil, zero value otherwise.

### GetDnnListOk

`func (o *BsfInfo) GetDnnListOk() (*[]string, bool)`

GetDnnListOk returns a tuple with the DnnList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnnList

`func (o *BsfInfo) SetDnnList(v []string)`

SetDnnList sets DnnList field to given value.

### HasDnnList

`func (o *BsfInfo) HasDnnList() bool`

HasDnnList returns a boolean if a field has been set.

### GetIpDomainList

`func (o *BsfInfo) GetIpDomainList() []string`

GetIpDomainList returns the IpDomainList field if non-nil, zero value otherwise.

### GetIpDomainListOk

`func (o *BsfInfo) GetIpDomainListOk() (*[]string, bool)`

GetIpDomainListOk returns a tuple with the IpDomainList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpDomainList

`func (o *BsfInfo) SetIpDomainList(v []string)`

SetIpDomainList sets IpDomainList field to given value.

### HasIpDomainList

`func (o *BsfInfo) HasIpDomainList() bool`

HasIpDomainList returns a boolean if a field has been set.

### GetIpv4AddressRanges

`func (o *BsfInfo) GetIpv4AddressRanges() []Ipv4AddressRange`

GetIpv4AddressRanges returns the Ipv4AddressRanges field if non-nil, zero value otherwise.

### GetIpv4AddressRangesOk

`func (o *BsfInfo) GetIpv4AddressRangesOk() (*[]Ipv4AddressRange, bool)`

GetIpv4AddressRangesOk returns a tuple with the Ipv4AddressRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4AddressRanges

`func (o *BsfInfo) SetIpv4AddressRanges(v []Ipv4AddressRange)`

SetIpv4AddressRanges sets Ipv4AddressRanges field to given value.

### HasIpv4AddressRanges

`func (o *BsfInfo) HasIpv4AddressRanges() bool`

HasIpv4AddressRanges returns a boolean if a field has been set.

### GetIpv6PrefixRanges

`func (o *BsfInfo) GetIpv6PrefixRanges() []Ipv6PrefixRange`

GetIpv6PrefixRanges returns the Ipv6PrefixRanges field if non-nil, zero value otherwise.

### GetIpv6PrefixRangesOk

`func (o *BsfInfo) GetIpv6PrefixRangesOk() (*[]Ipv6PrefixRange, bool)`

GetIpv6PrefixRangesOk returns a tuple with the Ipv6PrefixRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6PrefixRanges

`func (o *BsfInfo) SetIpv6PrefixRanges(v []Ipv6PrefixRange)`

SetIpv6PrefixRanges sets Ipv6PrefixRanges field to given value.

### HasIpv6PrefixRanges

`func (o *BsfInfo) HasIpv6PrefixRanges() bool`

HasIpv6PrefixRanges returns a boolean if a field has been set.

### GetRxDiamHost

`func (o *BsfInfo) GetRxDiamHost() string`

GetRxDiamHost returns the RxDiamHost field if non-nil, zero value otherwise.

### GetRxDiamHostOk

`func (o *BsfInfo) GetRxDiamHostOk() (*string, bool)`

GetRxDiamHostOk returns a tuple with the RxDiamHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxDiamHost

`func (o *BsfInfo) SetRxDiamHost(v string)`

SetRxDiamHost sets RxDiamHost field to given value.

### HasRxDiamHost

`func (o *BsfInfo) HasRxDiamHost() bool`

HasRxDiamHost returns a boolean if a field has been set.

### GetRxDiamRealm

`func (o *BsfInfo) GetRxDiamRealm() string`

GetRxDiamRealm returns the RxDiamRealm field if non-nil, zero value otherwise.

### GetRxDiamRealmOk

`func (o *BsfInfo) GetRxDiamRealmOk() (*string, bool)`

GetRxDiamRealmOk returns a tuple with the RxDiamRealm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxDiamRealm

`func (o *BsfInfo) SetRxDiamRealm(v string)`

SetRxDiamRealm sets RxDiamRealm field to given value.

### HasRxDiamRealm

`func (o *BsfInfo) HasRxDiamRealm() bool`

HasRxDiamRealm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


