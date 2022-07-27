package cbclient

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestGetResourceHandler(t *testing.T) {
	// Register the test with gomega
	RegisterTestingT(t)

	// Setup mock server with scripted responses
	// Setup requests buffer
	server, requests := mockServer(responsesForResourceHandler)
	Expect(server).NotTo(BeNil())
	Expect(requests).NotTo(BeNil())

	// Setup CloudBolt Client
	client := getClient(server)
	Expect(client).NotTo(BeNil())

	rhName := "My Test Resource Handler"
	rh, err := client.GetResourceHandler(rhName)
	Expect(rh).NotTo(BeNil())
	Expect(err).NotTo(HaveOccurred())

	// This should have made three requests:
	// 1+2. Fail to get RH, get a token
	// 3. Successfully getting the order
	Expect(len(*requests)).To(Equal(3))

	// The last request is the one we care about
	Expect((*requests)[2].URL.Path).To(Equal("/api/v3/cmp/resourceHandlers/"))

	// The CloudBolt Order object should be parsed correctly
	Expect(rh.Links.Self.Href).To(Equal("/api/v3/cmp/resourceHandlers/RH-nza16uyn/"))
	Expect(rh.Links.Self.Title).To(Equal("My Test Resource Handler"))
	Expect(rh.Name).To(Equal("My Test Resource Handler"))
	Expect(rh.ID).To(Equal("RH-nza16uyn"))
}

func TestGetResourceHandlerById(t *testing.T) {
	// Register the test with gomega
	RegisterTestingT(t)

	// Setup mock server with scripted responses
	// Setup requests buffer
	server, requests := mockServer(responsesForResourceHandlerById)
	Expect(server).NotTo(BeNil())
	Expect(requests).NotTo(BeNil())

	// Setup CloudBolt Client
	client := getClient(server)
	Expect(client).NotTo(BeNil())

	rhId := "RH-nza16uyn"
	rh, err := client.GetResourceHandlerById(rhId)
	Expect(rh).NotTo(BeNil())
	Expect(err).NotTo(HaveOccurred())

	// This should have made three requests:
	// 1+2. Fail to get order, get a token
	// 3. Successfully getting the order
	Expect(len(*requests)).To(Equal(3))

	// The last request is the one we care about
	Expect((*requests)[2].URL.Path).To(Equal("/api/v3/cmp/resourceHandlers/RH-nza16uyn/"))

	// The CloudBolt Order object should be parsed correctly
	Expect(rh.Links.Self.Href).To(Equal("/api/v3/cmp/resourceHandlers/RH-nza16uyn/"))
	Expect(rh.Links.Self.Title).To(Equal("My Test Resource Handler"))
	Expect(rh.Name).To(Equal("My Test Resource Handler"))
	Expect(rh.ID).To(Equal("RH-nza16uyn"))
}
