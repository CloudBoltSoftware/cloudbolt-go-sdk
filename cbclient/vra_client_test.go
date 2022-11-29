package cbclient

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestGetVraPolicy(t *testing.T) {
	// Register the test with gomega
	RegisterTestingT(t)

	// Setup mock server with scripted responses
	// Setup requests buffer
	server, requests := mockServer(responsesForVraPolicy)
	Expect(server).NotTo(BeNil())
	Expect(requests).NotTo(BeNil())

	// Setup CloudBolt Client
	client := getClient(server)
	Expect(client).NotTo(BeNil())

	policyName := "My_Vra_Policy"
	policy, err := client.GetVraPolicy(policyName)
	Expect(policy).NotTo(BeNil())
	Expect(err).NotTo(HaveOccurred())

	// This should have made three requests:
	// 1+2. Fail to get, get a token
	// 3. Successful Get
	Expect(len(*requests)).To(Equal(3))

	// The last request is the one we care about
	Expect((*requests)[2].URL.Path).To(Equal("/api/v3/onefuse/vraPolicies/"))

	// The CloudBolt Order object should be parsed correctly
	Expect(policy.Links.Self.Href).To(Equal("/api/v3/onefuse/vraPolicies/1/"))
	Expect(policy.Links.Self.Title).To(Equal("My_Vra_Policy"))
	Expect(policy.Links.Workspace.Href).To(Equal("/api/v3/onefuse/workspaces/2/"))
	Expect(policy.Links.Workspace.Title).To(Equal("Default"))
	Expect(policy.Links.Endpoint.Href).To(Equal("/api/v3/onefuse/endpoints/27/"))
	Expect(policy.Links.Endpoint.Title).To(Equal("My_Vra_Policy_Endpoint"))
	Expect(policy.Name).To(Equal("My_Vra_Policy"))
	Expect(policy.ID).To(Equal(1))
	Expect(policy.Description).To(Equal("A Vra Policy"))
}

func TestGetVraDeploymentById(t *testing.T) {
	// Register the test with gomega
	RegisterTestingT(t)

	// Setup mock server with scripted responses
	// Setup requests buffer
	server, requests := mockServer(responsesForVraDeployment)
	Expect(server).NotTo(BeNil())
	Expect(requests).NotTo(BeNil())

	// Setup CloudBolt Client
	client := getClient(server)
	Expect(client).NotTo(BeNil())

	vraDeployment, err := client.GetVraDeploymentById("1")
	Expect(vraDeployment).NotTo(BeNil())
	Expect(err).NotTo(HaveOccurred())

	// This should have made three requests:
	// 1+2. Fail to get, get a token
	// 3. Successful Get
	Expect(len(*requests)).To(Equal(3))

	// The last request is the one we care about
	Expect((*requests)[2].URL.Path).To(Equal("/api/v3/onefuse/vraDeployments/1/"))

	// The CloudBolt Order object should be parsed correctly
	Expect(vraDeployment.Links.Self.Href).To(Equal("/api/v3/onefuse/vraDeployments/1/"))
	Expect(vraDeployment.Links.Self.Title).To(Equal("vRealize Automation Deployment id 1"))
	Expect(vraDeployment.Links.Workspace.Href).To(Equal("/api/v3/onefuse/workspaces/1/"))
	Expect(vraDeployment.Links.Workspace.Title).To(Equal("Default"))
	Expect(vraDeployment.Links.Policy.Href).To(Equal("/api/v3/onefuse/vraPolicies/1/"))
	Expect(vraDeployment.Links.Policy.Title).To(Equal("Vra_Policy_1"))
	Expect(vraDeployment.Links.JobMetadata.Href).To(Equal("/api/v3/onefuse/jobMetadata/1/"))
	Expect(vraDeployment.Links.JobMetadata.Title).To(Equal("Job Metadata Record id 1"))
	Expect(vraDeployment.ID).To(Equal(1))
	Expect(vraDeployment.BlueprintName).To(Equal("vRA Blueprint Name"))
	Expect(vraDeployment.ProjectName).To(Equal("vRA Project Name"))
	Expect(vraDeployment.Archived).To(Equal(false))
	Expect(vraDeployment.DeploymentName).NotTo(BeNil())
}

func TestGetVraDeployment(t *testing.T) {
	// Register the test with gomega
	RegisterTestingT(t)

	// Setup mock server with scripted responses
	// Setup requests buffer
	server, requests := mockServer(responsesForVraDeployment)
	Expect(server).NotTo(BeNil())
	Expect(requests).NotTo(BeNil())

	// Setup CloudBolt Client
	client := getClient(server)
	Expect(client).NotTo(BeNil())

	vraDeploymentPath := "/api/v3/onefuse/vraDeployments/1/"
	vraDeployment, err := client.GetVraDeployment(vraDeploymentPath)
	Expect(vraDeployment).NotTo(BeNil())
	Expect(err).NotTo(HaveOccurred())

	// This should have made three requests:
	// 1+2. Fail to get, get a token
	// 3. Successful Get
	Expect(len(*requests)).To(Equal(3))

	// The last request is the one we care about
	Expect((*requests)[2].URL.Path).To(Equal("/api/v3/onefuse/vraDeployments/1/"))

	// The CloudBolt Order object should be parsed correctly
	Expect(vraDeployment.Links.Self.Href).To(Equal("/api/v3/onefuse/vraDeployments/1/"))
	Expect(vraDeployment.Links.Self.Title).To(Equal("vRealize Automation Deployment id 1"))
	Expect(vraDeployment.Links.Workspace.Href).To(Equal("/api/v3/onefuse/workspaces/1/"))
	Expect(vraDeployment.Links.Workspace.Title).To(Equal("Default"))
	Expect(vraDeployment.Links.Policy.Href).To(Equal("/api/v3/onefuse/vraPolicies/1/"))
	Expect(vraDeployment.Links.Policy.Title).To(Equal("Vra_Policy_1"))
	Expect(vraDeployment.Links.JobMetadata.Href).To(Equal("/api/v3/onefuse/jobMetadata/1/"))
	Expect(vraDeployment.Links.JobMetadata.Title).To(Equal("Job Metadata Record id 1"))
	Expect(vraDeployment.ID).To(Equal(1))
	Expect(vraDeployment.BlueprintName).To(Equal("vRA Blueprint Name"))
	Expect(vraDeployment.ProjectName).To(Equal("vRA Project Name"))
	Expect(vraDeployment.Archived).To(Equal(false))
	Expect(vraDeployment.DeploymentName).NotTo(BeNil())
}

func TestCreateVraDeployment(t *testing.T) {
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

	newVraDeployment := VraDeployment{
		PolicyID:           1,
		WorkspaceURL:       "/api/v3/onefuse/workspaces/1/",
		DeploymentName:     "My Vra Deployment",
		TemplateProperties: nil,
	}

	jobStatus, err := client.CreateVraDeployment(&newVraDeployment)
	Expect(jobStatus).NotTo(BeNil())
	Expect(err).NotTo(HaveOccurred())

	// This should have made three requests:
	// 1+2. Fail to get, get a token
	// 3. Successful Get
	Expect(len(*requests)).To(Equal(3))

	// The last request is the one we care about
	Expect((*requests)[2].URL.Path).To(Equal("/api/v3/onefuse/vraDeployments/"))

	// The CloudBolt Order object should be parsed correctly
	verifyJobStatus(jobStatus)
}

func TestDeleteVraDeployment(t *testing.T) {
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

	jobStatus, err := client.DeleteVraDeployment("1")
	Expect(jobStatus).NotTo(BeNil())
	Expect(err).NotTo(HaveOccurred())

	// This should have made three requests:
	// 1+2. Fail to get, get a token
	// 3. Successful Get
	Expect(len(*requests)).To(Equal(3))

	// The last request is the one we care about
	Expect((*requests)[2].URL.Path).To(Equal("/api/v3/onefuse/vraDeployments/1/"))

	// The CloudBolt Order object should be parsed correctly
	verifyJobStatus(jobStatus)
}
