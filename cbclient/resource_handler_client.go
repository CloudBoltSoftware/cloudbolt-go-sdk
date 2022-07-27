package cbclient

import (
	"encoding/json"
	"fmt"
	"net/url"
)

type CloudBoltResourceHandlerResult struct {
	CloudBoltResult
	Embedded struct {
		ResourceHandlers []CloudBoltReferenceFields `json:"resourceHandlers"`
	} `json:"_embedded"`
}

// GetResourceHandler accepts the name of a Resource Handler
//
func (c *CloudBoltClient) GetResourceHandler(name string) (*CloudBoltReferenceFields, error) {
	apiurl := c.baseURL
	apiurl.Path = c.apiEndpoint("resourceHandlers")
	apiurl.RawQuery = fmt.Sprintf(filterByName, url.QueryEscape(name))

	resp, err := c.makeRequest("GET", apiurl.String(), nil)
	if err != nil {
		return nil, err
	}
	// TODO: HTTP Response handling

	// We Decode the data because we already have an io.Reader on hand
	var res CloudBoltResourceHandlerResult
	json.NewDecoder(resp.Body).Decode(&res)

	// TODO: Sanity check the decoded object
	if len(res.Embedded.ResourceHandlers) == 0 {
		return nil, fmt.Errorf(
			"Could not find resource handler with name %s. Does the user have permission to view this?",
			name,
		)
	}
	return &res.Embedded.ResourceHandlers[0], nil
}

func (c *CloudBoltClient) GetResourceHandlerById(id string) (*CloudBoltReferenceFields, error) {
	apiurl := c.baseURL
	apiurl.Path = c.apiEndpoint(
		"resourceHandlers",
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
