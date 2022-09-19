package cbclient

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestGetDNSPolicy(t *testing.T) {
	// Register the test with gomega
	RegisterTestingT(t)

	// Setup mock server with scripted responses
	// Setup requests buffer
	server, requests := mockServer(responsesForDNSPolicy)
	Expect(server).NotTo(BeNil())
	Expect(requests).NotTo(BeNil())

	// Setup CloudBolt Client
	client := getClient(server)
	Expect(client).NotTo(BeNil())

	policyName := "my_infoblox_dns_policy"
	policy, err := client.GetDNSPolicy(policyName)
	Expect(policy).NotTo(BeNil())
	Expect(err).NotTo(HaveOccurred())

	// This should have made three requests:
	// 1+2. Fail to get policy, get a token
	// 3. Successfully getting the order
	Expect(len(*requests)).To(Equal(3))

	// The last request is the one we care about
	Expect((*requests)[2].URL.Path).To(Equal("/api/v3/onefuse/dnsPolicies/"))

	// The CloudBolt Order object should be parsed correctly
	Expect(policy.Links.Self.Href).To(Equal("/api/v3/onefuse/dnsPolicies/9/"))
	Expect(policy.Links.Self.Title).To(Equal("my_infoblox_dns_policy"))
	Expect(policy.Links.Workspace.Href).To(Equal("/api/v3/onefuse/workspaces/2/"))
	Expect(policy.Links.Workspace.Title).To(Equal("Default"))
	Expect(policy.Links.Endpoint.Href).To(Equal("/api/v3/onefuse/endpoints/14/"))
	Expect(policy.Links.Endpoint.Title).To(Equal("my_infoblox_dns_policy_Endpoint"))
	Expect(policy.Name).To(Equal("my_infoblox_dns_policy"))
	Expect(policy.ID).To(Equal(9))
	Expect(policy.Description).To(Equal("An Infoblox DNS Policy created through automated tests"))
}

func TestGetDNSReservationById(t *testing.T) {
	// Register the test with gomega
	RegisterTestingT(t)

	// Setup mock server with scripted responses
	// Setup requests buffer
	server, requests := mockServer(responsesForDNSReservation)
	Expect(server).NotTo(BeNil())
	Expect(requests).NotTo(BeNil())

	// Setup CloudBolt Client
	client := getClient(server)
	Expect(client).NotTo(BeNil())

	dnsReservation, err := client.GetDNSReservationById("10")
	Expect(dnsReservation).NotTo(BeNil())
	Expect(err).NotTo(HaveOccurred())

	// This should have made three requests:
	// 1+2. Fail to get reservation, get a token
	// 3. Successfully getting the order
	Expect(len(*requests)).To(Equal(3))

	// The last request is the one we care about
	Expect((*requests)[2].URL.Path).To(Equal("/api/v3/onefuse/dnsReservations/10/"))

	// The CloudBolt Order object should be parsed correctly
	Expect(dnsReservation.Links.Self.Href).To(Equal("/api/v3/dnsReservations/10/"))
	Expect(dnsReservation.Links.Self.Title).To(Equal("test010"))
	Expect(dnsReservation.Links.Workspace.Href).To(Equal("/api/v3/onefuse/workspaces/1/"))
	Expect(dnsReservation.Links.Workspace.Title).To(Equal("Default"))
	Expect(dnsReservation.Links.Policy.Href).To(Equal("/api/v3/onefuse/dnsPolicies/1/"))
	Expect(dnsReservation.Links.Policy.Title).To(Equal("infoblox1"))
	Expect(dnsReservation.Links.JobMetadata.Href).To(Equal("/api/v3/onefuse/jobMetadata/86/"))
	Expect(dnsReservation.Links.JobMetadata.Title).To(Equal("Job Metadata Record id 86"))
	Expect(dnsReservation.ID).To(Equal(10))
	Expect(dnsReservation.Name).To(Equal("test010"))
}

func TestGetDNSReservationBy(t *testing.T) {
	// Register the test with gomega
	RegisterTestingT(t)

	// Setup mock server with scripted responses
	// Setup requests buffer
	server, requests := mockServer(responsesForDNSReservation)
	Expect(server).NotTo(BeNil())
	Expect(requests).NotTo(BeNil())

	// Setup CloudBolt Client
	client := getClient(server)
	Expect(client).NotTo(BeNil())

	dnsReservationPath := "/api/v3/onefuse/dnsReservations/10/"
	dnsReservation, err := client.GetDNSReservation(dnsReservationPath)
	Expect(dnsReservation).NotTo(BeNil())
	Expect(err).NotTo(HaveOccurred())

	// This should have made three requests:
	// 1+2. Fail to get reservation, get a token
	// 3. Successfully getting the order
	Expect(len(*requests)).To(Equal(3))

	// The last request is the one we care about
	Expect((*requests)[2].URL.Path).To(Equal("/api/v3/onefuse/dnsReservations/10/"))

	// The CloudBolt Order object should be parsed correctly
	Expect(dnsReservation.Links.Self.Href).To(Equal("/api/v3/dnsReservations/10/"))
	Expect(dnsReservation.Links.Self.Title).To(Equal("test010"))
	Expect(dnsReservation.Links.Workspace.Href).To(Equal("/api/v3/onefuse/workspaces/1/"))
	Expect(dnsReservation.Links.Workspace.Title).To(Equal("Default"))
	Expect(dnsReservation.Links.Policy.Href).To(Equal("/api/v3/onefuse/dnsPolicies/1/"))
	Expect(dnsReservation.Links.Policy.Title).To(Equal("infoblox1"))
	Expect(dnsReservation.Links.JobMetadata.Href).To(Equal("/api/v3/onefuse/jobMetadata/86/"))
	Expect(dnsReservation.Links.JobMetadata.Title).To(Equal("Job Metadata Record id 86"))
	Expect(dnsReservation.ID).To(Equal(10))
	Expect(dnsReservation.Name).To(Equal("test010"))
}

func TestCreateDNSReservation(t *testing.T) {
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

	newDNSRecord := DNSReservation{
		Name:         "test010",
		PolicyID:     1,
		WorkspaceURL: "/api/v3/onefuse/workspaces/1/",
		Value:        "My DNS Reservation",
		TemplateProperties: map[string]interface{}{
			"Environment": "Production",
			"Application": "Web Servers",
			"OS":          "Linux",
			"suffix":      "example.com",
		},
	}

	jobStatus, err := client.CreateDNSReservation(&newDNSRecord)
	Expect(jobStatus).NotTo(BeNil())
	Expect(err).NotTo(HaveOccurred())

	// This should have made three requests:
	// 1+2. Fail to get reservation, get a token
	// 3. Successfully getting the order
	Expect(len(*requests)).To(Equal(3))

	// The last request is the one we care about
	Expect((*requests)[2].URL.Path).To(Equal("/api/v3/onefuse/dnsReservations/"))

	// The CloudBolt Order object should be parsed correctly
	verifyJobStatus(jobStatus)
}

func TestDeleteDNSReservation(t *testing.T) {
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

	jobStatus, err := client.DeleteDNSReservation("10")
	Expect(jobStatus).NotTo(BeNil())
	Expect(err).NotTo(HaveOccurred())

	// This should have made three requests:
	// 1+2. Fail to get reservation, get a token
	// 3. Successfully getting the order
	Expect(len(*requests)).To(Equal(3))

	// The last request is the one we care about
	Expect((*requests)[2].URL.Path).To(Equal("/api/v3/onefuse/dnsReservations/10/"))

	// The CloudBolt Order object should be parsed correctly
	verifyJobStatus(jobStatus)
}
