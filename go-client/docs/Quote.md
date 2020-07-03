# Quote

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [optional] [default to null]
**Name** | **string** |  | [optional] [default to null]
**Created** | **string** | The date created | [optional] [default to null]
**Updated** | **string** | The date last updated | [optional] [default to null]
**CustomerRef** | **string** |  | [optional] [default to null]
**ContextData** | [***interface{}**](interface{}.md) |  | [optional] [default to null]
**CurrentValidation** | [***ValidationResponse**](ValidationResponse.md) |  | [optional] [default to null]
**QuoteType** | **string** |  | [optional] [default to null]
**Items** | [**[]QuoteItem**](QuoteItem.md) |  | [optional] [default to null]
**PricingSummary** | [***IQuotePricingSummary**](IQuotePricingSummary.md) |  | [optional] [default to null]
**OrderId** | **string** | The orderId the quote was created from (for quotes of type &#39;change&#39;). | [optional] [default to null]
**CpqOrderId** | **string** | The Internal ID of the Generic Order generated for this quote on submission. | [optional] [default to null]
**CpqSupplementalOrderId** | **string** | The Internal ID of the Supplemental Order generated for this quote on submission. | [optional] [default to null]
**OrderVersion** | **string** | The version of the order the quote was created from (for quotes of type &#39;change&#39;) | [optional] [default to null]
**PricingDate** | **string** | Pricing Issue Date | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


