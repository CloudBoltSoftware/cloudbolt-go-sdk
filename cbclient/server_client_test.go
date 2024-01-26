package cbclient

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestGetServer(t *testing.T) {
	// Register the test with gomega
	RegisterTestingT(t)

	// Setup mock server with scripted responses
	// Setup requests buffer
	server, requests := mockServer(responsesForGetServer)
	Expect(server).NotTo(BeNil())
	Expect(requests).NotTo(BeNil())

	// Setup CloudBolt Client
	client := getClient(server)
	Expect(client).NotTo(BeNil())

	// Define a serverPath parameter value
	serverPath := "/api/v3/cmp/servers/SVR-yrk09wht/"

	// Get the CloudBolt Server object
	// Expect no errors to occur
	cbServer, err := client.GetServer(serverPath)
	Expect(cbServer).NotTo(BeNil())
	Expect(err).NotTo(HaveOccurred())

	// This should have made three requests:
	// 1+2. Fail to get server, get a token
	// 3. Successfully getting the server
	Expect(len(*requests)).To(Equal(3))

	// The last request is the one we care about
	Expect((*requests)[2].URL.Path).To(Equal("/api/v3/cmp/servers/SVR-yrk09wht/"))

	// The CloudBolt Server object should be parsed correctly
	Expect(cbServer.Links.Self.Href).To(Equal("/api/v3/cmp/servers/SVR-yrk09wht/"))
	Expect(cbServer.Links.Self.Title).To(Equal("myawsinstance1"))
	Expect(cbServer.ID).To(Equal("SVR-yrk09wht"))
	Expect(cbServer.Hostname).To(Equal("myawsinstance1"))
	Expect(cbServer.IP).To(Equal("3.17.176.101"))
	Expect(cbServer.Status).To(Equal("ACTIVE"))
	Expect(cbServer.Mac).To(Equal("02:99:e2:0f:18:b2"))
	Expect(cbServer.PowerStatus).To(Equal("POWERON"))
	Expect(cbServer.DateAddedToCloudbolt).To(Equal("2022-04-08 11:55:07.056038"))
	Expect(cbServer.CPUCount).To(Equal(1))
	Expect(cbServer.MemorySizeGB).To(Equal("0.5000"))
	Expect(cbServer.DiskSizeGB).To(Equal(8))
	Expect(cbServer.Notes).To(Equal(""))
	Expect(len(cbServer.Labels)).To(Equal(0))
	Expect(cbServer.OsFamily).To(Equal("Amazon Linux"))
	Expect(cbServer.RateBreakdown).To(Not(BeNil()))
	Expect(cbServer.Attributes).To(Not(BeNil()))
	Expect(cbServer.Credentials).To(Not(BeNil()))
	Expect(len(cbServer.Disks)).To(Equal(1))
	Expect(len(cbServer.Networks)).To(Equal(1))
	Expect(cbServer.TechSpecificAttributes).To(Not(BeNil()))
}

func TestGetServerById(t *testing.T) {
	// Register the test with gomega
	RegisterTestingT(t)

	// Setup mock server with scripted responses
	// Setup requests buffer
	server, requests := mockServer(responsesForGetServer)
	Expect(server).NotTo(BeNil())
	Expect(requests).NotTo(BeNil())

	// Setup CloudBolt Client
	client := getClient(server)
	Expect(client).NotTo(BeNil())

	// Define a serverPath parameter value
	serverId := "SVR-yrk09wht"

	// Get the CloudBolt Server object
	// Expect no errors to occur
	cbServer, err := client.GetServerById(serverId)
	Expect(cbServer).NotTo(BeNil())
	Expect(err).NotTo(HaveOccurred())

	// This should have made three requests:
	// 1+2. Fail to get server, get a token
	// 3. Successfully getting the server
	Expect(len(*requests)).To(Equal(3))

	// The last request is the one we care about
	Expect((*requests)[2].URL.Path).To(Equal("/api/v3/cmp/servers/SVR-yrk09wht/"))

	// The CloudBolt Server object should be parsed correctly
	Expect(cbServer.Links.Self.Href).To(Equal("/api/v3/cmp/servers/SVR-yrk09wht/"))
	Expect(cbServer.Links.Self.Title).To(Equal("myawsinstance1"))
	Expect(cbServer.ID).To(Equal("SVR-yrk09wht"))
	Expect(cbServer.Hostname).To(Equal("myawsinstance1"))
	Expect(cbServer.IP).To(Equal("3.17.176.101"))
	Expect(cbServer.Status).To(Equal("ACTIVE"))
	Expect(cbServer.Mac).To(Equal("02:99:e2:0f:18:b2"))
	Expect(cbServer.PowerStatus).To(Equal("POWERON"))
	Expect(cbServer.DateAddedToCloudbolt).To(Equal("2022-04-08 11:55:07.056038"))
	Expect(cbServer.CPUCount).To(Equal(1))
	Expect(cbServer.MemorySizeGB).To(Equal("0.5000"))
	Expect(cbServer.DiskSizeGB).To(Equal(8))
	Expect(cbServer.Notes).To(Equal(""))
	Expect(len(cbServer.Labels)).To(Equal(0))
	Expect(cbServer.OsFamily).To(Equal("Amazon Linux"))
	Expect(cbServer.RateBreakdown).To(Not(BeNil()))
	Expect(cbServer.Attributes).To(Not(BeNil()))
	Expect(cbServer.Credentials).To(Not(BeNil()))
	Expect(len(cbServer.Disks)).To(Equal(1))
	Expect(len(cbServer.Networks)).To(Equal(1))
	Expect(cbServer.TechSpecificAttributes).To(Not(BeNil()))
}

func TestGetServerByHostname(t *testing.T) {
	// Register the test with gomega
	RegisterTestingT(t)

	// Setup mock server with scripted responses
	// Setup requests buffer
	server, requests := mockServer(responsesForGetServerByHostname)
	Expect(server).NotTo(BeNil())
	Expect(requests).NotTo(BeNil())

	// Setup CloudBolt Client
	client := getClient(server)
	Expect(client).NotTo(BeNil())

	// Define a serverPath parameter value
	hostname := "myawsinstance1"

	// Get the CloudBolt Server object
	// Expect no errors to occur
	cbServer, err := client.GetServerByHostname(hostname)
	Expect(cbServer).NotTo(BeNil())
	Expect(err).NotTo(HaveOccurred())

	// This should have made three requests:
	// 1+2. Fail to get server, get a token
	// 3. Successfully getting the server
	Expect(len(*requests)).To(Equal(3))

	// The last request is the one we care about
	Expect((*requests)[2].URL.Path).To(Equal("/api/v3/cmp/servers/"))

	// The CloudBolt Server object should be parsed correctly
	Expect(cbServer.Links.Self.Href).To(Equal("/api/v3/cmp/servers/SVR-yrk09wht/"))
	Expect(cbServer.Links.Self.Title).To(Equal("myawsinstance1"))
	Expect(cbServer.ID).To(Equal("SVR-yrk09wht"))
	Expect(cbServer.Hostname).To(Equal("myawsinstance1"))
	Expect(cbServer.IP).To(Equal("3.17.176.101"))
	Expect(cbServer.Status).To(Equal("ACTIVE"))
	Expect(cbServer.Mac).To(Equal("02:99:e2:0f:18:b2"))
	Expect(cbServer.PowerStatus).To(Equal("POWERON"))
	Expect(cbServer.DateAddedToCloudbolt).To(Equal("2022-04-08 11:55:07.056038"))
	Expect(cbServer.CPUCount).To(Equal(1))
	Expect(cbServer.MemorySizeGB).To(Equal("0.5000"))
	Expect(cbServer.DiskSizeGB).To(Equal(8))
	Expect(cbServer.Notes).To(Equal(""))
	Expect(len(cbServer.Labels)).To(Equal(0))
	Expect(cbServer.OsFamily).To(Equal("Amazon Linux"))
	Expect(cbServer.RateBreakdown).To(Not(BeNil()))
	Expect(cbServer.Attributes).To(Not(BeNil()))
	Expect(cbServer.Credentials).To(Not(BeNil()))
	Expect(len(cbServer.Disks)).To(Equal(1))
	Expect(len(cbServer.Networks)).To(Equal(1))
	Expect(cbServer.TechSpecificAttributes).To(Not(BeNil()))
}

func TestDecomServerOrderResponse(t *testing.T) {
	// Register the test with gomega
	RegisterTestingT(t)

	// Setup mock server with scripted responses
	// Setup requests buffer
	server, requests := mockServer(responsesForDecomServerOrder)
	Expect(server).NotTo(BeNil())
	Expect(requests).NotTo(BeNil())

	// Setup CloudBolt Client
	client := getClient(server)
	Expect(client).NotTo(BeNil())

	// Define a serverPath parameter value
	serverId := "SVR-d7xr7for"

	// Decom Server
	// Expect no errors to occur
	order, err := client.DecomServer(serverId)
	Expect(order).NotTo(BeNil())
	Expect(err).NotTo(HaveOccurred())

	// This should have made three requests:
	// 1+2. Fail to get server, get a token
	// 3. Successfully getting the server
	Expect(len(*requests)).To(Equal(3))

	// The last request is the one we care about
	Expect((*requests)[2].URL.Path).To(Equal("/api/v3/cmp/servers/SVR-d7xr7for/decommission/"))
	Expect(order.Links.Self.Href).To(Equal("/api/v3/cmp/orders/ORD-ijudvhqv/"))
	Expect(order.Links.Self.Title).To(Equal("Deletion of myawsinstance2"))
	Expect(order.ID).To(Equal("ORD-ijudvhqv"))
}

func TestDecomServerJobResponse(t *testing.T) {
	// Register the test with gomega
	RegisterTestingT(t)

	// Setup mock server with scripted responses
	// Setup requests buffer
	server, requests := mockServer(responsesForDecomServerJob)
	Expect(server).NotTo(BeNil())
	Expect(requests).NotTo(BeNil())

	// Setup CloudBolt Client
	client := getClient(server)
	Expect(client).NotTo(BeNil())

	// Define a serverPath parameter value
	serverId := "SVR-668kqo0f"

	// Decom Server
	// Expect no errors to occur
	order, err := client.DecomServer(serverId)
	Expect(order).NotTo(BeNil())
	Expect(err).NotTo(HaveOccurred())

	// This should have made three requests:
	// 1+2. Fail to get server, get a token
	// 3. Successfully getting the server
	Expect(len(*requests)).To(Equal(3))

	// The last request is the one we care about
	Expect((*requests)[2].URL.Path).To(Equal("/api/v3/cmp/servers/SVR-668kqo0f/decommission/"))
	Expect(order.Links.Self.Href).To(Equal("/api/v3/cmp/jobs/JOB-80uh0rmr/"))
	Expect(order.Links.Self.Title).To(Equal("Delete Server Job 502"))
	Expect(order.ID).To(Equal("JOB-80uh0rmr"))
}
