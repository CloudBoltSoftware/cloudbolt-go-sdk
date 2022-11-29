package cbclient

const aMicrosoftEndpointList string = `{
    "_links": {
        "self": {
            "href": "/api/v3/onefuse/endpoints/?page=1&filter=name%3AQA_Microsoft_Endpoint_47393314",
            "title": "List of Endpoints - Page 1 of 1"
        }
    },
    "total": 1,
    "count": 1,
    "_embedded": {
        "endpoints": [
            {
                "_links": {
                    "self": {
                        "href": "/api/v3/onefuse/endpoints/7/",
                        "title": "My_Microsoft_Endpoint"
                    },
                    "workspace": {
                        "href": "/api/v3/onefuse/workspaces/2/",
                        "title": "Default"
                    },
                    "credential": {
                        "href": "/api/v3/onefuse/moduleCredentials/11/",
                        "title": "My_Credentials"
                    }
                },
                "id": 7,
                "type": "microsoft",
                "name": "My_Microsoft_Endpoint",
                "description": "A Microsoft Endpoint",
                "singleThreaded": true,
                "host": "microsoftqa01.mydomain.net",
                "port": 443,
                "ssl": true,
                "jumpHost": null,
                "jumpPort": 22,
                "temporaryDirectoryForScripts": "C:\\Windows\\temp",
                "sharePathForTemporaryDirectory": ""
            }
        ]
    }
}`

func responsesForMicrosoftEndpoint(i int) (string, int) {
	return bodyForGetMicrosoftEndpoint(i), missingTokenStatusPattern(i)
}

func bodyForGetMicrosoftEndpoint(i int) string {
	return missingTokenBodyPattern(
		aMicrosoftEndpointList,
	)[i]
}
