/*
 * CPQ
 *
 * CPQ Web API Documentation
 *
 * API version: 2.9.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type IQuotePricingSummary struct {
	TotalPriceSummary *ITotalProfitSummary `json:"TotalPriceSummary,omitempty"`
	PaidTotalPriceSummary *ITotalProfitSummary `json:"PaidTotalPriceSummary,omitempty"`
}
