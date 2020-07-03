/*
 * CPQ
 *
 * CPQ Web API Documentation
 *
 * API version: 2.9.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type WorkflowValidationError400 struct {
	ResponseCode string `json:"responseCode"`
	HttpStatus float32 `json:"httpStatus"`
	ExceptionType string `json:"exceptionType"`
	ResponseText string `json:"responseText"`
	ResponseBody []WorkflowValidationResponseBody `json:"responseBody"`
}