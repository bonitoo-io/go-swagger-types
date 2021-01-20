# AuthorizationAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**OrgID** | Pointer to **string** | ID of org that authorization is scoped to. | [optional] 
**Permissions** | Pointer to [**[]Permission**](Permission.md) | List of permissions for an auth.  An auth must have at least one Permission. | [optional] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Token** | Pointer to **string** | Passed via the Authorization Header and Token Authentication type. | [optional] [readonly] 
**UserID** | Pointer to **string** | ID of user that created and owns the token. | [optional] [readonly] 
**User** | Pointer to **string** | Name of user that created and owns the token. | [optional] [readonly] 
**Org** | Pointer to **string** | Name of the org token is scoped to. | [optional] [readonly] 
**Links** | Pointer to [**AuthorizationAllOfLinks**](Authorization_allOf_links.md) |  | [optional] 

## Methods

### NewAuthorizationAllOf

`func NewAuthorizationAllOf() *AuthorizationAllOf`

NewAuthorizationAllOf instantiates a new AuthorizationAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizationAllOfWithDefaults

`func NewAuthorizationAllOfWithDefaults() *AuthorizationAllOf`

NewAuthorizationAllOfWithDefaults instantiates a new AuthorizationAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *AuthorizationAllOf) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AuthorizationAllOf) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AuthorizationAllOf) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AuthorizationAllOf) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AuthorizationAllOf) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AuthorizationAllOf) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AuthorizationAllOf) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AuthorizationAllOf) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetOrgID

`func (o *AuthorizationAllOf) GetOrgID() string`

GetOrgID returns the OrgID field if non-nil, zero value otherwise.

### GetOrgIDOk

`func (o *AuthorizationAllOf) GetOrgIDOk() (*string, bool)`

GetOrgIDOk returns a tuple with the OrgID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgID

`func (o *AuthorizationAllOf) SetOrgID(v string)`

SetOrgID sets OrgID field to given value.

### HasOrgID

`func (o *AuthorizationAllOf) HasOrgID() bool`

HasOrgID returns a boolean if a field has been set.

### GetPermissions

`func (o *AuthorizationAllOf) GetPermissions() []Permission`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *AuthorizationAllOf) GetPermissionsOk() (*[]Permission, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *AuthorizationAllOf) SetPermissions(v []Permission)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *AuthorizationAllOf) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetId

`func (o *AuthorizationAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuthorizationAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuthorizationAllOf) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AuthorizationAllOf) HasId() bool`

HasId returns a boolean if a field has been set.

### GetToken

`func (o *AuthorizationAllOf) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AuthorizationAllOf) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AuthorizationAllOf) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *AuthorizationAllOf) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUserID

`func (o *AuthorizationAllOf) GetUserID() string`

GetUserID returns the UserID field if non-nil, zero value otherwise.

### GetUserIDOk

`func (o *AuthorizationAllOf) GetUserIDOk() (*string, bool)`

GetUserIDOk returns a tuple with the UserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserID

`func (o *AuthorizationAllOf) SetUserID(v string)`

SetUserID sets UserID field to given value.

### HasUserID

`func (o *AuthorizationAllOf) HasUserID() bool`

HasUserID returns a boolean if a field has been set.

### GetUser

`func (o *AuthorizationAllOf) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *AuthorizationAllOf) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *AuthorizationAllOf) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *AuthorizationAllOf) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetOrg

`func (o *AuthorizationAllOf) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *AuthorizationAllOf) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *AuthorizationAllOf) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *AuthorizationAllOf) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetLinks

`func (o *AuthorizationAllOf) GetLinks() AuthorizationAllOfLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AuthorizationAllOf) GetLinksOk() (*AuthorizationAllOfLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AuthorizationAllOf) SetLinks(v AuthorizationAllOfLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AuthorizationAllOf) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


