# ChatCompletionResponseMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Role** | **string** | The role of the author of this message. | 
**Content** | **string** | The contents of the message | 

## Methods

### NewChatCompletionResponseMessage

`func NewChatCompletionResponseMessage(role string, content string, ) *ChatCompletionResponseMessage`

NewChatCompletionResponseMessage instantiates a new ChatCompletionResponseMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatCompletionResponseMessageWithDefaults

`func NewChatCompletionResponseMessageWithDefaults() *ChatCompletionResponseMessage`

NewChatCompletionResponseMessageWithDefaults instantiates a new ChatCompletionResponseMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRole

`func (o *ChatCompletionResponseMessage) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ChatCompletionResponseMessage) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ChatCompletionResponseMessage) SetRole(v string)`

SetRole sets Role field to given value.


### GetContent

`func (o *ChatCompletionResponseMessage) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ChatCompletionResponseMessage) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ChatCompletionResponseMessage) SetContent(v string)`

SetContent sets Content field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


