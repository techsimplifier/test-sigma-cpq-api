/*
 * CPQ
 *
 * CPQ Web API Documentation
 *
 * API version: 2.9.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type SparseCreateQuoteItemResponse struct {
	// The unique ID of the quote in which sparse items are created
	QuoteId string `json:"quoteId,omitempty"`
	// Unique quote items ids
	SparseQuoteItemIds []string `json:"sparseQuoteItemIds,omitempty"`
}