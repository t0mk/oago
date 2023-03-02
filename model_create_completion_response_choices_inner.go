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

// checks if the CreateCompletionResponseChoicesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateCompletionResponseChoicesInner{}

// CreateCompletionResponseChoicesInner struct for CreateCompletionResponseChoicesInner
type CreateCompletionResponseChoicesInner struct {
	Text *string `json:"text,omitempty"`
	Index *int32 `json:"index,omitempty"`
	Logprobs NullableCreateCompletionResponseChoicesInnerLogprobs `json:"logprobs,omitempty"`
	FinishReason *string `json:"finish_reason,omitempty"`
}

// NewCreateCompletionResponseChoicesInner instantiates a new CreateCompletionResponseChoicesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCompletionResponseChoicesInner() *CreateCompletionResponseChoicesInner {
	this := CreateCompletionResponseChoicesInner{}
	return &this
}

// NewCreateCompletionResponseChoicesInnerWithDefaults instantiates a new CreateCompletionResponseChoicesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCompletionResponseChoicesInnerWithDefaults() *CreateCompletionResponseChoicesInner {
	this := CreateCompletionResponseChoicesInner{}
	return &this
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *CreateCompletionResponseChoicesInner) GetText() string {
	if o == nil || IsNil(o.Text) {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCompletionResponseChoicesInner) GetTextOk() (*string, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *CreateCompletionResponseChoicesInner) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *CreateCompletionResponseChoicesInner) SetText(v string) {
	o.Text = &v
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *CreateCompletionResponseChoicesInner) GetIndex() int32 {
	if o == nil || IsNil(o.Index) {
		var ret int32
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCompletionResponseChoicesInner) GetIndexOk() (*int32, bool) {
	if o == nil || IsNil(o.Index) {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *CreateCompletionResponseChoicesInner) HasIndex() bool {
	if o != nil && !IsNil(o.Index) {
		return true
	}

	return false
}

// SetIndex gets a reference to the given int32 and assigns it to the Index field.
func (o *CreateCompletionResponseChoicesInner) SetIndex(v int32) {
	o.Index = &v
}

// GetLogprobs returns the Logprobs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateCompletionResponseChoicesInner) GetLogprobs() CreateCompletionResponseChoicesInnerLogprobs {
	if o == nil || IsNil(o.Logprobs.Get()) {
		var ret CreateCompletionResponseChoicesInnerLogprobs
		return ret
	}
	return *o.Logprobs.Get()
}

// GetLogprobsOk returns a tuple with the Logprobs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateCompletionResponseChoicesInner) GetLogprobsOk() (*CreateCompletionResponseChoicesInnerLogprobs, bool) {
	if o == nil {
		return nil, false
	}
	return o.Logprobs.Get(), o.Logprobs.IsSet()
}

// HasLogprobs returns a boolean if a field has been set.
func (o *CreateCompletionResponseChoicesInner) HasLogprobs() bool {
	if o != nil && o.Logprobs.IsSet() {
		return true
	}

	return false
}

// SetLogprobs gets a reference to the given NullableCreateCompletionResponseChoicesInnerLogprobs and assigns it to the Logprobs field.
func (o *CreateCompletionResponseChoicesInner) SetLogprobs(v CreateCompletionResponseChoicesInnerLogprobs) {
	o.Logprobs.Set(&v)
}
// SetLogprobsNil sets the value for Logprobs to be an explicit nil
func (o *CreateCompletionResponseChoicesInner) SetLogprobsNil() {
	o.Logprobs.Set(nil)
}

// UnsetLogprobs ensures that no value is present for Logprobs, not even an explicit nil
func (o *CreateCompletionResponseChoicesInner) UnsetLogprobs() {
	o.Logprobs.Unset()
}

// GetFinishReason returns the FinishReason field value if set, zero value otherwise.
func (o *CreateCompletionResponseChoicesInner) GetFinishReason() string {
	if o == nil || IsNil(o.FinishReason) {
		var ret string
		return ret
	}
	return *o.FinishReason
}

// GetFinishReasonOk returns a tuple with the FinishReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCompletionResponseChoicesInner) GetFinishReasonOk() (*string, bool) {
	if o == nil || IsNil(o.FinishReason) {
		return nil, false
	}
	return o.FinishReason, true
}

// HasFinishReason returns a boolean if a field has been set.
func (o *CreateCompletionResponseChoicesInner) HasFinishReason() bool {
	if o != nil && !IsNil(o.FinishReason) {
		return true
	}

	return false
}

// SetFinishReason gets a reference to the given string and assigns it to the FinishReason field.
func (o *CreateCompletionResponseChoicesInner) SetFinishReason(v string) {
	o.FinishReason = &v
}

func (o CreateCompletionResponseChoicesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateCompletionResponseChoicesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Text) {
		toSerialize["text"] = o.Text
	}
	if !IsNil(o.Index) {
		toSerialize["index"] = o.Index
	}
	if o.Logprobs.IsSet() {
		toSerialize["logprobs"] = o.Logprobs.Get()
	}
	if !IsNil(o.FinishReason) {
		toSerialize["finish_reason"] = o.FinishReason
	}
	return toSerialize, nil
}

type NullableCreateCompletionResponseChoicesInner struct {
	value *CreateCompletionResponseChoicesInner
	isSet bool
}

func (v NullableCreateCompletionResponseChoicesInner) Get() *CreateCompletionResponseChoicesInner {
	return v.value
}

func (v *NullableCreateCompletionResponseChoicesInner) Set(val *CreateCompletionResponseChoicesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCompletionResponseChoicesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCompletionResponseChoicesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCompletionResponseChoicesInner(val *CreateCompletionResponseChoicesInner) *NullableCreateCompletionResponseChoicesInner {
	return &NullableCreateCompletionResponseChoicesInner{value: val, isSet: true}
}

func (v NullableCreateCompletionResponseChoicesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCompletionResponseChoicesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


