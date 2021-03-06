/*
 * CPQ
 *
 * CPQ Web API Documentation
 *
 * API version: 2.9.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type CardinalityRule struct {
	// The new value for maximum cardinality
	Max float32 `json:"max,omitempty"`
	// The new value for minimum cardinality
	Min float32 `json:"min,omitempty"`
	// The path to the product relation (from the offering root) that is to have its cardinality changed
	Path []string `json:"path,omitempty"`
}
