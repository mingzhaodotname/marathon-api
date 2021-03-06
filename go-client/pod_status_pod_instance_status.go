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

type PodStatusPodInstanceStatus struct {

	// Unique ID of this pod instance in the cluster. TODO(jdef) Probably represents the Mesos executor ID. 
	Id string `json:"id"`

	Status PodStatusPodInstanceState `json:"status"`

	// Time at which the status code was last modified. 
	StatusSince time.Time `json:"statusSince"`

	// Human-friendly explanation for reason of the current status. 
	Message string `json:"message,omitempty"`

	// Set of status conditions that apply to this pod instance. 
	Conditions []PodStatusStatusCondition `json:"conditions,omitempty"`

	// Hostname that this instance was launched on. May be an IP address if the agent was configured to advertise its hostname that way. 
	AgentHostname string `json:"agentHostname,omitempty"`

	// Sum of all resources allocated for this pod instance. May include additional, system-allocated resources for the default executor. 
	Resources PodResourcesResources `json:"resources,omitempty"`

	// Status of the networks to which this instance is attached. 
	Networks []PodStatusNetworkStatus `json:"networks,omitempty"`

	// status for each running container of this instance.
	Containers []PodStatusContainerStatus `json:"containers,omitempty"`

	// Location of the version of the pod specification this instance was created from. 
	SpecReference StringsUri `json:"specReference,omitempty"`

	// Time that this status was last checked and updated (even if nothing changed) 
	LastUpdated time.Time `json:"lastUpdated"`

	// Time that this status was last modified (some aspect of status did change) 
	LastChanged time.Time `json:"lastChanged"`
}
