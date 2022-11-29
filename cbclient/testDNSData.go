package cbclient

const aDNSPolicyList string = `{
    "_links": {
        "self": {
            "href": "/api/v3/onefuse/dnsPolicies/?page=1&filter=name%3Amy_infoblox_dns_policy",
            "title": "List of Dns Policies - Page 1 of 1"
        }
    },
    "total": 1,
    "count": 1,
    "_embedded": {
        "dnsPolicies": [
            {
                "_links": {
                    "self": {
                        "href": "/api/v3/onefuse/dnsPolicies/9/",
                        "title": "my_infoblox_dns_policy"
                    },
                    "workspace": {
                        "href": "/api/v3/onefuse/workspaces/2/",
                        "title": "Default"
                    },
                    "endpoint": {
                        "href": "/api/v3/onefuse/endpoints/14/",
                        "title": "my_infoblox_dns_policy_Endpoint"
                    }
                },
                "id": 9,
                "type": "infoblox",
                "name": "my_infoblox_dns_policy",
                "description": "An Infoblox DNS Policy created through automated tests",
                "createARecord": true,
                "preValidateARecord": false,
                "postValidateARecord": false,
                "createPtrRecord": false,
                "preValidatePtrRecord": false,
                "postValidatePtrRecord": false,
                "createCNameRecord": false,
                "preValidateCNameRecord": false,
                "postValidateCNameRecord": false,
                "postValidationSleepSeconds": "5",
                "validationTimeoutSeconds": "180",
                "hostnameOverride": null,
                "createHostRecord": false,
                "removeFixedAddressRecord": false
            }
        ]
    }
}`

const aDNSReservation string = `{
  "_links": {
    "self": {
      "href": "/api/v3/dnsReservations/10/",
      "title": "test010"
    },
    "workspace": {
      "href": "/api/v3/onefuse/workspaces/1/",
      "title": "Default"
    },
    "policy": {
      "href": "/api/v3/onefuse/dnsPolicies/1/",
      "title": "infoblox1"
    },
    "jobMetadata": {
      "href": "/api/v3/onefuse/jobMetadata/86/",
      "title": "Job Metadata Record id 86"
    }
  },
  "name": "test010",
  "id": 10,
  "records": [
    {
      "type": "a",
      "name": "test010.infoblox.example.net",
      "value": "32.30.29.150"
    },
    {
      "type": "ptr",
      "name": "test010.infoblox.example.net",
      "value": "32.30.29.150"
    },
    {
      "type": "host",
      "name": "test010.infoblox.example.net",
      "value": "32.30.29.150"
    }
  ]
}`

func responsesForDNSPolicy(i int) (string, int) {
	return bodyForGetDNSPolicy(i), missingTokenStatusPattern(i)
}

func bodyForGetDNSPolicy(i int) string {
	return missingTokenBodyPattern(
		aDNSPolicyList,
	)[i]
}

func responsesForDNSReservation(i int) (string, int) {
	return bodyForGetDNSReservation(i), missingTokenStatusPattern(i)
}

func bodyForGetDNSReservation(i int) string {
	return missingTokenBodyPattern(
		aDNSReservation,
	)[i]
}
