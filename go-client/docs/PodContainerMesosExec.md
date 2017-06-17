# PodContainerMesosExec

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Command** | [**AppHealthCommandMesosCommand**](app.health.command.MesosCommand.md) | Command specification executed by Mesos, not parsed by Marathon. The presence of &#x60;command.shell&#x60; implies &#x60;overrideEntrypoint&#x3D;true&#x60;.  | [default to null]
**OverrideEntrypoint** | **bool** | When true argv[0] will be used as the entrypoint/exec of the container. Otherwise the contents of argv[] are appended as arguments.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


