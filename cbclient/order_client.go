package cbclient

import (
	"encoding/json"
	"log"
)

type CloudBoltOrder struct {
	Links struct {
		Self       CloudBoltHALItem   `json:"self"`
		Group      CloudBoltHALItem   `json:"group"`
		Owner      CloudBoltHALItem   `json:"owner"`
		ApprovedBy CloudBoltHALItem   `json:"approvedBy"`
		Jobs       []CloudBoltHALItem `json:"jobs"`
		Duplicate  CloudBoltHALItem   `json:"duplicate"`
	} `json:"_links"`
	Name            string `json:"name"`
	ID              string `json:"id"`
	Status          string `json:"status"`
	Rate            string `json:"rate"`
	CreateDate      string `json:"createDate"`
	ApproveDate     string `json:"approveDate"`
	DeploymentItems []struct {
		ID                 string                 `json:"id"`
		ResourceName       string                 `json:"resourceName"`
		ResourceParameters map[string]interface{} `json:"resourceParameters"`
		Blueprint          struct {
			Href  string `json:"href"`
			Title string `json:"title"`
		} `json:"blueprint"`
		BlueprintItemsArguments map[string]interface{} `json:"blueprintItemsArguments"`
		ItemType                string                 `json:"itemType"`
	} `json:"deploymentItems"`
}

// GetOrder fetches an Order from CloudBolt
// - Order ID (orderID) e.g., "123"; formatted into a string like "/api/v2/orders/123"
func (c *CloudBoltClient) GetOrder(orderID string) (*CloudBoltOrder, error) {
	apiurl := c.baseURL
	apiurl.Path = c.apiEndpoint("cmp", "orders", orderID)

	resp, err := c.makeRequest("GET", apiurl.String(), nil)
	if err != nil {
		log.Fatalln(err)

		return nil, err
	}

	// We Decode the data because we already have an io.Reader on hand
	var order CloudBoltOrder
	json.NewDecoder(resp.Body).Decode(&order)

	return &order, nil
}
