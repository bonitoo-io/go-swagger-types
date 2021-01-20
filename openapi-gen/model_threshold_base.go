/*
 * Influx API Service
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ThresholdBase struct for ThresholdBase
type ThresholdBase struct {
	Level *CheckStatusLevel `json:"level,omitempty"`
	// If true, only alert if all values meet threshold.
	AllValues *bool `json:"allValues,omitempty"`
}

// NewThresholdBase instantiates a new ThresholdBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThresholdBase() *ThresholdBase {
	this := ThresholdBase{}
	return &this
}

// NewThresholdBaseWithDefaults instantiates a new ThresholdBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThresholdBaseWithDefaults() *ThresholdBase {
	this := ThresholdBase{}
	return &this
}

// GetLevel returns the Level field value if set, zero value otherwise.
func (o *ThresholdBase) GetLevel() CheckStatusLevel {
	if o == nil || o.Level == nil {
		var ret CheckStatusLevel
		return ret
	}
	return *o.Level
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThresholdBase) GetLevelOk() (*CheckStatusLevel, bool) {
	if o == nil || o.Level == nil {
		return nil, false
	}
	return o.Level, true
}

// HasLevel returns a boolean if a field has been set.
func (o *ThresholdBase) HasLevel() bool {
	if o != nil && o.Level != nil {
		return true
	}

	return false
}

// SetLevel gets a reference to the given CheckStatusLevel and assigns it to the Level field.
func (o *ThresholdBase) SetLevel(v CheckStatusLevel) {
	o.Level = &v
}

// GetAllValues returns the AllValues field value if set, zero value otherwise.
func (o *ThresholdBase) GetAllValues() bool {
	if o == nil || o.AllValues == nil {
		var ret bool
		return ret
	}
	return *o.AllValues
}

// GetAllValuesOk returns a tuple with the AllValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThresholdBase) GetAllValuesOk() (*bool, bool) {
	if o == nil || o.AllValues == nil {
		return nil, false
	}
	return o.AllValues, true
}

// HasAllValues returns a boolean if a field has been set.
func (o *ThresholdBase) HasAllValues() bool {
	if o != nil && o.AllValues != nil {
		return true
	}

	return false
}

// SetAllValues gets a reference to the given bool and assigns it to the AllValues field.
func (o *ThresholdBase) SetAllValues(v bool) {
	o.AllValues = &v
}

func (o ThresholdBase) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Level != nil {
		toSerialize["level"] = o.Level
	}
	if o.AllValues != nil {
		toSerialize["allValues"] = o.AllValues
	}
	return json.Marshal(toSerialize)
}

type NullableThresholdBase struct {
	value *ThresholdBase
	isSet bool
}

func (v NullableThresholdBase) Get() *ThresholdBase {
	return v.value
}

func (v *NullableThresholdBase) Set(val *ThresholdBase) {
	v.value = val
	v.isSet = true
}

func (v NullableThresholdBase) IsSet() bool {
	return v.isSet
}

func (v *NullableThresholdBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThresholdBase(val *ThresholdBase) *NullableThresholdBase {
	return &NullableThresholdBase{value: val, isSet: true}
}

func (v NullableThresholdBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThresholdBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


