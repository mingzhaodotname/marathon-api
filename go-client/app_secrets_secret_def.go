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

// A secret declaration
type AppSecretsSecretDef struct {

	// The source of the secrets value. The format dependes on the secret store 
	Source string `json:"source"`
}