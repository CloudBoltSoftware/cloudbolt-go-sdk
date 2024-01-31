package cbclient

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
)

// CloudBoltServer stores metadata about servers in CloudBolt.
type CloudBoltServer struct {
	Links struct {
		Self            CloudBoltHALItem         `json:"self"`
		Owner           CloudBoltHALItem         `json:"owner"`
		Group           CloudBoltHALItem         `json:"group"`
		Environment     CloudBoltHALItem         `json:"environment"`
		ResourceHandler CloudBoltHALItem         `json:"resource-handler"`
		Actions         []map[string]interface{} `json:"actions"`
		ProvisionJob    CloudBoltHALItem         `json:"provision-job"`
		OsBuild         CloudBoltHALItem         `json:"os-build"`
		Jobs            CloudBoltHALItem         `json:"jobs"`
		History         CloudBoltHALItem         `json:"history"`
	} `json:"_links"`
	ID                   string        `json:"id"`
	Hostname             string        `json:"hostname"`
	PowerStatus          string        `json:"powerStatus"`
	Status               string        `json:"status"`
	IP                   string        `json:"ipAddress"`
	Mac                  string        `json:"mac"`
	DateAddedToCloudbolt string        `json:"dateAddedToCloudBolt"`
	CPUCount             int           `json:"cpuCount"`
	MemorySizeGB         string        `json:"memorySizeGb"`
	DiskSizeGB           int           `json:"diskSizeGB"`
	OsFamily             string        `json:"osFamily"`
	Notes                string        `json:"notes"`
	Labels               []interface{} `json:"labels"`
	Credentials          struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Key      string `json:"key"`
	} `json:"credentials"`
	RateBreakdown          map[string]interface{}   `json:"rateBreakdown"`
	Disks                  []map[string]interface{} `json:"disks"`
	Snapshots              []map[string]interface{} `json:"snapshots"`
	Networks               []map[string]interface{} `json:"networks"`
	Attributes             []map[string]interface{} `json:"attributes"`
	TechSpecificAttributes map[string]interface{}   `json:"techSpecificAttributes"`
}

type CloudBoltServerResult struct {
	CloudBoltResult
	Embedded struct {
		Servers []CloudBoltServer `json:"servers"`
	} `json:"_embedded"`
}

type CloudBoltDecomServerResult struct {
	Links struct {
		Self CloudBoltHALItem `json:"self"`
	} `json:"_links"`
	ID string `json:"id"`
}

// GetServer fetches a Server object from CloudBolt at the given path
// - Server Path (serverPath) e.g., "/api/v2/servers/123/"
func (c *CloudBoltClient) GetServer(serverPath string) (*CloudBoltServer, error) {
	apiurl := c.baseURL
	apiurl.Path = serverPath

	resp, err := c.makeRequest("GET", apiurl.String(), nil)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	// We Decode the data because we already have an io.Reader on hand
	var svr CloudBoltServer
	json.NewDecoder(resp.Body).Decode(&svr)

	return &svr, nil
}

func (c *CloudBoltClient) GetServerById(id string) (*CloudBoltServer, error) {
	apiurl := c.baseURL
	apiurl.Path = c.apiEndpoint("cmp", "servers", id)

	resp, err := c.makeRequest("GET", apiurl.String(), nil)
	if err != nil {
		return nil, err
	}

	var svr CloudBoltServer
	json.NewDecoder(resp.Body).Decode(&svr)

	return &svr, nil
}

func (c *CloudBoltClient) GetServerByHostname(hostname string) (*CloudBoltServer, error) {
	apiurl := c.baseURL
	apiurl.Path = c.apiEndpoint("cmp", "servers")
	apiurl.RawQuery = fmt.Sprintf("filter=hostname:%s;status:ACTIVE", url.QueryEscape(hostname))

	resp, err := c.makeRequest("GET", apiurl.String(), nil)
	if err != nil {
		return nil, err
	}

	// We Decode the data because we already have an io.Reader on hand
	var res CloudBoltServerResult
	json.NewDecoder(resp.Body).Decode(&res)

	// TODO: Sanity check the decoded object
	if len(res.Embedded.Servers) == 0 {
		return nil, fmt.Errorf(
			"Could not find active server with hostname %s. Does the user have permission to view this?",
			hostname,
		)
	}

	if len(res.Embedded.Servers) > 1 {
		return nil, fmt.Errorf(
			"More than one active server with hostname %s found.",
			hostname,
		)
	}

	return &res.Embedded.Servers[0], nil
}

func (c *CloudBoltClient) DecomServer(serverId string) (*CloudBoltDecomServerResult, error) {
	apiurl := c.baseURL
	apiurl.Path = c.apiEndpoint(
		"cmp",
		"servers",
		serverId,
		"decommission",
	)

	resp, err := c.makeRequest("POST", apiurl.String(), nil)
	if err != nil {
		return nil, err
	}

	// We Decode the data because we already have an io.Reader on hand
	var decomResult CloudBoltDecomServerResult
	json.NewDecoder(resp.Body).Decode(&decomResult)

	return &decomResult, nil
}
