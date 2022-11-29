package cbclient

const aNamingPolicyList string = `{
    "_links": {
        "self": {
            "href": "/api/v3/onefuse/namingPolicies/?page=1&filter=name%3AMy_Naming_Policy",
            "title": "List of Naming Policies - Page 1 of 1"
        }
    },
    "total": 1,
    "count": 1,
    "_embedded": {
        "namingPolicies": [
            {
                "_links": {
                    "self": {
                        "href": "/api/v3/onefuse/namingPolicies/1/",
                        "title": "My_Naming_Policy"
                    },
                    "workspace": {
                        "href": "/api/v3/onefuse/workspaces/2/",
                        "title": "Default"
                    },
                    "namingSequences": [
                        {
                            "href": "/api/v3/onefuse/namingSequences/61/",
                            "title": "My_NamingSequence"
                        }
                    ],
                    "validationPolicies": []
                },
                "name": "My_Naming_Policy",
                "id": 1,
                "template": "{{sequence.QA_BASE10_NamingSequence10807033}}",
                "description": "A Naming Policy",
                "dnsSuffix": "test.com"
            }
        ]
    }
}`

const aCustomName string = `{
    "_links": {
      "self": {
        "href": "/api/v3/onefuse/customNames/1/",
        "title": "ldw001"
      },
      "workspace": {
        "href": "/api/v3/onefuse/workspaces/1/",
        "title": "Default"
      },
      "policy": {
        "href": "/api/v3/onefuse/namingPolicies/1/",
        "title": "global"
      },
      "jobMetadata": {
        "href": "/api/v3/onefuse/jobMetadata/1/",
        "title": "Job Metadata Record id 1"
      }
    },
    "name": "ldw001",
    "id": 1,
    "dnsSuffix": "example.com"
}`

func responsesForNamingPolicy(i int) (string, int) {
	return bodyForGetNamingPolicy(i), missingTokenStatusPattern(i)
}

func bodyForGetNamingPolicy(i int) string {
	return missingTokenBodyPattern(
		aNamingPolicyList,
	)[i]
}

func responsesForGetCustomName(i int) (string, int) {
	return bodyForGetCustomName(i), missingTokenStatusPattern(i)
}

func bodyForGetCustomName(i int) string {
	return missingTokenBodyPattern(
		aCustomName,
	)[i]
}
