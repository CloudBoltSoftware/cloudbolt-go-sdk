package cbclient

const BodyForGetObject = `{
    "_links": {
        "self": {
            "href": "/api/v2/thing1/?page=1",
            "title": "List of Thing1 - Page 1 of 1"
        }
    },
    "total": 1,
    "count": 1,
    "_embedded": [
        {
            "_links": {
                "self": {
                    "href": "/api/v2/thing1/6gd6uheu/",
                    "title": "TerraformEnvironment01"
                }
            },
            "name": "thing2",
            "id": "3"
        }
    ]
}
`
