// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of redis instanceNodeInfo
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const instanceNodeInfo = pulumi.output(tencentcloud.Redis.getInstanceNodeInfo({
 *     instanceId: "crs-c1nl9rpv",
 * }));
 * ```
 */
export function getInstanceNodeInfo(args: GetInstanceNodeInfoArgs, opts?: pulumi.InvokeOptions): Promise<GetInstanceNodeInfoResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:Redis/getInstanceNodeInfo:getInstanceNodeInfo", {
        "instanceId": args.instanceId,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getInstanceNodeInfo.
 */
export interface GetInstanceNodeInfoArgs {
    /**
     * The ID of instance.
     */
    instanceId: string;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getInstanceNodeInfo.
 */
export interface GetInstanceNodeInfoResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceId: string;
    /**
     * Proxy node information.
     */
    readonly proxies: outputs.Redis.GetInstanceNodeInfoProxy[];
    /**
     * Number of proxy nodes.
     */
    readonly proxyCount: number;
    /**
     * Redis node information.
     */
    readonly redis: outputs.Redis.GetInstanceNodeInfoRedi[];
    /**
     * Number of redis nodes.
     */
    readonly redisCount: number;
    readonly resultOutputFile?: string;
}

export function getInstanceNodeInfoOutput(args: GetInstanceNodeInfoOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetInstanceNodeInfoResult> {
    return pulumi.output(args).apply(a => getInstanceNodeInfo(a, opts))
}

/**
 * A collection of arguments for invoking getInstanceNodeInfo.
 */
export interface GetInstanceNodeInfoOutputArgs {
    /**
     * The ID of instance.
     */
    instanceId: pulumi.Input<string>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
}
