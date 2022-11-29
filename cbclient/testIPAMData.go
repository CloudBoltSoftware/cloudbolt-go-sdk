package cbclient

const aIPAMPolicyList string = `{
    "_links": {
        "self": {
            "href": "/api/v3/onefuse/ipamPolicies/?page=1&filter=name%3AMY_IPAM_POLICY",
            "title": "List of Ipam Policies - Page 1 of 1"
        }
    },
    "total": 1,
    "count": 1,
    "_embedded": {
        "ipamPolicies": [
            {
                "_links": {
                    "self": {
                        "href": "/api/v3/onefuse/ipamPolicies/3/",
                        "title": "MY_IPAM_POLICY"
                    },
                    "workspace": {
                        "href": "/api/v3/onefuse/workspaces/2/",
                        "title": "Default"
                    },
                    "endpoint": {
                        "href": "/api/v3/onefuse/endpoints/8/",
                        "title": "MY_IPAM_POLICY_Endpoint"
                    }
                },
                "id": 3,
                "type": "solarwinds",
                "name": "MY_IPAM_POLICY",
                "description": "An IPAM Policy created through automated tests",
                "hostnameOverride": "{{request.hostname}}",
                "updateConflictNameWithDns": false,
                "primaryDns": null,
                "secondaryDns": null,
                "dnsSuffix": null,
                "dnsSearchSuffixes": null,
                "nicLabel": null,
                "conflictNameTemplate": null,
                "subnets": [
                    {
                        "subnet": "10.192.50.0/24",
                        "gateway": "10.192.50.1",
                        "network": "mynetwork_10.192.50.0_24",
                        "netmask": "255.255.255.0"
                    }
                ]
            }
        ]
    }
}`

const aIPAMReservation string = `{
    "_links": {
      "self": {
        "href": "/api/v3/ipamReservations/10/",
        "title": "test010"
      },
      "workspace": {
        "href": "/api/v3/onefuse/workspaces/1/",
        "title": "Default"
      },
      "policy": {
        "href": "/api/v3/onefuse/ipamPolicies/1/",
        "title": "infoblox1"
      },
      "jobMetadata": {
        "href": "/api/v3/onefuse/jobMetadata/86/",
        "title": "Job Metadata Record id 86"
      }
    },
    "name": "test010",
    "id": 10,
    "primaryDNS": "8.8.8.8",
    "secondaryDns": "10.0.1.1",
    "dnsSuffix": "12.0.0.1",
    "dnsSearchSuffixes": "10.0.0.1",
    "nicLabel": "NIC 1",
    "subnet": "10.0.1.1/24",
    "gateway": "10.0.1.1",
    "network": "10.0.1.2",
    "netmask": "255.255.255.0"
}`

func responsesForIPAMPolicy(i int) (string, int) {
	return bodyForGetIPAMPolicy(i), missingTokenStatusPattern(i)
}

func bodyForGetIPAMPolicy(i int) string {
	return missingTokenBodyPattern(
		aIPAMPolicyList,
	)[i]
}

func responsesForIPAMReservation(i int) (string, int) {
	return bodyForGetIPAMReservation(i), missingTokenStatusPattern(i)
}

func bodyForGetIPAMReservation(i int) string {
	return missingTokenBodyPattern(
		aIPAMReservation,
	)[i]
}
