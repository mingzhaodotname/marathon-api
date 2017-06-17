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

type PodVolumesVolumeMount struct {

	// The name of the volume to reference.
	Name StringsName `json:"name"`

	// The path inside the container at which the volume is mounted.
	MountPath StringsPath `json:"mountPath"`

	ReadOnly bool `json:"readOnly,omitempty"`
}
