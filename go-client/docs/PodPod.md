# PodPod

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | [**StringsPathId**](strings.PathId.md) |  | [default to null]
**Labels** | [**AppLabelKvLabels**](app.label.KVLabels.md) | Metadata as key/value pair. Useful when passing directives to be interpreted by Mesos modules.  | [optional] [default to null]
**Version** | [**time.Time**](time.Time.md) | The version of the definition, immutable | [optional] [default to null]
**User** | **string** | The OS user to use to run the tasks on the agent. May be overridden by a container.  | [optional] [default to null]
**Environment** | [**AppEnvEnvVars**](app.env.EnvVars.md) | Environment Variables to set at the pod level. Individual containers may override them  | [optional] [default to null]
**Containers** | [**[]PodContainerPodContainer**](pod.container.PodContainer.md) |  | [default to null]
**Secrets** | [**AppSecretsSecrets**](app.secrets.Secrets.md) |  | [optional] [default to null]
**Volumes** | [**[]PodVolumesPodVolume**](pod.volumes.PodVolume.md) | Volumes defined on a pod level that are mounted into containers | [optional] [default to null]
**Networks** | [**[]AppNetworkNetwork**](app.network.Network.md) | Network settings are defined on a pod level. All containers share the same network stack. At this time, only one stack is supported.  | [optional] [default to null]
**Scaling** | [**PodPodScalingPolicy**](pod.PodScalingPolicy.md) |  | [optional] [default to null]
**Scheduling** | [**PodPodSchedulingPolicy**](pod.PodSchedulingPolicy.md) |  | [optional] [default to null]
**ExecutorResources** | [**PodResourcesExecutorResources**](pod.resources.ExecutorResources.md) | The resources to allocate to the executor. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


