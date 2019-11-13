package cbclient

import (
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

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
		w.Write([]byte(bodyForGetObject))
	}))

	uri, err := url.Parse(server.URL)
	Expect(err).NotTo(HaveOccurred())
	Expect(uri).NotTo(BeNil())

	client := CloudBoltClient{
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
		w.Write([]byte(bodyForGetGroup(len(requests))))

		requests = append(requests, r)
	}))

	uri, err := url.Parse(server.URL)
	Expect(err).NotTo(HaveOccurred())
	Expect(uri).NotTo(BeNil())

	client := CloudBoltClient{
		BaseURL:    *uri,
		HTTPClient: &http.Client{},
		Token:      "TestGetGroup Token",
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

	Expect(group.Links.Self.Href).To(Equal("/api/v2/groups/GRP-th3gr0up/"))
	Expect(group.Links.Self.Title).To(Equal("the group"))
	Expect(group.Name).To(Equal("the group"))
	Expect(group.ID).To(Equal("6"))
}

func TestDeployBlueprint(t *testing.T) {
	RegisterTestingT(t)

	var requests []*http.Request

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Add("Content-Type", "application/json")
		w.Write([]byte(bodyForDeployBlueprint(len(requests))))

		requests = append(requests, r)
	}))

	uri, err := url.Parse(server.URL)
	Expect(err).NotTo(HaveOccurred())
	Expect(uri).NotTo(BeNil())

	client := CloudBoltClient{
		BaseURL:    *uri,
		HTTPClient: &http.Client{},
		Token:      "TestGetGroup Token",
	}

	bpItems := bpOrderItems()

	order, err := client.DeployBlueprint("group name", "bp path", "resource name", bpItems)
	Expect(err).NotTo(HaveOccurred())
	Expect(order).NotTo(BeNil())

	/*
		{
			Links:{
				Self:{
					Href:/api/v2/orders/42/
					Title:Order id 42
				}
				Group:{
					Href:/api/v2/groups/GRP-ncrggwao/
					Title:TerraformOrderGroup01
				}
				Owner:{
					Href:/api/v2/users/2/
					Title:Some User
				}
				ApprovedBy:{
					Href:/api/v2/users/2/
					Title:Some User
				}
				Actions:{
					Href:
					Title:
				}
				Jobs:[
					{
					Href:/api/v2/jobs/1285/
					Title:Job id 1285
					}
				]
			}
			Name:
			ID:42
			Status:ACTIVE
			Rate:0.00/month
			CreateDate:2019-10-28T18:39:32.099576
			ApproveDate:2019-10-28T18:39:32.420531
			Items:{
				DeployItems:[
					{
					Blueprint:/api/v2/blueprints/BP-4gph95o9/
					BlueprintItemsArguments:{
						BuildItemBuildServer:{
							Attributes:{
								Hostname:
								Quantity:0
							}
							OsBuild:
							Environment:
							Parameters:map[]
						}
					}
					ResourceName:resource name
					ResourceParameters:{ }
					}
				]
			}
		}
	*/
	Expect(requests[0].URL.Path).To(Equal("/api/v2/orders/"))
	Expect(requests[0].Header["Authorization"]).To(Equal([]string{"Bearer TestGetGroup Token"}))

	Expect(order.Links.Self.Href).To(Equal("/api/v2/orders/101/"))
	Expect(order.Links.Self.Title).To(Equal("Order id 101"))
	Expect(order.Links.Group.Href).To(Equal("/api/v2/groups/GRP-th3gr0up/"))
	Expect(order.Links.Group.Title).To(Equal("the group"))
	Expect(order.Links.Owner.Href).To(Equal("/api/v2/users/42/"))
	Expect(order.Links.Owner.Title).To(Equal("the owner"))
	Expect(order.Links.ApprovedBy.Href).To(Equal("/api/v2/users/42/"))
	Expect(order.Links.ApprovedBy.Title).To(Equal("the owner"))
	Expect(order.Links.Actions.Href).To(Equal("/api/v2/actions/2019"))
	Expect(order.Links.Actions.Title).To(Equal("the action"))
	Expect(order.Links.Jobs[0].Href).To(Equal("/api/v2/jobs/1234/"))
	Expect(order.Links.Jobs[0].Title).To(Equal("Job id 1234"))
	Expect(order.Name).To(Equal("the order"))
	Expect(order.ID).To(Equal("1602"))
	Expect(order.Status).To(Equal("ACTIVE"))
	Expect(order.Rate).To(Equal("0.12/month"))
	Expect(order.Items.DeployItems[0].Blueprint).To(Equal("/api/v2/blueprints/BP-ab1u3prt"))
	Expect(order.Items.DeployItems[0].BlueprintItemsArguments.BuildItemBuildServer.Attributes.Hostname).To(Equal("the hostname"))
	Expect(order.Items.DeployItems[0].BlueprintItemsArguments.BuildItemBuildServer.Attributes.Quantity).To(Equal(1))
	Expect(order.Items.DeployItems[0].BlueprintItemsArguments.BuildItemBuildServer.OsBuild).To(Equal("/api/v2/os-builds/OSB-th3058ld/"))
	Expect(order.Items.DeployItems[0].BlueprintItemsArguments.BuildItemBuildServer.Environment).To(Equal("/api/v2/environments/ENV-th153nv5/"))
	Expect(order.Items.DeployItems[0].ResourceName).To(Equal("the resource"))
	// TODO: Test these too
	// Expect(order.Items.DeployItems[0].BlueprintItemsArguments.BuildItemBuildServer.Parameters).To(Equal(...))
	// Expect(order.Items.DeployItems[0].ResourceParameters).To(Equal(...))
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
