package cbclient

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/url"
	"strconv"
)

type AnsibleTowerPolicyResult struct {
	CloudBoltResult
	Embedded struct {
		AnsibleTowerPolicies []AnsibleTowerPolicy `json:"ansibleTowerPolicies"`
	} `json:"_embedded"`
}

type AnsibleTowerPolicy struct {
	Links *struct {
		Self      CloudBoltHALItem `json:"self,omitempty"`
		Workspace CloudBoltHALItem `json:"workspace,omitempty"`
		Endpoint  CloudBoltHALItem `json:"endpoint,omitempty"`
	} `json:"_links,omitempty"`
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type AnsibleTowerDeployment struct {
	Links *struct {
		Self        CloudBoltHALItem `json:"self,omitempty"`
		Workspace   CloudBoltHALItem `json:"workspace,omitempty"`
		Policy      CloudBoltHALItem `json:"policy,omitempty"`
		JobMetadata CloudBoltHALItem `json:"jobMetadata,omitempty"`
	} `json:"_links,omitempty"`
	ID                     int      `json:"id,omitempty"`
	PolicyID               int      `json:"policyId,omitempty"`
	Policy                 string   `json:"policy,omitempty"`
	WorkspaceURL           string   `json:"workspace,omitempty"`
	Limit                  string   `json:"limit,omitempty"`
	InventoryName          string   `json:"inventoryName,omitempty"`
	Hosts                  []string `json:"hosts,omitempty"`
	Archived               bool     `json:"archived,omitempty"`
	ProvisioningJobResults []struct {
		Output          string `json:"output"`
		Status          string `json:"status"`
		JobTemplateName string `json:"jobTemplateName"`
	} `json:"provisioningJobResults,omitempty"`
	DeprovisioningJobResults *struct {
		Output          string `json:"output"`
		Status          string `json:"status"`
		JobTemplateName string `json:"jobTemplateName"`
	} `json:"deprovisioningJobResults,omitempty"`
	TemplateProperties map[string]interface{} `json:"templateProperties"`
}

func (c *CloudBoltClient) GetAnsibleTowerPolicy(name string) (*AnsibleTowerPolicy, error) {
	log.Println("onefuse.apiClient: GetAnsibleTowerPolicyByName")

	apiurl := c.baseURL
	apiurl.Path = c.apiEndpoint("onefuse", "ansibleTowerPolicies")
	apiurl.RawQuery = fmt.Sprintf(filterByName, url.QueryEscape(name))

	resp, err := c.makeRequest("GET", apiurl.String(), nil)
	if err != nil {
		return nil, err
	}

	// TODO: HTTP Response handling

	var res AnsibleTowerPolicyResult
	json.NewDecoder(resp.Body).Decode(&res)

	// TODO: Sanity check the decoded object
	if len(res.Embedded.AnsibleTowerPolicies) == 0 {
		return nil, fmt.Errorf(
			"Could not find Ansible Tower Policy with name %s. Does the user have permission to view this?",
			name,
		)
	}

	return &res.Embedded.AnsibleTowerPolicies[0], nil
}

func (c *CloudBoltClient) CreateAnsibleTowerDeployment(ansibleTowerDeployment *AnsibleTowerDeployment) (*OneFuseJobStatus, error) {
	log.Println("onefuse.apiClient: CreateAnsibleTowerDeployment")

	if ansibleTowerDeployment.WorkspaceURL == "" {
		workspace, err := c.GetDefaultWorkSpace()

		if err != nil {
			return nil, err
		}

		ansibleTowerDeployment.WorkspaceURL = workspace.Links.Self.Href
	}

	if ansibleTowerDeployment.Policy == "" {
		if ansibleTowerDeployment.PolicyID != 0 {
			ansibleTowerDeployment.Policy = c.apiEndpoint("onefuse", "ansibleTowerPolicies", strconv.Itoa(ansibleTowerDeployment.PolicyID))
		} else {
			return nil, errors.New("onefuse.apiClient: Ansible Tower Deployment Create requires a PolicyID or Policy URL")
		}
	} else {
		return nil, errors.New("onefuse.apiClient: Ansible Tower Deployment Create requires a PolicyID or Policy URL")
	}

	reqJSON, err := json.Marshal(ansibleTowerDeployment)
	if err != nil {
		return nil, err
	}

	apiurl := c.baseURL
	apiurl.Path = c.apiEndpoint("onefuse", "ansibleTowerDeployments")

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

func (c *CloudBoltClient) GetAnsibleTowerDeployment(ansibleDeploymentPath string) (*AnsibleTowerDeployment, error) {
	apiurl := c.baseURL
	apiurl.Path = ansibleDeploymentPath

	resp, err := c.makeRequest("GET", apiurl.String(), nil)
	if err != nil {
		log.Fatalln(err)

		return nil, err
	}

	// We Decode the data because we already have an io.Reader on hand
	var ansibleDeloyment AnsibleTowerDeployment
	json.NewDecoder(resp.Body).Decode(&ansibleDeloyment)

	return &ansibleDeloyment, nil
}

func (c *CloudBoltClient) GetAnsibleTowerDeploymentById(ansibleDeploymentId string) (*AnsibleTowerDeployment, error) {
	apiurl := c.baseURL
	apiurl.Path = c.apiEndpoint("onefuse", "ansibleTowerDeployments", ansibleDeploymentId)

	resp, err := c.makeRequest("GET", apiurl.String(), nil)
	if err != nil {
		log.Fatalln(err)

		return nil, err
	}

	// We Decode the data because we already have an io.Reader on hand
	var ansibleDeployment AnsibleTowerDeployment
	json.NewDecoder(resp.Body).Decode(&ansibleDeployment)

	return &ansibleDeployment, nil
}

func (c *CloudBoltClient) DeleteAnsibleTowerDeployment(ansibleDeploymentId string) (*OneFuseJobStatus, error) {
	apiurl := c.baseURL
	apiurl.Path = c.apiEndpoint("onefuse", "ansibleTowerDeployments", ansibleDeploymentId)

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
