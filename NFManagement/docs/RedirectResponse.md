# RedirectResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cause** | Pointer to **string** |  | [optional] 
**TargetScp** | Pointer to **string** | String providing an URI formatted according to RFC 3986 | [optional] 

## Methods

### NewRedirectResponse

`func NewRedirectResponse() *RedirectResponse`

NewRedirectResponse instantiates a new RedirectResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRedirectResponseWithDefaults

`func NewRedirectResponseWithDefaults() *RedirectResponse`

NewRedirectResponseWithDefaults instantiates a new RedirectResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCause

`func (o *RedirectResponse) GetCause() string`

GetCause returns the Cause field if non-nil, zero value otherwise.

### GetCauseOk

`func (o *RedirectResponse) GetCauseOk() (*string, bool)`

GetCauseOk returns a tuple with the Cause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCause

`func (o *RedirectResponse) SetCause(v string)`

SetCause sets Cause field to given value.

### HasCause

`func (o *RedirectResponse) HasCause() bool`

HasCause returns a boolean if a field has been set.

### GetTargetScp

`func (o *RedirectResponse) GetTargetScp() string`

GetTargetScp returns the TargetScp field if non-nil, zero value otherwise.

### GetTargetScpOk

`func (o *RedirectResponse) GetTargetScpOk() (*string, bool)`

GetTargetScpOk returns a tuple with the TargetScp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetScp

`func (o *RedirectResponse) SetTargetScp(v string)`

SetTargetScp sets TargetScp field to given value.

### HasTargetScp

`func (o *RedirectResponse) HasTargetScp() bool`

HasTargetScp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


