package cbclient

import (
	"encoding/json"
	"log"
)

// CloudBoltResource contains metadata about Resources (e.g., "Services") in CloudBolt
type CloudBoltResource struct {
	Links struct {
		Self           CloudBoltHALItem   `json:"self"`
		ResourceType   CloudBoltHALItem   `json:"resourceType"`
		Blueprint      CloudBoltHALItem   `json:"blueprint"`
		Owner          CloudBoltHALItem   `json:"owner"`
		Group          CloudBoltHALItem   `json:"group"`
		Jobs           []CloudBoltHALItem `json:"jobs"`
		ParentResource CloudBoltHALItem   `json:"parentResource"`
		Servers        []CloudBoltHALItem `json:"servers"`
		Actions        []CloudBoltHALItem `json:"actions"`
	} `json:"_links"`
	Name       string                   `json:"name"`
	ID         string                   `json:"id"`
	Created    string                   `json:"created"`
	Status     string                   `json:"status"`
	Attributes []map[string]interface{} `json:"attributes"`
}

// GetResource fetches a Resource object from CloudBolt at the given path
// - Resource Path (resourcePath) e.g., "/api/v2/resources/service/123/"
func (c *CloudBoltClient) GetResource(resourcePath string) (*CloudBoltResource, error) {
	apiurl := c.baseURL
	apiurl.Path = resourcePath

	// log.Printf("[!!] apiurl in GetResource: %+v (%+v)", apiurl.String(), apiurl)

	resp, err := c.makeRequest("GET", apiurl.String(), nil)
	if err != nil {
		log.Fatalln(err)

		return nil, err
	}

	// We Decode the data because we already have an io.Reader on hand
	var res CloudBoltResource
	json.NewDecoder(resp.Body).Decode(&res)

	return &res, nil
}
