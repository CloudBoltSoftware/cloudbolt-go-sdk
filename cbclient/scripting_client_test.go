package cbclient

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestGetScriptingPolicy(t *testing.T) {
	// Register the test with gomega
	RegisterTestingT(t)

	// Setup mock server with scripted responses
	// Setup requests buffer
	server, requests := mockServer(responsesForScriptingPolicy)
	Expect(server).NotTo(BeNil())
	Expect(requests).NotTo(BeNil())

	// Setup CloudBolt Client
	client := getClient(server)
	Expect(client).NotTo(BeNil())

	policyName := "My_Scripting_Policy"
	policy, err := client.GetScriptingPolicy(policyName)
	Expect(policy).NotTo(BeNil())
	Expect(err).NotTo(HaveOccurred())

	// This should have made three requests:
	// 1+2. Fail to get, get a token
	// 3. Successful Get
	Expect(len(*requests)).To(Equal(3))

	// The last request is the one we care about
	Expect((*requests)[2].URL.Path).To(Equal("/api/v3/onefuse/scriptingPolicies/"))

	// The CloudBolt Order object should be parsed correctly
	Expect(policy.Links.Self.Href).To(Equal("/api/v3/onefuse/scriptingPolicies/901/"))
	Expect(policy.Links.Self.Title).To(Equal("My_Scripting_Policy"))
	Expect(policy.Links.Workspace.Href).To(Equal("/api/v3/onefuse/workspaces/2/"))
	Expect(policy.Links.Workspace.Title).To(Equal("Default"))
	Expect(policy.Name).To(Equal("My_Scripting_Policy"))
	Expect(policy.ID).To(Equal(901))
	Expect(policy.Description).To(Equal("A Scripting Policy created through automated tests"))
}

func TestCreateScriptingDeployment(t *testing.T) {
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

	newScriptingDeployment := ScriptingDeployment{
		PolicyID:     1,
		WorkspaceURL: "/api/v3/onefuse/workspaces/1/",
		TemplateProperties: map[string]interface{}{
			"ownerScript": "john.doe@example.com",
			"Environment": "production",
			"OS":          "Linux",
			"Application": "Web Servers",
			"suffix":      "example.com",
		},
	}

	jobStatus, err := client.CreateScriptingDeployment(&newScriptingDeployment)
	Expect(jobStatus).NotTo(BeNil())
	Expect(err).NotTo(HaveOccurred())

	// This should have made three requests:
	// 1+2. Fail to get, get a token
	// 3. Successful Get
	Expect(len(*requests)).To(Equal(3))

	// The last request is the one we care about
	Expect((*requests)[2].URL.Path).To(Equal("/api/v3/onefuse/scriptingDeployments/"))

	// The CloudBolt Order object should be parsed correctly
	verifyJobStatus(jobStatus)
}

func TestGetScriptingDeployment(t *testing.T) {
	// Register the test with gomega
	RegisterTestingT(t)

	// Setup mock server with scripted responses
	// Setup requests buffer
	server, requests := mockServer(responsesForScriptingDeployment)
	Expect(server).NotTo(BeNil())
	Expect(requests).NotTo(BeNil())

	// Setup CloudBolt Client
	client := getClient(server)
	Expect(client).NotTo(BeNil())

	scriptingDeploymentPath := "/api/v3/onefuse/scriptingDeployments/67/"
	scriptingDeployment, err := client.GetScriptingDeployment(scriptingDeploymentPath)
	Expect(scriptingDeployment).NotTo(BeNil())
	Expect(err).NotTo(HaveOccurred())

	// This should have made three requests:
	// 1+2. Fail to get, get a token
	// 3. Successful Get
	Expect(len(*requests)).To(Equal(3))

	// The last request is the one we care about
	Expect((*requests)[2].URL.Path).To(Equal("/api/v3/onefuse/scriptingDeployments/67/"))

	// The CloudBolt Order object should be parsed correctly
	Expect(scriptingDeployment.Links.Self.Href).To(Equal("/api/v3/onefuse/scriptingDeployments/67/"))
	Expect(scriptingDeployment.Links.Self.Title).To(Equal("Scripting Deployment id 67"))
	Expect(scriptingDeployment.Links.Workspace.Href).To(Equal("/api/v3/onefuse/workspaces/1/"))
	Expect(scriptingDeployment.Links.Workspace.Title).To(Equal("Default"))
	Expect(scriptingDeployment.Links.Policy.Href).To(Equal("/api/v3/onefuse/scriptingPolicies/1/"))
	Expect(scriptingDeployment.Links.Policy.Title).To(Equal("qalnxtst4_script"))
	Expect(scriptingDeployment.Links.JobMetadata.Href).To(Equal("/api/v3/onefuse/jobMetadata/650/"))
	Expect(scriptingDeployment.Links.JobMetadata.Title).To(Equal("Job Metadata Record id 650"))
	Expect(scriptingDeployment.ID).To(Equal(67))
	Expect(scriptingDeployment.Hostname).To(Equal("qalnxtst4.sovlabs.net"))
	Expect(scriptingDeployment.ProvisioningDetails).NotTo(BeNil())
	Expect(scriptingDeployment.DeprovisioningDetails).NotTo(BeNil())
	Expect(scriptingDeployment.Archived).To(Equal(false))
}

func TestGetScriptingDeploymentById(t *testing.T) {
	// Register the test with gomega
	RegisterTestingT(t)

	// Setup mock server with scripted responses
	// Setup requests buffer
	server, requests := mockServer(responsesForScriptingDeployment)
	Expect(server).NotTo(BeNil())
	Expect(requests).NotTo(BeNil())

	// Setup CloudBolt Client
	client := getClient(server)
	Expect(client).NotTo(BeNil())

	scriptingDeployment, err := client.GetScriptingDeploymentById("67")
	Expect(scriptingDeployment).NotTo(BeNil())
	Expect(err).NotTo(HaveOccurred())

	// This should have made three requests:
	// 1+2. Fail to get, get a token
	// 3. Successful Get
	Expect(len(*requests)).To(Equal(3))

	// The last request is the one we care about
	Expect((*requests)[2].URL.Path).To(Equal("/api/v3/onefuse/scriptingDeployments/67/"))

	// The CloudBolt Order object should be parsed correctly
	Expect(scriptingDeployment.Links.Self.Href).To(Equal("/api/v3/onefuse/scriptingDeployments/67/"))
	Expect(scriptingDeployment.Links.Self.Title).To(Equal("Scripting Deployment id 67"))
	Expect(scriptingDeployment.Links.Workspace.Href).To(Equal("/api/v3/onefuse/workspaces/1/"))
	Expect(scriptingDeployment.Links.Workspace.Title).To(Equal("Default"))
	Expect(scriptingDeployment.Links.Policy.Href).To(Equal("/api/v3/onefuse/scriptingPolicies/1/"))
	Expect(scriptingDeployment.Links.Policy.Title).To(Equal("qalnxtst4_script"))
	Expect(scriptingDeployment.Links.JobMetadata.Href).To(Equal("/api/v3/onefuse/jobMetadata/650/"))
	Expect(scriptingDeployment.Links.JobMetadata.Title).To(Equal("Job Metadata Record id 650"))
	Expect(scriptingDeployment.ID).To(Equal(67))
	Expect(scriptingDeployment.Hostname).To(Equal("qalnxtst4.sovlabs.net"))
	Expect(scriptingDeployment.ProvisioningDetails).NotTo(BeNil())
	Expect(scriptingDeployment.DeprovisioningDetails).NotTo(BeNil())
	Expect(scriptingDeployment.Archived).To(Equal(false))
}

func TestDeleteScriptingDeployment(t *testing.T) {
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

	jobStatus, err := client.DeleteScriptingDeployment("67")
	Expect(jobStatus).NotTo(BeNil())
	Expect(err).NotTo(HaveOccurred())

	// This should have made three requests:
	// 1+2. Fail to get, get a token
	// 3. Successful Get
	Expect(len(*requests)).To(Equal(3))

	// The last request is the one we care about
	Expect((*requests)[2].URL.Path).To(Equal("/api/v3/onefuse/scriptingDeployments/67/"))

	// The CloudBolt Order object should be parsed correctly
	verifyJobStatus(jobStatus)
}
