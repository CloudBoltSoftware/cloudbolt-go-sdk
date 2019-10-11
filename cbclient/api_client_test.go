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
