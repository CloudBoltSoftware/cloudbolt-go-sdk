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

// TODO: Improve the quality of this SDK code.

// CloudBoltObject ...
type CloudBoltObject struct {
	Links struct {
		Self struct {
			Href  string `json:"href"`
			Title string `json:"title"`
		} `json:"self"`
	} `json:"_links"`
	Name string `json:"name"`
	ID   string `json:"id"`
}

// CloudBoltClient ...
type CloudBoltClient struct {
	// TODO: Make members not public
	BaseURL    url.URL
	HTTPClient *http.Client
	Token      string
	username   string
	password   string
}

// CloudBoltResult ...
type CloudBoltResult struct {
	Links struct {
		Self struct {
			Href  string `json:"href"`
			Title string `json:"title"`
		} `json:"self"`
	} `json:"_links"`
	Total    int               `json:"total"`
	Count    int               `json:"count"`
	Embedded []CloudBoltObject `json:"_embedded"` // TODO: Maybe call this Items?
}

// CloudBoltActionResult ...
type CloudBoltActionResult struct {
	RunActionJob struct {
		Self struct {
			Href  string `json:"href"`
			Title string `json:"title"`
		} `json:"self"`
	} `json:"run-action-job"`
}

// CloudBoltHALItem ...
type CloudBoltHALItem struct {
	Href  string `json:"href"`
	Title string `json:"title"`
}

// CloudBoltOrder ...
type CloudBoltOrder struct {
	Links struct {
		Self struct {
			Href  string `json:"href"`
			Title string `json:"title"`
		} `json:"self"`
		Group struct {
			Href  string `json:"href"`
			Title string `json:"title"`
		} `json:"group"`
		Owner struct {
			Href  string `json:"href"`
			Title string `json:"title"`
		} `json:"owner"`
		ApprovedBy struct {
			Href  string `json:"href"`
			Title string `json:"title"`
		} `json:"approved-by"`
		Actions struct {
			Href  string `json:"href"`
			Title string `json:"title"`
		} `json:"actions"`
		Jobs []CloudBoltHALItem `json:"jobs"`
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

// CloudBoltJob ...
type CloudBoltJob struct {
	Links struct {
		Self struct {
			Href  string `json:"href"`
			Title string `json:"title"`
		} `json:"self"`
		Owner struct {
			Href  string `json:"href"`
			Title string `json:"title"`
		} `json:"owner"`
		Parent struct {
			Href  string `json:"href"`
			Title string `json:"title"`
		} `json:"parent"`
		Subjobs      []interface{} `json:"subjobs"`
		Prerequisite struct {
		} `json:"prerequisite"`
		DependentJobs []interface{} `json:"dependent-jobs"`
		Order         struct {
			Href  string `json:"href"`
			Title string `json:"title"`
		} `json:"order"`
		Resource struct {
			Href  string `json:"href"`
			Title string `json:"title"`
		} `json:"resource"`
		Servers []struct {
			Href  string `json:"href"`
			Title string `json:"title"`
		} `json:"servers"`
		LogUrls struct {
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

// CloudBoltGroup ...
type CloudBoltGroup struct {
	Links struct {
		Self struct {
			Href  string `json:"href"`
			Title string `json:"title"`
		} `json:"self"`
		Parent struct {
			Href  string `json:"href"`
			Title string `json:"title"`
		} `json:"parent"`
		Subgroups []struct {
			Href  string `json:"href"`
			Title string `json:"title"`
		} `json:"subgroups"`
		Environments          []interface{} `json:"environments"`
		OrderableEnvironments struct {
			Href  string `json:"href"`
			Title string `json:"title"`
		} `json:"orderable-environments"`
	} `json:"_links"`
	Name         string `json:"name"`
	ID           string `json:"id"`
	Type         string `json:"type"`
	Rate         string `json:"rate"`
	AutoApproval bool   `json:"auto-approval"`
}

// CloudBoltResource ...
type CloudBoltResource struct {
	Links struct {
		Self struct {
			Href  string `json:"href"`
			Title string `json:"title"`
		} `json:"self"`
		Blueprint struct {
			Href  string `json:"href"`
			Title string `json:"title"`
		} `json:"blueprint"`
		Owner struct {
			Href  string `json:"href"`
			Title string `json:"title"`
		} `json:"owner"`
		Group struct {
			Href  string `json:"href"`
			Title string `json:"title"`
		} `json:"group"`
		ResourceType struct {
			Href  string `json:"href"`
			Title string `json:"title"`
		} `json:"resource-type"`
		Servers []struct {
			Href  string `json:"href"`
			Title string `json:"title"`
			Tier  string `json:"tier"`
		} `json:"servers"`
		Actions []struct {
			Delete struct {
				Href  string `json:"href"`
				Title string `json:"title"`
			} `json:"Delete,omitempty"`
			Scale struct {
				Href  string `json:"href"`
				Title string `json:"title"`
			} `json:"Scale,omitempty"`
		} `json:"actions"`
		Jobs struct {
			Href  string `json:"href"`
			Title string `json:"title"`
		} `json:"jobs"`
		History struct {
			Href  string `json:"href"`
			Title string `json:"title"`
		} `json:"history"`
	} `json:"_links"`
	Name        string `json:"name"`
	ID          string `json:"id"`
	Status      string `json:"status"`
	InstallDate string `json:"install-date"`
}

// CloudBoltServer ...
type CloudBoltServer struct {
	Links struct {
		Self struct {
			Href  string `json:"href"`
			Title string `json:"title"`
		} `json:"self"`
		Owner struct {
			Href  string `json:"href"`
			Title string `json:"title"`
		} `json:"owner"`
		Group struct {
			Href  string `json:"href"`
			Title string `json:"title"`
		} `json:"group"`
		Environment struct {
			Href  string `json:"href"`
			Title string `json:"title"`
		} `json:"environment"`
		ResourceHandler struct {
			Href  string `json:"href"`
			Title string `json:"title"`
		} `json:"resource-handler"`
		Actions []struct {
			PowerOn struct {
				Href  string `json:"href"`
				Title string `json:"title"`
			} `json:"power_on,omitempty"`
			PowerOff struct {
				Href  string `json:"href"`
				Title string `json:"title"`
			} `json:"power_off,omitempty"`
			Reboot struct {
				Href  string `json:"href"`
				Title string `json:"title"`
			} `json:"reboot,omitempty"`
			RefreshInfo struct {
				Href  string `json:"href"`
				Title string `json:"title"`
			} `json:"refresh_info,omitempty"`
			Snapshot struct {
				Title string `json:"title"`
				Href  string `json:"href"`
			} `json:"snapshot,omitempty"`
			AdHocScript struct {
				Href  string `json:"href"`
				Title string `json:"title"`
			} `json:"Ad Hoc Script,omitempty"`
		} `json:"actions"`
		ProvisionJob struct {
			Href  string `json:"href"`
			Title string `json:"title"`
		} `json:"provision-job"`
		OsBuild struct {
			Href  string `json:"href"`
			Title string `json:"title"`
		} `json:"os-build"`
		Jobs struct {
			Href  string `json:"href"`
			Title string `json:"title"`
		} `json:"jobs"`
		History struct {
			Href  string `json:"href"`
			Title string `json:"title"`
		} `json:"history"`
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
	TechSpecificDetails struct {
		VmwareLinkedClone bool   `json:"vmware-linked-clone"`
		VmwareCluster     string `json:"vmware-cluster"`
	} `json:"tech-specific-details"`
}

// New ...
// TODO: In each other method try to do the action; if we get an auth error,
// try to get a new token and try the action again.
func New(protocol string, host string, port string, username string, password string) (CloudBoltClient, error) {
	cbClient := CloudBoltClient{
		// TODO: Make username and password data members of CloudBoltClient
		username: username,
		password: password,
		// TODO: Consider accepting HTTPClient as argument
		// Conditionally create new one?
		HTTPClient: &http.Client{
			Timeout: time.Second * 10,
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true, // TODO make this configurable
				},
			},
		},
		BaseURL: url.URL{
			Scheme: protocol,
			Host:   fmt.Sprintf("%s:%s", host, port),
		},
		// Token empty
	}

	// TODO: Conditional logging
	// log.Printf("[!!] cbClient: %+v", cbClient)

	return cbClient, nil
}

// performRequestWithAuthRetry ...
func (cbClient *CloudBoltClient) performRequestWithAuthRetry(req *http.Request) (*http.Response, error) {
	resp, err := cbClient.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == 403 { // use http.StatusAuthNotPermitted or whatever...
		// TODO: Turn this block into it's own function
		reqJSON, err := json.Marshal(struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}{
			Username: cbClient.username,
			Password: cbClient.password,
		})

		apiurl := cbClient.BaseURL
		// TODO: Make api-token-path a static string at the top of the file
		apiurl.Path = "/api/v2/api-token-auth/"

		// TODO: Conditional logging
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

		resp, err = cbClient.HTTPClient.Do(req)
	}

	return resp, err
}

// GetCloudBoltObject ...
// TODO: cbClient should by convention be `c`
// TODO: make this receiver a pointer
//       cbClient CloudBoltClient -> cbClient *CloudBoltClient
func (cbClient CloudBoltClient) GetCloudBoltObject(objPath string, objName string) (CloudBoltObject, error) {
	apiurl := cbClient.BaseURL
	// TODO: This is a magic string
	apiurl.Path = fmt.Sprintf("/api/v2/%s/", objPath)
	apiurl.RawQuery = fmt.Sprintf("filter=name:%s", url.QueryEscape(objName))

	// TODO: Conditional logging
	// log.Printf("[!!] apiurl in GetCloudBoltObject: %+v (%+v)", apiurl.String(), apiurl)

	req, err := http.NewRequest("GET", apiurl.String(), nil)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", cbClient.Token))
	req.Header.Add("Content-Type", "application/json")

	// TODO: Use this performRequestWithAuthRetry method
	// resp, err := performRequestWithAuthRetry(req)
	resp, err := cbClient.HTTPClient.Do(req)
	if err != nil {
		log.Fatalln(err)
		return CloudBoltObject{}, err
	}

	// TODO: Conditional logging
	// log.Printf("[!!] HTTP response: %+v", resp)

	// TODO: HTTP Response handling

	var res CloudBoltResult
	// TODO: Consider using json.Unmarshal
	// TODO: Un-chain the things
	// TODO: Handle all json.*.Decode() errors
	// TODO: Read resp.Body into a byte buffer
	// err := json.Unmarshal(bodyByteBuffer, &res)
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return CloudBoltObject{}, err
	}
	// TODO: Maybe defer resp.Body.Close()

	// TODO: Sanity check the decoded object

	// TODO: Conditional logging
	// log.Printf("[!!] CloudBoltResult response %+v", res) // HERE IS WHERE THE PANIC IS!!!
	if len(res.Embedded) == 0 {
		return CloudBoltObject{}, fmt.Errorf("Could not find %s with name %s. Does the user have permission to view this?", objPath, objName)
	}
	return res.Embedded[0], nil
}

// verifyGroup ...
// TODO: make this receiver a pointer
//       cbClient CloudBoltClient -> cbClient *CloudBoltClient
func (cbClient CloudBoltClient) verifyGroup(groupPath string, parentPath string) (bool, error) {
	var group CloudBoltGroup
	var parent string
	var nextParentPath string

	apiurl := cbClient.BaseURL
	apiurl.Path = groupPath

	// TODO: Conditional logging
	// log.Printf("[!!] apiurl in verifyGroup: %+v (%+v)", apiurl.String(), apiurl)

	req, err := http.NewRequest("GET", apiurl.String(), nil)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", cbClient.Token))
	req.Header.Add("Content-Type", "application/json")

	// TODO: Conditional logging
	// log.Printf("[!!] req: %+v", req)

	resp, err := cbClient.HTTPClient.Do(req)
	// TODO: Conditional logging
	// log.Printf("[!!] resp: %+v", resp)
	if err != nil {
		// TODO: Conditional logging
		// log.Printf("[!!] request err was not nil: %+v", err)
		log.Fatalln(err)

		return false, err
	}

	// TODO: This is a hack?
	if resp.StatusCode >= 300 {
		// TODO: Conditional logging
		// log.Printf("[!!] request returned a bad status: %+v", resp.Status)
		log.Fatalln(resp.Status)

		return false, errors.New(resp.Status)
	}

	json.NewDecoder(resp.Body).Decode(&group)

	// TODO: Conditional logging
	// log.Printf("[!!] group : %+v", group)

	nextIndex := strings.LastIndex(parentPath, "/")

	// TODO: Conditional logging
	// log.Printf("[!!] nextIndex : %+v", nextIndex)

	// TODO: Conditional logging
	// log.Printf("[!!] parentPath: %+v", parentPath)
	// TODO: Conditional logging
	// log.Printf("[!!] strings.LastIndex(parentPath, '/')+1: %+v", strings.LastIndex(parentPath, "/")+1)
	// TODO: hard to read, factor out into smaller functions
	// Might be a way to do this logic with a JSON decoder method
	if nextIndex >= 0 {
		parent = parentPath[strings.LastIndex(parentPath, "/")+1:]
		nextParentPath = parentPath[:strings.LastIndex(parentPath, "/")]
		// TODO: Conditional logging
		// log.Printf("[!!] parent: %+v, %+v", parent, nextParentPath)
	} else {
		parent = parentPath
		// TODO: Conditional logging
		// log.Printf("[!!] parent: %+v", parent)
	}

	// TODO: Conditional logging
	// log.Printf("[!!] group.Links.Parent.Title: %+v", group.Links.Parent.Title)
	if group.Links.Parent.Title != parent {
		return false, nil
	}

	// TODO: Conditional logging
	// log.Printf("[!!] nextParentPath: %+v", nextParentPath)
	if nextParentPath != "" {
		return cbClient.verifyGroup(group.Links.Parent.Href, nextParentPath)
	}

	return true, nil
}

// GetGroup ...
// TODO: make this receiver a pointer
//       cbClient CloudBoltClient -> cbClient *CloudBoltClient
func (cbClient CloudBoltClient) GetGroup(groupPath string) (CloudBoltObject, error) {
	var res CloudBoltResult
	var group string
	var parentPath string
	var groupFound bool

	groupPath = strings.Trim(groupPath, "/")
	nextIndex := strings.LastIndex(groupPath, "/")

	// TODO: Conditional logging
	// log.Printf("[!!] groupPath: %+v", groupPath)
	// TODO: Conditional logging
	// log.Printf("[!!] strings.LastIndex(groupPath, '/')+1: %+v", strings.LastIndex(groupPath, "/")+1)
	if nextIndex >= 0 {
		group = groupPath[strings.LastIndex(groupPath, "/")+1:]
		parentPath = groupPath[:strings.LastIndex(groupPath, "/")]
	} else {
		group = groupPath
	}

	apiurl := cbClient.BaseURL
	// TODO: This is a magic string
	apiurl.Path = "/api/v2/groups/"
	apiurl.RawQuery = fmt.Sprintf("filter=name:%s", url.QueryEscape(group))

	// TODO: Conditional logging
	// log.Printf("[!!] apiurl in GetGroup: %+v (%+v)", apiurl.String(), apiurl)

	req, err := http.NewRequest("GET", apiurl.String(), nil)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", cbClient.Token))
	req.Header.Add("Content-Type", "application/json")

	resp, err := cbClient.HTTPClient.Do(req)
	if err != nil {
		log.Fatalln(err)

		return CloudBoltObject{}, err
	}

	// TODO: Smarter decode this JSON body
	json.NewDecoder(resp.Body).Decode(&res)

	// TODO: Conditional logging
	// log.Printf("[!!] JSON decoded response: %+v", res)

	for _, v := range res.Embedded {
		groupFound, err = cbClient.verifyGroup(v.Links.Self.Href, parentPath)

		if groupFound {
			return v, nil
		}
	}

	return CloudBoltObject{}, fmt.Errorf("Group (%s): Not Found", groupPath)
}

// DeployBlueprint ...
// TODO: make this receiver a pointer
//       cbClient CloudBoltClient -> cbClient *CloudBoltClient
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

	// TODO: Conditional logging
	// log.Printf("[!!] JSON payload in POST request to Deploy Blueprint:\n%s", string(reqJSON))

	apiurl := cbClient.BaseURL
	// TODO: This is a magic string
	apiurl.Path = "/api/v2/orders/"

	// TODO: Conditional logging
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

// GetOrder ...
// TODO: make this receiver a pointer
//       cbClient CloudBoltClient -> cbClient *CloudBoltClient
func (cbClient CloudBoltClient) GetOrder(orderID string) (CloudBoltOrder, error) {
	var order CloudBoltOrder

	apiurl := cbClient.BaseURL
	// TODO: This is a magic string
	// API_ORDERS_FMT_STR
	apiurl.Path = fmt.Sprintf("/api/v2/orders/%s", orderID)

	// TODO: Conditional logging
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

// GetJob ...
// TODO: make this receiver a pointer
//       cbClient CloudBoltClient -> cbClient *CloudBoltClient
func (cbClient CloudBoltClient) GetJob(jobPath string) (CloudBoltJob, error) {
	var job CloudBoltJob

	apiurl := cbClient.BaseURL
	apiurl.Path = jobPath

	// TODO: Conditional logging
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

// GetResource ...
// TODO: make this receiver a pointer
//       cbClient CloudBoltClient -> cbClient *CloudBoltClient
func (cbClient CloudBoltClient) GetResource(resourcePath string) (CloudBoltResource, error) {
	var res CloudBoltResource

	apiurl := cbClient.BaseURL
	apiurl.Path = resourcePath

	// TODO: Conditional logging
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

// GetServer ...
// TODO: make this receiver a pointer
//       cbClient CloudBoltClient -> cbClient *CloudBoltClient
func (cbClient CloudBoltClient) GetServer(serverPath string) (CloudBoltServer, error) {
	var svr CloudBoltServer

	apiurl := cbClient.BaseURL
	apiurl.Path = serverPath

	// TODO: Conditional logging
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

// SubmitAction ...
// TODO: make this receiver a pointer
//       cbClient CloudBoltClient -> cbClient *CloudBoltClient
func (cbClient CloudBoltClient) SubmitAction(actionPath string) (CloudBoltActionResult, error) {
	var actionRes CloudBoltActionResult

	apiurl := cbClient.BaseURL
	apiurl.Path = actionPath

	// TODO: Conditional logging
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

// DecomOrder ...
// TODO: make this receiver a pointer
//       cbClient CloudBoltClient -> cbClient *CloudBoltClient
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
	// TODO: This is a magic string
	apiurl.Path = "/api/v2/orders/"

	// TODO: Conditional logging
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
