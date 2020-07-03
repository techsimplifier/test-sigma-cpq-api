# \NegotiationCustomerMappingsApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiCustomerMappingsCustomerCustomerIdGet**](NegotiationCustomerMappingsApi.md#ApiCustomerMappingsCustomerCustomerIdGet) | **Get** /api/customerMappings/customer/{customerId} | 
[**ApiCustomerMappingsCustomerCustomerIdStartDateStartDateEndDateEndDateGet**](NegotiationCustomerMappingsApi.md#ApiCustomerMappingsCustomerCustomerIdStartDateStartDateEndDateEndDateGet) | **Get** /api/customerMappings/customer/{customerId}/startDate/{startDate}/endDate/{endDate} | 
[**ApiCustomerMappingsCustomerIdContractIdGet**](NegotiationCustomerMappingsApi.md#ApiCustomerMappingsCustomerIdContractIdGet) | **Get** /api/customerMappings/{customerId}/contractId | 
[**ApiCustomerMappingsCustomerIdIdentifierIdentifierPut**](NegotiationCustomerMappingsApi.md#ApiCustomerMappingsCustomerIdIdentifierIdentifierPut) | **Put** /api/customerMappings/{customerId}/identifier/{identifier} | 
[**ApiCustomerMappingsPost**](NegotiationCustomerMappingsApi.md#ApiCustomerMappingsPost) | **Post** /api/customerMappings | 
[**ApiCustomerMappingsStartDateStartDateEndDateEndDateGet**](NegotiationCustomerMappingsApi.md#ApiCustomerMappingsStartDateStartDateEndDateEndDateGet) | **Get** /api/customerMappings/startDate/{startDate}/endDate/{endDate} | 
[**ApiCustomersCustomerIdContractsGet**](NegotiationCustomerMappingsApi.md#ApiCustomersCustomerIdContractsGet) | **Get** /api/customers/{customerId}/contracts | 


# **ApiCustomerMappingsCustomerCustomerIdGet**
> ApiCustomerMappingsCustomerCustomerIdGet(ctx, customerId, contentType, optional)


Get all customer mappings by customerId

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customerId** | **string**| The identifier of customer mapping | 
  **contentType** | **string**| Specifies the format of the data that is being sent to the API. Only JSON data is accepted | 
 **optional** | ***NegotiationCustomerMappingsApiApiCustomerMappingsCustomerCustomerIdGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NegotiationCustomerMappingsApiApiCustomerMappingsCustomerCustomerIdGetOpts struct

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

# **ApiCustomerMappingsCustomerCustomerIdStartDateStartDateEndDateEndDateGet**
> ApiCustomerMappingsCustomerCustomerIdStartDateStartDateEndDateEndDateGet(ctx, customerId, startDate, endDate, contentType, optional)


Get customer mappings by customerId and date range

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customerId** | **string**| The identifier of customer mapping | 
  **startDate** | **string**| Start date range of query in ISO 8601 format where customer mappings in effective | 
  **endDate** | **string**| End date range of query in ISO 8601 format where customer mappings in effective | 
  **contentType** | **string**| Specifies the format of the data that is being sent to the API. Only JSON data is accepted | 
 **optional** | ***NegotiationCustomerMappingsApiApiCustomerMappingsCustomerCustomerIdStartDateStartDateEndDateEndDateGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NegotiationCustomerMappingsApiApiCustomerMappingsCustomerCustomerIdStartDateStartDateEndDateEndDateGetOpts struct

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

# **ApiCustomerMappingsCustomerIdContractIdGet**
> ApiCustomerMappingsCustomerIdContractIdGet(ctx, customerId, startDate, contentType, optional)


Get contract rule set identifier of customer mappings by customer Id and date

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customerId** | **string**| The identifier of customer mapping | 
  **startDate** | **string**| Start date in ISO 8601 format where customer mapping is in effective | 
  **contentType** | **string**| Specifies the format of the data that is being sent to the API. Only JSON data is accepted | 
 **optional** | ***NegotiationCustomerMappingsApiApiCustomerMappingsCustomerIdContractIdGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NegotiationCustomerMappingsApiApiCustomerMappingsCustomerIdContractIdGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **clientRequestSource** | **optional.String**| The client request header that identifies which client had made a request to CPQ Server. | 
 **cpqLanguage** | **optional.String**| The language code (in ISO 639-1 format) that error messages are to be translated to for returning back to the client. If not supplied it defaults to en-GB | 
 **isContractId** | **optional.Bool**| If set to true, return contract Id as contract rule set identifier, if set to false, return fc_Id as contract rule set identifier.  If value is not specified, default to true | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiCustomerMappingsCustomerIdIdentifierIdentifierPut**
> ApiCustomerMappingsCustomerIdIdentifierIdentifierPut(ctx, fcUser, customerMapping, customerId, identifier, contentType, optional)


Update customer mappings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fcUser** | **string**| The name of the user that is updating the custom mapping | 
  **customerMapping** | [**CustomerMappingModel**](CustomerMappingModel.md)| Customer mapping that is to be written to the database. | 
  **customerId** | **string**| The identifier of customer mapping to be updated | 
  **identifier** | **string**| The identifier of a contract rule set that assoicated with the customer mapping to be updated - can be provided either as a guid value or as a value which matches the pattern ^[a-z0-9]{6}$ | 
  **contentType** | **string**| Specifies the format of the data that is being sent to the API. Only JSON data is accepted | 
 **optional** | ***NegotiationCustomerMappingsApiApiCustomerMappingsCustomerIdIdentifierIdentifierPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NegotiationCustomerMappingsApiApiCustomerMappingsCustomerIdIdentifierIdentifierPutOpts struct

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

# **ApiCustomerMappingsPost**
> ApiCustomerMappingsPost(ctx, customerMappings, contentType, optional)


Bulk create customer mappings that maps customer Id to contract rule set identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customerMappings** | [**CreateCustomerMappingsRequest**](CreateCustomerMappingsRequest.md)| Array of customer mapping format that is to be written to the database. | 
  **contentType** | **string**| Specifies the format of the data that is being sent to the API. Only JSON data is accepted | 
 **optional** | ***NegotiationCustomerMappingsApiApiCustomerMappingsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NegotiationCustomerMappingsApiApiCustomerMappingsPostOpts struct

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

# **ApiCustomerMappingsStartDateStartDateEndDateEndDateGet**
> ApiCustomerMappingsStartDateStartDateEndDateEndDateGet(ctx, startDate, endDate, contentType, optional)


Get customer mappings by date range

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **startDate** | **string**| Start date range of query in ISO 8601 format where customer mappings in effective | 
  **endDate** | **string**| End date range of query in ISO 8601 format where customer mappings in effective | 
  **contentType** | **string**| Specifies the format of the data that is being sent to the API. Only JSON data is accepted | 
 **optional** | ***NegotiationCustomerMappingsApiApiCustomerMappingsStartDateStartDateEndDateEndDateGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NegotiationCustomerMappingsApiApiCustomerMappingsStartDateStartDateEndDateEndDateGetOpts struct

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

# **ApiCustomersCustomerIdContractsGet**
> ApiCustomersCustomerIdContractsGet(ctx, customerId, contentType, optional)


Get all customer mappings and summary contracts by customerId

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customerId** | **string**| The identifier of customer mapping | 
  **contentType** | **string**| Specifies the format of the data that is being sent to the API. Only JSON data is accepted | 
 **optional** | ***NegotiationCustomerMappingsApiApiCustomersCustomerIdContractsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NegotiationCustomerMappingsApiApiCustomersCustomerIdContractsGetOpts struct

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

