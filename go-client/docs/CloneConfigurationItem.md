# CloneConfigurationItem

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndDate** | **string** | The end date which is to be set on the newly cloned contract. If not supplied the endDate from the original contract is used. If this value is supplied it must fall chronologically after the startDate | [optional] [default to null]
**FcId** | **string** | The fc_Id value to be set on the newly cloned contract. This value must be unique across all supplied configuration items and also be unique with respect to existing entries in the database | [optional] [default to null]
**KeepRules** | **bool** | Indicates if the selection rules from the original contract are to be kept and modified when it is cloned. If specified as true then the contents of the ruleTypeItems array are taken into account, otherwise a draft rule set is created with just the offerId and discountGuid values kept from the original contract | [optional] [default to null]
**Name** | **string** | The name which is to be set on the newly cloned contract. This value must be unique across all supplied configuration items and if not supplied will result in the cloned contract having the name of the original contract with the text &#39; - (clone)&#39; appended | [optional] [default to null]
**RuleTypeItems** | [**[]CloneRuleTypeItem**](CloneRuleTypeItem.md) | Describes which selection rules are to be kept on the original offering or standalone discount entries | [optional] [default to null]
**StartDate** | **string** | The start date which is to be set on the newly cloned contract. If not supplied the startDate from the original contract is used | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


