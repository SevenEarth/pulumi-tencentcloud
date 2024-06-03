// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of clb instanceTraffic
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const instanceTraffic = tencentcloud.Clb.getInstanceTraffic({
 *     loadBalancerRegion: "ap-guangzhou",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getInstanceTraffic(args?: GetInstanceTrafficArgs, opts?: pulumi.InvokeOptions): Promise<GetInstanceTrafficResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Clb/getInstanceTraffic:getInstanceTraffic", {
        "loadBalancerRegion": args.loadBalancerRegion,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getInstanceTraffic.
 */
export interface GetInstanceTrafficArgs {
    /**
     * CLB instance region. If this parameter is not passed in, CLB instances in all regions will be returned.
     */
    loadBalancerRegion?: string;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getInstanceTraffic.
 */
export interface GetInstanceTrafficResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly loadBalancerRegion?: string;
    /**
     * Information of CLB instances sorted by outbound bandwidth from highest to lowest. Note: This field may return null, indicating that no valid values can be obtained.
     */
    readonly loadBalancerTraffics: outputs.Clb.GetInstanceTrafficLoadBalancerTraffic[];
    readonly resultOutputFile?: string;
}
/**
 * Use this data source to query detailed information of clb instanceTraffic
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const instanceTraffic = tencentcloud.Clb.getInstanceTraffic({
 *     loadBalancerRegion: "ap-guangzhou",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getInstanceTrafficOutput(args?: GetInstanceTrafficOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetInstanceTrafficResult> {
    return pulumi.output(args).apply((a: any) => getInstanceTraffic(a, opts))
}

/**
 * A collection of arguments for invoking getInstanceTraffic.
 */
export interface GetInstanceTrafficOutputArgs {
    /**
     * CLB instance region. If this parameter is not passed in, CLB instances in all regions will be returned.
     */
    loadBalancerRegion?: pulumi.Input<string>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
}
