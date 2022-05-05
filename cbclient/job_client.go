package cbclient

import (
	"encoding/json"
	"log"
)

// CloudBoltJob contains metadata about a Job.
// Useful for getting the status of a running or completed job.
type CloudBoltJob struct {
	Links struct {
		Self          CloudBoltHALItem   `json:"self"`
		Owner         CloudBoltHALItem   `json:"owner"`
		Parent        CloudBoltHALItem   `json:"parent"`
		Subjobs       []CloudBoltHALItem `json:"subjobs"`
		Prerequisite  CloudBoltHALItem   `json:"prerequisite"`
		DependentJobs []CloudBoltHALItem `json:"dependent-jobs"`
		Order         CloudBoltHALItem   `json:"order"`
		Resource      CloudBoltHALItem   `json:"resource"`
		Servers       []CloudBoltHALItem `json:"servers"`
	} `json:"_links"`
	ID             string `json:"id"`
	Type           string `json:"type"`
	Status         string `json:"status"`
	WorkerPid      int    `json:"workerPid"`
	WorkerHostname string `json:"workerHostname"`
	CanBeRequeued  bool   `json:"canBeRequeued"`
	CreatedDate    string `json:"createdDate"`
	UpdatedDate    string `json:"updatedDate"`
	StartDate      string `json:"startDate"`
	EndDate        string `json:"endDate"`
	Output         string `json:"output"`
	Errors         string `json:"errors"`
	TasksDone      int    `json:"tasksDone"`
	TotalTasks     int    `json:"totalTasks"`
	Label          string `json:"label"`
	ExecutionState string `json:"executionState"`
}

// GetJob fetches the Job object from CloudBolt at the given path
// - Job Path (jobPath) e.g., "/api/v2/jobs/123/"
func (c *CloudBoltClient) GetJob(jobPath string) (*CloudBoltJob, error) {
	apiurl := c.baseURL
	apiurl.Path = jobPath

	resp, err := c.makeRequest("GET", apiurl.String(), nil)
	if err != nil {
		log.Fatalln(err)

		return nil, err
	}

	// We Decode the data because we already have an io.Reader on hand
	var job CloudBoltJob
	json.NewDecoder(resp.Body).Decode(&job)

	return &job, nil
}
