// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cbs

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
	case "tencentcloud:Cbs/snapshot:Snapshot":
		r = &Snapshot{}
	case "tencentcloud:Cbs/snapshotPolicy:SnapshotPolicy":
		r = &SnapshotPolicy{}
	case "tencentcloud:Cbs/snapshotPolicyAttachment:SnapshotPolicyAttachment":
		r = &SnapshotPolicyAttachment{}
	case "tencentcloud:Cbs/storage:Storage":
		r = &Storage{}
	case "tencentcloud:Cbs/storageAttachment:StorageAttachment":
		r = &StorageAttachment{}
	case "tencentcloud:Cbs/storageSet:StorageSet":
		r = &StorageSet{}
	case "tencentcloud:Cbs/storageSetAttachment:StorageSetAttachment":
		r = &StorageSetAttachment{}
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
		"Cbs/snapshot",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Cbs/snapshotPolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Cbs/snapshotPolicyAttachment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Cbs/storage",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Cbs/storageAttachment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Cbs/storageSet",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Cbs/storageSetAttachment",
		&module{version},
	)
}
