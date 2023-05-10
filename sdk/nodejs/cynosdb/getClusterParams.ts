// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of cynosdb clusterParams
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const clusterParams = pulumi.output(tencentcloud.Cynosdb.getClusterParams({
 *     clusterId: "cynosdbmysql-bws8h88b",
 *     paramName: "innodb_checksum_algorithm",
 * }));
 * ```
 */
export function getClusterParams(args: GetClusterParamsArgs, opts?: pulumi.InvokeOptions): Promise<GetClusterParamsResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:Cynosdb/getClusterParams:getClusterParams", {
        "clusterId": args.clusterId,
        "paramName": args.paramName,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getClusterParams.
 */
export interface GetClusterParamsArgs {
    /**
     * The ID of cluster.
     */
    clusterId: string;
    /**
     * Parameter name.
     */
    paramName?: string;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getClusterParams.
 */
export interface GetClusterParamsResult {
    readonly clusterId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Instance parameter list. Note: This field may return null, indicating that no valid value can be obtained.
     */
    readonly items: outputs.Cynosdb.GetClusterParamsItem[];
    /**
     * The name of parameter.
     */
    readonly paramName?: string;
    readonly resultOutputFile?: string;
}

export function getClusterParamsOutput(args: GetClusterParamsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetClusterParamsResult> {
    return pulumi.output(args).apply(a => getClusterParams(a, opts))
}

/**
 * A collection of arguments for invoking getClusterParams.
 */
export interface GetClusterParamsOutputArgs {
    /**
     * The ID of cluster.
     */
    clusterId: pulumi.Input<string>;
    /**
     * Parameter name.
     */
    paramName?: pulumi.Input<string>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
}