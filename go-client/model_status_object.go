/*
 * CPQ
 *
 * CPQ Web API Documentation
 *
 * API version: 2.9.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type StatusObject struct {
	Version string `json:"version,omitempty"`
	Routes []StatusObjectRoutes `json:"routes,omitempty"`
	DbStatus *StatusObjectDbStatus `json:"dbStatus,omitempty"`
}
