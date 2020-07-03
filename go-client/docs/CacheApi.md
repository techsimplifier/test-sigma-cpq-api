# \CacheApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiCacheCacheTypeWarmupIdentifierGet**](CacheApi.md#ApiCacheCacheTypeWarmupIdentifierGet) | **Get** /api/cache/{cacheType}/warmup/{identifier} | 


# **ApiCacheCacheTypeWarmupIdentifierGet**
> ApiCacheCacheTypeWarmupIdentifierGet(ctx, cacheType, identifier)


This API is responsible for pre-warming the CPQ caches, use the identifier parameter to specify the item to pre-warm

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cacheType** | **string**| The type of cache to warm, currently only &#39;specification is supported&#39; | 
  **identifier** | **string**| The identifier to pre-warm | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

