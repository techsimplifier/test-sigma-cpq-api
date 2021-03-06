/*
 * CPQ
 *
 * CPQ Web API Documentation
 *
 * API version: 2.9.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type CloneRuleTypeItem struct {
	// The identifier of either the offering or standalone discount on the original contract.
	Identifier string `json:"identifier,omitempty"`
	// Indicates which entity type the identifier refers to. If 'true' then it refers to an Offering entry and if 'false' then it refers to a Standalone Discount entry.
	IsOffering bool `json:"isOffering,omitempty"`
	// The list of selection rules which are to be copied to the cloned rule set from the original offering or standalone discount entry. For Standalone Discounts you can only specify 'standaloneDiscount' here, whereas the other values can be used with Offerings. Specifying 'all' means that all existing selection rules are to be cloned. An error is returned if you specify a ruleType which does not exist on the original entity
	RuleTypes []string `json:"ruleTypes,omitempty"`
}
