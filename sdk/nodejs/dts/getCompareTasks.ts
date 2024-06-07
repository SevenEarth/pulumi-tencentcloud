// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of dts compareTasks
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const compareTasks = tencentcloud.Dts.getCompareTasks({
 *     jobId: "",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getCompareTasks(args: GetCompareTasksArgs, opts?: pulumi.InvokeOptions): Promise<GetCompareTasksResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Dts/getCompareTasks:getCompareTasks", {
        "jobId": args.jobId,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getCompareTasks.
 */
export interface GetCompareTasksArgs {
    /**
     * job id.
     */
    jobId: string;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getCompareTasks.
 */
export interface GetCompareTasksResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * job id.
     */
    readonly jobId: string;
    /**
     * compare task list.
     */
    readonly lists: outputs.Dts.GetCompareTasksList[];
    readonly resultOutputFile?: string;
}
/**
 * Use this data source to query detailed information of dts compareTasks
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const compareTasks = tencentcloud.Dts.getCompareTasks({
 *     jobId: "",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getCompareTasksOutput(args: GetCompareTasksOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetCompareTasksResult> {
    return pulumi.output(args).apply((a: any) => getCompareTasks(a, opts))
}

/**
 * A collection of arguments for invoking getCompareTasks.
 */
export interface GetCompareTasksOutputArgs {
    /**
     * job id.
     */
    jobId: pulumi.Input<string>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
}
