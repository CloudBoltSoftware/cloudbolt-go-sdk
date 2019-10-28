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

	var requests []*http.Request

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requests = append(requests, r)
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
		Token:      "TestGetCloudBoltObject Token",
	}

	obj, err := client.GetCloudBoltObject("things", "Thing 2")
	Expect(err).NotTo(HaveOccurred())
	Expect(obj).NotTo(BeNil())

	Expect(requests[0].URL.Path).To(Equal("/api/v2/things/"))
	Expect(requests[0].URL.RawQuery).To(Equal("filter=name:Thing+2"))
	Expect(requests[0].Header["Authorization"]).To(Equal([]string{"Bearer TestGetCloudBoltObject Token"}))

	Expect(obj.Links.Self.Href).To(Equal("/api/v2/things/XYZ-abcdefgh/"))
	Expect(obj.Links.Self.Title).To(Equal("Thing 2"))
	Expect(obj.Name).To(Equal("Thing 2"))
	Expect(obj.ID).To(Equal("3"))

}

func TestVerifyGroup(t *testing.T) {
	RegisterTestingT(t)

	// TODO: This is here to force the test to fail
	Expect(nil).NotTo(BeNil())
}

func TestGetGroup(t *testing.T) {
	RegisterTestingT(t)

	var requests []*http.Request

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Add("Content-Type", "application/json")
		w.Write([]byte(BodyForGetGroup(len(requests))))

		requests = append(requests, r)
	}))

	uri, err := url.Parse(server.URL)
	Expect(err).NotTo(HaveOccurred())
	Expect(uri).NotTo(BeNil())

	client := CloudBoltClient {
		BaseURL: *uri,
		HTTPClient: &http.Client{},
		Token: "TestGetGroup Token",
	}

	group, err := client.GetGroup("group name")
	Expect(err).NotTo(HaveOccurred())
	Expect(group).NotTo(BeNil())

	Expect(requests[0].URL.Path).To(Equal("/api/v2/groups/"))
	Expect(requests[0].URL.RawQuery).To(Equal("filter=name:group+name"))
	Expect(requests[0].Header["Authorization"]).To(Equal([]string{"Bearer TestGetGroup Token"}))

	Expect(requests[1].URL.Path).To(Equal("/api/v2/groups/GRP-th3gr0up/"))
	Expect(requests[1].URL.RawQuery).To(Equal(""))
	Expect(requests[1].Header["Authorization"]).To(Equal([]string{"Bearer TestGetGroup Token"}))

	Expect(len(requests)).To(Equal(2))

	Expect(group.Links.Self.Href).To(Equal("/api/v2/groups/GRP-th3gr0up"))
	Expect(group.Links.Self.Title).To(Equal("the group"))
	Expect(group.Name).To(Equal("the group"))
	Expect(group.ID).To(Equal(6))
}

func TestDeployBlueprint(t *testing.T) {
	RegisterTestingT(t)

	// TODO: This is here to force the test to fail
	Expect(nil).NotTo(BeNil())
}

func TestGetOrder(t *testing.T) {
	RegisterTestingT(t)

	// TODO: This is here to force the test to fail
	Expect(nil).NotTo(BeNil())
}

func TestGetJob(t *testing.T) {
	RegisterTestingT(t)

	// TODO: This is here to force the test to fail
	Expect(nil).NotTo(BeNil())
}

func TestGetResource(t *testing.T) {
	RegisterTestingT(t)

	// TODO: This is here to force the test to fail
	Expect(nil).NotTo(BeNil())
}

func TestGetServer(t *testing.T) {
	RegisterTestingT(t)

	// TODO: This is here to force the test to fail
	Expect(nil).NotTo(BeNil())
}

func TestSubmitAction(t *testing.T) {
	RegisterTestingT(t)

	// TODO: This is here to force the test to fail
	Expect(nil).NotTo(BeNil())
}

func TestDecomOrder(t *testing.T) {
	RegisterTestingT(t)

	// TODO: This is here to force the test to fail
	Expect(nil).NotTo(BeNil())
}
