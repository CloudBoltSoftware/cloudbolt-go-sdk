package cbclient

const anOrder string = `{
    "_links": {
        "self": {
            "href": "/api/v3/cmp/orders/ORD-e9v87uia/",
            "title": "Installation of My Simple Blueprint"
        },
        "group": {
            "href": "/api/v3/cloudbolt/groups/GRP-yfbbsfht/",
            "title": "My Org"
        },
        "owner": {
            "href": "/api/v3/cloudbolt/users/USR-mxpqe1x7/",
            "title": "user001"
        },
        "approvedBy": {
            "href": "/api/v3/cloudbolt/users/USR-mxpqe1x7/",
            "title": "user001"
        },
        "jobs": [
            {
                "href": "/api/v3/cmp/jobs/JOB-9nrax3gb/",
                "title": "Deploy Blueprint Job 1011"
            },
            {
                "href": "/api/v3/cmp/jobs/JOB-kb0tuw1e/",
                "title": "Provision Server Job 1012"
            }
        ],
        "duplicate": {
            "href": "/api/v3/cmp/orders/ORD-e9v87uia/duplicate/",
            "title": "Duplicate Order"
        }
    },
    "name": "Installation of My Simple Blueprint",
    "id": "ORD-e9v87uia",
    "status": "SUCCESS",
    "rate": "4.18/month",
    "createDate": "2022-04-10T10:04:04.218104",
    "approveDate": "2022-04-10T10:04:15.041024",
    "deploymentItems": [
        {
            "id": "OI-1p0bajs6",
            "resourceName": "My Simple Blueprint",
            "resourceParameters": {
                "bpParam1": "bp1 value",
                "bpParam2": "10",
                "bpParam3": "3.14",
                "bpParam4": "True"
            },
            "blueprint": {
                "href": "/api/v3/cmp/blueprints/BP-esnjtp7u/",
                "title": "My Simple Blueprint"
            },
            "blueprintItemsArguments": {
                "plugin-bdi-olk0xwve": {
                    "tierType": "plugin",
                    "parameters": {
                        "bpParam1": "bp1 value",
                        "bpParam2": "10",
                        "bpParam3": "3.14",
                        "bpParam4": "True"
                    }
                },
                "server-bdi-743tlxxu": {
                    "tierType": "server",
                    "parameters": {
                        "instanceType": "t2.nano",
                        "envParam2": "20",
                        "grpParam2": "30",
                        "bpParam1": "bp1 value",
                        "bpParam2": "10",
                        "envParam1": "env1 value",
                        "grpParam1": "grp1 value",
                        "bpParam3": "3.14",
                        "bpParam4": "True"
                    },
                    "environment": {
                        "href": "/api/v3/cmp/environments/ENV-1tytr2pu/",
                        "title": "(MY AWS Resource Handler) us-east-2 vpc-46c8382e"
                    },
                    "attributes": {
                        "quantity": 1
                    }
                }
            },
            "itemType": "blueprint"
        }
    ]
}`

const anOrderStatus string = `{
    "status": "FAILURE",
    "outputMessages": [
        "Job 101: Output for Job 101",
        "Job 102: Output for Job 102"
    ],
    "errorMessages": [
        "Job 101: Error for Job 101",
        "Job 102: Error for Job 102"
    ]
}`

func responsesForGetOrder(i int) (string, int) {
	return bodyForGetOrder(i), missingTokenStatusPattern(i)
}

func bodyForGetOrder(i int) string {
	return missingTokenBodyPattern(
		anOrder,
	)[i]
}

func responsesForGetOrderStatus(i int) (string, int) {
	return bodyForGetOrderStatus(i), missingTokenStatusPattern(i)
}

func bodyForGetOrderStatus(i int) string {
	return missingTokenBodyPattern(
		anOrderStatus,
	)[i]
}
