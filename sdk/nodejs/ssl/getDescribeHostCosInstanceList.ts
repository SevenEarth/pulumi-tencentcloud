// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of ssl describeHostCosInstanceList
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const describeHostCosInstanceList = pulumi.output(tencentcloud.Ssl.getDescribeHostCosInstanceList({
 *     certificateId: "8u8DII0l",
 *     resourceType: "cos",
 * }));
 * ```
 */
export function getDescribeHostCosInstanceList(args: GetDescribeHostCosInstanceListArgs, opts?: pulumi.InvokeOptions): Promise<GetDescribeHostCosInstanceListResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:Ssl/getDescribeHostCosInstanceList:getDescribeHostCosInstanceList", {
        "certificateId": args.certificateId,
        "filters": args.filters,
        "isCache": args.isCache,
        "resourceType": args.resourceType,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getDescribeHostCosInstanceList.
 */
export interface GetDescribeHostCosInstanceListArgs {
    /**
     * Certificate ID to be deployed.
     */
    certificateId: string;
    /**
     * List of filter parameters.
     */
    filters?: inputs.Ssl.GetDescribeHostCosInstanceListFilter[];
    /**
     * Whether to query the cache, 1: Yes; 0: No, the default is the query cache, the cache is half an hour.
     */
    isCache?: number;
    /**
     * Deploy resource type cos.
     */
    resourceType: string;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getDescribeHostCosInstanceList.
 */
export interface GetDescribeHostCosInstanceListResult {
    /**
     * Current cache read timeNote: This field may return NULL, indicating that the valid value cannot be obtained.
     */
    readonly asyncCacheTime: string;
    /**
     * Asynchronous refresh current execution numberNote: This field may return NULL, indicating that the valid value cannot be obtained.
     */
    readonly asyncOffset: number;
    /**
     * The total number of asynchronous refreshNote: This field may return NULL, indicating that the valid value cannot be obtained.
     */
    readonly asyncTotalNum: number;
    readonly certificateId: string;
    readonly filters?: outputs.Ssl.GetDescribeHostCosInstanceListFilter[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * COS instance listNote: This field may return NULL, indicating that the valid value cannot be obtained.
     */
    readonly instanceLists: outputs.Ssl.GetDescribeHostCosInstanceListInstanceList[];
    readonly isCache?: number;
    readonly resourceType: string;
    readonly resultOutputFile?: string;
}

export function getDescribeHostCosInstanceListOutput(args: GetDescribeHostCosInstanceListOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDescribeHostCosInstanceListResult> {
    return pulumi.output(args).apply(a => getDescribeHostCosInstanceList(a, opts))
}

/**
 * A collection of arguments for invoking getDescribeHostCosInstanceList.
 */
export interface GetDescribeHostCosInstanceListOutputArgs {
    /**
     * Certificate ID to be deployed.
     */
    certificateId: pulumi.Input<string>;
    /**
     * List of filter parameters.
     */
    filters?: pulumi.Input<pulumi.Input<inputs.Ssl.GetDescribeHostCosInstanceListFilterArgs>[]>;
    /**
     * Whether to query the cache, 1: Yes; 0: No, the default is the query cache, the cache is half an hour.
     */
    isCache?: pulumi.Input<number>;
    /**
     * Deploy resource type cos.
     */
    resourceType: pulumi.Input<string>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
}