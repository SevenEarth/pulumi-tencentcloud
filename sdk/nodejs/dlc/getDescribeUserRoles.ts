// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of dlc describeUserRoles
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const describeUserRoles = pulumi.output(tencentcloud.Dlc.getDescribeUserRoles({
 *     fuzzy: "1",
 * }));
 * ```
 */
export function getDescribeUserRoles(args?: GetDescribeUserRolesArgs, opts?: pulumi.InvokeOptions): Promise<GetDescribeUserRolesResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:Dlc/getDescribeUserRoles:getDescribeUserRoles", {
        "fuzzy": args.fuzzy,
        "resultOutputFile": args.resultOutputFile,
        "sortBy": args.sortBy,
        "sorting": args.sorting,
    }, opts);
}

/**
 * A collection of arguments for invoking getDescribeUserRoles.
 */
export interface GetDescribeUserRolesArgs {
    /**
     * List according to ARN blur.
     */
    fuzzy?: string;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
    /**
     * The return results are sorted according to this field.
     */
    sortBy?: string;
    /**
     * Positive or inverted, such as DESC.
     */
    sorting?: string;
}

/**
 * A collection of values returned by getDescribeUserRoles.
 */
export interface GetDescribeUserRolesResult {
    readonly fuzzy?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly resultOutputFile?: string;
    readonly sortBy?: string;
    readonly sorting?: string;
    /**
     * User role information.
     */
    readonly userRoles: outputs.Dlc.GetDescribeUserRolesUserRole[];
}

export function getDescribeUserRolesOutput(args?: GetDescribeUserRolesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDescribeUserRolesResult> {
    return pulumi.output(args).apply(a => getDescribeUserRoles(a, opts))
}

/**
 * A collection of arguments for invoking getDescribeUserRoles.
 */
export interface GetDescribeUserRolesOutputArgs {
    /**
     * List according to ARN blur.
     */
    fuzzy?: pulumi.Input<string>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
    /**
     * The return results are sorted according to this field.
     */
    sortBy?: pulumi.Input<string>;
    /**
     * Positive or inverted, such as DESC.
     */
    sorting?: pulumi.Input<string>;
}