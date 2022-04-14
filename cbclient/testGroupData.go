package cbclient

const listOfGroups string = `{
    "_links": {
        "self": {
            "href": "/api/v3/cmp/groups/?page=1&filter=name%3AMy+Org",
            "title": "List of Groups - Page 1 of 1"
        }
    },
    "total": 2,
    "count": 2,
    "_embedded": {
        "groups": [
			{
				"_links": {
					"self": {
						"href": "/api/v3/cmp/groups/GRP-zg550a1x/",
						"title": "the childgroup"
					}
				},
				"name": "the childgroup",
				"id": "GRP-zg550a1x",
				"type": "Organization",
				"rate": "0",
				"autoApproval": false,
				"parent": {}
			},
            {
				"_links": {
					"self": {
						"href": "/api/v3/cmp/groups/GRP-zg550a1z/",
						"title": "the childgroup"
					}
				},
				"name": "the childgroup",
				"id": "GRP-zg550a1z",
				"type": "Organization",
				"rate": "0",
				"autoApproval": false,
				"parent": {
					"href": "/api/v3/cmp/groups/GRP-uz64vfht/",
					"title": "the subgroup"
				}
            }
        ]
    }
}`

const aGroup string = `{
    "_links": {
        "self": {
            "href": "/api/v3/cmp/groups/GRP-yfbbsfht/",
            "title": "the group"
        }
    },
    "name": "the group",
    "id": "GRP-yfbbsfht",
    "type": "Organization",
    "rate": "8.3520000000",
    "autoApproval": false,
    "parent": {}
}`

const aSubGroup string = `{
    "_links": {
        "self": {
            "href": "/api/v3/cmp/groups/GRP-uz64vfht/",
            "title": "the subgroup"
        }
    },
    "name": "the subgroup",
    "id": "GRP-uz64vfht",
    "type": "Organization",
    "rate": "0",
    "autoApproval": false,
    "parent": {
        "href": "/api/v3/cmp/groups/GRP-yfbbsfht/",
        "title": "the group"
    }
}`

const aChildGroup string = `{
    "_links": {
        "self": {
            "href": "/api/v3/cmp/groups/GRP-zg550a1z/",
            "title": "the childgroup"
        }
    },
    "name": "the childgroup",
    "id": "GRP-zg550a1z",
    "type": "Organization",
    "rate": "0",
    "autoApproval": false,
    "parent": {
        "href": "/api/v3/cmp/groups/GRP-uz64vfht/",
        "title": "the subgroup"
    }
}`

const yetAnotherGroup string = `{

}`

func responsesForVerifyGroup(i int) (string, int) {
	return bodyForVerifyGroup(i), missingTokenStatusPattern(i)
}

func bodyForVerifyGroup(i int) string {
	return missingTokenBodyPattern(
		aChildGroup,
		aSubGroup,
		aGroup, // Necessary?
	)[i]
}

func responsesForGroupById(i int) (string, int) {
	return bodyForGetGroupById(i), missingTokenStatusPattern(i)
}

func bodyForGetGroupById(i int) string {
	return missingTokenBodyPattern(
		aGroup,
	)[i]
}

func responsesForGetGroup(i int) (string, int) {
	return bodyForGetGroup(i), missingTokenStatusPattern(i)
}

// bodyForGetGroup: A slice of responses for the GetGroup test.
// This is a function because we cannot delcare const slices.
// Since it's a function we accept an index parameter `i` for convenience.
func bodyForGetGroup(i int) string {
	return missingTokenBodyPattern(
		listOfGroups,
		yetAnotherGroup,
		aChildGroup,
		aSubGroup,
		aGroup, // Necessary?
	)[i]
}
