# \ConfigurationApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiConfigurationStepNavigatorDetailsTypeIdGet**](ConfigurationApi.md#ApiConfigurationStepNavigatorDetailsTypeIdGet) | **Get** /api/configuration/stepNavigatorDetails/{type}/{id} | 


# **ApiConfigurationStepNavigatorDetailsTypeIdGet**
> GetStepNavigatorDetails ApiConfigurationStepNavigatorDetailsTypeIdGet(ctx, type_, id)


This API is responsible for fetching step navigator details based on 'type' and 'id' from database

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **type_** | **string**| The unique identifier type should be quote/order | 
  **id** | **string**| The unique identifier of quote or order to be retrieved | 

### Return type

[**GetStepNavigatorDetails**](GetStepNavigatorDetails.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

