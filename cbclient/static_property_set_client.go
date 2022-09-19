package cbclient

import (
	"encoding/json"
	"fmt"
	"net/url"
)

type StaticPropertySetResult struct {
	CloudBoltResult
	Embedded struct {
		PropertySets []StaticPropertySet `json:"propertySets"`
	} `json:"_embedded"`
}

type StaticPropertySet struct {
	Links *struct {
		Self      CloudBoltHALItem `json:"self,omitempty"`
		Workspace CloudBoltHALItem `json:"workspace,omitempty"`
	} `json:"_links,omitempty"`
	ID          int                    `json:"id,omitempty"`
	Name        string                 `json:"name,omitempty"`
	Description string                 `json:"description,omitempty"`
	Properties  map[string]interface{} `json:"properties,omitempty"`
	Raw         string
}

func (c *CloudBoltClient) GetStaticPropertySet(name string) (*StaticPropertySet, error) {
	apiurl := c.baseURL
	apiurl.Path = c.apiEndpoint("onefuse", "propertySets")
	apiurl.RawQuery = fmt.Sprintf(filterByName, url.QueryEscape(name))

	resp, err := c.makeRequest("GET", apiurl.String(), nil)
	if err != nil {
		return nil, err
	}

	// TODO: HTTP Response handling

	var res StaticPropertySetResult
	json.NewDecoder(resp.Body).Decode(&res)

	// TODO: Sanity check the decoded object
	if len(res.Embedded.PropertySets) == 0 {
		return nil, fmt.Errorf(
			"Could not find Static Property Set with name %s. Does the user have permission to view this?",
			name,
		)
	}

	return &res.Embedded.PropertySets[0], nil
}
