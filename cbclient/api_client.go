package cbclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"
)

const filterByName string = "filter=name:%s"

// CloudBoltObject stores the generic output of objects.
// Most objects in CloudBolt include Links.Self.Href, Links.Self.Title, Name, and ID
type CloudBoltObject struct {
	Links struct {
		Self CloudBoltHALItem `json:"self"`
	} `json:"_links"`
	Name string `json:"name"`
	ID   string `json:"id"`
}

// CloudBoltClient stores the important metadata necessary to make API requests.
// - BaseURL follows the pattern "https://cloudbolt.myco.ext:443/".
// - HTTPClient is a client used to make the API calls.
// - Token is retrieved in `New` and is included in the Bearer Token of request headers.
type CloudBoltClient struct {
	baseURL    url.URL
	httpClient *http.Client
	password   string
	token      string
	username   string
	apiVersion string
}

// CloudBoltResult stores the response of paginated calls like `/api/v2/blueprints/`
// These include a link to the page and an `embedded` list of response objects.
type CloudBoltResult struct {
	Links struct {
		Self CloudBoltHALItem `json:"self"`
	} `json:"_links"`
	Total    int               `json:"total"`
	Count    int               `json:"count"`
	Embedded []CloudBoltObject `json:"_embedded"` // TODO: Maybe call this Items?
}

// CloudBoltActionResult stores metadata about the result of running an action.
type CloudBoltActionResult struct {
	RunActionJob struct {
		Self CloudBoltHALItem `json:"self"`
	} `json:"run-action-job"`
}

// CloudBoltHALItem stores an object's title and API endpoint.
// This is a common pattern in the CloudBolt API, so it gets used a lot.
type CloudBoltHALItem struct {
	Href  string `json:"href"`
	Title string `json:"title"`
}

// CloudBoltOrder contains metadata about a CloudBolt Order
// Returned by DeployBlueprint, GetOrder, and DecomOrder
type CloudBoltOrder struct {
	Links struct {
		Self       CloudBoltHALItem   `json:"self"`
		Group      CloudBoltHALItem   `json:"group"`
		Owner      CloudBoltHALItem   `json:"owner"`
		ApprovedBy CloudBoltHALItem   `json:"approved-by"`
		Actions    CloudBoltHALItem   `json:"actions"`
		Jobs       []CloudBoltHALItem `json:"jobs"`
	} `json:"_links"`
	Name        string `json:"name"`
	ID          string `json:"id"`
	Status      string `json:"status"`
	Rate        string `json:"rate"`
	CreateDate  string `json:"create-date"`
	ApproveDate string `json:"approve-date"`
	Items       struct {
		DeployItems []struct {
			Blueprint               string `json:"blueprint"`
			BlueprintItemsArguments struct {
				BuildItemBuildServer struct {
					Attributes struct {
						Hostname string `json:"hostname"`
						Quantity int    `json:"quantity"`
					} `json:"attributes"`
					OsBuild     string                 `json:"os-build,omitempty"`
					Environment string                 `json:"environment,omitempty"`
					Parameters  map[string]interface{} `json:"parameters"`
				} `json:"build-item-Server"`
			} `json:"blueprint-items-arguments"`
			ResourceName       string `json:"resource-name"`
			ResourceParameters struct {
			} `json:"resource-parameters"`
		} `json:"deploy-items"`
	} `json:"items"`
}

// CloudBoltJob contains metadata about a Job.
// Useful for getting the status of a running or completed job.
type CloudBoltJob struct {
	Links struct {
		Self         CloudBoltHALItem `json:"self"`
		Owner        CloudBoltHALItem `json:"owner"`
		Parent       CloudBoltHALItem `json:"parent"`
		Subjobs      []interface{}    `json:"subjobs"`
		Prerequisite struct {
		} `json:"prerequisite"`
		DependentJobs []interface{}      `json:"dependent-jobs"`
		Order         CloudBoltHALItem   `json:"order"`
		Resource      CloudBoltHALItem   `json:"resource"`
		Servers       []CloudBoltHALItem `json:"servers"`
		LogUrls       struct {
			RawLog string `json:"raw-log"`
			ZipLog string `json:"zip-log"`
		} `json:"log_urls"`
	} `json:"_links"`
	Status   string `json:"status"`
	Type     string `json:"type"`
	Progress struct {
		TotalTasks int      `json:"total-tasks"`
		Completed  int      `json:"completed"`
		Messages   []string `json:"messages"`
	} `json:"progress"`
	StartDate string `json:"start-date"`
	EndDate   string `json:"end-date"`
	Output    string `json:"output"`
}

// CloudBoltGroup contains metadata about a Group in CloudBolt
type CloudBoltGroup struct {
	Links struct {
		Self                  CloudBoltHALItem   `json:"self"`
		Parent                CloudBoltHALItem   `json:"parent"`
		Subgroups             []CloudBoltHALItem `json:"subgroups"`
		Environments          []interface{}      `json:"environments"`
		OrderableEnvironments CloudBoltHALItem   `json:"orderable-environments"`
	} `json:"_links"`
	Name         string `json:"name"`
	ID           string `json:"id"`
	Type         string `json:"type"`
	Rate         string `json:"rate"`
	AutoApproval bool   `json:"auto-approval"`
}

// CloudBoltResource contains metadata about Resources (e.g., "Services") in CloudBolt
type CloudBoltResource struct {
	Links struct {
		Self         CloudBoltHALItem `json:"self"`
		Blueprint    CloudBoltHALItem `json:"blueprint"`
		Owner        CloudBoltHALItem `json:"owner"`
		Group        CloudBoltHALItem `json:"group"`
		ResourceType CloudBoltHALItem `json:"resource-type"`
		Servers      []struct {
			Href  string `json:"href"`
			Title string `json:"title"`
			Tier  string `json:"tier"`
		} `json:"servers"`
		Actions []struct {
			Delete CloudBoltHALItem `json:"Delete,omitempty"`
			Scale  CloudBoltHALItem `json:"Scale,omitempty"`
		} `json:"actions"`
		Jobs    CloudBoltHALItem `json:"jobs"`
		History CloudBoltHALItem `json:"history"`
	} `json:"_links"`
	Name        string `json:"name"`
	ID          string `json:"id"`
	Status      string `json:"status"`
	InstallDate string `json:"install-date"`
}

// CloudBoltServer stores metadata about servers in CloudBolt.
type CloudBoltServer struct {
	Links struct {
		Self            CloudBoltHALItem `json:"self"`
		Owner           CloudBoltHALItem `json:"owner"`
		Group           CloudBoltHALItem `json:"group"`
		Environment     CloudBoltHALItem `json:"environment"`
		ResourceHandler CloudBoltHALItem `json:"resource-handler"`
		Actions         []struct {
			PowerOn     CloudBoltHALItem `json:"power_on,omitempty"`
			PowerOff    CloudBoltHALItem `json:"power_off,omitempty"`
			Reboot      CloudBoltHALItem `json:"reboot,omitempty"`
			RefreshInfo CloudBoltHALItem `json:"refresh_info,omitempty"`
			Snapshot    CloudBoltHALItem `json:"snapshot,omitempty"`
			AdHocScript CloudBoltHALItem `json:"Ad Hoc Script,omitempty"`
		} `json:"actions"`
		ProvisionJob CloudBoltHALItem `json:"provision-job"`
		OsBuild      CloudBoltHALItem `json:"os-build"`
		Jobs         CloudBoltHALItem `json:"jobs"`
		History      CloudBoltHALItem `json:"history"`
	} `json:"_links"`
	Hostname             string        `json:"hostname"`
	PowerStatus          string        `json:"power-status"`
	Status               string        `json:"status"`
	IP                   string        `json:"ip"`
	Mac                  string        `json:"mac"`
	DateAddedToCloudbolt string        `json:"date-added-to-cloudbolt"`
	CPUCnt               int           `json:"cpu-cnt"`
	MemSize              string        `json:"mem-size"`
	DiskSize             string        `json:"disk-size"`
	OsFamily             string        `json:"os-family"`
	Notes                string        `json:"notes"`
	Labels               []interface{} `json:"labels"`
	Credentials          struct {
		Username string `json:"username"`
		Password string `json:"password"`
	} `json:"credentials"`
	// TODO: This should be a more generic map
	Disks []struct {
		UUID             string `json:"uuid"`
		DiskSize         int    `json:"disk-size"`
		Name             string `json:"name"`
		Datastore        string `json:"datastore"`
		ProvisioningType string `json:"provisioning-type"`
	} `json:"disks"`
	Networks []struct {
		Name          string      `json:"name"`
		Network       string      `json:"network"`
		Mac           string      `json:"mac"`
		IP            interface{} `json:"ip"`
		PrivateIP     string      `json:"private-ip"`
		AdditionalIps string      `json:"additional-ips"`
	} `json:"networks"`
	Parameters struct {
	} `json:"parameters"`
	// TODO: This should be a more generic map
	TechSpecificDetails struct {
		VmwareLinkedClone bool   `json:"vmware-linked-clone"`
		VmwareCluster     string `json:"vmware-cluster"`
	} `json:"tech-specific-details"`
}

type cloudBoltUserAuthCreds struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type cloudBoltUserAuthToken struct {
	Token string `json:"token"`
}

// New returns an initialized CloudBoltClient object.
// Accepts as input:
// - HTTP Protocol (protocol) e.g., "https"
// - HTTP Host (host) e.g., "cloudbolt.intranet"
// - HTTP Port (port) e.g., "443"
// - apiVersion (apiVersion) e.g., "v2"
// - Username (username) e.g., "myUserName"
// - Password (password) e.g., "My Passphrase!"
// - User-provided *HTTPClient (httpClient); provide `nil` to get a server with the following defaults:
//   - Timeout set to 60 seconds
//   Provide a custom http.Client if you require unique certificate, timeout, etc., configured.
//
// New does not make any API calls.
// CloudBoltClient.Authenticate must be called to initialize CloudBoltClient.token.
// This is done automatically when a request receives an HTTP Authorization error.
func New(protocol string, host string, port string, apiVersion string, username string, password string, httpClient *http.Client) *CloudBoltClient {
	baseURL := url.URL{
		Scheme: protocol,
		Host:   fmt.Sprintf("%s:%s", host, port),
	}

	var client *http.Client
	if httpClient == nil {
		// The given Client is `nil` so we provide a sane default
		client = &http.Client{
			Timeout: 60 * time.Second,
		}
	} else {
		// The user provided an HTTP Client so we pass that through to the CloudBoltClient
		client = httpClient
	}

	return &CloudBoltClient{
		baseURL:    baseURL,
		httpClient: client,
		username:   username,
		password:   password,
		apiVersion: apiVersion,
	}
}

// Authenticate forces the CloudBoltClient to re-authenticate
// Returns an error if there is an HTTP error, or if the HTTP Status Code is >=400
func (c *CloudBoltClient) Authenticate() (int, error) {
	// Craft the JSON payload used to request an API token
	userCreds := cloudBoltUserAuthCreds{
		Username: c.username,
		Password: c.password,
	}
	reqJSON, err := json.Marshal(userCreds)
	if err != nil {
		return -1, err
	}

	// Put the username+password into a bytes buffer for the API request
	reqJSONBytes := bytes.NewBuffer(reqJSON)

	// Craft the URL api-token request endpoint based on the API version
	apiurl := c.baseURL
	apiurl.Path = c.apiEndpoint("api-token-auth")

	// Make the POST request to get the API token
	req, err := http.NewRequest("POST", apiurl.String(), reqJSONBytes)
	if err != nil {
		return -1, err
	}
	req.Header.Set("Content-Type", "application/json")

	// Execute the HTTP request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return -1, fmt.Errorf("Failed to create the API client. %s", err)
	}

	// We received a bad HTTP request, so forward that to the caller before trying to parse the response
	if resp.StatusCode >= 400 {
		return resp.StatusCode, fmt.Errorf("Received bad HTTP response %d: %s", resp.StatusCode, resp.Status)
	}

	// We Decode the data because we already have an io.Reader on hand
	var userAuthData cloudBoltUserAuthToken
	json.NewDecoder(resp.Body).Decode(&userAuthData)

	// Set the CloudBoltClient token as that parsed Token value
	c.token = userAuthData.Token

	// Return the HTTP status code and a nil error for success
	return resp.StatusCode, nil
}

// makeRequest wrapps most HTTP requests by adding:
// - Set Content Type to JSON
// - Set Accept to JSON
// - Calling authWrappedRequest
func (c *CloudBoltClient) makeRequest(req *http.Request) (*http.Response, error) {
	// Sending and Accepting JSON
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	return c.authWrappedRequest(req)
}

// makeRequest wraps the normal HTTP request by re-authenticating if we get an
// "Unauthorized" HTTP response.
//
// if the first attempt at the request returns a 401 or 403 HTTP Status Code
// it Attempts exactly one call to CloudBoltClient.Authenticate() and resets the request token.
func (c *CloudBoltClient) authWrappedRequest(req *http.Request) (*http.Response, error) {
	// Add the Auth token to the request
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))

	// Attempt to make the given HTTP request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	switch resp.StatusCode {
	// Handle common HTTP "auth" related Status Codes
	case 401, 403:
		// Re-authenticate with the API
		_, err := c.Authenticate()
		if err != nil {
			return nil, err
		}

		// Re-add the auth token which should have updated
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))

		// Make the request again
		// We assume we have valid Auth credentials
		resp, err := c.httpClient.Do(req)
		if err != nil {
			return nil, err
		}

		// Return the post-re-auth response
		// This may still get a bad HTTP error
		return resp, nil

	// If we didn't get one of the those 40x Status Codes,
	// then pass through the original result
	default:
		// Return the original response
		return resp, nil
	}
}

// apiEndpoint standardizes getting a CloudBolt API endpoint
// Pass in a variadic number of entries and they are formatted like so:
// apiEndpoint("foo", "bar", "baz") -> /api/{apiVersion}/foo/bar/baz/
//
// Only returns the path, not the prepending "https://host:port"
func (c *CloudBoltClient) apiEndpoint(paths ...string) string {
	// Create a slice ["api", "someVersion"]
	basePathSlice := []string{
		"api",
		c.apiVersion,
	}

	// Concatenate basePathSlice with the user provided paths
	fullPath := append(
		basePathSlice,
		paths...,
	)

	// Format the array of path entries as "api/{apiVersion}/path/to/thing"
	// Note this does not include the root and trailing "/" which we want
	formattedPath := path.Join(fullPath...)

	// Formats the result to include the root and trailing slashes
	// e.g., "/api/{apiVersion}/path/to/thing/"
	return fmt.Sprintf("/%s/", formattedPath)
}

// GetCloudBoltObject fetches a given object of type "objPath" with the name "objName"
// e.g., GetCloudBoltObject("users", "Susan") gets the user with username "Susan"
//
// Caveat:
//   This makes a generic request to `/api/v2/objPath/?filter=name:objName`
//   it returns the first element of the `_embedded` list, so if multiple objects are
//   returned from this query, only the first one will be returned.
//
//   So don't lean on this without some sanity checks.
func (c *CloudBoltClient) GetCloudBoltObject(objPath string, objName string) (*CloudBoltObject, error) {
	apiurl := c.baseURL
	apiurl.Path = c.apiEndpoint(objPath)
	apiurl.RawQuery = fmt.Sprintf(filterByName, url.QueryEscape(objName))

	// log.Printf("[!!] apiurl in GetCloudBoltObject: %+v (%+v)", apiurl.String(), apiurl)

	req, err := http.NewRequest("GET", apiurl.String(), nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.makeRequest(req)
	if err != nil {
		return nil, err
	}
	// log.Printf("[!!] HTTP response: %+v", resp)

	// TODO: HTTP Response handling

	// We Decode the data because we already have an io.Reader on hand
	var res CloudBoltResult
	json.NewDecoder(resp.Body).Decode(&res)

	// TODO: Sanity check the decoded object

	// log.Printf("[!!] CloudBoltResult response %+v", res) // HERE IS WHERE THE PANIC IS!!!
	if len(res.Embedded) == 0 {
		return nil, fmt.Errorf(
			"Could not find %s with name %s. Does the user have permission to view this?",
			objPath,
			objName,
		)
	}
	return &res.Embedded[0], nil
}

// verifyGroup checks that all a given group is the one we intended to fetch.
//
// groupPath is the API path to the group, e.g., "/api/v2/groups/GRP-123456"
//
// If a group has no parents, "parentPath" should be empty.
// If a group has parents, it should be of the format "root-level-parent/sub-parent/.../closest-parent"
func (c *CloudBoltClient) verifyGroup(groupPath string, parentPath string) (bool, error) {
	// log.Printf("Verifying group %+v with parent(s) %+v\n", groupPath, parentPath)
	var parent string
	var nextParentPath string

	apiurl := c.baseURL
	apiurl.Path = groupPath

	req, err := http.NewRequest("GET", apiurl.String(), nil)
	if err != nil {
		return false, err
	}

	resp, err := c.makeRequest(req)
	if err != nil {
		return false, err
	}
	if resp.StatusCode >= 300 {
		log.Fatalln(resp.Status)

		return false, errors.New(resp.Status)
	}

	// We Decode the data because we already have an io.Reader on hand
	var group CloudBoltGroup
	json.NewDecoder(resp.Body).Decode(&group)

	nextIndex := strings.LastIndex(parentPath, "/")

	// log.Printf("[!!] parentPath: %+v", parentPath)
	// log.Printf("[!!] strings.LastIndex(parentPath, '/')+1: %+v", nextIndex+1)
	if nextIndex >= 0 {
		parent = parentPath[nextIndex+1:]
		nextParentPath = parentPath[:nextIndex]
		// log.Printf("[!!] NextIndex >= 0 so parent: %+v, nextParentPath %+v", parent, nextParentPath)
	} else {
		parent = parentPath
		// log.Printf("[!!] NextIndex < 0 so parent: %+v", parent)
	}

	if group.Links.Parent.Title != parent {
		// log.Printf("[!!] group.Links.Parent.Title '%+v' !=? parent '%+v'\nReturning False\n", group.Links.Parent.Title, parent)
		return false, nil
	}

	// log.Printf("[!!] nextParentPath: %+v", nextParentPath)
	if nextParentPath != "" {
		// log.Printf("[!!] NextParentPath is not empty, making recursive call in verifyGroup\n")
		return c.verifyGroup(group.Links.Parent.Href, nextParentPath)
	}

	// log.Printf("[!!] Group verified, returning true\n")
	return true, nil
}

// GetGroup accepts a groupPath string parameter of the following format:
// "/my parent group/some subgroup/a child group/" or just "my parent group"
//
// verifyGroup recursively verifies that this is a valid group/subgroup.
func (c *CloudBoltClient) GetGroup(groupPath string) (*CloudBoltObject, error) {
	var group string
	var parentPath string
	var groupFound bool

	groupPath = strings.Trim(groupPath, "/")
	nextIndex := strings.LastIndex(groupPath, "/")

	// log.Printf("[!!] groupPath: %+v", groupPath)
	// log.Printf("[!!] nextIndex %+v", nextIndex)
	// log.Printf("[!!] strings.LastIndex(groupPath, '/')+1: %+v", strings.LastIndex(groupPath, "/")+1)
	if nextIndex >= 0 {
		group = groupPath[nextIndex+1:]
		parentPath = groupPath[:nextIndex]
		// log.Printf("[!!] group: %+v // parentGroup: %+v\n", group, parentPath)
	} else {
		group = groupPath
	}

	apiurl := c.baseURL
	apiurl.Path = c.apiEndpoint("groups")
	apiurl.RawQuery = fmt.Sprintf(filterByName, url.QueryEscape(group))

	// log.Printf("[!!] apiurl in GetGroup: %+v (%+v)", apiurl.String(), apiurl)

	req, err := http.NewRequest("GET", apiurl.String(), nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.makeRequest(req)
	if err != nil {
		return nil, err
	}

	// We Decode the data because we already have an io.Reader on hand
	var res CloudBoltResult
	json.NewDecoder(resp.Body).Decode(&res)

	for _, v := range res.Embedded {
		// log.Printf("Group is %+v\n", v)
		groupFound, err = c.verifyGroup(v.Links.Self.Href, parentPath)

		if groupFound {
			return &v, nil
		}
	}

	return nil, fmt.Errorf("Group (%s): Not Found", groupPath)
}

// DeployBlueprint deploys the given:
// - Blueprint (bpPath) e.g., "/api/v2/blueprints/BG-6ic2tw7x/"
// - Group (grpPath) e.g., "/api/v2/groups/GRP-5ukhved7/"
// - Resource Name (resourceName) e.g., "My Resource Name"
// - Blueprint Items/request parameters (bpItems) e.g.,
//       []map[string]interface{}{
//           map[string]interface{}{
//               "bp-item-name": "bp item name",
//               "bp-item-paramas": map[string]interface{}}
//                   "some-param": "param value",
//                   "other-param": "foo bar baz",
//               },
//               "environment":     "bp environment",
//               "osbuild":         "bp osbuild",
//           }
//       }
func (c *CloudBoltClient) DeployBlueprint(grpPath string, bpPath string, resourceName string, bpItems []map[string]interface{}) (*CloudBoltOrder, error) {
	deployItems := make([]map[string]interface{}, 0)

	for _, v := range bpItems {
		bpItem := map[string]interface{}{
			"blueprint": bpPath,
			"blueprint-items-arguments": map[string]interface{}{
				v["bp-item-name"].(string): map[string]interface{}{
					"attributes": map[string]interface{}{
						"quantity": 1,
					},
					"parameters": v["bp-item-paramas"].(map[string]interface{}),
				},
			},
			"resource-name": resourceName,
		}

		env, ok := v["environment"]
		if ok {
			bpItem["blueprint-items-arguments"].(map[string]interface{})[v["bp-item-name"].(string)].(map[string]interface{})["environment"] = env
		}

		osb, ok := v["os-build"]
		if ok {
			bpItem["blueprint-items-arguments"].(map[string]interface{})[v["bp-item-name"].(string)].(map[string]interface{})["os-build"] = osb
		}

		deployItems = append(deployItems, bpItem)
	}

	reqData := map[string]interface{}{
		"group": grpPath,
		"items": map[string]interface{}{
			"deploy-items": deployItems,
		},
		"submit-now": "true",
	}

	reqJSON, err := json.Marshal(reqData)
	if err != nil {
		return nil, err
	}

	// log.Printf("[!!] JSON payload in POST request to Deploy Blueprint:\n%s", string(reqJSON))

	apiurl := c.baseURL
	apiurl.Path = c.apiEndpoint("orders")

	// log.Printf("[!!] apiurl in DeployBlueprint: %+v (%+v)", apiurl.String(), apiurl)

	reqBody := bytes.NewBuffer(reqJSON)

	req, err := http.NewRequest("POST", apiurl.String(), reqBody)
	if err != nil {
		return nil, err
	}

	resp, err := c.makeRequest(req)
	if err != nil {
		return nil, err
	}

	// Handle some common HTTP errors
	switch {
	case resp.StatusCode >= 500:
		buf := new(bytes.Buffer)
		buf.ReadFrom(resp.Body)
		respBody := buf.String()
		return nil, fmt.Errorf("received a server error: %s", respBody)
	case resp.StatusCode >= 400:
		buf := new(bytes.Buffer)
		buf.ReadFrom(resp.Body)
		respBody := buf.String()
		return nil, fmt.Errorf("received an HTTP client error: %s", respBody)
	default:
		// We Decode the data because we already have an io.Reader on hand
		var order CloudBoltOrder
		json.NewDecoder(resp.Body).Decode(&order)

		return &order, nil
	}
}

// GetOrder fetches an Order from CloudBolt
// - Order ID (orderID) e.g., "123"; formatted into a string like "/api/v2/orders/123"
func (c *CloudBoltClient) GetOrder(orderID string) (*CloudBoltOrder, error) {
	apiurl := c.baseURL
	apiurl.Path = c.apiEndpoint("orders", orderID)

	// log.Printf("[!!] apiurl in GetOrder: %+v (%+v)", apiurl.String(), apiurl)

	req, err := http.NewRequest("GET", apiurl.String(), nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.makeRequest(req)
	if err != nil {
		log.Fatalln(err)

		return nil, err
	}

	// We Decode the data because we already have an io.Reader on hand
	var order CloudBoltOrder
	json.NewDecoder(resp.Body).Decode(&order)

	return &order, nil
}

// GetJob fetches the Job object from CloudBolt at the given path
// - Job Path (jobPath) e.g., "/api/v2/jobs/123/"
func (c *CloudBoltClient) GetJob(jobPath string) (*CloudBoltJob, error) {
	apiurl := c.baseURL
	apiurl.Path = jobPath

	// log.Printf("[!!] GetJob: %+v (%+v)", apiurl.String(), apiurl)

	req, err := http.NewRequest("GET", apiurl.String(), nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.makeRequest(req)
	if err != nil {
		log.Fatalln(err)

		return nil, err
	}

	// We Decode the data because we already have an io.Reader on hand
	var job CloudBoltJob
	json.NewDecoder(resp.Body).Decode(&job)

	return &job, nil
}

// GetResource fetches a Resource object from CloudBolt at the given path
// - Resource Path (resourcePath) e.g., "/api/v2/resources/service/123/"
func (c *CloudBoltClient) GetResource(resourcePath string) (*CloudBoltResource, error) {
	apiurl := c.baseURL
	apiurl.Path = resourcePath

	// log.Printf("[!!] apiurl in GetResource: %+v (%+v)", apiurl.String(), apiurl)

	req, err := http.NewRequest("GET", apiurl.String(), nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.makeRequest(req)
	if err != nil {
		log.Fatalln(err)

		return nil, err
	}

	// We Decode the data because we already have an io.Reader on hand
	var res CloudBoltResource
	json.NewDecoder(resp.Body).Decode(&res)

	return &res, nil
}

// GetServer fetches a Server object from CloudBolt at the given path
// - Server Path (serverPath) e.g., "/api/v2/servers/123/"
func (c *CloudBoltClient) GetServer(serverPath string) (*CloudBoltServer, error) {
	apiurl := c.baseURL
	apiurl.Path = serverPath

	// log.Printf("[!!] apiurl in GetServer: %+v (%+v)", apiurl.String(), apiurl)

	req, err := http.NewRequest("GET", apiurl.String(), nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.makeRequest(req)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	// We Decode the data because we already have an io.Reader on hand
	var svr CloudBoltServer
	json.NewDecoder(resp.Body).Decode(&svr)

	return &svr, nil
}

// SubmitAction runs an action on the CloudBolt server
// - Action Path (actionPath) e.g., "/api/v2/actions/123/"
func (c *CloudBoltClient) SubmitAction(actionPath string) (*CloudBoltActionResult, error) {
	apiurl := c.baseURL
	apiurl.Path = actionPath

	// log.Printf("[!!] apiurl in SubmitAction: %+v (%+v)", apiurl.String(), apiurl)

	req, err := http.NewRequest("POST", apiurl.String(), nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.makeRequest(req)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	// We Decode the data because we already have an io.Reader on hand
	var actionRes CloudBoltActionResult
	json.NewDecoder(resp.Body).Decode(&actionRes)

	return &actionRes, nil
}

// DecomOrder orders the deletion of a CloudBolt resource
// - Group Path (grpPath) e.g., "/api/v2/groups/GRP-123/"
// - Environment Path (envPath) e.g., "/api/v2/environments/ENV-123/"
// - Servers (servers) e.g.,
//       []string{
//           `/api/v2/servers/123/`,
//           `/api/v2/servers/4567/`,
//           `/api/v2/servers/891011/`,
//       }
func (c *CloudBoltClient) DecomOrder(grpPath string, envPath string, servers []string) (*CloudBoltOrder, error) {
	decomItems := make([]map[string]interface{}, 0)

	decomItem := make(map[string]interface{})
	decomItem["environment"] = envPath
	decomItem["servers"] = servers

	decomItems = append(decomItems, decomItem)

	reqData := map[string]interface{}{
		"group": grpPath,
		"items": map[string]interface{}{
			"decom-items": decomItems,
		},
		"submit-now": "true",
	}

	reqJSON, err := json.Marshal(reqData)
	if err != nil {
		return nil, err
	}

	apiurl := c.baseURL
	apiurl.Path = c.apiEndpoint("orders")

	// log.Printf("[!!] apiurl in DecomOrder: %+v (%+v)", apiurl.String(), apiurl)

	req, err := http.NewRequest("POST", apiurl.String(), bytes.NewBuffer(reqJSON))
	if err != nil {
		return nil, err
	}

	resp, err := c.makeRequest(req)
	if err != nil {
		return nil, err
	}

	// We Decode the data because we already have an io.Reader on hand
	var order CloudBoltOrder
	json.NewDecoder(resp.Body).Decode(&order)

	return &order, nil
}
