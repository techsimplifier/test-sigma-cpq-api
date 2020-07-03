/*
 * CPQ
 *
 * CPQ Web API Documentation
 *
 * API version: 2.9.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type OrderCancellationResponse struct {
	Success bool `json:"success,omitempty"`
	Status string `json:"status,omitempty"`
	OrderItems []OrderItems `json:"orderItems,omitempty"`
	ErrorCode string `json:"errorCode,omitempty"`
	ErrorContext *OrderErrorContext `json:"errorContext,omitempty"`
}
