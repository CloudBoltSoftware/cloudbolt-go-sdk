package cbclient

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"strconv"
)

type NamingPolicyResult struct {
	CloudBoltResult
	Embedded struct {
		NamingPolicies []NamingPolicy `json:"namingPolicies"`
	} `json:"_embedded"`
}

type NamingPolicy struct {
	Links *struct {
		Self      CloudBoltHALItem `json:"self,omitempty"`
		Workspace CloudBoltHALItem `json:"workspace,omitempty"`
	} `json:"_links,omitempty"`
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type CustomName struct {
	Id        int
	Name      string
	DnsSuffix string
}

func (c *CloudBoltClient) GetNamingPolicy(name string) (*NamingPolicy, error) {
	apiurl := c.baseURL
	apiurl.Path = c.apiEndpoint("onefuse", "namingPolicies")
	apiurl.RawQuery = fmt.Sprintf(filterByName, url.QueryEscape(name))

	resp, err := c.makeRequest("GET", apiurl.String(), nil)
	if err != nil {
		return nil, err
	}

	// TODO: HTTP Response handling

	var res NamingPolicyResult
	json.NewDecoder(resp.Body).Decode(&res)

	// TODO: Sanity check the decoded object
	if len(res.Embedded.NamingPolicies) == 0 {
		return nil, fmt.Errorf(
			"Could not find Naming Policy with name %s. Does the user have permission to view this?",
			name,
		)
	}

	return &res.Embedded.NamingPolicies[0], nil
}

func (c *CloudBoltClient) GenerateCustomName(namingPolicyID string, workspaceID string, templateProperties map[string]interface{}) (*OneFuseJobStatus, error) {
	log.Println("onefuse.apiClient: GenerateCustomName")

	if workspaceID == "" {
		workspace, err := c.GetDefaultWorkSpace()

		if err != nil {
			return nil, err
		}

		workspaceID = strconv.Itoa(workspace.ID)
	}

	postBody := map[string]interface{}{
		"policy":             c.apiEndpoint("onefuse", "namingPolicies", namingPolicyID),
		"templateProperties": templateProperties,
		"workspace":          c.apiEndpoint("onefuse", "workspaces", workspaceID),
	}

	reqJSON, err := json.Marshal(postBody)
	if err != nil {
		return nil, err
	}

	apiurl := c.baseURL
	apiurl.Path = c.apiEndpoint("onefuse", "customNames")

	resp, err := c.makeRequest("POST", apiurl.String(), reqJSON)
	if err != nil {
		return nil, err
	}

	// Handle some common HTTP errors
	job_status, err := checkOneFuseResponse(resp)
	if err != nil {
		return nil, err
	}

	return job_status, nil
}

func (c *CloudBoltClient) GetCustomName(customNamePath string) (*CustomName, error) {
	apiurl := c.baseURL
	apiurl.Path = customNamePath

	resp, err := c.makeRequest("GET", apiurl.String(), nil)
	if err != nil {
		log.Fatalln(err)

		return nil, err
	}

	// We Decode the data because we already have an io.Reader on hand
	var customName CustomName
	json.NewDecoder(resp.Body).Decode(&customName)

	return &customName, nil
}

func (c *CloudBoltClient) GetCustomNameById(customNameId string) (*CustomName, error) {
	apiurl := c.baseURL
	apiurl.Path = c.apiEndpoint("onefuse", "customNames", customNameId)

	resp, err := c.makeRequest("GET", apiurl.String(), nil)
	if err != nil {
		log.Fatalln(err)

		return nil, err
	}

	// We Decode the data because we already have an io.Reader on hand
	var customName CustomName
	json.NewDecoder(resp.Body).Decode(&customName)

	return &customName, nil
}

func (c *CloudBoltClient) DeleteCustomName(customNameId string) (*OneFuseJobStatus, error) {
	apiurl := c.baseURL
	apiurl.Path = c.apiEndpoint("onefuse", "customNames", customNameId)

	resp, err := c.makeRequest("DELETE", apiurl.String(), nil)
	if err != nil {
		return nil, err
	}

	// Handle some common HTTP errors
	job_status, err := checkOneFuseResponse(resp)
	if err != nil {
		return nil, err
	}

	return job_status, nil
}
