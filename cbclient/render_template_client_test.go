package cbclient

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestRenderTemplate(t *testing.T) {
	// Register the test with gomega
	RegisterTestingT(t)

	// Setup mock server with scripted responses
	// Setup requests buffer
	server, requests := mockServer(responsesForRenderedTemplate)
	Expect(server).NotTo(BeNil())
	Expect(requests).NotTo(BeNil())

	// Setup CloudBolt Client
	client := getClient(server)
	Expect(client).NotTo(BeNil())

	template := "Environment: {{ Environment }}"

	templateProperties := map[string]interface{}{
		"Datacenter":  "2205",
		"Environment": "prod",
		"OS":          "linux",
		"Application": "web",
		"dnsSuffix":   "example.com",
		"ProjectCode": "pro",
	}

	// Deploy the Blueprint Order
	// Expect no errors to occur
	renderedTemplate, err := client.RenderTemplate(template, templateProperties)
	Expect(err).NotTo(HaveOccurred())
	Expect(renderedTemplate).NotTo(BeNil())

	// This should have made three requests:
	// 1+2. Fail to get, get a token
	// 3. Successful Get
	Expect(len(*requests)).To(Equal(3))

	// We expect that one call to be to the order's endpoint
	Expect((*requests)[2].URL.Path).To(Equal("/api/v3/onefuse/templateTester/"))
	Expect((*requests)[2].Header["Authorization"]).To(Equal([]string{"Bearer Testing Token"}))

	// The CloudBolt Deploy Blueprint Order object should be parsed correctly
	Expect(renderedTemplate.Value).To(Equal("Environment: prod"))
}
