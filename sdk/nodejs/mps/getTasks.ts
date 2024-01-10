// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of mps tasks
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const tasks = pulumi.output(tencentcloud.Mps.getTasks({
 *     limit: 20,
 *     status: "FINISH",
 * }));
 * ```
 */
export function getTasks(args: GetTasksArgs, opts?: pulumi.InvokeOptions): Promise<GetTasksResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:Mps/getTasks:getTasks", {
        "limit": args.limit,
        "resultOutputFile": args.resultOutputFile,
        "scrollToken": args.scrollToken,
        "status": args.status,
    }, opts);
}

/**
 * A collection of arguments for invoking getTasks.
 */
export interface GetTasksArgs {
    /**
     * Return the number of records, default value: 10, maximum value: 100.
     */
    limit?: number;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
    /**
     * Page turning flag, used when pulling in batches: when a single request cannot pull all the data, the interface will return a ScrollToken, and the next request will carry this Token, and it will be obtained from the next record.
     */
    scrollToken?: string;
    /**
     * Filter condition: task status, optional values: WAITING, PROCESSING, FINISH.
     */
    status: string;
}

/**
 * A collection of values returned by getTasks.
 */
export interface GetTasksResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly limit?: number;
    readonly resultOutputFile?: string;
    readonly scrollToken: string;
    readonly status: string;
    /**
     * Task list.
     */
    readonly taskSets: outputs.Mps.GetTasksTaskSet[];
}

export function getTasksOutput(args: GetTasksOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetTasksResult> {
    return pulumi.output(args).apply(a => getTasks(a, opts))
}

/**
 * A collection of arguments for invoking getTasks.
 */
export interface GetTasksOutputArgs {
    /**
     * Return the number of records, default value: 10, maximum value: 100.
     */
    limit?: pulumi.Input<number>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
    /**
     * Page turning flag, used when pulling in batches: when a single request cannot pull all the data, the interface will return a ScrollToken, and the next request will carry this Token, and it will be obtained from the next record.
     */
    scrollToken?: pulumi.Input<string>;
    /**
     * Filter condition: task status, optional values: WAITING, PROCESSING, FINISH.
     */
    status: pulumi.Input<string>;
}
