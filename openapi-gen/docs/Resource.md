# Resource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | Pointer to **NullableString** | If ID is set that is a permission for a specific resource. if it is not set it is a permission for all resources of that resource type. | [optional] 
**Name** | Pointer to **NullableString** | Optional name of the resource if the resource has a name field. | [optional] 
**OrgID** | Pointer to **NullableString** | If orgID is set that is a permission for all resources owned my that org. if it is not set it is a permission for all resources of that resource type. | [optional] 
**Org** | Pointer to **NullableString** | Optional name of the organization of the organization with orgID. | [optional] 

## Methods

### NewResource

`func NewResource(type_ string, ) *Resource`

NewResource instantiates a new Resource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceWithDefaults

`func NewResourceWithDefaults() *Resource`

NewResourceWithDefaults instantiates a new Resource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Resource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Resource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Resource) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *Resource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Resource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Resource) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Resource) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *Resource) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *Resource) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *Resource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Resource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Resource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Resource) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Resource) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Resource) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOrgID

`func (o *Resource) GetOrgID() string`

GetOrgID returns the OrgID field if non-nil, zero value otherwise.

### GetOrgIDOk

`func (o *Resource) GetOrgIDOk() (*string, bool)`

GetOrgIDOk returns a tuple with the OrgID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgID

`func (o *Resource) SetOrgID(v string)`

SetOrgID sets OrgID field to given value.

### HasOrgID

`func (o *Resource) HasOrgID() bool`

HasOrgID returns a boolean if a field has been set.

### SetOrgIDNil

`func (o *Resource) SetOrgIDNil(b bool)`

 SetOrgIDNil sets the value for OrgID to be an explicit nil

### UnsetOrgID
`func (o *Resource) UnsetOrgID()`

UnsetOrgID ensures that no value is present for OrgID, not even an explicit nil
### GetOrg

`func (o *Resource) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *Resource) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *Resource) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *Resource) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### SetOrgNil

`func (o *Resource) SetOrgNil(b bool)`

 SetOrgNil sets the value for Org to be an explicit nil

### UnsetOrg
`func (o *Resource) UnsetOrg()`

UnsetOrg ensures that no value is present for Org, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


