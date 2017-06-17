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

type MetricsHistogram struct {

	// The number of samples
	Count int64 `json:"count"`

	// The minimum value
	Min float32 `json:"min"`

	// The maximum value
	Max float32 `json:"max"`

	// The 50th percentile median value
	P50 float32 `json:"p50"`

	// The 75th percentile median value
	P75 float32 `json:"p75"`

	// The 98th percentile median value
	P98 float32 `json:"p98"`

	// The 99th percentile median value
	P99 float32 `json:"p99"`

	// The 999th percentile median value
	P999 float32 `json:"p999"`

	// The average of all samples
	Mean float32 `json:"mean"`

	// metadata about the metric
	Tags map[string]string `json:"tags"`

	// The unit of measurement
	Unit MetricsUnitOfMeasurement `json:"unit"`
}
