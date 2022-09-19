package cbclient

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestGetModulePolicy(t *testing.T) {
	// Register the test with gomega
	RegisterTestingT(t)

	// Setup mock server with scripted responses
	// Setup requests buffer
	server, requests := mockServer(responsesForModulePolicy)
	Expect(server).NotTo(BeNil())
	Expect(requests).NotTo(BeNil())

	// Setup CloudBolt Client
	client := getClient(server)
	Expect(client).NotTo(BeNil())

	policyName := "My_Module_Policy"
	policy, err := client.GetModulePolicy(policyName)
	Expect(policy).NotTo(BeNil())
	Expect(err).NotTo(HaveOccurred())

	// This should have made three requests:
	// 1+2. Fail to get, get a token
	// 3. Successful Get
	Expect(len(*requests)).To(Equal(3))

	// The last request is the one we care about
	Expect((*requests)[2].URL.Path).To(Equal("/api/v3/onefuse/modulePolicies/"))

	Expect(policy.Links.Self.Href).To(Equal("/api/v3/onefuse/modulePolicies/180/"))
	Expect(policy.Links.Self.Title).To(Equal("My_Module_Policy"))
	Expect(policy.Links.Workspace.Href).To(Equal("/api/v3/onefuse/workspaces/2/"))
	Expect(policy.Links.Workspace.Title).To(Equal("Default"))
	Expect(policy.Links.Blueprint.Href).To(Equal("/api/v3/onefuse/modules/BP-vjsxfc2z/"))
	Expect(policy.Links.Blueprint.Title).To(Equal("My_Blueprint"))
	Expect(policy.Name).To(Equal("My_Module_Policy"))
	Expect(policy.ID).To(Equal(180))
}

func TestCreateModuleDeployment(t *testing.T) {
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

	moduleDeployment := ModuleDeployment{
		WorkspaceURL: "/api/v3/onefuse/workspaces/2/",
		PolicyID:     180,
		TemplateProperties: map[string]interface{}{
			"title":    "Some Title Here",
			"message":  "Some message here.",
			"to_email": "test@mail.com",
		},
	}

	jobStatus, err := client.CreateModuleDeployment(&moduleDeployment)
	Expect(jobStatus).NotTo(BeNil())
	Expect(err).NotTo(HaveOccurred())

	// This should have made three requests:
	// 1+2. Fail to get, get a token
	// 3. Successful Get
	Expect(len(*requests)).To(Equal(3))

	// The last request is the one we care about
	Expect((*requests)[2].URL.Path).To(Equal("/api/v3/onefuse/moduleManagedObjects/"))
	verifyJobStatus(jobStatus)
}

func TestDeleteModuleDeployment(t *testing.T) {
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

	jobStatus, err := client.DeleteModuleDeployment("15")
	Expect(jobStatus).NotTo(BeNil())
	Expect(err).NotTo(HaveOccurred())

	// This should have made three requests:
	// 1+2. Fail to get, get a token
	// 3. Successful Get
	Expect(len(*requests)).To(Equal(3))

	// The last request is the one we care about
	Expect((*requests)[2].URL.Path).To(Equal("/api/v3/onefuse/moduleManagedObjects/15/"))
	verifyJobStatus(jobStatus)
}

func TestGetModuleDeployment(t *testing.T) {
	// Register the test with gomega
	RegisterTestingT(t)

	// Setup mock server with scripted responses
	// Setup requests buffer
	server, requests := mockServer(responsesForModuleDeployment)
	Expect(server).NotTo(BeNil())
	Expect(requests).NotTo(BeNil())

	// Setup CloudBolt Client
	client := getClient(server)
	Expect(client).NotTo(BeNil())

	moduleDeploymentPath := "/api/v3/onefuse/moduleManagedObjects/75/"
	moduleDeployment, err := client.GetModuleDeployment(moduleDeploymentPath)
	Expect(moduleDeployment).NotTo(BeNil())
	Expect(err).NotTo(HaveOccurred())

	// This should have made three requests:
	// 1+2. Fail to get, get a token
	// 3. Successful Get
	Expect(len(*requests)).To(Equal(3))

	// The last request is the one we care about
	Expect((*requests)[2].URL.Path).To(Equal("/api/v3/onefuse/moduleManagedObjects/75/"))

	// The CloudBolt Order object should be parsed correctly
	Expect(moduleDeployment.Links.Self.Href).To(Equal("/api/v3/onefuse/moduleManagedObjects/75/"))
	Expect(moduleDeployment.Links.Self.Title).To(Equal("Module Managed Object id 75"))
	Expect(moduleDeployment.Links.Workspace.Href).To(Equal("/api/v3/onefuse/workspaces/1/"))
	Expect(moduleDeployment.Links.Workspace.Title).To(Equal("Default"))
	Expect(moduleDeployment.Links.Policy.Href).To(Equal("/api/v3/onefuse/modulePolicies/1/"))
	Expect(moduleDeployment.Links.Policy.Title).To(Equal("automated_module_policy"))
	Expect(moduleDeployment.Links.JobMetadata.Href).To(Equal("/api/v3/onefuse/jobMetadata/662/"))
	Expect(moduleDeployment.Links.JobMetadata.Title).To(Equal("Job Metadata Record id 662"))
	Expect(moduleDeployment.ID).To(Equal(75))
	Expect(moduleDeployment.Name).To(Equal("vip-dev-ap012 (member pp-atltlap004.example.com)"))
	Expect(len(moduleDeployment.ProvisioningJobResults)).To(Equal(1))
	Expect(len(moduleDeployment.DeprovisioningJobResults)).To(Equal(0))
	Expect(moduleDeployment.Archived).To(Equal(false))
}

func TestGetModuleDeploymentId(t *testing.T) {
	// Register the test with gomega
	RegisterTestingT(t)

	// Setup mock server with scripted responses
	// Setup requests buffer
	server, requests := mockServer(responsesForModuleDeployment)
	Expect(server).NotTo(BeNil())
	Expect(requests).NotTo(BeNil())

	// Setup CloudBolt Client
	client := getClient(server)
	Expect(client).NotTo(BeNil())

	moduleDeployment, err := client.GetModuleDeploymentById("75")
	Expect(moduleDeployment).NotTo(BeNil())
	Expect(err).NotTo(HaveOccurred())

	// This should have made three requests:
	// 1+2. Fail to get, get a token
	// 3. Successful Get
	Expect(len(*requests)).To(Equal(3))

	// The last request is the one we care about
	Expect((*requests)[2].URL.Path).To(Equal("/api/v3/onefuse/moduleManagedObjects/75/"))

	// The CloudBolt Order object should be parsed correctly
	Expect(moduleDeployment.Links.Self.Href).To(Equal("/api/v3/onefuse/moduleManagedObjects/75/"))
	Expect(moduleDeployment.Links.Self.Title).To(Equal("Module Managed Object id 75"))
	Expect(moduleDeployment.Links.Workspace.Href).To(Equal("/api/v3/onefuse/workspaces/1/"))
	Expect(moduleDeployment.Links.Workspace.Title).To(Equal("Default"))
	Expect(moduleDeployment.Links.Policy.Href).To(Equal("/api/v3/onefuse/modulePolicies/1/"))
	Expect(moduleDeployment.Links.Policy.Title).To(Equal("automated_module_policy"))
	Expect(moduleDeployment.Links.JobMetadata.Href).To(Equal("/api/v3/onefuse/jobMetadata/662/"))
	Expect(moduleDeployment.Links.JobMetadata.Title).To(Equal("Job Metadata Record id 662"))
	Expect(moduleDeployment.ID).To(Equal(75))
	Expect(moduleDeployment.Name).To(Equal("vip-dev-ap012 (member pp-atltlap004.example.com)"))
	Expect(len(moduleDeployment.ProvisioningJobResults)).To(Equal(1))
	Expect(len(moduleDeployment.DeprovisioningJobResults)).To(Equal(0))
	Expect(moduleDeployment.Archived).To(Equal(false))
}
