/*
 * CPQ
 *
 * CPQ Web API Documentation
 *
 * API version: 2.9.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type CloneConfigurationItem struct {
	// The end date which is to be set on the newly cloned contract. If not supplied the endDate from the original contract is used. If this value is supplied it must fall chronologically after the startDate
	EndDate string `json:"endDate,omitempty"`
	// The fc_Id value to be set on the newly cloned contract. This value must be unique across all supplied configuration items and also be unique with respect to existing entries in the database
	FcId string `json:"fc_Id,omitempty"`
	// Indicates if the selection rules from the original contract are to be kept and modified when it is cloned. If specified as true then the contents of the ruleTypeItems array are taken into account, otherwise a draft rule set is created with just the offerId and discountGuid values kept from the original contract
	KeepRules bool `json:"keepRules,omitempty"`
	// The name which is to be set on the newly cloned contract. This value must be unique across all supplied configuration items and if not supplied will result in the cloned contract having the name of the original contract with the text ' - (clone)' appended
	Name string `json:"name,omitempty"`
	// Describes which selection rules are to be kept on the original offering or standalone discount entries
	RuleTypeItems []CloneRuleTypeItem `json:"ruleTypeItems,omitempty"`
	// The start date which is to be set on the newly cloned contract. If not supplied the startDate from the original contract is used
	StartDate string `json:"startDate,omitempty"`
}
