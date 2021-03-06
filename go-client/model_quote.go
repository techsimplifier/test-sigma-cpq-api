/*
 * CPQ
 *
 * CPQ Web API Documentation
 *
 * API version: 2.9.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type Quote struct {
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	// The date created
	Created string `json:"created,omitempty"`
	// The date last updated
	Updated string `json:"updated,omitempty"`
	CustomerRef string `json:"customerRef,omitempty"`
	ContextData *interface{} `json:"contextData,omitempty"`
	CurrentValidation *ValidationResponse `json:"currentValidation,omitempty"`
	QuoteType string `json:"quoteType,omitempty"`
	Items []QuoteItem `json:"items,omitempty"`
	PricingSummary *IQuotePricingSummary `json:"pricingSummary,omitempty"`
	// The orderId the quote was created from (for quotes of type 'change').
	OrderId string `json:"orderId,omitempty"`
	// The Internal ID of the Generic Order generated for this quote on submission.
	CpqOrderId string `json:"cpqOrderId,omitempty"`
	// The Internal ID of the Supplemental Order generated for this quote on submission.
	CpqSupplementalOrderId string `json:"cpqSupplementalOrderId,omitempty"`
	// The version of the order the quote was created from (for quotes of type 'change')
	OrderVersion string `json:"orderVersion,omitempty"`
	// Pricing Issue Date
	PricingDate string `json:"pricingDate,omitempty"`
}
