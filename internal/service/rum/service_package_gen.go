// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package rum

import (
	"context"
	aws_sdkv1 "github.com/aws/aws-sdk-go/aws"
	session_sdkv1 "github.com/aws/aws-sdk-go/aws/session"
	cloudwatchrum_sdkv1 "github.com/aws/aws-sdk-go/service/cloudwatchrum"
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
	return []*types.ServicePackageSDKDataSource{}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceAppMonitor,
			TypeName: "aws_rum_app_monitor",
			Name:     "App Monitor",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceMetricsDestination,
			TypeName: "aws_rum_metrics_destination",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.RUM
}

func (p *servicePackage) NewConn(ctx context.Context, sess *session_sdkv1.Session, endpoint string) *cloudwatchrum_sdkv1.CloudWatchRUM {
	return cloudwatchrum_sdkv1.New(sess.Copy(&aws_sdkv1.Config{Endpoint: aws_sdkv1.String(endpoint)}))
}

var ServicePackage = &servicePackage{}
