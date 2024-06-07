// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "tencentcloud:ApiGateway/api:Api":
		r = &Api{}
	case "tencentcloud:ApiGateway/apiApp:ApiApp":
		r = &ApiApp{}
	case "tencentcloud:ApiGateway/apiAppAttachment:ApiAppAttachment":
		r = &ApiAppAttachment{}
	case "tencentcloud:ApiGateway/apiDoc:ApiDoc":
		r = &ApiDoc{}
	case "tencentcloud:ApiGateway/apiKey:ApiKey":
		r = &ApiKey{}
	case "tencentcloud:ApiGateway/apiKeyAttachment:ApiKeyAttachment":
		r = &ApiKeyAttachment{}
	case "tencentcloud:ApiGateway/customDomain:CustomDomain":
		r = &CustomDomain{}
	case "tencentcloud:ApiGateway/importOpenApi:ImportOpenApi":
		r = &ImportOpenApi{}
	case "tencentcloud:ApiGateway/ipStrategy:IpStrategy":
		r = &IpStrategy{}
	case "tencentcloud:ApiGateway/plugin:Plugin":
		r = &Plugin{}
	case "tencentcloud:ApiGateway/pluginAttachment:PluginAttachment":
		r = &PluginAttachment{}
	case "tencentcloud:ApiGateway/service:Service":
		r = &Service{}
	case "tencentcloud:ApiGateway/serviceRelease:ServiceRelease":
		r = &ServiceRelease{}
	case "tencentcloud:ApiGateway/strategyAttachment:StrategyAttachment":
		r = &StrategyAttachment{}
	case "tencentcloud:ApiGateway/updateApiAppKey:UpdateApiAppKey":
		r = &UpdateApiAppKey{}
	case "tencentcloud:ApiGateway/updateService:UpdateService":
		r = &UpdateService{}
	case "tencentcloud:ApiGateway/upstream:Upstream":
		r = &Upstream{}
	case "tencentcloud:ApiGateway/usagePlan:UsagePlan":
		r = &UsagePlan{}
	case "tencentcloud:ApiGateway/usagePlanAttachment:UsagePlanAttachment":
		r = &UsagePlanAttachment{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := internal.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"ApiGateway/api",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"ApiGateway/apiApp",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"ApiGateway/apiAppAttachment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"ApiGateway/apiDoc",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"ApiGateway/apiKey",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"ApiGateway/apiKeyAttachment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"ApiGateway/customDomain",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"ApiGateway/importOpenApi",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"ApiGateway/ipStrategy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"ApiGateway/plugin",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"ApiGateway/pluginAttachment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"ApiGateway/service",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"ApiGateway/serviceRelease",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"ApiGateway/strategyAttachment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"ApiGateway/updateApiAppKey",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"ApiGateway/updateService",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"ApiGateway/upstream",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"ApiGateway/usagePlan",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"ApiGateway/usagePlanAttachment",
		&module{version},
	)
}
