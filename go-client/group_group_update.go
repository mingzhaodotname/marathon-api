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

import (
	"time"
)

type GroupGroupUpdate struct {

	Id StringsPathId `json:"id"`

	Apps []AppApp `json:"apps,omitempty"`

	Groups []GroupGroupUpdate `json:"groups,omitempty"`

	Dependencies []StringsPathId `json:"dependencies,omitempty"`

	ScaleBy float64 `json:"scaleBy,omitempty"`

	Version time.Time `json:"version,omitempty"`
}
