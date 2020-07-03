# \SalesApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiSalesNegotiatedSpecificationsDiscountsGet**](SalesApi.md#ApiSalesNegotiatedSpecificationsDiscountsGet) | **Get** /api/sales/negotiatedSpecifications/discounts | 
[**ApiSalesNegotiatedSpecificationsOfferingsGet**](SalesApi.md#ApiSalesNegotiatedSpecificationsOfferingsGet) | **Get** /api/sales/negotiatedSpecifications/offerings | 
[**ApiSalesOfferingsGet**](SalesApi.md#ApiSalesOfferingsGet) | **Get** /api/sales/offerings | 


# **ApiSalesNegotiatedSpecificationsDiscountsGet**
> ApiSalesNegotiatedSpecificationsDiscountsGet(ctx, optional)


Retrieve the Negotiated Specifications for the standalone discounts referenced on a contract rule set.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SalesApiApiSalesNegotiatedSpecificationsDiscountsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SalesApiApiSalesNegotiatedSpecificationsDiscountsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientRequestSource** | **optional.String**| The client request header that identifies which client had made a request to CPQ Server. | 
 **cpqLanguage** | **optional.String**| The language code (in ISO 639-1 format) that error messages are to be translated to for returning back to the client. If not supplied it defaults to en-GB | 
 **identifier** | **optional.String**| The identifier of a contract rule set - can be provided either as a guid value or as a value which matches the pattern ^[a-z0-9]{6}$ | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiSalesNegotiatedSpecificationsOfferingsGet**
> ApiSalesNegotiatedSpecificationsOfferingsGet(ctx, optional)


Retrieve the Negotiated Specifications for the offerings referenced on a contract rule set.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SalesApiApiSalesNegotiatedSpecificationsOfferingsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SalesApiApiSalesNegotiatedSpecificationsOfferingsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientRequestSource** | **optional.String**| The client request header that identifies which client had made a request to CPQ Server. | 
 **cpqLanguage** | **optional.String**| The language code (in ISO 639-1 format) that error messages are to be translated to for returning back to the client. If not supplied it defaults to en-GB | 
 **identifier** | **optional.String**| The identifier of a contract rule set - can be provided either as a guid value or as a value which matches the pattern ^[a-z0-9]{6}$ | 
 **entityId** | **optional.String**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiSalesOfferingsGet**
> ApiSalesOfferingsGet(ctx, optional)


Retrieves the base specifications for those offerings which are referenced as part of a contract rule set. If the identifier is not supplied then instead all of the Packages are retrieved from Catalog Services restricted to contain only their top level information.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SalesApiApiSalesOfferingsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SalesApiApiSalesOfferingsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientRequestSource** | **optional.String**| The client request header that identifies which client had made a request to CPQ Server. | 
 **cpqLanguage** | **optional.String**| The language code (in ISO 639-1 format) that error messages are to be translated to for returning back to the client. If not supplied it defaults to en-GB | 
 **identifier** | **optional.String**| The identifier of a contract rule set - can be provided either as a guid value or as a value which matches the pattern ^[a-z0-9]{6}$ | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

