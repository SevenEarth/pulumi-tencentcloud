// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of teo zoneAvailablePlans
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const zoneAvailablePlans = pulumi.output(tencentcloud.Teo.getZoneAvailablePlans());
 * ```
 */
export function getZoneAvailablePlans(args?: GetZoneAvailablePlansArgs, opts?: pulumi.InvokeOptions): Promise<GetZoneAvailablePlansResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:Teo/getZoneAvailablePlans:getZoneAvailablePlans", {
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getZoneAvailablePlans.
 */
export interface GetZoneAvailablePlansArgs {
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getZoneAvailablePlans.
 */
export interface GetZoneAvailablePlansResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Zone plans which current account can use.
     */
    readonly planInfoLists: outputs.Teo.GetZoneAvailablePlansPlanInfoList[];
    readonly resultOutputFile?: string;
}

export function getZoneAvailablePlansOutput(args?: GetZoneAvailablePlansOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetZoneAvailablePlansResult> {
    return pulumi.output(args).apply(a => getZoneAvailablePlans(a, opts))
}

/**
 * A collection of arguments for invoking getZoneAvailablePlans.
 */
export interface GetZoneAvailablePlansOutputArgs {
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
}
