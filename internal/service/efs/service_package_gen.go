// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package efs

import (
	"context"
	aws_sdkv1 "github.com/aws/aws-sdk-go/aws"
	session_sdkv1 "github.com/aws/aws-sdk-go/aws/session"
	efs_sdkv1 "github.com/aws/aws-sdk-go/service/efs"
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
			Factory:  DataSourceAccessPoint,
			TypeName: "aws_efs_access_point",
		},
		{
			Factory:  DataSourceAccessPoints,
			TypeName: "aws_efs_access_points",
		},
		{
			Factory:  DataSourceFileSystem,
			TypeName: "aws_efs_file_system",
		},
		{
			Factory:  DataSourceMountTarget,
			TypeName: "aws_efs_mount_target",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceAccessPoint,
			TypeName: "aws_efs_access_point",
			Name:     "Access Point",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "id",
			},
		},
		{
			Factory:  ResourceBackupPolicy,
			TypeName: "aws_efs_backup_policy",
		},
		{
			Factory:  ResourceFileSystem,
			TypeName: "aws_efs_file_system",
			Name:     "File System",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "id",
			},
		},
		{
			Factory:  ResourceFileSystemPolicy,
			TypeName: "aws_efs_file_system_policy",
		},
		{
			Factory:  ResourceMountTarget,
			TypeName: "aws_efs_mount_target",
		},
		{
			Factory:  ResourceReplicationConfiguration,
			TypeName: "aws_efs_replication_configuration",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.EFS
}

func (p *servicePackage) NewConn(ctx context.Context, sess *session_sdkv1.Session, endpoint string) *efs_sdkv1.EFS {
	return efs_sdkv1.New(sess.Copy(&aws_sdkv1.Config{Endpoint: aws_sdkv1.String(endpoint)}))
}

var ServicePackage = &servicePackage{}
