// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of css watermarks
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const watermarks = pulumi.output(tencentcloud.Css.getWatermarks());
 * ```
 */
export function getWatermarks(args?: GetWatermarksArgs, opts?: pulumi.InvokeOptions): Promise<GetWatermarksResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:Css/getWatermarks:getWatermarks", {
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getWatermarks.
 */
export interface GetWatermarksArgs {
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getWatermarks.
 */
export interface GetWatermarksResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly resultOutputFile?: string;
    /**
     * Watermark information list.
     */
    readonly watermarkLists: outputs.Css.GetWatermarksWatermarkList[];
}

export function getWatermarksOutput(args?: GetWatermarksOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetWatermarksResult> {
    return pulumi.output(args).apply(a => getWatermarks(a, opts))
}

/**
 * A collection of arguments for invoking getWatermarks.
 */
export interface GetWatermarksOutputArgs {
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
}