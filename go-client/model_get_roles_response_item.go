/*
 * CPQ
 *
 * CPQ Web API Documentation
 *
 * API version: 2.9.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type GetRolesResponseItem struct {
	RoleId string `json:"roleId,omitempty"`
	RoleName string `json:"roleName,omitempty"`
	Description string `json:"description,omitempty"`
	Active bool `json:"active,omitempty"`
	IsAdministrator bool `json:"isAdministrator,omitempty"`
	Claims []ClaimsData `json:"claims,omitempty"`
	// The date created
	Created string `json:"created,omitempty"`
	CreatedBy string `json:"createdBy,omitempty"`
	// The date last updated
	Modified string `json:"modified,omitempty"`
	ModifiedBy string `json:"modifiedBy,omitempty"`
}
