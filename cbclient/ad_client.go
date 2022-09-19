package cbclient

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/url"
	"strconv"
)

type ADPolicyResult struct {
	CloudBoltResult
	Embedded struct {
		ADPolicies []ADPolicy `json:"microsoftADPolicies"`
	} `json:"_embedded"`
}

type ADPolicy struct {
	Links *struct {
		Self      CloudBoltHALItem `json:"self,omitempty"`
		Workspace CloudBoltHALItem `json:"workspace,omitempty"`
		Endpoint  CloudBoltHALItem `json:"microsoftEndpoint"`
	} `json:"_links,omitempty"`
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type MicrosoftADPolicy struct {
	Links *struct {
		Self              CloudBoltHALItem `json:"self,omitempty"`
		Workspace         CloudBoltHALItem `json:"workspace,omitempty"`
		MicrosoftEndpoint CloudBoltHALItem `json:"microsoftEndpoint,omitempty"`
	} `json:"_links,omitempty"`
	Name                   string   `json:"name,omitempty"`
	ID                     int      `json:"id,omitempty"`
	Description            string   `json:"description,omitempty"`
	MicrosoftEndpointID    int      `json:"microsoftEndpointId,omitempty"`
	MicrosoftEndpoint      string   `json:"microsoftEndpoint,omitempty"`
	ComputerNameLetterCase string   `json:"computerNameLetterCase,omitempty"`
	WorkspaceURL           string   `json:"workspace,omitempty"`
	OU                     string   `json:"ou,omitempty"`
	CreateOU               bool     `json:"createOrganizationalUnit,omitempty"`
	RemoveOU               bool     `json:"removeOrganizationalUnit,omitempty"`
	SecurityGroups         []string `json:"securityGroups,omitempty"`
}

type MicrosoftADComputerAccount struct {
	Links *struct {
		Self        CloudBoltHALItem `json:"self,omitempty"`
		Workspace   CloudBoltHALItem `json:"workspace,omitempty"`
		Policy      CloudBoltHALItem `json:"policy,omitempty"`
		JobMetadata CloudBoltHALItem `json:"jobMetadata,omitempty"`
	} `json:"_links,omitempty"`
	ID                 int                    `json:"id,omitempty"`
	Name               string                 `json:"name,omitempty"`
	FinalOU            string                 `json:"finalOu"`
	PolicyID           int                    `json:"policyId,omitempty"`
	Policy             string                 `json:"policy,omitempty"`
	WorkspaceURL       string                 `json:"workspace,omitempty"`
	TemplateProperties map[string]interface{} `json:"templateProperties"`
}

func (c *CloudBoltClient) GetADPolicy(name string) (*ADPolicy, error) {
	apiurl := c.baseURL
	apiurl.Path = c.apiEndpoint("onefuse", "microsoftADPolicies")
	apiurl.RawQuery = fmt.Sprintf(filterByName, url.QueryEscape(name))

	resp, err := c.makeRequest("GET", apiurl.String(), nil)
	if err != nil {
		return nil, err
	}

	// TODO: HTTP Response handling

	var res ADPolicyResult
	json.NewDecoder(resp.Body).Decode(&res)

	// TODO: Sanity check the decoded object
	if len(res.Embedded.ADPolicies) == 0 {
		return nil, fmt.Errorf(
			"Could not find AD Policy with name %s. Does the user have permission to view this?",
			name,
		)
	}

	return &res.Embedded.ADPolicies[0], nil
}

func (c *CloudBoltClient) CreateMicrosoftADComputerAccount(computerAccount *MicrosoftADComputerAccount) (*OneFuseJobStatus, error) {
	log.Println("onefuse.apiClient: CreateMicrosoftADComputerAccount")

	if computerAccount.WorkspaceURL == "" {
		workspace, err := c.GetDefaultWorkSpace()

		if err != nil {
			return nil, err
		}

		computerAccount.WorkspaceURL = workspace.Links.Self.Href
	}

	if computerAccount.Policy == "" {
		if computerAccount.PolicyID != 0 {
			computerAccount.Policy = c.apiEndpoint("onefuse", "microsoftADPolicies", strconv.Itoa(computerAccount.PolicyID))
		} else {
			return nil, errors.New("onefuse.apiClient: Microsoft AD Computer Account Create requires a PolicyID or Policy URL")
		}
	} else {
		return nil, errors.New("onefuse.apiClient: Microsoft AD Computer Account Create requires a PolicyID or Policy URL")
	}

	reqJSON, err := json.Marshal(computerAccount)
	if err != nil {
		return nil, err
	}

	apiurl := c.baseURL
	apiurl.Path = c.apiEndpoint("onefuse", "microsoftADComputerAccounts")

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

func (c *CloudBoltClient) GetMicrosoftADComputerAccount(computerAccountPath string) (*MicrosoftADComputerAccount, error) {
	apiurl := c.baseURL
	apiurl.Path = computerAccountPath

	resp, err := c.makeRequest("GET", apiurl.String(), nil)
	if err != nil {
		log.Fatalln(err)

		return nil, err
	}

	// We Decode the data because we already have an io.Reader on hand
	var computerAccount MicrosoftADComputerAccount
	json.NewDecoder(resp.Body).Decode(&computerAccount)

	return &computerAccount, nil
}

func (c *CloudBoltClient) GetMicrosoftADComputerAccountById(computerAccountId string) (*MicrosoftADComputerAccount, error) {
	apiurl := c.baseURL
	apiurl.Path = c.apiEndpoint("onefuse", "microsoftADComputerAccounts", computerAccountId)

	resp, err := c.makeRequest("GET", apiurl.String(), nil)
	if err != nil {
		log.Fatalln(err)

		return nil, err
	}

	// We Decode the data because we already have an io.Reader on hand
	var computerAccount MicrosoftADComputerAccount
	json.NewDecoder(resp.Body).Decode(&computerAccount)

	return &computerAccount, nil
}

func (c *CloudBoltClient) DeleteMicrosoftADComputerAccount(computerAccountId string) (*OneFuseJobStatus, error) {
	apiurl := c.baseURL
	apiurl.Path = c.apiEndpoint("onefuse", "microsoftADComputerAccounts", computerAccountId)

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

func (c *CloudBoltClient) CreateMicrosoftADPolicy(newPolicy *MicrosoftADPolicy) (*MicrosoftADPolicy, error) {
	log.Println("onefuse.apiClient: CreateModuleDeployment")

	if newPolicy.WorkspaceURL == "" {
		workspace, err := c.GetDefaultWorkSpace()

		if err != nil {
			return nil, err
		}

		newPolicy.WorkspaceURL = workspace.Links.Self.Href
	}

	reqJSON, err := json.Marshal(newPolicy)
	if err != nil {
		return nil, err
	}

	apiurl := c.baseURL
	apiurl.Path = c.apiEndpoint("onefuse", "microsoftADPolicies")

	resp, err := c.makeRequest("POST", apiurl.String(), reqJSON)
	if err != nil {
		return nil, err
	}

	// Handle some common HTTP errors
	err = checkHttpStatus(resp)
	if err != nil {
		return nil, err
	}

	var adPolicy MicrosoftADPolicy
	json.NewDecoder(resp.Body).Decode(&adPolicy)

	return &adPolicy, nil
}

func (c *CloudBoltClient) GetMicrosoftADPolicyByID(policyId string) (*MicrosoftADPolicy, error) {
	apiurl := c.baseURL
	apiurl.Path = c.apiEndpoint("onefuse", "microsoftADPolicies", policyId)

	resp, err := c.makeRequest("GET", apiurl.String(), nil)
	if err != nil {
		log.Fatalln(err)

		return nil, err
	}

	// We Decode the data because we already have an io.Reader on hand
	var adPolicy MicrosoftADPolicy
	json.NewDecoder(resp.Body).Decode(&adPolicy)

	return &adPolicy, nil
}

func (c *CloudBoltClient) UpdateMicrosoftADPolicy(policyId string, updatedPolicy *MicrosoftADPolicy) (*MicrosoftADPolicy, error) {
	reqJSON, err := json.Marshal(updatedPolicy)
	if err != nil {
		return nil, err
	}

	apiurl := c.baseURL
	apiurl.Path = c.apiEndpoint("onefuse", "microsoftADPolicies", policyId)

	resp, err := c.makeRequest("PUT", apiurl.String(), reqJSON)
	if err != nil {
		return nil, err
	}

	// Handle some common HTTP errors
	err = checkHttpStatus(resp)
	if err != nil {
		return nil, err
	}

	var adPolicy MicrosoftADPolicy
	json.NewDecoder(resp.Body).Decode(&adPolicy)

	return &adPolicy, nil
}

func (c *CloudBoltClient) DeleteMicrosoftADPolicy(policyId string) error {
	apiurl := c.baseURL
	apiurl.Path = c.apiEndpoint("onefuse", "microsoftADPolicies", policyId)

	resp, err := c.makeRequest("DELETE", apiurl.String(), nil)
	if err != nil {
		return err
	}

	// Handle some common HTTP errors
	err = checkHttpStatus(resp)
	if err != nil {
		return err
	}

	return nil
}
