package cbclient

// We follow a pattern in this file of the following:
// Declare a function `bodyFor<TEST NAME>` which returns a response for the i-th
// API query in <TEST NAME>.
// This tends to be implemented as:
// func bodyForTESTNAME(i int) string {
//     return []string{`response 1`,`response 2`}[i]
// }
// This is probably not the most performant or scalable, but it works for now.

/* Orders parse to the following struct:
CloudBoltOrder {
	Links:{
		Self:{
			Href:/api/v2/orders/ORDER-ID/
			Title:Order id ORDER-ID
		}
		Group:{
			Href:/api/v2/groups/GRP-ID/
			Title:GROUP NAME
		}
		Owner:{
			Href:/api/v2/users/2/
			Title:USER NAME
		}
		ApprovedBy:{
			Href:/api/v2/users/USER-ID/
			Title:USER NAME
		}
		Actions:{
			Href:/api/v2/actions/ACTION-ID/
			Title:ACTION NAME
		}
		Jobs:[
			{
				Href:/api/v2/jobs/JOB-ID/
				Title:Job id JOB-ID
			}
		]
	}
	Name:
	ID:ORDER-ID
	Status:ORDER-STATUS
	Rate:0.00/month
	CreateDate:2019-10-28T18:39:32.099576
	ApproveDate:2019-10-28T18:39:32.420531
	Items:{
		DeployItems:[
			{
				Blueprint:/api/v2/blueprints/BP-ID/
				BlueprintItemsArguments:{
					BuildItemBuildServer:{
						Attributes:{
							Hostname:
							Quantity:0
						}
						OsBuild:
						Environment:
						Parameters:map[]
					}
				}
				ResourceName:resource name
				ResourceParameters:{ }
			}
		]
	}
}
*/
const anOrder string = `{
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
			"href": "/api/v2/actions/2019/",
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
}`

// List of groups for the query:
// `/api/v2/groups/?filter=name:the+childgroup`
const listOfGroups string = `{
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
					"href": "/api/v2/groups/GRP-y3tan0thrgrp/",
					"title": "the childgroup"
				}
			},
			"name": "the childgroup",
			"id": "215"
		},
		{
			"_links": {
				"self": {
					"href": "/api/v2/groups/GRP-an0thrgrp/",
					"title": "the childgroup"
				}
			},
			"name": "the childgroup",
			"id": "512"
		}
	]
}`

const aGroup string = `{
	"_links": {
		"self": {
			"href": "/api/v2/groups/GRP-th3gr0up/",
			"title": "the group"
		},
		"parent": { },
		"subgroups": [
			{
				"href": "/api/v2/groups/GRP-an0thrgrp/",
				"title": "the subgroup"
			}
		],
		"environments": [
			{
				"href": "/api/v2/environments/ENV-th153nv5/",
				"title": "the environment"
			}
		],
		"orderable-environments": {
			"href": "/api/v2/groups/GRP-th3gr0up/",
			"title": "Orderable Environments For 'the group'"
		}
	},
	"name": "the group",
	"id": "6",
	"type": "Mega Organization",
	"rate": "0.00/month",
	"auto-approval": true
}`

const aSubGroup string = `{
	"_links": {
		"self": {
			"href": "/api/v2/groups/GRP-a5ubgrp/",
			"title": "the subgroup"
		},
		"parent": {
            "href": "/api/v2/groups/GRP-th3gr0up/",
            "title": "the group"
        },
		"subgroups": [
			{
				"href": "/api/v2/groups/GRP-an0thrgrp/",
				"title": "the subgroup"
			}
		],
		"environments": [],
		"orderable-environments": {
			"href": "/api/v2/groups/GRP-a5ubgrp/",
			"title": "Orderable Environments For 'the subgroup'"
		}
	},
	"name": "the subgroup",
	"id": "511",
	"type": "Super Organization",
	"rate": "0.00/month",
	"auto-approval": false
}`

const aChildGroup string = `{
	"_links": {
		"self": {
			"href": "/api/v2/groups/GRP-an0thrgrp/",
			"title": "the childgroup"
		},
		"parent": {
            "href": "/api/v2/groups/GRP-a5ubgrp/",
            "title": "the subgroup"
        },
		"subgroups": [],
		"environments": [],
		"orderable-environments": {
			"href": "/api/v2/groups/GRP-an0thrgrp/",
			"title": "Orderable Environments For 'the subgroup'"
		}
	},
	"name": "the childgroup",
	"id": "512",
	"type": "Super Duper Organization",
	"rate": "0.00/month",
	"auto-approval": false
}`

const yetAnotherGroup string = `{
	"_links": {
		"self": {
			"href": "/api/v2/groups/GRP-y3tan0thrgrp/",
			"title": "the childgroup"
		},
		"parent": {},
		"subgroups": [],
		"environments": [],
		"orderable-environments": {
			"href": "/api/v2/groups/GRP-y3tan0thrgrp/",
			"title": "Orderable Environments For 'the childgroup'"
		}
	},
	"name": "the childgroup",
	"id": "215",
	"type": "Just another Organization",
	"rate": "0.00/month",
	"auto-approval": false
}`

const anObject string = `{
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
}`

const aJob string = `{
	"_links": {
		"self": {
			"href": "/api/v2/jobs/1234/",
			"title": "Job id 1234"
		},
		"owner": {
			"href": "/api/v2/users/42/",
			"title": "the owner"
		},
		"parent": {},
		"subjobs": [
			{
				"href": "/api/v2/jobs/1235/",
				"title": "Job id 1235"
			}
		],
		"prerequisite": {},
		"dependent-jobs": [],
		"order": {
			"href": "/api/v2/orders/101/",
			"title": "Order id 101"
		},
		"resource": {
			"href": "/api/v2/resources/big_service/2048/",
			"title": "A Big Service 2048"
		},
		"servers": [
			{
				"href": "/api/v2/servers/128/",
				"title": "a-server-128"
			}
		],
		"log_urls": {
			"raw-log": "/api/v2/jobs/1234/log-download-txt/",
			"zip-log": "/api/v2/jobs/1234/log-download"
		}
	},
	"status": "SUCCESS",
	"type": "Deploy Blueprint",
	"progress": {
		"total-tasks": 2,
		"completed": 2,
		"messages": [
			"Deploying blueprint A Big Service.",
			"Executing build steps for 1 build items.",
			"Starting The server build item",
			"Created and waiting on Provision Server Job 1235",
			"Blueprint deployment succeeded"
		]
	},
	"start-date": "2019-10-29T18:49:54.777786",
	"end-date": "2019-10-29T18:50:06.556124",
	"output": "Blueprint deployment succeeded"
}`

const aResource string = `
{
	"_links": {
		"self": {
			"href": "/api/v2/resources/big_service/2048/",
			"title": "A Big Service 2048"
		},
		"blueprint": {
		"href": "/api/v2/blueprints/BP-ab1u3prt",
			"title": "a blueprint"
		},
		"owner": {
			"href": "/api/v2/users/42/",
			"title": "the owner"
		},
		"group": {
			"href": "/api/v2/groups/GRP-th3gr0up/",
			"title": "the group"
		},
		"resource-type": {
			"href": "/api/v2/resource-types/4096/",
			"title": "Big Service"
		},
		"actions": [
			{
				"Delete": {
					"href": "/api/v2/resources/big_service/2048/actions/1/",
					"title": "Run 'Delete' on 'A Big Service 2048'"
				}
			},
			{
				"Scale": {
					"href": "/api/v2/resources/big_service/2048/actions/2/",
					"title": "Run 'Scale' on 'A Big Service 2048'"
				}
			}
		],
		"jobs": {
			"href": "/api/v2/resources/big_service/2048/related-jobs/",
			"title": "Related Jobs For Resource 'A Big Service 2048'"
		},
		"history": {
			"href": "/api/v2/resources/big_service/2048/history/",
			"title": "History For Resource 'A Big Service 2048'"
		}
	},
	"name": "A Big Service 2048",
	"id": "2048",
	"status": "Historical",
	"install-date": "2019-10-29T20:46:34.093868"
}
`

const aServer string = `{
	"_links": {
		"self": {
			"href": "/api/v2/servers/128/",
			"title": "a-server-128"
		},
		"owner": {
			"href": "/api/v2/users/42/",
			"title": "the owner"
		},
		"group": {
			"href": "/api/v2/groups/GRP-th3gr0up/",
			"title": "the group"
		},
		"environment": {
			"href": "/api/v2/environments/ENV-th153nv5/",
			"title": "the environment"
		},
		"resource-handler": {
			"href": "/api/v2/resource-handlers/404/",
			"title": "Resource Handler Found...ish"
		},
		"actions": [
			{
				"power_on": {
					"href": "/api/v2/servers/128/actions/poweron/",
					"title": "Power on 'a-server-128'"
				}
			},
			{
				"power_off": {
					"href": "/api/v2/servers/128/actions/poweroff/",
					"title": "Power off 'a-server-128'"
				}
			},
			{
				"reboot": {
					"href": "/api/v2/servers/128/actions/reboot/",
					"title": "Reboot 'a-server-128'"
				}
			},
			{
				"refresh_info": {
					"href": "/api/v2/servers/128/actions/refresh-info/",
					"title": "Refresh Info for 'a-server-128'"
				}
			},
			{
				"Ad Hoc Script": {
					"href": "/api/v2/servers/128/actions/1/",
					"title": "Run 'Ad Hoc Script' on 'a-server-128'"
				}
			}
		],
		"jobs": {
			"href": "/api/v2/servers/128/related-jobs/",
			"title": "Related Jobs For Server 'a-server-128'"
		},
		"history": {
			"href": "/api/v2/servers/5/history/",
			"title": "History For Server 'a-server-128'"
		}
	},
	"hostname": "a-server-128",
	"power-status": "POWEROFF",
	"status": "ACTIVE",
	"ip": "1.2.3.4",
	"mac": "aa:bb:cc:dd:ee:ff",
	"date-added-to-cloudbolt": "2019-11-01T18:44:26.670691",
	"cpu-cnt": 3,
	"mem-size": "1.2500 GB",
	"disk-size": "56 GB",
	"os-family": "Linux -&gt; SomeOS",
	"labels": [],
	"credentials": {
		"username": "TotallyNotRoot",
		"password": "not set",
		"key": "A CLOUDBOLT KEY 123"
	},
	"disks": [
		{
			"uuid": "vol-0123456789abcdef1",
			"disk-size": 13,
			"name": "also-vol-0123456789abcdef1",
			"datastore": "a-datastore",
			"provisioning-type": "some-provisioning-type"
		}
	],
	"networks": [
		{
			"name": "NIC 0",
			"network": "myswitch",
			"mac": "00:11:22:33:44:55",
			"ip": "1.2.3.4",
			"private-ip": "5.6.7.8",
			"additional-ips": "9.10.11.12"
		}
	],
	"parameters": {},
	"tech-specific-details": {
		"uuid": "i-abcdefghijklmnopqrst",
		"zone": "aa-northsouth-2b",
		"security_group": "sg-12345678910111213141",
		"instance_type": "abc.a1-abcde.fempto",
		"keypair_name": "gopher-keypair"
	}
}`

const aSubmitActionResponseBody string = `{
	"run-action-job": {
		"self": {
			"href": "/api/v2/jobs/1234",
			"title": "foo"
		}
	}
}`

const anUnauthorizedResponseBody string = `{
	"Status": "401 Unauthorized"
}` // TODO: Make this accurate

const anAuthRequestResponseBody string = `{
	"token": "Testing Token"
}`

func missingTokenStatusPattern(i int) int {
	switch i {
	// The first time the user tries to authenticate, they get a 401 Unauthorized
	case 0:
		return 401
	// This is what we expect the status to be from every successful
	// GET and POST request made to the API with a valid Auth token
	default:
		return 200
	}
}

// Wraps a given variadic number of responses with the normal "request an Auth token" script.
func missingTokenBodyPattern(responses ...string) []string {
	return append(
		[]string{
			anUnauthorizedResponseBody,
			anAuthRequestResponseBody,
		},
		responses...,
	)
}

/*
HTTP response script for TestNew() API calls
*/
func responsesForNew(i int) (string, int) {
	return bodyForNew(i), 200
}

// Since New() makes no API calls, TestNew() should make no API calls as well.
// We still pass this because we need the test to be _able_ to make API calls.
// Those would just raise an error, which we want to catch in the tests.
func bodyForNew(i int) string {
	return []string{}[i]
}

/*
HTTP response script for TestAuthenticate() API calls
*/
func responsesForAuthenticate(i int) (string, int) {
	return bodyForAuthenticate(i), 200
}

func bodyForAuthenticate(i int) string {
	return []string{
		anAuthRequestResponseBody,
	}[i]
}

/*
HTTP response script for TestAuthWrappedRequest() API calls
*/
func responsesForAuthWrappedRequest(i int) (string, int) {
	return bodyForAuthWrappedRequest(i), missingTokenStatusPattern(i)
}

func bodyForAuthWrappedRequest(i int) string {
	return missingTokenBodyPattern(
		`{"foo": "bar"}`,
	)[i]
}

// Used to verify that when the request response is _not_ 401 or 403,
// We just return the HTTP response without requesting a new token.
func responsesForAuthWrappedRequestWithToken(i int) (string, int) {
	return bodyForAuthWrappedRequestWithToken(i), 200
}

// Since we are verifying that we make only the object request and not a token request,
// we are only returning the requested object, not wrapping the requested object
// in `missingTokenBodyPattern`.
func bodyForAuthWrappedRequestWithToken(i int) string {
	return []string{
		`{"foo": "bar"}`,
	}[i]
}

/*
HTTP response script for TestGetCloudBolObject() API calls
*/
func responseForGetCloudBoltObject(i int) (string, int) {
	return bodyForGetCloudBoltObject(i), missingTokenStatusPattern(i)
}

func bodyForGetCloudBoltObject(i int) string {
	return missingTokenBodyPattern(
		anObject,
	)[i]
}

/*
HTTP response script for TestGetGroup() API calls
*/
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

/*
HTTP response script for TestDeployBlueprint() API calls
*/
func responsesForDeployBlueprint(i int) (string, int) {
	return bodyForDeployBlueprint(i), missingTokenStatusPattern(i)
}

func bodyForDeployBlueprint(i int) string {
	return missingTokenBodyPattern(
		anOrder,
	)[i]
}

func bpOrderItems() []map[string]interface{} {
	bpParameters := map[string]interface{}{
		"some-param":  "param value",
		"other-param": "foo bar baz",
	}

	bpItem := map[string]interface{}{
		"bp-item-name":    "bp item name",
		"bp-item-paramas": bpParameters,
		"environment":     "bp environment",
		"osbuild":         "bp osbuild",
	}

	return []map[string]interface{}{
		bpItem,
	}
}

/*
HTTP response script for TestVerifyGroup() API calls
*/
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

/*
HTTP response script for TestGetOrder() API calls
*/
func responsesForGetOrder(i int) (string, int) {
	return bodyForGetOrder(i), missingTokenStatusPattern(i)
}

func bodyForGetOrder(i int) string {
	return missingTokenBodyPattern(
		anOrder,
	)[i]
}

/*
HTTP response script for TestGetJob() API calls
*/
func responsesForGetJob(i int) (string, int) {
	return bodyForGetJob(i), missingTokenStatusPattern(i)
}

func bodyForGetJob(i int) string {
	return missingTokenBodyPattern(
		aJob,
	)[i]
}

/*
HTTP response script for TestGetResource() API calls
*/
func responsesForGetResource(i int) (string, int) {
	return bodyForGetResource(i), missingTokenStatusPattern(i)
}

func bodyForGetResource(i int) string {
	return missingTokenBodyPattern(
		aResource,
	)[i]
}

/*
HTTP response script for TestGetServer() API calls
*/
func responsesForGetServer(i int) (string, int) {
	return bodyForGetServer(i), missingTokenStatusPattern(i)
}

func bodyForGetServer(i int) string {
	return missingTokenBodyPattern(
		aServer,
	)[i]
}

/*
HTTP response script for TestSubmitAction() API calls
*/
func responsesForSubmitAction(i int) (string, int) {
	return bodyForSubmitAction(i), missingTokenStatusPattern(i)
}

func bodyForSubmitAction(i int) string {
	return missingTokenBodyPattern(
		aSubmitActionResponseBody,
	)[i]
}

/*
HTTP response script for TestDecomOrder() API calls
*/
func responsesForDecomOrder(i int) (string, int) {
	return bodyForDecomOrder(i), missingTokenStatusPattern(i)
}

func bodyForDecomOrder(i int) string {
	return missingTokenBodyPattern(
		anOrder, // TODO: aDecomOrder?
	)[i]
}
