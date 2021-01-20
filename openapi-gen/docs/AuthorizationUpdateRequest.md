# AuthorizationUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | If inactive the token is inactive and requests using the token will be rejected. | [optional] [default to "active"]
**Description** | Pointer to **string** | A description of the token. | [optional] 

## Methods

### NewAuthorizationUpdateRequest

`func NewAuthorizationUpdateRequest() *AuthorizationUpdateRequest`

NewAuthorizationUpdateRequest instantiates a new AuthorizationUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizationUpdateRequestWithDefaults

`func NewAuthorizationUpdateRequestWithDefaults() *AuthorizationUpdateRequest`

NewAuthorizationUpdateRequestWithDefaults instantiates a new AuthorizationUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *AuthorizationUpdateRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AuthorizationUpdateRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AuthorizationUpdateRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AuthorizationUpdateRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDescription

`func (o *AuthorizationUpdateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AuthorizationUpdateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AuthorizationUpdateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AuthorizationUpdateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


