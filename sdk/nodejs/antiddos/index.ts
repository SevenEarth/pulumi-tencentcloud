// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./ccBlackWhiteIp";
export * from "./ccPrecisionPolicy";
export * from "./ddosBlackWhiteIp";
export * from "./ddosGeoIpBlockConfig";
export * from "./ddosSpeedLimitConfig";
export * from "./defaultAlarmThreshold";
export * from "./getBasicDeviceStatus";
export * from "./getBgpBizTrend";
export * from "./getListListener";
export * from "./getOverviewAttackTrend";
export * from "./getOverviewCcTrend";
export * from "./getOverviewDdosEventList";
export * from "./getOverviewDdosTrend";
export * from "./getOverviewIndex";
export * from "./getPendingRiskInfo";
export * from "./ipAlarmThresholdConfig";
export * from "./packetFilterConfig";
export * from "./portAclConfig";
export * from "./schedulingDomainUserName";

// Import resources to register:
import { CcBlackWhiteIp } from "./ccBlackWhiteIp";
import { CcPrecisionPolicy } from "./ccPrecisionPolicy";
import { DdosBlackWhiteIp } from "./ddosBlackWhiteIp";
import { DdosGeoIpBlockConfig } from "./ddosGeoIpBlockConfig";
import { DdosSpeedLimitConfig } from "./ddosSpeedLimitConfig";
import { DefaultAlarmThreshold } from "./defaultAlarmThreshold";
import { IpAlarmThresholdConfig } from "./ipAlarmThresholdConfig";
import { PacketFilterConfig } from "./packetFilterConfig";
import { PortAclConfig } from "./portAclConfig";
import { SchedulingDomainUserName } from "./schedulingDomainUserName";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "tencentcloud:Antiddos/ccBlackWhiteIp:CcBlackWhiteIp":
                return new CcBlackWhiteIp(name, <any>undefined, { urn })
            case "tencentcloud:Antiddos/ccPrecisionPolicy:CcPrecisionPolicy":
                return new CcPrecisionPolicy(name, <any>undefined, { urn })
            case "tencentcloud:Antiddos/ddosBlackWhiteIp:DdosBlackWhiteIp":
                return new DdosBlackWhiteIp(name, <any>undefined, { urn })
            case "tencentcloud:Antiddos/ddosGeoIpBlockConfig:DdosGeoIpBlockConfig":
                return new DdosGeoIpBlockConfig(name, <any>undefined, { urn })
            case "tencentcloud:Antiddos/ddosSpeedLimitConfig:DdosSpeedLimitConfig":
                return new DdosSpeedLimitConfig(name, <any>undefined, { urn })
            case "tencentcloud:Antiddos/defaultAlarmThreshold:DefaultAlarmThreshold":
                return new DefaultAlarmThreshold(name, <any>undefined, { urn })
            case "tencentcloud:Antiddos/ipAlarmThresholdConfig:IpAlarmThresholdConfig":
                return new IpAlarmThresholdConfig(name, <any>undefined, { urn })
            case "tencentcloud:Antiddos/packetFilterConfig:PacketFilterConfig":
                return new PacketFilterConfig(name, <any>undefined, { urn })
            case "tencentcloud:Antiddos/portAclConfig:PortAclConfig":
                return new PortAclConfig(name, <any>undefined, { urn })
            case "tencentcloud:Antiddos/schedulingDomainUserName:SchedulingDomainUserName":
                return new SchedulingDomainUserName(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("tencentcloud", "Antiddos/ccBlackWhiteIp", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Antiddos/ccPrecisionPolicy", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Antiddos/ddosBlackWhiteIp", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Antiddos/ddosGeoIpBlockConfig", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Antiddos/ddosSpeedLimitConfig", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Antiddos/defaultAlarmThreshold", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Antiddos/ipAlarmThresholdConfig", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Antiddos/packetFilterConfig", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Antiddos/portAclConfig", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Antiddos/schedulingDomainUserName", _module)
