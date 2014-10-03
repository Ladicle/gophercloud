package volumetypes

import (
	"github.com/racker/perigee"
	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/pagination"
)

// CreateOpts are options for creating a volume type.
type CreateOpts struct {
	// OPTIONAL. See VolumeType.
	ExtraSpecs map[string]interface{}
	// OPTIONAL. See VolumeType.
	Name string
}

// Create will create a new volume, optionally wih CreateOpts. To extract the
// created volume type object, call the Extract method on the CreateResult.
func Create(client *gophercloud.ServiceClient, opts *CreateOpts) CreateResult {
	type volumeType struct {
		ExtraSpecs map[string]interface{} `json:"extra_specs,omitempty"`
		Name       *string                `json:"name,omitempty"`
	}

	type request struct {
		VolumeType volumeType `json:"volume_type"`
	}

	reqBody := request{
		VolumeType: volumeType{},
	}

	reqBody.VolumeType.Name = gophercloud.MaybeString(opts.Name)
	reqBody.VolumeType.ExtraSpecs = opts.ExtraSpecs

	var res CreateResult
	_, res.Err = perigee.Request("POST", createURL(client), perigee.Options{
		MoreHeaders: client.Provider.AuthenticatedHeaders(),
		OkCodes:     []int{200, 201},
		ReqBody:     &reqBody,
		Results:     &res.Resp,
	})
	return res
}

// Delete will delete the volume type with the provided ID.
func Delete(client *gophercloud.ServiceClient, id string) error {
	_, err := perigee.Request("DELETE", deleteURL(client, id), perigee.Options{
		MoreHeaders: client.Provider.AuthenticatedHeaders(),
		OkCodes:     []int{202},
	})
	return err
}

// Get will retrieve the volume type with the provided ID. To extract the volume
// type from the result, call the Extract method on the GetResult.
func Get(client *gophercloud.ServiceClient, id string) GetResult {
	var res GetResult
	_, err := perigee.Request("GET", getURL(client, id), perigee.Options{
		MoreHeaders: client.Provider.AuthenticatedHeaders(),
		OkCodes:     []int{200},
		Results:     &res.Resp,
	})
	res.Err = err
	return res
}

// List returns all volume types.
func List(client *gophercloud.ServiceClient) pagination.Pager {
	createPage := func(r pagination.LastHTTPResponse) pagination.Page {
		return ListResult{pagination.SinglePageBase(r)}
	}

	return pagination.NewPager(client, listURL(client), createPage)
}
