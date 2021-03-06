/*
 * CPQ
 *
 * CPQ Web API Documentation
 *
 * API version: 2.9.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type Role struct {
	RoleName string `json:"roleName,omitempty"`
	Description string `json:"description,omitempty"`
	Claims []RoleClaims `json:"claims,omitempty"`
}
