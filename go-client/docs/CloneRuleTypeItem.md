# CloneRuleTypeItem

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | **string** | The identifier of either the offering or standalone discount on the original contract. | [optional] [default to null]
**IsOffering** | **bool** | Indicates which entity type the identifier refers to. If &#39;true&#39; then it refers to an Offering entry and if &#39;false&#39; then it refers to a Standalone Discount entry. | [optional] [default to null]
**RuleTypes** | **[]string** | The list of selection rules which are to be copied to the cloned rule set from the original offering or standalone discount entry. For Standalone Discounts you can only specify &#39;standaloneDiscount&#39; here, whereas the other values can be used with Offerings. Specifying &#39;all&#39; means that all existing selection rules are to be cloned. An error is returned if you specify a ruleType which does not exist on the original entity | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


