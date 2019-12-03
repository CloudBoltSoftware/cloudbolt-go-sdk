# CloudBolt Go SDK

Welcome to the CloudBolt Golang SDK!

## Getting started

A quick example to get you started:

```go
import (
    "github.com/cloudbolt/cloudbolt-go-sdk/cbclient"
)

func main() {
    /*
    New(protocol, host, port, username, password)
    */
    client, err := cbclient.New("https", "cloudbolt.intranet", "443", "aUser", "aPassword")

    /*
    GetCloudBoltObject(object type, object name)
    */
    obj, err := client.GetCloudBoltObject("os-builds", "CentOS Linux 7 x86_64")

    /*
    GetGroup(group hierarchy path)
    */
    group, err := client.GetGroup("/Dev Org/Infra")

    /*
    DeployBlueprint(group path, blueprint path, resource name, [blueprint parameters])
    */
    bpParameters := map[string]interface{}{
        "some-param":  "param value",
        "other-param": "foo bar baz",
    }
    bpItem := map[string]interface{}{
        "bp-item-name":    "bp item name",
        "bp-item-params": bpParameters,
        "environment":     "bp environment",
        "osbuild":         "bp osbuild",
    }
    bpItems := []map[string]interface{}{
        bpItem,
    }
    order, err := client.DeployBlueprint("/api/v2/groups/GRP-y2n9xx58/", "/api/v2/blueprints/BP-6ic2tw7x/", "Name of New Resource", bpItems)

    /*
    GetOrder(order ID)
    */
    order, err := client.GetOrder("123")

    /*
    GetJob(job path)
    */
    job, err := client.GetJob("/api/v2/jobs/123/")

    /*
    GetResource(resource path)
    */
    resource, err := client.GetResource("/api/v2/resources/service/123/")

    /*
    GetServer(server path)
    */
    server, err := client.GetServer("/api/v2/servers/123/")

    /*
    SubmitAction(action path)
    */
    action, err := client.SubmitAction()

    /*
    DecomOrder(group path, environment path, [servers])
    */
    servers := []string{
        "/api/v2/servers/127/",
        "/api/v2/servers/513/",
        "/api/v2/servers/128/",
    }
    order, err := client.DecomOrder("/api/v2/groups/GRP-fsqpo11g/", "/api/v2/environments/ENV-mppchwvg/", servers)
}
```

## Testing

The quick answer to "how do I test this" is:

```sh
$ cd cbclient
[cbclient]$ go test
```

For a longer answer, read [TESTING.md](./TESTING.md).

## Updating Dependencies

If you want to make any changes to the dependencies, run `scripts/update-deps.sh`.
This will read `go.mod` and make any necessary changes to `go.sum`.
