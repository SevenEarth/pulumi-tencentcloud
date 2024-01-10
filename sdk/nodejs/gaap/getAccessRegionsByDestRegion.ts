// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of gaap access regions by dest region
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const accessRegionsByDestRegion = pulumi.output(tencentcloud.Gaap.getAccessRegionsByDestRegion({
 *     destRegion: "SouthChina",
 * }));
 * ```
 */
export function getAccessRegionsByDestRegion(args: GetAccessRegionsByDestRegionArgs, opts?: pulumi.InvokeOptions): Promise<GetAccessRegionsByDestRegionResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:Gaap/getAccessRegionsByDestRegion:getAccessRegionsByDestRegion", {
        "destRegion": args.destRegion,
        "ipAddressVersion": args.ipAddressVersion,
        "packageType": args.packageType,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getAccessRegionsByDestRegion.
 */
export interface GetAccessRegionsByDestRegionArgs {
    /**
     * Origin region.
     */
    destRegion: string;
    /**
     * IP version, can be taken as IPv4 or IPv6, with a default value of IPv4.
     */
    ipAddressVersion?: string;
    /**
     * Channel package type, where Thunder represents a standard proxy group, Accelerator represents a game accelerator proxy, and CrossBorder represents a cross-border proxy.
     */
    packageType?: string;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getAccessRegionsByDestRegion.
 */
export interface GetAccessRegionsByDestRegionResult {
    /**
     * List of available acceleration zone information.
     */
    readonly accessRegionSets: outputs.Gaap.GetAccessRegionsByDestRegionAccessRegionSet[];
    readonly destRegion: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ipAddressVersion?: string;
    readonly packageType?: string;
    readonly resultOutputFile?: string;
}

export function getAccessRegionsByDestRegionOutput(args: GetAccessRegionsByDestRegionOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAccessRegionsByDestRegionResult> {
    return pulumi.output(args).apply(a => getAccessRegionsByDestRegion(a, opts))
}

/**
 * A collection of arguments for invoking getAccessRegionsByDestRegion.
 */
export interface GetAccessRegionsByDestRegionOutputArgs {
    /**
     * Origin region.
     */
    destRegion: pulumi.Input<string>;
    /**
     * IP version, can be taken as IPv4 or IPv6, with a default value of IPv4.
     */
    ipAddressVersion?: pulumi.Input<string>;
    /**
     * Channel package type, where Thunder represents a standard proxy group, Accelerator represents a game accelerator proxy, and CrossBorder represents a cross-border proxy.
     */
    packageType?: pulumi.Input<string>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
}
