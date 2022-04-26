package cbclient

const aEnvironment string = `{
	"_links": {
		"self": {
			"href": "/api/v3/cmp/environments/ENV-1tytr2pu/",
			"title": "MY AWS Environment"
		},
		"resourceHandler": {
			"href": "/api/v3/cmp/resourceHandlers/aws/RH-amtie2vv/",
			"title": "MY AWS Resource Handler"
		},
		"techSpecificParameters": {
			"href": "/api/v3/cmp/environments/ENV-1tytr2pu/techSpecificParameters/",
			"title": "Tech-Specific options for MY AWS Environment"
		},
		"techSpecificParameters:refresh": {
			"href": "/api/v3/cmp/environments/ENV-1tytr2pu/techSpecificParameters:refresh/",
			"title": "Synchronize Tech-Specific options for MY AWS Environment"
		},
		"techSpecificParameters:remove": {
			"href": "/api/v3/cmp/environments/ENV-1tytr2pu/techSpecificParameters:remove/",
			"title": "Remove Tech-Specific Fields for MY AWS Environment"
		},
		"networks": {
			"href": "/api/v3/cmp/environments/ENV-1tytr2pu/networks/",
			"title": "Networks for MY AWS Environment"
		},
		"osBuilds": [
			{
				"href": "/api/v3/cmp/osBuilds/OSB-z69hjvki/",
				"title": "amzn2-ami-hvm-2.0.20210721.2-x86_64-gp2"
			},
			{
				"href": "/api/v3/cmp/osBuilds/OSB-ttr56map/",
				"title": "CentOS 7.9.2009 x86_64"
			}
		],
		"groups": [
			{
				"href": "/api/v3/cmp/groups/GRP-z8ki3ltv/",
				"title": "Default"
			},
			{
				"href": "/api/v3/cmp/groups/GRP-yfbbsfht/",
				"title": "My Org"
			}
		]
	},
	"name": "MY AWS Environment",
	"id": "ENV-1tytr2pu",
	"description": null,
	"autoApproval": false,
	"serverQuota": "Unlimited",
	"rateQuota": "Unlimited",
	"cpuQuota": "Unlimited",
	"memoryQuota": "Unlimited",
	"diskQuota": "Unlimited",
	"awsRegion": "us-east-2",
	"vpcId": "vpc-123xyz",
	"techTypeSlug": "aws",
	"techSpecificParameters": [
		{
			"name": "aws_availability_zone",
			"type": "String",
			"label": "Availability Zone",
			"options": [
				"us-east-2a",
				"us-east-2b",
				"us-east-2c"
			]
		},
		{
			"name": "instance_type",
			"type": "String",
			"label": "Instance type",
			"options": [
				"t2.nano",
				"t2.small"
			]
		},
		{
			"name": "key_name",
			"type": "String",
			"label": "Key pair name",
			"options": [
				"Harvard",
				"cloudbolt-east2",
				"laltomardev"
			]
		},
		{
			"name": "sec_groups",
			"type": "String",
			"label": "Security groups",
			"options": [
				"cloudbolt",
				"default"
			]
		},
		{
			"name": "aws_elastic_ip",
			"type": "String",
			"label": "Elastic IP",
			"options": []
		},
		{
			"name": "delete_ebs_volumes_on_termination",
			"type": "Boolean",
			"label": "Auto-Delete EBS Volumes on Termination",
			"options": [
				"True"
			]
		},
		{
			"name": "ebs_volume_type",
			"type": "String",
			"label": "EBS Volume Type",
			"options": [
				"standard",
				"gp2",
				"io1",
				"st1",
				"sc1",
				"gp3"
			]
		},
		{
			"name": "iops",
			"type": "Integer",
			"label": "IOPS",
			"options": []
		},
		{
			"name": "aws_host",
			"type": "String",
			"label": "Dedicated Host",
			"options": []
		},
		{
			"name": "aws_host_group",
			"type": "String",
			"label": "Dedicated Host Group",
			"options": []
		}
	]
}`

const aEnvironmentList string = `{
    "_links": {
        "self": {
            "href": "/api/v3/cmp/environments/?page=1&filter=name%3A%28MY+AWS+Resource+Handler%29+us-east-2+vpc-123xyz",
            "title": "List of Environments - Page 1 of 1"
        }
    },
    "total": 1,
    "count": 1,
    "_embedded": {
        "environments": [
            {
                "_links": {
                    "self": {
                        "href": "/api/v3/cmp/environments/ENV-1tytr2pu/",
                        "title": "MY AWS Environment"
                    },
                    "resourceHandler": {
                        "href": "/api/v3/cmp/resourceHandlers/aws/RH-amtie2vv/",
                        "title": "MY AWS Resource Handler"
                    },
                    "techSpecificParameters": {
                        "href": "/api/v3/cmp/environments/ENV-1tytr2pu/techSpecificParameters/",
                        "title": "Tech-Specific options for MY AWS Environment"
                    },
                    "techSpecificParameters:refresh": {
                        "href": "/api/v3/cmp/environments/ENV-1tytr2pu/techSpecificParameters:refresh/",
                        "title": "Synchronize Tech-Specific options for MY AWS Environment"
                    },
                    "techSpecificParameters:remove": {
                        "href": "/api/v3/cmp/environments/ENV-1tytr2pu/techSpecificParameters:remove/",
                        "title": "Remove Tech-Specific Fields for MY AWS Environment"
                    },
                    "networks": {
                        "href": "/api/v3/cmp/environments/ENV-1tytr2pu/networks/",
                        "title": "Networks for MY AWS Environment"
                    },
                    "osBuilds": [
                        {
                            "href": "/api/v3/cmp/osBuilds/OSB-z69hjvki/",
                            "title": "amzn2-ami-hvm-2.0.20210721.2-x86_64-gp2"
                        },
                        {
                            "href": "/api/v3/cmp/osBuilds/OSB-ttr56map/",
                            "title": "CentOS 7.9.2009 x86_64"
                        }
                    ],
                    "groups": [
                        {
                            "href": "/api/v3/cmp/groups/GRP-z8ki3ltv/",
                            "title": "Default"
                        },
                        {
                            "href": "/api/v3/cmp/groups/GRP-yfbbsfht/",
                            "title": "My Org"
                        }
                    ]
                },
                "name": "MY AWS Environment",
                "id": "ENV-1tytr2pu",
                "description": null,
                "autoApproval": false,
                "serverQuota": "Unlimited",
                "rateQuota": "Unlimited",
                "cpuQuota": "Unlimited",
                "memoryQuota": "Unlimited",
                "diskQuota": "Unlimited",
                "awsRegion": "us-east-2",
                "vpcId": "vpc-123xyz",
                "techTypeSlug": "aws",
                "techSpecificParameters": [
                    {
                        "name": "aws_availability_zone",
                        "type": "String",
                        "label": "Availability Zone",
                        "options": [
                            "us-east-2a",
                            "us-east-2b",
                            "us-east-2c"
                        ]
                    },
                    {
                        "name": "instance_type",
                        "type": "String",
                        "label": "Instance type",
                        "options": [
                            "t2.nano",
							"t2.small"
                        ]
                    },
                    {
                        "name": "key_name",
                        "type": "String",
                        "label": "Key pair name",
                        "options": [
                            "Harvard",
                            "cloudbolt-east2",
                            "laltomardev"
                        ]
                    },
                    {
                        "name": "sec_groups",
                        "type": "String",
                        "label": "Security groups",
                        "options": [
                            "cloudbolt",
                            "default"
                        ]
                    },
                    {
                        "name": "aws_elastic_ip",
                        "type": "String",
                        "label": "Elastic IP",
                        "options": []
                    },
                    {
                        "name": "delete_ebs_volumes_on_termination",
                        "type": "Boolean",
                        "label": "Auto-Delete EBS Volumes on Termination",
                        "options": [
                            "True"
                        ]
                    },
                    {
                        "name": "ebs_volume_type",
                        "type": "String",
                        "label": "EBS Volume Type",
                        "options": [
                            "standard",
                            "gp2",
                            "io1",
                            "st1",
                            "sc1",
                            "gp3"
                        ]
                    },
                    {
                        "name": "iops",
                        "type": "Integer",
                        "label": "IOPS",
                        "options": []
                    },
                    {
                        "name": "aws_host",
                        "type": "String",
                        "label": "Dedicated Host",
                        "options": []
                    },
                    {
                        "name": "aws_host_group",
                        "type": "String",
                        "label": "Dedicated Host Group",
                        "options": []
                    }
                ]
            }
        ]
    }
}`

func responsesForEnvironment(i int) (string, int) {
	return bodyForGetEnvironment(i), missingTokenStatusPattern(i)
}

func bodyForGetEnvironment(i int) string {
	return missingTokenBodyPattern(
		aEnvironmentList,
	)[i]
}

func responsesForEnvironmentById(i int) (string, int) {
	return bodyForGetEnvironmentById(i), missingTokenStatusPattern(i)
}

func bodyForGetEnvironmentById(i int) string {
	return missingTokenBodyPattern(
		aEnvironment,
	)[i]
}
