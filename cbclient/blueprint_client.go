package cbclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
)

type CloudBoltBlueprintResult struct {
	CloudBoltResult
	Embedded struct {
		Blueprints []CloudBoltReferenceFields `json:"blueprints"`
	} `json:"_embedded"`
}

// GetBlueprint accepts the name of a Blueprint
//
func (c *CloudBoltClient) GetBlueprint(name string) (*CloudBoltReferenceFields, error) {
	apiurl := c.baseURL
	apiurl.Path = c.apiEndpoint("blueprints")
	apiurl.RawQuery = fmt.Sprintf(filterByName, url.QueryEscape(name))

	resp, err := c.makeRequest("GET", apiurl.String(), nil)
	if err != nil {
		return nil, err
	}

	// TODO: HTTP Response handling

	// We Decode the data because we already have an io.Reader on hand
	var res CloudBoltBlueprintResult
	json.NewDecoder(resp.Body).Decode(&res)

	// TODO: Sanity check the decoded object

	// log.Printf("[!!] CloudBoltResult response %+v", res) // HERE IS WHERE THE PANIC IS!!!
	if len(res.Embedded.Blueprints) == 0 {
		return nil, fmt.Errorf(
			"Could not find blueprint with name %s. Does the user have permission to view this?",
			name,
		)
	}
	return &res.Embedded.Blueprints[0], nil
}

func (c *CloudBoltClient) GetBlueprintById(id string) (*CloudBoltReferenceFields, error) {
	apiurl := c.baseURL
	apiurl.Path = c.apiEndpoint("blueprints", id)

	resp, err := c.makeRequest("GET", apiurl.String(), nil)
	if err != nil {
		return nil, err
	}

	// TODO: HTTP Response handling

	// We Decode the data because we already have an io.Reader on hand
	var res CloudBoltReferenceFields
	json.NewDecoder(resp.Body).Decode(&res)

	// TODO: Sanity check the decoded object

	return &res, nil
}

func (c *CloudBoltClient) DeployBlueprint(grpPath string, blueprintID string, resourceName string, bpParams map[string]interface{}, bpItems []map[string]interface{}) (*CloudBoltOrder, error) {
	deployItems := make(map[string]interface{})

	for _, v := range bpItems {
		deployItems[v["bp-item-name"].(string)] = map[string]interface{}{
			"parameters": v["bp-item-paramas"].(map[string]interface{}),
		}

		env, ok := v["environment"]
		if ok {
			deployItems[v["bp-item-name"].(string)].(map[string]interface{})["environment"] = env
		}

		osb, ok := v["osbuild"]
		if ok {
			deployItems[v["bp-item-name"].(string)].(map[string]interface{})["osBuild"] = osb
		}
	}

	reqData := map[string]interface{}{
		"group":           grpPath,
		"deploymentItems": deployItems,
	}

	if bpParams != nil {
		reqData["parameters"] = bpParams
	}

	reqJSON, err := json.Marshal(reqData)
	if err != nil {
		return nil, err
	}

	apiurl := c.baseURL
	apiurl.Path = c.apiEndpoint("blueprints", blueprintID, "deploy")
	log.Printf("%s", apiurl.String())

	// log.Printf("[!!] apiurl in DeployBlueprint: %+v (%+v)", apiurl.String(), apiurl)

	resp, err := c.makeRequest("POST", apiurl.String(), reqJSON)
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
