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

type AppHealthCommandCheck struct {

	// Command line executed by the default shell. This process has to return with a zero exit code to indicate the process is healthy. Return codes other than null signal, the task is unhealthy. 
	Value string `json:"value"`
}
