package cbclient

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
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

type CloudBoltResourceResult struct {
	CloudBoltResult
	Embedded struct {
		Resources []CloudBoltResource `json:"resources"`
	} `json:"_embedded"`
}


func (c *CloudBoltClient) GetResourceById(id string) (*CloudBoltResource, error) {
	apiurl := c.baseURL
	apiurl.Path = c.apiEndpoint("cmp", "resources", id)

	resp, err := c.makeRequest("GET", apiurl.String(), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, ErrNotFound
	}

	var res CloudBoltResource
	json.NewDecoder(resp.Body).Decode(&res)

	return &res, nil
}

func (c *CloudBoltClient) GetResourceByName(name string) (*CloudBoltResource, error) {
	apiurl := c.baseURL
	apiurl.Path = c.apiEndpoint("cmp", "resources")
	apiurl.RawQuery = fmt.Sprintf(filterByName+";status:ACTIVE", url.QueryEscape(name))

	resp, err := c.makeRequest("GET", apiurl.String(), nil)
	if err != nil {
		return nil, err
	}

	// We Decode the data because we already have an io.Reader on hand
	var res CloudBoltResourceResult
	json.NewDecoder(resp.Body).Decode(&res)

	// TODO: Sanity check the decoded object
	if len(res.Embedded.Resources) == 0 {
		return nil, fmt.Errorf(
			"Could not find resorce with name %s. Does the user have permission to view this?",
			name,
		)
	}

	if len(res.Embedded.Resources) > 1 {
		return nil, fmt.Errorf(
			"More than one resource with name %s found.",
			name,
		)
	}

	return &res.Embedded.Resources[0], nil
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

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, ErrNotFound
	}

	// We Decode the data because we already have an io.Reader on hand
	var res CloudBoltResource
	json.NewDecoder(resp.Body).Decode(&res)

	return &res, nil
}
