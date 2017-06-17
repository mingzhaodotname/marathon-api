# QueueProcessedOffersSummary

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProcessedOffersCount** | **float32** | The number of processed offers for this launch attempt. | [default to null]
**UnusedOffersCount** | **float32** | The number of unused offers for this launch attempt. | [default to null]
**LastUnusedOfferAt** | [**time.Time**](time.Time.md) | Point in time when the last unused offer has been processed. | [optional] [default to null]
**LastUsedOfferAt** | [**time.Time**](time.Time.md) | Point in time when the last used offer has been processed. | [optional] [default to null]
**RejectSummaryLastOffers** | [**[]QueueDeclinedOfferStep**](queue.DeclinedOfferStep.md) | Summarizes the reasons for not using offers for all last offers. This summary accumulates all processed offers by only taking the last offer per agent into account.  | [optional] [default to null]
**RejectSummaryLaunchAttempt** | [**[]QueueDeclinedOfferStep**](queue.DeclinedOfferStep.md) | Summarizes the reasons for not using offers for the launch attempt of this run specification. This summary accumulates all processed offers since the start of the launch attempt.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


