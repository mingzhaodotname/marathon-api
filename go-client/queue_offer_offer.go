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

type QueueOfferOffer struct {

	// The id of this offer
	Id string `json:"id"`

	// the hostname of the slave
	Hostname string `json:"hostname"`

	// the id of the slave
	AgentId string `json:"agentId"`

	// all offered ressources
	Resources []QueueOfferOfferResource `json:"resources"`

	// all attributes of the agent
	Attributes []QueueOfferAgentAttribute `json:"attributes"`
}