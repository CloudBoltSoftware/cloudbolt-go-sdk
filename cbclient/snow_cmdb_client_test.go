package cbclient

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestGetServiceNowCMDBPolicy(t *testing.T) {
	// Register the test with gomega
	RegisterTestingT(t)

	// Setup mock server with scripted responses
	// Setup requests buffer
	server, requests := mockServer(responsesForServiceNowCMDBPolicy)
	Expect(server).NotTo(BeNil())
	Expect(requests).NotTo(BeNil())

	// Setup CloudBolt Client
	client := getClient(server)
	Expect(client).NotTo(BeNil())

	policyName := "My_SNOW_CMDB_Policy"
	policy, err := client.GetServiceNowCMDBPolicy(policyName)
	Expect(policy).NotTo(BeNil())
	Expect(err).NotTo(HaveOccurred())

	// This should have made three requests:
	// 1+2. Fail to get, get a token
	// 3. Successful Get
	Expect(len(*requests)).To(Equal(3))

	// The last request is the one we care about
	Expect((*requests)[2].URL.Path).To(Equal("/api/v3/onefuse/servicenowCMDBPolicies/"))

	// The CloudBolt Order object should be parsed correctly
	Expect(policy.Links.Self.Href).To(Equal("/api/v3/onefuse/servicenowCMDBPolicies/224/"))
	Expect(policy.Links.Self.Title).To(Equal("My_SNOW_CMDB_Policy"))
	Expect(policy.Links.Workspace.Href).To(Equal("/api/v3/onefuse/workspaces/2/"))
	Expect(policy.Links.Workspace.Title).To(Equal("Default"))
	Expect(policy.Links.Endpoint.Href).To(Equal("/api/v3/onefuse/endpoints/7723/"))
	Expect(policy.Links.Endpoint.Title).To(Equal("My_SNOW_CMDB_Policy_Endpoint"))
	Expect(policy.Name).To(Equal("My_SNOW_CMDB_Policy"))
	Expect(policy.ID).To(Equal(224))
	Expect(policy.Description).To(Equal("A ServiceNow CMDB Policy"))
}

func TestGetServicenowCMDBDeploymentById(t *testing.T) {
	// Register the test with gomega
	RegisterTestingT(t)

	// Setup mock server with scripted responses
	// Setup requests buffer
	server, requests := mockServer(responsesForServiceNowCMDBDeployment)
	Expect(server).NotTo(BeNil())
	Expect(requests).NotTo(BeNil())

	// Setup CloudBolt Client
	client := getClient(server)
	Expect(client).NotTo(BeNil())

	snowDeployment, err := client.GetServicenowCMDBDeploymentById("24")
	Expect(snowDeployment).NotTo(BeNil())
	Expect(err).NotTo(HaveOccurred())

	// This should have made three requests:
	// 1+2. Fail to get, get a token
	// 3. Successful Get
	Expect(len(*requests)).To(Equal(3))

	// The last request is the one we care about
	Expect((*requests)[2].URL.Path).To(Equal("/api/v3/onefuse/servicenowCMDBDeployments/24/"))

	// The CloudBolt Order object should be parsed correctly
	Expect(snowDeployment.Links.Self.Href).To(Equal("/api/v3/onefuse/servicenowCMDBDeployments/24/"))
	Expect(snowDeployment.Links.Self.Title).To(Equal("Service Now Deployment id 24"))
	Expect(snowDeployment.Links.Workspace.Href).To(Equal("/api/v3/onefuse/workspaces/1/"))
	Expect(snowDeployment.Links.Workspace.Title).To(Equal("Default"))
	Expect(snowDeployment.Links.Policy.Href).To(Equal("/api/v3/onefuse/servicenowCMDBPolicies/4/"))
	Expect(snowDeployment.Links.Policy.Title).To(Equal("test_serviceNow_CMDB_policy_updated_581"))
	Expect(snowDeployment.Links.JobMetadata.Href).To(Equal("/api/v3/onefuse/jobMetadata/434/"))
	Expect(snowDeployment.Links.JobMetadata.Title).To(Equal("Job Metadata Record id 434"))
	Expect(snowDeployment.ID).To(Equal(24))
}

func TestGetServicenowCMDBDeployment(t *testing.T) {
	// Register the test with gomega
	RegisterTestingT(t)

	// Setup mock server with scripted responses
	// Setup requests buffer
	server, requests := mockServer(responsesForServiceNowCMDBDeployment)
	Expect(server).NotTo(BeNil())
	Expect(requests).NotTo(BeNil())

	// Setup CloudBolt Client
	client := getClient(server)
	Expect(client).NotTo(BeNil())

	snowDeploymentPath := "/api/v3/onefuse/servicenowCMDBDeployments/24/"
	snowDeployment, err := client.GetServicenowCMDBDeployment(snowDeploymentPath)
	Expect(snowDeployment).NotTo(BeNil())
	Expect(err).NotTo(HaveOccurred())

	// This should have made three requests:
	// 1+2. Fail to get, get a token
	// 3. Successful Get
	Expect(len(*requests)).To(Equal(3))

	// The last request is the one we care about
	Expect((*requests)[2].URL.Path).To(Equal("/api/v3/onefuse/servicenowCMDBDeployments/24/"))

	// The CloudBolt Order object should be parsed correctly
	Expect(snowDeployment.Links.Self.Href).To(Equal("/api/v3/onefuse/servicenowCMDBDeployments/24/"))
	Expect(snowDeployment.Links.Self.Title).To(Equal("Service Now Deployment id 24"))
	Expect(snowDeployment.Links.Workspace.Href).To(Equal("/api/v3/onefuse/workspaces/1/"))
	Expect(snowDeployment.Links.Workspace.Title).To(Equal("Default"))
	Expect(snowDeployment.Links.Policy.Href).To(Equal("/api/v3/onefuse/servicenowCMDBPolicies/4/"))
	Expect(snowDeployment.Links.Policy.Title).To(Equal("test_serviceNow_CMDB_policy_updated_581"))
	Expect(snowDeployment.Links.JobMetadata.Href).To(Equal("/api/v3/onefuse/jobMetadata/434/"))
	Expect(snowDeployment.Links.JobMetadata.Title).To(Equal("Job Metadata Record id 434"))
	Expect(snowDeployment.ID).To(Equal(24))
}

func TestCreateServicenowCMDBDeployment(t *testing.T) {
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

	newServicenowCMDBDeployment := ServicenowCMDBDeployment{
		PolicyID:     1,
		WorkspaceURL: "/api/v3/onefuse/workspaces/1/",
		TemplateProperties: map[string]interface{}{
			"className":      "cmdb_ci_linux_server",
			"sys_class_name": "cmdb_ci_linux_server",
			"name":           "host14",
			"internal_id":    "uuid-for-host14",
			"host_name":      "host14.example.net",
			"nic0_ip":        "127.0.1.14",
			"name2":          "host36",
			"internal_id2":   "uuid-for-host35",
			"nic0_ip2":       "127.0.1.36",
			"host_name2":     "host36.example.net",
			"os_name":        "linux",
			"state":          "ON",
			"disk_size":      "800",
			"disk_size2":     "700",
		},
	}

	jobStatus, err := client.CreateServicenowCMDBDeployment(&newServicenowCMDBDeployment)
	Expect(jobStatus).NotTo(BeNil())
	Expect(err).NotTo(HaveOccurred())

	// This should have made three requests:
	// 1+2. Fail to get, get a token
	// 3. Successful Get
	Expect(len(*requests)).To(Equal(3))

	// The last request is the one we care about
	Expect((*requests)[2].URL.Path).To(Equal("/api/v3/onefuse/servicenowCMDBDeployments/"))

	// The CloudBolt Order object should be parsed correctly
	verifyJobStatus(jobStatus)
}

func TestDeleteServicenowCMDBDeployment(t *testing.T) {
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

	jobStatus, err := client.DeleteServicenowCMDBDeployment("24")
	Expect(jobStatus).NotTo(BeNil())
	Expect(err).NotTo(HaveOccurred())

	// This should have made three requests:
	// 1+2. Fail to get, get a token
	// 3. Successful Get
	Expect(len(*requests)).To(Equal(3))

	// The last request is the one we care about
	Expect((*requests)[2].URL.Path).To(Equal("/api/v3/onefuse/servicenowCMDBDeployments/24/"))

	// The CloudBolt Order object should be parsed correctly
	verifyJobStatus(jobStatus)
}
