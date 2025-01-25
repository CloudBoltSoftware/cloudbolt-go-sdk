package cbclient

import (
	"encoding/json"
	"fmt"
	"log"
)

// SubmitAction runs an action on the CloudBolt resource or server
func (c *CloudBoltClient) SubmitAction(actionPath string, resourcePath string, parameters map[string]interface{}) (*CloudBoltRunActionResult, error) {
	apiurl := c.baseURL
	apiurl.Path = fmt.Sprintf("%srunAction/", actionPath)

	reqData := map[string]interface{}{
		"resource": resourcePath,
	}

	if parameters != nil {
		reqData["parameters"] = parameters
	}

	reqJSON, err := json.Marshal(reqData)
	if err != nil {
		return nil, err
	}

	log.Printf("[!!] JSON payload in POST request to Deploy Blueprint:\n%s", string(reqJSON))

	resp, err := c.makeRequest("POST", apiurl.String(), reqJSON)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	// We Decode the data because we already have an io.Reader on hand
	var actionRes CloudBoltRunActionResult
	json.NewDecoder(resp.Body).Decode(&actionRes)

	return &actionRes, nil
}
