/* 
 * Marathon REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * OpenAPI spec version: 2.0
 * 
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package swagger

// Configures exponential backoff behavior when launching potentially sick apps. This prevents sandboxes associated with consecutively failing tasks from filling up the hard disk on Mesos slaves. The backoff period is multiplied by the factor for each consecutive failure until it reaches maxLaunchDelaySeconds. This applies also to tasks that are killed due to failing too many health checks. 
type PodPodSchedulingBackoffStrategy struct {

	// The initial backoff (seconds) applied when a launched instance fails.
	Backoff float32 `json:"backoff,omitempty"`

	// The factor applied to the current backoff to determine the new backoff.
	BackoffFactor float32 `json:"backoffFactor,omitempty"`

	// The maximum backoff (seconds) applied when subsequent failures are detected.
	MaxLaunchDelay float32 `json:"maxLaunchDelay,omitempty"`
}
