# \NegotiationContractsApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiContractsIdentifierClonePost**](NegotiationContractsApi.md#ApiContractsIdentifierClonePost) | **Post** /api/contracts/{identifier}/clone | 
[**ApiContractsIdentifierEndDatePut**](NegotiationContractsApi.md#ApiContractsIdentifierEndDatePut) | **Put** /api/contracts/{identifier}/endDate | 
[**ApiContractsIdentifierPut**](NegotiationContractsApi.md#ApiContractsIdentifierPut) | **Put** /api/contracts/{identifier} | 
[**ApiContractsIdentifierSelectableItemsGet**](NegotiationContractsApi.md#ApiContractsIdentifierSelectableItemsGet) | **Get** /api/contracts/{identifier}/selectableItems | 
[**ApiContractsIdentifierStatusStatusGet**](NegotiationContractsApi.md#ApiContractsIdentifierStatusStatusGet) | **Get** /api/contracts/{identifier}/status/{status} | 
[**ApiContractsPost**](NegotiationContractsApi.md#ApiContractsPost) | **Post** /api/contracts | 
[**ApiContractsStatusStatusGet**](NegotiationContractsApi.md#ApiContractsStatusStatusGet) | **Get** /api/contracts/status/{status} | 


# **ApiContractsIdentifierClonePost**
> ApiContractsIdentifierClonePost(ctx, contentType, fcUser, identifier, cloneConfigurationItems, optional)


Clone a framework contract into new copies each of which may contain or not contain the selection rules of the original

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contentType** | **string**| Specifies the format of the data that is being sent to the API. Only JSON data is accepted | 
  **fcUser** | **string**| The name of the user that is cloning the framework contract rule set | 
  **identifier** | **string**| The identifier of a contract rule set that is to be cloned - can be provided either as a guid value or as a value which matches the pattern ^[a-z0-9]{6}$ | 
  **cloneConfigurationItems** | [**CloneContractRequest**](CloneContractRequest.md)| The configurations which are to be applied to the existing contract to create cloned copies. | 
 **optional** | ***NegotiationContractsApiApiContractsIdentifierClonePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NegotiationContractsApiApiContractsIdentifierClonePostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **clientRequestSource** | **optional.String**| The client request header that identifies which client had made a request to CPQ Server. | 
 **cpqLanguage** | **optional.String**| The language code (in ISO 639-1 format) that error messages are to be translated to for returning back to the client. If not supplied it defaults to en-GB | 
 **count** | **optional.Float32**| A numeric value indicating how many framework contract clones that are to be created. This value should match the number of elements the clone configuration items array supplied in the request body. If not supplied it is defaulted to 1 | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiContractsIdentifierEndDatePut**
> ApiContractsIdentifierEndDatePut(ctx, fcUser, identifier, endDate, optional)


Update a framework contract rule and create negotiated specifications for the referenced offering and standalone discount specifications

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fcUser** | **string**| The name of the user that is creating the framework contract rule set | 
  **identifier** | **string**| The identifier of a contract rule set - can be provided either as a guid value or as a value which matches the pattern ^[a-z0-9]{6}$ | 
  **endDate** | **string**| The end date that you wish to set for the contract rule set. The value supplied must conform to ISO 8601 format (YYYY-MM-DD) or (YYYY-MM-DDTHH:mm:ss.SSSZ) | 
 **optional** | ***NegotiationContractsApiApiContractsIdentifierEndDatePutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NegotiationContractsApiApiContractsIdentifierEndDatePutOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **clientRequestSource** | **optional.String**| The client request header that identifies which client had made a request to CPQ Server. | 
 **cpqLanguage** | **optional.String**| The language code (in ISO 639-1 format) that error messages are to be translated to for returning back to the client. If not supplied it defaults to en-GB | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiContractsIdentifierPut**
> ApiContractsIdentifierPut(ctx, fcUser, identifier, contract, contentType, optional)


Update a framework contract rule and create negotiated specifications for the referenced offering and standalone discount specifications

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fcUser** | **string**| The name of the user that is creating the framework contract rule set | 
  **identifier** | **string**| The identifier of a contract rule set - can be provided either as a guid value or as a value which matches the pattern ^[a-z0-9]{6}$ | 
  **contract** | [**ContractModel**](ContractModel.md)| The framework contract rule set in JSON format that is to be written to the database. | 
  **contentType** | **string**| Specifies the format of the data that is being sent to the API. Only JSON data is accepted | 
 **optional** | ***NegotiationContractsApiApiContractsIdentifierPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NegotiationContractsApiApiContractsIdentifierPutOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **clientRequestSource** | **optional.String**| The client request header that identifies which client had made a request to CPQ Server. | 
 **cpqLanguage** | **optional.String**| The language code (in ISO 639-1 format) that error messages are to be translated to for returning back to the client. If not supplied it defaults to en-GB | 
 **performTransformation** | **optional.Bool**| Perform transformation of referenced product specifications using the given framework contract into transformed specifications or not. If not supplied defaults to true | 
 **returnNegotiatedSpecifications** | **optional.Bool**| Return the transformed specifications in response or not. If not supplied defaults to false | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiContractsIdentifierSelectableItemsGet**
> ApiContractsIdentifierSelectableItemsGet(ctx, identifier, optional)


Query the referenced offerings and standalone discounts on the framework contract with the specified rule set to identify those specification structures which can be used to generate selection rules. The data is returned in a hierarchical format for each offering or standalone discount.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identifier** | **string**| The identifier of a contract rule set for which you wish to return the referenced specification items which can converted into selection rules. This identifier can be provided either as a guid value or as a value which matches the pattern ^[a-z0-9]{6}$ | 
 **optional** | ***NegotiationContractsApiApiContractsIdentifierSelectableItemsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NegotiationContractsApiApiContractsIdentifierSelectableItemsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientRequestSource** | **optional.String**| The client request header that identifies which client had made a request to CPQ Server. | 
 **cpqLanguage** | **optional.String**| The language code (in ISO 639-1 format) that error messages are to be translated to for returning back to the client. If not supplied it defaults to en-GB | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiContractsIdentifierStatusStatusGet**
> ApiContractsIdentifierStatusStatusGet(ctx, identifier, status, optional)


Retrieve the details on the contract rule set which matches the supplied identifier and status.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identifier** | **string**| The identifier of a contract rule set - can be provided either as a guid value or as a value which matches the pattern ^[a-z0-9]{6}$ | 
  **status** | **string**| The status value which matching rule sets should have | 
 **optional** | ***NegotiationContractsApiApiContractsIdentifierStatusStatusGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NegotiationContractsApiApiContractsIdentifierStatusStatusGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientRequestSource** | **optional.String**| The client request header that identifies which client had made a request to CPQ Server. | 
 **cpqLanguage** | **optional.String**| The language code (in ISO 639-1 format) that error messages are to be translated to for returning back to the client. If not supplied it defaults to en-GB | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiContractsPost**
> ApiContractsPost(ctx, fcUser, contract, contentType, optional)


Create a framework contract rule and create negotiated specifications for the referenced offering and standalone discount specifications

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fcUser** | **string**| The name of the user that is creating the framework contract rule set | 
  **contract** | [**CreateContractModel**](CreateContractModel.md)| The framework contract rule set in JSON format that is to be written to the database. | 
  **contentType** | **string**| Specifies the format of the data that is being sent to the API. Only JSON data is accepted | 
 **optional** | ***NegotiationContractsApiApiContractsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NegotiationContractsApiApiContractsPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **clientRequestSource** | **optional.String**| The client request header that identifies which client had made a request to CPQ Server. | 
 **cpqLanguage** | **optional.String**| The language code (in ISO 639-1 format) that error messages are to be translated to for returning back to the client. If not supplied it defaults to en-GB | 
 **performTransformation** | **optional.Bool**| Perform transformation of referenced product specifications using the given framework contract into transformed specifications or not. If not supplied defaults to true | 
 **returnNegotiatedSpecifications** | **optional.Bool**| Return the transformed specifications in response or not. If not supplied defaults to false | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiContractsStatusStatusGet**
> ApiContractsStatusStatusGet(ctx, status, optional)


Retrieve the details of all contract rule sets which match the supplied status value

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **status** | **string**| The status value which matching rule sets should have | 
 **optional** | ***NegotiationContractsApiApiContractsStatusStatusGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NegotiationContractsApiApiContractsStatusStatusGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientRequestSource** | **optional.String**| The client request header that identifies which client had made a request to CPQ Server. | 
 **cpqLanguage** | **optional.String**| The language code (in ISO 639-1 format) that error messages are to be translated to for returning back to the client. If not supplied it defaults to en-GB | 
 **fullDocument** | **optional.Bool**| If true then the contract rule set will return the details of the offerings and selection rules in the response otherwise just the top level top level rule set details will be returned | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

