# \CandidateApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiCandidateTransformAfterLocationCopyPost**](CandidateApi.md#ApiCandidateTransformAfterLocationCopyPost) | **Post** /api/candidate/transform/afterLocationCopy | 
[**ApiCandidateTransformAfterLocationSelectionPost**](CandidateApi.md#ApiCandidateTransformAfterLocationSelectionPost) | **Post** /api/candidate/transform/afterLocationSelection | 


# **ApiCandidateTransformAfterLocationCopyPost**
> AfterLocationCopyResponse ApiCandidateTransformAfterLocationCopyPost(ctx, afterLocationCopyRequest)


Endpoint that allows you to make transformations to the product candidate after a location copy operation has occurred. This will call the registered transformation plugin

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **afterLocationCopyRequest** | [**AfterLocationCopyRequest**](AfterLocationCopyRequest.md)| The request parameters to the API | 

### Return type

[**AfterLocationCopyResponse**](AfterLocationCopyResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiCandidateTransformAfterLocationSelectionPost**
> AfterLocationSelectionResponse ApiCandidateTransformAfterLocationSelectionPost(ctx, afterLocationSelectionRequest)


Endpoint that allows you to make transformations to the product candidate after a location selection operation has occurred. This will call the registered transformation plugin

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **afterLocationSelectionRequest** | [**AfterLocationSelectionRequest**](AfterLocationSelectionRequest.md)| The request parameters to the API | 

### Return type

[**AfterLocationSelectionResponse**](AfterLocationSelectionResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

