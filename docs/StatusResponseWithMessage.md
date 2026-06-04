# StatusResponseWithMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int32** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 

## Methods

### NewStatusResponseWithMessage

`func NewStatusResponseWithMessage() *StatusResponseWithMessage`

NewStatusResponseWithMessage instantiates a new StatusResponseWithMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusResponseWithMessageWithDefaults

`func NewStatusResponseWithMessageWithDefaults() *StatusResponseWithMessage`

NewStatusResponseWithMessageWithDefaults instantiates a new StatusResponseWithMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *StatusResponseWithMessage) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *StatusResponseWithMessage) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *StatusResponseWithMessage) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *StatusResponseWithMessage) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *StatusResponseWithMessage) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StatusResponseWithMessage) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StatusResponseWithMessage) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *StatusResponseWithMessage) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCode

`func (o *StatusResponseWithMessage) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *StatusResponseWithMessage) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *StatusResponseWithMessage) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *StatusResponseWithMessage) HasCode() bool`

HasCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


