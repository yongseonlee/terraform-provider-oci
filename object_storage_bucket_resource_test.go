package main

import (
	"testing"
	"time"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client/mocks"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"

	"github.com/stretchr/testify/suite"
)

type ResourceObjectStorageBucketTestSuite struct {
	suite.Suite
	Client       *mocks.BareMetalClient
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	TimeCreated  baremetal.Time
	Config       string
	ResourceName string
	Res          *baremetal.Bucket
}

func (s *ResourceObjectStorageBucketTestSuite) SetupTest() {
	s.Client = &mocks.BareMetalClient{}

	s.Provider = Provider(
		func(d *schema.ResourceData) (interface{}, error) {
			return s.Client, nil
		},
	)

	s.Providers = map[string]terraform.ResourceProvider{
		"baremetal": s.Provider,
	}

	s.TimeCreated = baremetal.Time{Time: time.Now()}

	s.Config = `
		resource "baremetal_object_storage_bucket" "t" {
			compartment_id = "compartment_id"
			name = "name"
			namespace = "namespace"
			metadata = {
				"foo" = "bar"
			}
		}
	`

	s.Config += testProviderConfig

	s.ResourceName = "baremetal_object_storage_bucket.t"
	metadata := map[string]string{
		"foo": "bar",
	}
	s.Res = &baremetal.Bucket{
		CompartmentID: "compartment_id",
		Name:          "name",
		Namespace:     "namespace",
		Metadata:      metadata,
		CreatedBy:     "created_by",
		TimeCreated:   s.TimeCreated,
	}
	s.Res.ETag = "etag"
	s.Res.RequestID = "opcrequestid"

	opts := &baremetal.CreateBucketOptions{
		Metadata: metadata,
	}
	s.Client.On(
		"CreateBucket",
		"compartment_id",
		"name",
		"namespace",
		opts).Return(s.Res, nil)
	s.Client.On("DeleteBucket", "name", "namespace", (*baremetal.IfMatchOptions)(nil)).Return(nil)
}

func (s *ResourceObjectStorageBucketTestSuite) TestCreateResourceCoreBucket() {
	s.Client.On("GetBucket", "name", "namespace").Return(s.Res, nil).Times(2)
	s.Client.On("GetBucket", "name", "namespace").Return(nil, nil)

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "compartment_id", s.Res.CompartmentID),
					resource.TestCheckResourceAttr(s.ResourceName, "name", s.Res.Name),
					resource.TestCheckResourceAttr(s.ResourceName, "namespace", s.Res.Namespace),
//					resource.TestCheckResourceAttr(s.ResourceName, "created_by", s.Res.CreatedBy),
//					resource.TestCheckResourceAttr(s.ResourceName, "time_created", s.Res.TimeCreated.String()),
				),
			},
		},
	})
}

func TestResourceObjectStorageBucketTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceObjectStorageBucketTestSuite))
}
