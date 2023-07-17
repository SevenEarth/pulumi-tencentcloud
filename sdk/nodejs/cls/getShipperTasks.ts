// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of cls shipperTasks
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const shipperTasks = pulumi.output(tencentcloud.Cls.getShipperTasks({
 *     endTime: 160749910800,
 *     shipperId: "dbde3c9b-ea16-4032-bc2a-d8fa65567a8e",
 *     startTime: 160749910700,
 * }));
 * ```
 */
export function getShipperTasks(args: GetShipperTasksArgs, opts?: pulumi.InvokeOptions): Promise<GetShipperTasksResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:Cls/getShipperTasks:getShipperTasks", {
        "endTime": args.endTime,
        "resultOutputFile": args.resultOutputFile,
        "shipperId": args.shipperId,
        "startTime": args.startTime,
    }, opts);
}

/**
 * A collection of arguments for invoking getShipperTasks.
 */
export interface GetShipperTasksArgs {
    /**
     * end time(ms).
     */
    endTime: number;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
    /**
     * shipper id.
     */
    shipperId: string;
    /**
     * start time(ms).
     */
    startTime: number;
}

/**
 * A collection of values returned by getShipperTasks.
 */
export interface GetShipperTasksResult {
    /**
     * end time(ms).
     */
    readonly endTime: number;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly resultOutputFile?: string;
    /**
     * shipper id.
     */
    readonly shipperId: string;
    /**
     * start time(ms).
     */
    readonly startTime: number;
    /**
     * .
     */
    readonly tasks: outputs.Cls.GetShipperTasksTask[];
}

export function getShipperTasksOutput(args: GetShipperTasksOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetShipperTasksResult> {
    return pulumi.output(args).apply(a => getShipperTasks(a, opts))
}

/**
 * A collection of arguments for invoking getShipperTasks.
 */
export interface GetShipperTasksOutputArgs {
    /**
     * end time(ms).
     */
    endTime: pulumi.Input<number>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
    /**
     * shipper id.
     */
    shipperId: pulumi.Input<string>;
    /**
     * start time(ms).
     */
    startTime: pulumi.Input<number>;
}
