# PodContainerPodContainer

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | [**StringsName**](strings.Name.md) | The name of this container | [default to null]
**Exec** | [**PodContainerMesosExec**](pod.container.MesosExec.md) |  | [optional] [default to null]
**Resources** | [**PodResourcesResources**](pod.resources.Resources.md) | The resources to allocate to the container. | [default to null]
**Endpoints** | [**[]AppNetworkEndpoint**](app.network.Endpoint.md) | The ports needed to run this container | [optional] [default to null]
**Image** | [**PodContainerImage**](pod.container.Image.md) | The filesystem image to populate the container with | [optional] [default to null]
**Environment** | [**AppEnvEnvVars**](app.env.EnvVars.md) |  | [optional] [default to null]
**User** | **string** | The OS user to use to run the tasks on the agent. Values here override a \&quot;user\&quot; value specified in the pod definition.  | [optional] [default to null]
**HealthCheck** | [**AppHealthHealthCheck**](app.health.HealthCheck.md) | All healthchecks to perform on the container | [optional] [default to null]
**VolumeMounts** | [**[]PodVolumesVolumeMount**](pod.volumes.VolumeMount.md) | All volume mounts | [optional] [default to null]
**Artifacts** | [**[]AppArtifactArtifact**](app.artifact.Artifact.md) | All artifacts to download before the task runs | [optional] [default to null]
**Labels** | [**AppLabelKvLabels**](app.label.KVLabels.md) | Metadata as key/value pair. Useful when passing directives to be interpreted by Mesos modules.  | [optional] [default to null]
**Lifecycle** | [**PodContainerLifecycle**](pod.container.Lifecycle.md) |  | [optional] [default to null]
**Tty** | [**AppContainerCommonsTty**](app.containerCommons.TTY.md) | Describes the information about (pseudo) TTY that can be attached to the process of this container.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


