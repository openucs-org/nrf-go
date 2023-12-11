# \AccessTokenRequestApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccessTokenRequest**](AccessTokenRequestApi.md#AccessTokenRequest) | **Post** /oauth2/token | Access Token Request



## AccessTokenRequest

> AccessTokenRsp AccessTokenRequest(ctx).GrantType(grantType).NfInstanceId(nfInstanceId).Scope(scope).ContentEncoding(contentEncoding).AcceptEncoding(acceptEncoding).NfType(nfType).TargetNfType(targetNfType).TargetNfInstanceId(targetNfInstanceId).RequesterPlmn(requesterPlmn).RequesterPlmnList(requesterPlmnList).RequesterSnssaiList(requesterSnssaiList).RequesterFqdn(requesterFqdn).RequesterSnpnList(requesterSnpnList).TargetPlmn(targetPlmn).TargetSnssaiList(targetSnssaiList).TargetNsiList(targetNsiList).TargetNfSetId(targetNfSetId).TargetNfServiceSetId(targetNfServiceSetId).HnrfAccessTokenUri(hnrfAccessTokenUri).Execute()

Access Token Request

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    grantType := "grantType_example" // string | 
    nfInstanceId := TODO // string | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.
    scope := "scope_example" // string | 
    contentEncoding := "contentEncoding_example" // string | Content-Encoding, described in IETF RFC 7231 (optional)
    acceptEncoding := "acceptEncoding_example" // string | Accept-Encoding, described in IETF RFC 7231 (optional)
    nfType := *openapiclient.NewNFType() // NFType |  (optional)
    targetNfType := *openapiclient.NewNFType() // NFType |  (optional)
    targetNfInstanceId := TODO // string | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122. (optional)
    requesterPlmn := *openapiclient.NewPlmnId("Mcc_example", "Mnc_example") // PlmnId |  (optional)
    requesterPlmnList := []openapiclient.PlmnId{*openapiclient.NewPlmnId("Mcc_example", "Mnc_example")} // []PlmnId |  (optional)
    requesterSnssaiList := []openapiclient.Snssai{*openapiclient.NewSnssai(int32(123))} // []Snssai |  (optional)
    requesterFqdn := "requesterFqdn_example" // string | Fully Qualified Domain Name (optional)
    requesterSnpnList := []openapiclient.PlmnIdNid{*openapiclient.NewPlmnIdNid("Mcc_example", "Mnc_example")} // []PlmnIdNid |  (optional)
    targetPlmn :=  // PlmnId |  (optional)
    targetSnssaiList := []openapiclient.Snssai{*openapiclient.NewSnssai(int32(123))} // []Snssai |  (optional)
    targetNsiList := []string{"Inner_example"} // []string |  (optional)
    targetNfSetId := "targetNfSetId_example" // string | NF Set Identifier (see clause 28.12 of 3GPP TS 23.003), formatted as the following string \\\" set<Set ID>.<nftype>set.5gc.mnc<MNC>.mcc<MCC>\\\", or \\\"set<SetID>. <NFType>set.5gc.nid<NID>.mnc<MNC>.mcc<MCC>\\\" with <MCC> encoded as defined in clause 5.4.2 (\\\"Mcc\\\" data type definition) <MNC> encoded as defined in clause 5.4.2 (\\\"Mnc\\\" data type definition) <NFType> encoded as a value defined in Table 6.1.6.3.3-1 of 3GPP TS 29.510 but with lower case characters <Set ID> encoded as a string of characters consisting of alphabetic characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that shall end with either an alphabetic character or a digit. (optional)
    targetNfServiceSetId := "targetNfServiceSetId_example" // string | NF Service Set Identifier (see clause 28.12 of 3GPP TS 23.003) formatted as the following string  \\\" set<Set ID>.sn<Service Name>.nfi<NF Instance ID>.5gc.mnc<MNC>.mcc<MCC>\\\">\\\", or \\\"set<SetID>.sn<ServiceName>.nfi<NFInstanceID>.5gc.nid<NID>.mnc<MNC>.mcc<MCC>\\\" with <MCC> encoded as defined in clause 5.4.2 (\\\"Mcc\\\" data type definition)  <MNC> encoded as defined in clause 5.4.2 (\\\"Mnc\\\" data type definition)  <NID> encoded as defined in clause 5.4.2 (\\\"Nid\\\" data type definition) <NFInstanceId> encoded as defined in clause 5.3.2 <ServiceName> encoded as defined in 3GPP TS 29.510 <Set ID> encoded as a string of characters consisting of alphabetic characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that shall end with either an alphabetic character or a digit. (optional)
    hnrfAccessTokenUri := "hnrfAccessTokenUri_example" // string | String providing an URI formatted according to RFC 3986 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccessTokenRequestApi.AccessTokenRequest(context.Background()).GrantType(grantType).NfInstanceId(nfInstanceId).Scope(scope).ContentEncoding(contentEncoding).AcceptEncoding(acceptEncoding).NfType(nfType).TargetNfType(targetNfType).TargetNfInstanceId(targetNfInstanceId).RequesterPlmn(requesterPlmn).RequesterPlmnList(requesterPlmnList).RequesterSnssaiList(requesterSnssaiList).RequesterFqdn(requesterFqdn).RequesterSnpnList(requesterSnpnList).TargetPlmn(targetPlmn).TargetSnssaiList(targetSnssaiList).TargetNsiList(targetNsiList).TargetNfSetId(targetNfSetId).TargetNfServiceSetId(targetNfServiceSetId).HnrfAccessTokenUri(hnrfAccessTokenUri).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessTokenRequestApi.AccessTokenRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccessTokenRequest`: AccessTokenRsp
    fmt.Fprintf(os.Stdout, "Response from `AccessTokenRequestApi.AccessTokenRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAccessTokenRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **grantType** | **string** |  | 
 **nfInstanceId** | [**string**](string.md) | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122. | 
 **scope** | **string** |  | 
 **contentEncoding** | **string** | Content-Encoding, described in IETF RFC 7231 | 
 **acceptEncoding** | **string** | Accept-Encoding, described in IETF RFC 7231 | 
 **nfType** | [**NFType**](NFType.md) |  | 
 **targetNfType** | [**NFType**](NFType.md) |  | 
 **targetNfInstanceId** | [**string**](string.md) | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122. | 
 **requesterPlmn** | [**PlmnId**](PlmnId.md) |  | 
 **requesterPlmnList** | [**[]PlmnId**](PlmnId.md) |  | 
 **requesterSnssaiList** | [**[]Snssai**](Snssai.md) |  | 
 **requesterFqdn** | **string** | Fully Qualified Domain Name | 
 **requesterSnpnList** | [**[]PlmnIdNid**](PlmnIdNid.md) |  | 
 **targetPlmn** | [**PlmnId**](PlmnId.md) |  | 
 **targetSnssaiList** | [**[]Snssai**](Snssai.md) |  | 
 **targetNsiList** | **[]string** |  | 
 **targetNfSetId** | **string** | NF Set Identifier (see clause 28.12 of 3GPP TS 23.003), formatted as the following string \\\&quot; set&lt;Set ID&gt;.&lt;nftype&gt;set.5gc.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\\\&quot;, or \\\&quot;set&lt;SetID&gt;. &lt;NFType&gt;set.5gc.nid&lt;NID&gt;.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\\\&quot; with &lt;MCC&gt; encoded as defined in clause 5.4.2 (\\\&quot;Mcc\\\&quot; data type definition) &lt;MNC&gt; encoded as defined in clause 5.4.2 (\\\&quot;Mnc\\\&quot; data type definition) &lt;NFType&gt; encoded as a value defined in Table 6.1.6.3.3-1 of 3GPP TS 29.510 but with lower case characters &lt;Set ID&gt; encoded as a string of characters consisting of alphabetic characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that shall end with either an alphabetic character or a digit. | 
 **targetNfServiceSetId** | **string** | NF Service Set Identifier (see clause 28.12 of 3GPP TS 23.003) formatted as the following string  \\\&quot; set&lt;Set ID&gt;.sn&lt;Service Name&gt;.nfi&lt;NF Instance ID&gt;.5gc.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\\\&quot;&gt;\\\&quot;, or \\\&quot;set&lt;SetID&gt;.sn&lt;ServiceName&gt;.nfi&lt;NFInstanceID&gt;.5gc.nid&lt;NID&gt;.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\\\&quot; with &lt;MCC&gt; encoded as defined in clause 5.4.2 (\\\&quot;Mcc\\\&quot; data type definition)  &lt;MNC&gt; encoded as defined in clause 5.4.2 (\\\&quot;Mnc\\\&quot; data type definition)  &lt;NID&gt; encoded as defined in clause 5.4.2 (\\\&quot;Nid\\\&quot; data type definition) &lt;NFInstanceId&gt; encoded as defined in clause 5.3.2 &lt;ServiceName&gt; encoded as defined in 3GPP TS 29.510 &lt;Set ID&gt; encoded as a string of characters consisting of alphabetic characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that shall end with either an alphabetic character or a digit. | 
 **hnrfAccessTokenUri** | **string** | String providing an URI formatted according to RFC 3986 | 

### Return type

[**AccessTokenRsp**](AccessTokenRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

