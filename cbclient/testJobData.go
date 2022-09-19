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

const aJobStatus = `{
    "_links": {
        "self": {
            "href": "/api/v3/onefuse/jobStatus/3280/",
            "title": "Job Metadata Record id 3280"
        },
        "jobMetadata": {
            "href": "/api/v3/onefuse/jobMetadata/3280/",
            "title": "Job Metadata Record id 3280"
        },
        "workspace": {
            "href": "/api/v3/onefuse/workspaces/2/",
            "title": "Default"
        },
        "policy": {
            "href": "/api/v3/onefuse/modulePolicies/1/",
            "title": "1F_Notification"
        },
        "managedObject": {
            "href": "/api/v3/onefuse/moduleManagedObjects/15/",
            "title": "My Awesome Subject"
        }
    },
    "jobStateDescription": "Successful",
    "startTime": "2022-08-18T17:23:08.386458Z",
    "endTime": "2022-08-18T17:23:21.956198Z",
    "id": 3280,
    "jobType": "Provision Email Notification",
    "jobState": "Successful",
    "jobId": "18f65067-9935-439d-9c3b-c623dc2d5b71",
    "jobEngineId": 2888,
    "jobTrackingId": "3474c59f-6ca0-4d99-82ea-e1b98fca71c6",
    "source": "api",
    "requester": "admin",
    "module": "Notifications",
    "duration": 13569,
    "policyName": "1F_Notification",
    "isPluggable": true,
    "workspace": "Default",
    "actionName": null,
    "renderedTemplates": {},
    "resolvedProperties": {
        "title": "Some Title Here",
        "message": "Some message here.",
        "to_email": "laltomare@cloudbolt.io",
        "OneFuse_Tracking_Id": "3474c59f-6ca0-4d99-82ea-e1b98fca71c6"
    },
    "managedObject": {
        "_links": {
            "self": {
                "href": "/api/v3/onefuse/moduleManagedObjects/15/",
                "title": "Module Managed Object id 15"
            },
            "workspace": {
                "href": "/api/v3/onefuse/workspaces/2/",
                "title": "Default"
            },
            "policy": {
                "href": "/api/v3/onefuse/modulePolicies/1/",
                "title": "1F_Notification"
            },
            "jobMetadata": {
                "href": "/api/v3/onefuse/jobMetadata/3280/",
                "title": "Job Metadata Record id 3280"
            }
        },
        "name": "My Awesome Subject",
        "id": 15,
        "provisioningJobResults": [
            {
                "action": "provision"
            }
        ],
        "deprovisioningJobResults": [],
        "updateJobResults": [],
        "displayDetails": {
            "subject": "My Awesome Subject",
            "from": "CloudBolt <support@cloudbolt.io>",
            "recipients": {
                "to": "['laltomare@cloudbolt.io']"
            },
            "body": "<h1>Some Title Here</h1><p>Some message here.</p>",
            "smtpEndpoint": "smtp",
            "useTls": "true",
            "emailFormat": "html"
        },
        "archived": false,
        "resource": {
            "name": "My Awesome Subject",
            "id": "15",
            "status": "Active",
            "install-date": "2022-08-18T17:23:11.922712",
            "attributes": {
                "OneFuse Notification SMTP Endpoint": "smtp",
                "OneFuse Notification Use TLS": "true",
                "OneFuse Notification Email Format": "html",
                "OneFuse Notification From Address": "support@cloudbolt.io",
                "OneFuse Notification From Name": "CloudBolt",
                "OneFuse Notification Reply To": "",
                "OneFuse Notification CC Addresses": "",
                "OneFuse Notification BCC Addresses": "",
                "OneFuse Notification Subject": "My Awesome Subject",
                "OneFuse Notification Body": "<h1>Some Title Here</h1><p>Some message here.</p>",
                "OneFuse Notification To Addresses": "['laltomare@cloudbolt.io']",
                "OneFuse Tracking ID": "3474c59f-6ca0-4d99-82ea-e1b98fca71c6",
                "OneFuse Managed Object ID": "15"
            }
        }
    }
}`

func responsesForGetJob(i int) (string, int) {
	return bodyForGetJob(i), missingTokenStatusPattern(i)
}

func bodyForGetJob(i int) string {
	return missingTokenBodyPattern(
		aJob,
	)[i]
}

func responsesForGetJobStatus(i int) (string, int) {
	return bodyForGetJobStatus(i), missingTokenStatusPattern(i)
}

func bodyForGetJobStatus(i int) string {
	return missingTokenBodyPattern(
		aJobStatus,
	)[i]
}
