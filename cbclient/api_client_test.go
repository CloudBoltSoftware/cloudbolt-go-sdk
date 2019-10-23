package cbclient

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"net"

	. "github.com/onsi/gomega"
)

/*
// here is an example of how to run these tests, it includes
// reading from a response body

func TestHttpTestExample(t *testing.T) {
	RegisterTestingT(t)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write([]byte(`{"Hello": "This is dog"}`))
	}))
	defer server.Close()

	req, err := http.NewRequest("GET", server.URL, nil)
	Expect(err).NotTo(HaveOccurred())

	client := &http.Client{}

	resp, err := client.Do(req)
	Expect(err).NotTo(HaveOccurred())

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	Expect(err).NotTo(HaveOccurred())
	defer resp.Body.Close()

	Expect(string(bodyBytes)).To(Equal(`{"Hello": "This is dog"}`))
}
*/

func TestNew(t *testing.T) {
	RegisterTestingT(t)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Add("Content-Type", "application/json")
		w.Write([]byte(`{"token": "this is a testing token"}`))
	}))
	protocol := "http"
	uri, err := url.Parse(server.URL)
	Expect(err).NotTo(HaveOccurred())
	host, port, err := net.SplitHostPort(uri.Host)
	Expect(err).NotTo(HaveOccurred())
	username := "testUser"
	password := "testPass"

	client, err := New(protocol, host, port, username, password)
	Expect(err).NotTo(HaveOccurred())
	Expect(client).NotTo(BeNil())
	Expect(client.Token).NotTo(BeNil())
	Expect(client.Token).To(Equal("this is a testing token"))
}

func TestGetCloudBoltObject(t *testing.T) {
	RegisterTestingT(t)

	var receivedRequest *http.Request

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		receivedRequest = r
		w.WriteHeader(http.StatusOK)
		w.Header().Add("Content-Type", "application/json")
		w.Write([]byte(BodyForGetObject))
	}))

	uri, err := url.Parse(server.URL)
	Expect(err).NotTo(HaveOccurred())
	Expect(uri).NotTo(BeNil())

	client := CloudBoltClient {
		BaseURL:    *uri,
		HTTPClient: &http.Client{},
		Token:      "Whatever the heck we want",
	}

	obj, err := client.GetCloudBoltObject("thing1", "thing2")
	Expect(err).NotTo(HaveOccurred())
	Expect(obj).NotTo(BeNil())

	// TODO make some assertions about the body of `obj`

	Expect(receivedRequest.URL.Path).To(Equal("/api/v2/thing1/"))
	Expect(receivedRequest.URL.RawQuery).To(Equal("filter=name:thing2"))
	Expect(receivedRequest.Header["Authorization"]).To(Equal([]string{"Bearer Whatever the heck we want"}))
}

func TestVerifyGroup(t *testing.T) {
	RegisterTestingT(t)

	Expect(nil).NotTo(BeNil())
}

func TestGetGroup(t *testing.T) {
	RegisterTestingT(t)

	Expect(nil).NotTo(BeNil())
}

func TestDeployBlueprint(t *testing.T) {
	RegisterTestingT(t)

	Expect(nil).NotTo(BeNil())
}

func TestGetOrder(t *testing.T) {
	RegisterTestingT(t)

	Expect(nil).NotTo(BeNil())
}

func TestGetJob(t *testing.T) {
	RegisterTestingT(t)

	Expect(nil).NotTo(BeNil())
}

func TestGetResource(t *testing.T) {
	RegisterTestingT(t)

	Expect(nil).NotTo(BeNil())
}

func TestGetServer(t *testing.T) {
	RegisterTestingT(t)

	Expect(nil).NotTo(BeNil())
}

func TestSubmitAction(t *testing.T) {
	RegisterTestingT(t)

	Expect(nil).NotTo(BeNil())
}

func TestDecomOrder(t *testing.T) {
	RegisterTestingT(t)

	Expect(nil).NotTo(BeNil())
}
