# PodPodSchedulingBackoffStrategy

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Backoff** | **float32** | The initial backoff (seconds) applied when a launched instance fails. | [optional] [default to null]
**BackoffFactor** | **float32** | The factor applied to the current backoff to determine the new backoff. | [optional] [default to null]
**MaxLaunchDelay** | **float32** | The maximum backoff (seconds) applied when subsequent failures are detected. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


