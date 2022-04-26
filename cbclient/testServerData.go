package cbclient

const aServer string = `{
    "_links": {
        "self": {
            "href": "/api/v3/cmp/servers/SVR-yrk09wht/",
            "title": "myawsinstance1"
        }
    },
    "id": "SVR-yrk09wht",
    "hostname": "myawsinstance1",
    "ipAddress": "3.17.176.101",
    "status": "ACTIVE",
    "mac": "02:99:e2:0f:18:b2",
    "powerStatus": "POWERON",
    "dateAddedToCloudbolt": "2022-04-08 11:55:07.056038",
    "cpuCount": 1,
    "memorySizeGb": "0.5000",
    "diskSizeGb": 8,
    "notes": "",
    "labels": [],
    "osFamily": "Amazon Linux",
    "rateBreakdown": {
        "total": "$ 4.18/month",
        "hardware": "$ 4.18/month",
        "software": "-",
        "extra": "-"
    },
    "attributes": [
        {
            "name": "delete_ebs_volumes_on_termination",
            "type": "BOOL",
            "value": true
        },
        {
            "name": "ebs_volume_type",
            "type": "STR",
            "value": "standard"
        }
    ],
    "credentials": {
        "username": "root",
        "password": "not set",
        "key": "cloudbolt-east2"
    },
    "disks": [
        {
            "uuid": "vol-037494719ec2192d1",
            "diskSize": 8,
            "name": "vol-037494719ec2192d1",
            "availabilityZone": "us-east-2a",
            "volumeType": "standard",
            "encrypted": false
        }
    ],
    "networks": [
        {
            "name": "",
            "network": "subnet-214ab049",
            "mac": "02:99:e2:0f:18:b2",
            "ip": "3.17.176.215",
            "privateIp": "172.31.13.128",
            "additionalIps": ""
        }
    ],
    "techSpecificAttributes": {
        "ec2Region": "us-east-2",
        "vpcId": "vpc-46c8382e",
        "instanceId": "i-094c9eed88b8acaa0",
        "instanceType": "t2.nano",
        "ipAddress": "3.17.176.215",
        "elasticIp": null,
        "privateIpAddress": "172.31.13.128",
        "publicDnsName": "ec2-3-17-176-215.us-east-2.compute.amazonaws.com",
        "privateDnsName": "ip-172-31-13-128.us-east-2.compute.internal",
        "availabilityZone": "us-east-2a",
        "keyName": "cloudbolt-east2",
        "profileArn": null,
        "hostId": null,
        "hostGroupArn": null,
        "securityGroupsJson": "[\"cloudbolt\"]",
        "tagsJson": "{\"Name\": \"myawsinstance1\"}",
        "type": "ec2_server_info"
    }
}`

const aDecomServerOrder string = `{
    "_links": {
        "self": {
            "href": "/api/v3/cmp/orders/ORD-ijudvhqv/",
            "title": "Deletion of myawsinstance2"
        },
        "group": {
            "href": "/api/v3/cloudbolt/groups/GRP-z8ki3ltv/",
            "title": "Default"
        },
        "owner": {
            "href": "/api/v3/cloudbolt/users/USR-jlspg3az/",
            "title": "user001"
        },
        "approvedBy": {
            "href": "/api/v3/cloudbolt/users/USR-jlspg3az/",
            "title": "user001"
        },
        "jobs": [
            {
                "href": "/api/v3/cmp/jobs/JOB-64leem1j/",
                "title": "Delete Server Job 54"
            }
        ],
        "duplicate": {
            "href": "/api/v3/cmp/orders/ORD-ijudvhqv/duplicate/",
            "title": "Duplicate Order"
        }
    },
    "name": "Deletion of myawsinstance2",
    "id": "ORD-ijudvhqv",
    "status": "SUCCESS",
    "rate": "0.00/month",
    "createDate": "2022-03-23T16:42:09.930071",
    "approveDate": "2022-03-23T16:42:10.284359",
    "deploymentItems": [
        {
            "id": "OI-7ilqk55l",
            "environment": {
                "href": "/api/v3/cmp/environments/ENV-su349w6z/",
                "title": "OCI Test 1 "
            },
            "servers": [
                {
                    "href": "/api/v3/cmp/servers/SVR-d7xr7for/",
                    "title": "myawsinstance2"
                }
            ],
            "itemType": "decomServer"
        }
    ]
}`

const aDecomServerJob string = `{
    "_links": {
        "self": {
            "href": "/api/v3/cmp/jobs/JOB-80uh0rmr/",
            "title": "Delete Server Job 502"
        },
        "owner": {
            "href": "/api/v3/cmp/users/USR-jlspg3az/",
            "title": "admin"
        },
        "parent": {},
        "subjobs": [],
        "prerequisite": {},
        "dependentJobs": [],
        "order": {
            "href": "/api/v3/cmp/orders/ORD-puqjyty7/",
            "title": "Deletion of myawsinstance1"
        },
        "servers": [
            {
                "href": "/api/v3/cmp/servers/SVR-668kqo0f/",
                "title": "myawsinstance1"
            }
        ]
    },
    "id": "JOB-80uh0rmr",
    "type": "decom",
    "status": "SUCCESS",
    "workerPid": 31,
    "workerHostname": "worker00@ecf3dd1fc72a",
    "canBeRequeued": true,
    "createdDate": "2022-04-06 08:09:03.576682",
    "updatedDate": "2022-04-06 08:09:06.192433",
    "startDate": "2022-04-06 08:09:04.474991",
    "endDate": "2022-04-06 08:09:06.192198",
    "output": "Deletion job completed successfully",
    "errors": "",
    "tasksDone": 1,
    "totalTasks": 1,
    "label": "",
    "executionState": ""
}`

func responsesForGetServer(i int) (string, int) {
	return bodyForGetServer(i), missingTokenStatusPattern(i)
}

func bodyForGetServer(i int) string {
	return missingTokenBodyPattern(
		aServer,
	)[i]
}

func responsesForDecomServerOrder(i int) (string, int) {
	return bodyForDecomServerOrder(i), missingTokenStatusPattern(i)
}

func bodyForDecomServerOrder(i int) string {
	return missingTokenBodyPattern(
		aDecomServerOrder,
	)[i]
}

func responsesForDecomServerJob(i int) (string, int) {
	return bodyForDecomServerJob(i), missingTokenStatusPattern(i)
}

func bodyForDecomServerJob(i int) string {
	return missingTokenBodyPattern(
		aDecomServerJob,
	)[i]
}
