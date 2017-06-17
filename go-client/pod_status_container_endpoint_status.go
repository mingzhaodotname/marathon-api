/* 
 * Marathon REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * OpenAPI spec version: 2.0
 * 
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package swagger

type PodStatusContainerEndpointStatus struct {

	// name of the endpoint
	Name StringsName `json:"name"`

	AllocatedHostPort AppNumberPort `json:"allocatedHostPort,omitempty"`

	// true if a health check is defined for this endpoint and is passing
	Healthy bool `json:"healthy,omitempty"`
}