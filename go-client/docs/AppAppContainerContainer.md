# AppAppContainerContainer

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type_** | [**AppAppContainerEngineType**](app.appContainer.EngineType.md) |  | [default to null]
**Docker** | [**AppAppContainerDockerContainer**](app.appContainer.DockerContainer.md) |  | [optional] [default to null]
**Appc** | [**AppAppContainerAppCContainer**](app.appContainer.AppCContainer.md) |  | [optional] [default to null]
**Volumes** | [**[]PodVolumesAppVolume**](pod.volumes.AppVolume.md) |  | [optional] [default to null]
**PortMappings** | [**[]AppAppContainerContainerPortMapping**](app.appContainer.ContainerPortMapping.md) | Map container ports to host and service ports when using bridge- or container-mode networking. If left unspecified, Marathon will provide a single, default port mapping. To obtain a container with no port mappings, specify an empty array here.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


