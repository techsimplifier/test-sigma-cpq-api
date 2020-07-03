# ContractModel

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContractId** | **string** | UUID representing the instance of the document entry in the Contracts collection / table. This is a surrogate key which will be replaced each time the framework contract is added or modified. | [optional] [default to null]
**CreationDate** | **string** | The date on which the Contract was created in ISO 8601 format | [optional] [default to null]
**Document** | [***interface{}**](interface{}.md) | The body of the Framework Contract which contains the selection rules for filtering down a product specification to a framework contract specification | [optional] [default to null]
**EndDate** | **string** | End date from when the Contract is no longer effect represented in ISO 8601 format | [optional] [default to null]
**FcId** | **string** | digit identifier of the Framework Contract. It is possible for multiple documents to have the same identifier as there will be old versions of the contracts in the database | [optional] [default to null]
**Name** | **string** | The (optional) name of the Framework Contract e.g. &#39;Bronze&#39; or &#39;Silver&#39; used to identify | [optional] [default to null]
**StartDate** | **string** | Start date from when the Contract is to take effect represented in ISO 8601 format | [optional] [default to null]
**Status** | **string** | The current status of the Framework Contract | [optional] [default to null]
**Version** | **float32** | The current version of the Framework Contract - upon initial insertion it will be set to 1 | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


