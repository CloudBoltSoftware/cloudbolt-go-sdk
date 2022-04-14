package cbclient

const aOSBuildList string = `{
    "_links": {
        "self": {
            "href": "/api/v3/cmp/osBuilds/?page=1&filter=pk%3A2",
            "title": "List of Os Builds - Page 1 of 1"
        }
    },
    "total": 1,
    "count": 1,
    "_embedded": {
        "osBuilds": [
            {
                "_links": {
                    "self": {
                        "href": "/api/v3/cmp/osBuilds/OSB-z69hjvki/",
                        "title": "amzn2-ami-hvm-2.0.20210721.2-x86_64-gp2"
                    },
                    "environments": [
                        {
                            "href": "/api/v3/cmp/environments/ENV-1tytr2pu/",
                            "title": "MY AWS Environment"
                        }
                    ],
                    "images": [
                        {
                            "href": "/api/v3/cmp/osImages/IMG-ua0vvi31/",
                            "title": "amzn2-ami-hvm-2.0.20210721.2-x86_64-gp2"
                        }
                    ]
                },
                "name": "amzn2-ami-hvm-2.0.20210721.2-x86_64-gp2",
                "id": "OSB-z69hjvki",
                "description": null,
                "osFamily": "Amazon Linux"
            }
        ]
    }
}`

const aOSBuild string = `{
    "_links": {
        "self": {
            "href": "/api/v3/cmp/osBuilds/OSB-z69hjvki/",
            "title": "amzn2-ami-hvm-2.0.20210721.2-x86_64-gp2"
        },
        "environments": [
            {
                "href": "/api/v3/cmp/environments/ENV-1tytr2pu/",
                "title": "(MY AWS Environment"
            }
        ],
        "images": [
            {
                "href": "/api/v3/cmp/osImages/IMG-ua0vvi31/",
                "title": "amzn2-ami-hvm-2.0.20210721.2-x86_64-gp2"
            }
        ]
    },
    "name": "amzn2-ami-hvm-2.0.20210721.2-x86_64-gp2",
    "id": "OSB-z69hjvki",
    "description": null,
    "osFamily": "Amazon Linux"
}`

func responsesForOSBuild(i int) (string, int) {
	return bodyForGetOSBuild(i), missingTokenStatusPattern(i)
}

func bodyForGetOSBuild(i int) string {
	return missingTokenBodyPattern(
		aOSBuildList,
	)[i]
}

func responsesForOSBuildById(i int) (string, int) {
	return bodyForGetOSBuildById(i), missingTokenStatusPattern(i)
}

func bodyForGetOSBuildById(i int) string {
	return missingTokenBodyPattern(
		aOSBuild,
	)[i]
}
