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

// Resource Allocations
type PodResourcesResources struct {

	// The number of CPU shares this pod needs per instance. This number does not have to be integer, but can be a fraction.
	Cpus float64 `json:"cpus"`

	// The amount of memory in MB that is needed for the pod instance
	Mem float64 `json:"mem"`

	// How much disk space is needed for this application. This number does not have to be an integer, but can be a fraction.
	Disk float64 `json:"disk,omitempty"`

	// The amount of GPU cores that are needed for the pod
	Gpus int32 `json:"gpus,omitempty"`
}
