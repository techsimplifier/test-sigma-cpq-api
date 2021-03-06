/*
 * CPQ
 *
 * CPQ Web API Documentation
 *
 * API version: 2.9.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type UpdatedQuoteItem struct {
	QuoteId string `json:"quoteId,omitempty"`
	ItemId string `json:"itemId,omitempty"`
	Name string `json:"name,omitempty"`
	// The date item was updated
	Updated string `json:"updated,omitempty"`
}
