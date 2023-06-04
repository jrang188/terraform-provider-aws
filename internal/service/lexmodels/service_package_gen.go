// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package lexmodels

import (
	"context"
	aws_sdkv1 "github.com/aws/aws-sdk-go/aws"
	session_sdkv1 "github.com/aws/aws-sdk-go/aws/session"
	lexmodelbuildingservice_sdkv1 "github.com/aws/aws-sdk-go/service/lexmodelbuildingservice"
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
			Factory:  DataSourceBot,
			TypeName: "aws_lex_bot",
		},
		{
			Factory:  DataSourceBotAlias,
			TypeName: "aws_lex_bot_alias",
		},
		{
			Factory:  DataSourceIntent,
			TypeName: "aws_lex_intent",
		},
		{
			Factory:  DataSourceSlotType,
			TypeName: "aws_lex_slot_type",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceBot,
			TypeName: "aws_lex_bot",
		},
		{
			Factory:  ResourceBotAlias,
			TypeName: "aws_lex_bot_alias",
		},
		{
			Factory:  ResourceIntent,
			TypeName: "aws_lex_intent",
		},
		{
			Factory:  ResourceSlotType,
			TypeName: "aws_lex_slot_type",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.LexModels
}

func (p *servicePackage) NewConn(ctx context.Context, sess *session_sdkv1.Session, endpoint string) *lexmodelbuildingservice_sdkv1.LexModelBuildingService {
	return lexmodelbuildingservice_sdkv1.New(sess.Copy(&aws_sdkv1.Config{Endpoint: aws_sdkv1.String(endpoint)}))
}

var ServicePackage = &servicePackage{}
