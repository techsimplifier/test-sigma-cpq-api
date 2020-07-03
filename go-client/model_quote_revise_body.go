/*
 * CPQ
 *
 * CPQ Web API Documentation
 *
 * API version: 2.9.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type QuoteReviseBody struct {
	// id of the quote from the external CRM
	ExternalQuoteId string `json:"externalQuoteId,omitempty"`
	// The new version number of the quote
	NewQuoteVersion float32 `json:"newQuoteVersion,omitempty"`
}