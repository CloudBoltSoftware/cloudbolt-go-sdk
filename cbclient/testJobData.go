package cbclient

const aJob string = `{
    "_links": {
        "self": {
            "href": "/api/v3/cmp/jobs/JOB-9nrax3gb/",
            "title": "Deploy Blueprint Job 1011"
        },
        "owner": {
            "href": "/api/v3/cmp/users/USR-mxpqe1x7/",
            "title": "user001"
        },
        "parent": {},
        "subjobs": [
            {
                "href": "/api/v3/cmp/jobs/JOB-kb0tuw1e/",
                "title": "Provision Server Job 1012"
            },
            {
                "href": "/api/v3/cmp/jobs/JOB-t2js3lwf/",
                "title": "My Action Job 1013"
            }
        ],
        "prerequisite": {},
        "dependentJobs": [],
        "order": {
            "href": "/api/v3/cmp/orders/ORD-e9v87uia/",
            "title": "Installation of My Simple Blueprint"
        },
        "resource": {
            "href": "/api/v3/cmp/resources/RSC-hjt2wha2/",
            "title": "My Simple Blueprint"
        },
        "servers": [
            {
                "href": "/api/v3/cmp/servers/SVR-srb5y8r3/",
                "title": "myawainstance1"
            }
        ]
    },
    "id": "JOB-9nrax3gb",
    "type": "deploy_blueprint",
    "status": "SUCCESS",
    "workerPid": 20258,
    "workerHostname": "worker00@42975d51567f",
    "canBeRequeued": true,
    "createdDate": "2022-04-10 10:04:15.071344",
    "updatedDate": "2022-04-10 10:07:43.519722",
    "startDate": "2022-04-10 10:04:15.675759",
    "endDate": "2022-04-10 10:07:43.519530",
    "output": "Blueprint deployed successfully",
    "errors": "",
    "tasksDone": 3,
    "totalTasks": 3,
    "label": "",
    "executionState": ""
}`

func responsesForGetJob(i int) (string, int) {
	return bodyForGetJob(i), missingTokenStatusPattern(i)
}

func bodyForGetJob(i int) string {
	return missingTokenBodyPattern(
		aJob,
	)[i]
}
