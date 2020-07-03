# ProductCandidateForPricing

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ID** | **string** | The unique ID of the product candidate. Will be set during the database write operation if not provided | [optional] [default to null]
**EntityID** | **string** | ID of the entity according to its specification | [optional] [default to null]
**ChildEntity** | [**[]IEntity**](IEntity.md) | An array of related entities | [optional] [default to null]
**CharacteristicUse** | [**[]ErrorUnknown**](.md) | An array of characteristic uses. Can be configurable characteristic or fact instances | [optional] [default to null]
**RateAttribute** | [**[]ErrorUnknown**](.md) | An array of rating attributes | [optional] [default to null]
**ConfiguredValue** | [**[]ErrorUnknown**](.md) | An array of configured value uses. These are user defined characteristic instances | [optional] [default to null]
**LinkedEntity** | [**[]ErrorUnknown**](.md) | An optional array of strings that represent other entities that this product is linked to | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


