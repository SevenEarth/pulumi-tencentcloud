// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of chdfs accessGroups
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const accessGroups = pulumi.output(tencentcloud.Chdfs.getAccessGroups({
 *     vpcId: "vpc-pewdpc0d",
 * }));
 * ```
 */
export function getAccessGroups(args?: GetAccessGroupsArgs, opts?: pulumi.InvokeOptions): Promise<GetAccessGroupsResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:Chdfs/getAccessGroups:getAccessGroups", {
        "ownerUin": args.ownerUin,
        "resultOutputFile": args.resultOutputFile,
        "vpcId": args.vpcId,
    }, opts);
}

/**
 * A collection of arguments for invoking getAccessGroups.
 */
export interface GetAccessGroupsArgs {
    /**
     * get groups belongs to the owner uin, must set but only can use one of VpcId and OwnerUin to get the groups.
     */
    ownerUin?: number;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
    /**
     * get groups belongs to the vpc id, must set but only can use one of VpcId and OwnerUin to get the groups.
     */
    vpcId?: string;
}

/**
 * A collection of values returned by getAccessGroups.
 */
export interface GetAccessGroupsResult {
    /**
     * access group list.
     */
    readonly accessGroups: outputs.Chdfs.GetAccessGroupsAccessGroup[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ownerUin?: number;
    readonly resultOutputFile?: string;
    /**
     * VPC ID.
     */
    readonly vpcId?: string;
}

export function getAccessGroupsOutput(args?: GetAccessGroupsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAccessGroupsResult> {
    return pulumi.output(args).apply(a => getAccessGroups(a, opts))
}

/**
 * A collection of arguments for invoking getAccessGroups.
 */
export interface GetAccessGroupsOutputArgs {
    /**
     * get groups belongs to the owner uin, must set but only can use one of VpcId and OwnerUin to get the groups.
     */
    ownerUin?: pulumi.Input<number>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
    /**
     * get groups belongs to the vpc id, must set but only can use one of VpcId and OwnerUin to get the groups.
     */
    vpcId?: pulumi.Input<string>;
}
