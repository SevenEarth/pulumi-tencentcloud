// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of sqlserver datasourceCrossRegionZone
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const example = tencentcloud.Sqlserver.getCrossRegionZone({
 *     instanceId: "mssql-qelbzgwf",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getCrossRegionZone(args: GetCrossRegionZoneArgs, opts?: pulumi.InvokeOptions): Promise<GetCrossRegionZoneResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Sqlserver/getCrossRegionZone:getCrossRegionZone", {
        "instanceId": args.instanceId,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getCrossRegionZone.
 */
export interface GetCrossRegionZoneArgs {
    /**
     * Instance ID in the format of mssql-j8kv137v.
     */
    instanceId: string;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getCrossRegionZone.
 */
export interface GetCrossRegionZoneResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceId: string;
    /**
     * The string ID of the region where the standby machine is located, such as: ap-guangzhou.
     */
    readonly region: string;
    readonly resultOutputFile?: string;
    /**
     * The string ID of the availability zone where the standby machine is located, such as: ap-guangzhou-1.
     */
    readonly zone: string;
}
/**
 * Use this data source to query detailed information of sqlserver datasourceCrossRegionZone
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const example = tencentcloud.Sqlserver.getCrossRegionZone({
 *     instanceId: "mssql-qelbzgwf",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getCrossRegionZoneOutput(args: GetCrossRegionZoneOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetCrossRegionZoneResult> {
    return pulumi.output(args).apply((a: any) => getCrossRegionZone(a, opts))
}

/**
 * A collection of arguments for invoking getCrossRegionZone.
 */
export interface GetCrossRegionZoneOutputArgs {
    /**
     * Instance ID in the format of mssql-j8kv137v.
     */
    instanceId: pulumi.Input<string>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
}
