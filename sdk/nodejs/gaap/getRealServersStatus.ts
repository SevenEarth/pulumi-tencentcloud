// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of gaap real servers status
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const realServersStatus = tencentcloud.Gaap.getRealServersStatus({
 *     realServerIds: ["rs-3mlpbuut"],
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getRealServersStatus(args: GetRealServersStatusArgs, opts?: pulumi.InvokeOptions): Promise<GetRealServersStatusResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Gaap/getRealServersStatus:getRealServersStatus", {
        "realServerIds": args.realServerIds,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getRealServersStatus.
 */
export interface GetRealServersStatusArgs {
    /**
     * Real Server Ids.
     */
    realServerIds: string[];
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getRealServersStatus.
 */
export interface GetRealServersStatusResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly realServerIds: string[];
    /**
     * Real Server Status Set.
     */
    readonly realServerStatusSets: outputs.Gaap.GetRealServersStatusRealServerStatusSet[];
    readonly resultOutputFile?: string;
}
/**
 * Use this data source to query detailed information of gaap real servers status
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const realServersStatus = tencentcloud.Gaap.getRealServersStatus({
 *     realServerIds: ["rs-3mlpbuut"],
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getRealServersStatusOutput(args: GetRealServersStatusOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRealServersStatusResult> {
    return pulumi.output(args).apply((a: any) => getRealServersStatus(a, opts))
}

/**
 * A collection of arguments for invoking getRealServersStatus.
 */
export interface GetRealServersStatusOutputArgs {
    /**
     * Real Server Ids.
     */
    realServerIds: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
}
