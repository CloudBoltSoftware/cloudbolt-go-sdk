package cbclient

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
)

type WorkspaceResult struct {
	CloudBoltResult
	Embedded struct {
		Workspaces []Workspace `json:"workspaces"`
	} `json:"_embedded"`
}

type Workspace struct {
	Links *struct {
		Self CloudBoltHALItem `json:"self,omitempty"`
	} `json:"_links,omitempty"`
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

func (c *CloudBoltClient) GetDefaultWorkSpace() (*Workspace, error) {
	log.Println("onefuse.apiClient: GetDefaultWorkSpace")

	return c.GetWorkSpace("Default")
}

func (c *CloudBoltClient) GetWorkSpace(name string) (*Workspace, error) {
	log.Println("onefuse.apiClient: GetWorkSpace")

	apiurl := c.baseURL
	apiurl.Path = c.apiEndpoint("onefuse", "workspaces")
	apiurl.RawQuery = fmt.Sprintf(filterByName, url.QueryEscape(name))

	resp, err := c.makeRequest("GET", apiurl.String(), nil)
	if err != nil {
		return nil, err
	}

	// TODO: HTTP Response handling

	var res WorkspaceResult
	json.NewDecoder(resp.Body).Decode(&res)

	// TODO: Sanity check the decoded object
	if len(res.Embedded.Workspaces) == 0 {
		return nil, fmt.Errorf(
			"Could not find Workspace with name %s. Does the user have permission to view this?",
			name,
		)
	}

	return &res.Embedded.Workspaces[0], nil
}
