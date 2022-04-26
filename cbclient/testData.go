package cbclient

// We follow a pattern in this file of the following:
// Declare a function `bodyFor<TEST NAME>` which returns a response for the i-th
// API query in <TEST NAME>.
// This tends to be implemented as:
// func bodyForTESTNAME(i int) string {
//     return []string{`response 1`,`response 2`}[i]
// }
// This is probably not the most performant or scalable, but it works for now.

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
