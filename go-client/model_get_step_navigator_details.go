/*
 * CPQ
 *
 * CPQ Web API Documentation
 *
 * API version: 2.9.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type GetStepNavigatorDetails struct {
	Name string `json:"name,omitempty"`
	QuoteNumber string `json:"quoteNumber,omitempty"`
	QuoteType float32 `json:"quoteType,omitempty"`
	PricingDate string `json:"pricingDate,omitempty"`
	CurrentValidation *GetStepNavigatorDetailsCurrentValidation `json:"currentValidation,omitempty"`
}