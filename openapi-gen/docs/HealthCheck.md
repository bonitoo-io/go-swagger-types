# HealthCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Message** | Pointer to **string** |  | [optional] 
**Checks** | Pointer to [**[]HealthCheck**](HealthCheck.md) |  | [optional] 
**Status** | **string** |  | 

## Methods

### NewHealthCheck

`func NewHealthCheck(name string, status string, ) *HealthCheck`

NewHealthCheck instantiates a new HealthCheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealthCheckWithDefaults

`func NewHealthCheckWithDefaults() *HealthCheck`

NewHealthCheckWithDefaults instantiates a new HealthCheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *HealthCheck) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HealthCheck) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HealthCheck) SetName(v string)`

SetName sets Name field to given value.


### GetMessage

`func (o *HealthCheck) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *HealthCheck) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *HealthCheck) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *HealthCheck) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetChecks

`func (o *HealthCheck) GetChecks() []HealthCheck`

GetChecks returns the Checks field if non-nil, zero value otherwise.

### GetChecksOk

`func (o *HealthCheck) GetChecksOk() (*[]HealthCheck, bool)`

GetChecksOk returns a tuple with the Checks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecks

`func (o *HealthCheck) SetChecks(v []HealthCheck)`

SetChecks sets Checks field to given value.

### HasChecks

`func (o *HealthCheck) HasChecks() bool`

HasChecks returns a boolean if a field has been set.

### GetStatus

`func (o *HealthCheck) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HealthCheck) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HealthCheck) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


