package cbclient

import (
	"bytes"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	. "github.com/onsi/gomega"
)

type fn func(int) string

// The mockRequests data type; an array of http Requests.
type mockRequests []*http.Request

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
	bodyBytes, _ := ioutil.ReadAll(r.Body)
	// Restore the io.ReadCloser to its original state
	r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes)) // Use the content

	return r
}

func mockServer(responseFunc fn) (*httptest.Server, *mockRequests) {
	var requests mockRequests

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requests.append(copyRequest(r))

		w.WriteHeader(http.StatusOK)
		w.Header().Add("Content-Type", "application/json")
		w.Write([]byte(responseFunc(len(requests) - 1)))
	}))

	return server, &requests
}

func bodyToString(b io.ReadCloser) string {
	bodyBytes, _ := ioutil.ReadAll(b)
	bodyString := string(bodyBytes)
	return bodyString
}

func getClient(server *httptest.Server) *CloudBoltClient {
	uri, err := url.Parse(server.URL)
	Expect(err).NotTo(HaveOccurred())
	Expect(uri).NotTo(BeNil())

	token := "Testing Token"

	return &CloudBoltClient{
		BaseURL:    *uri,
		HTTPClient: &http.Client{},
		Token:      token,
	}
}

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
	// Register the test with omega
	RegisterTestingT(t)

	server, _ := mockServer(bodyForTestNew)

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
	// Register the test with omega
	RegisterTestingT(t)

	// Setup mock server
	// Setup requests buffer
	server, requests := mockServer(bodyForGetCloudBoltObject)
	Expect(server).NotTo(BeNil())
	Expect(requests).NotTo(BeNil())

	// Setup CloudBolt Client
	client := getClient(server)
	Expect(client).NotTo(BeNil())

	obj, err := client.GetCloudBoltObject("things", "Thing 2")
	Expect(err).NotTo(HaveOccurred())
	Expect(obj).NotTo(BeNil())

	Expect((*requests)[0].URL.Path).To(Equal("/api/v2/things/"))
	Expect((*requests)[0].URL.RawQuery).To(Equal("filter=name:Thing+2"))
	Expect((*requests)[0].Header["Authorization"]).To(Equal([]string{"Bearer Testing Token"}))

	Expect(obj.Links.Self.Href).To(Equal("/api/v2/things/XYZ-abcdefgh/"))
	Expect(obj.Links.Self.Title).To(Equal("Thing 2"))
	Expect(obj.Name).To(Equal("Thing 2"))
	Expect(obj.ID).To(Equal("3"))

}

func TestVerifyGroup(t *testing.T) {
	// Register the test with omega
	RegisterTestingT(t)

	// Setup mock server
	server, _ := mockServer(bodyForVerifyGroup)
	Expect(server).NotTo(BeNil())

	// Setup CloudBolt Client
	client := getClient(server)
	Expect(client).NotTo(BeNil())

	sampleGroupPath := "/api/v2/groups/GRP-an0thrgrp/"
	sampleParentPath := "the group/the subgroup"

	good, err := client.verifyGroup(sampleGroupPath, sampleParentPath)
	Expect(good).To(BeTrue())
	Expect(err).NotTo(HaveOccurred())
}

// This is a fun test, let's break down what exactly happens.
// If you look in `testData` at `bodyForGetGroup` you see we return four things:
//   - listOfGroups: a response to the query /api/v2/groups/?filter=name:the+childgroup
//   - yetAnotherGroup: a decoy group with the same name. This is allowed in
//     CloudBolt since group names only need to be unique _within_ a subgroup.
//   - aChildGroup: The real group we are looking for.
//   - aSubGroup: The parent of aChildGroup, used to verify this is the "real" group.
//   - aGroup: The parent of aSubGroup, also used to verify this is the "real" group.
// The calls look like this:
//   1. Call to the list of groups.
//   2. Try to verify yetAnotherGroup, which has no parents so it fails.
//   3. Try to verify aChildGroup, it has the correct parent, so verify the parent.
//   4. Try to verify aSubGroup, which also has the correct parent, and reaches
//      the root of the search so we return success in `verifyGroup`, passing the
//      test and finishing the call to GetGroup().
func TestGetGroup(t *testing.T) {
	// Register the test with omega
	RegisterTestingT(t)

	// Setup mock server
	// Setup requests buffer
	server, requests := mockServer(bodyForGetGroup)
	Expect(server).NotTo(BeNil())
	Expect(requests).NotTo(BeNil())

	// Setup CloudBolt Client
	client := getClient(server)
	Expect(client).NotTo(BeNil())

	group, err := client.GetGroup("/the group/the subgroup/the childgroup/")
	Expect(err).NotTo(HaveOccurred())
	Expect(group).NotTo(BeNil())

	Expect(len((*requests))).To(Equal(4))

	Expect((*requests)[0].URL.Path).To(Equal("/api/v2/groups/"))
	Expect((*requests)[0].URL.RawQuery).To(Equal("filter=name:the+childgroup"))
	Expect((*requests)[0].Header["Authorization"]).To(Equal([]string{"Bearer Testing Token"}))

	Expect((*requests)[1].URL.Path).To(Equal("/api/v2/groups/GRP-y3tan0thrgrp/"))
	Expect((*requests)[1].URL.RawQuery).To(Equal(""))
	Expect((*requests)[1].Header["Authorization"]).To(Equal([]string{"Bearer Testing Token"}))

	Expect((*requests)[2].URL.Path).To(Equal("/api/v2/groups/GRP-an0thrgrp/"))
	Expect((*requests)[2].URL.RawQuery).To(Equal(""))
	Expect((*requests)[2].Header["Authorization"]).To(Equal([]string{"Bearer Testing Token"}))

	Expect(group.Links.Self.Href).To(Equal("/api/v2/groups/GRP-an0thrgrp/"))
	Expect(group.Links.Self.Title).To(Equal("the childgroup"))
	Expect(group.Name).To(Equal("the childgroup"))
	Expect(group.ID).To(Equal("512"))
}

func TestDeployBlueprint(t *testing.T) {
	// Register the test with omega
	RegisterTestingT(t)

	// Setup mock server
	// Setup requests buffer
	server, requests := mockServer(bodyForDeployBlueprint)
	Expect(server).NotTo(BeNil())
	Expect(requests).NotTo(BeNil())

	// Setup CloudBolt Client
	client := getClient(server)
	Expect(client).NotTo(BeNil())

	// Get the order items we are going to deploy
	bpItems := bpOrderItems()

	// Submit the order
	// Expect no errors to occur
	cbOrder, err := client.DeployBlueprint("group name", "bp path", "resource name", bpItems)
	Expect(err).NotTo(HaveOccurred())
	Expect(cbOrder).NotTo(BeNil())

	Expect((*requests)[0].URL.Path).To(Equal("/api/v2/orders/"))
	Expect((*requests)[0].Header["Authorization"]).To(Equal([]string{"Bearer Testing Token"}))

	// An example parsed payload we should get back from client.DeployBlueprint(...)
	// Below this comment we make assertions about the parsed response
	// TODO: Consider if we should just compare order to some `expectedOrder` object...
	Expect(cbOrder.Links.Self.Href).To(Equal("/api/v2/orders/101/"))
	Expect(cbOrder.Links.Self.Title).To(Equal("Order id 101"))
	Expect(cbOrder.Links.Group.Href).To(Equal("/api/v2/groups/GRP-th3gr0up/"))
	Expect(cbOrder.Links.Group.Title).To(Equal("the group"))
	Expect(cbOrder.Links.Owner.Href).To(Equal("/api/v2/users/42/"))
	Expect(cbOrder.Links.Owner.Title).To(Equal("the owner"))
	Expect(cbOrder.Links.ApprovedBy.Href).To(Equal("/api/v2/users/42/"))
	Expect(cbOrder.Links.ApprovedBy.Title).To(Equal("the owner"))
	Expect(cbOrder.Links.Actions.Href).To(Equal("/api/v2/actions/2019/"))
	Expect(cbOrder.Links.Actions.Title).To(Equal("the action"))
	Expect(cbOrder.Links.Jobs[0].Href).To(Equal("/api/v2/jobs/1234/"))
	Expect(cbOrder.Links.Jobs[0].Title).To(Equal("Job id 1234"))
	Expect(cbOrder.Name).To(Equal("the order"))
	Expect(cbOrder.ID).To(Equal("1602"))
	Expect(cbOrder.Status).To(Equal("ACTIVE"))
	Expect(cbOrder.Rate).To(Equal("0.12/month"))
	Expect(cbOrder.Items.DeployItems[0].Blueprint).To(Equal("/api/v2/blueprints/BP-ab1u3prt"))
	Expect(cbOrder.Items.DeployItems[0].BlueprintItemsArguments.BuildItemBuildServer.Attributes.Hostname).To(Equal("the hostname"))
	Expect(cbOrder.Items.DeployItems[0].BlueprintItemsArguments.BuildItemBuildServer.Attributes.Quantity).To(Equal(1))
	Expect(cbOrder.Items.DeployItems[0].BlueprintItemsArguments.BuildItemBuildServer.OsBuild).To(Equal("/api/v2/os-builds/OSB-th3058ld/"))
	Expect(cbOrder.Items.DeployItems[0].BlueprintItemsArguments.BuildItemBuildServer.Environment).To(Equal("/api/v2/environments/ENV-th153nv5/"))
	Expect(cbOrder.Items.DeployItems[0].ResourceName).To(Equal("the resource"))
	Expect(cbOrder.Items.DeployItems[0].Blueprint).To(Equal("/api/v2/blueprints/BP-ab1u3prt"))
	Expect(cbOrder.Items.DeployItems[0].ResourceName).To(Equal("the resource"))
	// Expect(cbOrder.Items.DeployItems[0].ResourceParameters).To(...)
	// TODO: We could make assertions about cbOrder.Items.DeployItems[*].BlueprintItemArguments.*
}

func TestGetOrder(t *testing.T) {
	// Register the test with omega
	RegisterTestingT(t)

	// Setup mock server
	// Setup requests buffer
	server, requests := mockServer(bodyForGetOrder)
	Expect(server).NotTo(BeNil())
	Expect(requests).NotTo(BeNil())

	// Setup CloudBolt Client
	client := getClient(server)
	Expect(client).NotTo(BeNil())

	// Define an orderID parameter value
	orderID := "101"

	// Get the CloudBolt Order object
	// Expect no errors to occur
	cbOrder, err := client.GetOrder(orderID)
	Expect(cbOrder).NotTo(BeNil())
	Expect(err).NotTo(HaveOccurred())

	// We expect that `GetOrder` only makes one call to the API
	Expect(len(*requests)).To(Equal(1))
	// We expect that one call to be to the order's endpoint
	Expect((*requests)[0].URL.Path).To(Equal("/api/v2/orders/101"))

	Expect(cbOrder.Links.Self.Href).To(Equal("/api/v2/orders/101/"))
	Expect(cbOrder.Links.Self.Title).To(Equal("Order id 101"))
	Expect(cbOrder.Links.Group.Href).To(Equal("/api/v2/groups/GRP-th3gr0up/"))
	Expect(cbOrder.Links.Group.Title).To(Equal("the group"))
	Expect(cbOrder.Links.Owner.Href).To(Equal("/api/v2/users/42/"))
	Expect(cbOrder.Links.Owner.Title).To(Equal("the owner"))
	Expect(cbOrder.Links.ApprovedBy.Href).To(Equal("/api/v2/users/42/"))
	Expect(cbOrder.Links.ApprovedBy.Title).To(Equal("the owner"))
	Expect(cbOrder.Links.Actions.Href).To(Equal("/api/v2/actions/2019/"))
	Expect(cbOrder.Links.Actions.Title).To(Equal("the action"))
	Expect(cbOrder.Links.Jobs).NotTo(BeEmpty())
	Expect(cbOrder.Links.Jobs[0].Href).To(Equal("/api/v2/jobs/1234/"))
	Expect(cbOrder.Links.Jobs[0].Title).To(Equal("Job id 1234"))
	Expect(cbOrder.Name).To(Equal("the order"))
	Expect(cbOrder.ID).To(Equal("1602"))
	Expect(cbOrder.Status).To(Equal("ACTIVE"))
	Expect(cbOrder.Status).To(Equal("ACTIVE"))
	Expect(cbOrder.Items.DeployItems).NotTo(BeEmpty())
	Expect(cbOrder.Items.DeployItems[0].Blueprint).To(Equal("/api/v2/blueprints/BP-ab1u3prt"))
	Expect(cbOrder.Items.DeployItems[0].ResourceName).To(Equal("the resource"))
	// Expect(cbOrder.Items.DeployItems[0].ResourceParameters).To(...)
	// TODO: We could make assertions about cbOrder.Items.DeployItems[*].BlueprintItemArguments.*
}

func TestGetJob(t *testing.T) {
	// Register the test with omega
	RegisterTestingT(t)

	// Setup mock server
	// Setup requests buffer
	server, requests := mockServer(bodyForGetJob)
	Expect(server).NotTo(BeNil())
	Expect(requests).NotTo(BeNil())

	// Setup CloudBolt Client
	client := getClient(server)
	Expect(client).NotTo(BeNil())

	// Define a jobPath parameter value
	jobPath := "/path/to/job_id"

	// Get the CloudBolt Job object
	// Expect no errors to occur
	cbJob, err := client.GetJob(jobPath)
	Expect(cbJob).NotTo(BeNil())
	Expect(err).NotTo(HaveOccurred())

	Expect(len(*requests)).To(Equal(1))
	Expect((*requests)[0].URL.Path).To(Equal("/path/to/job_id"))

	Expect(cbJob.Links.Self.Href).To(Equal("/api/v2/jobs/1234/"))
	Expect(cbJob.Links.Self.Title).To(Equal("Job id 1234"))
	Expect(cbJob.Links.Owner.Href).To(Equal("/api/v2/users/42/"))
	Expect(cbJob.Links.Owner.Title).To(Equal("the owner"))
	// Expect(cbJob.Links.Parent).To(...)
	Expect(cbJob.Links.Subjobs).NotTo(BeEmpty()) // TODO: More sub-job tests
	// Expect(cbJob.Links.Prerequisite).To(...)
	Expect(cbJob.Links.Order.Href).To(Equal("/api/v2/orders/101/"))
	Expect(cbJob.Links.Order.Title).To(Equal("Order id 101"))
	Expect(cbJob.Links.Resource.Href).To(Equal("/api/v2/resources/big_service/2048/"))
	Expect(cbJob.Links.Resource.Title).To(Equal("A Big Service 2048"))
	Expect(cbJob.Links.Servers).NotTo(BeEmpty())
	Expect(cbJob.Links.Servers[0].Href).To(Equal("/api/v2/servers/128/"))
	Expect(cbJob.Links.Servers[0].Title).To(Equal("a-server-128"))
	Expect(cbJob.Links.LogUrls.RawLog).To(Equal("/api/v2/jobs/1234/log-download-txt/"))
	Expect(cbJob.Links.LogUrls.ZipLog).To(Equal("/api/v2/jobs/1234/log-download"))
	Expect(cbJob.Status).To(Equal("SUCCESS"))
	Expect(cbJob.Type).To(Equal("Deploy Blueprint"))
	Expect(cbJob.Progress.TotalTasks).To(Equal(2))
	Expect(cbJob.Progress.Completed).To(Equal(2))
	Expect(cbJob.Progress.Messages).NotTo(BeEmpty())
	Expect(cbJob.Progress.Messages).To(ContainElement("Deploying blueprint A Big Service."))
	Expect(cbJob.Progress.Messages).To(ContainElement("Starting The server build item"))
	Expect(cbJob.Progress.Messages).To(ContainElement("Blueprint deployment succeeded"))
	Expect(cbJob.Output).To(Equal("Blueprint deployment succeeded"))
}

func TestGetResource(t *testing.T) {
	// Register test with omega
	RegisterTestingT(t)

	// Setup mock server
	// Setup requests buffer
	server, requests := mockServer(bodyForGetResource)
	Expect(server).NotTo(BeNil())
	Expect(requests).NotTo(BeNil())

	// Setup CloudBolt Client
	client := getClient(server)
	Expect(client).NotTo(BeNil())

	// Define a resourcePath parameter value
	resourcePath := "/path/to/resource_id"

	// Get the CloudBolt Resource object
	// Expect no errors to occur
	cbResource, err := client.GetResource(resourcePath)
	Expect(cbResource).NotTo(BeNil())
	Expect(err).NotTo(HaveOccurred())

	Expect(len(*requests)).To(Equal(1))
	Expect((*requests)[0].URL.Path).To(Equal("/path/to/resource_id"))

	Expect(cbResource.Links.Self.Href).To(Equal("/api/v2/resources/big_service/2048/"))
	Expect(cbResource.Links.Self.Title).To(Equal("A Big Service 2048"))
	Expect(cbResource.Links.Blueprint.Href).To(Equal("/api/v2/blueprints/BP-ab1u3prt"))
	Expect(cbResource.Links.Blueprint.Title).To(Equal("a blueprint"))
	Expect(cbResource.Links.Owner.Href).To(Equal("/api/v2/users/42/"))
	Expect(cbResource.Links.Owner.Title).To(Equal("the owner"))
	Expect(cbResource.Links.Group.Href).To(Equal("/api/v2/groups/GRP-th3gr0up/"))
	Expect(cbResource.Links.Group.Title).To(Equal("the group"))
	Expect(cbResource.Links.ResourceType.Href).To(Equal("/api/v2/resource-types/4096/"))
	Expect(cbResource.Links.ResourceType.Title).To(Equal("Big Service"))
	Expect(cbResource.Links.Actions).NotTo(BeEmpty())
	Expect(cbResource.Links.Actions[0].Delete.Href).To(Equal("/api/v2/resources/big_service/2048/actions/1/"))
	Expect(cbResource.Links.Actions[0].Delete.Title).To(Equal("Run 'Delete' on 'A Big Service 2048'"))
	Expect(cbResource.Links.Actions[1].Scale.Href).To(Equal("/api/v2/resources/big_service/2048/actions/2/"))
	Expect(cbResource.Links.Actions[1].Scale.Title).To(Equal("Run 'Scale' on 'A Big Service 2048'"))
	Expect(cbResource.Links.Jobs.Href).To(Equal("/api/v2/resources/big_service/2048/related-jobs/"))
	Expect(cbResource.Links.Jobs.Title).To(Equal("Related Jobs For Resource 'A Big Service 2048'"))
	Expect(cbResource.Links.History.Href).To(Equal("/api/v2/resources/big_service/2048/history/"))
	Expect(cbResource.Links.History.Title).To(Equal("History For Resource 'A Big Service 2048'"))
	Expect(cbResource.Name).To(Equal("A Big Service 2048"))
	Expect(cbResource.ID).To(Equal("2048"))
	Expect(cbResource.Status).To(Equal("Historical"))
	Expect(cbResource.InstallDate).To(Equal("2019-10-29T20:46:34.093868"))
}

func TestGetServer(t *testing.T) {
	// Register the test with omega
	RegisterTestingT(t)

	// Setup mock server
	// Setup requests buffer
	server, requests := mockServer(bodyForGetServer)
	Expect(server).NotTo(BeNil())
	Expect(requests).NotTo(BeNil())

	// Setup CloudBolt Client
	client := getClient(server)
	Expect(client).NotTo(BeNil())

	// Define a serverPath parameter value
	serverPath := "/path/to/server_id"

	// Get the CloudBolt Server object
	// Expect no errors to occur
	cbServer, err := client.GetServer(serverPath)
	Expect(cbServer).NotTo(BeNil())
	Expect(err).NotTo(HaveOccurred())

	Expect(len(*requests)).To(Equal(1))
	Expect((*requests)[0].URL.Path).To(Equal("/path/to/server_id"))

	// Make assertions about the CloudBolt Server object
	Expect(cbServer.Links.Self.Href).To(Equal("/api/v2/servers/128/"))
	Expect(cbServer.Links.Self.Title).To(Equal("a-server-128"))
	Expect(cbServer.Links.Owner.Href).To(Equal("/api/v2/users/42/"))
	Expect(cbServer.Links.Owner.Title).To(Equal("the owner"))
	Expect(cbServer.Links.Group.Href).To(Equal("/api/v2/groups/GRP-th3gr0up/"))
	Expect(cbServer.Links.Group.Title).To(Equal("the group"))
	Expect(cbServer.Links.Environment.Href).To(Equal("/api/v2/environments/ENV-th153nv5/"))
	Expect(cbServer.Links.Environment.Title).To(Equal("the environment"))
	Expect(cbServer.Links.ResourceHandler.Href).To(Equal("/api/v2/resource-handlers/404/"))
	Expect(cbServer.Links.ResourceHandler.Title).To(Equal("Resource Handler Found...ish"))
	Expect(cbServer.Links.Actions[0].PowerOn.Href).To(Equal("/api/v2/servers/128/actions/poweron/"))
	Expect(cbServer.Links.Actions[0].PowerOn.Title).To(Equal("Power on 'a-server-128'"))
	Expect(cbServer.Links.Actions[1].PowerOff.Href).To(Equal("/api/v2/servers/128/actions/poweroff/"))
	Expect(cbServer.Links.Actions[1].PowerOff.Title).To(Equal("Power off 'a-server-128'"))
	Expect(cbServer.Links.Actions[2].Reboot.Href).To(Equal("/api/v2/servers/128/actions/reboot/"))
	Expect(cbServer.Links.Actions[2].Reboot.Title).To(Equal("Reboot 'a-server-128'"))
	Expect(cbServer.Links.Actions[3].RefreshInfo.Href).To(Equal("/api/v2/servers/128/actions/refresh-info/"))
	Expect(cbServer.Links.Actions[3].RefreshInfo.Title).To(Equal("Refresh Info for 'a-server-128'"))
	Expect(cbServer.Links.Actions[4].AdHocScript.Href).To(Equal("/api/v2/servers/128/actions/1/"))
	Expect(cbServer.Links.Actions[4].AdHocScript.Title).To(Equal("Run 'Ad Hoc Script' on 'a-server-128'"))
	Expect(cbServer.Links.Jobs.Href).To(Equal("/api/v2/servers/128/related-jobs/"))
	Expect(cbServer.Links.Jobs.Title).To(Equal("Related Jobs For Server 'a-server-128'"))
	Expect(cbServer.Links.History.Href).To(Equal("/api/v2/servers/5/history/"))
	Expect(cbServer.Links.History.Title).To(Equal("History For Server 'a-server-128'"))
	Expect(cbServer.Hostname).To(Equal("a-server-128"))
	Expect(cbServer.PowerStatus).To(Equal("POWEROFF"))
	Expect(cbServer.Status).To(Equal("ACTIVE"))
	Expect(cbServer.IP).To(Equal("1.2.3.4"))
	Expect(cbServer.Mac).To(Equal("aa:bb:cc:dd:ee:ff"))
	Expect(cbServer.DateAddedToCloudbolt).To(Equal("2019-11-01T18:44:26.670691"))
	Expect(cbServer.CPUCnt).To(Equal(3))
	Expect(cbServer.MemSize).To(Equal("1.2500 GB"))
	Expect(cbServer.DiskSize).To(Equal("56 GB"))
	Expect(cbServer.OsFamily).To(Equal("Linux -&gt; SomeOS"))
	Expect(cbServer.Labels).To(BeEmpty())
	Expect(cbServer.Credentials.Username).To(Equal("TotallyNotRoot"))
	Expect(cbServer.Credentials.Password).To(Equal("not set"))
	Expect(cbServer.Disks[0].UUID).To(Equal("vol-0123456789abcdef1"))
	Expect(cbServer.Disks[0].DiskSize).To(Equal(13))
	Expect(cbServer.Disks[0].Name).To(Equal("also-vol-0123456789abcdef1"))
	Expect(cbServer.Disks[0].Datastore).To(Equal("a-datastore"))
	Expect(cbServer.Disks[0].ProvisioningType).To(Equal("some-provisioning-type"))
	Expect(cbServer.Networks[0].Name).To(Equal("NIC 0"))
	Expect(cbServer.Networks[0].Network).To(Equal("myswitch"))
	Expect(cbServer.Networks[0].Mac).To(Equal("00:11:22:33:44:55"))
	Expect(cbServer.Networks[0].IP).To(Equal("1.2.3.4"))
	Expect(cbServer.Networks[0].PrivateIP).To(Equal("5.6.7.8"))
	Expect(cbServer.Networks[0].AdditionalIps).To(Equal("9.10.11.12"))
	// Expect(cbServer.Parameters).To(...)
	// Expect(cbServer.TechSpecificDetails.*).To(...)
}

func TestSubmitAction(t *testing.T) {
	// Register the test with omega
	RegisterTestingT(t)

	// Setup mock server
	// Setup requests buffer
	server, requests := mockServer(bodyForSubmitAction)
	Expect(server).NotTo(BeNil())
	Expect(requests).NotTo(BeNil())

	// Setup CloudBolt Client
	client := getClient(server)
	Expect(client).NotTo(BeNil())

	// Define an actionPath parameter value
	actionPath := "/path/to/action_id"

	// Get the CloudBolt Action-Result object
	// Expect no errors to occur
	cbActionResult, err := client.SubmitAction(actionPath)
	Expect(cbActionResult).NotTo(BeNil())
	Expect(err).NotTo(HaveOccurred())

	Expect(len(*requests)).To(Equal(1))
	Expect((*requests)[0].URL.Path).To(Equal("/path/to/action_id"))

	Expect(cbActionResult.RunActionJob.Self.Href).To(Equal("/api/v2/jobs/1234"))
	Expect(cbActionResult.RunActionJob.Self.Title).To(Equal("foo"))
}

func TestDecomOrder(t *testing.T) {
	// Register the test with omega
	RegisterTestingT(t)

	// Setup mock server
	// Setup requests buffer
	server, requests := mockServer(bodyForDecomOrder)
	Expect(server).NotTo(BeNil())
	Expect(requests).NotTo(BeNil())

	// Setup CloudBolt Client
	client := getClient(server)
	Expect(client).NotTo(BeNil())

	// Define a parameter values for group, environment, and a slice of servers for decom
	grpPath := "/path/to/group_id"
	envPath := "/path/to/env_id"
	servers := []string{`/path/to/server1_id`, `/path/to/server2_id`, `/path/to/server3_id`}

	// Get the CloudBolt Decom-Order object
	// Expect no errors to occur
	cbOrder, err := client.DecomOrder(grpPath, envPath, servers)
	Expect(cbOrder).NotTo(BeNil())
	Expect(err).NotTo(HaveOccurred())

	Expect(len(*requests)).To(Equal(1))
	Expect((*requests)[0].URL.Path).To(Equal("/api/v2/orders/"))
	Expect((*requests)[0].Method).To(Equal("POST"))
	expectedJSON := `{"group":"/path/to/group_id","items":{"decom-items":[{"environment":"/path/to/env_id","servers":["/path/to/server1_id","/path/to/server2_id","/path/to/server3_id"]}]},"submit-now":"true"}`
	requestBody := bodyToString((*requests)[0].Body)
	Expect(requestBody).To(MatchJSON(expectedJSON))

	Expect(cbOrder.Links.Self.Href).To(Equal("/api/v2/orders/101/"))
	Expect(cbOrder.Links.Self.Title).To(Equal("Order id 101"))
	Expect(cbOrder.Links.Group.Href).To(Equal("/api/v2/groups/GRP-th3gr0up/"))
	Expect(cbOrder.Links.Group.Title).To(Equal("the group"))
	Expect(cbOrder.Links.Owner.Href).To(Equal("/api/v2/users/42/"))
	Expect(cbOrder.Links.Owner.Title).To(Equal("the owner"))
	Expect(cbOrder.Links.ApprovedBy.Href).To(Equal("/api/v2/users/42/"))
	Expect(cbOrder.Links.ApprovedBy.Title).To(Equal("the owner"))
	Expect(cbOrder.Links.Actions.Href).To(Equal("/api/v2/actions/2019/"))
	Expect(cbOrder.Links.Actions.Title).To(Equal("the action"))
	Expect(cbOrder.Links.Jobs[0].Href).To(Equal("/api/v2/jobs/1234/"))
	Expect(cbOrder.Links.Jobs[0].Title).To(Equal("Job id 1234"))
	Expect(cbOrder.Name).To(Equal("the order"))
	Expect(cbOrder.ID).To(Equal("1602"))
	Expect(cbOrder.Status).To(Equal("ACTIVE"))
}
