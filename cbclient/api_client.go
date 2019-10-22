package cbclient

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

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
//
// Note that the CloudBoltClient will change in future version of the CloudBolt SDK:
// - All members will be private.
// - It will have the following members:
//   - apiVersion
//   - username
//   - password
//   - token
//   - httpClient
// For more information about changes to the CloudBoltClient struct, see the docstring for New.
type CloudBoltClient struct {
	BaseURL    url.URL
	HTTPClient *http.Client
	Token      string
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

// New creates a CloudBoltClient object.
// Note that this _does_ make a call to the API to retrieve a token.
//
// This behavior is expected to change in future version of the SDK.
// - New will accept as input an *http.Client, username, and password, and apiVersion.
// - New will make no API calls and will initialize an empty `token`.
// - Each cbClient function will make a request with the current `CloudBoltClient.token`.
//   If that call fails, it will re-auth and try again.
// - cbClient will get a new `Authenticate` method to force this re-auth.
func New(protocol string, host string, port string, username string, password string) (CloudBoltClient, error) {
	var cbClient CloudBoltClient
	cbClient.HTTPClient = &http.Client{
		Timeout: time.Second * 10,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true, // TODO make this configurable
			},
		},
	}

	cbClient.BaseURL = url.URL{
		Scheme: protocol,
		Host:   fmt.Sprintf("%s:%s", host, port),
	}

	reqJSON, err := json.Marshal(struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{
		Username: username,
		Password: password,
	})

	if err != nil {
		log.Fatalln(err)
		return cbClient, err
	}

	apiurl := cbClient.BaseURL
	apiurl.Path = "/api/v2/api-token-auth/"

	// log.Printf("[!!] apiurl in New: %+v (%+v)", apiurl.String(), apiurl)

	resp, err := cbClient.HTTPClient.Post(apiurl.String(), "application/json", bytes.NewBuffer(reqJSON))
	if err != nil {
		log.Fatalf("Failed to create the API client. %s", err)
	}

	userAuthData := struct {
		Token string `json:"token"`
	}{}

	json.NewDecoder(resp.Body).Decode(&userAuthData)
	cbClient.Token = userAuthData.Token

	// log.Printf("[!!] cbClient: %+v", cbClient)

	return cbClient, nil
}

// GetCloudBoltObject fetches a given object of type "objPath" with the name "objName"
// e.g., GetCloudBoltObject("users", "Susan") gets the user with username "Susan"
func (cbClient CloudBoltClient) GetCloudBoltObject(objPath string, objName string) (CloudBoltObject, error) {
	apiurl := cbClient.BaseURL
	apiurl.Path = fmt.Sprintf("/api/v2/%s/", objPath)
	apiurl.RawQuery = fmt.Sprintf("filter=name:%s", url.QueryEscape(objName))

	// log.Printf("[!!] apiurl in GetCloudBoltObject: %+v (%+v)", apiurl.String(), apiurl)

	req, err := http.NewRequest("GET", apiurl.String(), nil)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", cbClient.Token))
	req.Header.Add("Content-Type", "application/json")

	resp, err := cbClient.HTTPClient.Do(req)
	if err != nil {
		log.Fatalln(err)

		return CloudBoltObject{}, err // Consider return nil, err
	}
	// log.Printf("[!!] HTTP response: %+v", resp)

	// TODO: HTTP Response handling

	var res CloudBoltResult
	json.NewDecoder(resp.Body).Decode(&res)

	// TODO: Sanity check the decoded object

	// log.Printf("[!!] CloudBoltResult response %+v", res) // HERE IS WHERE THE PANIC IS!!!
	if len(res.Embedded) == 0 {
		return CloudBoltObject{}, fmt.Errorf("Could not find %s with name %s. Does the user have permission to view this?", objPath, objName)
	}
	return res.Embedded[0], nil
}

// verifyGroup checks that all a given group is the one we intended to fetch.
//
// groupPath is the API path to the group, e.g., "/api/v2/groups/GRP-123456"
//
// If a group has no parents, "parentPath" should be empty.
// If a group has parents, it should be of the format "root-level-parent/sub-parent/.../closest-parent"
func (cbClient CloudBoltClient) verifyGroup(groupPath string, parentPath string) (bool, error) {
	// log.Printf("Verifying group %+v with parent(s) %+v\n", groupPath, parentPath)
	var group CloudBoltGroup
	var parent string
	var nextParentPath string

	apiurl := cbClient.BaseURL
	apiurl.Path = groupPath

	req, err := http.NewRequest("GET", apiurl.String(), nil)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", cbClient.Token))
	req.Header.Add("Content-Type", "application/json")

	resp, err := cbClient.HTTPClient.Do(req)
	if err != nil {
		log.Fatalln(err)

		return false, err
	}
	if resp.StatusCode >= 300 {
		log.Fatalln(resp.Status)

		return false, errors.New(resp.Status)
	}

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
		return cbClient.verifyGroup(group.Links.Parent.Href, nextParentPath)
	}

	// log.Printf("[!!] Group verified, returning true\n")
	return true, nil
}

// GetGroup accepts a groupPath string parameter of the following format:
// "/my parent group/some subgroup/a child group/" or just "my parent group"
//
// verifyGroup recursively verifies that this is a valid group/subgroup.
func (cbClient CloudBoltClient) GetGroup(groupPath string) (CloudBoltObject, error) {
	var res CloudBoltResult
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

	apiurl := cbClient.BaseURL
	apiurl.Path = "/api/v2/groups/"
	apiurl.RawQuery = fmt.Sprintf("filter=name:%s", url.QueryEscape(group))

	// log.Printf("[!!] apiurl in GetGroup: %+v (%+v)", apiurl.String(), apiurl)

	req, err := http.NewRequest("GET", apiurl.String(), nil)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", cbClient.Token))
	req.Header.Add("Content-Type", "application/json")

	resp, err := cbClient.HTTPClient.Do(req)
	if err != nil {
		log.Fatalln(err)

		return CloudBoltObject{}, err
	}

	json.NewDecoder(resp.Body).Decode(&res)

	for _, v := range res.Embedded {
		// log.Printf("Group is %+v\n", v)
		groupFound, err = cbClient.verifyGroup(v.Links.Self.Href, parentPath)

		if groupFound {
			return v, nil
		}
	}

	return CloudBoltObject{}, fmt.Errorf("Group (%s): Not Found", groupPath)
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
func (cbClient CloudBoltClient) DeployBlueprint(grpPath string, bpPath string, resourceName string, bpItems []map[string]interface{}) (CloudBoltOrder, error) {
	var order CloudBoltOrder

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
		log.Fatalln(err)
		return order, err
	}

	// log.Printf("[!!] JSON payload in POST request to Deploy Blueprint:\n%s", string(reqJSON))

	apiurl := cbClient.BaseURL
	apiurl.Path = "/api/v2/orders/"

	// log.Printf("[!!] apiurl in DeployBlueprint: %+v (%+v)", apiurl.String(), apiurl)

	reqBody := bytes.NewBuffer(reqJSON)

	req, err := http.NewRequest("POST", apiurl.String(), reqBody)
	if err != nil {
		log.Fatalln(err)
		return order, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", cbClient.Token))
	req.Header.Add("Content-Type", "application/json")
	// TODO: Make the API more responsive
	cbClient.HTTPClient.Timeout = 60 * time.Second

	resp, err := cbClient.HTTPClient.Do(req)
	if err != nil {
		log.Fatalln(err)
		return CloudBoltOrder{}, err
	}

	// Handle some common HTTP errors
	switch {
	case resp.StatusCode >= 500:
		buf := new(bytes.Buffer)
		buf.ReadFrom(resp.Body)
		respBody := buf.String()
		return CloudBoltOrder{}, fmt.Errorf("received a server error: %s", respBody)
	case resp.StatusCode >= 400:
		buf := new(bytes.Buffer)
		buf.ReadFrom(resp.Body)
		respBody := buf.String()
		return CloudBoltOrder{}, fmt.Errorf("received an HTTP client error: %s", respBody)
	default:
		json.NewDecoder(resp.Body).Decode(&order)

		return order, nil
	}
}

// GetOrder fetches an Order from CloudBolt
// - Order ID (orderID) e.g., "123"; formatted into a string like "/api/v2/orders/123"
func (cbClient CloudBoltClient) GetOrder(orderID string) (CloudBoltOrder, error) {
	var order CloudBoltOrder

	apiurl := cbClient.BaseURL
	apiurl.Path = fmt.Sprintf("/api/v2/orders/%s", orderID)

	// log.Printf("[!!] apiurl in GetOrder: %+v (%+v)", apiurl.String(), apiurl)

	req, err := http.NewRequest("GET", apiurl.String(), nil)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", cbClient.Token))
	req.Header.Add("Content-Type", "application/json")

	resp, err := cbClient.HTTPClient.Do(req)
	if err != nil {
		log.Fatalln(err)

		return CloudBoltOrder{}, err
	}

	json.NewDecoder(resp.Body).Decode(&order)

	return order, nil
}

// GetJob fetches the Job object from CloudBolt at the given path
// - Job Path (jobPath) e.g., "/api/v2/jobs/123/"
func (cbClient CloudBoltClient) GetJob(jobPath string) (CloudBoltJob, error) {
	var job CloudBoltJob

	apiurl := cbClient.BaseURL
	apiurl.Path = jobPath

	// log.Printf("[!!] GetJob: %+v (%+v)", apiurl.String(), apiurl)

	req, err := http.NewRequest("GET", apiurl.String(), nil)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", cbClient.Token))
	req.Header.Add("Content-Type", "application/json")

	resp, err := cbClient.HTTPClient.Do(req)
	if err != nil {
		log.Fatalln(err)

		return CloudBoltJob{}, err
	}

	json.NewDecoder(resp.Body).Decode(&job)

	return job, nil
}

// GetResource fetches a Resource object from CloudBolt at the given path
// - Resource Path (resourcePath) e.g., "/api/v2/resources/service/123/"
func (cbClient CloudBoltClient) GetResource(resourcePath string) (CloudBoltResource, error) {
	var res CloudBoltResource

	apiurl := cbClient.BaseURL
	apiurl.Path = resourcePath

	// log.Printf("[!!] apiurl in GetResource: %+v (%+v)", apiurl.String(), apiurl)

	req, err := http.NewRequest("GET", apiurl.String(), nil)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", cbClient.Token))
	req.Header.Add("Content-Type", "application/json")

	resp, err := cbClient.HTTPClient.Do(req)
	if err != nil {
		log.Fatalln(err)

		return CloudBoltResource{}, err
	}

	json.NewDecoder(resp.Body).Decode(&res)

	return res, nil
}

// GetServer fetches a Server object from CloudBolt at the given path
// - Server Path (serverPath) e.g., "/api/v2/servers/123/"
func (cbClient CloudBoltClient) GetServer(serverPath string) (CloudBoltServer, error) {
	var svr CloudBoltServer

	apiurl := cbClient.BaseURL
	apiurl.Path = serverPath

	// log.Printf("[!!] apiurl in GetServer: %+v (%+v)", apiurl.String(), apiurl)

	req, err := http.NewRequest("GET", apiurl.String(), nil)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", cbClient.Token))
	req.Header.Add("Content-Type", "application/json")

	resp, err := cbClient.HTTPClient.Do(req)
	if err != nil {
		log.Fatalln(err)
		return CloudBoltServer{}, err
	}

	json.NewDecoder(resp.Body).Decode(&svr)

	return svr, nil
}

// SubmitAction runs an action on the CloudBolt server
// - Action Path (actionPath) e.g., "/api/v2/actions/123/"
func (cbClient CloudBoltClient) SubmitAction(actionPath string) (CloudBoltActionResult, error) {
	var actionRes CloudBoltActionResult

	apiurl := cbClient.BaseURL
	apiurl.Path = actionPath

	// log.Printf("[!!] apiurl in SubmitAction: %+v (%+v)", apiurl.String(), apiurl)

	req, err := http.NewRequest("POST", apiurl.String(), nil)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", cbClient.Token))
	req.Header.Add("Content-Type", "application/json")

	resp, err := cbClient.HTTPClient.Do(req)
	if err != nil {
		log.Fatalln(err)
		return CloudBoltActionResult{}, err
	}

	json.NewDecoder(resp.Body).Decode(&actionRes)

	return actionRes, nil
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
func (cbClient CloudBoltClient) DecomOrder(grpPath string, envPath string, servers []string) (CloudBoltOrder, error) {
	var order CloudBoltOrder

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
		log.Fatalln(err)
		return CloudBoltOrder{}, err
	}

	apiurl := cbClient.BaseURL
	apiurl.Path = "/api/v2/orders/"

	// log.Printf("[!!] apiurl in DecomOrder: %+v (%+v)", apiurl.String(), apiurl)

	req, err := http.NewRequest("POST", apiurl.String(), bytes.NewBuffer(reqJSON))
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", cbClient.Token))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	resp, err := cbClient.HTTPClient.Do(req)
	if err != nil {
		log.Fatalln(err)
		return CloudBoltOrder{}, err
	}

	json.NewDecoder(resp.Body).Decode(&order)

	return order, nil
}
