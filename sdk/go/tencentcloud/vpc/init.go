// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

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
	case "tencentcloud:Vpc/acl:Acl":
		r = &Acl{}
	case "tencentcloud:Vpc/aclAttachment:AclAttachment":
		r = &AclAttachment{}
	case "tencentcloud:Vpc/bandwidthPackage:BandwidthPackage":
		r = &BandwidthPackage{}
	case "tencentcloud:Vpc/bandwidthPackageAttachment:BandwidthPackageAttachment":
		r = &BandwidthPackageAttachment{}
	case "tencentcloud:Vpc/classicLinkAttachment:ClassicLinkAttachment":
		r = &ClassicLinkAttachment{}
	case "tencentcloud:Vpc/dhcpAssociateAddress:DhcpAssociateAddress":
		r = &DhcpAssociateAddress{}
	case "tencentcloud:Vpc/dhcpIp:DhcpIp":
		r = &DhcpIp{}
	case "tencentcloud:Vpc/enableEndPointConnect:EnableEndPointConnect":
		r = &EnableEndPointConnect{}
	case "tencentcloud:Vpc/endPoint:EndPoint":
		r = &EndPoint{}
	case "tencentcloud:Vpc/endPointService:EndPointService":
		r = &EndPointService{}
	case "tencentcloud:Vpc/endPointServiceWhiteList:EndPointServiceWhiteList":
		r = &EndPointServiceWhiteList{}
	case "tencentcloud:Vpc/flowLog:FlowLog":
		r = &FlowLog{}
	case "tencentcloud:Vpc/flowLogConfig:FlowLogConfig":
		r = &FlowLogConfig{}
	case "tencentcloud:Vpc/instance:Instance":
		r = &Instance{}
	case "tencentcloud:Vpc/ipv6CidrBlock:Ipv6CidrBlock":
		r = &Ipv6CidrBlock{}
	case "tencentcloud:Vpc/ipv6SubnetCidrBlock:Ipv6SubnetCidrBlock":
		r = &Ipv6SubnetCidrBlock{}
	case "tencentcloud:Vpc/localGateway:LocalGateway":
		r = &LocalGateway{}
	case "tencentcloud:Vpc/netDetect:NetDetect":
		r = &NetDetect{}
	case "tencentcloud:Vpc/networkAclQuintuple:NetworkAclQuintuple":
		r = &NetworkAclQuintuple{}
	case "tencentcloud:Vpc/notifyRoutes:NotifyRoutes":
		r = &NotifyRoutes{}
	case "tencentcloud:Vpc/peerConnectAcceptOperation:PeerConnectAcceptOperation":
		r = &PeerConnectAcceptOperation{}
	case "tencentcloud:Vpc/peerConnectManager:PeerConnectManager":
		r = &PeerConnectManager{}
	case "tencentcloud:Vpc/peerConnectRejectOperation:PeerConnectRejectOperation":
		r = &PeerConnectRejectOperation{}
	case "tencentcloud:Vpc/resumeSnapshotInstance:ResumeSnapshotInstance":
		r = &ResumeSnapshotInstance{}
	case "tencentcloud:Vpc/snapshotPolicy:SnapshotPolicy":
		r = &SnapshotPolicy{}
	case "tencentcloud:Vpc/snapshotPolicyAttachment:SnapshotPolicyAttachment":
		r = &SnapshotPolicyAttachment{}
	case "tencentcloud:Vpc/snapshotPolicyConfig:SnapshotPolicyConfig":
		r = &SnapshotPolicyConfig{}
	case "tencentcloud:Vpc/trafficPackage:TrafficPackage":
		r = &TrafficPackage{}
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
		"Vpc/acl",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Vpc/aclAttachment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Vpc/bandwidthPackage",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Vpc/bandwidthPackageAttachment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Vpc/classicLinkAttachment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Vpc/dhcpAssociateAddress",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Vpc/dhcpIp",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Vpc/enableEndPointConnect",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Vpc/endPoint",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Vpc/endPointService",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Vpc/endPointServiceWhiteList",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Vpc/flowLog",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Vpc/flowLogConfig",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Vpc/instance",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Vpc/ipv6CidrBlock",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Vpc/ipv6SubnetCidrBlock",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Vpc/localGateway",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Vpc/netDetect",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Vpc/networkAclQuintuple",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Vpc/notifyRoutes",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Vpc/peerConnectAcceptOperation",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Vpc/peerConnectManager",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Vpc/peerConnectRejectOperation",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Vpc/resumeSnapshotInstance",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Vpc/snapshotPolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Vpc/snapshotPolicyAttachment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Vpc/snapshotPolicyConfig",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Vpc/trafficPackage",
		&module{version},
	)
}
