# PodPodUpgradeStrategy

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MinimumHealthCapacity** | **float32** | A number between 0and 1 that is multiplied with the instance count. This is the minimum number of healthy nodes that do not sacrifice overall application purpose. Marathon will make sure, during the upgrade process, that at any point of time this number of healthy instances are up.  | [optional] [default to null]
**MaximumOverCapacity** | **float32** | A number between 0 and 1 which is multiplied with the instance count. This is the maximum number of additional instances launched at any point of time during the upgrade process.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


