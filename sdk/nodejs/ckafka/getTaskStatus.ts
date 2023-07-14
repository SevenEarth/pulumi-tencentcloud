// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of ckafka taskStatus
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const taskStatus = pulumi.output(tencentcloud.Ckafka.getTaskStatus({
 *     flowId: 123456,
 * }));
 * ```
 */
export function getTaskStatus(args: GetTaskStatusArgs, opts?: pulumi.InvokeOptions): Promise<GetTaskStatusResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
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

export function getTaskStatusOutput(args: GetTaskStatusOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetTaskStatusResult> {
    return pulumi.output(args).apply(a => getTaskStatus(a, opts))
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
