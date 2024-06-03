// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of ccnCrossBorderRegionBandwidthLimits
 *
 * > **NOTE:** This resource is dedicated to Unicom.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const ccnRegionBandwidthLimits = tencentcloud.Ccn.getCrossBorderRegionBandwidthLimits({
 *     filters: [
 *         {
 *             name: "source-region",
 *             values: ["ap-guangzhou"],
 *         },
 *         {
 *             name: "destination-region",
 *             values: ["ap-shanghai"],
 *         },
 *     ],
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getCrossBorderRegionBandwidthLimits(args?: GetCrossBorderRegionBandwidthLimitsArgs, opts?: pulumi.InvokeOptions): Promise<GetCrossBorderRegionBandwidthLimitsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Ccn/getCrossBorderRegionBandwidthLimits:getCrossBorderRegionBandwidthLimits", {
        "filters": args.filters,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getCrossBorderRegionBandwidthLimits.
 */
export interface GetCrossBorderRegionBandwidthLimitsArgs {
    /**
     * Filter condition. Currently, only one value is supported. The supported fields, 1)source-region, the value is like ap-guangzhou; 2)destination-region, the value is like ap-shanghai; 3)ccn-ids,cloud network ID array, the value is like ccn-12345678; 4)user-account-id,user account ID, the value is like 12345678.
     */
    filters?: inputs.Ccn.GetCrossBorderRegionBandwidthLimitsFilter[];
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getCrossBorderRegionBandwidthLimits.
 */
export interface GetCrossBorderRegionBandwidthLimitsResult {
    /**
     * Info of cross region ccn instance.
     */
    readonly ccnBandwidthSets: outputs.Ccn.GetCrossBorderRegionBandwidthLimitsCcnBandwidthSet[];
    readonly filters?: outputs.Ccn.GetCrossBorderRegionBandwidthLimitsFilter[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly resultOutputFile?: string;
}
/**
 * Use this data source to query detailed information of ccnCrossBorderRegionBandwidthLimits
 *
 * > **NOTE:** This resource is dedicated to Unicom.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const ccnRegionBandwidthLimits = tencentcloud.Ccn.getCrossBorderRegionBandwidthLimits({
 *     filters: [
 *         {
 *             name: "source-region",
 *             values: ["ap-guangzhou"],
 *         },
 *         {
 *             name: "destination-region",
 *             values: ["ap-shanghai"],
 *         },
 *     ],
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getCrossBorderRegionBandwidthLimitsOutput(args?: GetCrossBorderRegionBandwidthLimitsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetCrossBorderRegionBandwidthLimitsResult> {
    return pulumi.output(args).apply((a: any) => getCrossBorderRegionBandwidthLimits(a, opts))
}

/**
 * A collection of arguments for invoking getCrossBorderRegionBandwidthLimits.
 */
export interface GetCrossBorderRegionBandwidthLimitsOutputArgs {
    /**
     * Filter condition. Currently, only one value is supported. The supported fields, 1)source-region, the value is like ap-guangzhou; 2)destination-region, the value is like ap-shanghai; 3)ccn-ids,cloud network ID array, the value is like ccn-12345678; 4)user-account-id,user account ID, the value is like 12345678.
     */
    filters?: pulumi.Input<pulumi.Input<inputs.Ccn.GetCrossBorderRegionBandwidthLimitsFilterArgs>[]>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
}
