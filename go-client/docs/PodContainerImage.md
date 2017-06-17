# PodContainerImage

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | [**PodContainerImageType**](pod.container.ImageType.md) |  | [default to null]
**Id** | [**StringsPath**](strings.Path.md) | address/identifier of the file system image | [default to null]
**PullConfig** | [**AppAppContainerDockerDockerPullConfig**](app.appContainer.docker.DockerPullConfig.md) | Name of a secret whose value contains a stringified Docker config.json  | [optional] [default to null]
**ForcePull** | **bool** | true if the image should always be pulled | [optional] [default to null]
**Labels** | [**AppLabelKvLabels**](app.label.KVLabels.md) | Used during image discovery and dependency resolution. Several well-known labels are defined: version: the version of this image os: (default: linux) operating system arch: (default: amd64) architecture  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


