// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of vpc limits
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const limits = tencentcloud.Vpc.getLimits({
 *     limitTypes: [
 *         "appid-max-vpcs",
 *         "vpc-max-subnets",
 *     ],
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getLimits(args: GetLimitsArgs, opts?: pulumi.InvokeOptions): Promise<GetLimitsResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Vpc/getLimits:getLimits", {
        "limitTypes": args.limitTypes,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getLimits.
 */
export interface GetLimitsArgs {
    /**
     * Quota name. A maximum of 100 quota types can be queried each time.
     */
    limitTypes: string[];
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getLimits.
 */
export interface GetLimitsResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly limitTypes: string[];
    readonly resultOutputFile?: string;
    /**
     * vpc limit.
     */
    readonly vpcLimitSets: outputs.Vpc.GetLimitsVpcLimitSet[];
}
/**
 * Use this data source to query detailed information of vpc limits
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const limits = tencentcloud.Vpc.getLimits({
 *     limitTypes: [
 *         "appid-max-vpcs",
 *         "vpc-max-subnets",
 *     ],
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getLimitsOutput(args: GetLimitsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetLimitsResult> {
    return pulumi.output(args).apply((a: any) => getLimits(a, opts))
}

/**
 * A collection of arguments for invoking getLimits.
 */
export interface GetLimitsOutputArgs {
    /**
     * Quota name. A maximum of 100 quota types can be queried each time.
     */
    limitTypes: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
}
