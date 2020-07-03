# SparseCreateQuoteItemErrorResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExceptionType** | **string** | The type of exception which occurred. For success responses this will not be present | [optional] [default to null]
**HttpStatus** | **float32** | The http status code which describes the response | [optional] [default to null]
**ResponseCode** | **string** | The unique code which describes the exception which occurred or is specified as SUCCESS for success responses | [optional] [default to null]
**ResponseText** | **string** | A fuller description of the exception which occurred that is translated to the requested language of the calling client. If the response is a success response this property will contain the word SUCCESS translated into the language of the calling client. | [optional] [default to null]
**ResponseBody** | [**[]SparseCreateQuoteItemValidationRestErrorResponse**](SparseCreateQuoteItemValidationRestErrorResponse.md) | The details of what parameter validation failed | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


