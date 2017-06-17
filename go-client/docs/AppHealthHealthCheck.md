# AppHealthHealthCheck

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Http** | [**AppHealthHttpHealthCheck**](app.health.HttpHealthCheck.md) |  | [optional] [default to null]
**Tcp** | [**AppHealthTcpHealthCheck**](app.health.TcpHealthCheck.md) |  | [optional] [default to null]
**Exec** | [**AppHealthCommandHealthCheck**](app.health.CommandHealthCheck.md) | Command that executes some health check process. Use with pods requires Mesos v1.2 or higher.  | [optional] [default to null]
**GracePeriodSeconds** | **int32** | Health check failures are ignored within this number of seconds of the task being started or until the task becomes healthy for the first time.  | [optional] [default to 300]
**IntervalSeconds** | **int32** | Interval between the health checks | [optional] [default to 60]
**MaxConsecutiveFailures** | **int32** | Number of consecutive failures until the task will be killed | [optional] [default to 3]
**TimeoutSeconds** | **int32** | Amount of time to wait for the health check to complete. | [optional] [default to 20]
**DelaySeconds** | **int32** | Amount of time to wait until starting the health checks. | [optional] [default to 15]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


