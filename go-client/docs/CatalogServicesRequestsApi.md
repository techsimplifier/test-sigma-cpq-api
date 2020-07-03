# \CatalogServicesRequestsApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiClassificationsClassificationRootIdGet**](CatalogServicesRequestsApi.md#ApiClassificationsClassificationRootIdGet) | **Get** /api/classifications/{classificationRootId} | 
[**ApiEntitiesEntityIdGet**](CatalogServicesRequestsApi.md#ApiEntitiesEntityIdGet) | **Get** /api/entities/{entityId} | 
[**ApiEntitiesMetaDataEntityIdGet**](CatalogServicesRequestsApi.md#ApiEntitiesMetaDataEntityIdGet) | **Get** /api/entities/metaData/{entityId} | 
[**ApiOffersByTypeGet**](CatalogServicesRequestsApi.md#ApiOffersByTypeGet) | **Get** /api/offers/byType | 
[**ApiOffersGet**](CatalogServicesRequestsApi.md#ApiOffersGet) | **Get** /api/offers | 
[**ApiOffersUntransformedGet**](CatalogServicesRequestsApi.md#ApiOffersUntransformedGet) | **Get** /api/offers/untransformed | 


# **ApiClassificationsClassificationRootIdGet**
> ClassificationsObject ApiClassificationsClassificationRootIdGet(ctx, classificationRootId, optional)


This API is responsible for fetching a list of classified products eligible for selection by a customer from catalog services, for a specific classificationRootId. The returned list of products can be used for selection and addition to the quote based on the customers choice.<br/> If a failed response is returned by the API, the responseCode value contains:  responseCode | Meaning   --- | ----   CSConnectionRequestError | The request was sent to Catalog Services and failed in this component.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **classificationRootId** | **string**| Classification root ID&lt;br/&gt;Example:TSales_Channel | 
 **optional** | ***CatalogServicesRequestsApiApiClassificationsClassificationRootIdGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CatalogServicesRequestsApiApiClassificationsClassificationRootIdGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientRequestSource** | **optional.String**| The client request header that identifies which client had made a request to CPQ Server. | 

### Return type

[**ClassificationsObject**](ClassificationsObject.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiEntitiesEntityIdGet**
> ProductSpecObject ApiEntitiesEntityIdGet(ctx, entityId, optional)


Gets the catalog defined offer specification for the entity with the specified Id. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **entityId** | **string**| entity ID | 
 **optional** | ***CatalogServicesRequestsApiApiEntitiesEntityIdGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CatalogServicesRequestsApiApiEntitiesEntityIdGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientRequestSource** | **optional.String**| The client request header that identifies which client had made a request to CPQ Server. | 

### Return type

[**ProductSpecObject**](ProductSpecObject.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiEntitiesMetaDataEntityIdGet**
> ProductSpecObject ApiEntitiesMetaDataEntityIdGet(ctx, entityId, optional)


Gets the metadata for the entity with the specified Id. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **entityId** | **string**| entity ID | 
 **optional** | ***CatalogServicesRequestsApiApiEntitiesMetaDataEntityIdGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CatalogServicesRequestsApiApiEntitiesMetaDataEntityIdGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientRequestSource** | **optional.String**| The client request header that identifies which client had made a request to CPQ Server. | 

### Return type

[**ProductSpecObject**](ProductSpecObject.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiOffersByTypeGet**
> GetOffersByTypeResponse ApiOffersByTypeGet(ctx, instanceTypeNames, optional)


Retrieve a collection of entities which match the supplied instance types that are available on the date specified. If no date is specified it defaults to today's date. An example of the call would look like:<br/><br/><b>/api/offers/byType?InstanceTypeNames=Package</b><br/><br/>Please note that this is an example only, and should not be used in production.<br/><br/>Please also note that only the top level properties defined on each entity are returned - no properties which are child entities or charges are returned.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **instanceTypeNames** | **string**| Comma separated list of entity types you would like returned. These can include&lt;br/&gt; Package,Bundle Promotion etc. Note that these types must be specified with the first letter capitalised | 
 **optional** | ***CatalogServicesRequestsApiApiOffersByTypeGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CatalogServicesRequestsApiApiOffersByTypeGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientRequestSource** | **optional.String**| The client request header that identifies which client had made a request to CPQ Server. | 
 **targetDate** | **optional.String**| The date in YYYY-MM-DD format which represents the date that the returned entities are marked as available (ie between AvailableStartDate and AvailableEndDate. If this parameter is not specified then it defaults to today&#39;s date. | 

### Return type

[**GetOffersByTypeResponse**](GetOffersByTypeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiOffersGet**
> GetOffersByClassificationResponse ApiOffersGet(ctx, instanceTypeNames, classificationElementName, classifications, xsltCode, optional)


Retrieve the offers which match the supplied classifications, facts, characteristics or categories. A fully populated endpoint would look like:<br/><br/><b>/api/offers?InstanceTypeNames=Package&Classifications=[TSales_Channel_Classification;e4a2a606-501f-f009-731a-64ba27c59d50;false]&ClassificationElementName=TSales_Channel_Classification&FactElementFilter=[Lookup_Colour;fed70bc9-c12d-4875-b2d3-e1930992440e;false]&Categories=[22,18;true]&Characteristics=[Lookup_Colour;846fa97b-b9a5-4671-9660-f3b0d554da4b;false]&xsltCode=offer_short&at[p1]=ID&el[p2]=name</b><br/><br/>Please note that this is an example only, and should not be used in production.<br/>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **instanceTypeNames** | **string**| Comma separated list of entity types you would like returned. These can include&lt;br/&gt; Package,Bundle Promotion etc. Note that these types must be specified with the first letter capitalised | 
  **classificationElementName** | **string**| That name of the element on each entity you wish returned that is of the Classification type specified in the data packets | 
  **classifications** | **string**| A list of classification packets that define which classification on which to match. Each packet has the following format:&lt;br/&gt;[classificationElementName; classificationGuids; MatchAllClassification]&lt;br/&gt;CLASSIFICATIONELEMENTNAME is the name of the schema element that holds the Classifications on which they query is filtering.&lt;br/&gt;CLASSIFICATIONGUIDS is a comma-separated list that represents all the Classifications on which the query is filtering.&lt;br/&gt;MATCHALLCLASSIFICATION is either TRUE or FALSE to specify whether the query results need to match ALL of these or just ANY. You can pass multiple packets into this query. | 
  **xsltCode** | **string**| The code of the XSLT stored in the Catalog Services database that transforms the XML output into a desired format. This is required and must return data in the required format. | 
 **optional** | ***CatalogServicesRequestsApiApiOffersGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CatalogServicesRequestsApiApiOffersGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **clientRequestSource** | **optional.String**| The client request header that identifies which client had made a request to CPQ Server. | 
 **characteristics** | **optional.String**| A list of characteristic packets that define the characteristic on which you want to match, with each packet in the following format:&lt;br/&gt;[characteristicElementName; characteristicGuids; MatchAllCharacteristics]&lt;br/&gt;CHARACTERISTICELEMENTNAME is the name of the schema element that holds the characteristics on which they query is filtering.&lt;br/&gt;         CHARACTERISTICGUIDS is a comma-separated list that represents all the characteristics on which the query is filtering.&lt;br/&gt;MATCHALLCHARACTERISTICS is either TRUE or FALSE to specify whether the query results need to match ALL of these or just ANY. You can pass multiple packets into this query.&lt;br/&gt;Add multiple characteristic and classification packets by appending each element set inside square brackets ([ ]). Combine each additional packet with the others using AND logic to specify that all must match. | 
 **factElementFilter** | **optional.String**| A list of fact values for those on which you want to query to match. Each packet takes the following format:&lt;br/&gt;[FactValue; FactValueGuids; MatchAllFactValues]&lt;br /&gt;FACTVALUE is the name of the schema element that holds the fact value on which you are filtering.&lt;br/&gt; FACTVALUEGUIDS is a comma-separated list of GUIDs representing the fact values on which you are filtering. &lt;br/&gt;MATCHALLFACTVALUES is a true or false value indicating whether the query matches on ALL fact values or just ANY fact values.Note: You can pass multiple fact value packets in this query. | 
 **categories** | **optional.String**| A data packet which takes the following format: &lt;br/&gt;[Categories;IncludeChildCategories].&lt;br/&gt; CATEGORIES is a comma separated list of category Ids.&lt;br/&gt; INCLUDECHILDCATEGORIES indicates whether entities should be returned that live in categories that are children of the category Id specified in the CATEGORIES parameter | 
 **atP1** | **optional.String**| A parameter in the parameter list for the above XSLT which are used to retrieve values which appear as XML attributes on a matching entity, there can be as many of these as you like as long as each key is unique, only one &#39;p1&#39; property. You cannot have at[p1] and el[p1] every subsequent parameter should be incremented | 
 **elP1** | **optional.String**| A parameter in the parameter list for the above XSLT which are used to retrieve values which appear as XML attributes on a matching entity, there can be as many of these as you like as long as each key is unique, only one &#39;p1&#39; property. You cannot have at[p1] and el[p1] every subsequent parameter should be incremented | 
 **atP2** | **optional.String**| A parameter in the parameter list for the above XSLT which are used to retrieve values which appear as XML attributes on a matching entity, there can be as many of these as you like as long as each key is unique, only one &#39;p1&#39; property. You cannot have at[p1] and el[p1] every subsequent parameter should be incremented | 
 **elP2** | **optional.String**| A parameter in the parameter list for the above XSLT which are used to retrieve values which appear as XML attributes on a matching entity, there can be as many of these as you like as long as each key is unique, only one &#39;p1&#39; property. You cannot have at[p1] and el[p1] every subsequent parameter should be incremented | 
 **atP3** | **optional.String**| A parameter in the parameter list for the above XSLT which are used to retrieve values which appear as XML attributes on a matching entity, there can be as many of these as you like as long as each key is unique, only one &#39;p1&#39; property. You cannot have at[p1] and el[p1] every subsequent parameter should be incremented | 
 **elP3** | **optional.String**| A parameter in the parameter list for the above XSLT which are used to retrieve values which appear as XML attributes on a matching entity, there can be as many of these as you like as long as each key is unique, only one &#39;p1&#39; property. You cannot have at[p1] and el[p1] every subsequent parameter should be incremented | 

### Return type

[**GetOffersByClassificationResponse**](GetOffersByClassificationResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiOffersUntransformedGet**
> GetOffersUntransformedResponse ApiOffersUntransformedGet(ctx, instanceTypeNames, classificationElementName, classifications, optional)


Retrieve the offers which match the supplied classifications, facts, characteristics or categories. The results are not subject to XML transform so specifying any XsltCode or XSLT parameters is not necessary and are ignored if specified.<br/><br/> A fully populated endpoint would look like:<br/><br/><b>/api/offers/untransformed?InstanceTypeNames=Package&Classifications=[TSales_Channel_Classification;e4a2a606-501f-f009-731a-64ba27c59d50;false]&ClassificationElementName=TSales_Channel_Classification&FactElementFilter=[Lookup_Colour;fed70bc9-c12d-4875-b2d3-e1930992440e;false]&Categories=[22,18;true]&Characteristics=[Lookup_Colour;846fa97b-b9a5-4671-9660-f3b0d554da4b;false]</b><br/><br/>Please note that this is an example only, and should not be used in production.<br/><br/>Please also note that calling this endpoint will result in potentially a very large amount of data being returned depending upon the structure of the models in Catalog and the Instance Types requested.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **instanceTypeNames** | **string**| Comma separated list of entity types you would like returned. These can include&lt;br/&gt; Package,Bundle Promotion etc. Note that these types must be specified with the first letter capitalised | 
  **classificationElementName** | **string**| That name of the element on each entity you wish returned that is of the Classification type specified in the data packets | 
  **classifications** | **string**| A list of classification packets that define which classification on which to match. Each packet has the following format:&lt;br/&gt;[classificationElementName; classificationGuids; MatchAllClassification]&lt;br/&gt;CLASSIFICATIONELEMENTNAME is the name of the schema element that holds the Classifications on which they query is filtering.&lt;br/&gt;CLASSIFICATIONGUIDS is a comma-separated list that represents all the Classifications on which the query is filtering.&lt;br/&gt;MATCHALLCLASSIFICATION is either TRUE or FALSE to specify whether the query results need to match ALL of these or just ANY. You can pass multiple packets into this query. | 
 **optional** | ***CatalogServicesRequestsApiApiOffersUntransformedGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CatalogServicesRequestsApiApiOffersUntransformedGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **clientRequestSource** | **optional.String**| The client request header that identifies which client had made a request to CPQ Server. | 
 **characteristics** | **optional.String**| A list of characteristic packets that define the characteristic on which you want to match, with each packet in the following format:&lt;br/&gt;[characteristicElementName; characteristicGuids; MatchAllCharacteristics]&lt;br/&gt;CHARACTERISTICELEMENTNAME is the name of the schema element that holds the characteristics on which they query is filtering.&lt;br/&gt;         CHARACTERISTICGUIDS is a comma-separated list that represents all the characteristics on which the query is filtering.&lt;br/&gt;MATCHALLCHARACTERISTICS is either TRUE or FALSE to specify whether the query results need to match ALL of these or just ANY. You can pass multiple packets into this query.&lt;br/&gt;Add multiple characteristic and classification packets by appending each element set inside square brackets ([ ]). Combine each additional packet with the others using AND logic to specify that all must match. | 
 **factElementFilter** | **optional.String**| A list of fact values for those on which you want to query to match. Each packet takes the following format:&lt;br/&gt;[FactValue; FactValueGuids; MatchAllFactValues]&lt;br /&gt;FACTVALUE is the name of the schema element that holds the fact value on which you are filtering.&lt;br/&gt; FACTVALUEGUIDS is a comma-separated list of GUIDs representing the fact values on which you are filtering. &lt;br/&gt;MATCHALLFACTVALUES is a true or false value indicating whether the query matches on ALL fact values or just ANY fact values.Note: You can pass multiple fact value packets in this query. | 
 **categories** | **optional.String**| A data packet which takes the following format: &lt;br/&gt;[Categories;IncludeChildCategories].&lt;br/&gt; CATEGORIES is a comma separated list of category Ids.&lt;br/&gt; INCLUDECHILDCATEGORIES indicates whether entities should be returned that live in categories that are children of the category Id specified in the CATEGORIES parameter | 

### Return type

[**GetOffersUntransformedResponse**](GetOffersUntransformedResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

