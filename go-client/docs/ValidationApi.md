# \ValidationApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiContractValidatePost**](ValidationApi.md#ApiContractValidatePost) | **Post** /api/contract/validate | 


# **ApiContractValidatePost**
> ApiContractValidatePost(ctx, isCreate, contract, contentType, optional)


Validate a framework contract rule set.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **isCreate** | **bool**| Indicates if the rule set we are validating is a new rule set or an update to an existing one. If this is true then there should be no existing contract rule set entries which have the same fc_Id in the database | 
  **contract** | [**ContractModel**](ContractModel.md)| The framework contract rule set in JSON format that is to be validated. | 
  **contentType** | **string**| Specifies the format of the data that is being sent to the API. Only JSON data is accepted | 
 **optional** | ***ValidationApiApiContractValidatePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ValidationApiApiContractValidatePostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **clientRequestSource** | **optional.String**| The client request header that identifies which client had made a request to CPQ Server. | 
 **cpqLanguage** | **optional.String**| The language code (in ISO 639-1 format) that error messages are to be translated to for returning back to the client. If not supplied it defaults to en-GB | 
 **performTransformation** | **optional.Bool**| Perform transformation of referenced product specifications using the given framework contract into transformed specifications or not. If not supplied defaults to false | 
 **returnNegotiatedSpecifications** | **optional.Bool**| Return the transformed specifications in response or not. If not supplied defaults to false | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

