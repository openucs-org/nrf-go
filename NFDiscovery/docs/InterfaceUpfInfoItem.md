# InterfaceUpfInfoItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InterfaceType** | [**UPInterfaceType**](UPInterfaceType.md) |  | 
**Ipv4EndpointAddresses** | Pointer to **[]string** |  | [optional] 
**Ipv6EndpointAddresses** | Pointer to [**[]Ipv6Addr**](Ipv6Addr.md) |  | [optional] 
**EndpointFqdn** | Pointer to **string** | Fully Qualified Domain Name | [optional] 
**NetworkInstance** | Pointer to **string** |  | [optional] 

## Methods

### NewInterfaceUpfInfoItem

`func NewInterfaceUpfInfoItem(interfaceType UPInterfaceType, ) *InterfaceUpfInfoItem`

NewInterfaceUpfInfoItem instantiates a new InterfaceUpfInfoItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInterfaceUpfInfoItemWithDefaults

`func NewInterfaceUpfInfoItemWithDefaults() *InterfaceUpfInfoItem`

NewInterfaceUpfInfoItemWithDefaults instantiates a new InterfaceUpfInfoItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterfaceType

`func (o *InterfaceUpfInfoItem) GetInterfaceType() UPInterfaceType`

GetInterfaceType returns the InterfaceType field if non-nil, zero value otherwise.

### GetInterfaceTypeOk

`func (o *InterfaceUpfInfoItem) GetInterfaceTypeOk() (*UPInterfaceType, bool)`

GetInterfaceTypeOk returns a tuple with the InterfaceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceType

`func (o *InterfaceUpfInfoItem) SetInterfaceType(v UPInterfaceType)`

SetInterfaceType sets InterfaceType field to given value.


### GetIpv4EndpointAddresses

`func (o *InterfaceUpfInfoItem) GetIpv4EndpointAddresses() []string`

GetIpv4EndpointAddresses returns the Ipv4EndpointAddresses field if non-nil, zero value otherwise.

### GetIpv4EndpointAddressesOk

`func (o *InterfaceUpfInfoItem) GetIpv4EndpointAddressesOk() (*[]string, bool)`

GetIpv4EndpointAddressesOk returns a tuple with the Ipv4EndpointAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4EndpointAddresses

`func (o *InterfaceUpfInfoItem) SetIpv4EndpointAddresses(v []string)`

SetIpv4EndpointAddresses sets Ipv4EndpointAddresses field to given value.

### HasIpv4EndpointAddresses

`func (o *InterfaceUpfInfoItem) HasIpv4EndpointAddresses() bool`

HasIpv4EndpointAddresses returns a boolean if a field has been set.

### GetIpv6EndpointAddresses

`func (o *InterfaceUpfInfoItem) GetIpv6EndpointAddresses() []Ipv6Addr`

GetIpv6EndpointAddresses returns the Ipv6EndpointAddresses field if non-nil, zero value otherwise.

### GetIpv6EndpointAddressesOk

`func (o *InterfaceUpfInfoItem) GetIpv6EndpointAddressesOk() (*[]Ipv6Addr, bool)`

GetIpv6EndpointAddressesOk returns a tuple with the Ipv6EndpointAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6EndpointAddresses

`func (o *InterfaceUpfInfoItem) SetIpv6EndpointAddresses(v []Ipv6Addr)`

SetIpv6EndpointAddresses sets Ipv6EndpointAddresses field to given value.

### HasIpv6EndpointAddresses

`func (o *InterfaceUpfInfoItem) HasIpv6EndpointAddresses() bool`

HasIpv6EndpointAddresses returns a boolean if a field has been set.

### GetEndpointFqdn

`func (o *InterfaceUpfInfoItem) GetEndpointFqdn() string`

GetEndpointFqdn returns the EndpointFqdn field if non-nil, zero value otherwise.

### GetEndpointFqdnOk

`func (o *InterfaceUpfInfoItem) GetEndpointFqdnOk() (*string, bool)`

GetEndpointFqdnOk returns a tuple with the EndpointFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointFqdn

`func (o *InterfaceUpfInfoItem) SetEndpointFqdn(v string)`

SetEndpointFqdn sets EndpointFqdn field to given value.

### HasEndpointFqdn

`func (o *InterfaceUpfInfoItem) HasEndpointFqdn() bool`

HasEndpointFqdn returns a boolean if a field has been set.

### GetNetworkInstance

`func (o *InterfaceUpfInfoItem) GetNetworkInstance() string`

GetNetworkInstance returns the NetworkInstance field if non-nil, zero value otherwise.

### GetNetworkInstanceOk

`func (o *InterfaceUpfInfoItem) GetNetworkInstanceOk() (*string, bool)`

GetNetworkInstanceOk returns a tuple with the NetworkInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInstance

`func (o *InterfaceUpfInfoItem) SetNetworkInstance(v string)`

SetNetworkInstance sets NetworkInstance field to given value.

### HasNetworkInstance

`func (o *InterfaceUpfInfoItem) HasNetworkInstance() bool`

HasNetworkInstance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


