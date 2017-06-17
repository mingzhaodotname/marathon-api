# AppNetworkEndpoint

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | [**StringsName**](strings.Name.md) | Name of this port. Should be unique in the context of the pod.  | [default to null]
**ContainerPort** | [**AppNumberPort**](app.number.Port.md) | The container port that a task&#39;s process is actually listening on. Required if the network mode is container  | [optional] [default to null]
**HostPort** | [**AppNumberAnyPort**](app.number.AnyPort.md) | Mapped port on the host. All host ports are allocated from resource offers.  | [optional] [default to null]
**Protocol** | [**[]StringsName**](strings.Name.md) | The protocol of this port, advertised in Mesos DiscoveryInfo (DI). Specifying more than one protocol here will result in the generation of multiple Port DI entries.  | [optional] [default to null]
**Labels** | [**AppLabelKvLabels**](app.label.KVLabels.md) | Metadata as key/value pair. Possible uses include VIPs, per-network-interface binding  | [optional] [default to null]
**NetworkNames** | [**[]StringsName**](strings.Name.md) | List of the container networks associated with this endpoint. If absent, then this endpoint is associated with all defined container networks (for this application). A single item list is mandatory when &#x60;hostPort&#x60; is specified and multiple container networks are defined.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


