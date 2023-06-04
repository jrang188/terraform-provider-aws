// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package dynamodb

import (
	"context"
	aws_sdkv1 "github.com/aws/aws-sdk-go/aws"
	session_sdkv1 "github.com/aws/aws-sdk-go/aws/session"
	dynamodb_sdkv1 "github.com/aws/aws-sdk-go/service/dynamodb"
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
			Factory:  DataSourceTable,
			TypeName: "aws_dynamodb_table",
		},
		{
			Factory:  DataSourceTableItem,
			TypeName: "aws_dynamodb_table_item",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceContributorInsights,
			TypeName: "aws_dynamodb_contributor_insights",
		},
		{
			Factory:  ResourceGlobalTable,
			TypeName: "aws_dynamodb_global_table",
		},
		{
			Factory:  ResourceKinesisStreamingDestination,
			TypeName: "aws_dynamodb_kinesis_streaming_destination",
		},
		{
			Factory:  ResourceTable,
			TypeName: "aws_dynamodb_table",
			Name:     "Table",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceTableItem,
			TypeName: "aws_dynamodb_table_item",
		},
		{
			Factory:  ResourceTableReplica,
			TypeName: "aws_dynamodb_table_replica",
			Name:     "Table Replica",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  ResourceTag,
			TypeName: "aws_dynamodb_tag",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.DynamoDB
}

func (p *servicePackage) NewConn(ctx context.Context, sess *session_sdkv1.Session, endpoint string) *dynamodb_sdkv1.DynamoDB {
	return dynamodb_sdkv1.New(sess.Copy(&aws_sdkv1.Config{Endpoint: aws_sdkv1.String(endpoint)}))
}

var ServicePackage = &servicePackage{}
