// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package css

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
	case "tencentcloud:Css/authenticateDomainOwnerOperation:AuthenticateDomainOwnerOperation":
		r = &AuthenticateDomainOwnerOperation{}
	case "tencentcloud:Css/backupStream:BackupStream":
		r = &BackupStream{}
	case "tencentcloud:Css/callbackRuleAttachment:CallbackRuleAttachment":
		r = &CallbackRuleAttachment{}
	case "tencentcloud:Css/callbackTemplate:CallbackTemplate":
		r = &CallbackTemplate{}
	case "tencentcloud:Css/domain:Domain":
		r = &Domain{}
	case "tencentcloud:Css/domainReferer:DomainReferer":
		r = &DomainReferer{}
	case "tencentcloud:Css/enableOptimalSwitching:EnableOptimalSwitching":
		r = &EnableOptimalSwitching{}
	case "tencentcloud:Css/liveTranscodeRuleAttachment:LiveTranscodeRuleAttachment":
		r = &LiveTranscodeRuleAttachment{}
	case "tencentcloud:Css/liveTranscodeTemplate:LiveTranscodeTemplate":
		r = &LiveTranscodeTemplate{}
	case "tencentcloud:Css/padRuleAttachment:PadRuleAttachment":
		r = &PadRuleAttachment{}
	case "tencentcloud:Css/padTemplate:PadTemplate":
		r = &PadTemplate{}
	case "tencentcloud:Css/playAuthKeyConfig:PlayAuthKeyConfig":
		r = &PlayAuthKeyConfig{}
	case "tencentcloud:Css/playDomainCertAttachment:PlayDomainCertAttachment":
		r = &PlayDomainCertAttachment{}
	case "tencentcloud:Css/pullStreamTask:PullStreamTask":
		r = &PullStreamTask{}
	case "tencentcloud:Css/pullStreamTaskRestart:PullStreamTaskRestart":
		r = &PullStreamTaskRestart{}
	case "tencentcloud:Css/pushAuthKeyConfig:PushAuthKeyConfig":
		r = &PushAuthKeyConfig{}
	case "tencentcloud:Css/recordRuleAttachment:RecordRuleAttachment":
		r = &RecordRuleAttachment{}
	case "tencentcloud:Css/recordTemplate:RecordTemplate":
		r = &RecordTemplate{}
	case "tencentcloud:Css/snapshotRuleAttachment:SnapshotRuleAttachment":
		r = &SnapshotRuleAttachment{}
	case "tencentcloud:Css/snapshotTemplate:SnapshotTemplate":
		r = &SnapshotTemplate{}
	case "tencentcloud:Css/startStreamMonitor:StartStreamMonitor":
		r = &StartStreamMonitor{}
	case "tencentcloud:Css/streamMonitor:StreamMonitor":
		r = &StreamMonitor{}
	case "tencentcloud:Css/timeshiftRuleAttachment:TimeshiftRuleAttachment":
		r = &TimeshiftRuleAttachment{}
	case "tencentcloud:Css/timeshiftTemplate:TimeshiftTemplate":
		r = &TimeshiftTemplate{}
	case "tencentcloud:Css/watermark:Watermark":
		r = &Watermark{}
	case "tencentcloud:Css/watermarkRuleAttachment:WatermarkRuleAttachment":
		r = &WatermarkRuleAttachment{}
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
		"Css/authenticateDomainOwnerOperation",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Css/backupStream",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Css/callbackRuleAttachment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Css/callbackTemplate",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Css/domain",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Css/domainReferer",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Css/enableOptimalSwitching",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Css/liveTranscodeRuleAttachment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Css/liveTranscodeTemplate",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Css/padRuleAttachment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Css/padTemplate",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Css/playAuthKeyConfig",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Css/playDomainCertAttachment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Css/pullStreamTask",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Css/pullStreamTaskRestart",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Css/pushAuthKeyConfig",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Css/recordRuleAttachment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Css/recordTemplate",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Css/snapshotRuleAttachment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Css/snapshotTemplate",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Css/startStreamMonitor",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Css/streamMonitor",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Css/timeshiftRuleAttachment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Css/timeshiftTemplate",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Css/watermark",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Css/watermarkRuleAttachment",
		&module{version},
	)
}
