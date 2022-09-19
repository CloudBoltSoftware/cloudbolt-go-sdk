package cbclient

const aRenderedTemplate string = `{
    "value": "Environment: prod",
    "resolvedProperties": {
        "Datacenter": "2205",
        "Environment": "prod",
        "OS": "linux",
        "Application": "web",
        "dnsSuffix": "example.com",
        "ProjectCode": "pro"
    }
}`

func responsesForRenderedTemplate(i int) (string, int) {
	return bodyForRenderedTemplate(i), missingTokenStatusPattern(i)
}

func bodyForRenderedTemplate(i int) string {
	return missingTokenBodyPattern(
		aRenderedTemplate,
	)[i]
}
