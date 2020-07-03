/*
 * CPQ
 *
 * CPQ Web API Documentation
 *
 * API version: 2.9.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type GetOrderEnrichmentTask200OrderEnrichmentTasks struct {
	ItemId string `json:"itemId,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
	State string `json:"state,omitempty"`
	Progress float32 `json:"progress,omitempty"`
	EnrichmentTasks []GetOrderEnrichmentTask200EnrichmentTasks `json:"enrichmentTasks,omitempty"`
}
