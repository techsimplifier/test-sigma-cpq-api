/*
 * CPQ
 *
 * CPQ Web API Documentation
 *
 * API version: 2.9.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type FinancialSummaryQuery struct {
	PluginToExecuteIdentifier *PluginName `json:"pluginToExecuteIdentifier,omitempty"`
	GroupBy *GroupBy `json:"groupBy,omitempty"`
	OptionsToInclude *OptionsToInclude `json:"optionsToInclude,omitempty"`
}