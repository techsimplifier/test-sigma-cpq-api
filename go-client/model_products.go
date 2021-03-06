/*
 * CPQ
 *
 * CPQ Web API Documentation
 *
 * API version: 2.9.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type Products struct {
	Entity *SparseEntity `json:"entity,omitempty"`
	// An array of sparse characteristics. Can be configurable characteristics or fact instances
	Characteristics []Characteristics `json:"characteristics,omitempty"`
	// An array of user defined characteristic uses
	UserDefinedCharacteristics []UserDefinedCharacteristics `json:"userDefinedCharacteristics,omitempty"`
}
