package cbclient

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestGetIPAMPolicy(t *testing.T) {
	// Register the test with gomega
	RegisterTestingT(t)

	// Setup mock server with scripted responses
	// Setup requests buffer
	server, requests := mockServer(responsesForIPAMPolicy)
	Expect(server).NotTo(BeNil())
	Expect(requests).NotTo(BeNil())

	// Setup CloudBolt Client
	client := getClient(server)
	Expect(client).NotTo(BeNil())

	policyName := "MY_IPAM_POLICY"
	policy, err := client.GetIPAMPolicy(policyName)
	Expect(policy).NotTo(BeNil())
	Expect(err).NotTo(HaveOccurred())

	// This should have made three requests:
	// 1+2. Fail to get, get a token
	// 3. Successful Get
	Expect(len(*requests)).To(Equal(3))

	// The last request is the one we care about
	Expect((*requests)[2].URL.Path).To(Equal("/api/v3/onefuse/ipamPolicies/"))

	// The CloudBolt Order object should be parsed correctly
	Expect(policy.Links.Self.Href).To(Equal("/api/v3/onefuse/ipamPolicies/3/"))
	Expect(policy.Links.Self.Title).To(Equal("MY_IPAM_POLICY"))
	Expect(policy.Links.Workspace.Href).To(Equal("/api/v3/onefuse/workspaces/2/"))
	Expect(policy.Links.Workspace.Title).To(Equal("Default"))
	Expect(policy.Links.Endpoint.Href).To(Equal("/api/v3/onefuse/endpoints/8/"))
	Expect(policy.Links.Endpoint.Title).To(Equal("MY_IPAM_POLICY_Endpoint"))
	Expect(policy.Name).To(Equal("MY_IPAM_POLICY"))
	Expect(policy.ID).To(Equal(3))
	Expect(policy.Description).To(Equal("An IPAM Policy created through automated tests"))
}

func TestGetIPAMReservationById(t *testing.T) {
	// Register the test with gomega
	RegisterTestingT(t)

	// Setup mock server with scripted responses
	// Setup requests buffer
	server, requests := mockServer(responsesForIPAMReservation)
	Expect(server).NotTo(BeNil())
	Expect(requests).NotTo(BeNil())

	// Setup CloudBolt Client
	client := getClient(server)
	Expect(client).NotTo(BeNil())

	ipamReservation, err := client.GetIPAMReservationById("10")
	Expect(ipamReservation).NotTo(BeNil())
	Expect(err).NotTo(HaveOccurred())

	// This should have made three requests:
	// 1+2. Fail to get, get a token
	// 3. Successful Get
	Expect(len(*requests)).To(Equal(3))

	// The last request is the one we care about
	Expect((*requests)[2].URL.Path).To(Equal("/api/v3/onefuse/ipamReservations/10/"))

	// The CloudBolt Order object should be parsed correctly
	Expect(ipamReservation.Links.Self.Href).To(Equal("/api/v3/ipamReservations/10/"))
	Expect(ipamReservation.Links.Self.Title).To(Equal("test010"))
	Expect(ipamReservation.Links.Workspace.Href).To(Equal("/api/v3/onefuse/workspaces/1/"))
	Expect(ipamReservation.Links.Workspace.Title).To(Equal("Default"))
	Expect(ipamReservation.Links.Policy.Href).To(Equal("/api/v3/onefuse/ipamPolicies/1/"))
	Expect(ipamReservation.Links.Policy.Title).To(Equal("infoblox1"))
	Expect(ipamReservation.Links.JobMetadata.Href).To(Equal("/api/v3/onefuse/jobMetadata/86/"))
	Expect(ipamReservation.Links.JobMetadata.Title).To(Equal("Job Metadata Record id 86"))
	Expect(ipamReservation.ID).To(Equal(10))
	Expect(ipamReservation.PrimaryDNS).To(Equal("8.8.8.8"))
	Expect(ipamReservation.SecondaryDNS).To(Equal("10.0.1.1"))
	Expect(ipamReservation.DNSSuffix).To(Equal("12.0.0.1"))
	Expect(ipamReservation.NicLabel).To(Equal("NIC 1"))
	Expect(ipamReservation.Subnet).To(Equal("10.0.1.1/24"))
	Expect(ipamReservation.Gateway).To(Equal("10.0.1.1"))
	Expect(ipamReservation.Network).To(Equal("10.0.1.2"))
	Expect(ipamReservation.Netmask).To(Equal("255.255.255.0"))
}

func TestGetIPAMReservation(t *testing.T) {
	// Register the test with gomega
	RegisterTestingT(t)

	// Setup mock server with scripted responses
	// Setup requests buffer
	server, requests := mockServer(responsesForIPAMReservation)
	Expect(server).NotTo(BeNil())
	Expect(requests).NotTo(BeNil())

	// Setup CloudBolt Client
	client := getClient(server)
	Expect(client).NotTo(BeNil())

	ipamReservationPath := "/api/v3/onefuse/ipamReservations/10/"
	ipamReservation, err := client.GetIPAMReservation(ipamReservationPath)
	Expect(ipamReservation).NotTo(BeNil())
	Expect(err).NotTo(HaveOccurred())

	// This should have made three requests:
	// 1+2. Fail to get, get a token
	// 3. Successful Get
	Expect(len(*requests)).To(Equal(3))

	// The last request is the one we care about
	Expect((*requests)[2].URL.Path).To(Equal("/api/v3/onefuse/ipamReservations/10/"))

	// The CloudBolt Order object should be parsed correctly
	Expect(ipamReservation.Links.Self.Href).To(Equal("/api/v3/ipamReservations/10/"))
	Expect(ipamReservation.Links.Self.Title).To(Equal("test010"))
	Expect(ipamReservation.Links.Workspace.Href).To(Equal("/api/v3/onefuse/workspaces/1/"))
	Expect(ipamReservation.Links.Workspace.Title).To(Equal("Default"))
	Expect(ipamReservation.Links.Policy.Href).To(Equal("/api/v3/onefuse/ipamPolicies/1/"))
	Expect(ipamReservation.Links.Policy.Title).To(Equal("infoblox1"))
	Expect(ipamReservation.Links.JobMetadata.Href).To(Equal("/api/v3/onefuse/jobMetadata/86/"))
	Expect(ipamReservation.Links.JobMetadata.Title).To(Equal("Job Metadata Record id 86"))
	Expect(ipamReservation.ID).To(Equal(10))
	Expect(ipamReservation.PrimaryDNS).To(Equal("8.8.8.8"))
	Expect(ipamReservation.SecondaryDNS).To(Equal("10.0.1.1"))
	Expect(ipamReservation.DNSSuffix).To(Equal("12.0.0.1"))
	Expect(ipamReservation.NicLabel).To(Equal("NIC 1"))
	Expect(ipamReservation.Subnet).To(Equal("10.0.1.1/24"))
	Expect(ipamReservation.Gateway).To(Equal("10.0.1.1"))
	Expect(ipamReservation.Network).To(Equal("10.0.1.2"))
	Expect(ipamReservation.Netmask).To(Equal("255.255.255.0"))
}

func TestCreateIPAMReservation(t *testing.T) {
	// Register the test with gomega
	RegisterTestingT(t)

	// Setup mock server with scripted responses
	// Setup requests buffer
	server, requests := mockServer(responsesForGetJobStatus)
	Expect(server).NotTo(BeNil())
	Expect(requests).NotTo(BeNil())

	// Setup CloudBolt Client
	client := getClient(server)
	Expect(client).NotTo(BeNil())

	newIPAMRecord := IPAMReservation{
		PolicyID:     650,
		WorkspaceURL: "/api/v3/onefuse/workspaces/1/",
		TemplateProperties: map[string]interface{}{
			"comment":                  "QA API Generated",
			"BlueCatConfigurationName": "Configuration 01",
			"BlueCatDnsView":           "bluecatDnsView",
			"config_name":              "Configuration 01",
			"gateway":                  "10.30.32.1",
			"netmask":                  "255.255.255.0",
			"network_name":             "default",
			"subnet":                   "10.30.32.0/24",
			"dns_suffix":               "dnssuffix.net",
			"hostname_override":        "test.hostname.override",
		},
	}

	jobStatus, err := client.CreateIPAMReservation(&newIPAMRecord)
	Expect(jobStatus).NotTo(BeNil())
	Expect(err).NotTo(HaveOccurred())

	// This should have made three requests:
	// 1+2. Fail to get, get a token
	// 3. Successful Get
	Expect(len(*requests)).To(Equal(3))

	// The last request is the one we care about
	Expect((*requests)[2].URL.Path).To(Equal("/api/v3/onefuse/ipamReservations/"))

	// The CloudBolt Order object should be parsed correctly
	verifyJobStatus(jobStatus)
}

func TestDeleteIPAMReservation(t *testing.T) {
	// Register the test with gomega
	RegisterTestingT(t)

	// Setup mock server with scripted responses
	// Setup requests buffer
	server, requests := mockServer(responsesForGetJobStatus)
	Expect(server).NotTo(BeNil())
	Expect(requests).NotTo(BeNil())

	// Setup CloudBolt Client
	client := getClient(server)
	Expect(client).NotTo(BeNil())

	jobStatus, err := client.DeleteIPAMReservation("10")
	Expect(jobStatus).NotTo(BeNil())
	Expect(err).NotTo(HaveOccurred())

	// This should have made three requests:
	// 1+2. Fail to get, get a token
	// 3. Successful Get
	Expect(len(*requests)).To(Equal(3))

	// The last request is the one we care about
	Expect((*requests)[2].URL.Path).To(Equal("/api/v3/onefuse/ipamReservations/10/"))

	// The CloudBolt Order object should be parsed correctly
	verifyJobStatus(jobStatus)
}
