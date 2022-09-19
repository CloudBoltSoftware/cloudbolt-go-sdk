package cbclient

const aModulePolicyList string = `{
    "_links": {
        "self": {
            "href": "/api/v3/onefuse/modulePolicies/?page=1&filter=name%3AMy_Module_Policy",
            "title": "List of Module Policies - Page 1 of 1"
        }
    },
    "total": 1,
    "count": 1,
    "_embedded": {
        "modulePolicies": [
            {
                "_links": {
                    "self": {
                        "href": "/api/v3/onefuse/modulePolicies/180/",
                        "title": "My_Module_Policy"
                    },
                    "blueprint": {
                        "href": "/api/v3/onefuse/modules/BP-vjsxfc2z/",
                        "title": "My_Blueprint"
                    },
                    "workspace": {
                        "href": "/api/v3/onefuse/workspaces/2/",
                        "title": "Default"
                    }
                },
                "name": "My_Module_Policy",
                "id": 180,
                "description": "A Module Policy created through automated tests",
                "policyTemplate": "{\"provisioningPayload\": {\"deploymentItems\": {\"plugin-bdi-zjn6ztq6\": {\"parameters\": {\"input_text\": \"{{text}}\",\"from\": \"en\",\"to\": \"ja\"}}}}}",
                "actionPayloads": ""
            }
        ]
    }
}`

const aModuleDeployment string = `{
    "_links": {
      "self": {
        "href": "/api/v3/onefuse/moduleManagedObjects/75/",
        "title": "Module Managed Object id 75"
      },
      "workspace": {
        "href": "/api/v3/onefuse/workspaces/1/",
        "title": "Default"
      },
      "policy": {
        "href": "/api/v3/onefuse/modulePolicies/1/",
        "title": "automated_module_policy"
      },
      "jobMetadata": {
        "href": "/api/v3/onefuse/jobMetadata/662/",
        "title": "Job Metadata Record id 662"
      }
    },
    "id": 75,
    "name": "vip-dev-ap012 (member pp-atltlap004.example.com)",
    "provisioningJobResults": [
      {
        "results": [
          {
            "code": 200,
            "message": "success",
            "lineCount": 24,
            "host": "localhost",
            "tenant": "tenant_vip-dev-ap012",
            "runTime": 2603
          }
        ],
        "action": "provision"
      }
    ],
    "deprovisioningJobResults": [],
    "updateJobResults": [],
    "archived": false,
    "resource": {}
}`

func responsesForModulePolicy(i int) (string, int) {
	return bodyForGetModulePolicy(i), missingTokenStatusPattern(i)
}

func bodyForGetModulePolicy(i int) string {
	return missingTokenBodyPattern(
		aModulePolicyList,
	)[i]
}

func responsesForModuleDeployment(i int) (string, int) {
	return bodyForGetModuleDeployment(i), missingTokenStatusPattern(i)
}

func bodyForGetModuleDeployment(i int) string {
	return missingTokenBodyPattern(
		aModuleDeployment,
	)[i]
}
