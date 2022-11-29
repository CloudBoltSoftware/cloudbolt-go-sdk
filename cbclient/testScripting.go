package cbclient

const aScriptingPolicyList string = `{
    "_links": {
        "self": {
            "href": "/api/v3/onefuse/scriptingPolicies/?page=1&filter=name%3AMy_Scripting_Policy",
            "title": "List of Scripting Policies - Page 1 of 1"
        }
    },
    "total": 1,
    "count": 1,
    "_embedded": {
        "scriptingPolicies": [
            {
                "_links": {
                    "self": {
                        "href": "/api/v3/onefuse/scriptingPolicies/901/",
                        "title": "My_Scripting_Policy"
                    },
                    "credential": {
                        "href": "/api/v3/onefuse/moduleCredentials/8328/",
                        "title": "My_Scripting_Policy_Credentials"
                    },
                    "workspace": {
                        "href": "/api/v3/onefuse/workspaces/2/",
                        "title": "Default"
                    }
                },
                "name": "My_Scripting_Policy",
                "id": 901,
                "description": "A Scripting Policy created through automated tests",
                "targetHost": "mytest.net",
                "provisionLaunchCommandTemplate": "sudo /bin/bash {{ scriptName }}",
                "provisionScript": "echo '{\"provisioning-message\":\"Hello, provisioning Script Template test\"}'\nexit 0",
                "provisionSuccessExitCodes": "0",
                "deprovisionScript": "echo '{\"deprovisioning-message\":\"Hello, deprovisioning Script Template test\"}'\nexit 0",
                "deprovisionLaunchCommandTemplate": "sudo /bin/bash {{ scriptName }}",
                "deprovisionSuccessExitCodes": "0"
            }
        ]
    }
}`

const aScriptingDeployment string = `{
    "_links": {
      "self": {
        "href": "/api/v3/onefuse/scriptingDeployments/67/",
        "title": "Scripting Deployment id 67"
      },
      "workspace": {
        "href": "/api/v3/onefuse/workspaces/1/",
        "title": "Default"
      },
      "policy": {
        "href": "/api/v3/onefuse/scriptingPolicies/1/",
        "title": "qalnxtst4_script"
      },
      "jobMetadata": {
        "href": "/api/v3/onefuse/jobMetadata/650/",
        "title": "Job Metadata Record id 650"
      }
    },
    "id": 67,
    "hostname": "qalnxtst4.sovlabs.net",
    "provisioningDetails": {
      "status": "successful",
      "output": [
        {
          "commitId": "b8f2b8b",
          "environment": "prod",
          "tagsAtCommit": "sample_tag",
          "project": "project_acme",
          "currentDate": "01/21/2021",
          "version": "1.0.0"
        },
        {
          "commitId": "a8f2a8a",
          "environment": "dev",
          "tagsAtCommit": "sample_tag",
          "project": "project_acme_v2",
          "currentDate": "01/21/2021",
          "version": "1.0.0"
        }
      ]
    },
    "deprovisioningDetails": {},
    "archived": false
}`

func responsesForScriptingPolicy(i int) (string, int) {
	return bodyForScriptingPolicy(i), missingTokenStatusPattern(i)
}

func bodyForScriptingPolicy(i int) string {
	return missingTokenBodyPattern(
		aScriptingPolicyList,
	)[i]
}

func responsesForScriptingDeployment(i int) (string, int) {
	return bodyForScriptingDeployment(i), missingTokenStatusPattern(i)
}

func bodyForScriptingDeployment(i int) string {
	return missingTokenBodyPattern(
		aScriptingDeployment,
	)[i]
}
