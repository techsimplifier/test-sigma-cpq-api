/*
 * CPQ
 *
 * CPQ Web API Documentation
 *
 * API version: 2.9.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type OrderCandidateRequest struct {
	ID string `json:"ID,omitempty"`
	OrderCandidate *OrderCandidate `json:"OrderCandidate,omitempty"`
	CustomerPortfolio *interface{} `json:"CustomerPortfolio,omitempty"`
	// The activation date
	ActivationDate string `json:"ActivationDate,omitempty"`
}