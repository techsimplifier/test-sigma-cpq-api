/*
 * CPQ
 *
 * CPQ Web API Documentation
 *
 * API version: 2.9.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type QuoteConvert200Response struct {
	Id string `json:"id"`
	QuoteReference string `json:"quoteReference,omitempty"`
	OrderReference string `json:"orderReference,omitempty"`
	Status string `json:"status,omitempty"`
	OrderCreationDate string `json:"orderCreationDate,omitempty"`
}