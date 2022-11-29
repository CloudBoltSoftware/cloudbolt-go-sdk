package cbclient

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/url"
	"strings"
)

type CloudBoltGroup struct {
	CloudBoltReferenceFields
	Parent CloudBoltHALItem `json:"parent"`
}

type CloudBoltGroupResult struct {
	CloudBoltResult
	Embedded struct {
		Groups []CloudBoltGroup `json:"groups"`
	} `json:"_embedded"`
}

// GetGroup accepts a groupPath string parameter of the following format:
// "/my parent group/some subgroup/a child group/" or just "my parent group"
//
// verifyGroup recursively verifies that this is a valid group/subgroup.
func (c *CloudBoltClient) GetGroup(groupPath string) (*CloudBoltGroup, error) {
	var group string
	var parentPath string
	var groupFound bool

	groupPath = strings.Trim(groupPath, "/")
	nextIndex := strings.LastIndex(groupPath, "/")

	if nextIndex >= 0 {
		group = groupPath[nextIndex+1:]
		parentPath = groupPath[:nextIndex]
	} else {
		group = groupPath
	}

	apiurl := c.baseURL
	apiurl.Path = c.apiEndpoint("cmp", "groups")
	apiurl.RawQuery = fmt.Sprintf(filterByName, url.QueryEscape(group))

	resp, err := c.makeRequest("GET", apiurl.String(), nil)
	if err != nil {
		return nil, err
	}

	// We Decode the data because we already have an io.Reader on hand
	var res CloudBoltGroupResult
	json.NewDecoder(resp.Body).Decode(&res)

	for _, v := range res.Embedded.Groups {
		groupFound, err = c.verifyGroup(v.Links.Self.Href, parentPath)

		if groupFound {
			return &v, nil
		}
	}

	return nil, fmt.Errorf("Group (%s): Not Found", groupPath)
}

func (c *CloudBoltClient) GetGroupById(id string) (*CloudBoltGroup, error) {
	apiurl := c.baseURL
	apiurl.Path = c.apiEndpoint("cmp", "groups", id)

	resp, err := c.makeRequest("GET", apiurl.String(), nil)
	if err != nil {
		return nil, err
	}

	var group CloudBoltGroup
	json.NewDecoder(resp.Body).Decode(&group)

	return &group, nil
}

// verifyGroup checks that all a given group is the one we intended to fetch.
//
// groupPath is the API path to the group, e.g., "/api/v2/groups/GRP-123456"
//
// If a group has no parents, "parentPath" should be empty.
// If a group has parents, it should be of the format "root-level-parent/sub-parent/.../closest-parent"
func (c *CloudBoltClient) verifyGroup(groupPath string, parentPath string) (bool, error) {
	var parent string
	var nextParentPath string

	apiurl := c.baseURL
	apiurl.Path = groupPath

	resp, err := c.makeRequest("GET", apiurl.String(), nil)
	if err != nil {
		return false, err
	}
	if resp.StatusCode >= 300 {
		log.Fatalln(resp.Status)

		return false, errors.New(resp.Status)
	}

	// We Decode the data because we already have an io.Reader on hand
	var group CloudBoltGroup
	json.NewDecoder(resp.Body).Decode(&group)

	nextIndex := strings.LastIndex(parentPath, "/")

	if nextIndex >= 0 {
		parent = parentPath[nextIndex+1:]
		nextParentPath = parentPath[:nextIndex]
	} else {
		parent = parentPath
	}

	if group.Parent.Title != parent {
		return false, nil
	}

	if nextParentPath != "" {
		return c.verifyGroup(group.Parent.Href, nextParentPath)
	}

	return true, nil
}
