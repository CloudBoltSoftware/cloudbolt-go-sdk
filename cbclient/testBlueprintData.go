package cbclient

const aBlueprint string = `{
    "_links": {
        "self": {
            "href": "/api/v3/cmp/blueprints/BP-esnjtp7u/",
            "title": "My Simple Blueprint"
        },
        "export": {
            "href": "/api/v3/cmp/blueprints/BP-esnjtp7u/export/",
            "title": "Export My Simple Blueprint"
        },
        "deploy": {
            "href": "/api/v3/cmp/blueprints/BP-esnjtp7u/deploy/",
            "title": "Deploy My Simple Blueprint"
        },
        "deploymentSchema": {
            "href": "/api/v3/cmp/blueprints/BP-esnjtp7u/deploymentSchema/",
            "title": "Deployment Schema for My Simple Blueprint"
        },
        "samplePayload": {
            "href": "/api/v3/cmp/blueprints/BP-esnjtp7u/samplePayload/",
            "title": "Sample Playload to deploy My Simple Blueprint"
        },
        "groupsThatCanManage": [
            {
                "href": "/api/v3/cmp/groups/GRP-z8ki3ltv/",
                "title": "Default"
            }
        ],
        "groupsThatCanDeploy": [
            {
                "href": "/api/v3/cmp/groups/GRP-z8ki3ltv/",
                "title": "Default"
            },
            {
                "href": "/api/v3/cmp/groups/GRP-u7a3tbab/",
                "title": "Dept 1"
            },
            {
                "href": "/api/v3/cmp/groups/GRP-uz64vfht/",
                "title": "Dept 2"
            },
            {
                "href": "/api/v3/cmp/groups/GRP-zg550a1z/",
                "title": "Dept 2a"
            },
            {
                "href": "/api/v3/cmp/groups/GRP-yfbbsfht/",
                "title": "My Org"
            }
        ],
        "managementActions": [],
        "resourceType": {
            "href": "/api/v3/cmp/resourceTypes/RT-bde9nds8/",
            "title": "service",
            "label": "Service"
        }
    },
    "name": "My Simple Blueprint",
    "id": "BP-esnjtp7u",
    "description": "",
    "anyGroupCanDeploy": true,
    "sequence": 0,
    "favorited": false,
    "resourceNameTemplate": null,
    "isOrderable": true,
    "autoHistoricalResources": false,
    "showRecipientFieldOnOrderForm": false,
    "remoteSourceUrl": "",
    "lastCached": null,
    "isManageable": true,
    "blueprintImage": null,
    "labels": [
        {
            "name": "cat1"
        },
        {
            "name": "cat2"
        }
    ],
    "deploymentItems": [
        {
            "id": "BDI-743tlxxu",
            "name": "My Instance",
            "description": "",
            "deploySeq": 1,
            "executeInParallel": false,
            "showOnOrderForm": true,
            "restrictApplications": false,
            "hostnameTemplate": "laltomarvm00X",
            "allEnvironmentsEnabled": false,
            "osBuild": null,
            "allowedOsFamilies": null,
            "applications": null,
            "environmentSelectionOrchestration": null,
            "tierType": "server"
        },
        {
            "id": "BDI-olk0xwve",
            "name": "My Action",
            "description": null,
            "deploySeq": 2,
            "executeInParallel": false,
            "showOnOrderForm": true,
            "actionName": "My Action",
            "continueOnFailure": false,
            "runOnScaleUp": true,
            "tierType": "plugin"
        }
    ],
    "teardownItems": [],
    "osFamilies": [],
    "needsConfiguration": false,
    "orderCount": 51,
    "status": "ACTIVE"
	}`

const aBlueprintList string = `{
    "_links": {
        "self": {
            "href": "/api/v3/cmp/blueprints/?page=1&filter=name%3AMy+Simple+Blueprint",
            "title": "List of Blueprints - Page 1 of 1"
        }
    },
    "total": 1,
    "count": 1,
    "_embedded": {
        "blueprints": [
            {
                "_links": {
                    "self": {
                        "href": "/api/v3/cmp/blueprints/BP-esnjtp7u/",
                        "title": "My Simple Blueprint"
                    },
                    "export": {
                        "href": "/api/v3/cmp/blueprints/BP-esnjtp7u/export/",
                        "title": "Export My Simple Blueprint"
                    },
                    "deploy": {
                        "href": "/api/v3/cmp/blueprints/BP-esnjtp7u/deploy/",
                        "title": "Deploy My Simple Blueprint"
                    },
                    "deploymentSchema": {
                        "href": "/api/v3/cmp/blueprints/BP-esnjtp7u/deploymentSchema/",
                        "title": "Deployment Schema for My Simple Blueprint"
                    },
                    "samplePayload": {
                        "href": "/api/v3/cmp/blueprints/BP-esnjtp7u/samplePayload/",
                        "title": "Sample Playload to deploy My Simple Blueprint"
                    },
                    "groupsThatCanManage": [
                        {
                            "href": "/api/v3/cmp/groups/GRP-z8ki3ltv/",
                            "title": "Default"
                        }
                    ],
                    "groupsThatCanDeploy": [
                        {
                            "href": "/api/v3/cmp/groups/GRP-z8ki3ltv/",
                            "title": "Default"
                        },
                        {
                            "href": "/api/v3/cmp/groups/GRP-u7a3tbab/",
                            "title": "Dept 1"
                        },
                        {
                            "href": "/api/v3/cmp/groups/GRP-uz64vfht/",
                            "title": "Dept 2"
                        },
                        {
                            "href": "/api/v3/cmp/groups/GRP-zg550a1z/",
                            "title": "Dept 2a"
                        },
                        {
                            "href": "/api/v3/cmp/groups/GRP-yfbbsfht/",
                            "title": "My Org"
                        }
                    ],
                    "managementActions": [],
                    "resourceType": {
                        "href": "/api/v3/cmp/resourceTypes/RT-bde9nds8/",
                        "title": "service",
                        "label": "Service"
                    }
                },
                "name": "My Simple Blueprint",
                "id": "BP-esnjtp7u",
                "description": "",
                "anyGroupCanDeploy": true,
                "sequence": 0,
                "favorited": false,
                "resourceNameTemplate": null,
                "isOrderable": true,
                "autoHistoricalResources": false,
                "showRecipientFieldOnOrderForm": false,
                "remoteSourceUrl": "",
                "lastCached": null,
                "isManageable": true,
                "blueprintImage": null,
                "labels": [
                    {
                        "name": "cat1"
                    },
                    {
                        "name": "cat2"
                    }
                ],
                "deploymentItems": [
                    {
                        "id": "BDI-743tlxxu",
                        "name": "My Instance",
                        "description": "",
                        "deploySeq": 1,
                        "executeInParallel": false,
                        "showOnOrderForm": true,
                        "restrictApplications": false,
                        "hostnameTemplate": "laltomarvm00X",
                        "allEnvironmentsEnabled": false,
                        "osBuild": null,
                        "allowedOsFamilies": null,
                        "applications": null,
                        "environmentSelectionOrchestration": null,
                        "tierType": "server"
                    },
                    {
                        "id": "BDI-olk0xwve",
                        "name": "My Action",
                        "description": null,
                        "deploySeq": 2,
                        "executeInParallel": false,
                        "showOnOrderForm": true,
                        "actionName": "My Action",
                        "continueOnFailure": false,
                        "runOnScaleUp": true,
                        "tierType": "plugin"
                    }
                ],
                "teardownItems": [],
                "osFamilies": [],
                "needsConfiguration": false,
                "orderCount": 51,
                "status": "ACTIVE"
            }
        ]
    }
}`

func responsesForBlueprint(i int) (string, int) {
	return bodyForGetBlueprint(i), missingTokenStatusPattern(i)
}

func bodyForGetBlueprint(i int) string {
	return missingTokenBodyPattern(
		aBlueprintList,
	)[i]
}

func responsesForBlueprintById(i int) (string, int) {
	return bodyForGetBlueprintById(i), missingTokenStatusPattern(i)
}

func bodyForGetBlueprintById(i int) string {
	return missingTokenBodyPattern(
		aBlueprint,
	)[i]
}

func responsesForDeployBlueprint(i int) (string, int) {
	return bodyForDeployBlueprint(i), missingTokenStatusPattern(i)
}

func bodyForDeployBlueprint(i int) string {
	return missingTokenBodyPattern(
		anOrder,
	)[i]
}
