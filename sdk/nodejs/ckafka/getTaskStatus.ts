// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of ckafka taskStatus
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const taskStatus = tencentcloud.Ckafka.getTaskStatus({
 *     flowId: 123456,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getTaskStatus(args: GetTaskStatusArgs, opts?: pulumi.InvokeOptions): Promise<GetTaskStatusResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Ckafka/getTaskStatus:getTaskStatus", {
        "flowId": args.flowId,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getTaskStatus.
 */
export interface GetTaskStatusArgs {
    /**
     * FlowId.
     */
    flowId: number;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getTaskStatus.
 */
export interface GetTaskStatusResult {
    readonly flowId: number;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly resultOutputFile?: string;
    /**
     * Result.
     */
    readonly results: outputs.Ckafka.GetTaskStatusResult[];
}
/**
 * Use this data source to query detailed information of ckafka taskStatus
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const taskStatus = tencentcloud.Ckafka.getTaskStatus({
 *     flowId: 123456,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getTaskStatusOutput(args: GetTaskStatusOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetTaskStatusResult> {
    return pulumi.output(args).apply((a: any) => getTaskStatus(a, opts))
}

/**
 * A collection of arguments for invoking getTaskStatus.
 */
export interface GetTaskStatusOutputArgs {
    /**
     * FlowId.
     */
    flowId: pulumi.Input<number>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
}
