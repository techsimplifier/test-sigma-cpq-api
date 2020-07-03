# QuoteItem

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [optional] [default to null]
**Name** | **string** |  | [optional] [default to null]
**Created** | **string** | The date item was created | [optional] [default to null]
**QuoteLastUpdated** | **string** | The date item was created | [optional] [default to null]
**ProductId** | **string** |  | [optional] [default to null]
**Description** | **string** |  | [optional] [default to null]
**ProductCandidate** | [***ProductCandidate**](ProductCandidate.md) |  | [optional] [default to null]
**PrePricedCandidate** | [***ProductCandidate**](ProductCandidate.md) |  | [optional] [default to null]
**PortfolioItemId** | **string** | The identifier of the portfolio item that this QuoteItem relates to. | [optional] [default to null]
**OfferSpecification** | [***interface{}**](interface{}.md) | the product specification for this offer | [optional] [default to null]
**CurrentPricing** | [***ICpqPricingResponse**](ICpqPricingResponse.md) |  | [optional] [default to null]
**CurrentValidation** | [***ValidationResponse**](ValidationResponse.md) |  | [optional] [default to null]
**ItemAction** | **string** | The portfolio action which is to be applied to the Quote.If no value is specified then the CPQ API will treat it as an Add. The only values which are allowed are &lt;Add | Update | Delete&gt; | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


