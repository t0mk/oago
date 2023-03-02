# CreateSearchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | Pointer to **string** |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**[]CreateSearchResponseDataInner**](CreateSearchResponseDataInner.md) |  | [optional] 

## Methods

### NewCreateSearchResponse

`func NewCreateSearchResponse() *CreateSearchResponse`

NewCreateSearchResponse instantiates a new CreateSearchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSearchResponseWithDefaults

`func NewCreateSearchResponseWithDefaults() *CreateSearchResponse`

NewCreateSearchResponseWithDefaults instantiates a new CreateSearchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *CreateSearchResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *CreateSearchResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *CreateSearchResponse) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *CreateSearchResponse) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetModel

`func (o *CreateSearchResponse) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *CreateSearchResponse) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *CreateSearchResponse) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *CreateSearchResponse) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetData

`func (o *CreateSearchResponse) GetData() []CreateSearchResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreateSearchResponse) GetDataOk() (*[]CreateSearchResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreateSearchResponse) SetData(v []CreateSearchResponseDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *CreateSearchResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


