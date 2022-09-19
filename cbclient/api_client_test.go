package cbclient

import (
	"net/http"
	"net/url"
	"testing"
	"time"

	. "github.com/onsi/gomega"
)

// Below is an example of how to write these tests, it includes
// reading from a response body
/*
func TestHttpTestExample(t *testing.T) {
	// Register this test with gomega
	RegisterTesting(t)

	// A pointer to the httptest Server and a pointer to a slice of Requests
	server, requests := mockServer(responseBodyFunction)

	// The CloudBolt Client object
	client := getClient(server)

	// Prepare what you will pass to the CloudBolt Client
	params := someCloudBoltData

	// Perform some action(s) with the CloudBolt Client
	// Receive an object in response that we can inspect
	obj, err := client.SomeCloudBoltAction(params)
	Expect(err).NotTo(HaveOccurred())

	// Make assertions about the response object
	Expect(obj.Some.Field).To(Equal("SomeValue"))
	// Expect(...).To(...)
}
*/

func TestNew(t *testing.T) {
	// Register the test with gomega
	RegisterTestingT(t)

	// Create a mock server with that accepts requests and responds with scripted responses
	// Create a buffer of received requests
	server, requests := mockServer(responsesForNew)
	Expect(server).NotTo(BeNil())
	Expect(requests).NotTo(BeNil())

	// Call our getClient function which also makes assertions about the process
	client := getClient(server)
	Expect(client).NotTo(BeNil())

	// calling New() (which happened in getClient()) should make no API calls
	Expect(len(*requests)).To(Equal(0))

	serverURL, err := url.Parse(server.URL)
	Expect(err).NotTo(HaveOccurred())

	Expect(client.token).To(BeEmpty())
	Expect(client.username).To(Equal("testUser"))
	Expect(client.baseURL).To(Equal(*serverURL))
	Expect(client.password).To(Equal("testPass"))
	Expect(client.domain).To(Equal("mydomain.com"))
	Expect(client.httpClient.Timeout).To(Equal(time.Duration(60 * time.Second)))
}

func TestAuthenticate(t *testing.T) {
	// Register the test with gomega
	RegisterTestingT(t)

	// Create a mock server to accept requests and respond with scripted responses
	// Create a buffer of requests
	server, requests := mockServer(responsesForAuthenticate)
	Expect(server).NotTo(BeNil())
	Expect(requests).NotTo(BeNil())

	// Initialize an CloudBolt API client
	client := getClient(server)
	Expect(client).NotTo(BeNil())

	// At this point we have made no API requests
	Expect(len(*requests)).To(Equal(0))

	// Manually authenticate with the server
	status, err := client.Authenticate()
	Expect(status).Should(Equal(200)) // Questionable
	Expect(err).NotTo(HaveOccurred())

	// We only expect the API client to have made 1 requests at this point
	Expect(len(*requests)).To(Equal(1))
	Expect((*requests)[0].URL.Path).To(Equal("/api/v3/cmp/apiToken/"))
	Expect((*requests)[0].Method).To(Equal("POST"))

	// The token we should have at this point is "Testing Token"
	Expect(client.token).To(Equal("Testing Token"))
}

// TestAuthWrapper is a fun test that verifies that when CloudBoltClient.token
// is empty, cbClient auto-re-auths with CloudBolt.
//
// The timeline look like this:
// 1. cbClient (bad token)  (Request for X)----> CloudBolt Server | Client makes request for X resource
// 2. cbClient (bad token)  <----(Unauthorized!) CloudBolt Server | Client receives "Unauthorized" response
// 3. cbClient (bad token)  (Request Token ----> CloudBolt Server | Client requests a new token
// 4. cbClient (bad token)  <-------(User Token) CloudBolt Server | Client receives an auth token
// 5. cbClient (with token) (Request for X)----> CloudBolt Server | Client re-requests X resource
// 6. cbClient (with token) <--------(X Payload) CloudBolt Server | Client successfully receives X resource
//
// In the above timeline, 'bad token' may be an empty or expired token.
func TestAuthWrappedRequest(t *testing.T) {
	// Register the test with gomega
	RegisterTestingT(t)

	// Setup mock server
	// Setup requests buffer
	server, requests := mockServer(responsesForAuthWrappedRequest)
	Expect(server).NotTo(BeNil())
	Expect(requests).NotTo(BeNil())

	// Setup CloudBolt Client
	client := getClient(server)
	Expect(client).NotTo(BeNil())

	// To call authWrapperRequest directly we need an http.Request
	// Use the mock server's URL
	apiurl := client.baseURL
	// Use a dummy path
	apiurl.Path = "/foo/"
	reqBody := []byte(`{"bar": "foo"}`)

	// Call authWrappedRequest implicitly via makeRequest
	resp, err := client.makeRequest("POST", apiurl.String(), reqBody)
	Expect(err).NotTo(HaveOccurred())
	Expect(resp).NotTo(BeNil())

	// This should have made three requests:
	// 1+2. Fail to get resource, get a token
	// 3. Successfully getting the object
	Expect(len(*requests)).To(Equal(3))

	// First request is for the thing
	Expect((*requests)[0].URL.Path).To(Equal("/foo/"))
	Expect((*requests)[0].Header["Authorization"]).To(Equal([]string{"Bearer"}))

	// Second request is for an API token
	Expect((*requests)[1].URL.Path).To(Equal("/api/v3/cmp/apiToken/"))
	Expect((*requests)[1].Method).To(Equal("POST"))

	// Third request is for the thing again with an API token
	Expect((*requests)[2].URL.Path).To(Equal("/foo/"))
	Expect((*requests)[2].Header["Authorization"]).To(Equal([]string{"Bearer Testing Token"}))

	// Check that the final response body was as expected
	respBody := bodyToString(resp.Body)
	Expect(respBody).To(MatchJSON(`{"foo": "bar"}`))
}

// This validates that when the server responds with a "good" http status,
// the auth wrapper does not make a request for a new token.
func TestAuthWrappedRequestWithValidToken(t *testing.T) {
	// Register the test with gomega
	RegisterTestingT(t)

	// Create server, requests, and client.
	server, requests := mockServer(responsesForAuthWrappedRequestWithToken)
	Expect(server).NotTo(BeNil())
	Expect(requests).NotTo(BeNil())

	// Get an API client
	client := getClient(server)
	Expect(client).NotTo(BeNil())

	// We should have made no API requests at this point
	Expect(len(*requests)).To(Equal(0))

	// To call authWrapper we need to create an http.Request object
	// Use the mock server base URL
	apiurl := client.baseURL
	// Set a dummy path
	apiurl.Path = "/foo/"
	// Create the HTTP request object
	req, err := http.NewRequest("GET", apiurl.String(), nil)
	Expect(err).NotTo(HaveOccurred())
	// Create second HTTP request object if the first fails
	reqBackup, err := http.NewRequest("GET", apiurl.String(), nil)
	Expect(err).NotTo(HaveOccurred())

	// call the request auth wrapper
	resp, err := client.authWrappedRequest(req, reqBackup)
	Expect(err).NotTo(HaveOccurred())
	Expect(resp).NotTo(BeNil())

	// The wrapper got a 200 response the first time so it should have made a total of 1 requests
	Expect(len(*requests)).To(Equal(1))

	// The body needs to be parsed and checked
	body := bodyToString(resp.Body)
	Expect(body).To(MatchJSON(`{"foo": "bar"}`))
}

func TestAPIEndpoint(t *testing.T) {
	RegisterTestingT(t)

	// Create an API client with a strange API version
	client := New("https", "mycloudbolt.test", "1234", "uname", "bar", "", nil)

	// When nothing is passed, this is the the base API path
	basePath := client.apiEndpoint()
	Expect(basePath).To(Equal("/api/v3/"))

	// When many things are passed it formats them as a path in order
	longEndpoint := client.apiEndpoint("cmp", "a", "b", "c")
	Expect(longEndpoint).To(Equal("/api/v3/cmp/a/b/c/"))
}
