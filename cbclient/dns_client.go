package cbclient

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/url"
	"strconv"
)

type DNSPolicyResult struct {
	CloudBoltResult
	Embedded struct {
		DNSPolicies []DNSPolicy `json:"dnsPolicies"`
	} `json:"_embedded"`
}

type DNSPolicy struct {
	Links *struct {
		Self      CloudBoltHALItem `json:"self,omitempty"`
		Workspace CloudBoltHALItem `json:"workspace,omitempty"`
		Endpoint  CloudBoltHALItem `json:"endpoint,omitempty"`
	} `json:"_links,omitempty"`
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type DNSReservation struct {
	Links *struct {
		Self        CloudBoltHALItem `json:"self,omitempty"`
		Workspace   CloudBoltHALItem `json:"workspace,omitempty"`
		Policy      CloudBoltHALItem `json:"policy,omitempty"`
		JobMetadata CloudBoltHALItem `json:"jobMetadata,omitempty"`
	} `json:"_links,omitempty"`
	ID                 int                    `json:"id,omitempty"`
	Name               string                 `json:"name,omitempty"`
	PolicyID           int                    `json:"policyId,omitempty"`
	Policy             string                 `json:"policy,omitempty"`
	WorkspaceURL       string                 `json:"workspace,omitempty"`
	Value              string                 `json:"value,omitempty"`
	Zones              []string               `json:"zones,omitempty"`
	TemplateProperties map[string]interface{} `json:"templateProperties"`
}

func (c *CloudBoltClient) GetDNSPolicy(name string) (*DNSPolicy, error) {
	apiurl := c.baseURL
	apiurl.Path = c.apiEndpoint("onefuse", "dnsPolicies")
	apiurl.RawQuery = fmt.Sprintf(filterByName, url.QueryEscape(name))

	resp, err := c.makeRequest("GET", apiurl.String(), nil)
	if err != nil {
		return nil, err
	}

	// TODO: HTTP Response handling

	var res DNSPolicyResult
	json.NewDecoder(resp.Body).Decode(&res)

	// TODO: Sanity check the decoded object
	if len(res.Embedded.DNSPolicies) == 0 {
		return nil, fmt.Errorf(
			"Could not find DNS Policy with name %s. Does the user have permission to view this?",
			name,
		)
	}

	return &res.Embedded.DNSPolicies[0], nil
}

func (c *CloudBoltClient) CreateDNSReservation(dnsRecord *DNSReservation) (*OneFuseJobStatus, error) {
	log.Println("onefuse.apiClient: CreateDNSReservation")

	if dnsRecord.WorkspaceURL == "" {
		workspace, err := c.GetDefaultWorkSpace()

		if err != nil {
			return nil, err
		}

		dnsRecord.WorkspaceURL = workspace.Links.Self.Href
	}

	if dnsRecord.Policy == "" {
		if dnsRecord.PolicyID != 0 {
			dnsRecord.Policy = c.apiEndpoint("onefuse", "dnsPolicies", strconv.Itoa(dnsRecord.PolicyID))
		} else {
			return nil, errors.New("onefuse.apiClient: DNS Reservation Create requires a PolicyID or Policy URL")
		}
	} else {
		return nil, errors.New("onefuse.apiClient: DNS Reservation Create requires a PolicyID or Policy URL")
	}

	reqJSON, err := json.Marshal(dnsRecord)
	if err != nil {
		return nil, err
	}

	apiurl := c.baseURL
	apiurl.Path = c.apiEndpoint("onefuse", "dnsReservations")

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

func (c *CloudBoltClient) GetDNSReservation(dnsReservationPath string) (*DNSReservation, error) {
	apiurl := c.baseURL
	apiurl.Path = dnsReservationPath

	resp, err := c.makeRequest("GET", apiurl.String(), nil)
	if err != nil {
		log.Fatalln(err)

		return nil, err
	}

	// We Decode the data because we already have an io.Reader on hand
	var dnsRecord DNSReservation
	json.NewDecoder(resp.Body).Decode(&dnsRecord)

	return &dnsRecord, nil
}

func (c *CloudBoltClient) GetDNSReservationById(dnsReservationId string) (*DNSReservation, error) {
	apiurl := c.baseURL
	apiurl.Path = c.apiEndpoint("onefuse", "dnsReservations", dnsReservationId)

	resp, err := c.makeRequest("GET", apiurl.String(), nil)
	if err != nil {
		log.Fatalln(err)

		return nil, err
	}

	// We Decode the data because we already have an io.Reader on hand
	var dnsRecord DNSReservation
	json.NewDecoder(resp.Body).Decode(&dnsRecord)

	return &dnsRecord, nil
}

func (c *CloudBoltClient) DeleteDNSReservation(dnsReservationId string) (*OneFuseJobStatus, error) {
	apiurl := c.baseURL
	apiurl.Path = c.apiEndpoint("onefuse", "dnsReservations", dnsReservationId)

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
