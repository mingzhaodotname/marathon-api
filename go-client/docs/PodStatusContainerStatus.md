# PodStatusContainerStatus

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | [**StringsName**](strings.Name.md) | name of the container specification (within the pod) | [default to null]
**Status** | [**PodStatusContainerState**](podStatus.ContainerState.md) |  | [default to null]
**StatusSince** | [**time.Time**](time.Time.md) | Time at which the status code was last modified.  | [default to null]
**Message** | **string** | Human-friendly explanation for the container&#39;s current status.  | [optional] [default to null]
**Conditions** | [**[]PodStatusStatusCondition**](podStatus.StatusCondition.md) | Set of status conditions that apply to this container.  | [optional] [default to null]
**ContainerId** | **string** | Unique ID of this container in the cluster. TODO(jdef) Probably represents the Mesos task ID.  | [optional] [default to null]
**Endpoints** | [**[]PodStatusContainerEndpointStatus**](podStatus.ContainerEndpointStatus.md) |  | [optional] [default to null]
**Resources** | [**PodResourcesResources**](pod.resources.Resources.md) | Resources in use by this container.  | [optional] [default to null]
**Termination** | [**PodStatusContainerTerminationState**](podStatus.ContainerTerminationState.md) |  | [optional] [default to null]
**LastUpdated** | [**time.Time**](time.Time.md) | Time that this status was last checked and updated (even if nothing changed)  | [default to null]
**LastChanged** | [**time.Time**](time.Time.md) | Time that this status was last modified (some aspect of status did change)  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


