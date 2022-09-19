package cbclient

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/url"
	"strconv"
)

type ScriptingPolicyResult struct {
	CloudBoltResult
	Embedded struct {
		ScriptingPolicies []ScriptingPolicy `json:"scriptingPolicies"`
	} `json:"_embedded"`
}

type ScriptingPolicy struct {
	Links *struct {
		Self      CloudBoltHALItem `json:"self,omitempty"`
		Workspace CloudBoltHALItem `json:"workspace,omitempty"`
	} `json:"_links,omitempty"`
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type ScriptingDeployment struct {
	Links *struct {
		Self        CloudBoltHALItem `json:"self,omitempty"`
		Workspace   CloudBoltHALItem `json:"workspace,omitempty"`
		Policy      CloudBoltHALItem `json:"policy,omitempty"`
		JobMetadata CloudBoltHALItem `json:"jobMetadata,omitempty"`
	} `json:"_links,omitempty"`
	ID                  int    `json:"id,omitempty"`
	PolicyID            int    `json:"policyId,omitempty"`
	Policy              string `json:"policy,omitempty"`
	WorkspaceURL        string `json:"workspace,omitempty"`
	Hostname            string `json:"hostname,omitempty"`
	ProvisioningDetails *struct {
		Status string   `json:"status"`
		Output []string `json:"output"`
	} `json:"provisioningDetails,omitempty"`
	DeprovisioningDetails *struct {
		Status string   `json:"status"`
		Output []string `json:"output"`
	} `json:"deprovisioningDetails,omitempty"`
	Archived           bool                   `json:"archived,omitempty"`
	TemplateProperties map[string]interface{} `json:"templateProperties"`
}

func (c *CloudBoltClient) GetScriptingPolicy(name string) (*ScriptingPolicy, error) {
	apiurl := c.baseURL
	apiurl.Path = c.apiEndpoint("onefuse", "scriptingPolicies")
	apiurl.RawQuery = fmt.Sprintf(filterByName, url.QueryEscape(name))

	resp, err := c.makeRequest("GET", apiurl.String(), nil)
	if err != nil {
		return nil, err
	}

	// TODO: HTTP Response handling

	var res ScriptingPolicyResult
	json.NewDecoder(resp.Body).Decode(&res)

	// TODO: Sanity check the decoded object
	if len(res.Embedded.ScriptingPolicies) == 0 {
		return nil, fmt.Errorf(
			"Could not find Naming Policy with name %s. Does the user have permission to view this?",
			name,
		)
	}

	return &res.Embedded.ScriptingPolicies[0], nil
}

func (c *CloudBoltClient) CreateScriptingDeployment(scriptionDeployment *ScriptingDeployment) (*OneFuseJobStatus, error) {
	log.Println("onefuse.apiClient: CreateScriptingDeployment")

	if scriptionDeployment.WorkspaceURL == "" {
		workspace, err := c.GetDefaultWorkSpace()

		if err != nil {
			return nil, err
		}

		scriptionDeployment.WorkspaceURL = workspace.Links.Self.Href
	}

	if scriptionDeployment.Policy == "" {
		if scriptionDeployment.PolicyID != 0 {
			scriptionDeployment.Policy = c.apiEndpoint("onefuse", "scriptingPolicies", strconv.Itoa(scriptionDeployment.PolicyID))
		} else {
			return nil, errors.New("onefuse.apiClient: Scripting Deployment Create requires a PolicyID or Policy URL")
		}
	} else {
		return nil, errors.New("onefuse.apiClient: Scripting Deployment Create requires a PolicyID or Policy URL")
	}

	reqJSON, err := json.Marshal(scriptionDeployment)
	if err != nil {
		return nil, err
	}

	apiurl := c.baseURL
	apiurl.Path = c.apiEndpoint("onefuse", "scriptingDeployments")

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

func (c *CloudBoltClient) GetScriptingDeployment(scriptingDeploymentPath string) (*ScriptingDeployment, error) {
	apiurl := c.baseURL
	apiurl.Path = scriptingDeploymentPath

	resp, err := c.makeRequest("GET", apiurl.String(), nil)
	if err != nil {
		log.Fatalln(err)

		return nil, err
	}

	// We Decode the data because we already have an io.Reader on hand
	var scriptingDeployment ScriptingDeployment
	json.NewDecoder(resp.Body).Decode(&scriptingDeployment)

	return &scriptingDeployment, nil
}

func (c *CloudBoltClient) GetScriptingDeploymentById(scriptingDeploymentId string) (*ScriptingDeployment, error) {
	apiurl := c.baseURL
	apiurl.Path = c.apiEndpoint("onefuse", "scriptingDeployments", scriptingDeploymentId)

	resp, err := c.makeRequest("GET", apiurl.String(), nil)
	if err != nil {
		log.Fatalln(err)

		return nil, err
	}

	// We Decode the data because we already have an io.Reader on hand
	var scriptingDeployment ScriptingDeployment
	json.NewDecoder(resp.Body).Decode(&scriptingDeployment)

	return &scriptingDeployment, nil
}

func (c *CloudBoltClient) DeleteScriptingDeployment(scriptingDeploymentId string) (*OneFuseJobStatus, error) {
	apiurl := c.baseURL
	apiurl.Path = c.apiEndpoint("onefuse", "scriptingDeployments", scriptingDeploymentId)

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
