# \StatusApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiDocsGet**](StatusApi.md#ApiDocsGet) | **Get** /api/docs | 
[**ApiStatusGet**](StatusApi.md#ApiStatusGet) | **Get** /api/status | 


# **ApiDocsGet**
> interface{} ApiDocsGet(ctx, )


Gets the current documentation for all of the available REST endpoints 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiStatusGet**
> StatusObject ApiStatusGet(ctx, )


Gets the current status of the REST endpoints 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**StatusObject**](StatusObject.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

