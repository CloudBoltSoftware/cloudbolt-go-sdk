package cbclient

const aStaticPropertySetList string = `{
    "_links": {
        "self": {
            "href": "/api/v3/onefuse/propertySets/?page=1&filter=name%3AMy_Static_Property_Set",
            "title": "List of Property Sets - Page 1 of 1"
        }
    },
    "total": 1,
    "count": 1,
    "_embedded": {
        "propertySets": [
            {
                "_links": {
                    "self": {
                        "href": "/api/v3/onefuse/propertySets/168/",
                        "title": "My_Static_Property_Set"
                    },
                    "workspace": {
                        "href": "/api/v3/onefuse/workspaces/2/",
                        "title": "Default"
                    }
                },
                "properties": {
                    "product": "OneFuse",
                    "organization": "CloudBolt Software"
                },
                "id": 168,
                "name": "My_Static_Property_Set",
                "description": "My Static Property Set"
            }
        ]
    }
}`

func responsesForStaticPropertySet(i int) (string, int) {
	return bodyForGetStaticPropertySet(i), missingTokenStatusPattern(i)
}

func bodyForGetStaticPropertySet(i int) string {
	return missingTokenBodyPattern(
		aStaticPropertySetList,
	)[i]
}
