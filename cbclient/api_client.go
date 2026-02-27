package cbclient

/*
Organization note:
1. Imports
2. Constants
3. Public types
4. Private types
6. Public funcs
7. Private funcs
*/

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"time"
)

const filterByName string = "filter=name:%s"

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
	domain     string
}

// CloudBoltResult stores the response of paginated calls like `/api/v2/blueprints/`
// These include a link to the page and an `embedded` list of response objects.
type CloudBoltResult struct {
	Links struct {
		Self CloudBoltHALItem `json:"self"`
	} `json:"_links"`
	Total int `json:"total"`
	Count int `json:"count"`
}

// CloudBoltHALItem stores an object's title and API endpoint.
// This is a common pattern in the CloudBolt API, so it gets used a lot.
type CloudBoltHALItem struct {
	Href  string `json:"href"`
	Title string `json:"title"`
}

type CloudBoltReferenceFields struct {
	Links struct {
		Self struct {
			Href  string `json:"href"`
			Title string `json:"title"`
		} `json:"self"`
	} `json:"_links"`
	Name string `json:"name"`
	ID   string `json:"id"`
}

type CloudBoltRunActionResult struct {
	ResourceHref string `json:"resource"`
	Results      struct {
		Job           CloudBoltJob   `json:"job"`
		Order         CloudBoltOrder `json:"order"`
		Status        string         `json:"status"`
		OutputMessage string         `json:"outputMessage"`
		ErrorMessage  string         `json:"errorMessage"`
	} `json:"results"`
}

var ErrNotFound = errors.New("CloudBolt Object Not Found")

// New returns an initialized CloudBoltClient object.
// Accepts as input:
// - HTTP Protocol (protocol) e.g., "https"
// - HTTP Host (host) e.g., "cloudbolt.intranet"
// - HTTP Port (port) e.g., "443"
// - Username (username) e.g., "myUserName"
// - Password (password) e.g., "My Passphrase!"
// - Domain (domain) e.g., "mydomain.com"
// - User-provided *HTTPClient (httpClient); provide `nil` to get a server with the following defaults:
//   - Timeout set to 60 seconds
//   Provide a custom http.Client if you require unique certificate, timeout, etc., configured.
//
// New does not make any API calls.
// CloudBoltClient.Authenticate must be called to initialize CloudBoltClient.token.
// This is done automatically when a request receives an HTTP Authorization error.
func New(protocol string, host string, port string, username string, password string, domain string, httpClient *http.Client) *CloudBoltClient {
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
		domain:     domain,
	}
}

// Authenticate forces the CloudBoltClient to re-authenticate
// Returns an error if there is an HTTP error, or if the HTTP Status Code is >=400
func (c *CloudBoltClient) Authenticate() (int, error) {
	// Craft the JSON payload used to request an API token
	var reqJSON []byte
	var err error

	if c.domain != "" {
		userCreds := struct {
			Username string `json:"username"`
			Password string `json:"password"`
			Domain   string `json:"domain"`
		}{
			Username: c.username,
			Password: c.password,
			Domain:   c.domain,
		}

		reqJSON, err = json.Marshal(userCreds)
		if err != nil {
			return -1, err
		}
	} else {
		userCreds := struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}{
			Username: c.username,
			Password: c.password,
		}

		reqJSON, err = json.Marshal(userCreds)
		if err != nil {
			return -1, err
		}
	}

	// Put the username+password into a bytes buffer for the API request
	reqJSONBuffer := bytes.NewBuffer(reqJSON)

	// Craft the URL api-token request endpoint based on the API version
	apiurl := c.baseURL
	apiurl.Path = c.apiEndpoint("cmp", "apiToken")

	// Make the POST request to get the API token
	req, err := http.NewRequest("POST", apiurl.String(), reqJSONBuffer)
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
	var userAuthData struct {
		Token string `json:"token"`
	}

	json.NewDecoder(resp.Body).Decode(&userAuthData)

	// Set the CloudBoltClient token as that parsed Token value
	c.token = userAuthData.Token

	// Return the HTTP status code and a nil error for success
	return resp.StatusCode, nil
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
		"v3",
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

// makeRequest wraps the normal HTTP request by re-authenticating if we get an
// "Unauthorized" HTTP response.
//
// if the first attempt at the request returns a 401 or 403 HTTP Status Code
// it Attempts exactly one call to CloudBoltClient.Authenticate() and resets the request token.
func (c *CloudBoltClient) authWrappedRequest(req *http.Request, backup *http.Request) (*http.Response, error) {
	// if c.token == "" {
	// 	log.Printf("Authenticating %+v", req)
	// 	time.Sleep(10 * time.Second)
	// 	_, err := c.Authenticate()
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// }

	// Add the Auth token to the request
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))

	// Attempt to make the given HTTP request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	// (Bluntly) Handles common HTTP "auth" related Status Codes
	if resp.StatusCode >= 400 {
		_, err := c.Authenticate()
		if err != nil {
			return nil, err
		}

		backup.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))

		resp, err := c.httpClient.Do(backup)
		if err != nil {
			return nil, err
		}

		return resp, nil
	}

	// If we didn't get one of the those 50x/40x Status Codes,
	// then pass through the original result
	// Return the original response
	return resp, nil
}

// makeRequest wraps what http.NewRequest would do:
// Creates an HTTP request
// Creates a duplicate if the body is not nil
// Calls authWrappedRequest with both requests
func (c *CloudBoltClient) makeRequest(method string, url string, body []byte) (*http.Response, error) {
	// Construct the initial request
	req, err := constructRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	// Default to assigning reqBackup to the current request
	reqBackup := req

	// If the Body is not nil, we cannot reuse the request object
	// So we generate a new request object from scratch
	if body != nil {
		reqBackup, err = constructRequest(method, url, body)
		if err != nil {
			return nil, err
		}
	}

	return c.authWrappedRequest(req, reqBackup)
}

// constructRequest generates a CloudBolt API HTTP request object.
// - Reads the body into a buffer.
// - Calls http.NewRequest
// - Sets ContentType and Accept to JSON
func constructRequest(method string, url string, body []byte) (*http.Request, error) {
	// Load the body into a buffer
	reqBody := bytes.NewBuffer(body)

	// Generate a new HTTP Request
	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		return nil, err
	}

	// Set JSON headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	return req, nil
}

func checkHttpStatus(resp *http.Response) error {
	switch {
	case resp.StatusCode >= 500:
		buf := new(bytes.Buffer)
		buf.ReadFrom(resp.Body)
		respBody := buf.String()
		return fmt.Errorf("received a server error: %s", respBody)
	case resp.StatusCode >= 400:
		buf := new(bytes.Buffer)
		buf.ReadFrom(resp.Body)
		respBody := buf.String()
		return fmt.Errorf("received an HTTP client error: %s", respBody)
	}

	return nil
}

func checkOneFuseResponse(resp *http.Response) (*OneFuseJobStatus, error) {
	err := checkHttpStatus(resp)
	if err != nil {
		return nil, err
	}

	// We Decode the data because we already have an io.Reader on hand
	var jobStatus OneFuseJobStatus
	json.NewDecoder(resp.Body).Decode(&jobStatus)

	return &jobStatus, nil
}
