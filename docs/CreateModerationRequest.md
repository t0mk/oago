# CreateModerationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Input** | [**CreateModerationRequestInput**](CreateModerationRequestInput.md) |  | 
**Model** | Pointer to **string** | Two content moderations models are available: &#x60;text-moderation-stable&#x60; and &#x60;text-moderation-latest&#x60;.  The default is &#x60;text-moderation-latest&#x60; which will be automatically upgraded over time. This ensures you are always using our most accurate model. If you use &#x60;text-moderation-stable&#x60;, we will provide advanced notice before updating the model. Accuracy of &#x60;text-moderation-stable&#x60; may be slightly lower than for &#x60;text-moderation-latest&#x60;.  | [optional] [default to "text-moderation-latest"]

## Methods

### NewCreateModerationRequest

`func NewCreateModerationRequest(input CreateModerationRequestInput, ) *CreateModerationRequest`

NewCreateModerationRequest instantiates a new CreateModerationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateModerationRequestWithDefaults

`func NewCreateModerationRequestWithDefaults() *CreateModerationRequest`

NewCreateModerationRequestWithDefaults instantiates a new CreateModerationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInput

`func (o *CreateModerationRequest) GetInput() CreateModerationRequestInput`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *CreateModerationRequest) GetInputOk() (*CreateModerationRequestInput, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *CreateModerationRequest) SetInput(v CreateModerationRequestInput)`

SetInput sets Input field to given value.


### GetModel

`func (o *CreateModerationRequest) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *CreateModerationRequest) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *CreateModerationRequest) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *CreateModerationRequest) HasModel() bool`

HasModel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


