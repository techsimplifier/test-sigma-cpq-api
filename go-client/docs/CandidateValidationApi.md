# \CandidateValidationApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiValidationEvaluateCompatibilityPost**](CandidateValidationApi.md#ApiValidationEvaluateCompatibilityPost) | **Post** /api/validation/evaluateCompatibility | 
[**ApiValidationPost**](CandidateValidationApi.md#ApiValidationPost) | **Post** /api/validation | 


# **ApiValidationEvaluateCompatibilityPost**
> EvaluateValidationResponse ApiValidationEvaluateCompatibilityPost(ctx, candidate, optional)


This API is responsible for evaluating compatibility rules for the given product candidate. This API utilizes catalog services functionality to evaluate compatibility rules and return a new product candidate structure with any violating entities marked as 'NotAvailable'. The compatibility rules that are evaluated by this route are defined in Catalog</b><br/> If a failed response is returned by the API, the responseCode value contains:  responseCode   | Meaning   --- | ----   CSConnectionRequestError | The error occurs when request was sent to Catalog Services, an error occurred whilst making a POST request to catalog services or the request did not supply a valid product candidate  MissingCatalogConnection | This error occurs when the catalog services connection information is missing from the configuration.  MissingDataTransformer   | This error occurs when the data transformer is missing.  MissingCandidate | This error occurs when the route was called without a candidate. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **candidate** | [**ProductCandidate**](ProductCandidate.md)| The candidate JSON to be validated | 
 **optional** | ***CandidateValidationApiApiValidationEvaluateCompatibilityPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CandidateValidationApiApiValidationEvaluateCompatibilityPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientRequestSource** | **optional.String**| The client request header that identifies which client had made a request to CPQ Server. | 
 **quoteId** | **optional.String**| The current quote id | 
 **quoteItemId** | **optional.String**| The current quote item id | 

### Return type

[**EvaluateValidationResponse**](EvaluateValidationResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiValidationPost**
> ValidationResponse ApiValidationPost(ctx, candidate, optional)


Posts `product candidate` objects to be validated 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **candidate** | [**interface{}**](interface{}.md)| The candidate JSON to be validated | 
 **optional** | ***CandidateValidationApiApiValidationPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CandidateValidationApiApiValidationPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientRequestSource** | **optional.String**| The client request header that identifies which client had made a request to CPQ Server. | 

### Return type

[**ValidationResponse**](ValidationResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

