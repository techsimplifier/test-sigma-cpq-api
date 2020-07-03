# \CatalogServicesProxyApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiCatalogCommercialSpecIdGet**](CatalogServicesProxyApi.md#ApiCatalogCommercialSpecIdGet) | **Get** /api/catalog/commercialSpec/{id} | 
[**ApiCatalogFindPost**](CatalogServicesProxyApi.md#ApiCatalogFindPost) | **Post** /api/catalog/find | 
[**ApiCatalogSpecIdGet**](CatalogServicesProxyApi.md#ApiCatalogSpecIdGet) | **Get** /api/catalog/spec/{id} | 


# **ApiCatalogCommercialSpecIdGet**
> interface{} ApiCatalogCommercialSpecIdGet(ctx, id, optional)


A proxy call that will call the catalog services '/sigma/find/productSpecification/commercial/guid/' endpoint with the given id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| entity ID | 
 **optional** | ***CatalogServicesProxyApiApiCatalogCommercialSpecIdGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CatalogServicesProxyApiApiCatalogCommercialSpecIdGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientRequestSource** | **optional.String**| The client request header that identifies which client had made a request to CPQ Server. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiCatalogFindPost**
> interface{} ApiCatalogFindPost(ctx, optional)


A proxy call that will call the catalog services '/sigma/find/filtered' endpoint with the given request body

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CatalogServicesProxyApiApiCatalogFindPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CatalogServicesProxyApiApiCatalogFindPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientRequestSource** | **optional.String**| The client request header that identifies which client had made a request to CPQ Server. | 
 **filters** | [**optional.Interface of interface{}**](interface{}.md)| Filter options | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiCatalogSpecIdGet**
> interface{} ApiCatalogSpecIdGet(ctx, id, optional)


A proxy call that will call the catalog services '/sigma/find/productSpecification/guid/' endpoint with the given id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| entity ID | 
 **optional** | ***CatalogServicesProxyApiApiCatalogSpecIdGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CatalogServicesProxyApiApiCatalogSpecIdGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientRequestSource** | **optional.String**| The client request header that identifies which client had made a request to CPQ Server. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

