// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scf

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
	case "tencentcloud:Scf/function:Function":
		r = &Function{}
	case "tencentcloud:Scf/functionAlias:FunctionAlias":
		r = &FunctionAlias{}
	case "tencentcloud:Scf/functionEventInvokeConfig:FunctionEventInvokeConfig":
		r = &FunctionEventInvokeConfig{}
	case "tencentcloud:Scf/functionVersion:FunctionVersion":
		r = &FunctionVersion{}
	case "tencentcloud:Scf/invokeFunction:InvokeFunction":
		r = &InvokeFunction{}
	case "tencentcloud:Scf/layer:Layer":
		r = &Layer{}
	case "tencentcloud:Scf/namespace:Namespace":
		r = &Namespace{}
	case "tencentcloud:Scf/provisionedConcurrencyConfig:ProvisionedConcurrencyConfig":
		r = &ProvisionedConcurrencyConfig{}
	case "tencentcloud:Scf/reservedConcurrencyConfig:ReservedConcurrencyConfig":
		r = &ReservedConcurrencyConfig{}
	case "tencentcloud:Scf/syncInvokeFunction:SyncInvokeFunction":
		r = &SyncInvokeFunction{}
	case "tencentcloud:Scf/terminateAsyncEvent:TerminateAsyncEvent":
		r = &TerminateAsyncEvent{}
	case "tencentcloud:Scf/triggerConfig:TriggerConfig":
		r = &TriggerConfig{}
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
		"Scf/function",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Scf/functionAlias",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Scf/functionEventInvokeConfig",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Scf/functionVersion",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Scf/invokeFunction",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Scf/layer",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Scf/namespace",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Scf/provisionedConcurrencyConfig",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Scf/reservedConcurrencyConfig",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Scf/syncInvokeFunction",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Scf/terminateAsyncEvent",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Scf/triggerConfig",
		&module{version},
	)
}
