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

// GroupInfo is a status report of a tree (or subtree) of the root-group hierarchy within Marathon. It is never used to update groups, it is only for reporting. TODO App should be AppInfo here, once we have a RAML type for that. 
type GroupGroupInfo struct {

	Id StringsPathId `json:"id"`

	Apps []AppApp `json:"apps,omitempty"`

	Pods []PodStatusPodStatus `json:"pods,omitempty"`

	// Groups can build a tree. Each group can contain sub-groups. The sub-groups are defined here. 
	Groups []GroupGroupInfo `json:"groups,omitempty"`

	// A list of services that this group depends on. An order is derived from the dependencies for performing start/stop and upgrade of the application. For example, an application /a relies on the services /b which itself relies on /c. To start all 3 applications, first /c is started than /b than /a. 
	Dependencies []StringsPathId `json:"dependencies,omitempty"`

	Version time.Time `json:"version,omitempty"`
}