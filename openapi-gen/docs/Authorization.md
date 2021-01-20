# Authorization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | If inactive the token is inactive and requests using the token will be rejected. | [optional] [default to "active"]
**Description** | Pointer to **string** | A description of the token. | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**OrgID** | **string** | ID of org that authorization is scoped to. | 
**Permissions** | [**[]Permission**](Permission.md) | List of permissions for an auth.  An auth must have at least one Permission. | 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Token** | Pointer to **string** | Passed via the Authorization Header and Token Authentication type. | [optional] [readonly] 
**UserID** | Pointer to **string** | ID of user that created and owns the token. | [optional] [readonly] 
**User** | Pointer to **string** | Name of user that created and owns the token. | [optional] [readonly] 
**Org** | Pointer to **string** | Name of the org token is scoped to. | [optional] [readonly] 
**Links** | Pointer to [**AuthorizationAllOfLinks**](Authorization_allOf_links.md) |  | [optional] 

## Methods

### NewAuthorization

`func NewAuthorization(orgID string, permissions []Permission, ) *Authorization`

NewAuthorization instantiates a new Authorization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizationWithDefaults

`func NewAuthorizationWithDefaults() *Authorization`

NewAuthorizationWithDefaults instantiates a new Authorization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *Authorization) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Authorization) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Authorization) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Authorization) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDescription

`func (o *Authorization) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Authorization) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Authorization) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Authorization) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Authorization) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Authorization) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Authorization) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Authorization) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Authorization) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Authorization) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Authorization) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Authorization) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetOrgID

`func (o *Authorization) GetOrgID() string`

GetOrgID returns the OrgID field if non-nil, zero value otherwise.

### GetOrgIDOk

`func (o *Authorization) GetOrgIDOk() (*string, bool)`

GetOrgIDOk returns a tuple with the OrgID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgID

`func (o *Authorization) SetOrgID(v string)`

SetOrgID sets OrgID field to given value.


### GetPermissions

`func (o *Authorization) GetPermissions() []Permission`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *Authorization) GetPermissionsOk() (*[]Permission, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *Authorization) SetPermissions(v []Permission)`

SetPermissions sets Permissions field to given value.


### GetId

`func (o *Authorization) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Authorization) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Authorization) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Authorization) HasId() bool`

HasId returns a boolean if a field has been set.

### GetToken

`func (o *Authorization) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *Authorization) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *Authorization) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *Authorization) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUserID

`func (o *Authorization) GetUserID() string`

GetUserID returns the UserID field if non-nil, zero value otherwise.

### GetUserIDOk

`func (o *Authorization) GetUserIDOk() (*string, bool)`

GetUserIDOk returns a tuple with the UserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserID

`func (o *Authorization) SetUserID(v string)`

SetUserID sets UserID field to given value.

### HasUserID

`func (o *Authorization) HasUserID() bool`

HasUserID returns a boolean if a field has been set.

### GetUser

`func (o *Authorization) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *Authorization) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *Authorization) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *Authorization) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetOrg

`func (o *Authorization) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *Authorization) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *Authorization) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *Authorization) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetLinks

`func (o *Authorization) GetLinks() AuthorizationAllOfLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Authorization) GetLinksOk() (*AuthorizationAllOfLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Authorization) SetLinks(v AuthorizationAllOfLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Authorization) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


