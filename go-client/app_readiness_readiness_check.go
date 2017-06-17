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

// Query these readiness checks to determine if a task is ready to serve requests. 
type AppReadinessReadinessCheck struct {

	// The name used to identify this readiness check
	Name string `json:"name,omitempty"`

	Protocol StringsHttpScheme `json:"protocol,omitempty"`

	// Path to endpoint exposed by the task that will provide readiness status. 
	Path StringsPath `json:"path,omitempty"`

	// Name of the port to query as described in the portDefinitions. 
	PortName StringsLegacyName `json:"portName,omitempty"`

	// Number of seconds to wait between readiness checks. 
	IntervalSeconds int32 `json:"intervalSeconds,omitempty"`

	// Number of seconds after which a health check is considered a failure regardless of the response. Must be smaller than intervalSeconds. 
	TimeoutSeconds int32 `json:"timeoutSeconds,omitempty"`

	// The HTTP(s) status codes to treat as 'ready'
	HttpStatusCodesForReady []int32 `json:"httpStatusCodesForReady,omitempty"`

	// If and only if true, preserve the last readiness check responses and expose them in the API as part of a deployment. 
	PreserveLastResponse bool `json:"preserveLastResponse,omitempty"`
}
