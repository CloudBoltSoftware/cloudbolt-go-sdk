# Testing CloudBolt Go SDK

Thank you for looking into testing the CloudBolt Go SDK!

If you are interested in improving the tests, please read this doc.
It should explain the design for tests how to extend them.

## Overview

The tests for the CloudBolt Go SDK live in the `cbclient/` directory next to the client code.

```text
cbclient/
├── api_client.go       API Client code
├── api_client_test.go  Tests for API Client
└── testData.go         Test Data used by API client
```

## Running tests

To run the tests, make sure you have `go>=0.13.0` installed and ready.

After that, run the following in your shell:

```sh
go test ./cbclient/
```

Success looks like this:

```sh
$ go test ./cbclient/
ok      github.com/cloudboltsoftware/cloudbolt-go-sdk/cbclient/cbclient 0.434s
```

and a failure looks like this:

```sh
$ go test ./cbclient/
--- FAIL: TestSomeFunction (0.00s)
    api_client_test.go:547:
        Expected
            <nil>: aValue
        to be aDifferentValue
FAIL
FAIL    github.com/cloudboltsoftware/cloudbolt-go-sdk/cbclient/cbclient 0.195s
FAIL
```

## Anatomy of a test

The tests tend to follow a similar structure, presented in go-pseudoscope in the following sections.

The general pattern is this:

1. Register the test.

2. Create a simple HTTP server which serves scripted responses. e.g., request
   1 to the server will return response from a list, request 2 will get back
   response 2, etc.

3. Create an API client object.

4. Call some client function(s).

5. Make assertions about the requests made.

6. Make some assertions about the client function(s) resulting object.

### Main test function

```go
func TestSomeFunction(t *testing.T) {
    // Register the test with our framework
    RegisterTestingT(t)

    // Spawn a testing server with a function for scripted responses
    // Below we go into what `responsesForSomeFunction` should look like
    server, requests := mockServer(responsesForSomeFunction)
    Expect(server).NotTo(BeNil())
    Expect(requests).NoTo(BeNil())

    // Create a CloudBolt client with the server
    client := getClient(server)
    Expect(client).NoTo(BeNil())

    // Perform some action
    obj, err := client.SomeFunction(someInput)
    Expect(err).NotTo(HaveOccured())
    Expect(obj).NotTo(BeNil())

    // Make assertions about the requests made by the client
    Expect(((*requests)[0].Property).To(Equal(expectedValue))
    ...

    // Make assertions about the object we got back
    Expect(obj.Property).To(Equal(someExpectedValue))
    ...
}
```

### Scripted responses function

Our tests design pattern only works because we know ahead of time what requests `client.SomeFunction()` should make at run-time.
For example: we know based on our implementation of `client.SomeFunction()` that it should make 2 requests.
Knowing this we need to make sure our `mockServer` returns 2 correct-looking payloads.

To do this declare a function that generates responses for our `mockServer()` to serve.

That function is passed as an input to `mockServer()` and is called on each HTTP request.
e.g., `responsesForSomeFunction()` is called in our mockServer as `responsesForSomeFunction(request#)`.

Response functions are declared in `testData.go` and also tend to look the same.

```go
// Declare some response strings
const someResponse string = `{
    "some_key": "some_value",
    "another_key": "more_values"
}`
const someOtherResponse string = `{
    "a_key": "another_value"
}`

// Write a function that accepts as input an index (i) and returns a string (body) and int (status code).
// This is the response for that tests's i-th request response.
func responsesForSomeFunction(i int) (string, int) {
    return bodyForSomeFunction(i), statusForSomeFunction(i)
}

// Bodies for the i-th request response.
func bodyForSomeFunction(i int) string {
    return []string{
        someResponse,
        someOtherResponse,
    }[i]
}

// Status codes for the i-th request response.
func statusForSomeFunction(i int) int {
    return []int{
        401,
        200
    }[i]
}
```
