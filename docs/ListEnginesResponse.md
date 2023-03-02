# ListEnginesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** |  | 
**Data** | [**[]Engine**](Engine.md) |  | 

## Methods

### NewListEnginesResponse

`func NewListEnginesResponse(object string, data []Engine, ) *ListEnginesResponse`

NewListEnginesResponse instantiates a new ListEnginesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListEnginesResponseWithDefaults

`func NewListEnginesResponseWithDefaults() *ListEnginesResponse`

NewListEnginesResponseWithDefaults instantiates a new ListEnginesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *ListEnginesResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ListEnginesResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ListEnginesResponse) SetObject(v string)`

SetObject sets Object field to given value.


### GetData

`func (o *ListEnginesResponse) GetData() []Engine`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListEnginesResponse) GetDataOk() (*[]Engine, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListEnginesResponse) SetData(v []Engine)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


