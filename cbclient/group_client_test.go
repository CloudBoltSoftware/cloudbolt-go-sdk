package cbclient

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestVerifyGroup(t *testing.T) {
	// Register the test with gomega
	RegisterTestingT(t)

	// Setup mock server with scripted responses
	// Setup requests buffer
	server, requests := mockServer(responsesForVerifyGroup)
	Expect(server).NotTo(BeNil())

	// Setup CloudBolt Client
	client := getClient(server)
	Expect(client).NotTo(BeNil())

	sampleGroupPath := "/api/v3/cmp/groups/GRP-zg550a1z/"
	sampleParentPath := "the group/the subgroup"

	good, err := client.verifyGroup(sampleGroupPath, sampleParentPath)
	Expect(good).To(BeTrue())
	Expect(err).NotTo(HaveOccurred())

	// We expect that to find this group we needed to make 4 API calls
	// 1+2. Fail to get group, get a token
	// 3. make request to /api/v2/groups/GRP-an0thrgrp/ to verify `the subgroup` is this group's parent
	// 4. make request to /api/v2/groups/... to verify `the group` is `the subgroup`'s parent.
	Expect(len(*requests)).To(Equal(4))
}

func TestGetGroupById(t *testing.T) {
	// Register the test with gomega
	RegisterTestingT(t)

	// Setup mock server with scripted responses
	// Setup requests buffer
	server, requests := mockServer(responsesForGroupById)
	Expect(server).NotTo(BeNil())
	Expect(requests).NotTo(BeNil())

	// Setup CloudBolt Client
	client := getClient(server)
	Expect(client).NotTo(BeNil())

	environmentId := "GRP-yfbbsfht"
	environment, err := client.GetGroupById(environmentId)
	Expect(environment).NotTo(BeNil())
	Expect(err).NotTo(HaveOccurred())

	// This should have made three requests:
	// 1+2. Fail to get group, get a token
	// 3. Successfully getting the order
	Expect(len(*requests)).To(Equal(3))

	// The last request is the one we care about
	Expect((*requests)[2].URL.Path).To(Equal("/api/v3/cmp/groups/GRP-yfbbsfht/"))

	// The CloudBolt Order object should be parsed correctly
	Expect(environment.Links.Self.Href).To(Equal("/api/v3/cmp/groups/GRP-yfbbsfht/"))
	Expect(environment.Links.Self.Title).To(Equal("the group"))
	Expect(environment.Name).To(Equal("the group"))
	Expect(environment.ID).To(Equal("GRP-yfbbsfht"))
}

// This is a fun test, let's break down what exactly happens.
// If you look in `testData` at `responsesForGetGroup` you see we return four things:
//   - listOfGroups: a response to the query /api/v2/groups/?filter=name:the+childgroup
//   - yetAnotherGroup: a decoy group with the same name. This is allowed in
//     CloudBolt since group names only need to be unique _within_ a subgroup.
//   - aChildGroup: The real group we are looking for.
//   - aSubGroup: The parent of aChildGroup, used to verify this is the "real" group.
//   - aGroup: The parent of aSubGroup, also used to verify this is the "real" group.
// The calls look like this:
//   1. Call to the list of groups.
//   2. Try to verify yetAnotherGroup, which has no parents so it fails.
//   3. Try to verify aChildGroup, it has the correct parent, so verify the parent.
//   4. Try to verify aSubGroup, which also has the correct parent, and reaches
//      the root of the search so we return success in `verifyGroup`, passing the
//      test and finishing the call to GetGroup().
func TestGetGroup(t *testing.T) {
	// Register the test with gomega
	RegisterTestingT(t)

	// Setup mock server with scripted responses
	// Setup requests buffer
	server, requests := mockServer(responsesForGetGroup)
	Expect(server).NotTo(BeNil())
	Expect(requests).NotTo(BeNil())

	// Setup CloudBolt Client
	client := getClient(server)
	Expect(client).NotTo(BeNil())

	group, err := client.GetGroup("/the group/the subgroup/the childgroup/")
	Expect(err).NotTo(HaveOccurred())
	Expect(group).NotTo(BeNil())

	Expect(len((*requests))).To(Equal(6))

	Expect((*requests)[2].URL.Path).To(Equal("/api/v3/cmp/groups/"))
	Expect((*requests)[2].URL.RawQuery).To(Equal("filter=name:the+childgroup"))
	Expect((*requests)[2].Header["Authorization"]).To(Equal([]string{"Bearer Testing Token"}))

	Expect((*requests)[3].URL.Path).To(Equal("/api/v3/cmp/groups/GRP-zg550a1x/"))
	Expect((*requests)[3].URL.RawQuery).To(Equal(""))
	Expect((*requests)[3].Header["Authorization"]).To(Equal([]string{"Bearer Testing Token"}))

	Expect((*requests)[4].URL.Path).To(Equal("/api/v3/cmp/groups/GRP-zg550a1z/"))
	Expect((*requests)[4].URL.RawQuery).To(Equal(""))
	Expect((*requests)[4].Header["Authorization"]).To(Equal([]string{"Bearer Testing Token"}))

	Expect((*requests)[5].URL.Path).To(Equal("/api/v3/cmp/groups/GRP-uz64vfht/"))
	Expect((*requests)[5].URL.RawQuery).To(Equal(""))
	Expect((*requests)[5].Header["Authorization"]).To(Equal([]string{"Bearer Testing Token"}))

	// The CloudBolt Group object should be parsed correctly
	Expect(group.Links.Self.Href).To(Equal("/api/v3/cmp/groups/GRP-zg550a1z/"))
	Expect(group.Links.Self.Title).To(Equal("the childgroup"))
	Expect(group.Name).To(Equal("the childgroup"))
	Expect(group.ID).To(Equal("GRP-zg550a1z"))
}
