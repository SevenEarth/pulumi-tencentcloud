// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query the COS batch.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const cosBatchs = tencentcloud.Cos.getBatchs({
 *     appid: "xxxxxx",
 *     uin: "xxxxxx",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getBatchs(args: GetBatchsArgs, opts?: pulumi.InvokeOptions): Promise<GetBatchsResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Cos/getBatchs:getBatchs", {
        "appid": args.appid,
        "jobStatuses": args.jobStatuses,
        "resultOutputFile": args.resultOutputFile,
        "uin": args.uin,
    }, opts);
}

/**
 * A collection of arguments for invoking getBatchs.
 */
export interface GetBatchsArgs {
    /**
     * Appid.
     */
    appid: number;
    /**
     * The task status information you need to query. If you do not specify a task status, COS returns the status of all tasks that have been executed, including those that are in progress. If you specify a task status, COS returns the task in the specified state. Optional task states include: Active, Cancelled, Cancelling, Complete, Completing, Failed, Failing, New, Paused, Pausing, Preparing, Ready, Suspended.
     */
    jobStatuses?: string;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
    /**
     * Uin.
     */
    uin: string;
}

/**
 * A collection of values returned by getBatchs.
 */
export interface GetBatchsResult {
    readonly appid: number;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly jobStatuses?: string;
    /**
     * Multiple batch processing task information.
     */
    readonly jobs: outputs.Cos.GetBatchsJob[];
    readonly resultOutputFile?: string;
    readonly uin: string;
}
/**
 * Use this data source to query the COS batch.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const cosBatchs = tencentcloud.Cos.getBatchs({
 *     appid: "xxxxxx",
 *     uin: "xxxxxx",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getBatchsOutput(args: GetBatchsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetBatchsResult> {
    return pulumi.output(args).apply((a: any) => getBatchs(a, opts))
}

/**
 * A collection of arguments for invoking getBatchs.
 */
export interface GetBatchsOutputArgs {
    /**
     * Appid.
     */
    appid: pulumi.Input<number>;
    /**
     * The task status information you need to query. If you do not specify a task status, COS returns the status of all tasks that have been executed, including those that are in progress. If you specify a task status, COS returns the task in the specified state. Optional task states include: Active, Cancelled, Cancelling, Complete, Completing, Failed, Failing, New, Paused, Pausing, Preparing, Ready, Suspended.
     */
    jobStatuses?: pulumi.Input<string>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
    /**
     * Uin.
     */
    uin: pulumi.Input<string>;
}
