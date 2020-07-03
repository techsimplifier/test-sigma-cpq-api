# \PricingApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiQuotesQuoteIdPriceQuotePost**](PricingApi.md#ApiQuotesQuoteIdPriceQuotePost) | **Post** /api/quotes/{quoteId}/priceQuote | 
[**ApiQuotesQuoteIdPricingPost**](PricingApi.md#ApiQuotesQuoteIdPricingPost) | **Post** /api/quotes/{quoteId}/pricing | 


# **ApiQuotesQuoteIdPriceQuotePost**
> Quote ApiQuotesQuoteIdPriceQuotePost(ctx, quoteId, quote, optional)


<br/>This API is responsible for pricing the modified Quote Item in the context of a Quote by utilizing the underlying Catalog Services using OrderCandidate pricing services with additional behaviour including the execution of any registered CPQ Pricing Plugins.Importantly the API does not save or commit the passed Quote to the database it is simply priced and returned.<br/><br/> If a failed response is returned by the API, the responseCode value contains:  responseCode | Meaning   --- | ----   CSConnectionRequestError | The request was sent to Catalog Services and failed in this component.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **quoteId** | **string**| Supply a valid guid for quoteId and it should be present in the CPQ system &lt;br/&gt;Example:841b8d04-2c18-44c4-aa18-a3ae0fc17558 | 
  **quote** | [**Quote**](Quote.md)| The modified quote item whose price you wish to obtain. | 
 **optional** | ***PricingApiApiQuotesQuoteIdPriceQuotePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PricingApiApiQuotesQuoteIdPriceQuotePostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientRequestSource** | **optional.String**| The client request header that identifies which client had made a request to CPQ Server. | 

### Return type

[**Quote**](Quote.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiQuotesQuoteIdPricingPost**
> ICpqPricingResponse ApiQuotesQuoteIdPricingPost(ctx, quoteId, productCandidate, optional)


<br/>This API is responsible for pricing a Product Candidate in the context of a Quote by utilizing the underlying Catalog Services pricing services with additional behaviour including the execution of any registered CPQ Pricing Plugins.The pricing response consist of summary of all the charges that the consumer has been charged.Importantly the API does not save or commit the passed Product Candidate to the database it is simply priced and returned.<br/><br/> If a failed response is returned by the API, the responseCode value contains:  responseCode | Meaning   --- | ----   CSConnectionRequestError | The request was sent to Catalog Services and failed in this component.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **quoteId** | **string**| Supply a valid guid for quoteId and it should be present in the CPQ system &lt;br/&gt;Example:841b8d04-2c18-44c4-aa18-a3ae0fc17558 | 
  **productCandidate** | [**ProductCandidateForPricing**](ProductCandidateForPricing.md)| The product candidate whose price you wish to obtain. | 
 **optional** | ***PricingApiApiQuotesQuoteIdPricingPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PricingApiApiQuotesQuoteIdPricingPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientRequestSource** | **optional.String**| The client request header that identifies which client had made a request to CPQ Server. | 
 **cpqLanguage** | **optional.String**| The language code (in ISO 639-1 format) that error messages are to be translated to for returning back to the client. If not supplied it defaults to en-GB&lt;br/&gt;Example:en-US | 

### Return type

[**ICpqPricingResponse**](ICpqPricingResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

