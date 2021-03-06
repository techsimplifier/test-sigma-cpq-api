
/*
 * CPQ
 *
 * CPQ Web API Documentation
 *
 * API version: 2.9.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"fmt"
	"github.com/antihax/optional"
)

// Linger please
var (
	_ context.Context
)

type NegotiationCustomerMappingsApiService service

/* 
NegotiationCustomerMappingsApiService
Get all customer mappings by customerId
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param customerId The identifier of customer mapping
 * @param contentType Specifies the format of the data that is being sent to the API. Only JSON data is accepted
 * @param optional nil or *NegotiationCustomerMappingsApiApiCustomerMappingsCustomerCustomerIdGetOpts - Optional Parameters:
     * @param "ClientRequestSource" (optional.String) -  The client request header that identifies which client had made a request to CPQ Server.
     * @param "CpqLanguage" (optional.String) -  The language code (in ISO 639-1 format) that error messages are to be translated to for returning back to the client. If not supplied it defaults to en-GB


*/

type NegotiationCustomerMappingsApiApiCustomerMappingsCustomerCustomerIdGetOpts struct { 
	ClientRequestSource optional.String
	CpqLanguage optional.String
}

func (a *NegotiationCustomerMappingsApiService) ApiCustomerMappingsCustomerCustomerIdGet(ctx context.Context, customerId string, contentType string, localVarOptionals *NegotiationCustomerMappingsApiApiCustomerMappingsCustomerCustomerIdGetOpts) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/customerMappings/customer/{customerId}"
	localVarPath = strings.Replace(localVarPath, "{"+"customerId"+"}", fmt.Sprintf("%v", customerId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.ClientRequestSource.IsSet() {
		localVarHeaderParams["clientRequestSource"] = parameterToString(localVarOptionals.ClientRequestSource.Value(), "")
	}
	if localVarOptionals != nil && localVarOptionals.CpqLanguage.IsSet() {
		localVarHeaderParams["cpq-language"] = parameterToString(localVarOptionals.CpqLanguage.Value(), "")
	}
	localVarHeaderParams["Content-Type"] = parameterToString(contentType, "")
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}


	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/* 
NegotiationCustomerMappingsApiService
Get customer mappings by customerId and date range
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param customerId The identifier of customer mapping
 * @param startDate Start date range of query in ISO 8601 format where customer mappings in effective
 * @param endDate End date range of query in ISO 8601 format where customer mappings in effective
 * @param contentType Specifies the format of the data that is being sent to the API. Only JSON data is accepted
 * @param optional nil or *NegotiationCustomerMappingsApiApiCustomerMappingsCustomerCustomerIdStartDateStartDateEndDateEndDateGetOpts - Optional Parameters:
     * @param "ClientRequestSource" (optional.String) -  The client request header that identifies which client had made a request to CPQ Server.
     * @param "CpqLanguage" (optional.String) -  The language code (in ISO 639-1 format) that error messages are to be translated to for returning back to the client. If not supplied it defaults to en-GB


*/

type NegotiationCustomerMappingsApiApiCustomerMappingsCustomerCustomerIdStartDateStartDateEndDateEndDateGetOpts struct { 
	ClientRequestSource optional.String
	CpqLanguage optional.String
}

func (a *NegotiationCustomerMappingsApiService) ApiCustomerMappingsCustomerCustomerIdStartDateStartDateEndDateEndDateGet(ctx context.Context, customerId string, startDate string, endDate string, contentType string, localVarOptionals *NegotiationCustomerMappingsApiApiCustomerMappingsCustomerCustomerIdStartDateStartDateEndDateEndDateGetOpts) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/customerMappings/customer/{customerId}/startDate/{startDate}/endDate/{endDate}"
	localVarPath = strings.Replace(localVarPath, "{"+"customerId"+"}", fmt.Sprintf("%v", customerId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"startDate"+"}", fmt.Sprintf("%v", startDate), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"endDate"+"}", fmt.Sprintf("%v", endDate), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.ClientRequestSource.IsSet() {
		localVarHeaderParams["clientRequestSource"] = parameterToString(localVarOptionals.ClientRequestSource.Value(), "")
	}
	if localVarOptionals != nil && localVarOptionals.CpqLanguage.IsSet() {
		localVarHeaderParams["cpq-language"] = parameterToString(localVarOptionals.CpqLanguage.Value(), "")
	}
	localVarHeaderParams["Content-Type"] = parameterToString(contentType, "")
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}


	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/* 
NegotiationCustomerMappingsApiService
Get contract rule set identifier of customer mappings by customer Id and date
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param customerId The identifier of customer mapping
 * @param startDate Start date in ISO 8601 format where customer mapping is in effective
 * @param contentType Specifies the format of the data that is being sent to the API. Only JSON data is accepted
 * @param optional nil or *NegotiationCustomerMappingsApiApiCustomerMappingsCustomerIdContractIdGetOpts - Optional Parameters:
     * @param "ClientRequestSource" (optional.String) -  The client request header that identifies which client had made a request to CPQ Server.
     * @param "CpqLanguage" (optional.String) -  The language code (in ISO 639-1 format) that error messages are to be translated to for returning back to the client. If not supplied it defaults to en-GB
     * @param "IsContractId" (optional.Bool) -  If set to true, return contract Id as contract rule set identifier, if set to false, return fc_Id as contract rule set identifier.  If value is not specified, default to true


*/

type NegotiationCustomerMappingsApiApiCustomerMappingsCustomerIdContractIdGetOpts struct { 
	ClientRequestSource optional.String
	CpqLanguage optional.String
	IsContractId optional.Bool
}

func (a *NegotiationCustomerMappingsApiService) ApiCustomerMappingsCustomerIdContractIdGet(ctx context.Context, customerId string, startDate string, contentType string, localVarOptionals *NegotiationCustomerMappingsApiApiCustomerMappingsCustomerIdContractIdGetOpts) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/customerMappings/{customerId}/contractId"
	localVarPath = strings.Replace(localVarPath, "{"+"customerId"+"}", fmt.Sprintf("%v", customerId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("startDate", parameterToString(startDate, ""))
	if localVarOptionals != nil && localVarOptionals.IsContractId.IsSet() {
		localVarQueryParams.Add("isContractId", parameterToString(localVarOptionals.IsContractId.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.ClientRequestSource.IsSet() {
		localVarHeaderParams["clientRequestSource"] = parameterToString(localVarOptionals.ClientRequestSource.Value(), "")
	}
	if localVarOptionals != nil && localVarOptionals.CpqLanguage.IsSet() {
		localVarHeaderParams["cpq-language"] = parameterToString(localVarOptionals.CpqLanguage.Value(), "")
	}
	localVarHeaderParams["Content-Type"] = parameterToString(contentType, "")
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}


	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/* 
NegotiationCustomerMappingsApiService
Update customer mappings
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param fcUser The name of the user that is updating the custom mapping
 * @param customerMapping Customer mapping that is to be written to the database.
 * @param customerId The identifier of customer mapping to be updated
 * @param identifier The identifier of a contract rule set that assoicated with the customer mapping to be updated - can be provided either as a guid value or as a value which matches the pattern ^[a-z0-9]{6}$
 * @param contentType Specifies the format of the data that is being sent to the API. Only JSON data is accepted
 * @param optional nil or *NegotiationCustomerMappingsApiApiCustomerMappingsCustomerIdIdentifierIdentifierPutOpts - Optional Parameters:
     * @param "ClientRequestSource" (optional.String) -  The client request header that identifies which client had made a request to CPQ Server.
     * @param "CpqLanguage" (optional.String) -  The language code (in ISO 639-1 format) that error messages are to be translated to for returning back to the client. If not supplied it defaults to en-GB


*/

type NegotiationCustomerMappingsApiApiCustomerMappingsCustomerIdIdentifierIdentifierPutOpts struct { 
	ClientRequestSource optional.String
	CpqLanguage optional.String
}

func (a *NegotiationCustomerMappingsApiService) ApiCustomerMappingsCustomerIdIdentifierIdentifierPut(ctx context.Context, fcUser string, customerMapping CustomerMappingModel, customerId string, identifier string, contentType string, localVarOptionals *NegotiationCustomerMappingsApiApiCustomerMappingsCustomerIdIdentifierIdentifierPutOpts) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/customerMappings/{customerId}/identifier/{identifier}"
	localVarPath = strings.Replace(localVarPath, "{"+"customerId"+"}", fmt.Sprintf("%v", customerId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"identifier"+"}", fmt.Sprintf("%v", identifier), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(fcUser) < 1 {
		return nil, reportError("fcUser must have at least 1 elements")
	}
	if strlen(fcUser) > 30 {
		return nil, reportError("fcUser must have less than 30 elements")
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.ClientRequestSource.IsSet() {
		localVarHeaderParams["clientRequestSource"] = parameterToString(localVarOptionals.ClientRequestSource.Value(), "")
	}
	localVarHeaderParams["fc-user"] = parameterToString(fcUser, "")
	if localVarOptionals != nil && localVarOptionals.CpqLanguage.IsSet() {
		localVarHeaderParams["cpq-language"] = parameterToString(localVarOptionals.CpqLanguage.Value(), "")
	}
	localVarHeaderParams["Content-Type"] = parameterToString(contentType, "")
	// body params
	localVarPostBody = &customerMapping
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}


	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/* 
NegotiationCustomerMappingsApiService
Bulk create customer mappings that maps customer Id to contract rule set identifier
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param customerMappings Array of customer mapping format that is to be written to the database.
 * @param contentType Specifies the format of the data that is being sent to the API. Only JSON data is accepted
 * @param optional nil or *NegotiationCustomerMappingsApiApiCustomerMappingsPostOpts - Optional Parameters:
     * @param "ClientRequestSource" (optional.String) -  The client request header that identifies which client had made a request to CPQ Server.
     * @param "CpqLanguage" (optional.String) -  The language code (in ISO 639-1 format) that error messages are to be translated to for returning back to the client. If not supplied it defaults to en-GB


*/

type NegotiationCustomerMappingsApiApiCustomerMappingsPostOpts struct { 
	ClientRequestSource optional.String
	CpqLanguage optional.String
}

func (a *NegotiationCustomerMappingsApiService) ApiCustomerMappingsPost(ctx context.Context, customerMappings CreateCustomerMappingsRequest, contentType string, localVarOptionals *NegotiationCustomerMappingsApiApiCustomerMappingsPostOpts) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/customerMappings"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.ClientRequestSource.IsSet() {
		localVarHeaderParams["clientRequestSource"] = parameterToString(localVarOptionals.ClientRequestSource.Value(), "")
	}
	if localVarOptionals != nil && localVarOptionals.CpqLanguage.IsSet() {
		localVarHeaderParams["cpq-language"] = parameterToString(localVarOptionals.CpqLanguage.Value(), "")
	}
	localVarHeaderParams["Content-Type"] = parameterToString(contentType, "")
	// body params
	localVarPostBody = &customerMappings
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}


	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/* 
NegotiationCustomerMappingsApiService
Get customer mappings by date range
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param startDate Start date range of query in ISO 8601 format where customer mappings in effective
 * @param endDate End date range of query in ISO 8601 format where customer mappings in effective
 * @param contentType Specifies the format of the data that is being sent to the API. Only JSON data is accepted
 * @param optional nil or *NegotiationCustomerMappingsApiApiCustomerMappingsStartDateStartDateEndDateEndDateGetOpts - Optional Parameters:
     * @param "ClientRequestSource" (optional.String) -  The client request header that identifies which client had made a request to CPQ Server.
     * @param "CpqLanguage" (optional.String) -  The language code (in ISO 639-1 format) that error messages are to be translated to for returning back to the client. If not supplied it defaults to en-GB


*/

type NegotiationCustomerMappingsApiApiCustomerMappingsStartDateStartDateEndDateEndDateGetOpts struct { 
	ClientRequestSource optional.String
	CpqLanguage optional.String
}

func (a *NegotiationCustomerMappingsApiService) ApiCustomerMappingsStartDateStartDateEndDateEndDateGet(ctx context.Context, startDate string, endDate string, contentType string, localVarOptionals *NegotiationCustomerMappingsApiApiCustomerMappingsStartDateStartDateEndDateEndDateGetOpts) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/customerMappings/startDate/{startDate}/endDate/{endDate}"
	localVarPath = strings.Replace(localVarPath, "{"+"startDate"+"}", fmt.Sprintf("%v", startDate), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"endDate"+"}", fmt.Sprintf("%v", endDate), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.ClientRequestSource.IsSet() {
		localVarHeaderParams["clientRequestSource"] = parameterToString(localVarOptionals.ClientRequestSource.Value(), "")
	}
	if localVarOptionals != nil && localVarOptionals.CpqLanguage.IsSet() {
		localVarHeaderParams["cpq-language"] = parameterToString(localVarOptionals.CpqLanguage.Value(), "")
	}
	localVarHeaderParams["Content-Type"] = parameterToString(contentType, "")
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}


	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/* 
NegotiationCustomerMappingsApiService
Get all customer mappings and summary contracts by customerId
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param customerId The identifier of customer mapping
 * @param contentType Specifies the format of the data that is being sent to the API. Only JSON data is accepted
 * @param optional nil or *NegotiationCustomerMappingsApiApiCustomersCustomerIdContractsGetOpts - Optional Parameters:
     * @param "ClientRequestSource" (optional.String) -  The client request header that identifies which client had made a request to CPQ Server.
     * @param "CpqLanguage" (optional.String) -  The language code (in ISO 639-1 format) that error messages are to be translated to for returning back to the client. If not supplied it defaults to en-GB


*/

type NegotiationCustomerMappingsApiApiCustomersCustomerIdContractsGetOpts struct { 
	ClientRequestSource optional.String
	CpqLanguage optional.String
}

func (a *NegotiationCustomerMappingsApiService) ApiCustomersCustomerIdContractsGet(ctx context.Context, customerId string, contentType string, localVarOptionals *NegotiationCustomerMappingsApiApiCustomersCustomerIdContractsGetOpts) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/customers/{customerId}/contracts"
	localVarPath = strings.Replace(localVarPath, "{"+"customerId"+"}", fmt.Sprintf("%v", customerId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.ClientRequestSource.IsSet() {
		localVarHeaderParams["clientRequestSource"] = parameterToString(localVarOptionals.ClientRequestSource.Value(), "")
	}
	if localVarOptionals != nil && localVarOptionals.CpqLanguage.IsSet() {
		localVarHeaderParams["cpq-language"] = parameterToString(localVarOptionals.CpqLanguage.Value(), "")
	}
	localVarHeaderParams["Content-Type"] = parameterToString(contentType, "")
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}


	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}
