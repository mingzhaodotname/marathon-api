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

import (
	"time"
)

// Detailed version information
type AppVersionInfoVersionInfo struct {

	// Contains the timestamp of the last change to this pod which was not simply a scaling or restarting configuration
	LastScalingAt time.Time `json:"lastScalingAt"`

	// Contains the timestamp of the last change including changes like scaling or restarting.
	LastConfigChangeAt time.Time `json:"lastConfigChangeAt"`
}
