// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of monitor tmpRegions
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const tmpRegions = pulumi.output(tencentcloud.Monitor.getTmpRegions({
 *     payMode: 1,
 * }));
 * ```
 */
export function getTmpRegions(args?: GetTmpRegionsArgs, opts?: pulumi.InvokeOptions): Promise<GetTmpRegionsResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:Monitor/getTmpRegions:getTmpRegions", {
        "payMode": args.payMode,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getTmpRegions.
 */
export interface GetTmpRegionsArgs {
    /**
     * Pay mode. `1`-Prepaid, `2`-Postpaid, `3`-All regions (default is all regions if not filled in).
     */
    payMode?: number;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getTmpRegions.
 */
export interface GetTmpRegionsResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly payMode?: number;
    /**
     * Region set.
     */
    readonly regionSets: outputs.Monitor.GetTmpRegionsRegionSet[];
    readonly resultOutputFile?: string;
}

export function getTmpRegionsOutput(args?: GetTmpRegionsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetTmpRegionsResult> {
    return pulumi.output(args).apply(a => getTmpRegions(a, opts))
}

/**
 * A collection of arguments for invoking getTmpRegions.
 */
export interface GetTmpRegionsOutputArgs {
    /**
     * Pay mode. `1`-Prepaid, `2`-Postpaid, `3`-All regions (default is all regions if not filled in).
     */
    payMode?: pulumi.Input<number>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
}