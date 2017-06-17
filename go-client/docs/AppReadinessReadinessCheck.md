# AppReadinessReadinessCheck

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name used to identify this readiness check | [optional] [default to null]
**Protocol** | [**StringsHttpScheme**](strings.HttpScheme.md) |  | [optional] [default to null]
**Path** | [**StringsPath**](strings.Path.md) | Path to endpoint exposed by the task that will provide readiness status.  | [optional] [default to null]
**PortName** | [**StringsLegacyName**](strings.LegacyName.md) | Name of the port to query as described in the portDefinitions.  | [optional] [default to null]
**IntervalSeconds** | **int32** | Number of seconds to wait between readiness checks.  | [optional] [default to 30]
**TimeoutSeconds** | **int32** | Number of seconds after which a health check is considered a failure regardless of the response. Must be smaller than intervalSeconds.  | [optional] [default to 10]
**HttpStatusCodesForReady** | **[]int32** | The HTTP(s) status codes to treat as &#39;ready&#39; | [optional] [default to null]
**PreserveLastResponse** | **bool** | If and only if true, preserve the last readiness check responses and expose them in the API as part of a deployment.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


