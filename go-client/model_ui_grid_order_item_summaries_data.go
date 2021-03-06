/*
 * CPQ
 *
 * CPQ Web API Documentation
 *
 * API version: 2.9.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type UiGridOrderItemSummariesData struct {
	// The length of total summary items
	TotalItems float32 `json:"TotalItems,omitempty"`
	OrderItems []UiGridOrderItemSummary `json:"OrderItems,omitempty"`
}
