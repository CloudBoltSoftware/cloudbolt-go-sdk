package cbclient

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestGetMicrosoftEndpoint(t *testing.T) {
	// Register the test with gomega
	RegisterTestingT(t)

	// Setup mock server with scripted responses
	// Setup requests buffer
	server, requests := mockServer(responsesForMicrosoftEndpoint)
	Expect(server).NotTo(BeNil())
	Expect(requests).NotTo(BeNil())

	// Setup CloudBolt Client
	client := getClient(server)
	Expect(client).NotTo(BeNil())

	endpointName := "My_Microsoft_Endpoint"
	endpoint, err := client.GetMicrosoftEndpoint(endpointName)
	Expect(endpoint).NotTo(BeNil())
	Expect(err).NotTo(HaveOccurred())

	// This should have made three requests:
	// 1+2. Fail to get endpoint, get a token
	// 3. Successfully getting the endpoint
	Expect(len(*requests)).To(Equal(3))

	// The last request is the one we care about
	Expect((*requests)[2].URL.Path).To(Equal("/api/v3/onefuse/endpoints/"))

	// The CloudBolt Order object should be parsed correctly
	Expect(endpoint.Links.Self.Href).To(Equal("/api/v3/onefuse/endpoints/7/"))
	Expect(endpoint.Links.Self.Title).To(Equal("My_Microsoft_Endpoint"))
	Expect(endpoint.Links.Workspace.Href).To(Equal("/api/v3/onefuse/workspaces/2/"))
	Expect(endpoint.Links.Workspace.Title).To(Equal("Default"))
	Expect(endpoint.Links.Credential.Href).To(Equal("/api/v3/onefuse/moduleCredentials/11/"))
	Expect(endpoint.Links.Credential.Title).To(Equal("My_Credentials"))
	Expect(endpoint.Name).To(Equal("My_Microsoft_Endpoint"))
	Expect(endpoint.ID).To(Equal(7))
	Expect(endpoint.Description).To(Equal("A Microsoft Endpoint"))
	Expect(endpoint.Host).To(Equal("microsoftqa01.mydomain.net"))
	Expect(endpoint.Port).To(Equal(443))
	Expect(endpoint.SSL).To(Equal(true))

}
