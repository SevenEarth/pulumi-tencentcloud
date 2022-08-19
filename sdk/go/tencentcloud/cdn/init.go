// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cdn

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "tencentcloud:Cdn/domain:Domain":
		r = &Domain{}
	case "tencentcloud:Cdn/urlPurge:UrlPurge":
		r = &UrlPurge{}
	case "tencentcloud:Cdn/urlPush:UrlPush":
		r = &UrlPush{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := tencentcloud.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Cdn/domain",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Cdn/urlPurge",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Cdn/urlPush",
		&module{version},
	)
}
