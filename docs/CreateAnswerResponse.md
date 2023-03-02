# CreateAnswerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | Pointer to **string** |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**SearchModel** | Pointer to **string** |  | [optional] 
**Completion** | Pointer to **string** |  | [optional] 
**Answers** | Pointer to **[]string** |  | [optional] 
**SelectedDocuments** | Pointer to [**[]CreateAnswerResponseSelectedDocumentsInner**](CreateAnswerResponseSelectedDocumentsInner.md) |  | [optional] 

## Methods

### NewCreateAnswerResponse

`func NewCreateAnswerResponse() *CreateAnswerResponse`

NewCreateAnswerResponse instantiates a new CreateAnswerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAnswerResponseWithDefaults

`func NewCreateAnswerResponseWithDefaults() *CreateAnswerResponse`

NewCreateAnswerResponseWithDefaults instantiates a new CreateAnswerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *CreateAnswerResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *CreateAnswerResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *CreateAnswerResponse) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *CreateAnswerResponse) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetModel

`func (o *CreateAnswerResponse) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *CreateAnswerResponse) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *CreateAnswerResponse) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *CreateAnswerResponse) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetSearchModel

`func (o *CreateAnswerResponse) GetSearchModel() string`

GetSearchModel returns the SearchModel field if non-nil, zero value otherwise.

### GetSearchModelOk

`func (o *CreateAnswerResponse) GetSearchModelOk() (*string, bool)`

GetSearchModelOk returns a tuple with the SearchModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchModel

`func (o *CreateAnswerResponse) SetSearchModel(v string)`

SetSearchModel sets SearchModel field to given value.

### HasSearchModel

`func (o *CreateAnswerResponse) HasSearchModel() bool`

HasSearchModel returns a boolean if a field has been set.

### GetCompletion

`func (o *CreateAnswerResponse) GetCompletion() string`

GetCompletion returns the Completion field if non-nil, zero value otherwise.

### GetCompletionOk

`func (o *CreateAnswerResponse) GetCompletionOk() (*string, bool)`

GetCompletionOk returns a tuple with the Completion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletion

`func (o *CreateAnswerResponse) SetCompletion(v string)`

SetCompletion sets Completion field to given value.

### HasCompletion

`func (o *CreateAnswerResponse) HasCompletion() bool`

HasCompletion returns a boolean if a field has been set.

### GetAnswers

`func (o *CreateAnswerResponse) GetAnswers() []string`

GetAnswers returns the Answers field if non-nil, zero value otherwise.

### GetAnswersOk

`func (o *CreateAnswerResponse) GetAnswersOk() (*[]string, bool)`

GetAnswersOk returns a tuple with the Answers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswers

`func (o *CreateAnswerResponse) SetAnswers(v []string)`

SetAnswers sets Answers field to given value.

### HasAnswers

`func (o *CreateAnswerResponse) HasAnswers() bool`

HasAnswers returns a boolean if a field has been set.

### GetSelectedDocuments

`func (o *CreateAnswerResponse) GetSelectedDocuments() []CreateAnswerResponseSelectedDocumentsInner`

GetSelectedDocuments returns the SelectedDocuments field if non-nil, zero value otherwise.

### GetSelectedDocumentsOk

`func (o *CreateAnswerResponse) GetSelectedDocumentsOk() (*[]CreateAnswerResponseSelectedDocumentsInner, bool)`

GetSelectedDocumentsOk returns a tuple with the SelectedDocuments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedDocuments

`func (o *CreateAnswerResponse) SetSelectedDocuments(v []CreateAnswerResponseSelectedDocumentsInner)`

SetSelectedDocuments sets SelectedDocuments field to given value.

### HasSelectedDocuments

`func (o *CreateAnswerResponse) HasSelectedDocuments() bool`

HasSelectedDocuments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


