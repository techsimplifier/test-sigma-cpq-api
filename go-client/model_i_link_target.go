/*
 * CPQ
 *
 * CPQ Web API Documentation
 *
 * API version: 2.9.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ILinkTarget struct {
	// The id of the portfolio item that is linked
	PortfolioItemID string `json:"PortfolioItemID,omitempty"`
	// The action to be performed on this LinkTarget. Only available in Order Items
	Action string `json:"Action,omitempty"`
}
