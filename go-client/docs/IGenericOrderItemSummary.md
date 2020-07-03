# IGenericOrderItemSummary

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ID** | **string** | Unique ID for generic order item instance | [optional] [default to null]
**EntityID** | **string** | ID of the entity according to its specification | [optional] [default to null]
**PortfolioItemID** | **string** | Unique portfolio item ID of the the generic order item | [optional] [default to null]
**ItemAction** | **string** | The portfolio action what will be applied to the Quote | [optional] [default to null]
**ChangeType** | **string** |  | [optional] [default to null]
**EntityPath** | **string** | Path used by Sigma Catalog | [optional] [default to null]
**CharacteristicUse** | [**[]ICharacteristicUse**](ICharacteristicUse.md) |  An array of characteristic uses. Can be configurable characteristic or fact instances | [optional] [default to null]
**ConfiguredValue** | [**[]IConfiguredValueUse**](IConfiguredValueUse.md) | An array of configured value uses. These are user defined characteristic instances | [optional] [default to null]
**Rate** | [***IEntityRate**](IEntityRate.md) |  | [optional] [default to null]
**Children** | [**[]IGenericOrderItemSummary**](IGenericOrderItemSummary.md) | An array of generic order items that are children of the current generic order item | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


