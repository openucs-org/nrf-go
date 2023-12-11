# NfInstanceInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NrfDiscApiUri** | Pointer to **string** | String providing an URI formatted according to RFC 3986 | [optional] 
**PreferredSearch** | Pointer to [**PreferredSearch**](PreferredSearch.md) |  | [optional] 

## Methods

### NewNfInstanceInfo

`func NewNfInstanceInfo() *NfInstanceInfo`

NewNfInstanceInfo instantiates a new NfInstanceInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNfInstanceInfoWithDefaults

`func NewNfInstanceInfoWithDefaults() *NfInstanceInfo`

NewNfInstanceInfoWithDefaults instantiates a new NfInstanceInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNrfDiscApiUri

`func (o *NfInstanceInfo) GetNrfDiscApiUri() string`

GetNrfDiscApiUri returns the NrfDiscApiUri field if non-nil, zero value otherwise.

### GetNrfDiscApiUriOk

`func (o *NfInstanceInfo) GetNrfDiscApiUriOk() (*string, bool)`

GetNrfDiscApiUriOk returns a tuple with the NrfDiscApiUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNrfDiscApiUri

`func (o *NfInstanceInfo) SetNrfDiscApiUri(v string)`

SetNrfDiscApiUri sets NrfDiscApiUri field to given value.

### HasNrfDiscApiUri

`func (o *NfInstanceInfo) HasNrfDiscApiUri() bool`

HasNrfDiscApiUri returns a boolean if a field has been set.

### GetPreferredSearch

`func (o *NfInstanceInfo) GetPreferredSearch() PreferredSearch`

GetPreferredSearch returns the PreferredSearch field if non-nil, zero value otherwise.

### GetPreferredSearchOk

`func (o *NfInstanceInfo) GetPreferredSearchOk() (*PreferredSearch, bool)`

GetPreferredSearchOk returns a tuple with the PreferredSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredSearch

`func (o *NfInstanceInfo) SetPreferredSearch(v PreferredSearch)`

SetPreferredSearch sets PreferredSearch field to given value.

### HasPreferredSearch

`func (o *NfInstanceInfo) HasPreferredSearch() bool`

HasPreferredSearch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


