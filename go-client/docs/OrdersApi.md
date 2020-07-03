# \OrdersApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiOrdersOrderIdCancelPost**](OrdersApi.md#ApiOrdersOrderIdCancelPost) | **Post** /api/orders/{orderId}/cancel | 
[**ApiOrdersOrderIdEnrichmentTasksGet**](OrdersApi.md#ApiOrdersOrderIdEnrichmentTasksGet) | **Get** /api/orders/{orderId}/enrichmentTasks | 
[**ApiOrdersOrderIdGet**](OrdersApi.md#ApiOrdersOrderIdGet) | **Get** /api/orders/{orderId} | 
[**ApiOrdersOrderIdToChangeQuotePost**](OrdersApi.md#ApiOrdersOrderIdToChangeQuotePost) | **Post** /api/orders/{orderId}/toChangeQuote | 
[**ApiOrdersQuoteIdGetOrderItemSummariesGet**](OrdersApi.md#ApiOrdersQuoteIdGetOrderItemSummariesGet) | **Get** /api/orders/{quoteId}/getOrderItemSummaries/ | 
[**ApiOrdersQuoteIdGetOrderItemSummariesPost**](OrdersApi.md#ApiOrdersQuoteIdGetOrderItemSummariesPost) | **Post** /api/orders/{quoteId}/getOrderItemSummaries/ | 
[**ApiOrdersQuoteIdGetOrderItemSummariesQuoteItemIdGet**](OrdersApi.md#ApiOrdersQuoteIdGetOrderItemSummariesQuoteItemIdGet) | **Get** /api/orders/{quoteId}/getOrderItemSummaries/{quoteItemId} | 
[**ApiOrdersQuoteIdSubmitPost**](OrdersApi.md#ApiOrdersQuoteIdSubmitPost) | **Post** /api/orders/{quoteId}/submit | 
[**ApiOrdersQuoteIdTransformGet**](OrdersApi.md#ApiOrdersQuoteIdTransformGet) | **Get** /api/orders/{quoteId}/transform | 
[**ApiOrdersQuoteIdTransformOssjGet**](OrdersApi.md#ApiOrdersQuoteIdTransformOssjGet) | **Get** /api/orders/{quoteId}/transform/ossj | 
[**ApiOrdersQuoteIdTransformOssjWithOrderCharacteristicsGet**](OrdersApi.md#ApiOrdersQuoteIdTransformOssjWithOrderCharacteristicsGet) | **Get** /api/orders/{quoteId}/transform/ossjWithOrderCharacteristics | 


# **ApiOrdersOrderIdCancelPost**
> OrderCancellationResponse ApiOrdersOrderIdCancelPost(ctx, orderId, optional)


This endpoint allows the cancellation of an order that is not past PONR(Point of No Return) on an external Order Management System.  In order to use this endpoint, an Order Adapter plugin must be registered than implements the `IOrderAdapter` interface. This can be found in the `sigma-cpq.d.ts` file within the `sigma-cpq-typings-bundle` npm module.  When the endpoint is called, CPQ first retrieves the order from the order management system to ensure it is not past PONR, then POSTs a cancellation request to the Order Management System. The response is transformed into an 'IOrderCancellationResponse' structure and returned to the callee. If there is a server error, a 500 response will be returned. Any errors returned from the Order Management System will be transformed and returned to the callee in a readable format.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orderId** | **string**| Order ID | 
 **optional** | ***OrdersApiApiOrdersOrderIdCancelPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrdersApiApiOrdersOrderIdCancelPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientRequestSource** | **optional.String**| The client request header that identifies which client had made a request to CPQ Server. | 

### Return type

[**OrderCancellationResponse**](OrderCancellationResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiOrdersOrderIdEnrichmentTasksGet**
> GetOrderEnrichmentTask200 ApiOrdersOrderIdEnrichmentTasksGet(ctx, orderId)


This API is responsible for fetching all the enrichment tasks by order ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orderId** | **string**| The unique identifier of the orderId of the converted quote | 

### Return type

[**GetOrderEnrichmentTask200**](GetOrderEnrichmentTask_200.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiOrdersOrderIdGet**
> GenericOrder ApiOrdersOrderIdGet(ctx, orderId, optional)


Retrieves an order related detailed information from the order management system with the unique identifier provided.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orderId** | **string**| The unique identifier of the order to be retrieved | 
 **optional** | ***OrdersApiApiOrdersOrderIdGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrdersApiApiOrdersOrderIdGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientRequestSource** | **optional.String**| The client request header that identifies which client had made a request to CPQ Server. | 

### Return type

[**GenericOrder**](GenericOrder.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiOrdersOrderIdToChangeQuotePost**
> Quote ApiOrdersOrderIdToChangeQuotePost(ctx, orderId, optional)


This route will require an OrderAdapter plugin be registered and will get the specified order using the plugin and create a change quote from it.  If the 'returnQuote=true' query param is supplied, the change quote will be returned, otherwise it will just return the id of the change quote.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orderId** | **string**| Order ID | 
 **optional** | ***OrdersApiApiOrdersOrderIdToChangeQuotePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrdersApiApiOrdersOrderIdToChangeQuotePostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **returnquote** | **optional.String**| A boolean value used to specify if you want the change quote returned. If false or param is omitted then just the change quote id is returned | 
 **include** | **optional.String**| This is the CRM&#39;s reference of the opportunity. | 
 **opportunityId** | **optional.String**| Comma separated list of parameters to return on the quote. List of available options are: specification, candidate, validation, pricing, quote. Leaving it blank would omit the include params. | 
 **clientRequestSource** | **optional.String**| The client request header that identifies which client had made a request to CPQ Server. | 

### Return type

[**Quote**](Quote.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiOrdersQuoteIdGetOrderItemSummariesGet**
> IGenericOrderItemSummary ApiOrdersQuoteIdGetOrderItemSummariesGet(ctx, quoteId, optional)


Get all order item review summaries associated with the provided quoteId

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **quoteId** | **string**| Quote ID | 
 **optional** | ***OrdersApiApiOrdersQuoteIdGetOrderItemSummariesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrdersApiApiOrdersQuoteIdGetOrderItemSummariesGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientRequestSource** | **optional.String**| The client request header that identifies which client had made a request to CPQ Server. | 

### Return type

[**IGenericOrderItemSummary**](IGenericOrderItemSummary.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiOrdersQuoteIdGetOrderItemSummariesPost**
> UiGridOrderItemSummariesResponse ApiOrdersQuoteIdGetOrderItemSummariesPost(ctx, quoteId, orderItemSummariesParameters)


Get all order item review summaries associated with the provided quoteId in UI grid format

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **quoteId** | **string**| Quote ID | 
  **orderItemSummariesParameters** | [**OrderItemSummariesParameters**](OrderItemSummariesParameters.md)| The query parameters for orderItemSummaries | 

### Return type

[**UiGridOrderItemSummariesResponse**](UiGridOrderItemSummariesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiOrdersQuoteIdGetOrderItemSummariesQuoteItemIdGet**
> IGenericOrderItemSummary ApiOrdersQuoteIdGetOrderItemSummariesQuoteItemIdGet(ctx, quoteId, quoteItemId, optional)


Get the exact order item review summary using provided quoteId and quoteItemId

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **quoteId** | **string**| Quote ID | 
  **quoteItemId** | **string**| Quote Item ID | 
 **optional** | ***OrdersApiApiOrdersQuoteIdGetOrderItemSummariesQuoteItemIdGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrdersApiApiOrdersQuoteIdGetOrderItemSummariesQuoteItemIdGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientRequestSource** | **optional.String**| The client request header that identifies which client had made a request to CPQ Server. | 

### Return type

[**IGenericOrderItemSummary**](IGenericOrderItemSummary.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiOrdersQuoteIdSubmitPost**
> OrderSubmissionResponse ApiOrdersQuoteIdSubmitPost(ctx, quoteId, orderSubmissionRequest, optional)


This endpoint allows the submission of a quote to an external Order Management system.   In order to use this endpoint an order adapter must be registered than implements the `IOrderAdapter` interface.   When the endpoint is called CPQ generates an order candidate from the quote specified by the given quoteID, this order candidate is then passed to the implementer defined order adapter's transform method which is responsible for transforming the generic order into the format required by whichever order management system the order will be submitted to.   After transformatiom the transformed order structure is supplied to the submit method on the same order adapter, which handles the actual submission process.   The following is a list of errors that may be returned from the Order Adapter `submit()` method, for either a new or supplemental order:  ErrorId | Meaning   --- | ----   OrderCouldNotBeCancelled | The specified order could not be cancelled   OrderCouldNotBeSuspended | The specified order could not be suspended   OrderCouldNotBeResumed | The specified order could not be resumed   OrderStateInvalid | The specified orders state is invalid   OrderNotFound | An order could not be found with the specified ID   OrdersNotFound | No orders could be found   OrderContainsDuplicateExtKeys | The specified order contains a duplicated external key   OrderItemNotFoundOnOrder | The specified order contains an item that could not be found   OrderSpecNotFound | The spec referenced by the supplied order could not be found   OrderAlreadyExists | The supplied order key already exists   SupplementalOrderNotValid | The supplied supplemental order is invalid   OrderCouldNotBeProcessed | There was an error processing the supplied order.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **quoteId** | **string**| Supply a valid guid for quoteId and it should be present in the CPQ system.&lt;br/&gt;Example:841b8d04-2c18-44c4-aa18-a3ae0fc17558 | 
  **orderSubmissionRequest** | [**OrderSubmissionRequest**](OrderSubmissionRequest.md)| last updated date of quote | 
 **optional** | ***OrdersApiApiOrdersQuoteIdSubmitPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrdersApiApiOrdersQuoteIdSubmitPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientRequestSource** | **optional.String**| The client request header that identifies which client had made a request to CPQ Server. | 

### Return type

[**OrderSubmissionResponse**](OrderSubmissionResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiOrdersQuoteIdTransformGet**
> ApiOrdersQuoteIdTransformGet(ctx, quoteId, optional)


This route requires that an OrderAdapter plugin be registered, this route will return the result of the transformation carried out by the transform function of the registered OrderAdapter

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **quoteId** | **string**| Quote ID | 
 **optional** | ***OrdersApiApiOrdersQuoteIdTransformGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrdersApiApiOrdersQuoteIdTransformGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientRequestSource** | **optional.String**| The client request header that identifies which client had made a request to CPQ Server. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiOrdersQuoteIdTransformOssjGet**
> ApiOrdersQuoteIdTransformOssjGet(ctx, quoteId, optional)


This route is optional and requires the plug-ins to be installed    Performs the custom OOSJ transformations on the supplied quote, by calling the registered plugin

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **quoteId** | **string**| Quote ID | 
 **optional** | ***OrdersApiApiOrdersQuoteIdTransformOssjGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrdersApiApiOrdersQuoteIdTransformOssjGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientRequestSource** | **optional.String**| The client request header that identifies which client had made a request to CPQ Server. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiOrdersQuoteIdTransformOssjWithOrderCharacteristicsGet**
> ApiOrdersQuoteIdTransformOssjWithOrderCharacteristicsGet(ctx, quoteId, optional)


This route is optional and requires the plug-ins to be installed    Performs the custom OOSJ transformations including adding the characteristic data, on the supplied quote, by calling the registered plugin

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **quoteId** | **string**| Quote ID | 
 **optional** | ***OrdersApiApiOrdersQuoteIdTransformOssjWithOrderCharacteristicsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrdersApiApiOrdersQuoteIdTransformOssjWithOrderCharacteristicsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientRequestSource** | **optional.String**| The client request header that identifies which client had made a request to CPQ Server. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

