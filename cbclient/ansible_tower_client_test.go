package cbclient

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestGetAnsibleTowerPolicy(t *testing.T) {
	// Register the test with gomega
	RegisterTestingT(t)

	// Setup mock server with scripted responses
	// Setup requests buffer
	server, requests := mockServer(responsesForAnsibleTowerPolicy)
	Expect(server).NotTo(BeNil())
	Expect(requests).NotTo(BeNil())

	// Setup CloudBolt Client
	client := getClient(server)
	Expect(client).NotTo(BeNil())

	policyName := "My_Ansible_Tower_Policy"
	policy, err := client.GetAnsibleTowerPolicy(policyName)
	Expect(policy).NotTo(BeNil())
	Expect(err).NotTo(HaveOccurred())

	// This should have made three requests:
	// 1+2. Fail to get, get a token
	// 3. Successful Get
	Expect(len(*requests)).To(Equal(3))

	// The last request is the one we care about
	Expect((*requests)[2].URL.Path).To(Equal("/api/v3/onefuse/ansibleTowerPolicies/"))

	// The CloudBolt Order object should be parsed correctly
	Expect(policy.Links.Self.Href).To(Equal("/api/v3/onefuse/ansibleTowerPolicies/6/"))
	Expect(policy.Links.Self.Title).To(Equal("My_Ansible_Tower_Policy"))
	Expect(policy.Links.Workspace.Href).To(Equal("/api/v3/onefuse/workspaces/2/"))
	Expect(policy.Links.Workspace.Title).To(Equal("Default"))
	Expect(policy.Links.Endpoint.Href).To(Equal("/api/v3/onefuse/endpoints/192/"))
	Expect(policy.Links.Endpoint.Title).To(Equal("My_Ansible_Tower_Endpoint"))
	Expect(policy.Name).To(Equal("My_Ansible_Tower_Policy"))
	Expect(policy.ID).To(Equal(6))
	Expect(policy.Description).To(Equal("An Ansible Policy created through automated tests"))
}

func TestGetAnsibleDeployment(t *testing.T) {
	// Register the test with gomega
	RegisterTestingT(t)

	// Setup mock server with scripted responses
	// Setup requests buffer
	server, requests := mockServer(responsesForAnsibleTowerDeoplyment)
	Expect(server).NotTo(BeNil())
	Expect(requests).NotTo(BeNil())

	// Setup CloudBolt Client
	client := getClient(server)
	Expect(client).NotTo(BeNil())

	deploymentPath := "/api/v3/onefuse/ansibleTowerDeployments/4/"
	deployment, err := client.GetAnsibleTowerDeployment(deploymentPath)
	Expect(deployment).NotTo(BeNil())
	Expect(err).NotTo(HaveOccurred())

	// This should have made three requests:
	// 1+2. Fail to get, get a token
	// 3. Successful Get
	Expect(len(*requests)).To(Equal(3))

	// The last request is the one we care about
	Expect((*requests)[2].URL.Path).To(Equal("/api/v3/onefuse/ansibleTowerDeployments/4/"))

	// The CloudBolt Order object should be parsed correctly
	Expect(deployment.Links.Self.Href).To(Equal("/api/v3/onefuse/ansibleTowerDeployments/4/"))
	Expect(deployment.Links.Self.Title).To(Equal("Ansible Tower Deployment id 4"))
	Expect(deployment.Links.Workspace.Href).To(Equal("/api/v3/onefuse/workspaces/2/"))
	Expect(deployment.Links.Workspace.Title).To(Equal("Default"))
	Expect(deployment.Links.Policy.Href).To(Equal("/api/v3/onefuse/ansibleTowerPolicies/1/"))
	Expect(deployment.Links.Policy.Title).To(Equal("atPolicy01"))
	Expect(deployment.ID).To(Equal(4))
}

func TestGetAnsibleDeploymentById(t *testing.T) {
	// Register the test with gomega
	RegisterTestingT(t)

	// Setup mock server with scripted responses
	// Setup requests buffer
	server, requests := mockServer(responsesForAnsibleTowerDeoplyment)
	Expect(server).NotTo(BeNil())
	Expect(requests).NotTo(BeNil())

	// Setup CloudBolt Client
	client := getClient(server)
	Expect(client).NotTo(BeNil())

	deployment, err := client.GetAnsibleTowerDeploymentById("4")
	Expect(deployment).NotTo(BeNil())
	Expect(err).NotTo(HaveOccurred())

	// This should have made three requests:
	// 1+2. Fail to get, get a token
	// 3. Successful Get
	Expect(len(*requests)).To(Equal(3))

	// The last request is the one we care about
	Expect((*requests)[2].URL.Path).To(Equal("/api/v3/onefuse/ansibleTowerDeployments/4/"))

	// The CloudBolt Order object should be parsed correctly
	Expect(deployment.Links.Self.Href).To(Equal("/api/v3/onefuse/ansibleTowerDeployments/4/"))
	Expect(deployment.Links.Self.Title).To(Equal("Ansible Tower Deployment id 4"))
	Expect(deployment.Links.Workspace.Href).To(Equal("/api/v3/onefuse/workspaces/2/"))
	Expect(deployment.Links.Workspace.Title).To(Equal("Default"))
	Expect(deployment.Links.Policy.Href).To(Equal("/api/v3/onefuse/ansibleTowerPolicies/1/"))
	Expect(deployment.Links.Policy.Title).To(Equal("atPolicy01"))
	Expect(deployment.ID).To(Equal(4))
}

func TestDeleteAnsibleDeployment(t *testing.T) {
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

	jobStatus, err := client.DeleteAnsibleTowerDeployment("4")
	Expect(jobStatus).NotTo(BeNil())
	Expect(err).NotTo(HaveOccurred())

	// This should have made three requests:
	// 1+2. Fail to get, get a token
	// 3. Successful Get
	Expect(len(*requests)).To(Equal(3))

	// The last request is the one we care about
	Expect((*requests)[2].URL.Path).To(Equal("/api/v3/onefuse/ansibleTowerDeployments/4/"))

	// The CloudBolt Order object should be parsed correctly
	verifyJobStatus(jobStatus)
}

func TestCreateAnsibleTowerDeployment(t *testing.T) {
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

	newAnsibleTowerDeployment := AnsibleTowerDeployment{
		PolicyID:     1,
		WorkspaceURL: "/api/v3/onefuse/workspaces/2/",
		Hosts:        []string{"host1", "host2"},
		Limit:        "host1",
		TemplateProperties: map[string]interface{}{
			"property1": "value1",
			"property2": "value2",
		},
	}

	jobStatus, err := client.CreateAnsibleTowerDeployment(&newAnsibleTowerDeployment)
	Expect(jobStatus).NotTo(BeNil())
	Expect(err).NotTo(HaveOccurred())

	// This should have made three requests:
	// 1+2. Fail to get, get a token
	// 3. Successful Get
	Expect(len(*requests)).To(Equal(3))

	// The last request is the one we care about
	Expect((*requests)[2].URL.Path).To(Equal("/api/v3/onefuse/ansibleTowerDeployments/"))

	// The CloudBolt Order object should be parsed correctly
	verifyJobStatus(jobStatus)
}
