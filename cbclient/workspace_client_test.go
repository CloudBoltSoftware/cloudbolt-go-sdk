package cbclient

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestGetWorkspace(t *testing.T) {
	// Register the test with gomega
	RegisterTestingT(t)

	// Setup mock server with scripted responses
	// Setup requests buffer
	server, requests := mockServer(responsesForWorkspaceList)
	Expect(server).NotTo(BeNil())
	Expect(requests).NotTo(BeNil())

	// Setup CloudBolt Client
	client := getClient(server)
	Expect(client).NotTo(BeNil())

	workspaceName := "Default"
	workspace, err := client.GetWorkSpace(workspaceName)
	Expect(workspace).NotTo(BeNil())
	Expect(err).NotTo(HaveOccurred())

	// This should have made three requests:
	// 1+2. Fail to get, get a token
	// 3. Successful Get
	Expect(len(*requests)).To(Equal(3))

	// The last request is the one we care about
	Expect((*requests)[2].URL.Path).To(Equal("/api/v3/onefuse/workspaces/"))

	// The CloudBolt Order object should be parsed correctly
	Expect(workspace.Links.Self.Href).To(Equal("/api/v3/onefuse/workspaces/2/"))
	Expect(workspace.Links.Self.Title).To(Equal("Default"))
	Expect(workspace.Name).To(Equal("Default"))
	Expect(workspace.ID).To(Equal(2))
}

func TestGetDefaultWorkspace(t *testing.T) {
	// Register the test with gomega
	RegisterTestingT(t)

	// Setup mock server with scripted responses
	// Setup requests buffer
	server, requests := mockServer(responsesForWorkspaceList)
	Expect(server).NotTo(BeNil())
	Expect(requests).NotTo(BeNil())

	// Setup CloudBolt Client
	client := getClient(server)
	Expect(client).NotTo(BeNil())

	workspace, err := client.GetDefaultWorkSpace()
	Expect(workspace).NotTo(BeNil())
	Expect(err).NotTo(HaveOccurred())

	// This should have made three requests:
	// 1+2. Fail to get, get a token
	// 3. Successful Get
	Expect(len(*requests)).To(Equal(3))

	// The last request is the one we care about
	Expect((*requests)[2].URL.Path).To(Equal("/api/v3/onefuse/workspaces/"))

	// The CloudBolt Order object should be parsed correctly
	Expect(workspace.Links.Self.Href).To(Equal("/api/v3/onefuse/workspaces/2/"))
	Expect(workspace.Links.Self.Title).To(Equal("Default"))
	Expect(workspace.Name).To(Equal("Default"))
	Expect(workspace.ID).To(Equal(2))
}
