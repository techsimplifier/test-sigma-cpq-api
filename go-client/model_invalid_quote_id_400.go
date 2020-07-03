/*
 * CPQ
 *
 * CPQ Web API Documentation
 *
 * API version: 2.9.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Bad request - invalid quote id or invalid request body
type InvalidQuoteId400 struct {
	ResponseCode string `json:"responseCode"`
	HttpStatus float32 `json:"httpStatus"`
	ResponseText string `json:"responseText"`
	ExceptionType string `json:"exceptionType,omitempty"`
}
