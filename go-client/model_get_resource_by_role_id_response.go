/*
 * CPQ
 *
 * CPQ Web API Documentation
 *
 * API version: 2.9.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type GetResourceByRoleIdResponse struct {
	Roles string `json:"roles,omitempty"`
	Resources []string `json:"resources,omitempty"`
}