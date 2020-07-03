/*
 * CPQ
 *
 * CPQ Web API Documentation
 *
 * API version: 2.9.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type IGenericOrderItemSummary struct {
	// Unique ID for generic order item instance
	ID string `json:"ID,omitempty"`
	// ID of the entity according to its specification
	EntityID string `json:"EntityID,omitempty"`
	// Unique portfolio item ID of the the generic order item
	PortfolioItemID string `json:"PortfolioItemID,omitempty"`
	// The portfolio action what will be applied to the Quote
	ItemAction string `json:"ItemAction,omitempty"`
	// 
	ChangeType string `json:"ChangeType,omitempty"`
	// Path used by Sigma Catalog
	EntityPath string `json:"EntityPath,omitempty"`
	//  An array of characteristic uses. Can be configurable characteristic or fact instances
	CharacteristicUse []ICharacteristicUse `json:"CharacteristicUse,omitempty"`
	// An array of configured value uses. These are user defined characteristic instances
	ConfiguredValue []IConfiguredValueUse `json:"ConfiguredValue,omitempty"`
	Rate *IEntityRate `json:"Rate,omitempty"`
	// An array of generic order items that are children of the current generic order item
	Children []IGenericOrderItemSummary `json:"Children,omitempty"`
}