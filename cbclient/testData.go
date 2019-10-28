package cbclient

const BodyForGetObject = `{
    "_links": {
        "self": {
            "href": "/api/v2/things/?page=1",
            "title": "List of Things - Page 1 of 1"
        }
    },
    "total": 1,
    "count": 1,
    "_embedded": [
        {
            "_links": {
                "self": {
                    "href": "/api/v2/things/XYZ-abcdefgh/",
                    "title": "Thing 2"
                }
            },
            "name": "Thing 2",
            "id": "3"
        }
    ]
}
`

// A slice of responses for the GetGroup test.
// This is a function because we cannot delcare const slices.
// Since it's a function we accept an index parameter `i` for coveneince.
func BodyForGetGroup(i int) string {
	return []string{`{
		    "_links": {
			"self": {
			    "href": "/api/v2/groups/?page=1",
			    "title": "List of Groups - Page 1 of 1"
			}
		    },
		    "total": 1,
		    "count": 1,
		    "_embedded": [
			{
			    "_links": {
				"self": {
				    "href": "/api/v2/groups/GRP-th3gr0up/",
				    "title": "the group"
				}
			    },
			    "name": "the group",
			    "id": "6"
			}
		    ]
		}`,
		`{
		    "_links": {
			"self": {
			    "href": "/api/v2/groups/GRP-th3gr0up/",
			    "title": "the group"
			},
			"parent": { },
			"subgroups": [],
			"environments": [],
			"orderable-environments": {
			    "href": "/api/v2/groups/GRP-th3gr0up/",
			    "title": "Orderable Environments For 'the group'"
			}
		    },
		    "name": "the group",
		    "id": "6",
		    "type": "Organization",
		    "rate": "0.00/month",
		    "auto-approval": false
		}`}[i]
}
