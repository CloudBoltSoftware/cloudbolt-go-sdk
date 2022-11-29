package cbclient

const aAnsibleTowerPolicyList string = `{
    "_links": {
        "self": {
            "href": "/api/v3/onefuse/ansibleTowerPolicies/?page=1&filter=name%3AMy_Ansible_Tower_Policy",
            "title": "List of Ansible Tower Policies - Page 1 of 1"
        }
    },
    "total": 1,
    "count": 1,
    "_embedded": {
        "ansibleTowerPolicies": [
            {
                "_links": {
                    "self": {
                        "href": "/api/v3/onefuse/ansibleTowerPolicies/6/",
                        "title": "My_Ansible_Tower_Policy"
                    },
                    "workspace": {
                        "href": "/api/v3/onefuse/workspaces/2/",
                        "title": "Default"
                    },
                    "endpoint": {
                        "href": "/api/v3/onefuse/endpoints/192/",
                        "title": "My_Ansible_Tower_Endpoint"
                    }
                },
                "id": 6,
                "name": "My_Ansible_Tower_Policy",
                "description": "An Ansible Policy created through automated tests",
                "organizationName": "{{org_name}}",
                "provisioningJobTemplates": "[{\"name\": \"{{job_template_name}}\"}]",
                "deprovisioningJobTemplates": "[{\"name\": \"{{job_template_name}}\"}]",
                "verifyPromptOnLaunchForLimit": false,
                "machineCredentialOverride": null,
                "extraVarsOverride": null,
                "inventoryName": null,
                "groups": null
            }
        ]
    }
}
`
const aAnsibleTowerDeployment = `{
    "_links": {
      "self": {
        "href": "/api/v3/onefuse/ansibleTowerDeployments/4/",
        "title": "Ansible Tower Deployment id 4"
      },
      "workspace": {
        "href": "/api/v3/onefuse/workspaces/2/",
        "title": "Default"
      },
      "policy": {
        "href": "/api/v3/onefuse/ansibleTowerPolicies/1/",
        "title": "atPolicy01"
      },
      "jobMetadata": {
        "href": "/api/v3/onefuse/jobMetadata/7/",
        "title": "Job Metadata Record id 7"
      }
    },
    "id": 4,
    "limit": "rb*",
    "inventoryName": "rb-inv001",
    "hosts": [],
    "provisioningJobResults": [
      {
        "output": "Full Ansible Play Log",
        "status": "successful",
        "jobTemplateName": "qa-sleep"
      }
    ],
    "deprovisioningJobResults": {},
    "archived": false
  }
`

func responsesForAnsibleTowerPolicy(i int) (string, int) {
	return bodyForGetAnsibleTowerPolicy(i), missingTokenStatusPattern(i)
}

func bodyForGetAnsibleTowerPolicy(i int) string {
	return missingTokenBodyPattern(
		aAnsibleTowerPolicyList,
	)[i]
}

func responsesForAnsibleTowerDeoplyment(i int) (string, int) {
	return bodyForGetAnsibleTowerDeployment(i), missingTokenStatusPattern(i)
}

func bodyForGetAnsibleTowerDeployment(i int) string {
	return missingTokenBodyPattern(
		aAnsibleTowerDeployment,
	)[i]
}
