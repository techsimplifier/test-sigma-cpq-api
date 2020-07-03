# \ExternalCallsApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiQuotesQuoteIdExternalLookupapiNameGet**](ExternalCallsApi.md#ApiQuotesQuoteIdExternalLookupapiNameGet) | **Get** /api/quotes/{quoteId}/externalLookup/&#39;apiName&#39; | 
[**ApiQuotesQuoteIdExternalReservationapiNamePost**](ExternalCallsApi.md#ApiQuotesQuoteIdExternalReservationapiNamePost) | **Post** /api/quotes/{quoteId}/externalReservation/&#39;apiName&#39; | 
[**ApiQuotesQuoteIdExternalVerificationapiNamePost**](ExternalCallsApi.md#ApiQuotesQuoteIdExternalVerificationapiNamePost) | **Post** /api/quotes/{quoteId}/externalVerification/&#39;apiName&#39; | 


# **ApiQuotesQuoteIdExternalLookupapiNameGet**
> ExternalLookupObject ApiQuotesQuoteIdExternalLookupapiNameGet(ctx, quoteId, optional)


Gets the appropriate data from external system by executing registered plugin for the same.<br/>Where 'apiName' would be the name you registered for the externalLookup plugin.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **quoteId** | **string**| entity ID | 
 **optional** | ***ExternalCallsApiApiQuotesQuoteIdExternalLookupapiNameGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExternalCallsApiApiQuotesQuoteIdExternalLookupapiNameGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientRequestSource** | **optional.String**| The client request header that identifies which client had made a request to CPQ Server. | 

### Return type

[**ExternalLookupObject**](ExternalLookupObject.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiQuotesQuoteIdExternalReservationapiNamePost**
> ExternalReservationObject ApiQuotesQuoteIdExternalReservationapiNamePost(ctx, quoteId, reservationValuesArray, optional)


Reserves the appropriate data in an external system by executing registered plugin for the same.<br/>Where 'apiName' would be the name you registered for the externalReservation plugin.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **quoteId** | **string**| entity ID | 
  **reservationValuesArray** | [**ExternalVerificationValueType**](ExternalVerificationValueType.md)| Array of values to be reserved | 
 **optional** | ***ExternalCallsApiApiQuotesQuoteIdExternalReservationapiNamePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExternalCallsApiApiQuotesQuoteIdExternalReservationapiNamePostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientRequestSource** | **optional.String**| The client request header that identifies which client had made a request to CPQ Server. | 

### Return type

[**ExternalReservationObject**](ExternalReservationObject.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiQuotesQuoteIdExternalVerificationapiNamePost**
> ExternalVerificationObject ApiQuotesQuoteIdExternalVerificationapiNamePost(ctx, quoteId, verificationValuesArray, optional)


Posts the verification value array to external system with quote id and returns the verification result.<br>Where 'apiName' would be the name you registered for the externalVerification plugin.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **quoteId** | **string**| entity ID | 
  **verificationValuesArray** | [**ExternalVerificationValueType**](ExternalVerificationValueType.md)| Array of values to be verified | 
 **optional** | ***ExternalCallsApiApiQuotesQuoteIdExternalVerificationapiNamePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExternalCallsApiApiQuotesQuoteIdExternalVerificationapiNamePostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientRequestSource** | **optional.String**| The client request header that identifies which client had made a request to CPQ Server. | 

### Return type

[**ExternalVerificationObject**](ExternalVerificationObject.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

