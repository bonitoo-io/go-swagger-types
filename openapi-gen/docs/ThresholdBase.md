# ThresholdBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Level** | Pointer to [**CheckStatusLevel**](CheckStatusLevel.md) |  | [optional] 
**AllValues** | Pointer to **bool** | If true, only alert if all values meet threshold. | [optional] 

## Methods

### NewThresholdBase

`func NewThresholdBase() *ThresholdBase`

NewThresholdBase instantiates a new ThresholdBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThresholdBaseWithDefaults

`func NewThresholdBaseWithDefaults() *ThresholdBase`

NewThresholdBaseWithDefaults instantiates a new ThresholdBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLevel

`func (o *ThresholdBase) GetLevel() CheckStatusLevel`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *ThresholdBase) GetLevelOk() (*CheckStatusLevel, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *ThresholdBase) SetLevel(v CheckStatusLevel)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *ThresholdBase) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetAllValues

`func (o *ThresholdBase) GetAllValues() bool`

GetAllValues returns the AllValues field if non-nil, zero value otherwise.

### GetAllValuesOk

`func (o *ThresholdBase) GetAllValuesOk() (*bool, bool)`

GetAllValuesOk returns a tuple with the AllValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllValues

`func (o *ThresholdBase) SetAllValues(v bool)`

SetAllValues sets AllValues field to given value.

### HasAllValues

`func (o *ThresholdBase) HasAllValues() bool`

HasAllValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


