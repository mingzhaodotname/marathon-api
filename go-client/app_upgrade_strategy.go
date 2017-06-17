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

// During an upgrade all instances of an application get replaced by a new version. The upgradeStrategy controls how Marathon stops old versions and launches new versions. 
type AppUpgradeStrategy struct {

	// A number between 0 and 1 which is multiplied with the instance count. This is the maximum number of additional instances launched at any point of time during the upgrade process. 
	MaximumOverCapacity float64 `json:"maximumOverCapacity"`

	// A number between 0 and 1 that is multiplied with the instance count. This is the minimum number of healthy nodes that do not sacrifice overall application purpose. Marathon will make sure, during the upgrade process, that at any point of time this number of healthy instances are up. 
	MinimumHealthCapacity float64 `json:"minimumHealthCapacity"`
}
