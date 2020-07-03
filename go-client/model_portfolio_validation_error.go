/*
 * CPQ
 *
 * CPQ Web API Documentation
 *
 * API version: 2.9.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type PortfolioValidationError struct {
	EntityID string `json:"EntityID,omitempty"`
	EntityUniqueCode string `json:"EntityUniqueCode,omitempty"`
	ChildID string `json:"ChildID,omitempty"`
	ChildUniqueCode string `json:"ChildUniqueCode,omitempty"`
	ErrorCode string `json:"ErrorCode,omitempty"`
	Message string `json:"Message,omitempty"`
	ExtraInfo string `json:"ExtraInfo,omitempty"`
}