# AppAppContainerDockerContainer

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credential** | [**AppAppContainerDockerCredentials**](app.appContainer.DockerCredentials.md) | The credentials to fetch this container. Please note: this property is supported only with the Mesos containerizer, not the docker containerizer.  | [optional] [default to null]
**PullConfig** | [**AppAppContainerDockerDockerPullConfig**](app.appContainer.docker.DockerPullConfig.md) | Docker&#39;s config.json given as a secret name into a secret store which corresponding value is a content of &#x60;~/.docker/config.json&#x60;. It is supported only with Mesos containerizer.  | [optional] [default to null]
**ForcePullImage** | **bool** | The container will be pulled, regardless if it is already available on the local system.  | [optional] [default to null]
**Image** | **string** | The name of the docker image to use | [default to null]
**Network** | [**AppAppContainerDockerNetwork**](app.appContainer.DockerNetwork.md) |  | [optional] [default to null]
**Parameters** | [**[]AppAppContainerDockerParameter**](app.appContainer.DockerParameter.md) | Allowing arbitrary parameters to be passed to docker CLI. Note that anything passed to this field is not guaranteed to be supported moving forward, as we might move away from the docker CLI.  | [optional] [default to null]
**PortMappings** | [**[]AppAppContainerContainerPortMapping**](app.appContainer.ContainerPortMapping.md) |  | [optional] [default to null]
**Privileged** | **bool** | Run this docker image in privileged mode Please note: this property is supported only with the docker containerizer, not the Mesos containerizer.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


