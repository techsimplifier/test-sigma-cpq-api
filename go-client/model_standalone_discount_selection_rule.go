/*
 * CPQ
 *
 * CPQ Web API Documentation
 *
 * API version: 2.9.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type StandaloneDiscountSelectionRule struct {
	// The collection of identifiers for those discount values that are to be retained on the negotiated specification. Each entry is in UUID format
	DiscountRates []string `json:"discountRates,omitempty"`
}
