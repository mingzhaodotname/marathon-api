# PodResourcesResources

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cpus** | **float64** | The number of CPU shares this pod needs per instance. This number does not have to be integer, but can be a fraction. | [default to 1.0]
**Mem** | **float64** | The amount of memory in MB that is needed for the pod instance | [default to 128.0]
**Disk** | **float64** | How much disk space is needed for this application. This number does not have to be an integer, but can be a fraction. | [optional] [default to 0.0]
**Gpus** | **int32** | The amount of GPU cores that are needed for the pod | [optional] [default to 0]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


