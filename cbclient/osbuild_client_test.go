package cbclient

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestGetOSBuild(t *testing.T) {
	// Register the test with gomega
	RegisterTestingT(t)

	// Setup mock server with scripted responses
	// Setup requests buffer
	server, requests := mockServer(responsesForOSBuild)
	Expect(server).NotTo(BeNil())
	Expect(requests).NotTo(BeNil())

	// Setup CloudBolt Client
	client := getClient(server)
	Expect(client).NotTo(BeNil())

	osbName := "amzn2-ami-hvm-2.0.20210721.2-x86_64-gp2"
	osb, err := client.GetOSBuild(osbName)
	Expect(osb).NotTo(BeNil())
	Expect(err).NotTo(HaveOccurred())

	// This should have made three requests:
	// 1+2. Fail to get order, get a token
	// 3. Successfully getting the order
	Expect(len(*requests)).To(Equal(3))

	// The last request is the one we care about
	Expect((*requests)[2].URL.Path).To(Equal("/api/v3/cmp/osBuilds/"))

	// The CloudBolt Order object should be parsed correctly
	Expect(osb.Links.Self.Href).To(Equal("/api/v3/cmp/osBuilds/OSB-z69hjvki/"))
	Expect(osb.Links.Self.Title).To(Equal("amzn2-ami-hvm-2.0.20210721.2-x86_64-gp2"))
	Expect(osb.Name).To(Equal("amzn2-ami-hvm-2.0.20210721.2-x86_64-gp2"))
	Expect(osb.ID).To(Equal("OSB-z69hjvki"))
}

func TestGetOSBuildById(t *testing.T) {
	// Register the test with gomega
	RegisterTestingT(t)

	// Setup mock server with scripted responses
	// Setup requests buffer
	server, requests := mockServer(responsesForOSBuildById)
	Expect(server).NotTo(BeNil())
	Expect(requests).NotTo(BeNil())

	// Setup CloudBolt Client
	client := getClient(server)
	Expect(client).NotTo(BeNil())

	osbId := "OSB-z69hjvki"
	osb, err := client.GetOSBuildById(osbId)
	Expect(osb).NotTo(BeNil())
	Expect(err).NotTo(HaveOccurred())

	// This should have made three requests:
	// 1+2. Fail to get order, get a token
	// 3. Successfully getting the order
	Expect(len(*requests)).To(Equal(3))

	// The last request is the one we care about
	Expect((*requests)[2].URL.Path).To(Equal("/api/v3/cmp/osBuilds/OSB-z69hjvki/"))

	// The CloudBolt Order object should be parsed correctly
	Expect(osb.Links.Self.Href).To(Equal("/api/v3/cmp/osBuilds/OSB-z69hjvki/"))
	Expect(osb.Links.Self.Title).To(Equal("amzn2-ami-hvm-2.0.20210721.2-x86_64-gp2"))
	Expect(osb.Name).To(Equal("amzn2-ami-hvm-2.0.20210721.2-x86_64-gp2"))
	Expect(osb.ID).To(Equal("OSB-z69hjvki"))
}
