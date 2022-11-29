package cbclient

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/url"
	"strconv"
)

type IPAMPolicyResult struct {
	CloudBoltResult
	Embedded struct {
		IPAMPolicies []IPAMPolicy `json:"ipamPolicies"`
	} `json:"_embedded"`
}

type IPAMPolicy struct {
	Links *struct {
		Self      CloudBoltHALItem `json:"self,omitempty"`
		Workspace CloudBoltHALItem `json:"workspace,omitempty"`
		Endpoint  CloudBoltHALItem `json:"endpoint,omitempty"`
	} `json:"_links,omitempty"`
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type IPAMReservation struct {
	Links *struct {
		Self        CloudBoltHALItem `json:"self,omitempty"`
		Workspace   CloudBoltHALItem `json:"workspace,omitempty"`
		Policy      CloudBoltHALItem `json:"policy,omitempty"`
		JobMetadata CloudBoltHALItem `json:"jobMetadata,omitempty"`
	} `json:"_links,omitempty"`
	ID                 int                    `json:"id,omitempty"`
	Hostname           string                 `json:"hostname,omitempty"`
	PolicyID           int                    `json:"policyId,omitempty"`
	Policy             string                 `json:"policy,omitempty"`
	WorkspaceURL       string                 `json:"workspace,omitempty"`
	IPaddress          string                 `json:"ipAddress,omitempty"`
	Gateway            string                 `json:"gateway,omitempty"`
	PrimaryDNS         string                 `json:"primaryDns"`
	SecondaryDNS       string                 `json:"secondaryDns"`
	Network            string                 `json:"network,omitempty"`
	Subnet             string                 `json:"subnet,omitempty"`
	DNSSuffix          string                 `json:"dnsSuffix,omitempty"`
	Netmask            string                 `json:"netmask,omitempty"`
	NicLabel           string                 `json:"nicLabel,omitempty"`
	TemplateProperties map[string]interface{} `json:"template_properties,omitempty"`
}

func (c *CloudBoltClient) GetIPAMPolicy(name string) (*IPAMPolicy, error) {
	apiurl := c.baseURL
	apiurl.Path = c.apiEndpoint("onefuse", "ipamPolicies")
	apiurl.RawQuery = fmt.Sprintf(filterByName, url.QueryEscape(name))

	resp, err := c.makeRequest("GET", apiurl.String(), nil)
	if err != nil {
		return nil, err
	}

	// TODO: HTTP Response handling

	var res IPAMPolicyResult
	json.NewDecoder(resp.Body).Decode(&res)

	// TODO: Sanity check the decoded object
	if len(res.Embedded.IPAMPolicies) == 0 {
		return nil, fmt.Errorf(
			"Could not find IPAM Policy with name %s. Does the user have permission to view this?",
			name,
		)
	}

	return &res.Embedded.IPAMPolicies[0], nil
}

func (c *CloudBoltClient) CreateIPAMReservation(ipamRecord *IPAMReservation) (*OneFuseJobStatus, error) {
	log.Println("onefuse.apiClient: CreateIPAMReservation")

	if ipamRecord.WorkspaceURL == "" {
		workspace, err := c.GetDefaultWorkSpace()

		if err != nil {
			return nil, err
		}

		ipamRecord.WorkspaceURL = workspace.Links.Self.Href
	}

	log.Println("onefuse.apiClient: CreateModuleDeployment")
	if ipamRecord.Policy == "" {
		if ipamRecord.PolicyID != 0 {
			ipamRecord.Policy = c.apiEndpoint("onefuse", "ipamPolicies", strconv.Itoa(ipamRecord.PolicyID))
		} else {
			return nil, errors.New("onefuse.apiClient: IPAM Reservation Create requires a PolicyID or Policy URL")
		}
	} else {
		return nil, errors.New("onefuse.apiClient: IPAM Reservation Create requires a PolicyID or Policy URL")
	}

	reqJSON, err := json.Marshal(ipamRecord)
	if err != nil {
		return nil, err
	}

	apiurl := c.baseURL
	apiurl.Path = c.apiEndpoint("onefuse", "ipamReservations")

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

func (c *CloudBoltClient) GetIPAMReservation(ipamReservationPath string) (*IPAMReservation, error) {
	apiurl := c.baseURL
	apiurl.Path = ipamReservationPath

	resp, err := c.makeRequest("GET", apiurl.String(), nil)
	if err != nil {
		log.Fatalln(err)

		return nil, err
	}

	// We Decode the data because we already have an io.Reader on hand
	var ipamRecord IPAMReservation
	json.NewDecoder(resp.Body).Decode(&ipamRecord)

	return &ipamRecord, nil
}

func (c *CloudBoltClient) GetIPAMReservationById(ipamReservationId string) (*IPAMReservation, error) {
	apiurl := c.baseURL
	apiurl.Path = c.apiEndpoint("onefuse", "ipamReservations", ipamReservationId)

	resp, err := c.makeRequest("GET", apiurl.String(), nil)
	if err != nil {
		log.Fatalln(err)

		return nil, err
	}

	// We Decode the data because we already have an io.Reader on hand
	var ipamRecord IPAMReservation
	json.NewDecoder(resp.Body).Decode(&ipamRecord)

	return &ipamRecord, nil
}

func (c *CloudBoltClient) DeleteIPAMReservation(ipamReservationId string) (*OneFuseJobStatus, error) {
	apiurl := c.baseURL
	apiurl.Path = c.apiEndpoint("onefuse", "ipamReservations", ipamReservationId)

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
