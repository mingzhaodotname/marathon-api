# PodVolumesPersistentVolume

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type_** | [**PodVolumesPersistentVolumeType**](pod.volumes.PersistentVolumeType.md) |  | [optional] [default to null]
**Size** | **int64** | The size of the persistence volume in MB. | [default to null]
**MaxSize** | **int64** | For &#x60;mount&#x60; mesos disk resources, the optional maximum size of an exclusive mount volume to be considered.  | [optional] [default to null]
**Constraints** | [**[]AppConstraintsVolumeConstraint**](app.constraints.VolumeConstraint.md) | Constraints restricting where new persistent volumes should be created. Currently, it is only possible to constrain the path of the disk resource by regular expression.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


