package cbclient

const aResource string = `{
    "_links": {
        "self": {
            "href": "/api/v3/cmp/resources/RSC-no9aztne/",
            "title": "My Simple Blueprint"
        },
        "resourceType": {
            "href": "/api/v3/cmp/resourceTypes/RT-bde9nds8/",
            "title": "service"
        },
        "blueprint": {
            "href": "/api/v3/cmp/blueprints/BP-esnjtp7u/",
            "title": "My Simple Blueprint"
        },
        "owner": {
            "href": "/api/v3/cloudbolt/users/USR-mxpqe1x7/",
            "title": "user001"
        },
        "group": {
            "href": "/api/v3/cloudbolt/groups/GRP-yfbbsfht/",
            "title": "My Org"
        },
        "jobs": [
            {
                "href": "/api/v3/cmp/jobs/JOB-9nrax3gb/",
                "title": "Deploy Blueprint Job 1011"
            },
            {
                "href": "/api/v3/cmp/jobs/JOB-t2js3lwf/",
                "title": "My Action Job 1013"
            },
            {
                "href": "/api/v3/cmp/jobs/JOB-8i53zztl/",
                "title": "My Simple Resource Action Job 1016"
            }
        ],
        "parentResource": {},
        "actions": [
            {
                "href": "/api/v3/cmp/resourceActions/RSA-hxfync2x/",
                "title": "Scale"
            },
            {
                "href": "/api/v3/cmp/resourceActions/RSA-aq3b3gxm/",
                "title": "My Resource Action"
            },
            {
                "href": "/api/v3/cmp/resourceActions/RSA-beim3g0e/",
                "title": "Delete"
            }
        ],
        "servers": [
            {
                "href": "/api/v3/cmp/servers/SVR-srb5y8r3/",
                "title": "myawsinstance"
            }
        ]
    },
    "name": "My Simple Blueprint",
    "id": "RSC-no9aztne",
    "created": "2022-04-10 10:04:15",
    "status": "ACTIVE",
    "attributes": [
        {
            "name": "bp_param1",
            "type": "STR",
            "value": "bp1 value"
        },
        {
            "name": "bp_param2",
            "type": "INT",
            "value": 10
        },
        {
            "name": "bp_param3",
            "type": "DEC",
            "value": 3.14
        },
        {
            "name": "bp_param4",
            "type": "BOOL",
            "value": true
        }
    ]
}`

const aResourceList string = `{
    "_links": {
        "self": {
            "href": "/api/v3/cmp/resoures/?page=1&filter=name%3AMy+Simple+Blueprint",
            "title": "List of Resources - Page 1 of 1"
        }
    },
    "total": 1,
    "count": 1,
    "_embedded": {
        "resources": [
            {
                "_links": {
                    "self": {
                        "href": "/api/v3/cmp/resources/RSC-no9aztne/",
                        "title": "My Simple Blueprint"
                    },
                    "resourceType": {
                        "href": "/api/v3/cmp/resourceTypes/RT-bde9nds8/",
                        "title": "service"
                    },
                    "blueprint": {
                        "href": "/api/v3/cmp/blueprints/BP-esnjtp7u/",
                        "title": "My Simple Blueprint"
                    },
                    "owner": {
                        "href": "/api/v3/cloudbolt/users/USR-mxpqe1x7/",
                        "title": "user001"
                    },
                    "group": {
                        "href": "/api/v3/cloudbolt/groups/GRP-yfbbsfht/",
                        "title": "My Org"
                    },
                    "jobs": [
                        {
                            "href": "/api/v3/cmp/jobs/JOB-9nrax3gb/",
                            "title": "Deploy Blueprint Job 1011"
                        },
                        {
                            "href": "/api/v3/cmp/jobs/JOB-t2js3lwf/",
                            "title": "My Action Job 1013"
                        },
                        {
                            "href": "/api/v3/cmp/jobs/JOB-8i53zztl/",
                            "title": "My Simple Resource Action Job 1016"
                        }
                    ],
                    "parentResource": {},
                    "actions": [
                        {
                            "href": "/api/v3/cmp/resourceActions/RSA-hxfync2x/",
                            "title": "Scale"
                        },
                        {
                            "href": "/api/v3/cmp/resourceActions/RSA-aq3b3gxm/",
                            "title": "My Resource Action"
                        },
                        {
                            "href": "/api/v3/cmp/resourceActions/RSA-beim3g0e/",
                            "title": "Delete"
                        }
                    ],
                    "servers": [
                        {
                            "href": "/api/v3/cmp/servers/SVR-srb5y8r3/",
                            "title": "myawsinstance"
                        }
                    ]
                },
                "name": "My Simple Blueprint",
                "id": "RSC-no9aztne",
                "created": "2022-04-10 10:04:15",
                "status": "ACTIVE",
                "attributes": [
                    {
                        "name": "bp_param1",
                        "type": "STR",
                        "value": "bp1 value"
                    },
                    {
                        "name": "bp_param2",
                        "type": "INT",
                        "value": 10
                    },
                    {
                        "name": "bp_param3",
                        "type": "DEC",
                        "value": 3.14
                    },
                    {
                        "name": "bp_param4",
                        "type": "BOOL",
                        "value": true
                    }
                ]
            }
        ]
    }
}
`

func responsesForGetResource(i int) (string, int) {
	return bodyForGetResource(i), missingTokenStatusPattern(i)
}

func bodyForGetResource(i int) string {
	return missingTokenBodyPattern(
		aResource,
	)[i]
}

func responsesForGetResourceById(i int) (string, int) {
	return bodyForGetResourceById(i), missingTokenStatusPattern(i)
}

func bodyForGetResourceById(i int) string {
	return missingTokenBodyPattern(
		aResource,
	)[i]
}

func responsesForGetResourceByName(i int) (string, int) {
	return bodyForGetResourceByName(i), missingTokenStatusPattern(i)
}

func bodyForGetResourceByName(i int) string {
	return missingTokenBodyPattern(
		aResourceList,
	)[i]
}
