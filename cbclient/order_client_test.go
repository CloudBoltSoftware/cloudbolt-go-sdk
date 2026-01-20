package cbclient

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestGetOrder(t *testing.T) {
	// Register the test with gomega
	RegisterTestingT(t)

	// Setup mock server with scripted responses
	// Setup requests buffer
	server, requests := mockServer(responsesForGetOrder)
	Expect(server).NotTo(BeNil())
	Expect(requests).NotTo(BeNil())

	// Setup CloudBolt Client
	client := getClient(server)
	Expect(client).NotTo(BeNil())

	// Define an orderID parameter value
	orderID := "ORD-e9v87uia"

	// Get the CloudBolt Order object
	// Expect no errors to occur
	order, err := client.GetOrder(orderID)
	Expect(order).NotTo(BeNil())
	Expect(err).NotTo(HaveOccurred())

	// This should have made three requests:
	// 1+2. Fail to get order, get a token
	// 3. Successfully getting the order
	Expect(len(*requests)).To(Equal(3))

	// The last request is the one we care about
	Expect((*requests)[2].URL.Path).To(Equal("/api/v3/cmp/orders/ORD-e9v87uia/"))

	// The CloudBolt Order object should be parsed correctly
	Expect(order.Links.Self.Href).To(Equal("/api/v3/cmp/orders/ORD-e9v87uia/"))
	Expect(order.Links.Self.Title).To(Equal("Installation of My Simple Blueprint"))
	Expect(order.Links.Group.Href).To(Equal("/api/v3/cloudbolt/groups/GRP-yfbbsfht/"))
	Expect(order.Links.Group.Title).To(Equal("My Org"))
	Expect(order.Links.Owner.Href).To(Equal("/api/v3/cloudbolt/users/USR-mxpqe1x7/"))
	Expect(order.Links.Owner.Title).To(Equal("user001"))
	Expect(order.Links.ApprovedBy.Href).To(Equal("/api/v3/cloudbolt/users/USR-mxpqe1x7/"))
	Expect(order.Links.ApprovedBy.Title).To(Equal("user001"))
	Expect(len(order.Links.Jobs)).To(Equal(2))
	Expect(order.Links.Jobs[0].Href).To(Equal("/api/v3/cmp/jobs/JOB-9nrax3gb/"))
	Expect(order.Links.Jobs[0].Title).To(Equal("Deploy Blueprint Job 1011"))
	Expect(order.Links.Duplicate.Href).To(Equal("/api/v3/cmp/orders/ORD-e9v87uia/duplicate/"))
	Expect(order.Links.Duplicate.Title).To(Equal("Duplicate Order"))
	Expect(order.Name).To(Equal("Installation of My Simple Blueprint"))
	Expect(order.ID).To(Equal("ORD-e9v87uia"))
	Expect(order.Status).To(Equal("SUCCESS"))
	Expect(order.Rate).To(Equal("4.18/month"))
	Expect(len(order.DeploymentItems)).To(Equal(1))
	Expect(order.DeploymentItems[0].ID).To(Equal("OI-1p0bajs6"))
	Expect(order.DeploymentItems[0].ResourceName).To(Equal("My Simple Blueprint"))
	Expect(order.DeploymentItems[0].ResourceParameters).To(Not(BeNil()))
	Expect(order.DeploymentItems[0].Blueprint.Href).To(Equal("/api/v3/cmp/blueprints/BP-esnjtp7u/"))
	Expect(order.DeploymentItems[0].Blueprint.Title).To(Equal("My Simple Blueprint"))
	Expect(order.DeploymentItems[0].BlueprintItemsArguments).To(Not(BeNil()))
	Expect(order.DeploymentItems[0].ItemType).To(Equal("blueprint"))
}


func TestGetOrderStatus(t *testing.T) {
	// Register the test with gomega
	RegisterTestingT(t)

	// Setup mock server with scripted responses
	// Setup requests buffer
	server, requests := mockServer(responsesForGetOrderStatus)
	Expect(server).NotTo(BeNil())
	Expect(requests).NotTo(BeNil())

	// Setup CloudBolt Client
	client := getClient(server)
	Expect(client).NotTo(BeNil())

	// Define an orderID parameter value
	orderID := "ORD-e9v87uia"

	// Get the CloudBolt Order object
	// Expect no errors to occur
	orderStatus, err := client.GetOrderStatus(orderID)
	Expect(orderStatus).NotTo(BeNil())
	Expect(err).NotTo(HaveOccurred())

	// This should have made three requests:
	// 1+2. Fail to get order, get a token
	// 3. Successfully getting the order
	Expect(len(*requests)).To(Equal(3))

	Expect(orderStatus.Status).To(Equal("FAILURE"))
	Expect(len(orderStatus.OutputMessages)).To(Equal(2))
	Expect(orderStatus.OutputMessages[0]).To(Equal("Job 101: Output for Job 101"))
	Expect(orderStatus.OutputMessages[1]).To(Equal("Job 102: Output for Job 102"))
	Expect(len(orderStatus.ErrorMessages)).To(Equal(2))
	Expect(orderStatus.ErrorMessages[0]).To(Equal("Job 101: Error for Job 101"))
	Expect(orderStatus.ErrorMessages[1]).To(Equal("Job 102: Error for Job 102"))
}