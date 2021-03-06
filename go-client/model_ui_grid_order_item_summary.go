/*
 * CPQ
 *
 * CPQ Web API Documentation
 *
 * API version: 2.9.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type UiGridOrderItemSummary struct {
	// Unique ID for generic order item instance
	ID string `json:"ID,omitempty"`
	// ID of the entity according to its specification
	EntityID string `json:"EntityID,omitempty"`
	// The Name to of the order Item
	DisplayID string `json:"DisplayID,omitempty"`
	// 
	ChangeType string `json:"ChangeType,omitempty"`
	// The portfolio action what will be applied to the Quote
	ItemAction string `json:"ItemAction,omitempty"`
	// The tree level in which it should be displayed on UI grid tree
	TreeLevel float32 `json:"$$treeLevel,omitempty"`
	// The name of the Characteristic
	CharacteristicName string `json:"CharacteristicName,omitempty"`
	// The value of the Characteristic
	CharacteristicValue string `json:"CharacteristicValue,omitempty"`
	// The Primary Location assigned to this Order Item
	Location string `json:"Location,omitempty"`
}
