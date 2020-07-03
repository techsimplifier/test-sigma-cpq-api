/*
 * CPQ
 *
 * CPQ Web API Documentation
 *
 * API version: 2.9.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type QuoteItemCopy struct {
	// Optional parameter for the copied item name
	Name string `json:"name,omitempty"`
	// The quote version number represented as the last updated timestamp
	QuoteLastUpdated string `json:"quoteLastUpdated,omitempty"`
}
