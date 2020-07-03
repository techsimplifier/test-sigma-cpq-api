# FcValidationItem

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContextPath** | **string** | The path to the data element in framework contract | [optional] [default to null]
**ResolutionText** | **string** | For some scenarios it is possible to provide additional information that will prevent any future calls from failing. This property will contain that additional description which will be translated into the language of the calling client | [optional] [default to null]
**ResponseCode** | **string** | The unique code which describes the exception which occurred or is specified as SUCCESS for success responses | [optional] [default to null]
**ResponseText** | **string** | A fuller description of the exception which occurred that is translated to the requested language of the calling client. If the response is a success response this property will contain the word SUCCESS translated into the language of the calling client. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


