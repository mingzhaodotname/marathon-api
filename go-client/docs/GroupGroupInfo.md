# GroupGroupInfo

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | [**StringsPathId**](strings.PathId.md) |  | [default to null]
**Apps** | [**[]AppApp**](app.App.md) |  | [optional] [default to null]
**Pods** | [**[]PodStatusPodStatus**](podStatus.PodStatus.md) |  | [optional] [default to null]
**Groups** | [**[]GroupGroupInfo**](group.GroupInfo.md) | Groups can build a tree. Each group can contain sub-groups. The sub-groups are defined here.  | [optional] [default to null]
**Dependencies** | [**[]StringsPathId**](strings.PathId.md) | A list of services that this group depends on. An order is derived from the dependencies for performing start/stop and upgrade of the application. For example, an application /a relies on the services /b which itself relies on /c. To start all 3 applications, first /c is started than /b than /a.  | [optional] [default to null]
**Version** | [**time.Time**](time.Time.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


