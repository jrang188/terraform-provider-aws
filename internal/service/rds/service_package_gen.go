// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package rds

import (
	"context"
	aws_sdkv1 "github.com/aws/aws-sdk-go/aws"
	session_sdkv1 "github.com/aws/aws-sdk-go/aws/session"
	rds_sdkv1 "github.com/aws/aws-sdk-go/service/rds"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{
		{
			Factory: newResourceExportTask,
		},
	}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  DataSourceClusterSnapshot,
			TypeName: "aws_db_cluster_snapshot",
		},
		{
			Factory:  DataSourceEventCategories,
			TypeName: "aws_db_event_categories",
		},
		{
			Factory:  DataSourceInstance,
			TypeName: "aws_db_instance",
		},
		{
			Factory:  DataSourceInstances,
			TypeName: "aws_db_instances",
		},
		{
			Factory:  DataSourceProxy,
			TypeName: "aws_db_proxy",
		},
		{
			Factory:  DataSourceSnapshot,
			TypeName: "aws_db_snapshot",
		},
		{
			Factory:  DataSourceSubnetGroup,
			TypeName: "aws_db_subnet_group",
		},
		{
			Factory:  DataSourceCertificate,
			TypeName: "aws_rds_certificate",
		},
		{
			Factory:  DataSourceCluster,
			TypeName: "aws_rds_cluster",
		},
		{
			Factory:  DataSourceClusters,
			TypeName: "aws_rds_clusters",
		},
		{
			Factory:  DataSourceEngineVersion,
			TypeName: "aws_rds_engine_version",
		},
		{
			Factory:  DataSourceOrderableInstance,
			TypeName: "aws_rds_orderable_db_instance",
		},
		{
			Factory:  DataSourceReservedOffering,
			TypeName: "aws_rds_reserved_instance_offering",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceClusterSnapshot,
			TypeName: "aws_db_cluster_snapshot",
			Name:     "Cluster Snapshot",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "db_cluster_snapshot_arn",
			},
		},
		{
			Factory:  ResourceEventSubscription,
			TypeName: "aws_db_event_subscription",
			Name:     "Event Subscription",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceInstance,
			TypeName: "aws_db_instance",
			Name:     "DB Instance",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceInstanceAutomatedBackupsReplication,
			TypeName: "aws_db_instance_automated_backups_replication",
		},
		{
			Factory:  ResourceInstanceRoleAssociation,
			TypeName: "aws_db_instance_role_association",
		},
		{
			Factory:  ResourceOptionGroup,
			TypeName: "aws_db_option_group",
			Name:     "DB Option Group",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceParameterGroup,
			TypeName: "aws_db_parameter_group",
			Name:     "DB Parameter Group",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceProxy,
			TypeName: "aws_db_proxy",
			Name:     "DB Proxy",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceProxyDefaultTargetGroup,
			TypeName: "aws_db_proxy_default_target_group",
		},
		{
			Factory:  ResourceProxyEndpoint,
			TypeName: "aws_db_proxy_endpoint",
			Name:     "DB Proxy Endpoint",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceProxyTarget,
			TypeName: "aws_db_proxy_target",
		},
		{
			Factory:  ResourceSnapshot,
			TypeName: "aws_db_snapshot",
			Name:     "DB Snapshot",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "db_snapshot_arn",
			},
		},
		{
			Factory:  ResourceSnapshotCopy,
			TypeName: "aws_db_snapshot_copy",
			Name:     "DB Snapshot",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "db_snapshot_arn",
			},
		},
		{
			Factory:  ResourceSubnetGroup,
			TypeName: "aws_db_subnet_group",
			Name:     "DB Subnet Group",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceCluster,
			TypeName: "aws_rds_cluster",
			Name:     "Cluster",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceClusterActivityStream,
			TypeName: "aws_rds_cluster_activity_stream",
		},
		{
			Factory:  ResourceClusterEndpoint,
			TypeName: "aws_rds_cluster_endpoint",
			Name:     "Cluster Endpoint",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceClusterInstance,
			TypeName: "aws_rds_cluster_instance",
			Name:     "Cluster Instance",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceClusterParameterGroup,
			TypeName: "aws_rds_cluster_parameter_group",
			Name:     "Cluster Parameter Group",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceClusterRoleAssociation,
			TypeName: "aws_rds_cluster_role_association",
		},
		{
			Factory:  ResourceGlobalCluster,
			TypeName: "aws_rds_global_cluster",
		},
		{
			Factory:  ResourceReservedInstance,
			TypeName: "aws_rds_reserved_instance",
			Name:     "Reserved Instance",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.RDS
}

func (p *servicePackage) NewConn(ctx context.Context, sess *session_sdkv1.Session, endpoint string) *rds_sdkv1.RDS {
	return rds_sdkv1.New(sess.Copy(&aws_sdkv1.Config{Endpoint: aws_sdkv1.String(endpoint)}))
}

var ServicePackage = &servicePackage{}
