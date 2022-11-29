package cbclient

const aServiceNowCMDBPolicyList string = `{
    "_links": {
        "self": {
            "href": "/api/v3/onefuse/servicenowCMDBPolicies/?page=1&filter=name%3AMy_SNOW_CMDB_Policy",
            "title": "List of Servicenow Cmdb Policies - Page 1 of 1"
        }
    },
    "total": 1,
    "count": 1,
    "_embedded": {
        "servicenowCMDBPolicies": [
            {
                "_links": {
                    "self": {
                        "href": "/api/v3/onefuse/servicenowCMDBPolicies/224/",
                        "title": "My_SNOW_CMDB_Policy"
                    },
                    "workspace": {
                        "href": "/api/v3/onefuse/workspaces/2/",
                        "title": "Default"
                    },
                    "endpoint": {
                        "href": "/api/v3/onefuse/endpoints/7723/",
                        "title": "My_SNOW_CMDB_Policy_Endpoint"
                    }
                },
                "name": "My_SNOW_CMDB_Policy",
                "id": 224,
                "description": "A ServiceNow CMDB Policy",
                "provisionTemplate": "{\"items\": [{\"className\": \"cmdb_ci_linux_server\", \"values\": {\"discovery_source\": \"Other Automated\", \"os\": \"GNU/Linux\", \"name\": \"{{ OneFuse_VmNic0.hostname }}-{{OneFuse_Suffix}}\", \"dns_domain\": \"{{ OneFuse_VmNic0.dnsSuffix }}\", \"host_name\": \"{{ OneFuse_VmNic0.hostname }}\", \"fqdn\": \"{{ OneFuse_VmNic0.fqdn }}\", \"ip_address\": \"{{ OneFuse_VmNic0.ipAddress }}\", \"serial_number\": \"vmware-{{ OneFuse_VmHardware.platformUuid }}-{{OneFuse_Suffix}}\", \"cpu_count\": \"{{ OneFuse_VmHardware.cpuCount }}\", \"disk_space\": \"{{ OneFuse_VmHardware.totalStorageGB }}\", \"ram\": \"{{ OneFuse_VmHardware.memoryMB }}\", \"virtual\": \"true\", \"state\": \"{{ OneFuse_VmHardware.powerState }}\", \"hardware_status\": \"installed\"}}]}",
                "updateTemplate": "{\"items\": [{\"className\": \"cmdb_ci_linux_server\", \"values\": {\"discovery_source\": \"Other Automated\", \"os\": \"GNU/Linux\", \"name\": \"{{ OneFuse_VmNic0.hostname }}-{{OneFuse_Suffix}}\", \"dns_domain\": \"{{ OneFuse_VmNic0.dnsSuffix }}\", \"host_name\": \"{{ OneFuse_VmNic0.hostname }}\", \"fqdn\": \"{{ OneFuse_VmNic0.fqdn }}\", \"ip_address\": \"{{ OneFuse_VmNic0.ipAddress }}\", \"serial_number\": \"vmware-{{ OneFuse_VmHardware.platformUuid }}-{{OneFuse_Suffix}}\", \"cpu_count\": \"{{ OneFuse_VmHardware.cpuCount }}\", \"disk_space\": \"{{ OneFuse_VmHardware.totalStorageGB }}\", \"ram\": \"{{ OneFuse_VmHardware.memoryMB }}\", \"virtual\": \"true\", \"state\": \"{{ OneFuse_VmHardware.powerState }}\", \"hardware_status\": \"installed\"}}]}",
                "deprovisionTemplate": "{\"items\": [{\"className\": \"cmdb_ci_linux_server\", \"values\": {\"name\": \"{{ OneFuse_VmNic0.hostname }}-{{OneFuse_Suffix}}\", \"discovery_source\": \"Other Automated\", \"serial_number\": \"vmware-{{ OneFuse_VmHardware.platformUuid }}-{{OneFuse_Suffix}}\", \"state\": \"{{ OneFuse_VmHardware.powerState }}\", \"hardware_status\": \"retired\"}}]}"
            }
        ]
    }
}`

const aServiceNowCMDBDeployment string = `{
    "_links": {
      "self": {
        "href": "/api/v3/onefuse/servicenowCMDBDeployments/24/",
        "title": "Service Now Deployment id 24"
      },
      "workspace": {
        "href": "/api/v3/onefuse/workspaces/1/",
        "title": "Default"
      },
      "policy": {
        "href": "/api/v3/onefuse/servicenowCMDBPolicies/4/",
        "title": "test_serviceNow_CMDB_policy_updated_581"
      },
      "jobMetadata": {
        "href": "/api/v3/onefuse/jobMetadata/434/",
        "title": "Job Metadata Record id 434"
      }
    },
    "id": 24,
    "configurationItemsInfo": [
      {
        "ciClassName": "cmdb_ci_linux_server",
        "ciName": "host14"
      },
      {
        "ciClassName": "cmdb_ci_linux_server",
        "ciName": "host15"
      }
    ],
    "executionDetails": {
      "latestExecution": {
        "lifecycle": "provision",
        "timestamp": "2021-03-01 23:57:45"
      },
      "response": {
        "result": {
          "items": [
            {
              "className": "cmdb_ci_linux_server",
              "operation": "UPDATE",
              "sysId": "d2132d7007...b6fd1e7c1ed019",
              "identifierEntrySysId": "556eb250c...8d4bea192d3ae92",
              "identificationAttempts": [
                {
                  "identifierName": "Hardware Rule",
                  "attemptResult": "SKIPPED",
                  "attributes": [
                    "serial_number",
                    "serial_number_type"
                  ],
                  "searchOnTable": "cmdb_serial_number",
                  "hybridEntryCiAttributes": []
                },
                {
                  "identifierName": "Hardware Rule",
                  "attemptResult": "SKIPPED",
                  "attributes": [
                    "serial_number"
                  ],
                  "searchOnTable": "cmdb_ci_hardware",
                  "hybridEntryCiAttributes": []
                },
                {
                  "identifierName": "Hardware Rule",
                  "attemptResult": "MATCHED",
                  "attributes": [
                    "name"
                  ],
                  "searchOnTable": "cmdb_ci_hardware",
                  "hybridEntryCiAttributes": []
                }
              ]
            },
            {
              "className": "cmdb_ci_linux_server",
              "operation": "UPDATE",
              "sysId": "d2132d7007...b6fd1e7c1ed01d",
              "identifierEntrySysId": "556eb250c...8d4bea192d3ae92",
              "identificationAttempts": [
                {
                  "identifierName": "Hardware Rule",
                  "attemptResult": "SKIPPED",
                  "attributes": [
                    "serial_number",
                    "serial_number_type"
                  ],
                  "searchOnTable": "cmdb_serial_number",
                  "hybridEntryCiAttributes": []
                },
                {
                  "identifierName": "Hardware Rule",
                  "attemptResult": "SKIPPED",
                  "attributes": [
                    "serial_number"
                  ],
                  "searchOnTable": "cmdb_ci_hardware",
                  "hybridEntryCiAttributes": []
                },
                {
                  "identifierName": "Hardware Rule",
                  "attemptResult": "MATCHED",
                  "attributes": [
                    "name"
                  ],
                  "searchOnTable": "cmdb_ci_hardware",
                  "hybridEntryCiAttributes": []
                }
              ]
            }
          ],
          "relations": []
        }
      }
    },
    "archived": false
}`

func responsesForServiceNowCMDBPolicy(i int) (string, int) {
	return bodyForGetServiceNowCMDBPolicy(i), missingTokenStatusPattern(i)
}

func bodyForGetServiceNowCMDBPolicy(i int) string {
	return missingTokenBodyPattern(
		aServiceNowCMDBPolicyList,
	)[i]
}

func responsesForServiceNowCMDBDeployment(i int) (string, int) {
	return bodyForGetServiceNowCMDBDeployment(i), missingTokenStatusPattern(i)
}

func bodyForGetServiceNowCMDBDeployment(i int) string {
	return missingTokenBodyPattern(
		aServiceNowCMDBDeployment,
	)[i]
}
