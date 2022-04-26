package cbclient

import (
	"bytes"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"

	. "github.com/onsi/gomega"
)

type fn func(int) (string, int)

// The mockRequests data type; an array of http Requests.
type mockRequests []*http.Request

// Appends an httpRequest to a list of mockRequests
func (req *mockRequests) append(r *http.Request) {
	tmp := *req
	*req = append(tmp, r)
}

// We do this bodyBytes, ReadAll, NoCloser, etc mess so we can preserve
// the request body when appending it to the mockRequests array.
//
// Honestly, if you asked me why, I don't think I could tell you why it works.
//
// Source: https://medium.com/@xoen/2c6911805361
func copyRequest(r *http.Request) *http.Request {
	// Read it into the bytes buffer
	bodyBytes, err := ioutil.ReadAll(r.Body)
	Expect(err).NotTo(HaveOccurred())

	// Restore the io.ReadCloser to its original state
	r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes)) // Use the content

	return r
}

// Create a mockServer that responds to incoming events with a "responseFunc" script.
//
// Each request to the server is indexed (e.g., first request is 0, then 1, then 2)
// That index is used to call the script: responseFunc(requestIndex)
// The response func returns a tuple of (responseBody, responseStatusCode)
// The server then writes the body and sets the status code accordingly
//
// Returns the server and a queue of requests that have come in.
// The server is used to create a CloudBoltClient object.
// The requests can be indexed and inspected to verify the API client is making the correct calls.
// Requests usage looks like: Expect((*requests)[0].URL.Path).To(Equal("/path/to/a-resource/"))
//
// There is a known issue with this mockServer that it doesn't correctly empty the body of a request.
// This means that when, for example, the same http.Request object is used multiple times, the request.Body is not depleted.
// This differs with real behavior in that request bodies are depleted upon being read and are empty, so POST Request objects cannot be re-used.
// To preverse requests for testing introspection we can't easily deplete request object Bodies.
// This is probably solvable but I haven't had the time to look into it.
func mockServer(responseFunc fn) (*httptest.Server, *mockRequests) {
	var requests mockRequests

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Duplicate this request and add it to the buffer for later inspection
		requests.append(copyRequest(r))

		// Get the body (string) and HTTP status code for this request
		body, status := responseFunc(len(requests) - 1)

		// Write the response
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(status)
		w.Write([]byte(body))
	}))

	// Return mock server with scripted responses
	// Return requests buffer
	return server, &requests
}

// bodyToString was created because I kept forgetting how
// to get something useful out of http.Response.Body
func bodyToString(b io.ReadCloser) string {
	bodyBytes, err := ioutil.ReadAll(b)
	Expect(err).NotTo(HaveOccurred())

	bodyString := string(bodyBytes)

	return bodyString
}

// getClient takes an httptest.Server and returns a pointer to a CloudBolt Client object
// Uses defaults when possbile, e.g., the HTTP Client default.
func getClient(server *httptest.Server) *CloudBoltClient {
	protocol := "http"
	uri, err := url.Parse(server.URL)
	Expect(err).NotTo(HaveOccurred())

	host, port, err := net.SplitHostPort(uri.Host)
	Expect(err).NotTo(HaveOccurred())

	username := "testUser"
	password := "testPass"
	domain := "mydomain.com"

	client := New(protocol, host, port, username, password, domain, nil)
	Expect(client).NotTo(BeNil())

	return client
}
