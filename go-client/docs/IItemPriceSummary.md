# IItemPriceSummary

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IgnoredCharges** | [**[]interface{}**](interface{}.md) | Summary information of ignored charges for the candidate | [optional] [default to null]
**RecurringCharges** | [**[]IRecurringChargeSummary**](IRecurringChargeSummary.md) | Summary information of the recurring charges for the candidate.This has to be any, as the pricing service could return an object or an array | [optional] [default to null]
**NonRecurringCharges** | [***INonRecurringChargeSummary**](INonRecurringChargeSummary.md) |  | [optional] [default to null]
**ActionBasedCharges** | [***INonRecurringActionBasedChargeSummary**](INonRecurringActionBasedChargeSummary.md) |  | [optional] [default to null]
**StandaloneCharges** | [***IStandaloneActionBasedChargeSummary**](IStandaloneActionBasedChargeSummary.md) |  | [optional] [default to null]
**IgnoredCosts** | [**[]interface{}**](interface{}.md) | Summary information of ignored costs for the candidate | [optional] [default to null]
**RecurringCosts** | [**[]IRecurringCostSummary**](IRecurringCostSummary.md) | Summary information of the recurring costs for the candidate.This has to be any, as the pricing service could return an object or an array | [optional] [default to null]
**NonRecurringCosts** | [***INonRecurringCostSummary**](INonRecurringCostSummary.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


