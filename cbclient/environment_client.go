package cbclient

import (
	"encoding/json"
	"fmt"
	"net/url"
)

type CloudBoltEnvironmentResult struct {
	CloudBoltResult
	Embedded struct {
		Environments []CloudBoltReferenceFields `json:"environments"`
	} `json:"_embedded"`
}

// GetEnvironment accepts the name of a Environment
//
func (c *CloudBoltClient) GetEnvironment(name string) (*CloudBoltReferenceFields, error) {
	apiurl := c.baseURL
	apiurl.Path = c.apiEndpoint("cmp", "environments")
	apiurl.RawQuery = fmt.Sprintf(filterByName, url.QueryEscape(name))

	resp, err := c.makeRequest("GET", apiurl.String(), nil)
	if err != nil {
		return nil, err
	}

	// TODO: HTTP Response handling

	var res CloudBoltEnvironmentResult
	json.NewDecoder(resp.Body).Decode(&res)

	// TODO: Sanity check the decoded object
	if len(res.Embedded.Environments) == 0 {
		return nil, fmt.Errorf(
			"Could not find enviornment with name %s. Does the user have permission to view this?",
			name,
		)
	}
	return &res.Embedded.Environments[0], nil
}

func (c *CloudBoltClient) GetEnvironmentById(id string) (*CloudBoltReferenceFields, error) {
	apiurl := c.baseURL
	apiurl.Path = c.apiEndpoint("cmp", "environments", id)

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
