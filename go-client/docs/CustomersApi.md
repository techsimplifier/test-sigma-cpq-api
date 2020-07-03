# \CustomersApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiCustomersCustomerIdGet**](CustomersApi.md#ApiCustomersCustomerIdGet) | **Get** /api/customers/{customerId} | 
[**ApiCustomersCustomerIdLocationsPost**](CustomersApi.md#ApiCustomersCustomerIdLocationsPost) | **Post** /api/customers/{customerId}/locations | 
[**ApiCustomersCustomerIdOrdersGet**](CustomersApi.md#ApiCustomersCustomerIdOrdersGet) | **Get** /api/customers/{customerId}/orders | 
[**ApiCustomersCustomerIdPortfolioByLocationPost**](CustomersApi.md#ApiCustomersCustomerIdPortfolioByLocationPost) | **Post** /api/customers/{customerId}/portfolioByLocation | 
[**ApiCustomersCustomerIdPortfolioGet**](CustomersApi.md#ApiCustomersCustomerIdPortfolioGet) | **Get** /api/customers/{customerId}/portfolio | 
[**ApiCustomersCustomerIdPortfolioValidateGet**](CustomersApi.md#ApiCustomersCustomerIdPortfolioValidateGet) | **Get** /api/customers/{customerId}/portfolio/validate | 
[**ApiCustomersCustomerIdQuotesGet**](CustomersApi.md#ApiCustomersCustomerIdQuotesGet) | **Get** /api/customers/{customerId}/quotes | 


# **ApiCustomersCustomerIdGet**
> CustomerAccount ApiCustomersCustomerIdGet(ctx, customerId, optional)


Gets the Customer details based on supplied customerId 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customerId** | **string**| Customer ID | 
 **optional** | ***CustomersApiApiCustomersCustomerIdGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomersApiApiCustomersCustomerIdGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientRequestSource** | **optional.String**| The client request header that identifies which client had made a request to CPQ Server. | 

### Return type

[**CustomerAccount**](CustomerAccount.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiCustomersCustomerIdLocationsPost**
> Locations ApiCustomersCustomerIdLocationsPost(ctx, customerId, optional)


Gets the `Locations` for a given customer based on the supplied query parameter. This endpoint will use the configured `ContextProvider` and make a call to the registered `GetLocations` method. The endpoint will directly return any errors returned from the plugin.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customerId** | **string**| Customer ID | 
 **optional** | ***CustomersApiApiCustomersCustomerIdLocationsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomersApiApiCustomersCustomerIdLocationsPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientRequestSource** | **optional.String**| The client request header that identifies which client had made a request to CPQ Server. | 
 **locationQueryParameters** | [**optional.Interface of LocationQueryParameters**](LocationQueryParameters.md)| The query options for location, includes &#39;filter&#39; and &#39;fields&#39; options | 

### Return type

[**Locations**](Locations.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiCustomersCustomerIdOrdersGet**
> OrderSummary ApiCustomersCustomerIdOrdersGet(ctx, customerId, optional)


Calls out to the configured order management system to get the orders for a given customer, this requires an OrderAdapter plugin to be registered.<br/>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customerId** | **string**| Customer ID | 
 **optional** | ***CustomersApiApiCustomersCustomerIdOrdersGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomersApiApiCustomersCustomerIdOrdersGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientRequestSource** | **optional.String**| The client request header that identifies which client had made a request to CPQ Server. | 
 **page** | **optional.String**| Page count | 
 **count** | **optional.String**| Number of records per page | 
 **sort** | [**optional.Interface of map[string]string**](string.md)| Sort field and direction | 

### Return type

[**OrderSummary**](OrderSummary.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiCustomersCustomerIdPortfolioByLocationPost**
> PortfolioByLocationResponse ApiCustomersCustomerIdPortfolioByLocationPost(ctx, customerId)


Gets the Customer Portfolio by location for the given customer based on supplied CustomerID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customerId** | **string**| Customer ID | 

### Return type

[**PortfolioByLocationResponse**](PortfolioByLocationResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiCustomersCustomerIdPortfolioGet**
> PortfolioItemSummaryResponse ApiCustomersCustomerIdPortfolioGet(ctx, customerId, optional)


Gets the Customer Portfolio for the given customer based on supplied CustomerID.<br/> if Portfolio type is supplied this end point gets the list of Portfolio item object containing name, id and entityID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customerId** | **string**| Customer ID | 
 **optional** | ***CustomersApiApiCustomersCustomerIdPortfolioGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomersApiApiCustomersCustomerIdPortfolioGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientRequestSource** | **optional.String**| The client request header that identifies which client had made a request to CPQ Server. | 
 **portfolioType** | **optional.String**| Used to retrieve portfolio summary by using portfolioType as &#39;summary&#39; | 

### Return type

[**PortfolioItemSummaryResponse**](PortfolioItemSummaryResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiCustomersCustomerIdPortfolioValidateGet**
> PortfolioValidation ApiCustomersCustomerIdPortfolioValidateGet(ctx, customerId, optional)


Retrieves the customerPortfolio from the contextProvider plugin and validates it returning any validation errors.<br/>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customerId** | **string**| Customer ID | 
 **optional** | ***CustomersApiApiCustomersCustomerIdPortfolioValidateGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomersApiApiCustomersCustomerIdPortfolioValidateGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientRequestSource** | **optional.String**| The client request header that identifies which client had made a request to CPQ Server. | 

### Return type

[**PortfolioValidation**](PortfolioValidation.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiCustomersCustomerIdQuotesGet**
> QuoteSummary ApiCustomersCustomerIdQuotesGet(ctx, customerId, optional)


Retrieves list of quotes associated to the provided unique CustomerID. Can pass optional filterTypes parameter.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customerId** | **string**| Customer ID | 
 **optional** | ***CustomersApiApiCustomersCustomerIdQuotesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomersApiApiCustomersCustomerIdQuotesGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientRequestSource** | **optional.String**| The client request header that identifies which client had made a request to CPQ Server. | 
 **filters** | **optional.String**| Comma separated list of values which will filter for the given quote types. List of available options are: 0 (sales), 1 (opportunity), 2 (change quotes). Leaving blank would omit the filter params and all quote types will be returned. | 

### Return type

[**QuoteSummary**](QuoteSummary.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

