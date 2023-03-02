# CreateSearchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | **string** | Query to search against the documents. | 
**Documents** | Pointer to **[]string** | Up to 200 documents to search over, provided as a list of strings.  The maximum document length (in tokens) is 2034 minus the number of tokens in the query.  You should specify either &#x60;documents&#x60; or a &#x60;file&#x60;, but not both.  | [optional] 
**File** | Pointer to **NullableString** | The ID of an uploaded file that contains documents to search over.  You should specify either &#x60;documents&#x60; or a &#x60;file&#x60;, but not both.  | [optional] 
**MaxRerank** | Pointer to **NullableInt32** | The maximum number of documents to be re-ranked and returned by search.  This flag only takes effect when &#x60;file&#x60; is set.  | [optional] [default to 200]
**ReturnMetadata** | Pointer to **NullableBool** | A special boolean flag for showing metadata. If set to &#x60;true&#x60;, each document entry in the returned JSON will contain a \&quot;metadata\&quot; field.  This flag only takes effect when &#x60;file&#x60; is set.  | [optional] [default to false]
**User** | Pointer to **string** | A unique identifier representing your end-user, which can help OpenAI to monitor and detect abuse. [Learn more](/docs/guides/safety-best-practices/end-user-ids).  | [optional] 

## Methods

### NewCreateSearchRequest

`func NewCreateSearchRequest(query string, ) *CreateSearchRequest`

NewCreateSearchRequest instantiates a new CreateSearchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSearchRequestWithDefaults

`func NewCreateSearchRequestWithDefaults() *CreateSearchRequest`

NewCreateSearchRequestWithDefaults instantiates a new CreateSearchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuery

`func (o *CreateSearchRequest) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *CreateSearchRequest) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *CreateSearchRequest) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetDocuments

`func (o *CreateSearchRequest) GetDocuments() []string`

GetDocuments returns the Documents field if non-nil, zero value otherwise.

### GetDocumentsOk

`func (o *CreateSearchRequest) GetDocumentsOk() (*[]string, bool)`

GetDocumentsOk returns a tuple with the Documents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocuments

`func (o *CreateSearchRequest) SetDocuments(v []string)`

SetDocuments sets Documents field to given value.

### HasDocuments

`func (o *CreateSearchRequest) HasDocuments() bool`

HasDocuments returns a boolean if a field has been set.

### SetDocumentsNil

`func (o *CreateSearchRequest) SetDocumentsNil(b bool)`

 SetDocumentsNil sets the value for Documents to be an explicit nil

### UnsetDocuments
`func (o *CreateSearchRequest) UnsetDocuments()`

UnsetDocuments ensures that no value is present for Documents, not even an explicit nil
### GetFile

`func (o *CreateSearchRequest) GetFile() string`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *CreateSearchRequest) GetFileOk() (*string, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *CreateSearchRequest) SetFile(v string)`

SetFile sets File field to given value.

### HasFile

`func (o *CreateSearchRequest) HasFile() bool`

HasFile returns a boolean if a field has been set.

### SetFileNil

`func (o *CreateSearchRequest) SetFileNil(b bool)`

 SetFileNil sets the value for File to be an explicit nil

### UnsetFile
`func (o *CreateSearchRequest) UnsetFile()`

UnsetFile ensures that no value is present for File, not even an explicit nil
### GetMaxRerank

`func (o *CreateSearchRequest) GetMaxRerank() int32`

GetMaxRerank returns the MaxRerank field if non-nil, zero value otherwise.

### GetMaxRerankOk

`func (o *CreateSearchRequest) GetMaxRerankOk() (*int32, bool)`

GetMaxRerankOk returns a tuple with the MaxRerank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRerank

`func (o *CreateSearchRequest) SetMaxRerank(v int32)`

SetMaxRerank sets MaxRerank field to given value.

### HasMaxRerank

`func (o *CreateSearchRequest) HasMaxRerank() bool`

HasMaxRerank returns a boolean if a field has been set.

### SetMaxRerankNil

`func (o *CreateSearchRequest) SetMaxRerankNil(b bool)`

 SetMaxRerankNil sets the value for MaxRerank to be an explicit nil

### UnsetMaxRerank
`func (o *CreateSearchRequest) UnsetMaxRerank()`

UnsetMaxRerank ensures that no value is present for MaxRerank, not even an explicit nil
### GetReturnMetadata

`func (o *CreateSearchRequest) GetReturnMetadata() bool`

GetReturnMetadata returns the ReturnMetadata field if non-nil, zero value otherwise.

### GetReturnMetadataOk

`func (o *CreateSearchRequest) GetReturnMetadataOk() (*bool, bool)`

GetReturnMetadataOk returns a tuple with the ReturnMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnMetadata

`func (o *CreateSearchRequest) SetReturnMetadata(v bool)`

SetReturnMetadata sets ReturnMetadata field to given value.

### HasReturnMetadata

`func (o *CreateSearchRequest) HasReturnMetadata() bool`

HasReturnMetadata returns a boolean if a field has been set.

### SetReturnMetadataNil

`func (o *CreateSearchRequest) SetReturnMetadataNil(b bool)`

 SetReturnMetadataNil sets the value for ReturnMetadata to be an explicit nil

### UnsetReturnMetadata
`func (o *CreateSearchRequest) UnsetReturnMetadata()`

UnsetReturnMetadata ensures that no value is present for ReturnMetadata, not even an explicit nil
### GetUser

`func (o *CreateSearchRequest) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *CreateSearchRequest) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *CreateSearchRequest) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *CreateSearchRequest) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


