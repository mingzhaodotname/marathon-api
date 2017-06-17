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

// An environment variable set to a secret
type AppEnvEnvVarSecret struct {

	// The name of the secret to refer to. At runtime, the value of the secret will be injected into the value of the variable. 
	Secret string `json:"secret"`
}
