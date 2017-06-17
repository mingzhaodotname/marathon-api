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

// Describes the information about (pseudo) TTY that can be attached to a process running in a container. 
type AppContainerCommonsTty struct {

	// The rows of the tty device.
	Rows int32 `json:"rows,omitempty"`

	// The columns of the tty device.
	Columns int32 `json:"columns,omitempty"`
}
