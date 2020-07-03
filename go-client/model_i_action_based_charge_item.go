/*
 * CPQ
 *
 * CPQ Web API Documentation
 *
 * API version: 2.9.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type IActionBasedChargeItem struct {
	ChildEntity []IPricingServiceResponseEntity `json:"ChildEntity,omitempty"`
	Rate *IEntityRate `json:"Rate,omitempty"`
	// The type of charge for the action based charge item
	ChargeType string `json:"ChargeType,omitempty"`
	// The exact type of the action based charge item
	ExactType string `json:"ExactType,omitempty"`
	// The price grouping of the action based charge item
	PriceGrouping string `json:"PriceGrouping,omitempty"`
	// The name of the action based charge item
	ChargeName string `json:"ChargeName,omitempty"`
}