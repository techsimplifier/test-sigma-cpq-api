# \WorkflowApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiWorkflowStateMachineStatesGet**](WorkflowApi.md#ApiWorkflowStateMachineStatesGet) | **Get** /api/workflow/{stateMachine}/states | 
[**ApiWorkflowStateMachineStatusStatusTransitionsGet**](WorkflowApi.md#ApiWorkflowStateMachineStatusStatusTransitionsGet) | **Get** /api/workflow/{stateMachine}/status/{status}/transitions | 
[**ApiWorkflowStateMachineTransitionTransitionPut**](WorkflowApi.md#ApiWorkflowStateMachineTransitionTransitionPut) | **Put** /api/workflow/{stateMachine}/transition/{transition} | 


# **ApiWorkflowStateMachineStatesGet**
> ApiWorkflowStateMachineStatesGet(ctx, stateMachine, optional)


For the specified state machine, retrieve a listing of all of the statuses that have been registered and also include a count of all framework contracts which are currently of those statuses

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **stateMachine** | **string**| The name of the state machine whose statuses are to be queried | 
 **optional** | ***WorkflowApiApiWorkflowStateMachineStatesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WorkflowApiApiWorkflowStateMachineStatesGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientRequestSource** | **optional.String**| The client request header that identifies which client had made a request to CPQ Server. | 
 **cpqLanguage** | **optional.String**| The language code (in ISO 639-1 format) that error messages are to be translated to for returning back to the client. If not supplied it defaults to en-GB | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiWorkflowStateMachineStatusStatusTransitionsGet**
> ApiWorkflowStateMachineStatusStatusTransitionsGet(ctx, stateMachine, status, optional)


For the supplied state machine and status return the details on which transitions are possible to select.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **stateMachine** | **string**| The name of the state machine to find valid transitions for | 
  **status** | **string**| The status from which to determine the next allowable transition | 
 **optional** | ***WorkflowApiApiWorkflowStateMachineStatusStatusTransitionsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WorkflowApiApiWorkflowStateMachineStatusStatusTransitionsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientRequestSource** | **optional.String**| The client request header that identifies which client had made a request to CPQ Server. | 
 **cpqLanguage** | **optional.String**| The language code (in ISO 639-1 format) that error messages are to be translated to for returning back to the client. If not supplied it defaults to en-GB | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiWorkflowStateMachineTransitionTransitionPut**
> SuccessResponse200 ApiWorkflowStateMachineTransitionTransitionPut(ctx, contentType, stateMachine, transition, workflowRequestBody, optional)


Execute a transition operation on a state machine to change the status of the documents which are referenced by the identifiers in the request body

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contentType** | **string**| Specifies the format of the data that is being sent to the API. Only JSON data is accepted | 
  **stateMachine** | **string**| The name of the state machine which represents the documents whose status you wish to change | 
  **transition** | **string**| The name of the transition which is present on the state machine which is responsible for changing the document status | 
  **workflowRequestBody** | [**WorkflowRequest**](WorkflowRequest.md)| The collection of workflowRequestBody which refer to the documents for their status changed as part of the workflow operation. These workflowRequestBody must have &#39;id&#39; in format (GUID or FC_ID or External OM ID [0-9 digit format]) otherwise an error will be returned .Generalized optional messages body for workflow API transition can be array or array of objects and should contain atleast one element.The messages can be submitted to indicate specific reasons for failure. | 
 **optional** | ***WorkflowApiApiWorkflowStateMachineTransitionTransitionPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WorkflowApiApiWorkflowStateMachineTransitionTransitionPutOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **clientRequestSource** | **optional.String**| The client request header that identifies which client had made a request to CPQ Server. | 
 **cpqLanguage** | **optional.String**| The language code (in ISO 639-1 format) that error messages are to be translated to for returning back to the client. If not supplied it defaults to en-GB | 
 **isSimulation** | **optional.Bool**| Determines whether after determining which identifiers are able to transition the changes are written to the database. If specified as true then no changes are made and if specified as false then if possible the database is updated. If not specified this parameter defaults to false | 
 **allowPartialTransitions** | **optional.Bool**| If specified as true then at least one identifier has to be able to transition to the target state for the operation to be considered successful. If specified as false then all identifiers have to be able to transition. The default value for this parameter is false. | 

### Return type

[**SuccessResponse200**](successResponse_200.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

