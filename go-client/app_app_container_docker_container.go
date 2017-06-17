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

type AppAppContainerDockerContainer struct {

	// The credentials to fetch this container. Please note: this property is supported only with the Mesos containerizer, not the docker containerizer. 
	Credential AppAppContainerDockerCredentials `json:"credential,omitempty"`

	// Docker's config.json given as a secret name into a secret store which corresponding value is a content of `~/.docker/config.json`. It is supported only with Mesos containerizer. 
	PullConfig AppAppContainerDockerDockerPullConfig `json:"pullConfig,omitempty"`

	// The container will be pulled, regardless if it is already available on the local system. 
	ForcePullImage bool `json:"forcePullImage,omitempty"`

	// The name of the docker image to use
	Image string `json:"image"`

	Network AppAppContainerDockerNetwork `json:"network,omitempty"`

	// Allowing arbitrary parameters to be passed to docker CLI. Note that anything passed to this field is not guaranteed to be supported moving forward, as we might move away from the docker CLI. 
	Parameters []AppAppContainerDockerParameter `json:"parameters,omitempty"`

	PortMappings []AppAppContainerContainerPortMapping `json:"portMappings,omitempty"`

	// Run this docker image in privileged mode Please note: this property is supported only with the docker containerizer, not the Mesos containerizer. 
	Privileged bool `json:"privileged,omitempty"`
}
