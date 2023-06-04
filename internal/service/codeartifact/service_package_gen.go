// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package codeartifact

import (
	"context"
	aws_sdkv1 "github.com/aws/aws-sdk-go/aws"
	session_sdkv1 "github.com/aws/aws-sdk-go/aws/session"
	codeartifact_sdkv1 "github.com/aws/aws-sdk-go/service/codeartifact"
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
			Factory:  DataSourceAuthorizationToken,
			TypeName: "aws_codeartifact_authorization_token",
		},
		{
			Factory:  DataSourceRepositoryEndpoint,
			TypeName: "aws_codeartifact_repository_endpoint",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceDomain,
			TypeName: "aws_codeartifact_domain",
			Name:     "Domain",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceDomainPermissionsPolicy,
			TypeName: "aws_codeartifact_domain_permissions_policy",
		},
		{
			Factory:  ResourceRepository,
			TypeName: "aws_codeartifact_repository",
			Name:     "Repository",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceRepositoryPermissionsPolicy,
			TypeName: "aws_codeartifact_repository_permissions_policy",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.CodeArtifact
}

func (p *servicePackage) NewConn(ctx context.Context, sess *session_sdkv1.Session, endpoint string) *codeartifact_sdkv1.CodeArtifact {
	return codeartifact_sdkv1.New(sess.Copy(&aws_sdkv1.Config{Endpoint: aws_sdkv1.String(endpoint)}))
}

var ServicePackage = &servicePackage{}
