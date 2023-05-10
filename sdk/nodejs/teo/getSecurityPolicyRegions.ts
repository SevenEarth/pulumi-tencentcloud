// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of teo securityPolicyRegions
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const securityPolicyRegions = pulumi.output(tencentcloud.Teo.getSecurityPolicyRegions());
 * ```
 */
export function getSecurityPolicyRegions(args?: GetSecurityPolicyRegionsArgs, opts?: pulumi.InvokeOptions): Promise<GetSecurityPolicyRegionsResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:Teo/getSecurityPolicyRegions:getSecurityPolicyRegions", {
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getSecurityPolicyRegions.
 */
export interface GetSecurityPolicyRegionsArgs {
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getSecurityPolicyRegions.
 */
export interface GetSecurityPolicyRegionsResult {
    /**
     * Region info.
     */
    readonly geoIps: outputs.Teo.GetSecurityPolicyRegionsGeoIp[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly resultOutputFile?: string;
}

export function getSecurityPolicyRegionsOutput(args?: GetSecurityPolicyRegionsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSecurityPolicyRegionsResult> {
    return pulumi.output(args).apply(a => getSecurityPolicyRegions(a, opts))
}

/**
 * A collection of arguments for invoking getSecurityPolicyRegions.
 */
export interface GetSecurityPolicyRegionsOutputArgs {
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
}
