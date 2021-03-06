/*
 * CPQ
 *
 * CPQ Web API Documentation
 *
 * API version: 2.9.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type PutQuoteItem struct {
	// The last time quote was updated. It should not contain time zone.
	QuoteLastUpdated string `json:"quoteLastUpdated,omitempty"`
	Item *QuoteItem `json:"item,omitempty"`
}
