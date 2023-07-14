// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package postgresql

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
	case "tencentcloud:Postgresql/backupDownloadRestrictionConfig:BackupDownloadRestrictionConfig":
		r = &BackupDownloadRestrictionConfig{}
	case "tencentcloud:Postgresql/backupPlanConfig:BackupPlanConfig":
		r = &BackupPlanConfig{}
	case "tencentcloud:Postgresql/baseBackup:BaseBackup":
		r = &BaseBackup{}
	case "tencentcloud:Postgresql/deleteLogBackupOperation:DeleteLogBackupOperation":
		r = &DeleteLogBackupOperation{}
	case "tencentcloud:Postgresql/disisolateDbInstanceOperation:DisisolateDbInstanceOperation":
		r = &DisisolateDbInstanceOperation{}
	case "tencentcloud:Postgresql/instance:Instance":
		r = &Instance{}
	case "tencentcloud:Postgresql/isolateDbInstanceOperation:IsolateDbInstanceOperation":
		r = &IsolateDbInstanceOperation{}
	case "tencentcloud:Postgresql/modifyAccountRemarkOperation:ModifyAccountRemarkOperation":
		r = &ModifyAccountRemarkOperation{}
	case "tencentcloud:Postgresql/modifySwitchTimePeriodOperation:ModifySwitchTimePeriodOperation":
		r = &ModifySwitchTimePeriodOperation{}
	case "tencentcloud:Postgresql/parameterTemplate:ParameterTemplate":
		r = &ParameterTemplate{}
	case "tencentcloud:Postgresql/readonlyAttachment:ReadonlyAttachment":
		r = &ReadonlyAttachment{}
	case "tencentcloud:Postgresql/readonlyGroup:ReadonlyGroup":
		r = &ReadonlyGroup{}
	case "tencentcloud:Postgresql/readonlyInstance:ReadonlyInstance":
		r = &ReadonlyInstance{}
	case "tencentcloud:Postgresql/rebalanceReadonlyGroupOperation:RebalanceReadonlyGroupOperation":
		r = &RebalanceReadonlyGroupOperation{}
	case "tencentcloud:Postgresql/renewDbInstanceOperation:RenewDbInstanceOperation":
		r = &RenewDbInstanceOperation{}
	case "tencentcloud:Postgresql/restartDbInstanceOperation:RestartDbInstanceOperation":
		r = &RestartDbInstanceOperation{}
	case "tencentcloud:Postgresql/securityGroupConfig:SecurityGroupConfig":
		r = &SecurityGroupConfig{}
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
		"Postgresql/backupDownloadRestrictionConfig",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Postgresql/backupPlanConfig",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Postgresql/baseBackup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Postgresql/deleteLogBackupOperation",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Postgresql/disisolateDbInstanceOperation",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Postgresql/instance",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Postgresql/isolateDbInstanceOperation",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Postgresql/modifyAccountRemarkOperation",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Postgresql/modifySwitchTimePeriodOperation",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Postgresql/parameterTemplate",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Postgresql/readonlyAttachment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Postgresql/readonlyGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Postgresql/readonlyInstance",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Postgresql/rebalanceReadonlyGroupOperation",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Postgresql/renewDbInstanceOperation",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Postgresql/restartDbInstanceOperation",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Postgresql/securityGroupConfig",
		&module{version},
	)
}
