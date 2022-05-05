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
    GetGroup(group hierarchy path)
    */
    group, err := client.GetGroup("/Dev Org/Infra")

    /*
    GetGroupById(group Id)
    */
    group, err := client.GetGroupById("GRP-abcd1234")

    /*
    GetBlueprint(blueprint name)
    */
    bp, err := client.GetBlueprint("My Blueprint")

    /*
    GetBlueprintById(group Id)
    */
    group, err := client.GetBlueprintById("BP-abcd1234")


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
