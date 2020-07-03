/*
 * CPQ
 *
 * CPQ Web API Documentation
 *
 * API version: 2.9.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type IEntityProfitSummary struct {
	// Unique identifier for the entity in relation to the specification
	EntityID string `json:"EntityID,omitempty"`
	// A list of unique identifiers for each instance that exists in the candidate.
	InstanceIDs []string `json:"InstanceIDs,omitempty"`
	// The total summary for each recurring period.
	Recurring []IPeriodProfitSummary `json:"Recurring,omitempty"`
	// The total non recurring summary.
	NonRecurring []IProfitSummary `json:"NonRecurring,omitempty"`
	// Amount of instances of the given entity
	Quantity float32 `json:"Quantity,omitempty"`
}
