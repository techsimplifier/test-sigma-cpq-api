# \AuthorizationApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiAuthorizationAdminAclGet**](AuthorizationApi.md#ApiAuthorizationAdminAclGet) | **Get** /api/authorization/admin/acl | 
[**ApiAuthorizationAdminAclResourceIdRolesGet**](AuthorizationApi.md#ApiAuthorizationAdminAclResourceIdRolesGet) | **Get** /api/authorization/admin/acl/{resourceId}/roles | 
[**ApiAuthorizationAdminClaimsClaimIdGet**](AuthorizationApi.md#ApiAuthorizationAdminClaimsClaimIdGet) | **Get** /api/authorization/admin/claims/{claimId} | 
[**ApiAuthorizationAdminClaimsGet**](AuthorizationApi.md#ApiAuthorizationAdminClaimsGet) | **Get** /api/authorization/admin/claims | 
[**ApiAuthorizationAdminClaimsPost**](AuthorizationApi.md#ApiAuthorizationAdminClaimsPost) | **Post** /api/authorization/admin/claims | 
[**ApiAuthorizationAdminResourcesBulkAssignRolesPost**](AuthorizationApi.md#ApiAuthorizationAdminResourcesBulkAssignRolesPost) | **Post** /api/authorization/admin/resources/bulk/assign/roles | 
[**ApiAuthorizationAdminResourcesBulkUnassignRolesPost**](AuthorizationApi.md#ApiAuthorizationAdminResourcesBulkUnassignRolesPost) | **Post** /api/authorization/admin/resources/bulk/unassign/roles | 
[**ApiAuthorizationAdminResourcesResourceIdRolesPut**](AuthorizationApi.md#ApiAuthorizationAdminResourcesResourceIdRolesPut) | **Put** /api/authorization/admin/resources/{resourceId}/roles | 
[**ApiAuthorizationAdminResourcesResourceIdRolesRemovePut**](AuthorizationApi.md#ApiAuthorizationAdminResourcesResourceIdRolesRemovePut) | **Put** /api/authorization/admin/resources/{resourceId}/roles/remove | 
[**ApiAuthorizationAdminRolesGet**](AuthorizationApi.md#ApiAuthorizationAdminRolesGet) | **Get** /api/authorization/admin/roles | 
[**ApiAuthorizationAdminRolesPost**](AuthorizationApi.md#ApiAuthorizationAdminRolesPost) | **Post** /api/authorization/admin/roles | 
[**ApiAuthorizationAdminRolesRoleIdAclGet**](AuthorizationApi.md#ApiAuthorizationAdminRolesRoleIdAclGet) | **Get** /api/authorization/admin/roles/{roleId}/acl | 
[**ApiAuthorizationAdminRolesRoleIdGet**](AuthorizationApi.md#ApiAuthorizationAdminRolesRoleIdGet) | **Get** /api/authorization/admin/roles/{roleId} | 
[**ApiAuthorizationAdminRolesRoleIdPut**](AuthorizationApi.md#ApiAuthorizationAdminRolesRoleIdPut) | **Put** /api/authorization/admin/roles/{roleId} | 
[**ApiAuthorizationAdminUpdateClaimsClaimIdPut**](AuthorizationApi.md#ApiAuthorizationAdminUpdateClaimsClaimIdPut) | **Put** /api/authorization/admin/update/claims/{claimId} | 
[**ApiAuthorizationUserAclGet**](AuthorizationApi.md#ApiAuthorizationUserAclGet) | **Get** /api/authorization/user/acl | 


# **ApiAuthorizationAdminAclGet**
> GetAclResponse ApiAuthorizationAdminAclGet(ctx, )


This API is responsible for fetching all ACLs from database. It can be accessed only by user who has admin role or got permission to access based on ACL entry in database.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GetAclResponse**](GetACLResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiAuthorizationAdminAclResourceIdRolesGet**
> GetRolesbyResourceIdResponse ApiAuthorizationAdminAclResourceIdRolesGet(ctx, resourceId, optional)


This API is  responsible for retrieving all the roles based on provided resource ID 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **resourceId** | **string**| resourceId of Resource | 
 **optional** | ***AuthorizationApiApiAuthorizationAdminAclResourceIdRolesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthorizationApiApiAuthorizationAdminAclResourceIdRolesGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientRequestSource** | **optional.String**| The client request header that identifies which client had made a request to CPQ Server. | 

### Return type

[**GetRolesbyResourceIdResponse**](GetRolesbyResourceIdResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiAuthorizationAdminClaimsClaimIdGet**
> GetClaimsResponseItem ApiAuthorizationAdminClaimsClaimIdGet(ctx, claimId, optional)


This API is responsible for fetching all claim details based on claimId from database

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **claimId** | **string**| The unique identifier of the claims collection to be retrieved | 
 **optional** | ***AuthorizationApiApiAuthorizationAdminClaimsClaimIdGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthorizationApiApiAuthorizationAdminClaimsClaimIdGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientRequestSource** | **optional.String**| The client request header that identifies which client had made a request to CPQ Server. | 

### Return type

[**GetClaimsResponseItem**](GetClaimsResponseItem.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiAuthorizationAdminClaimsGet**
> AddClaimsResponse ApiAuthorizationAdminClaimsGet(ctx, )


This API is  responsible for fetching all claims from database

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**AddClaimsResponse**](AddClaimsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiAuthorizationAdminClaimsPost**
> AddClaimsResponse ApiAuthorizationAdminClaimsPost(ctx, roles, optional)


This API allows us to post claims  that is to be saved to the DB, the claim is validated against the schema, we can insert bulk entires i.e array of claims

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **roles** | [**AddClaimsRequest**](AddClaimsRequest.md)| Array of Roles | 
 **optional** | ***AuthorizationApiApiAuthorizationAdminClaimsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthorizationApiApiAuthorizationAdminClaimsPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientRequestSource** | **optional.String**| The client request header that identifies which client had made a request to CPQ Server. | 

### Return type

[**AddClaimsResponse**](AddClaimsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiAuthorizationAdminResourcesBulkAssignRolesPost**
> EmptyResponse ApiAuthorizationAdminResourcesBulkAssignRolesPost(ctx, rolesResources)


This API allows us to bulk assign roles to resources

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **rolesResources** | [**BulkAssignRolesToResourcesBody**](BulkAssignRolesToResourcesBody.md)| Array of roles and resources | 

### Return type

[**EmptyResponse**](EmptyResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiAuthorizationAdminResourcesBulkUnassignRolesPost**
> EmptyResponse ApiAuthorizationAdminResourcesBulkUnassignRolesPost(ctx, rolesResources)


This API endpoint allows us to bulk unassign roles from resources

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **rolesResources** | [**BulkAssignRolesToResourcesBody**](BulkAssignRolesToResourcesBody.md)| Array of roles and resources | 

### Return type

[**EmptyResponse**](EmptyResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiAuthorizationAdminResourcesResourceIdRolesPut**
> AssignRolesSuccessResponse ApiAuthorizationAdminResourcesResourceIdRolesPut(ctx, resourceId, roleData, optional)


Assign roles to resource updates Acl object in the database Acl collection, for supplied resourceId. resourceId not supplied will no be updated

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **resourceId** | **string**| The unique identifier of the Acl collection to be retrieved | 
  **roleData** | [**AssignRoles**](AssignRoles.md)| Array of roleId to assign | 
 **optional** | ***AuthorizationApiApiAuthorizationAdminResourcesResourceIdRolesPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthorizationApiApiAuthorizationAdminResourcesResourceIdRolesPutOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientRequestSource** | **optional.String**| The client request header that identifies which client had made a request to CPQ Server. | 

### Return type

[**AssignRolesSuccessResponse**](AssignRolesSuccessResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiAuthorizationAdminResourcesResourceIdRolesRemovePut**
> RemoveRolesSuccessResponse ApiAuthorizationAdminResourcesResourceIdRolesRemovePut(ctx, resourceId, roleData, optional)


Removes roles from resource updates Acl object in the database Acl collection, for supplied resourceId. resourceId not supplied will no be updated

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **resourceId** | **string**| The unique identifier of the Acl collection to be retrieved | 
  **roleData** | [**RemoveRoles**](RemoveRoles.md)| Array of roleId to remove | 
 **optional** | ***AuthorizationApiApiAuthorizationAdminResourcesResourceIdRolesRemovePutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthorizationApiApiAuthorizationAdminResourcesResourceIdRolesRemovePutOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientRequestSource** | **optional.String**| The client request header that identifies which client had made a request to CPQ Server. | 

### Return type

[**RemoveRolesSuccessResponse**](RemoveRolesSuccessResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiAuthorizationAdminRolesGet**
> GetRolesResponse ApiAuthorizationAdminRolesGet(ctx, optional)


This API is  responsible for fetching all roles from database

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AuthorizationApiApiAuthorizationAdminRolesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthorizationApiApiAuthorizationAdminRolesGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientRequestSource** | **optional.String**| The client request header that identifies which client had made a request to CPQ Server. | 
 **isActive** | **optional.String**| isActive &#x3D; true for fetching active roles  | 

### Return type

[**GetRolesResponse**](GetRolesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiAuthorizationAdminRolesPost**
> AddRolesResponse ApiAuthorizationAdminRolesPost(ctx, roles, optional)


This API allows us to post roles  that is to be saved to the DB, the role is validated against the schema, we can insert bulk entires i.e array of roles

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **roles** | [**Roles**](Roles.md)| Array of Roles | 
 **optional** | ***AuthorizationApiApiAuthorizationAdminRolesPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthorizationApiApiAuthorizationAdminRolesPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientRequestSource** | **optional.String**| The client request header that identifies which client had made a request to CPQ Server. | 

### Return type

[**AddRolesResponse**](AddRolesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiAuthorizationAdminRolesRoleIdAclGet**
> GetResourceByRoleIdResponse ApiAuthorizationAdminRolesRoleIdAclGet(ctx, roleId, optional)


This API is  responsible for retrieving all the Resource based on provided role ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **roleId** | **string**| roleId of Role | 
 **optional** | ***AuthorizationApiApiAuthorizationAdminRolesRoleIdAclGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthorizationApiApiAuthorizationAdminRolesRoleIdAclGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientRequestSource** | **optional.String**| The client request header that identifies which client had made a request to CPQ Server. | 
 **isClientFunction** | **optional.String**| isClientFunction - true when requested for client asset | 

### Return type

[**GetResourceByRoleIdResponse**](GetResourceByRoleIdResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiAuthorizationAdminRolesRoleIdGet**
> GetRolesResponseItem ApiAuthorizationAdminRolesRoleIdGet(ctx, roleId, optional)


This API is  responsible for fetching all role details based on roleId from database

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **roleId** | **string**| The unique identifier of the claims collection to be retrieved | 
 **optional** | ***AuthorizationApiApiAuthorizationAdminRolesRoleIdGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthorizationApiApiAuthorizationAdminRolesRoleIdGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientRequestSource** | **optional.String**| The client request header that identifies which client had made a request to CPQ Server. | 

### Return type

[**GetRolesResponseItem**](GetRolesResponseItem.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiAuthorizationAdminRolesRoleIdPut**
> UpdateRoleResponse ApiAuthorizationAdminRolesRoleIdPut(ctx, roleId, roleData, optional)


This API allows us to update the single role entry that exists in the DB, only valid claims are allowed to be updated

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **roleId** | **string**| The unique identifier of the claims collection to be retrieved | 
  **roleData** | [**Role**](Role.md)| The role object to update | 
 **optional** | ***AuthorizationApiApiAuthorizationAdminRolesRoleIdPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthorizationApiApiAuthorizationAdminRolesRoleIdPutOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientRequestSource** | **optional.String**| The client request header that identifies which client had made a request to CPQ Server. | 

### Return type

[**UpdateRoleResponse**](UpdateRoleResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiAuthorizationAdminUpdateClaimsClaimIdPut**
> UpdateClaimResponse ApiAuthorizationAdminUpdateClaimsClaimIdPut(ctx, claimId, claimData, optional)


Updates claim object in the database Claims collection, for each supplied claimId. claimId not supplied will no be updated, no claims will be added or deleted. An updated object is returned on success. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **claimId** | **string**| The unique identifier of the claims collection to be retrieved | 
  **claimData** | [**Claims**](Claims.md)| The claim object to update | 
 **optional** | ***AuthorizationApiApiAuthorizationAdminUpdateClaimsClaimIdPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthorizationApiApiAuthorizationAdminUpdateClaimsClaimIdPutOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientRequestSource** | **optional.String**| The client request header that identifies which client had made a request to CPQ Server. | 

### Return type

[**UpdateClaimResponse**](UpdateClaimResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiAuthorizationUserAclGet**
> GetUserAclResponse ApiAuthorizationUserAclGet(ctx, optional)


This API is  responsible for fetching all ACLs from database based on user claims

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AuthorizationApiApiAuthorizationUserAclGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthorizationApiApiAuthorizationUserAclGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientRequestSource** | **optional.String**| The client request header that identifies which client had made a request to CPQ Server. | 
 **isClientFunction** | **optional.String**| isClientFunction - true when requested for client asset | 

### Return type

[**GetUserAclResponse**](GetUserACLResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

