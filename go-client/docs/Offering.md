# Offering

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardinalityRules** | [**[]CardinalityRule**](CardinalityRule.md) | The collection of entities that are to be retained on the negotiated specification. | [optional] [default to null]
**CharacteristicSelectionRules** | [**[]ConfigurableCharacteristicSelectionRule**](ConfigurableCharacteristicSelectionRule.md) | The collection of TSpecCharUse structure that are to be retained on the negotiated specification | [optional] [default to null]
**ChargeSelectionRules** | [**[]RateableItemSelectionRule**](RateableItemSelectionRule.md) | The collection of charges that are to be retained on the negotiated specification | [optional] [default to null]
**DiscountSelectionRules** | [**[]RateableItemSelectionRule**](RateableItemSelectionRule.md) | The collection of discounts that are to be retained on the negotiated specification | [optional] [default to null]
**FactSelectionRules** | [**[]FactSelectionRule**](FactSelectionRule.md) | The collection of TConfigurableFact structures that are to be retained on the negotiated specification | [optional] [default to null]
**OfferId** | **string** | The identifier of the product specification which is in UUID format. | [optional] [default to null]
**OfferName** | **string** | The name of the offer which has been selected - this is purely a UI concern | [optional] [default to null]
**UserDefinedCharacteristicSelectionRules** | [**[]UserCharacteristicSelectionRule**](UserCharacteristicSelectionRule.md) | The collection of TUserDefinedCharacteristics that are to be retained on the negotiated specification. This structure can be used to specify both internal and externally defined user characteristics. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


