# AppAppContainerAppCContainer

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Image** | **string** | The name of the AppC image to use | [default to null]
**Id** | **string** | An image ID is a string of the format &#39;hash-value&#39;, where &#39;hash&#39; is the hash algorithm used and &#39;value&#39; is the hex-encoded digest. Currently the only permitted hash algorithm is sha512.  | [optional] [default to null]
**Labels** | [**AppLabelKvLabels**](app.label.KVLabels.md) | Optional labels. Suggested labels: &#39;version&#39;, &#39;os&#39;, &#39;arch&#39;.  | [optional] [default to null]
**ForcePullImage** | **bool** | The container will be pulled, regardless if it is already available on the local system  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


