// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dnspod

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
	case "tencentcloud:Dnspod/customLine:CustomLine":
		r = &CustomLine{}
	case "tencentcloud:Dnspod/domainAlias:DomainAlias":
		r = &DomainAlias{}
	case "tencentcloud:Dnspod/domainInstance:DomainInstance":
		r = &DomainInstance{}
	case "tencentcloud:Dnspod/domainLock:DomainLock":
		r = &DomainLock{}
	case "tencentcloud:Dnspod/downloadSnapshotOperation:DownloadSnapshotOperation":
		r = &DownloadSnapshotOperation{}
	case "tencentcloud:Dnspod/modifyDomainOwnerOperation:ModifyDomainOwnerOperation":
		r = &ModifyDomainOwnerOperation{}
	case "tencentcloud:Dnspod/modifyRecordGroupOperation:ModifyRecordGroupOperation":
		r = &ModifyRecordGroupOperation{}
	case "tencentcloud:Dnspod/record:Record":
		r = &Record{}
	case "tencentcloud:Dnspod/recordGroup:RecordGroup":
		r = &RecordGroup{}
	case "tencentcloud:Dnspod/snapshotConfig:SnapshotConfig":
		r = &SnapshotConfig{}
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
		"Dnspod/customLine",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Dnspod/domainAlias",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Dnspod/domainInstance",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Dnspod/domainLock",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Dnspod/downloadSnapshotOperation",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Dnspod/modifyDomainOwnerOperation",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Dnspod/modifyRecordGroupOperation",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Dnspod/record",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Dnspod/recordGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Dnspod/snapshotConfig",
		&module{version},
	)
}
