package cbclient

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type RenderTemplateResponse struct {
	Value string `json:"value,omitempty"`
}

type RenderTemplateRequest struct {
	Template           string                 `json:"template,omitempty"`
	TemplateProperties map[string]interface{} `json:"template_properties,omitempty"`
}

func (c *CloudBoltClient) RenderTemplate(template string, templateProperties map[string]interface{}) (*RenderTemplateResponse, error) {
	requestBody := RenderTemplateRequest{
		Template:           template,
		TemplateProperties: templateProperties,
	}

	reqJSON, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	apiurl := c.baseURL
	apiurl.Path = c.apiEndpoint("onefuse", "templateTester")

	resp, err := c.makeRequest("POST", apiurl.String(), reqJSON)
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
		var renderTemplateResponse RenderTemplateResponse
		json.NewDecoder(resp.Body).Decode(&renderTemplateResponse)

		return &renderTemplateResponse, nil
	}
}
