// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of clb resources
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const resources = pulumi.output(tencentcloud.Clb.getResources({
 *     filters: [{
 *         name: "isp",
 *         values: ["BGP"],
 *     }],
 * }));
 * ```
 */
export function getResources(args?: GetResourcesArgs, opts?: pulumi.InvokeOptions): Promise<GetResourcesResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:Clb/getResources:getResources", {
        "filters": args.filters,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getResources.
 */
export interface GetResourcesArgs {
    /**
     * Filter to query the list of AZ resources as detailed below: zone - String - Optional - Filter by AZ, such as ap-guangzhou-1. isp -- String - Optional - Filter by the ISP. Values: BGP, CMCC, CUCC and CTCC.
     */
    filters?: inputs.Clb.GetResourcesFilter[];
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getResources.
 */
export interface GetResourcesResult {
    readonly filters?: outputs.Clb.GetResourcesFilter[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly resultOutputFile?: string;
    /**
     * List of resources supported by the AZ.
     */
    readonly zoneResourceSets: outputs.Clb.GetResourcesZoneResourceSet[];
}

export function getResourcesOutput(args?: GetResourcesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetResourcesResult> {
    return pulumi.output(args).apply(a => getResources(a, opts))
}

/**
 * A collection of arguments for invoking getResources.
 */
export interface GetResourcesOutputArgs {
    /**
     * Filter to query the list of AZ resources as detailed below: zone - String - Optional - Filter by AZ, such as ap-guangzhou-1. isp -- String - Optional - Filter by the ISP. Values: BGP, CMCC, CUCC and CTCC.
     */
    filters?: pulumi.Input<pulumi.Input<inputs.Clb.GetResourcesFilterArgs>[]>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
}
