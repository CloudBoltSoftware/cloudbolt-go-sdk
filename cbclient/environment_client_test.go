package cbclient

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestGetEnvironment(t *testing.T) {
	// Register the test with gomega
	RegisterTestingT(t)

	// Setup mock server with scripted responses
	// Setup requests buffer
	server, requests := mockServer(responsesForEnvironment)
	Expect(server).NotTo(BeNil())
	Expect(requests).NotTo(BeNil())

	// Setup CloudBolt Client
	client := getClient(server)
	Expect(client).NotTo(BeNil())

	environmentName := "MY AWS Environment"
	environment, err := client.GetEnvironment(environmentName)
	Expect(environment).NotTo(BeNil())
	Expect(err).NotTo(HaveOccurred())

	// This should have made three requests:
	// 1+2. Fail to get order, get a token
	// 3. Successfully getting the order
	Expect(len(*requests)).To(Equal(3))

	// The last request is the one we care about
	Expect((*requests)[2].URL.Path).To(Equal("/api/v3/cmp/environments/"))

	// The CloudBolt Order object should be parsed correctly
	Expect(environment.Links.Self.Href).To(Equal("/api/v3/cmp/environments/ENV-1tytr2pu/"))
	Expect(environment.Links.Self.Title).To(Equal("MY AWS Environment"))
	Expect(environment.Name).To(Equal("MY AWS Environment"))
	Expect(environment.ID).To(Equal("ENV-1tytr2pu"))
}

func TestGetEnvironmentById(t *testing.T) {
	// Register the test with gomega
	RegisterTestingT(t)

	// Setup mock server with scripted responses
	// Setup requests buffer
	server, requests := mockServer(responsesForEnvironmentById)
	Expect(server).NotTo(BeNil())
	Expect(requests).NotTo(BeNil())

	// Setup CloudBolt Client
	client := getClient(server)
	Expect(client).NotTo(BeNil())

	environmentId := "ENV-1tytr2pu"
	environment, err := client.GetEnvironmentById(environmentId)
	Expect(environment).NotTo(BeNil())
	Expect(err).NotTo(HaveOccurred())

	// This should have made three requests:
	// 1+2. Fail to get order, get a token
	// 3. Successfully getting the order
	Expect(len(*requests)).To(Equal(3))

	// The last request is the one we care about
	Expect((*requests)[2].URL.Path).To(Equal("/api/v3/cmp/environments/ENV-1tytr2pu/"))

	// The CloudBolt Order object should be parsed correctly
	Expect(environment.Links.Self.Href).To(Equal("/api/v3/cmp/environments/ENV-1tytr2pu/"))
	Expect(environment.Links.Self.Title).To(Equal("MY AWS Environment"))
	Expect(environment.Name).To(Equal("MY AWS Environment"))
	Expect(environment.ID).To(Equal("ENV-1tytr2pu"))
}
