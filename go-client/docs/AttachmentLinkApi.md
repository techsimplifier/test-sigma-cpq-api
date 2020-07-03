# \AttachmentLinkApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiAttachmentLinkAttachmentLinkIdDelete**](AttachmentLinkApi.md#ApiAttachmentLinkAttachmentLinkIdDelete) | **Delete** /api/attachmentLink/{attachmentLinkId} | 


# **ApiAttachmentLinkAttachmentLinkIdDelete**
> interface{} ApiAttachmentLinkAttachmentLinkIdDelete(ctx, attachmentLinkId, optional)


Delete the specified `AttachmentLink` object from the DB 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **attachmentLinkId** | **string**| ID of the attachment link to delete | 
 **optional** | ***AttachmentLinkApiApiAttachmentLinkAttachmentLinkIdDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AttachmentLinkApiApiAttachmentLinkAttachmentLinkIdDeleteOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientRequestSource** | **optional.String**| The client request header that identifies which client had made a request to CPQ Server. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

