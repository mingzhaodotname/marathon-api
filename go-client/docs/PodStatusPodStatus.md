# PodStatusPodStatus

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | [**StringsPathId**](strings.PathId.md) |  | [default to null]
**Spec** | [**PodPod**](pod.Pod.md) | The latest version of the pod specification (that has the same pod ID).  | [default to null]
**Status** | [**PodStatusPodState**](podStatus.PodState.md) |  | [default to null]
**StatusSince** | [**time.Time**](time.Time.md) | Time at which the status code was last modified.  | [default to null]
**Message** | **string** | Human-friendly explanation for reason of the current status.  | [optional] [default to null]
**Instances** | [**[]PodStatusPodInstanceStatus**](podStatus.PodInstanceStatus.md) |  | [optional] [default to null]
**TerminationHistory** | [**[]PodStatusTerminationHistory**](podStatus.TerminationHistory.md) | List of most recent instance terminations. TODO(jdef) determine how many items might show up here; current thinking is .. not many  | [optional] [default to null]
**LastUpdated** | [**time.Time**](time.Time.md) | Time that this status object was last checked and updated (even if nothing changed)  | [default to null]
**LastChanged** | [**time.Time**](time.Time.md) | Time that this status object was last modified (some aspect of status did change)  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


