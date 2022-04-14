package cbclient

import (
	"encoding/json"
	"testing"

	. "github.com/onsi/gomega"
)

func TestGetBlueprint(t *testing.T) {
	// Register the test with gomega
	RegisterTestingT(t)

	// Setup mock server with scripted responses
	// Setup requests buffer
	server, requests := mockServer(responsesForBlueprint)
	Expect(server).NotTo(BeNil())
	Expect(requests).NotTo(BeNil())

	// Setup CloudBolt Client
	client := getClient(server)
	Expect(client).NotTo(BeNil())

	// Define an orderID parameter value
	blueprintName := "My Simple Blueprint"

	// Get the CloudBolt Order object
	// Expect no errors to occur
	blueprint, err := client.GetBlueprint(blueprintName)
	Expect(blueprint).NotTo(BeNil())
	Expect(err).NotTo(HaveOccurred())

	// This should have made three requests:
	// 1+2. Fail to get order, get a token
	// 3. Successfully getting the order
	Expect(len(*requests)).To(Equal(3))

	// The last request is the one we care about
	Expect((*requests)[2].URL.Path).To(Equal("/api/v3/cmp/blueprints/"))

	// The CloudBolt Order object should be parsed correctly
	Expect(blueprint.Links.Self.Href).To(Equal("/api/v3/cmp/blueprints/BP-esnjtp7u/"))
	Expect(blueprint.Links.Self.Title).To(Equal("My Simple Blueprint"))
	Expect(blueprint.Name).To(Equal("My Simple Blueprint"))
	Expect(blueprint.ID).To(Equal("BP-esnjtp7u"))
}

func TestGetBlueprintById(t *testing.T) {
	// Register the test with gomega
	RegisterTestingT(t)

	// Setup mock server with scripted responses
	// Setup requests buffer
	server, requests := mockServer(responsesForBlueprintById)
	Expect(server).NotTo(BeNil())
	Expect(requests).NotTo(BeNil())

	// Setup CloudBolt Client
	client := getClient(server)
	Expect(client).NotTo(BeNil())

	// Define an orderID parameter value
	blueprintId := "BP-esnjtp7u"

	// Get the CloudBolt Order object
	// Expect no errors to occur
	blueprint, err := client.GetBlueprintById(blueprintId)
	Expect(blueprint).NotTo(BeNil())
	Expect(err).NotTo(HaveOccurred())

	// This should have made three requests:
	// 1+2. Fail to get order, get a token
	// 3. Successfully getting the order
	Expect(len(*requests)).To(Equal(3))

	// The last request is the one we care about
	Expect((*requests)[2].URL.Path).To(Equal("/api/v3/cmp/blueprints/BP-esnjtp7u/"))

	// The CloudBolt Order object should be parsed correctly
	Expect(blueprint.Links.Self.Href).To(Equal("/api/v3/cmp/blueprints/BP-esnjtp7u/"))
	Expect(blueprint.Links.Self.Title).To(Equal("My Simple Blueprint"))
	Expect(blueprint.Name).To(Equal("My Simple Blueprint"))
	Expect(blueprint.ID).To(Equal("BP-esnjtp7u"))
}

func TestDeployBlueprint(t *testing.T) {
	// Register the test with gomega
	RegisterTestingT(t)

	// Setup mock server with scripted responses
	// Setup requests buffer
	server, requests := mockServer(responsesForDeployBlueprint)
	Expect(server).NotTo(BeNil())
	Expect(requests).NotTo(BeNil())

	// Setup CloudBolt Client
	client := getClient(server)
	Expect(client).NotTo(BeNil())

	// Get the order items we are going to deploy
	bpItems := make([]map[string]interface{}, 0)
	bpItems = append(bpItems, map[string]interface{}{
		"bp-item-name": "plugin-bdi-olk0xwve",
		"bp-item-paramas": map[string]interface{}{
			"actiom_param1": "act1 value",
			"actiom_param2": "act2 value",
		},
	})

	bpItems = append(bpItems, map[string]interface{}{
		"bp-item-name": "server-bdi-743tlxxu",
		"bp-item-paramas": map[string]interface{}{
			"instance_type": "t2.nano",
			"env_param2":    20,
			"grp_param2":    30,
			"env_param1":    "env1 value",
			"grp_param1":    "grp1 value",
		},
		"environment": "/api/v3/cmp/environments/ENV-1tytr2pu/",
		"osbuild":     "/api/v3/cmp/osBuilds/OSB-z69hjvki/",
	})

	bpParams := map[string]interface{}{
		"bp_param1": "my parameter 1",
		"bp_param2": 2,
		"bp_param3": true,
		"bp_param4": 3.14,
	}

	grpPath := "/api/v3/cmp/groups/GRP-yfbbsfht/"
	bpPath := "/api/v3/cmp/blueprints/BP-esnjtp7u/"

	// Deploy the Blueprint Order
	// Expect no errors to occur
	order, err := client.DeployBlueprint(grpPath, bpPath, "resource name", bpParams, bpItems)
	Expect(err).NotTo(HaveOccurred())
	Expect(order).NotTo(BeNil())

	// This should have made three requests:
	// 1+2. Fail to get resource, get a token
	// 3. Successfully getting the object
	Expect(len(*requests)).To(Equal(3))

	// We expect that one call to be to the order's endpoint
	Expect((*requests)[2].URL.Path).To(Equal("/api/v3/cmp/blueprints/api/v3/cmp/blueprints/BP-esnjtp7u/deploy/"))
	Expect((*requests)[2].Header["Authorization"]).To(Equal([]string{"Bearer Testing Token"}))

	// Verify Blueprint Deployment Payload Is generated correctly.
	var bpReqJson struct {
		Group      string `json:"group"`
		Parameters struct {
			BpParam1 string  `json:"bp_param1"`
			BpParam2 int     `json:"bp_param2"`
			BpParam3 bool    `json:"bp_param3"`
			BpParam4 float64 `json:"bp_param4"`
		} `json:"parameters"`
		DeploymentItems struct {
			PluginBdiOlk0Xwve struct {
				Parameters struct {
					ActiomParam1 string `json:"actiom_param1"`
					ActiomParam2 string `json:"actiom_param2"`
				} `json:"parameters"`
			} `json:"plugin-bdi-olk0xwve"`
			ServerBdi743Tlxxu struct {
				Environment string `json:"environment"`
				OsBuild     string `json:"osBuild"`
				Parameters  struct {
					InstanceType string `json:"instance_type"`
					EnvParam2    int    `json:"env_param2"`
					GrpParam2    int    `json:"grp_param2"`
					EnvParam1    string `json:"env_param1"`
					GrpParam1    string `json:"grp_param1"`
				} `json:"parameters"`
			} `json:"server-bdi-743tlxxu"`
		} `json:"deploymentItems"`
	}

	json.NewDecoder((*requests)[2].Body).Decode(&bpReqJson)
	Expect(bpReqJson.Group).To(Equal("/api/v3/cmp/groups/GRP-yfbbsfht/"))
	Expect(bpReqJson.Parameters.BpParam1).To(Equal("my parameter 1"))
	Expect(bpReqJson.Parameters.BpParam2).To(Equal(2))
	Expect(bpReqJson.Parameters.BpParam3).To(Equal(true))
	Expect(bpReqJson.Parameters.BpParam4).To(Equal(3.14))
	Expect(bpReqJson.DeploymentItems.PluginBdiOlk0Xwve.Parameters.ActiomParam1).To(Equal("act1 value"))
	Expect(bpReqJson.DeploymentItems.PluginBdiOlk0Xwve.Parameters.ActiomParam2).To(Equal("act2 value"))
	Expect(bpReqJson.DeploymentItems.ServerBdi743Tlxxu.Environment).To(Equal("/api/v3/cmp/environments/ENV-1tytr2pu/"))
	Expect(bpReqJson.DeploymentItems.ServerBdi743Tlxxu.OsBuild).To(Equal("/api/v3/cmp/osBuilds/OSB-z69hjvki/"))
	Expect(bpReqJson.DeploymentItems.ServerBdi743Tlxxu.Parameters.EnvParam1).To(Equal("env1 value"))
	Expect(bpReqJson.DeploymentItems.ServerBdi743Tlxxu.Parameters.EnvParam2).To(Equal(20))
	Expect(bpReqJson.DeploymentItems.ServerBdi743Tlxxu.Parameters.GrpParam1).To(Equal("grp1 value"))
	Expect(bpReqJson.DeploymentItems.ServerBdi743Tlxxu.Parameters.GrpParam2).To(Equal(30))
	Expect(bpReqJson.DeploymentItems.ServerBdi743Tlxxu.Parameters.InstanceType).To(Equal("t2.nano"))

	// The CloudBolt Deploy Blueprint Order object should be parsed correctly
	Expect(order.Links.Self.Href).To(Equal("/api/v3/cmp/orders/ORD-e9v87uia/"))
	Expect(order.Links.Self.Title).To(Equal("Installation of My Simple Blueprint"))
	Expect(order.Links.Group.Href).To(Equal("/api/v3/cloudbolt/groups/GRP-yfbbsfht/"))
	Expect(order.Links.Group.Title).To(Equal("My Org"))
	Expect(order.Links.Owner.Href).To(Equal("/api/v3/cloudbolt/users/USR-mxpqe1x7/"))
	Expect(order.Links.Owner.Title).To(Equal("user001"))
	Expect(order.Links.ApprovedBy.Href).To(Equal("/api/v3/cloudbolt/users/USR-mxpqe1x7/"))
	Expect(order.Links.ApprovedBy.Title).To(Equal("user001"))
	Expect(len(order.Links.Jobs)).To(Equal(2))
	Expect(order.Links.Jobs[0].Href).To(Equal("/api/v3/cmp/jobs/JOB-9nrax3gb/"))
	Expect(order.Links.Jobs[0].Title).To(Equal("Deploy Blueprint Job 1011"))
	Expect(order.Links.Duplicate.Href).To(Equal("/api/v3/cmp/orders/ORD-e9v87uia/duplicate/"))
	Expect(order.Links.Duplicate.Title).To(Equal("Duplicate Order"))
	Expect(order.Name).To(Equal("Installation of My Simple Blueprint"))
	Expect(order.ID).To(Equal("ORD-e9v87uia"))
	Expect(order.Status).To(Equal("SUCCESS"))
	Expect(order.Rate).To(Equal("4.18/month"))
	Expect(len(order.DeploymentItems)).To(Equal(1))
	Expect(order.DeploymentItems[0].ID).To(Equal("OI-1p0bajs6"))
	Expect(order.DeploymentItems[0].ResourceName).To(Equal("My Simple Blueprint"))
	Expect(order.DeploymentItems[0].ResourceParameters).To(Not(BeNil()))
	Expect(order.DeploymentItems[0].Blueprint.Href).To(Equal("/api/v3/cmp/blueprints/BP-esnjtp7u/"))
	Expect(order.DeploymentItems[0].Blueprint.Title).To(Equal("My Simple Blueprint"))
	Expect(order.DeploymentItems[0].BlueprintItemsArguments).To(Not(BeNil()))
	Expect(order.DeploymentItems[0].ItemType).To(Equal("blueprint"))
}
