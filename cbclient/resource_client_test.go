package cbclient

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestGetResource(t *testing.T) {
	// Register test with gomega
	RegisterTestingT(t)

	// Setup mock server with scripted responses
	// Setup requests buffer
	server, requests := mockServer(responsesForGetResource)
	Expect(server).NotTo(BeNil())
	Expect(requests).NotTo(BeNil())

	// Setup CloudBolt Client
	client := getClient(server)
	Expect(client).NotTo(BeNil())

	// Define a resourcePath parameter value
	resourcePath := "/api/v3/cmp/resources/RSC-hjt2wha2/"

	// Get the CloudBolt Resource object
	// Expect no errors to occur
	resource, err := client.GetResource(resourcePath)
	Expect(resource).NotTo(BeNil())
	Expect(err).NotTo(HaveOccurred())

	// This should have made three requests:
	// 1+2. Fail to get resource, get a token
	// 3. Successfully getting the resource
	Expect(len(*requests)).To(Equal(3))

	// The last request is the one we care about
	Expect((*requests)[2].URL.Path).To(Equal("/api/v3/cmp/resources/RSC-hjt2wha2/"))

	// The CloudBolt Resource object should be parsed correctly
	Expect(resource.Links.Self.Href).To(Equal("/api/v3/cmp/resources/RSC-hjt2wha2/"))
	Expect(resource.Links.Self.Title).To(Equal("My Simple Blueprint"))
	Expect(resource.Links.ResourceType.Href).To(Equal("/api/v3/cmp/resourceTypes/RT-bde9nds8/"))
	Expect(resource.Links.ResourceType.Title).To(Equal("service"))
	Expect(resource.Links.Blueprint.Href).To(Equal("/api/v3/cmp/blueprints/BP-esnjtp7u/"))
	Expect(resource.Links.Blueprint.Title).To(Equal("My Simple Blueprint"))
	Expect(resource.Links.Owner.Href).To(Equal("/api/v3/cloudbolt/users/USR-mxpqe1x7/"))
	Expect(resource.Links.Owner.Title).To(Equal("user001"))
	Expect(resource.Links.Group.Href).To(Equal("/api/v3/cloudbolt/groups/GRP-yfbbsfht/"))
	Expect(resource.Links.Group.Title).To(Equal("My Org"))
	Expect(len(resource.Links.Jobs)).To(Equal(3))
	Expect(resource.Links.Jobs[0].Href).To(Equal("/api/v3/cmp/jobs/JOB-9nrax3gb/"))
	Expect(resource.Links.Jobs[0].Title).To(Equal("Deploy Blueprint Job 1011"))
	Expect(resource.Links.Jobs[1].Href).To(Equal("/api/v3/cmp/jobs/JOB-t2js3lwf/"))
	Expect(resource.Links.Jobs[1].Title).To(Equal("My Action Job 1013"))
	Expect(resource.Links.Jobs[2].Href).To(Equal("/api/v3/cmp/jobs/JOB-8i53zztl/"))
	Expect(resource.Links.Jobs[2].Title).To(Equal("My Simple Resource Action Job 1016"))
	Expect(resource.Links.ParentResource.Href).To(Equal(""))
	Expect(resource.Links.ParentResource.Title).To(Equal(""))
	Expect(len(resource.Links.Actions)).To(Equal(3))
	Expect(resource.Links.Actions[0].Href).To(Equal("/api/v3/cmp/resourceActions/RSA-hxfync2x/"))
	Expect(resource.Links.Actions[0].Title).To(Equal("Scale"))
	Expect(resource.Links.Actions[1].Href).To(Equal("/api/v3/cmp/resourceActions/RSA-aq3b3gxm/"))
	Expect(resource.Links.Actions[1].Title).To(Equal("My Resource Action"))
	Expect(resource.Links.Actions[2].Href).To(Equal("/api/v3/cmp/resourceActions/RSA-beim3g0e/"))
	Expect(resource.Links.Actions[2].Title).To(Equal("Delete"))
	Expect(len(resource.Links.Servers)).To(Equal(1))
	Expect(resource.Links.Servers[0].Href).To(Equal("/api/v3/cmp/servers/SVR-srb5y8r3/"))
	Expect(resource.Links.Servers[0].Title).To(Equal("myawsinstance"))
	Expect(resource.Name).To(Equal("My Simple Blueprint"))
	Expect(resource.ID).To(Equal("RSC-hjt2wha2"))
	Expect(resource.Created).To(Equal("2022-04-10 10:04:15"))
	Expect(resource.Status).To(Equal("ACTIVE"))
	Expect(resource.Attributes).To(Not(BeNil()))
}
