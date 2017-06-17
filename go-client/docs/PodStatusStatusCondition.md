# PodStatusStatusCondition

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | [**StringsName**](strings.Name.md) | Human and machine-readable name of this condition. For example \&quot;healthy\&quot;, \&quot;disk-full\&quot;.  | [default to null]
**LastChanged** | [**time.Time**](time.Time.md) | last time the value field was changed for this condition | [default to null]
**LastUpdated** | [**time.Time**](time.Time.md) | last time this condition was updated (value may not have changed) | [default to null]
**Value** | **string** | the state of the condition. may be boolean or some enumeration-derived value | [default to null]
**Reason** | **string** | a machine-readable value that systems use to reason about the state of the condition  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


