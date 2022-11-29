package cbclient

import (
	"encoding/json"
	"fmt"
	"net/url"
)

type EndpointsListResult struct {
	CloudBoltResult
	Embedded struct {
		Endpoints []MicrosoftEndpoint `json:"endpoints"` // TODO: Generalize to Endpoints
	} `json:"_embedded"`
}

type MicrosoftEndpoint struct {
	Links *struct {
		Self       CloudBoltHALItem `json:"self,omitempty"`
		Workspace  CloudBoltHALItem `json:"workspace,omitempty"`
		Credential CloudBoltHALItem `json:"credential,omitempty"`
	} `json:"_links,omitempty"`
	ID               int    `json:"id,omitempty"`
	Type             string `json:"type,omitempty"`
	Name             string `json:"name,omitempty"`
	Description      string `json:"description,omitempty"`
	Host             string `json:"host,omitempty"`
	Port             int    `json:"port,omitempty"`
	SSL              bool   `json:"ssl,omitempty"`
	MicrosoftVersion string `json:"microsoftVersion,omitempty"`
}

func (c *CloudBoltClient) GetMicrosoftEndpoint(name string) (*MicrosoftEndpoint, error) {
	apiurl := c.baseURL
	apiurl.Path = c.apiEndpoint("onefuse", "endpoints")
	apiurl.RawQuery = fmt.Sprintf(filterByName, url.QueryEscape(name))

	resp, err := c.makeRequest("GET", apiurl.String(), nil)
	if err != nil {
		return nil, err
	}

	// TODO: HTTP Response handling

	var res EndpointsListResult
	json.NewDecoder(resp.Body).Decode(&res)

	// TODO: Sanity check the decoded object
	if len(res.Embedded.Endpoints) == 0 {
		return nil, fmt.Errorf(
			"Could not find Microsoft Endpoint with name %s. Does the user have permission to view this?",
			name,
		)
	}

	return &res.Embedded.Endpoints[0], nil
}
