# AppUnreachableStrategyUnreachableEnabled

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InactiveAfterSeconds** | **int64** | If an instance is unreachable for longer than inactiveAfter seconds it is marked as inactive. This will trigger a new instance launch. The original task is not expunged yet. Must be less than expungeAfterSeconds.  The default value is set to 5 minutes (300 seconds).  | [optional] [default to 300]
**ExpungeAfterSeconds** | **int64** | If an instance is unreachable for longer than unreachableExpungeAfter seconds it will be expunged.  That means it will be killed if it ever comes back. Instances are usually marked as unreachable before they are expunged but they don&#39;t have to. This value is required to be greater than inactiveAfterSeconds.  The default value is set to 10 minutes (600 seconds).  | [optional] [default to 600]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


