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

func bpOrderItems() []map[string]interface{} {
	bpParameters := map[string]interface{}{
		"some-param": "param value",
		"other-param": "foo bar baz",
	}

	bpItem := map[string]interface{}{
		"bp-item-name":    "bp item name",
		"bp-item-paramas": bpParameters,
		"environment": "bp environment",
		"osbuild": "bp osbuild",
	}

	return []map[string]interface{}{ bpItem }
}

func BodyForDeployBlueprint(i int) string {
	return []string{`{
	    "_links": {
		"self": {
		    "href": "/api/v2/orders/101/",
		    "title": "Order id 101"
		},
		"group": {
		    "href": "/api/v2/groups/GRP-th3gr0up/",
		    "title": "the group"
		},
		"owner": {
		    "href": "/api/v2/users/42/",
		    "title": "the owner"
		},
		"approved-by": {
		    "href": "/api/v2/users/42/",
		    "title": "the owner"
		},
		"actions": {
	             "href": "/api/v2/actions/2019",
	             "title": "the action"
		},
		"jobs": [
		    {
			"href": "/api/v2/jobs/1234/",
			"title": "Job id 1234"
		    }
		]
	    },
	    "name": "the order",
	    "id": "1602",
	    "status": "ACTIVE",
	    "rate": "0.12/month",
	    "create-date": "2019-10-28T18:39:32.099576",
	    "approve-date": "2019-10-28T18:39:32.420531",
	    "items": {
		"deploy-items": [
		    {
			"blueprint": "/api/v2/blueprints/BP-ab1u3prt",
			"blueprint-items-arguments": {
			    "build-item-Server": {
				"attributes": {
					"hostname": "the hostname",
					"quantity": 1
				},
				"os-build": "/api/v2/os-builds/OSB-th3058ld/",
				"environment": "/api/v2/environments/ENV-th153nv5/",
				"parameters": { }
			    }
			},
			"resource-name": "the resource",
			"resource-parameters": {}
		    }
		]
	    }
	}`}[i]
}
