// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package lighthouse

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
	case "tencentcloud:Lighthouse/applyDiskBackup:ApplyDiskBackup":
		r = &ApplyDiskBackup{}
	case "tencentcloud:Lighthouse/applyInstanceSnapshot:ApplyInstanceSnapshot":
		r = &ApplyInstanceSnapshot{}
	case "tencentcloud:Lighthouse/blueprint:Blueprint":
		r = &Blueprint{}
	case "tencentcloud:Lighthouse/disk:Disk":
		r = &Disk{}
	case "tencentcloud:Lighthouse/diskAttachment:DiskAttachment":
		r = &DiskAttachment{}
	case "tencentcloud:Lighthouse/diskBackup:DiskBackup":
		r = &DiskBackup{}
	case "tencentcloud:Lighthouse/firewallRule:FirewallRule":
		r = &FirewallRule{}
	case "tencentcloud:Lighthouse/firewallTemplate:FirewallTemplate":
		r = &FirewallTemplate{}
	case "tencentcloud:Lighthouse/instance:Instance":
		r = &Instance{}
	case "tencentcloud:Lighthouse/keyPair:KeyPair":
		r = &KeyPair{}
	case "tencentcloud:Lighthouse/keyPairAttachment:KeyPairAttachment":
		r = &KeyPairAttachment{}
	case "tencentcloud:Lighthouse/rebootInstance:RebootInstance":
		r = &RebootInstance{}
	case "tencentcloud:Lighthouse/renewDisk:RenewDisk":
		r = &RenewDisk{}
	case "tencentcloud:Lighthouse/renewInstance:RenewInstance":
		r = &RenewInstance{}
	case "tencentcloud:Lighthouse/snapshot:Snapshot":
		r = &Snapshot{}
	case "tencentcloud:Lighthouse/startInstance:StartInstance":
		r = &StartInstance{}
	case "tencentcloud:Lighthouse/stopInstance:StopInstance":
		r = &StopInstance{}
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
		"Lighthouse/applyDiskBackup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Lighthouse/applyInstanceSnapshot",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Lighthouse/blueprint",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Lighthouse/disk",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Lighthouse/diskAttachment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Lighthouse/diskBackup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Lighthouse/firewallRule",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Lighthouse/firewallTemplate",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Lighthouse/instance",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Lighthouse/keyPair",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Lighthouse/keyPairAttachment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Lighthouse/rebootInstance",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Lighthouse/renewDisk",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Lighthouse/renewInstance",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Lighthouse/snapshot",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Lighthouse/startInstance",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Lighthouse/stopInstance",
		&module{version},
	)
}
