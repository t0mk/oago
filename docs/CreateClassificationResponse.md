# CreateClassificationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | Pointer to **string** |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**SearchModel** | Pointer to **string** |  | [optional] 
**Completion** | Pointer to **string** |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**SelectedExamples** | Pointer to [**[]CreateClassificationResponseSelectedExamplesInner**](CreateClassificationResponseSelectedExamplesInner.md) |  | [optional] 

## Methods

### NewCreateClassificationResponse

`func NewCreateClassificationResponse() *CreateClassificationResponse`

NewCreateClassificationResponse instantiates a new CreateClassificationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateClassificationResponseWithDefaults

`func NewCreateClassificationResponseWithDefaults() *CreateClassificationResponse`

NewCreateClassificationResponseWithDefaults instantiates a new CreateClassificationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *CreateClassificationResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *CreateClassificationResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *CreateClassificationResponse) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *CreateClassificationResponse) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetModel

`func (o *CreateClassificationResponse) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *CreateClassificationResponse) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *CreateClassificationResponse) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *CreateClassificationResponse) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetSearchModel

`func (o *CreateClassificationResponse) GetSearchModel() string`

GetSearchModel returns the SearchModel field if non-nil, zero value otherwise.

### GetSearchModelOk

`func (o *CreateClassificationResponse) GetSearchModelOk() (*string, bool)`

GetSearchModelOk returns a tuple with the SearchModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchModel

`func (o *CreateClassificationResponse) SetSearchModel(v string)`

SetSearchModel sets SearchModel field to given value.

### HasSearchModel

`func (o *CreateClassificationResponse) HasSearchModel() bool`

HasSearchModel returns a boolean if a field has been set.

### GetCompletion

`func (o *CreateClassificationResponse) GetCompletion() string`

GetCompletion returns the Completion field if non-nil, zero value otherwise.

### GetCompletionOk

`func (o *CreateClassificationResponse) GetCompletionOk() (*string, bool)`

GetCompletionOk returns a tuple with the Completion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletion

`func (o *CreateClassificationResponse) SetCompletion(v string)`

SetCompletion sets Completion field to given value.

### HasCompletion

`func (o *CreateClassificationResponse) HasCompletion() bool`

HasCompletion returns a boolean if a field has been set.

### GetLabel

`func (o *CreateClassificationResponse) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CreateClassificationResponse) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CreateClassificationResponse) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *CreateClassificationResponse) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetSelectedExamples

`func (o *CreateClassificationResponse) GetSelectedExamples() []CreateClassificationResponseSelectedExamplesInner`

GetSelectedExamples returns the SelectedExamples field if non-nil, zero value otherwise.

### GetSelectedExamplesOk

`func (o *CreateClassificationResponse) GetSelectedExamplesOk() (*[]CreateClassificationResponseSelectedExamplesInner, bool)`

GetSelectedExamplesOk returns a tuple with the SelectedExamples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedExamples

`func (o *CreateClassificationResponse) SetSelectedExamples(v []CreateClassificationResponseSelectedExamplesInner)`

SetSelectedExamples sets SelectedExamples field to given value.

### HasSelectedExamples

`func (o *CreateClassificationResponse) HasSelectedExamples() bool`

HasSelectedExamples returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


