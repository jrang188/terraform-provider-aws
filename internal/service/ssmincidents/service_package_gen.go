// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package ssmincidents

import (
	"context"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  DataSourceReplicationSet,
			TypeName: "aws_ssmincidents_replication_set",
		},
		{
			Factory:  DataSourceResponsePlan,
			TypeName: "aws_ssmincidents_response_plan",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceReplicationSet,
			TypeName: "aws_ssmincidents_replication_set",
			Name:     "Replication Set",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "id",
			},
		},
		{
			Factory:  ResourceResponsePlan,
			TypeName: "aws_ssmincidents_response_plan",
			Name:     "Response Plan",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "id",
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.SSMIncidents
}

var ServicePackage = &servicePackage{}
