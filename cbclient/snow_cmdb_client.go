package cbclient

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/url"
	"strconv"
)

type ServiceNowCMDBPolicyResult struct {
	CloudBoltResult
	Embedded struct {
		ServiceNowCMDBPolicies []ServiceNowCMDBPolicy `json:"servicenowCMDBPolicies"`
	} `json:"_embedded"`
}

type ServiceNowCMDBPolicy struct {
	Links *struct {
		Self      CloudBoltHALItem `json:"self,omitempty"`
		Workspace CloudBoltHALItem `json:"workspace,omitempty"`
		Endpoint  CloudBoltHALItem `json:"endpoint,omitempty"`
	} `json:"_links,omitempty"`
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type ServicenowCMDBDeployment struct {
	Links *struct {
		Self        CloudBoltHALItem `json:"self,omitempty"`
		Workspace   CloudBoltHALItem `json:"workspace,omitempty"`
		Policy      CloudBoltHALItem `json:"policy,omitempty"`
		JobMetadata CloudBoltHALItem `json:"jobMetadata,omitempty"`
	} `json:"_links,omitempty"`
	ID                     int                      `json:"id,omitempty"`
	PolicyID               int                      `json:"policyId,omitempty"`
	Policy                 string                   `json:"policy,omitempty"`
	WorkspaceURL           string                   `json:"workspace,omitempty"`
	ConfigurationItemsInfo []map[string]interface{} `json:"configurationItemsInfo,omitempty"`
	ExecutionDetails       map[string]interface{}   `json:"executionDetails,omitempty"`
	Archived               bool                     `json:"archived,omitempty"`
	TemplateProperties     map[string]interface{}   `json:"templateProperties"`
}

func (c *CloudBoltClient) GetServiceNowCMDBPolicy(name string) (*ServiceNowCMDBPolicy, error) {
	apiurl := c.baseURL
	apiurl.Path = c.apiEndpoint("onefuse", "servicenowCMDBPolicies")
	apiurl.RawQuery = fmt.Sprintf(filterByName, url.QueryEscape(name))

	resp, err := c.makeRequest("GET", apiurl.String(), nil)
	if err != nil {
		return nil, err
	}

	// TODO: HTTP Response handling

	var res ServiceNowCMDBPolicyResult
	json.NewDecoder(resp.Body).Decode(&res)

	// TODO: Sanity check the decoded object
	if len(res.Embedded.ServiceNowCMDBPolicies) == 0 {
		return nil, fmt.Errorf(
			"Could not find ServiceNow CMDB Policy with name %s. Does the user have permission to view this?",
			name,
		)
	}

	return &res.Embedded.ServiceNowCMDBPolicies[0], nil
}

func (c *CloudBoltClient) CreateServicenowCMDBDeployment(snowDeployment *ServicenowCMDBDeployment) (*OneFuseJobStatus, error) {
	log.Println("onefuse.apiClient: CreateServicenowCMDBDeployment")

	if snowDeployment.WorkspaceURL == "" {
		workspace, err := c.GetDefaultWorkSpace()

		if err != nil {
			return nil, err
		}

		snowDeployment.WorkspaceURL = workspace.Links.Self.Href
	}

	if snowDeployment.Policy == "" {
		if snowDeployment.PolicyID != 0 {
			snowDeployment.Policy = c.apiEndpoint("onefuse", "modulePolicies", strconv.Itoa(snowDeployment.PolicyID))
		} else {
			return nil, errors.New("onefuse.apiClient: ServiceNow CMDB Deployment Create requires a PolicyID or Policy URL")
		}
	} else {
		return nil, errors.New("onefuse.apiClient: ServiceNow CMDB Deployment Create requires a PolicyID or Policy URL")
	}

	reqJSON, err := json.Marshal(snowDeployment)
	if err != nil {
		return nil, err
	}

	apiurl := c.baseURL
	apiurl.Path = c.apiEndpoint("onefuse", "servicenowCMDBDeployments")

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

func (c *CloudBoltClient) GetServicenowCMDBDeployment(snowDeploymentPath string) (*ServicenowCMDBDeployment, error) {
	apiurl := c.baseURL
	apiurl.Path = snowDeploymentPath

	resp, err := c.makeRequest("GET", apiurl.String(), nil)
	if err != nil {
		log.Fatalln(err)

		return nil, err
	}

	// We Decode the data because we already have an io.Reader on hand
	var snowDeployment ServicenowCMDBDeployment
	json.NewDecoder(resp.Body).Decode(&snowDeployment)

	return &snowDeployment, nil
}

func (c *CloudBoltClient) GetServicenowCMDBDeploymentById(snowDeploymentId string) (*ServicenowCMDBDeployment, error) {
	apiurl := c.baseURL
	apiurl.Path = c.apiEndpoint("onefuse", "servicenowCMDBDeployments", snowDeploymentId)

	resp, err := c.makeRequest("GET", apiurl.String(), nil)
	if err != nil {
		log.Fatalln(err)

		return nil, err
	}

	// We Decode the data because we already have an io.Reader on hand
	var snowDeployment ServicenowCMDBDeployment
	json.NewDecoder(resp.Body).Decode(&snowDeployment)

	return &snowDeployment, nil
}

func (c *CloudBoltClient) DeleteServicenowCMDBDeployment(snowDeploymentId string) (*OneFuseJobStatus, error) {
	apiurl := c.baseURL
	apiurl.Path = c.apiEndpoint("onefuse", "servicenowCMDBDeployments", snowDeploymentId)

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
