# AppAppResidency

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RelaunchEscalationTimeoutSeconds** | **int64** | When a task using persistent local volumes cannot be restarted on the agent it&#39;s been pinned to, Marathon will try to launch this task on another node after this timeout. Defaults to 3600 (one hour).\&quot;,  | [default to 3600]
**TaskLostBehavior** | [**StringsTaskLostBehavior**](strings.TaskLostBehavior.md) |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


