/*
OpenAI API

APIs for sampling from and fine-tuning language models

API version: 1.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the CreateImageRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateImageRequest{}

// CreateImageRequest struct for CreateImageRequest
type CreateImageRequest struct {
	// A text description of the desired image(s). The maximum length is 1000 characters.
	Prompt string `json:"prompt"`
	// The number of images to generate. Must be between 1 and 10.
	N NullableInt32 `json:"n,omitempty"`
	// The size of the generated images. Must be one of `256x256`, `512x512`, or `1024x1024`.
	Size NullableString `json:"size,omitempty"`
	// The format in which the generated images are returned. Must be one of `url` or `b64_json`.
	ResponseFormat NullableString `json:"response_format,omitempty"`
	// A unique identifier representing your end-user, which can help OpenAI to monitor and detect abuse. [Learn more](/docs/guides/safety-best-practices/end-user-ids). 
	User *string `json:"user,omitempty"`
}



// NewCreateImageRequest instantiates a new CreateImageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateImageRequest(prompt string) *CreateImageRequest {
	this := CreateImageRequest{}
	this.Prompt = prompt
	var n int32 = 1
	this.N = *NewNullableInt32(&n)
	var size string = "1024x1024"
	this.Size = *NewNullableString(&size)
	var responseFormat string = "url"
	this.ResponseFormat = *NewNullableString(&responseFormat)
	return &this
}

// NewCreateImageRequestWithDefaults instantiates a new CreateImageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateImageRequestWithDefaults() *CreateImageRequest {
	this := CreateImageRequest{}
	var n int32 = 1
	this.N = *NewNullableInt32(&n)
	var size string = "1024x1024"
	this.Size = *NewNullableString(&size)
	var responseFormat string = "url"
	this.ResponseFormat = *NewNullableString(&responseFormat)
	return &this
}

// GetPrompt returns the Prompt field value
func (o *CreateImageRequest) GetPrompt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Prompt
}

// GetPromptOk returns a tuple with the Prompt field value
// and a boolean to check if the value has been set.
func (o *CreateImageRequest) GetPromptOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Prompt, true
}

// SetPrompt sets field value
func (o *CreateImageRequest) SetPrompt(v string) {
	o.Prompt = v
}

// GetN returns the N field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateImageRequest) GetN() int32 {
	if o == nil || IsNil(o.N.Get()) {
		var ret int32
		return ret
	}
	return *o.N.Get()
}

// GetNOk returns a tuple with the N field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateImageRequest) GetNOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.N.Get(), o.N.IsSet()
}

// HasN returns a boolean if a field has been set.
func (o *CreateImageRequest) HasN() bool {
	if o != nil && o.N.IsSet() {
		return true
	}

	return false
}

// SetN gets a reference to the given NullableInt32 and assigns it to the N field.
func (o *CreateImageRequest) SetN(v int32) {
	o.N.Set(&v)
}
// SetNNil sets the value for N to be an explicit nil
func (o *CreateImageRequest) SetNNil() {
	o.N.Set(nil)
}

// UnsetN ensures that no value is present for N, not even an explicit nil
func (o *CreateImageRequest) UnsetN() {
	o.N.Unset()
}

// GetSize returns the Size field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateImageRequest) GetSize() string {
	if o == nil || IsNil(o.Size.Get()) {
		var ret string
		return ret
	}
	return *o.Size.Get()
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateImageRequest) GetSizeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Size.Get(), o.Size.IsSet()
}

// HasSize returns a boolean if a field has been set.
func (o *CreateImageRequest) HasSize() bool {
	if o != nil && o.Size.IsSet() {
		return true
	}

	return false
}

// SetSize gets a reference to the given NullableString and assigns it to the Size field.
func (o *CreateImageRequest) SetSize(v string) {
	o.Size.Set(&v)
}
// SetSizeNil sets the value for Size to be an explicit nil
func (o *CreateImageRequest) SetSizeNil() {
	o.Size.Set(nil)
}

// UnsetSize ensures that no value is present for Size, not even an explicit nil
func (o *CreateImageRequest) UnsetSize() {
	o.Size.Unset()
}

// GetResponseFormat returns the ResponseFormat field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateImageRequest) GetResponseFormat() string {
	if o == nil || IsNil(o.ResponseFormat.Get()) {
		var ret string
		return ret
	}
	return *o.ResponseFormat.Get()
}

// GetResponseFormatOk returns a tuple with the ResponseFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateImageRequest) GetResponseFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResponseFormat.Get(), o.ResponseFormat.IsSet()
}

// HasResponseFormat returns a boolean if a field has been set.
func (o *CreateImageRequest) HasResponseFormat() bool {
	if o != nil && o.ResponseFormat.IsSet() {
		return true
	}

	return false
}

// SetResponseFormat gets a reference to the given NullableString and assigns it to the ResponseFormat field.
func (o *CreateImageRequest) SetResponseFormat(v string) {
	o.ResponseFormat.Set(&v)
}
// SetResponseFormatNil sets the value for ResponseFormat to be an explicit nil
func (o *CreateImageRequest) SetResponseFormatNil() {
	o.ResponseFormat.Set(nil)
}

// UnsetResponseFormat ensures that no value is present for ResponseFormat, not even an explicit nil
func (o *CreateImageRequest) UnsetResponseFormat() {
	o.ResponseFormat.Unset()
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *CreateImageRequest) GetUser() string {
	if o == nil || IsNil(o.User) {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateImageRequest) GetUserOk() (*string, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *CreateImageRequest) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *CreateImageRequest) SetUser(v string) {
	o.User = &v
}

func (o CreateImageRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateImageRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["prompt"] = o.Prompt
	if o.N.IsSet() {
		toSerialize["n"] = o.N.Get()
	}
	if o.Size.IsSet() {
		toSerialize["size"] = o.Size.Get()
	}
	if o.ResponseFormat.IsSet() {
		toSerialize["response_format"] = o.ResponseFormat.Get()
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	return toSerialize, nil
}

type NullableCreateImageRequest struct {
	value *CreateImageRequest
	isSet bool
}

func (v NullableCreateImageRequest) Get() *CreateImageRequest {
	return v.value
}

func (v *NullableCreateImageRequest) Set(val *CreateImageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateImageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateImageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateImageRequest(val *CreateImageRequest) *NullableCreateImageRequest {
	return &NullableCreateImageRequest{value: val, isSet: true}
}

func (v NullableCreateImageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateImageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


