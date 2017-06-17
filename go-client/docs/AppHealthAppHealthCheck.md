# AppHealthAppHealthCheck

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Command** | [**AppHealthCommandCheck**](app.health.CommandCheck.md) |  | [optional] [default to null]
**GracePeriodSeconds** | **int32** | Health check failures are ignored within this number of seconds of the task being started or until the task becomes healthy for the first time.  | [optional] [default to 300]
**IgnoreHttp1xx** | **bool** | Ignore HTTP 1xx responses | [optional] [default to null]
**IntervalSeconds** | **int32** | Number of seconds to wait between health checks | [optional] [default to 60]
**MaxConsecutiveFailures** | **int32** | Number of consecutive health check failures after which the unhealthy task should be killed.  | [optional] [default to 3]
**Path** | **string** | Path to endpoint exposed by the task that will provide health status. Note: only used if protocol &#x3D;&#x3D; HTTP[S].\&quot;  | [optional] [default to null]
**Port** | [**AppNumberAnyPort**](app.number.AnyPort.md) | The specific port to connect to. In case of dynamic ports, see portIndex.  | [optional] [default to null]
**PortIndex** | [**AppNumberAnyPort**](app.number.AnyPort.md) | Index in this app&#39;s ports array to be used for health requests. An index is used so the app can use random ports, like [0, 0, 0] for example, and tasks could be started with port environment variables like $PORT1.  | [optional] [default to null]
**Protocol** | [**AppHealthAppHealthCheckProtocol**](app.health.AppHealthCheckProtocol.md) |  | [optional] [default to null]
**TimeoutSeconds** | **int32** | Number of seconds after which a health check is considered a failure regardless of the response.  | [optional] [default to 20]
**DelaySeconds** | **int32** | Amount of time to wait until starting the health checks. | [optional] [default to 15]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


