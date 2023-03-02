# Engine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Object** | **string** |  | 
**Created** | **NullableInt32** |  | 
**Ready** | **bool** |  | 

## Methods

### NewEngine

`func NewEngine(id string, object string, created NullableInt32, ready bool, ) *Engine`

NewEngine instantiates a new Engine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEngineWithDefaults

`func NewEngineWithDefaults() *Engine`

NewEngineWithDefaults instantiates a new Engine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Engine) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Engine) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Engine) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *Engine) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *Engine) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *Engine) SetObject(v string)`

SetObject sets Object field to given value.


### GetCreated

`func (o *Engine) GetCreated() int32`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Engine) GetCreatedOk() (*int32, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Engine) SetCreated(v int32)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *Engine) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *Engine) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetReady

`func (o *Engine) GetReady() bool`

GetReady returns the Ready field if non-nil, zero value otherwise.

### GetReadyOk

`func (o *Engine) GetReadyOk() (*bool, bool)`

GetReadyOk returns a tuple with the Ready field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReady

`func (o *Engine) SetReady(v bool)`

SetReady sets Ready field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


