/*
 * CPQ
 *
 * CPQ Web API Documentation
 *
 * API version: 2.9.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type IEntityLink struct {
	// GUID of the EntityLink as defined on the Product Spec
	LinkTypeID string `json:"LinkTypeID,omitempty"`
	// The action to perform against this EntityLink. Only available on EntityLinks in Order Items
	Action string `json:"Action,omitempty"`
	ChangeType string `json:"ChangeType,omitempty"`
	// A list of all the portfolio items that this EntityLink targets
	Links []ILinkTarget `json:"Links,omitempty"`
	ItemSource string `json:"ItemSource,omitempty"`
}
