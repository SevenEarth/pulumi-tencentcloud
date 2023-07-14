// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./attachment";
export * from "./completeLifecycle";
export * from "./executeScalingPolicy";
export * from "./getAdvices";
export * from "./getInstances";
export * from "./getLastActivity";
export * from "./getLimits";
export * from "./getScalingConfigs";
export * from "./getScalingGroups";
export * from "./getScalingPolicies";
export * from "./lifecycleHook";
export * from "./loadBalancer";
export * from "./notification";
export * from "./protectInstances";
export * from "./removeInstances";
export * from "./scaleInInstances";
export * from "./scaleOutInstances";
export * from "./scalingConfig";
export * from "./scalingGroup";
export * from "./scalingGroupStatus";
export * from "./scalingPolicy";
export * from "./schedule";
export * from "./startInstances";
export * from "./stopInstances";

// Import resources to register:
import { Attachment } from "./attachment";
import { CompleteLifecycle } from "./completeLifecycle";
import { ExecuteScalingPolicy } from "./executeScalingPolicy";
import { LifecycleHook } from "./lifecycleHook";
import { LoadBalancer } from "./loadBalancer";
import { Notification } from "./notification";
import { ProtectInstances } from "./protectInstances";
import { RemoveInstances } from "./removeInstances";
import { ScaleInInstances } from "./scaleInInstances";
import { ScaleOutInstances } from "./scaleOutInstances";
import { ScalingConfig } from "./scalingConfig";
import { ScalingGroup } from "./scalingGroup";
import { ScalingGroupStatus } from "./scalingGroupStatus";
import { ScalingPolicy } from "./scalingPolicy";
import { Schedule } from "./schedule";
import { StartInstances } from "./startInstances";
import { StopInstances } from "./stopInstances";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "tencentcloud:As/attachment:Attachment":
                return new Attachment(name, <any>undefined, { urn })
            case "tencentcloud:As/completeLifecycle:CompleteLifecycle":
                return new CompleteLifecycle(name, <any>undefined, { urn })
            case "tencentcloud:As/executeScalingPolicy:ExecuteScalingPolicy":
                return new ExecuteScalingPolicy(name, <any>undefined, { urn })
            case "tencentcloud:As/lifecycleHook:LifecycleHook":
                return new LifecycleHook(name, <any>undefined, { urn })
            case "tencentcloud:As/loadBalancer:LoadBalancer":
                return new LoadBalancer(name, <any>undefined, { urn })
            case "tencentcloud:As/notification:Notification":
                return new Notification(name, <any>undefined, { urn })
            case "tencentcloud:As/protectInstances:ProtectInstances":
                return new ProtectInstances(name, <any>undefined, { urn })
            case "tencentcloud:As/removeInstances:RemoveInstances":
                return new RemoveInstances(name, <any>undefined, { urn })
            case "tencentcloud:As/scaleInInstances:ScaleInInstances":
                return new ScaleInInstances(name, <any>undefined, { urn })
            case "tencentcloud:As/scaleOutInstances:ScaleOutInstances":
                return new ScaleOutInstances(name, <any>undefined, { urn })
            case "tencentcloud:As/scalingConfig:ScalingConfig":
                return new ScalingConfig(name, <any>undefined, { urn })
            case "tencentcloud:As/scalingGroup:ScalingGroup":
                return new ScalingGroup(name, <any>undefined, { urn })
            case "tencentcloud:As/scalingGroupStatus:ScalingGroupStatus":
                return new ScalingGroupStatus(name, <any>undefined, { urn })
            case "tencentcloud:As/scalingPolicy:ScalingPolicy":
                return new ScalingPolicy(name, <any>undefined, { urn })
            case "tencentcloud:As/schedule:Schedule":
                return new Schedule(name, <any>undefined, { urn })
            case "tencentcloud:As/startInstances:StartInstances":
                return new StartInstances(name, <any>undefined, { urn })
            case "tencentcloud:As/stopInstances:StopInstances":
                return new StopInstances(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("tencentcloud", "As/attachment", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "As/completeLifecycle", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "As/executeScalingPolicy", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "As/lifecycleHook", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "As/loadBalancer", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "As/notification", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "As/protectInstances", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "As/removeInstances", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "As/scaleInInstances", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "As/scaleOutInstances", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "As/scalingConfig", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "As/scalingGroup", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "As/scalingGroupStatus", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "As/scalingPolicy", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "As/schedule", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "As/startInstances", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "As/stopInstances", _module)
