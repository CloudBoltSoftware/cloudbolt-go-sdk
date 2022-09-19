package cbclient

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/url"
	"strconv"
)

type VraPolicyResult struct {
	CloudBoltResult
	Embedded struct {
		VraPolicies []VraPolicy `json:"vraPolicies"`
	} `json:"_embedded"`
}

type VraPolicy struct {
	Links *struct {
		Self      CloudBoltHALItem `json:"self,omitempty"`
		Workspace CloudBoltHALItem `json:"workspace,omitempty"`
		Endpoint  CloudBoltHALItem `json:"endpoint,omitempty"`
	} `json:"_links,omitempty"`
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type VraDeployment struct {
	Links *struct {
		Self        CloudBoltHALItem `json:"self,omitempty"`
		Workspace   CloudBoltHALItem `json:"workspace,omitempty"`
		Policy      CloudBoltHALItem `json:"policy,omitempty"`
		JobMetadata CloudBoltHALItem `json:"jobMetadata,omitempty"`
	} `json:"_links,omitempty"`
	ID                 int                    `json:"id,omitempty"`
	PolicyID           int                    `json:"policyId,omitempty"`
	Policy             string                 `json:"policy,omitempty"`
	WorkspaceURL       string                 `json:"workspace,omitempty"`
	DeploymentName     string                 `json:"deploymentName,omitempty"`
	Name               string                 `json:"name,omitempty"`
	Archived           bool                   `json:"archived,omitempty"`
	TemplateProperties map[string]interface{} `json:"templateProperties"`
	DeploymentInfo     map[string]interface{} `json:"deploymentInfo,omitempty"`
	BlueprintName      string                 `json:"blueprintName,omitempty"`
	ProjectName        string                 `json:"projectName,omitempty"`
}

func (c *CloudBoltClient) GetVraPolicy(name string) (*VraPolicy, error) {
	apiurl := c.baseURL
	apiurl.Path = c.apiEndpoint("onefuse", "vraPolicies")
	apiurl.RawQuery = fmt.Sprintf(filterByName, url.QueryEscape(name))

	resp, err := c.makeRequest("GET", apiurl.String(), nil)
	if err != nil {
		return nil, err
	}

	// TODO: HTTP Response handling

	var res VraPolicyResult
	json.NewDecoder(resp.Body).Decode(&res)

	// TODO: Sanity check the decoded object
	if len(res.Embedded.VraPolicies) == 0 {
		return nil, fmt.Errorf(
			"Could not find vRA Policy with name %s. Does the user have permission to view this?",
			name,
		)
	}

	return &res.Embedded.VraPolicies[0], nil
}

func (c *CloudBoltClient) CreateVraDeployment(vraDeployment *VraDeployment) (*OneFuseJobStatus, error) {
	log.Println("onefuse.apiClient: CreateVraDeployment")

	if vraDeployment.WorkspaceURL == "" {
		workspace, err := c.GetDefaultWorkSpace()

		if err != nil {
			return nil, err
		}

		vraDeployment.WorkspaceURL = workspace.Links.Self.Href
	}

	if vraDeployment.Policy == "" {
		if vraDeployment.PolicyID != 0 {
			vraDeployment.Policy = c.apiEndpoint("onefuse", "modulePolicies", strconv.Itoa(vraDeployment.PolicyID))
		} else {
			return nil, errors.New("onefuse.apiClient: vRA Deployment Create requires a PolicyID or Policy URL")
		}
	} else {
		return nil, errors.New("onefuse.apiClient: vRA Deployment Create requires a PolicyID or Policy URL")
	}

	reqJSON, err := json.Marshal(vraDeployment)
	if err != nil {
		return nil, err
	}

	apiurl := c.baseURL
	apiurl.Path = c.apiEndpoint("onefuse", "vraDeployments")

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

func (c *CloudBoltClient) GetVraDeployment(vraDeploymentPath string) (*VraDeployment, error) {
	apiurl := c.baseURL
	apiurl.Path = vraDeploymentPath

	resp, err := c.makeRequest("GET", apiurl.String(), nil)
	if err != nil {
		log.Fatalln(err)

		return nil, err
	}

	// We Decode the data because we already have an io.Reader on hand
	var vraDeployment VraDeployment
	json.NewDecoder(resp.Body).Decode(&vraDeployment)

	return &vraDeployment, nil
}

func (c *CloudBoltClient) GetVraDeploymentById(vraDeploymentId string) (*VraDeployment, error) {
	apiurl := c.baseURL
	apiurl.Path = c.apiEndpoint("onefuse", "vraDeployments", vraDeploymentId)

	resp, err := c.makeRequest("GET", apiurl.String(), nil)
	if err != nil {
		log.Fatalln(err)

		return nil, err
	}

	// We Decode the data because we already have an io.Reader on hand
	var vraDeployment VraDeployment
	json.NewDecoder(resp.Body).Decode(&vraDeployment)

	return &vraDeployment, nil
}

func (c *CloudBoltClient) DeleteVraDeployment(vraDeploymentId string) (*OneFuseJobStatus, error) {
	apiurl := c.baseURL
	apiurl.Path = c.apiEndpoint("onefuse", "vraDeployments", vraDeploymentId)

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
