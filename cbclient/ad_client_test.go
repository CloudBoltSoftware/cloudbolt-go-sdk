package cbclient

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestGetADPolicy(t *testing.T) {
	// Register the test with gomega
	RegisterTestingT(t)

	// Setup mock server with scripted responses
	// Setup requests buffer
	server, requests := mockServer(responsesForADPolicy)
	Expect(server).NotTo(BeNil())
	Expect(requests).NotTo(BeNil())

	// Setup CloudBolt Client
	client := getClient(server)
	Expect(client).NotTo(BeNil())

	policyName := "My_AD_Policy"
	policy, err := client.GetADPolicy(policyName)
	Expect(policy).NotTo(BeNil())
	Expect(err).NotTo(HaveOccurred())

	// This should have made three requests:
	// 1+2. Fail to get, get a token
	// 3. Successful Get
	Expect(len(*requests)).To(Equal(3))

	// The last request is the one we care about
	Expect((*requests)[2].URL.Path).To(Equal("/api/v3/onefuse/microsoftADPolicies/"))

	// The CloudBolt Order object should be parsed correctly
	Expect(policy.Links.Self.Href).To(Equal("/api/v3/onefuse/microsoftADPolicies/17/"))
	Expect(policy.Links.Self.Title).To(Equal("My_AD_Policy"))
	Expect(policy.Links.Workspace.Href).To(Equal("/api/v3/onefuse/workspaces/2/"))
	Expect(policy.Links.Workspace.Title).To(Equal("Default"))
	Expect(policy.Links.Endpoint.Href).To(Equal("/api/v3/onefuse/endpoints/1539/"))
	Expect(policy.Links.Endpoint.Title).To(Equal("My_AD_Policy_Endpoint"))
	Expect(policy.Name).To(Equal("My_AD_Policy"))
	Expect(policy.ID).To(Equal(17))
	Expect(policy.Description).To(Equal("microsoft active directory policy description"))
}

func TestGetMicrosoftADComputerAccountById(t *testing.T) {
	// Register the test with gomega
	RegisterTestingT(t)

	// Setup mock server with scripted responses
	// Setup requests buffer
	server, requests := mockServer(responsesForComputerAccount)
	Expect(server).NotTo(BeNil())
	Expect(requests).NotTo(BeNil())

	// Setup CloudBolt Client
	client := getClient(server)
	Expect(client).NotTo(BeNil())

	computerAccount, err := client.GetMicrosoftADComputerAccountById("23")
	Expect(computerAccount).NotTo(BeNil())
	Expect(err).NotTo(HaveOccurred())

	// This should have made three requests:
	// 1+2. Fail to get, get a token
	// 3. Successful Get
	Expect(len(*requests)).To(Equal(3))

	// The last request is the one we care about
	Expect((*requests)[2].URL.Path).To(Equal("/api/v3/onefuse/microsoftADComputerAccounts/23/"))

	// The CloudBolt Order object should be parsed correctly
	Expect(computerAccount.Links.Self.Href).To(Equal("/api/v3/onefuse/microsoftADComputerAccounts/23/"))
	Expect(computerAccount.Links.Self.Title).To(Equal("testcompacct_490991"))
	Expect(computerAccount.Links.Workspace.Href).To(Equal("/api/v3/onefuse/workspaces/1/"))
	Expect(computerAccount.Links.Workspace.Title).To(Equal("Default"))
	Expect(computerAccount.Links.Policy.Href).To(Equal("/api/v3/onefuse/microsoftADPolicies/20/"))
	Expect(computerAccount.Links.Policy.Title).To(Equal("microsoft_ad_policy_78092336"))
	Expect(computerAccount.Links.JobMetadata.Href).To(Equal("/api/v3/onefuse/jobMetadata/597/"))
	Expect(computerAccount.Links.JobMetadata.Title).To(Equal("Job Metadata Record id 597"))
	Expect(computerAccount.ID).To(Equal(23))
	Expect(computerAccount.Name).To(Equal("testcompacct_490991"))
	Expect(computerAccount.FinalOU).To(Equal("OU=qa,OU=Environments,DC=example,DC=net"))
}

func TestGetMicrosoftADComputerAccount(t *testing.T) {
	// Register the test with gomega
	RegisterTestingT(t)

	// Setup mock server with scripted responses
	// Setup requests buffer
	server, requests := mockServer(responsesForComputerAccount)
	Expect(server).NotTo(BeNil())
	Expect(requests).NotTo(BeNil())

	// Setup CloudBolt Client
	client := getClient(server)
	Expect(client).NotTo(BeNil())

	computerAccountPath := "/api/v3/onefuse/microsoftADComputerAccounts/23/"
	computerAccount, err := client.GetMicrosoftADComputerAccount(computerAccountPath)
	Expect(computerAccount).NotTo(BeNil())
	Expect(err).NotTo(HaveOccurred())

	// This should have made three requests:
	// 1+2. Fail to get, get a token
	// 3. Successful Get
	Expect(len(*requests)).To(Equal(3))

	// The last request is the one we care about
	Expect((*requests)[2].URL.Path).To(Equal("/api/v3/onefuse/microsoftADComputerAccounts/23/"))

	// The CloudBolt Order object should be parsed correctly
	Expect(computerAccount.Links.Self.Href).To(Equal("/api/v3/onefuse/microsoftADComputerAccounts/23/"))
	Expect(computerAccount.Links.Self.Title).To(Equal("testcompacct_490991"))
	Expect(computerAccount.Links.Workspace.Href).To(Equal("/api/v3/onefuse/workspaces/1/"))
	Expect(computerAccount.Links.Workspace.Title).To(Equal("Default"))
	Expect(computerAccount.Links.Policy.Href).To(Equal("/api/v3/onefuse/microsoftADPolicies/20/"))
	Expect(computerAccount.Links.Policy.Title).To(Equal("microsoft_ad_policy_78092336"))
	Expect(computerAccount.Links.JobMetadata.Href).To(Equal("/api/v3/onefuse/jobMetadata/597/"))
	Expect(computerAccount.Links.JobMetadata.Title).To(Equal("Job Metadata Record id 597"))
	Expect(computerAccount.ID).To(Equal(23))
	Expect(computerAccount.Name).To(Equal("testcompacct_490991"))
	Expect(computerAccount.FinalOU).To(Equal("OU=qa,OU=Environments,DC=example,DC=net"))
}

func TestCreateMicrosoftADComputerAccount(t *testing.T) {
	// Register the test with gomega
	RegisterTestingT(t)

	// Setup mock server with scripted responses
	// Setup requests buffer
	server, requests := mockServer(responsesForGetJobStatus)
	Expect(server).NotTo(BeNil())
	Expect(requests).NotTo(BeNil())

	// Setup CloudBolt Client
	client := getClient(server)
	Expect(client).NotTo(BeNil())

	newComputerAccount := MicrosoftADComputerAccount{
		Name:         "testcompacct_490991",
		FinalOU:      "OU=qa,OU=Environments,DC=example,DC=net",
		PolicyID:     20,
		WorkspaceURL: "/api/v3/onefuse/workspaces/1/",
		TemplateProperties: map[string]interface{}{
			"org_name": "testorg",
		},
	}

	jobStatus, err := client.CreateMicrosoftADComputerAccount(&newComputerAccount)
	Expect(jobStatus).NotTo(BeNil())
	Expect(err).NotTo(HaveOccurred())

	// This should have made three requests:
	// 1+2. Fail to get, get a token
	// 3. Successful Get
	Expect(len(*requests)).To(Equal(3))

	// The last request is the one we care about
	Expect((*requests)[2].URL.Path).To(Equal("/api/v3/onefuse/microsoftADComputerAccounts/"))

	// The CloudBolt Order object should be parsed correctly
	verifyJobStatus(jobStatus)
}

func TestDeleteMicrosoftADComputerAccount(t *testing.T) {
	// Register the test with gomega
	RegisterTestingT(t)

	// Setup mock server with scripted responses
	// Setup requests buffer
	server, requests := mockServer(responsesForGetJobStatus)
	Expect(server).NotTo(BeNil())
	Expect(requests).NotTo(BeNil())

	// Setup CloudBolt Client
	client := getClient(server)
	Expect(client).NotTo(BeNil())

	jobStatus, err := client.DeleteMicrosoftADComputerAccount("23")
	Expect(jobStatus).NotTo(BeNil())
	Expect(err).NotTo(HaveOccurred())

	// This should have made three requests:
	// 1+2. Fail to get, get a token
	// 3. Successful Get
	Expect(len(*requests)).To(Equal(3))

	// The last request is the one we care about
	Expect((*requests)[2].URL.Path).To(Equal("/api/v3/onefuse/microsoftADComputerAccounts/23/"))

	// The CloudBolt Order object should be parsed correctly
	verifyJobStatus(jobStatus)
}

func TestGetMicrosoftADPolicyById(t *testing.T) {
	// Register the test with gomega
	RegisterTestingT(t)

	// Setup mock server with scripted responses
	// Setup requests buffer
	server, requests := mockServer(responsesForMicrosoftADPolicy)
	Expect(server).NotTo(BeNil())
	Expect(requests).NotTo(BeNil())

	// Setup CloudBolt Client
	client := getClient(server)
	Expect(client).NotTo(BeNil())

	policy, err := client.GetMicrosoftADPolicyByID("11")
	Expect(policy).NotTo(BeNil())
	Expect(err).NotTo(HaveOccurred())

	// This should have made three requests:
	// 1+2. Fail to get, get a token
	// 3. Successful Get
	Expect(len(*requests)).To(Equal(3))

	// The last request is the one we care about
	Expect((*requests)[2].URL.Path).To(Equal("/api/v3/onefuse/microsoftADPolicies/11/"))

	// The CloudBolt Order object should be parsed correctly
	verifyMicrosoftADPolicy(policy)
}

func TestCreateMicrosoftADPolicyById(t *testing.T) {
	// Register the test with gomega
	RegisterTestingT(t)

	// Setup mock server with scripted responses
	// Setup requests buffer
	server, requests := mockServer(responsesForMicrosoftADPolicy)
	Expect(server).NotTo(BeNil())
	Expect(requests).NotTo(BeNil())

	// Setup CloudBolt Client
	client := getClient(server)
	Expect(client).NotTo(BeNil())

	newPolicy := MicrosoftADPolicy{
		Name:                   "microsoft_ad_policy_38181060",
		Description:            "microsoft active directory policy description",
		OU:                     "microsoft active directory policy description",
		MicrosoftEndpointID:    506,
		ComputerNameLetterCase: "LOWER",
		WorkspaceURL:           "/api/v3/onefuse/workspaces/1/",
		CreateOU:               false,
		RemoveOU:               false,
		SecurityGroups: []string{
			"CN=TestSecurityGroup,OU=OneFuse,DC=example,DC=net",
			"CN=TestSecurityGroupTemp,OU=TEMPOU1,DC=example,DC=net",
		},
	}

	policy, err := client.CreateMicrosoftADPolicy(&newPolicy)
	Expect(policy).NotTo(BeNil())
	Expect(err).NotTo(HaveOccurred())

	// This should have made three requests:
	// 1+2. Fail to get, get a token
	// 3. Successful Get
	Expect(len(*requests)).To(Equal(3))

	// The last request is the one we care about
	Expect((*requests)[2].URL.Path).To(Equal("/api/v3/onefuse/microsoftADPolicies/"))

	// The CloudBolt Order object should be parsed correctly
	verifyMicrosoftADPolicy(policy)
}

func TestUpdateMicrosoftADPolicyById(t *testing.T) {
	// Register the test with gomega
	RegisterTestingT(t)

	// Setup mock server with scripted responses
	// Setup requests buffer
	server, requests := mockServer(responsesForMicrosoftADPolicy)
	Expect(server).NotTo(BeNil())
	Expect(requests).NotTo(BeNil())

	// Setup CloudBolt Client
	client := getClient(server)
	Expect(client).NotTo(BeNil())

	updatePolicy := MicrosoftADPolicy{
		Name:                   "microsoft_ad_policy_38181060",
		Description:            "microsoft active directory policy description",
		OU:                     "microsoft active directory policy description",
		MicrosoftEndpointID:    506,
		ComputerNameLetterCase: "LOWER",
		WorkspaceURL:           "/api/v3/onefuse/workspaces/1/",
		CreateOU:               false,
		RemoveOU:               false,
		SecurityGroups: []string{
			"CN=TestSecurityGroup,OU=OneFuse,DC=example,DC=net",
			"CN=TestSecurityGroupTemp,OU=TEMPOU1,DC=example,DC=net",
		},
	}

	policy, err := client.UpdateMicrosoftADPolicy("11", &updatePolicy)
	Expect(policy).NotTo(BeNil())
	Expect(err).NotTo(HaveOccurred())

	// This should have made three requests:
	// 1+2. Fail to get, get a token
	// 3. Successful Get
	Expect(len(*requests)).To(Equal(3))

	// The last request is the one we care about
	Expect((*requests)[2].URL.Path).To(Equal("/api/v3/onefuse/microsoftADPolicies/11/"))

	// The CloudBolt Order object should be parsed correctly
	verifyMicrosoftADPolicy(policy)
}

func TestDeleteMicrosoftADPolicyById(t *testing.T) {
	// Register the test with gomega
	RegisterTestingT(t)

	// Setup mock server with scripted responses
	// Setup requests buffer
	server, requests := mockServer(responsesForMicrosoftADPolicy)
	Expect(server).NotTo(BeNil())
	Expect(requests).NotTo(BeNil())

	// Setup CloudBolt Client
	client := getClient(server)
	Expect(client).NotTo(BeNil())

	err := client.DeleteMicrosoftADPolicy("11")
	Expect(err).NotTo(HaveOccurred())

	// This should have made three requests:
	// 1+2. Fail to get, get a token
	// 3. Successful Get
	Expect(len(*requests)).To(Equal(3))

	// The last request is the one we care about
	Expect((*requests)[2].URL.Path).To(Equal("/api/v3/onefuse/microsoftADPolicies/11/"))
}

func verifyMicrosoftADPolicy(policy *MicrosoftADPolicy) {
	Expect(policy.Links.Self.Href).To(Equal("/api/v3/onefuse/microsoftADPolicies/11/"))
	Expect(policy.Links.Self.Title).To(Equal("microsoft_ad_policy_38181060"))
	Expect(policy.Links.Workspace.Href).To(Equal("/api/v3/onefuse/workspaces/1/"))
	Expect(policy.Links.Workspace.Title).To(Equal("Default"))
	Expect(policy.Links.MicrosoftEndpoint.Href).To(Equal("/api/v3/onefuse/endpoints/506/"))
	Expect(policy.Links.MicrosoftEndpoint.Title).To(Equal("QA_Microsoft_Endpoint_91672960"))
	Expect(policy.ID).To(Equal(11))
	Expect(policy.Name).To(Equal("microsoft_ad_policy_38181060"))
	Expect(policy.Description).To(Equal("microsoft active directory policy description"))
	Expect(policy.MicrosoftEndpoint).To(Equal("Endpoint: microsoft: QA_Microsoft_Endpoint_91672960: dc01.example.net:443"))
	Expect(policy.ComputerNameLetterCase).To(Equal("LOWER"))
	Expect(policy.CreateOU).To(Equal(false))
	Expect(policy.RemoveOU).To(Equal(false))
	Expect(policy.OU).To(Equal("OU=qa,OU=Environments,DC=example,DC=net"))
	Expect(len(policy.SecurityGroups)).To(Equal(2))
	Expect(policy.SecurityGroups[0]).To(Equal("CN=TestSecurityGroup,OU=OneFuse,DC=example,DC=net"))
	Expect(policy.SecurityGroups[1]).To(Equal("CN=TestSecurityGroupTemp,OU=TEMPOU1,DC=example,DC=net"))
}
