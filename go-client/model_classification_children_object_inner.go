/*
 * CPQ
 *
 * CPQ Web API Documentation
 *
 * API version: 2.9.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ClassificationChildrenObjectInner struct {
	//  Indicates whether or not there is a nested structure underneath this entry
	IsLeaf bool `json:"isLeaf,omitempty"`
	// The unique identifier for the classification entry
	Guid string `json:"guid,omitempty"`
	// The name of the classification entry to display
	Name string `json:"name,omitempty"`
	// The name of the schema type of which the classification is an instance of
	SchemaXsdName string `json:"schemaXsdName,omitempty"`
	//  Indicates a sorting value
	Sort string `json:"sort,omitempty"`
	//  Another identifier for the classification entry
	Uid string `json:"uid,omitempty"`
	Children *ClassificationChildrenObject `json:"children,omitempty"`
}
