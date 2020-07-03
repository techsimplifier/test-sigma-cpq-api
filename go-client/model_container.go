/*
 * CPQ
 *
 * CPQ Web API Documentation
 *
 * API version: 2.9.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type Container struct {
	Id string `json:"id,omitempty"`
	QuoteItemIds []string `json:"QuoteItemIds,omitempty"`
	BreakdownSummary *BreakdownSummary `json:"breakdownSummary,omitempty"`
	Breakdown *Breakdown `json:"breakdown,omitempty"`
	Title string `json:"title,omitempty"`
}
