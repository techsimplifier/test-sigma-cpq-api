/*
 * CPQ
 *
 * CPQ Web API Documentation
 *
 * API version: 2.9.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type SuccessResponse200ResponseBody struct {
	SuccessfulIdentifiers []string `json:"successfulIdentifiers,omitempty"`
	FailedIdentifiers []ErrorUnknown `json:"failedIdentifiers,omitempty"`
}
