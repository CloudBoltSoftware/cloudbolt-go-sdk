package cbclient

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/url"
	"strconv"
)

type ModulePolicyResult struct {
	CloudBoltResult
	Embedded struct {
		ModulePolicies []ModulePolicy `json:"modulePolicies"`
	} `json:"_embedded"`
}

type ModulePolicy struct {
	Links *struct {
		Self      CloudBoltHALItem `json:"self,omitempty"`
		Workspace CloudBoltHALItem `json:"workspace,omitempty"`
		Blueprint CloudBoltHALItem `json:"blueprint,omitempty"`
	} `json:"_links,omitempty"`
	ID             int    `json:"id,omitempty"`
	Name           string `json:"name,omitempty"`
	Description    string `json:"description,omitempty"`
	PolicyTemplate string `json:"policyTemplate,omitempty"`
}

type ModuleDeployment struct {
	Links *struct {
		Self        CloudBoltHALItem `json:"self,omitempty"`
		Workspace   CloudBoltHALItem `json:"workspace,omitempty"`
		Policy      CloudBoltHALItem `json:"policy,omitempty"`
		JobMetadata CloudBoltHALItem `json:"jobMetadata,omitempty"`
	} `json:"_links,omitempty"`
	ID                       int                      `json:"id,omitempty"`
	PolicyID                 int                      `json:"policyId,omitempty"`
	Policy                   string                   `json:"policy,omitempty"`
	WorkspaceURL             string                   `json:"workspace,omitempty"`
	Name                     string                   `json:"name,omitempty"`
	Archived                 bool                     `json:"archived,omitempty"`
	TemplateProperties       map[string]interface{}   `json:"templateProperties"`
	ProvisioningJobResults   []map[string]interface{} `json:"provisioningJobResults,omitempty"`
	DeprovisioningJobResults []map[string]interface{} `json:"deprovisioningJobResults,omitempty"`
}

func (c *CloudBoltClient) GetModulePolicy(name string) (*ModulePolicy, error) {
	apiurl := c.baseURL
	apiurl.Path = c.apiEndpoint("onefuse", "modulePolicies")
	apiurl.RawQuery = fmt.Sprintf(filterByName, url.QueryEscape(name))

	resp, err := c.makeRequest("GET", apiurl.String(), nil)
	if err != nil {
		return nil, err
	}

	// TODO: HTTP Response handling

	var res ModulePolicyResult
	json.NewDecoder(resp.Body).Decode(&res)

	// TODO: Sanity check the decoded object
	if len(res.Embedded.ModulePolicies) == 0 {
		return nil, fmt.Errorf(
			"Could not find Module Policy with name %s. Does the user have permission to view this?",
			name,
		)
	}

	return &res.Embedded.ModulePolicies[0], nil
}

func (c *CloudBoltClient) CreateModuleDeployment(moduleDeployment *ModuleDeployment) (*OneFuseJobStatus, error) {
	log.Println("onefuse.apiClient: CreateModuleDeployment")

	if moduleDeployment.WorkspaceURL == "" {
		workspace, err := c.GetDefaultWorkSpace()

		if err != nil {
			return nil, err
		}

		moduleDeployment.WorkspaceURL = workspace.Links.Self.Href
	}

	log.Println("onefuse.apiClient: CreateModuleDeployment")
	if moduleDeployment.Policy == "" {
		if moduleDeployment.PolicyID != 0 {
			moduleDeployment.Policy = c.apiEndpoint("onefuse", "modulePolicies", strconv.Itoa(moduleDeployment.PolicyID))
		} else {
			return nil, errors.New("onefuse.apiClient: Module Deployment Create requires a PolicyID or Policy URL")
		}
	} else {
		return nil, errors.New("onefuse.apiClient: Module Deployment Create requires a PolicyID or Policy URL")
	}

	reqJSON, err := json.Marshal(moduleDeployment)
	if err != nil {
		return nil, err
	}

	apiurl := c.baseURL
	apiurl.Path = c.apiEndpoint("onefuse", "moduleManagedObjects")

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

func (c *CloudBoltClient) GetModuleDeployment(moduleDeploymentPath string) (*ModuleDeployment, error) {
	apiurl := c.baseURL
	apiurl.Path = moduleDeploymentPath

	resp, err := c.makeRequest("GET", apiurl.String(), nil)
	if err != nil {
		log.Fatalln(err)

		return nil, err
	}

	// We Decode the data because we already have an io.Reader on hand
	var moduleDeployment ModuleDeployment
	json.NewDecoder(resp.Body).Decode(&moduleDeployment)

	return &moduleDeployment, nil
}

func (c *CloudBoltClient) GetModuleDeploymentById(moduleDeploymentId string) (*ModuleDeployment, error) {
	apiurl := c.baseURL
	apiurl.Path = c.apiEndpoint("onefuse", "moduleManagedObjects", moduleDeploymentId)

	resp, err := c.makeRequest("GET", apiurl.String(), nil)
	if err != nil {
		log.Fatalln(err)

		return nil, err
	}

	// We Decode the data because we already have an io.Reader on hand
	var moduleDeployment ModuleDeployment
	json.NewDecoder(resp.Body).Decode(&moduleDeployment)

	return &moduleDeployment, nil
}

func (c *CloudBoltClient) DeleteModuleDeployment(moduleDeploymentId string) (*OneFuseJobStatus, error) {
	apiurl := c.baseURL
	apiurl.Path = c.apiEndpoint("onefuse", "moduleManagedObjects", moduleDeploymentId)

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
