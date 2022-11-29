package cbclient

const aADPolicyList string = `{
    "_links": {
        "self": {
            "href": "/api/v3/onefuse/microsoftADPolicies/?page=1&filter=name%3Amicrosoft_ad_policy_98045627",
            "title": "List of Microsoft Ad Policies - Page 1 of 1"
        }
    },
    "total": 1,
    "count": 1,
    "_embedded": {
        "microsoftADPolicies": [
            {
                "_links": {
                    "self": {
                        "href": "/api/v3/onefuse/microsoftADPolicies/17/",
                        "title": "My_AD_Policy"
                    },
                    "workspace": {
                        "href": "/api/v3/onefuse/workspaces/2/",
                        "title": "Default"
                    },
                    "microsoftEndpoint": {
                        "href": "/api/v3/onefuse/endpoints/1539/",
                        "title": "My_AD_Policy_Endpoint"
                    }
                },
                "name": "My_AD_Policy",
                "id": 17,
                "description": "microsoft active directory policy description",
                "computerNameLetterCase": "LOWER",
                "buildOu": "",
                "createBuildOu": true,
                "removeBuildOu": true,
                "ou": "ou=qa,ou=environments,dc=sovlabs,dc=net",
                "createOrganizationalUnit": false,
                "removeOrganizationalUnit": false,
                "deleteComputerAccountsByName": true,
                "securityGroups": []
            }
        ]
    }
}`

const aComputerAccount string = `{
    "_links": {
      "self": {
        "href": "/api/v3/onefuse/microsoftADComputerAccounts/23/",
        "title": "testcompacct_490991"
      },
      "workspace": {
        "href": "/api/v3/onefuse/workspaces/1/",
        "title": "Default"
      },
      "policy": {
        "href": "/api/v3/onefuse/microsoftADPolicies/20/",
        "title": "microsoft_ad_policy_78092336"
      },
      "jobMetadata": {
        "href": "/api/v3/onefuse/jobMetadata/597/",
        "title": "Job Metadata Record id 597"
      }
    },
    "name": "testcompacct_490991",
    "id": 23,
    "state": "final",
    "buildOu": "",
    "finalOu": "OU=qa,OU=Environments,DC=example,DC=net",
    "securityGroups": [
      "CN=TestSecurityGroup,OU=OneFuse,DC=example,DC=net",
      "CN=TestSecurityGroupTemp,OU=TEMPOU1,DC=example,DC=net"
    ]
}`

const aMicrosoftADPolicy string = `{
    "_links": {
      "self": {
        "href": "/api/v3/onefuse/microsoftADPolicies/11/",
        "title": "microsoft_ad_policy_38181060"
      },
      "workspace": {
        "href": "/api/v3/onefuse/workspaces/1/",
        "title": "Default"
      },
      "microsoftEndpoint": {
        "href": "/api/v3/onefuse/endpoints/506/",
        "title": "QA_Microsoft_Endpoint_91672960"
      }
    },
    "name": "microsoft_ad_policy_38181060",
    "id": 11,
    "description": "microsoft active directory policy description",
    "microsoftEndpoint": "Endpoint: microsoft: QA_Microsoft_Endpoint_91672960: dc01.example.net:443",
    "computerNameLetterCase": "LOWER",
    "buildOu": "",
    "createBuildOu": true,
    "removeBuildOu": true,
    "ou": "OU=qa,OU=Environments,DC=example,DC=net",
    "createOrganizationalUnit": false,
    "removeOrganizationalUnit": false,
    "deleteComputerAccountsByName": true,
    "securityGroups": [
        "CN=TestSecurityGroup,OU=OneFuse,DC=example,DC=net",
        "CN=TestSecurityGroupTemp,OU=TEMPOU1,DC=example,DC=net"
    ]
}`

func responsesForADPolicy(i int) (string, int) {
	return bodyForGetADPolicy(i), missingTokenStatusPattern(i)
}

func bodyForGetADPolicy(i int) string {
	return missingTokenBodyPattern(
		aADPolicyList,
	)[i]
}

func responsesForComputerAccount(i int) (string, int) {
	return bodyForGetComputerAccount(i), missingTokenStatusPattern(i)
}

func bodyForGetComputerAccount(i int) string {
	return missingTokenBodyPattern(
		aComputerAccount,
	)[i]
}

func responsesForMicrosoftADPolicy(i int) (string, int) {
	return bodyForGetMicrosoftADPolicy(i), missingTokenStatusPattern(i)
}

func bodyForGetMicrosoftADPolicy(i int) string {
	return missingTokenBodyPattern(
		aMicrosoftADPolicy,
	)[i]
}
