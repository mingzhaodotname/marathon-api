# PodVolumesHostVolume

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | [**StringsName**](strings.Name.md) | The name of the volume to reference. | [default to null]
**Host** | **string** | Absolute path of the file or directory on the agent, or else the relative path of the directory in the executor&#39;s sandbox. Host volumes are useful for mapping directories that exist on the agent apriori, or within the executor sandbox. No resources (Mesos or otherwise) are allocated for these types of volumes.  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


