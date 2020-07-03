
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

type ExternalCallsApiService service

/* 
ExternalCallsApiService
Gets the appropriate data from external system by executing registered plugin for the same.&lt;br/&gt;Where &#39;apiName&#39; would be the name you registered for the externalLookup plugin.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param quoteId entity ID
 * @param optional nil or *ExternalCallsApiApiQuotesQuoteIdExternalLookupapiNameGetOpts - Optional Parameters:
     * @param "ClientRequestSource" (optional.String) -  The client request header that identifies which client had made a request to CPQ Server.

@return ExternalLookupObject
*/

type ExternalCallsApiApiQuotesQuoteIdExternalLookupapiNameGetOpts struct { 
	ClientRequestSource optional.String
}

func (a *ExternalCallsApiService) ApiQuotesQuoteIdExternalLookupapiNameGet(ctx context.Context, quoteId string, localVarOptionals *ExternalCallsApiApiQuotesQuoteIdExternalLookupapiNameGetOpts) (ExternalLookupObject, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue ExternalLookupObject
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/quotes/{quoteId}/externalLookup/'apiName'"
	localVarPath = strings.Replace(localVarPath, "{"+"quoteId"+"}", fmt.Sprintf("%v", quoteId), -1)

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
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		
		if localVarHttpResponse.StatusCode == 200 {
			var v ExternalLookupObject
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/* 
ExternalCallsApiService
Reserves the appropriate data in an external system by executing registered plugin for the same.&lt;br/&gt;Where &#39;apiName&#39; would be the name you registered for the externalReservation plugin.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param quoteId entity ID
 * @param reservationValuesArray Array of values to be reserved
 * @param optional nil or *ExternalCallsApiApiQuotesQuoteIdExternalReservationapiNamePostOpts - Optional Parameters:
     * @param "ClientRequestSource" (optional.String) -  The client request header that identifies which client had made a request to CPQ Server.

@return ExternalReservationObject
*/

type ExternalCallsApiApiQuotesQuoteIdExternalReservationapiNamePostOpts struct { 
	ClientRequestSource optional.String
}

func (a *ExternalCallsApiService) ApiQuotesQuoteIdExternalReservationapiNamePost(ctx context.Context, quoteId string, reservationValuesArray ExternalVerificationValueType, localVarOptionals *ExternalCallsApiApiQuotesQuoteIdExternalReservationapiNamePostOpts) (ExternalReservationObject, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue ExternalReservationObject
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/quotes/{quoteId}/externalReservation/'apiName'"
	localVarPath = strings.Replace(localVarPath, "{"+"quoteId"+"}", fmt.Sprintf("%v", quoteId), -1)

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
	// body params
	localVarPostBody = &reservationValuesArray
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		
		if localVarHttpResponse.StatusCode == 200 {
			var v ExternalReservationObject
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/* 
ExternalCallsApiService
Posts the verification value array to external system with quote id and returns the verification result.&lt;br&gt;Where &#39;apiName&#39; would be the name you registered for the externalVerification plugin.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param quoteId entity ID
 * @param verificationValuesArray Array of values to be verified
 * @param optional nil or *ExternalCallsApiApiQuotesQuoteIdExternalVerificationapiNamePostOpts - Optional Parameters:
     * @param "ClientRequestSource" (optional.String) -  The client request header that identifies which client had made a request to CPQ Server.

@return ExternalVerificationObject
*/

type ExternalCallsApiApiQuotesQuoteIdExternalVerificationapiNamePostOpts struct { 
	ClientRequestSource optional.String
}

func (a *ExternalCallsApiService) ApiQuotesQuoteIdExternalVerificationapiNamePost(ctx context.Context, quoteId string, verificationValuesArray ExternalVerificationValueType, localVarOptionals *ExternalCallsApiApiQuotesQuoteIdExternalVerificationapiNamePostOpts) (ExternalVerificationObject, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue ExternalVerificationObject
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/quotes/{quoteId}/externalVerification/'apiName'"
	localVarPath = strings.Replace(localVarPath, "{"+"quoteId"+"}", fmt.Sprintf("%v", quoteId), -1)

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
	// body params
	localVarPostBody = &verificationValuesArray
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		
		if localVarHttpResponse.StatusCode == 200 {
			var v ExternalVerificationObject
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
