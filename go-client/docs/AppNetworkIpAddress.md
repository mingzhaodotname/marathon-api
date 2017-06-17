# AppNetworkIpAddress

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Discovery** | [**AppNetworkIpDiscovery**](app.network.IpDiscovery.md) | DEPRECATED. Please try to use portMappings instead.  | [optional] [default to null]
**Groups** | **[]string** | Ignored by Marathon, only exists here to preserve API compat (for now). | [optional] [default to null]
**Labels** | [**AppLabelKvLabels**](app.label.KVLabels.md) |  | [optional] [default to null]
**NetworkName** | **string** | If present, declares the name of the network that the container should join. In practice it is up to the IPAM to decide how to interpret this field  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


