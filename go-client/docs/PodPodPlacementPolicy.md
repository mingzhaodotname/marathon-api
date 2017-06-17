# PodPodPlacementPolicy

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Constraints** | [**[]AppConstraintsConstraint**](app.constraints.Constraint.md) |  | [optional] [default to null]
**AcceptedResourceRoles** | **[]string** | A list of resource roles. Marathon considers only resource offers with roles in this list for launching tasks of this app. If you do not specify this, Marathon considers all resource offers with roles that have been configured by the &#x60;--default_accepted_resource_roles&#x60; command line flag. If no &#x60;--default_accepted_resource_roles&#x60; was given on startup, Marathon considers all resource offers. To register Marathon for a role, you need to specify the &#x60;--mesos_role&#x60; command line flag on startup. If you want to assign all resources of a slave to a role, you can use the &#x60;--default_role&#x60; argument when starting up the slave. If you need a more fine-grained configuration, you can use the &#x60;--resources&#x60; argument to specify resource shares per role. See [the Mesos attribute and resources documentation](http://mesos.apache.org/documentation/latest/attributes-resources/) for details  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


