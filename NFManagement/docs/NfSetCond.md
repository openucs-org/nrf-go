# NfSetCond

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NfSetId** | **string** | NF Set Identifier (see clause 28.12 of 3GPP TS 23.003), formatted as the following string \&quot; set&lt;Set ID&gt;.&lt;nftype&gt;set.5gc.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot;, or \&quot;set&lt;SetID&gt;. &lt;NFType&gt;set.5gc.nid&lt;NID&gt;.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot; with &lt;MCC&gt; encoded as defined in clause 5.4.2 (\&quot;Mcc\&quot; data type definition) &lt;MNC&gt; encoded as defined in clause 5.4.2 (\&quot;Mnc\&quot; data type definition) &lt;NFType&gt; encoded as a value defined in Table 6.1.6.3.3-1 of 3GPP TS 29.510 but with lower case characters &lt;Set ID&gt; encoded as a string of characters consisting of alphabetic characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that shall end with either an alphabetic character or a digit. | 

## Methods

### NewNfSetCond

`func NewNfSetCond(nfSetId string, ) *NfSetCond`

NewNfSetCond instantiates a new NfSetCond object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNfSetCondWithDefaults

`func NewNfSetCondWithDefaults() *NfSetCond`

NewNfSetCondWithDefaults instantiates a new NfSetCond object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNfSetId

`func (o *NfSetCond) GetNfSetId() string`

GetNfSetId returns the NfSetId field if non-nil, zero value otherwise.

### GetNfSetIdOk

`func (o *NfSetCond) GetNfSetIdOk() (*string, bool)`

GetNfSetIdOk returns a tuple with the NfSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfSetId

`func (o *NfSetCond) SetNfSetId(v string)`

SetNfSetId sets NfSetId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


