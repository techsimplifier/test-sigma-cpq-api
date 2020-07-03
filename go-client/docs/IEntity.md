# IEntity

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ID** | **string** | Unique ID for the entity instance | [optional] [default to null]
**EntityID** | **string** | ID of the entity according to its specification | [optional] [default to null]
**ChildEntity** | [**[]IEntity**](IEntity.md) | An array of related entities | [optional] [default to null]
**CharacteristicUse** | [**[]ICharacteristicUse**](ICharacteristicUse.md) |  An array of characteristic uses. Can be configurable characteristic or fact instances | [optional] [default to null]
**RateAttribute** | [**[]IRatingAttribute**](IRatingAttribute.md) |  An array of rating attributes | [optional] [default to null]
**ConfiguredValue** | [**[]IConfiguredValueUse**](IConfiguredValueUse.md) |  An array of configured value uses. These are user defined characteristic instances | [optional] [default to null]
**IsNewForCustomer** | **bool** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


