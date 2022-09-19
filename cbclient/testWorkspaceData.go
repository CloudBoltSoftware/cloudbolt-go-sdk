package cbclient

const aWorkspaceList string = `{
    "_links": {
        "self": {
            "href": "/api/v3/onefuse/workspaces/?page=1",
            "title": "List of Workspaces - Page 1 of 1"
        }
    },
    "total": 1,
    "count": 1,
    "_embedded": {
        "workspaces": [
            {
                "_links": {
                    "self": {
                        "href": "/api/v3/onefuse/workspaces/2/",
                        "title": "Default"
                    }
                },
                "id": 2,
                "name": "Default"
            }
        ]
    }
}`

func responsesForWorkspaceList(i int) (string, int) {
	return bodyForGetWorkspace(i), missingTokenStatusPattern(i)
}

func bodyForGetWorkspace(i int) string {
	return missingTokenBodyPattern(
		aWorkspaceList,
	)[i]
}
