// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of sqlserver datasourceRegions
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const datasourceRegions = pulumi.output(tencentcloud.Sqlserver.getRegions());
 * ```
 */
export function getRegions(args?: GetRegionsArgs, opts?: pulumi.InvokeOptions): Promise<GetRegionsResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:Sqlserver/getRegions:getRegions", {
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getRegions.
 */
export interface GetRegionsArgs {
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getRegions.
 */
export interface GetRegionsResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Region information array.
     */
    readonly regionSets: outputs.Sqlserver.GetRegionsRegionSet[];
    readonly resultOutputFile?: string;
}

export function getRegionsOutput(args?: GetRegionsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRegionsResult> {
    return pulumi.output(args).apply(a => getRegions(a, opts))
}

/**
 * A collection of arguments for invoking getRegions.
 */
export interface GetRegionsOutputArgs {
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
}
