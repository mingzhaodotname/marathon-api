# QueueQueuePod

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **float32** | The number of instances left to launch. | [default to null]
**Delay** | [**QueueQueueDelay**](queue.QueueDelay.md) | If a runspec has failed to often the launch will be delayed. See backoff to tune this behavior. | [optional] [default to null]
**Since** | [**time.Time**](time.Time.md) | point in time since Marathon has started to launch tasks. | [default to null]
**ProcessedOffersSummary** | [**QueueProcessedOffersSummary**](queue.ProcessedOffersSummary.md) | Statistics for processed offers. | [default to null]
**LastUnusedOffers** | [**[]QueueUnusedOffer**](queue.UnusedOffer.md) | Last N unused offers, where N can be configured via xxx. | [optional] [default to null]
**Pod** | [**PodPod**](pod.Pod.md) |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


