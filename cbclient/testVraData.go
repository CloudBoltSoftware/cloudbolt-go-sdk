package cbclient

const aVraPolicyList string = `{
    "_links": {
        "self": {
            "href": "/api/v3/onefuse/vraPolicies/?page=1&filter=name%3AQA_vra8_VRA_POLICY98127886",
            "title": "List of Vra Policies - Page 1 of 1"
        }
    },
    "total": 1,
    "count": 1,
    "_embedded": {
        "vraPolicies": [
            {
                "_links": {
                    "self": {
                        "href": "/api/v3/onefuse/vraPolicies/1/",
                        "title": "My_Vra_Policy"
                    },
                    "workspace": {
                        "href": "/api/v3/onefuse/workspaces/2/",
                        "title": "Default"
                    },
                    "endpoint": {
                        "href": "/api/v3/onefuse/endpoints/27/",
                        "title": "My_Vra_Policy_Endpoint"
                    }
                },
                "id": 1,
                "type": "vra8",
                "name": "My_Vra_Policy",
                "description": "A Vra Policy",
                "cloudTemplateName": "my_naming_snow_1",
                "cloudTemplateVersionNumber": "1",
                "blueprintId": "ce8c92e6-7b00-4a77-971d-8af6a3113359",
                "projectName": "My Test Project",
                "cloudTemplateInputs": null,
                "userMapping": null
            }
        ]
    }
}`

const aVraDeployment string = `{
    "_links": {
      "self": {
        "href": "/api/v3/onefuse/vraDeployments/1/",
        "title": "vRealize Automation Deployment id 1"
      },
      "workspace": {
        "href": "/api/v3/onefuse/workspaces/1/",
        "title": "Default"
      },
      "policy": {
        "href": "/api/v3/onefuse/vraPolicies/1/",
        "title": "Vra_Policy_1"
      },
      "jobMetadata": {
        "href": "/api/v3/onefuse/jobMetadata/1/",
        "title": "Job Metadata Record id 1"
      }
    },
    "id": 1,
    "blueprintName": "vRA Blueprint Name",
    "projectName": "vRA Project Name",
    "deploymentInfo": [
      {
        "id": "88fedb90-5b48-4874-b34f-c82b1f10cf48",
        "name": "Deployment Name",
        "description": "Description",
        "orgId": "626f025c-f9e3-4a5d-9114-f12b487842f9",
        "blueprintId": "35396fb5-7e2a-48ea-a684-974440ccafc6",
        "blueprintVersion": 1,
        "createdAt": "2021-01-15T15:05:04.576Z",
        "createdBy": "vrasvc",
        "lastUpdatedAt": "2021-01-15T15:05:04.576Z",
        "lastUpdatedBy": "tango-blueprint-G7xnPk7qV1qORrlq",
        "inputs": [
          {
            "cpuCount": 1,
            "totalMemoryMB": 1024
          }
        ],
        "projectId": "2d054652-fe53-4b6b-ac5a-3c9dc83ce9d9",
        "status": "CREATE_SUCCESSFUL",
        "childResources": [
          {
            "id": "335cf0c5-e63a-49b3-8bf6-045d916d52bd",
            "name": "Cloud_vSphere_Machine_1",
            "type": "Cloud.vSphere.Machine",
            "dependsOn": [],
            "createdAt": "2021-01-26T20:04:37.115711Z",
            "properties": {
              "resourceId": "335cf0c5-e63a-49b3-8bf6-045d916d52bd",
              "resourceDescLink": "/resources/compute-descriptions/ce8f35e0-9e8f-42cc-9b6c-20dec5a24f81",
              "provisionGB": "0",
              "powerState": "ON",
              "zone": "Cluster1 / CB_Testing",
              "computeHostType": "ResourcePool",
              "id": "/resources/compute/335cf0c5-e63a-49b3-8bf6-045d916d52bd",
              "cpuCount": 1,
              "totalMemoryMB": 1024,
              "endpointType": "vsphere",
              "resourceName": "vrasvc-391",
              "softwareName": "Other (64-bit)",
              "name": "Cloud_vSphere_Machine_1",
              "resourceLink": "/resources/compute/335cf0c5-e63a-49b3-8bf6-045d916d52bd",
              "region": "SovLabs",
              "storage": {
                "disks": [
                  {
                    "vm": "VirtualMachine:vm-671603",
                    "name": "boot-disk",
                    "type": "HDD",
                    "vcUuid": "634f89cd-bcf8-421b-bfa3-65f09557abca",
                    "bootOrder": 1,
                    "encrypted": false,
                    "capacityGb": 0,
                    "persistent": false,
                    "independent": "false",
                    "provisionGB": "0",
                    "diskPlacementRef": "StoragePod:group-p100",
                    "provisioningType": "thin"
                  }
                ]
              },
              "networks": [
                {
                  "name": "dvs_312_10.30.12.0_24",
                  "assignment": "dynamic",
                  "deviceIndex": 0,
                  "mac_address": "00:50:56:a5:10:a5",
                  "resourceName": "dvs_312_10.30.12.0_24"
                }
              ],
              "osType": "LINUX",
              "resourcePool": "/resources/pools/6becacb96dc6e875-7f703c5265a63d87",
              "componentType": "Cloud.vSphere.Machine",
              "endpointId": "cf38f2a7-ecd1-4560-a14d-5f09f9bc0c54",
              "datastoreName": "XtremIO_5t_datastore5",
              "primaryMAC": "00:50:56:a5:10:a5",
              "computeHostRef": "ResourcePool:resgroup-654387",
              "imageRef": "MassTesting",
              "account": "vcenter01.sovlabs.net",
              "vcUuid": "634f89cd-bcf8-421b-bfa3-65f09557abca"
            },
            "state": "OK"
          }
        ]
      }
    ],
    "deprovisioningJobResults": {},
    "archived": false
  }`

func responsesForVraPolicy(i int) (string, int) {
	return bodyForGetVraPolicy(i), missingTokenStatusPattern(i)
}

func bodyForGetVraPolicy(i int) string {
	return missingTokenBodyPattern(
		aVraPolicyList,
	)[i]
}

func responsesForVraDeployment(i int) (string, int) {
	return bodyForGetVraDeployment(i), missingTokenStatusPattern(i)
}

func bodyForGetVraDeployment(i int) string {
	return missingTokenBodyPattern(
		aVraDeployment,
	)[i]
}
