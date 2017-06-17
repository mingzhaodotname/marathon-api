# PodStatusPodInstanceStatus

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique ID of this pod instance in the cluster. TODO(jdef) Probably represents the Mesos executor ID.  | [default to null]
**Status** | [**PodStatusPodInstanceState**](podStatus.PodInstanceState.md) |  | [default to null]
**StatusSince** | [**time.Time**](time.Time.md) | Time at which the status code was last modified.  | [default to null]
**Message** | **string** | Human-friendly explanation for reason of the current status.  | [optional] [default to null]
**Conditions** | [**[]PodStatusStatusCondition**](podStatus.StatusCondition.md) | Set of status conditions that apply to this pod instance.  | [optional] [default to null]
**AgentHostname** | **string** | Hostname that this instance was launched on. May be an IP address if the agent was configured to advertise its hostname that way.  | [optional] [default to null]
**Resources** | [**PodResourcesResources**](pod.resources.Resources.md) | Sum of all resources allocated for this pod instance. May include additional, system-allocated resources for the default executor.  | [optional] [default to null]
**Networks** | [**[]PodStatusNetworkStatus**](podStatus.NetworkStatus.md) | Status of the networks to which this instance is attached.  | [optional] [default to null]
**Containers** | [**[]PodStatusContainerStatus**](podStatus.ContainerStatus.md) | status for each running container of this instance. | [optional] [default to null]
**SpecReference** | [**StringsUri**](strings.Uri.md) | Location of the version of the pod specification this instance was created from.  | [optional] [default to null]
**LastUpdated** | [**time.Time**](time.Time.md) | Time that this status was last checked and updated (even if nothing changed)  | [default to null]
**LastChanged** | [**time.Time**](time.Time.md) | Time that this status was last modified (some aspect of status did change)  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


