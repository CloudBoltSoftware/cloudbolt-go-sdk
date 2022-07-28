package cbclient

const aResourceHandlerList string = `{
    "_links": {
        "self": {
            "href": "/api/v3/cmp/resourceHandlers/?page=1&filter=name%3AMy+Test+Resource+HAndler",
            "title": "List of Resource Handlers - Page 1 of 1"
        }
    },
    "total": 1,
    "count": 1,
    "_embedded": {
        "resourceHandlers": [
            {
                "_links": {
                    "self": {
                        "href": "/api/v3/cmp/resourceHandlers/RH-nza16uyn/",
                        "title": "My Test Resource Handler"
                    }
                },
                "name": "My Test Resource Handler",
                "id": "RH-nza16uyn",
                "type": "AWS resource handler"
            }
        ]
    }
}
`

const aResourceHandler = `{
    "_links": {
        "self": {
            "href": "/api/v3/cmp/resourceHandlers/RH-nza16uyn/",
            "title": "My Test Resource Handler"
        }
    },
    "name": "My Test Resource Handler",
    "id": "RH-nza16uyn",
    "type": "AWS resource handler"
}`

func responsesForResourceHandler(i int) (string, int) {
	return bodyForGetResourceHandler(i), missingTokenStatusPattern(i)
}

func bodyForGetResourceHandler(i int) string {
	return missingTokenBodyPattern(
		aResourceHandlerList,
	)[i]
}

func responsesForResourceHandlerById(i int) (string, int) {
	return bodyForGetResourceHandlerById(i), missingTokenStatusPattern(i)
}

func bodyForGetResourceHandlerById(i int) string {
	return missingTokenBodyPattern(
		aResourceHandler,
	)[i]
}
