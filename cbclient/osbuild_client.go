package cbclient

import (
	"encoding/json"
	"fmt"
	"net/url"
)

type CloudBoltOSBuildResult struct {
	CloudBoltResult
	Embedded struct {
		OSBuilds []CloudBoltReferenceFields `json:"osBuilds"`
	} `json:"_embedded"`
}

// GetOSBuild accepts the name of a OSBuild
//
func (c *CloudBoltClient) GetOSBuild(name string) (*CloudBoltReferenceFields, error) {
	apiurl := c.baseURL
	apiurl.Path = c.apiEndpoint("osBuilds")
	apiurl.RawQuery = fmt.Sprintf(filterByName, url.QueryEscape(name))

	resp, err := c.makeRequest("GET", apiurl.String(), nil)
	if err != nil {
		return nil, err
	}
	// TODO: HTTP Response handling

	// We Decode the data because we already have an io.Reader on hand
	var res CloudBoltOSBuildResult
	json.NewDecoder(resp.Body).Decode(&res)

	// TODO: Sanity check the decoded object
	if len(res.Embedded.OSBuilds) == 0 {
		return nil, fmt.Errorf(
			"Could not find enviornment with name %s. Does the user have permission to view this?",
			name,
		)
	}
	return &res.Embedded.OSBuilds[0], nil
}

func (c *CloudBoltClient) GetOSBuildById(id string) (*CloudBoltReferenceFields, error) {
	apiurl := c.baseURL
	apiurl.Path = c.apiEndpoint(
		"osBuilds",
		id,
	)

	resp, err := c.makeRequest("GET", apiurl.String(), nil)
	if err != nil {
		return nil, err
	}
	// TODO: HTTP Response handling

	// We Decode the data because we already have an io.Reader on hand
	var res CloudBoltReferenceFields
	json.NewDecoder(resp.Body).Decode(&res)

	return &res, nil
}
