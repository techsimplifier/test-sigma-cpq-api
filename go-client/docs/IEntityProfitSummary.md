# IEntityProfitSummary

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityID** | **string** | Unique identifier for the entity in relation to the specification | [optional] [default to null]
**InstanceIDs** | **[]string** | A list of unique identifiers for each instance that exists in the candidate. | [optional] [default to null]
**Recurring** | [**[]IPeriodProfitSummary**](IPeriodProfitSummary.md) | The total summary for each recurring period. | [optional] [default to null]
**NonRecurring** | [**[]IProfitSummary**](IProfitSummary.md) | The total non recurring summary. | [optional] [default to null]
**Quantity** | **float32** | Amount of instances of the given entity | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


