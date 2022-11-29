package cbclient

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestGetStaticPropertySet(t *testing.T) {
	// Register the test with gomega
	RegisterTestingT(t)

	// Setup mock server with scripted responses
	// Setup requests buffer
	server, requests := mockServer(responsesForStaticPropertySet)
	Expect(server).NotTo(BeNil())
	Expect(requests).NotTo(BeNil())

	// Setup CloudBolt Client
	client := getClient(server)
	Expect(client).NotTo(BeNil())

	propertySetName := "My_Static_Property_Set"
	propertySet, err := client.GetStaticPropertySet(propertySetName)
	Expect(propertySet).NotTo(BeNil())
	Expect(err).NotTo(HaveOccurred())

	// This should have made three requests:
	// 1+2. Fail to get, get a token
	// 3. Successful Get
	Expect(len(*requests)).To(Equal(3))

	// The last request is the one we care about
	Expect((*requests)[2].URL.Path).To(Equal("/api/v3/onefuse/propertySets/"))

	// The CloudBolt Order object should be parsed correctly
	Expect(propertySet.Links.Self.Href).To(Equal("/api/v3/onefuse/propertySets/168/"))
	Expect(propertySet.Links.Self.Title).To(Equal("My_Static_Property_Set"))
	Expect(propertySet.Links.Workspace.Href).To(Equal("/api/v3/onefuse/workspaces/2/"))
	Expect(propertySet.Links.Workspace.Title).To(Equal("Default"))
	Expect(propertySet.Name).To(Equal("My_Static_Property_Set"))
	Expect(propertySet.ID).To(Equal(168))
	Expect(propertySet.Description).To(Equal("My Static Property Set"))
	Expect(propertySet.Properties["product"]).To(Equal("OneFuse"))
	Expect(propertySet.Properties["organization"]).To(Equal("CloudBolt Software"))
}
