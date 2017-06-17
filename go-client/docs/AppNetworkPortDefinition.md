# AppNetworkPortDefinition

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Port** | [**AppNumberAnyPort**](app.number.AnyPort.md) | If requirePorts is set to true, then this port number will be used on the agent host. Otherwise if the requirePorts is set to false and this port number is not zero, then it will be used as a service port and a dynamic port will be used on the agent host. If it is set to zero, a dynamic port will be used on the host and a unique service port will be assigned by Marathon.  | [optional] [default to null]
**Labels** | [**AppLabelKvLabels**](app.label.KVLabels.md) |  | [optional] [default to null]
**Name** | [**StringsLegacyName**](strings.LegacyName.md) | Name of the service hosted on this port. If provided, it must be unique over all port definitions.  | [optional] [default to null]
**Protocol** | [**StringsNetworkProtocol**](strings.NetworkProtocol.md) | If this port is used for tcp or udp or both. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


