package cbclient

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestGetJob(t *testing.T) {
	// Register the test with gomega
	RegisterTestingT(t)

	// Setup mock server with scripted responses
	// Setup requests buffer
	server, requests := mockServer(responsesForGetJob)
	Expect(server).NotTo(BeNil())
	Expect(requests).NotTo(BeNil())

	// Setup CloudBolt Client
	client := getClient(server)
	Expect(client).NotTo(BeNil())

	// Define a jobPath parameter value
	jobPath := "/api/v3/cmp/jobs/JOB-9nrax3gb/"

	// Get the CloudBolt Job object
	// Expect no errors to occur
	job, err := client.GetJob(jobPath)
	Expect(job).NotTo(BeNil())
	Expect(err).NotTo(HaveOccurred())

	// This should have made three requests:
	// 1+2. Fail to get job, get a token
	// 3. Successfully getting the job
	Expect(len(*requests)).To(Equal(3))

	// The last request is the one we care about
	Expect((*requests)[2].URL.Path).To(Equal("/api/v3/cmp/jobs/JOB-9nrax3gb/"))

	// // The CloudBolt Job object should be parsed correctly
	Expect(job.Links.Self.Href).To(Equal("/api/v3/cmp/jobs/JOB-9nrax3gb/"))
	Expect(job.Links.Self.Title).To(Equal("Deploy Blueprint Job 1011"))
	Expect(job.Links.Owner.Href).To(Equal("/api/v3/cmp/users/USR-mxpqe1x7/"))
	Expect(job.Links.Owner.Title).To(Equal("user001"))
	Expect(job.Links.Parent.Href).To(Equal(""))
	Expect(job.Links.Parent.Title).To(Equal(""))
	Expect(len(job.Links.Subjobs)).To(Equal(2))
	Expect(job.Links.Subjobs[0].Href).To(Equal("/api/v3/cmp/jobs/JOB-kb0tuw1e/"))
	Expect(job.Links.Subjobs[0].Title).To(Equal("Provision Server Job 1012"))
	Expect(job.Links.Subjobs[1].Href).To(Equal("/api/v3/cmp/jobs/JOB-t2js3lwf/"))
	Expect(job.Links.Subjobs[1].Title).To(Equal("My Action Job 1013"))
	Expect(job.Links.Prerequisite.Href).To(Equal(""))
	Expect(job.Links.Prerequisite.Title).To(Equal(""))
	Expect(len(job.Links.DependentJobs)).To(Equal(0))
	Expect(job.Links.Order.Href).To(Equal("/api/v3/cmp/orders/ORD-e9v87uia/"))
	Expect(job.Links.Order.Title).To(Equal("Installation of My Simple Blueprint"))
	Expect(job.Links.Resource.Href).To(Equal("/api/v3/cmp/resources/RSC-hjt2wha2/"))
	Expect(job.Links.Resource.Title).To(Equal("My Simple Blueprint"))
	Expect(len(job.Links.Servers)).To(Equal(1))
	Expect(job.Links.Servers[0].Href).To(Equal("/api/v3/cmp/servers/SVR-srb5y8r3/"))
	Expect(job.Links.Servers[0].Title).To(Equal("myawainstance1"))
	Expect(job.ID).To(Equal("JOB-9nrax3gb"))
	Expect(job.Type).To(Equal("deploy_blueprint"))
	Expect(job.Status).To(Equal("SUCCESS"))
	Expect(job.WorkerPid).To(Equal(20258))
	Expect(job.WorkerHostname).To(Equal("worker00@42975d51567f"))
	Expect(job.CanBeRequeued).To(Equal(true))
	Expect(job.CreatedDate).To(Equal("2022-04-10 10:04:15.071344"))
	Expect(job.UpdatedDate).To(Equal("2022-04-10 10:07:43.519722"))
	Expect(job.StartDate).To(Equal("2022-04-10 10:04:15.675759"))
	Expect(job.EndDate).To(Equal("2022-04-10 10:07:43.519530"))
	Expect(job.Output).To(Equal("Blueprint deployed successfully"))
	Expect(job.Errors).To(Equal(""))
	Expect(job.TasksDone).To(Equal(3))
	Expect(job.TotalTasks).To(Equal(3))
	Expect(job.Label).To(Equal(""))
	Expect(job.ExecutionState).To(Equal(""))
}
