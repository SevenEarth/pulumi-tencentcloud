// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query target group information.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@tencentcloud_iac/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const clbBasic = new tencentcloud.clb.Instance("clbBasic", {
 *     networkType: "OPEN",
 *     clbName: "tf-clb-rule-basic",
 * });
 * const listenerBasic = new tencentcloud.clb.Listener("listenerBasic", {
 *     clbId: clbBasic.id,
 *     port: 1,
 *     protocol: "HTTP",
 *     listenerName: "listener_basic",
 * });
 * const ruleBasic = new tencentcloud.clb.ListenerRule("ruleBasic", {
 *     clbId: clbBasic.id,
 *     listenerId: listenerBasic.listenerId,
 *     domain: "abc.com",
 *     url: "/",
 *     sessionExpireTime: 30,
 *     scheduler: "WRR",
 *     targetType: "TARGETGROUP",
 * });
 * const test = new tencentcloud.clb.TargetGroup("test", {targetGroupName: "test-target-keep-1"});
 * const group = new tencentcloud.clb.TargetGroupAttachment("group", {
 *     clbId: clbBasic.id,
 *     listenerId: listenerBasic.listenerId,
 *     ruleId: ruleBasic.ruleId,
 *     targrtGroupId: test.id,
 * });
 * const targetGroupInfoId = tencentcloud.Clb.getTargetGroupsOutput({
 *     targetGroupId: test.id,
 * });
 * ```
 */
export function getTargetGroups(args?: GetTargetGroupsArgs, opts?: pulumi.InvokeOptions): Promise<GetTargetGroupsResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:Clb/getTargetGroups:getTargetGroups", {
        "resultOutputFile": args.resultOutputFile,
        "targetGroupId": args.targetGroupId,
        "targetGroupName": args.targetGroupName,
        "vpcId": args.vpcId,
    }, opts);
}

/**
 * A collection of arguments for invoking getTargetGroups.
 */
export interface GetTargetGroupsArgs {
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
    /**
     * ID of Target group. Mutually exclusive with `vpcId` and `targetGroupName`. `targetGroupId` is preferred.
     */
    targetGroupId?: string;
    /**
     * Name of target group. Mutually exclusive with `targetGroupId`. `targetGroupId` is preferred.
     */
    targetGroupName?: string;
    /**
     * Target group VPC ID. Mutually exclusive with `targetGroupId`. `targetGroupId` is preferred.
     */
    vpcId?: string;
}

/**
 * A collection of values returned by getTargetGroups.
 */
export interface GetTargetGroupsResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Target group info list.
     */
    readonly lists: outputs.Clb.GetTargetGroupsList[];
    readonly resultOutputFile?: string;
    /**
     * ID of Target group.
     */
    readonly targetGroupId?: string;
    /**
     * Target group VPC ID.
     */
    readonly targetGroupName?: string;
    /**
     * Name of target group.
     */
    readonly vpcId?: string;
}

export function getTargetGroupsOutput(args?: GetTargetGroupsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetTargetGroupsResult> {
    return pulumi.output(args).apply(a => getTargetGroups(a, opts))
}

/**
 * A collection of arguments for invoking getTargetGroups.
 */
export interface GetTargetGroupsOutputArgs {
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
    /**
     * ID of Target group. Mutually exclusive with `vpcId` and `targetGroupName`. `targetGroupId` is preferred.
     */
    targetGroupId?: pulumi.Input<string>;
    /**
     * Name of target group. Mutually exclusive with `targetGroupId`. `targetGroupId` is preferred.
     */
    targetGroupName?: pulumi.Input<string>;
    /**
     * Target group VPC ID. Mutually exclusive with `targetGroupId`. `targetGroupId` is preferred.
     */
    vpcId?: pulumi.Input<string>;
}
