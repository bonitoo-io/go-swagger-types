# Authorizations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**Links**](Links.md) |  | [optional] 
**Authorizations** | Pointer to [**[]Authorization**](Authorization.md) |  | [optional] 

## Methods

### NewAuthorizations

`func NewAuthorizations() *Authorizations`

NewAuthorizations instantiates a new Authorizations object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizationsWithDefaults

`func NewAuthorizationsWithDefaults() *Authorizations`

NewAuthorizationsWithDefaults instantiates a new Authorizations object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *Authorizations) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Authorizations) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Authorizations) SetLinks(v Links)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Authorizations) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetAuthorizations

`func (o *Authorizations) GetAuthorizations() []Authorization`

GetAuthorizations returns the Authorizations field if non-nil, zero value otherwise.

### GetAuthorizationsOk

`func (o *Authorizations) GetAuthorizationsOk() (*[]Authorization, bool)`

GetAuthorizationsOk returns a tuple with the Authorizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizations

`func (o *Authorizations) SetAuthorizations(v []Authorization)`

SetAuthorizations sets Authorizations field to given value.

### HasAuthorizations

`func (o *Authorizations) HasAuthorizations() bool`

HasAuthorizations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


