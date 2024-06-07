/*
maestro Service API

maestro Service API

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// checks if the ResourceBundleAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceBundleAllOf{}

// ResourceBundleAllOf struct for ResourceBundleAllOf
type ResourceBundleAllOf struct {
	Name            *string                  `json:"name,omitempty"`
	ConsumerName    *string                  `json:"consumer_name,omitempty"`
	Version         *int32                   `json:"version,omitempty"`
	CreatedAt       *time.Time               `json:"created_at,omitempty"`
	UpdatedAt       *time.Time               `json:"updated_at,omitempty"`
	Manifests       []map[string]interface{} `json:"manifests,omitempty"`
	DeleteOption    map[string]interface{}   `json:"delete_option,omitempty"`
	ManifestConfigs []map[string]interface{} `json:"manifest_configs,omitempty"`
	Status          map[string]interface{}   `json:"status,omitempty"`
}

// NewResourceBundleAllOf instantiates a new ResourceBundleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceBundleAllOf() *ResourceBundleAllOf {
	this := ResourceBundleAllOf{}
	return &this
}

// NewResourceBundleAllOfWithDefaults instantiates a new ResourceBundleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceBundleAllOfWithDefaults() *ResourceBundleAllOf {
	this := ResourceBundleAllOf{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ResourceBundleAllOf) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceBundleAllOf) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ResourceBundleAllOf) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ResourceBundleAllOf) SetName(v string) {
	o.Name = &v
}

// GetConsumerName returns the ConsumerName field value if set, zero value otherwise.
func (o *ResourceBundleAllOf) GetConsumerName() string {
	if o == nil || IsNil(o.ConsumerName) {
		var ret string
		return ret
	}
	return *o.ConsumerName
}

// GetConsumerNameOk returns a tuple with the ConsumerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceBundleAllOf) GetConsumerNameOk() (*string, bool) {
	if o == nil || IsNil(o.ConsumerName) {
		return nil, false
	}
	return o.ConsumerName, true
}

// HasConsumerName returns a boolean if a field has been set.
func (o *ResourceBundleAllOf) HasConsumerName() bool {
	if o != nil && !IsNil(o.ConsumerName) {
		return true
	}

	return false
}

// SetConsumerName gets a reference to the given string and assigns it to the ConsumerName field.
func (o *ResourceBundleAllOf) SetConsumerName(v string) {
	o.ConsumerName = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ResourceBundleAllOf) GetVersion() int32 {
	if o == nil || IsNil(o.Version) {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceBundleAllOf) GetVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ResourceBundleAllOf) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *ResourceBundleAllOf) SetVersion(v int32) {
	o.Version = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ResourceBundleAllOf) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceBundleAllOf) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ResourceBundleAllOf) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ResourceBundleAllOf) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ResourceBundleAllOf) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceBundleAllOf) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ResourceBundleAllOf) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *ResourceBundleAllOf) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetManifests returns the Manifests field value if set, zero value otherwise.
func (o *ResourceBundleAllOf) GetManifests() []map[string]interface{} {
	if o == nil || IsNil(o.Manifests) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Manifests
}

// GetManifestsOk returns a tuple with the Manifests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceBundleAllOf) GetManifestsOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Manifests) {
		return nil, false
	}
	return o.Manifests, true
}

// HasManifests returns a boolean if a field has been set.
func (o *ResourceBundleAllOf) HasManifests() bool {
	if o != nil && !IsNil(o.Manifests) {
		return true
	}

	return false
}

// SetManifests gets a reference to the given []map[string]interface{} and assigns it to the Manifests field.
func (o *ResourceBundleAllOf) SetManifests(v []map[string]interface{}) {
	o.Manifests = v
}

// GetDeleteOption returns the DeleteOption field value if set, zero value otherwise.
func (o *ResourceBundleAllOf) GetDeleteOption() map[string]interface{} {
	if o == nil || IsNil(o.DeleteOption) {
		var ret map[string]interface{}
		return ret
	}
	return o.DeleteOption
}

// GetDeleteOptionOk returns a tuple with the DeleteOption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceBundleAllOf) GetDeleteOptionOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.DeleteOption) {
		return map[string]interface{}{}, false
	}
	return o.DeleteOption, true
}

// HasDeleteOption returns a boolean if a field has been set.
func (o *ResourceBundleAllOf) HasDeleteOption() bool {
	if o != nil && !IsNil(o.DeleteOption) {
		return true
	}

	return false
}

// SetDeleteOption gets a reference to the given map[string]interface{} and assigns it to the DeleteOption field.
func (o *ResourceBundleAllOf) SetDeleteOption(v map[string]interface{}) {
	o.DeleteOption = v
}

// GetManifestConfigs returns the ManifestConfigs field value if set, zero value otherwise.
func (o *ResourceBundleAllOf) GetManifestConfigs() []map[string]interface{} {
	if o == nil || IsNil(o.ManifestConfigs) {
		var ret []map[string]interface{}
		return ret
	}
	return o.ManifestConfigs
}

// GetManifestConfigsOk returns a tuple with the ManifestConfigs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceBundleAllOf) GetManifestConfigsOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.ManifestConfigs) {
		return nil, false
	}
	return o.ManifestConfigs, true
}

// HasManifestConfigs returns a boolean if a field has been set.
func (o *ResourceBundleAllOf) HasManifestConfigs() bool {
	if o != nil && !IsNil(o.ManifestConfigs) {
		return true
	}

	return false
}

// SetManifestConfigs gets a reference to the given []map[string]interface{} and assigns it to the ManifestConfigs field.
func (o *ResourceBundleAllOf) SetManifestConfigs(v []map[string]interface{}) {
	o.ManifestConfigs = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ResourceBundleAllOf) GetStatus() map[string]interface{} {
	if o == nil || IsNil(o.Status) {
		var ret map[string]interface{}
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceBundleAllOf) GetStatusOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Status) {
		return map[string]interface{}{}, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ResourceBundleAllOf) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given map[string]interface{} and assigns it to the Status field.
func (o *ResourceBundleAllOf) SetStatus(v map[string]interface{}) {
	o.Status = v
}

func (o ResourceBundleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceBundleAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ConsumerName) {
		toSerialize["consumer_name"] = o.ConsumerName
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.Manifests) {
		toSerialize["manifests"] = o.Manifests
	}
	if !IsNil(o.DeleteOption) {
		toSerialize["delete_option"] = o.DeleteOption
	}
	if !IsNil(o.ManifestConfigs) {
		toSerialize["manifest_configs"] = o.ManifestConfigs
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableResourceBundleAllOf struct {
	value *ResourceBundleAllOf
	isSet bool
}

func (v NullableResourceBundleAllOf) Get() *ResourceBundleAllOf {
	return v.value
}

func (v *NullableResourceBundleAllOf) Set(val *ResourceBundleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceBundleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceBundleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceBundleAllOf(val *ResourceBundleAllOf) *NullableResourceBundleAllOf {
	return &NullableResourceBundleAllOf{value: val, isSet: true}
}

func (v NullableResourceBundleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceBundleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}