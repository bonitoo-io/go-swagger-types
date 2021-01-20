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

// AuthorizationAllOfLinks struct for AuthorizationAllOfLinks
type AuthorizationAllOfLinks struct {
	// URI of resource.
	Self *string `json:"self,omitempty"`
	// URI of resource.
	User *string `json:"user,omitempty"`
}

// NewAuthorizationAllOfLinks instantiates a new AuthorizationAllOfLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizationAllOfLinks() *AuthorizationAllOfLinks {
	this := AuthorizationAllOfLinks{}
	return &this
}

// NewAuthorizationAllOfLinksWithDefaults instantiates a new AuthorizationAllOfLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizationAllOfLinksWithDefaults() *AuthorizationAllOfLinks {
	this := AuthorizationAllOfLinks{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *AuthorizationAllOfLinks) GetSelf() string {
	if o == nil || o.Self == nil {
		var ret string
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationAllOfLinks) GetSelfOk() (*string, bool) {
	if o == nil || o.Self == nil {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *AuthorizationAllOfLinks) HasSelf() bool {
	if o != nil && o.Self != nil {
		return true
	}

	return false
}

// SetSelf gets a reference to the given string and assigns it to the Self field.
func (o *AuthorizationAllOfLinks) SetSelf(v string) {
	o.Self = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *AuthorizationAllOfLinks) GetUser() string {
	if o == nil || o.User == nil {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationAllOfLinks) GetUserOk() (*string, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *AuthorizationAllOfLinks) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *AuthorizationAllOfLinks) SetUser(v string) {
	o.User = &v
}

func (o AuthorizationAllOfLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Self != nil {
		toSerialize["self"] = o.Self
	}
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	return json.Marshal(toSerialize)
}

type NullableAuthorizationAllOfLinks struct {
	value *AuthorizationAllOfLinks
	isSet bool
}

func (v NullableAuthorizationAllOfLinks) Get() *AuthorizationAllOfLinks {
	return v.value
}

func (v *NullableAuthorizationAllOfLinks) Set(val *AuthorizationAllOfLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizationAllOfLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizationAllOfLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizationAllOfLinks(val *AuthorizationAllOfLinks) *NullableAuthorizationAllOfLinks {
	return &NullableAuthorizationAllOfLinks{value: val, isSet: true}
}

func (v NullableAuthorizationAllOfLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizationAllOfLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


