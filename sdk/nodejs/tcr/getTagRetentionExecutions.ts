// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of tcr tagRetentionExecutions
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const tagRetentionExecutions = pulumi.output(tencentcloud.Tcr.getTagRetentionExecutions({
 *     registryId: "tcr_ins_id",
 *     retentionId: 1,
 * }));
 * ```
 */
export function getTagRetentionExecutions(args: GetTagRetentionExecutionsArgs, opts?: pulumi.InvokeOptions): Promise<GetTagRetentionExecutionsResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:Tcr/getTagRetentionExecutions:getTagRetentionExecutions", {
        "registryId": args.registryId,
        "resultOutputFile": args.resultOutputFile,
        "retentionId": args.retentionId,
    }, opts);
}

/**
 * A collection of arguments for invoking getTagRetentionExecutions.
 */
export interface GetTagRetentionExecutionsArgs {
    /**
     * instance id.
     */
    registryId: string;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
    /**
     * retention id.
     */
    retentionId: number;
}

/**
 * A collection of values returned by getTagRetentionExecutions.
 */
export interface GetTagRetentionExecutionsResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly registryId: string;
    readonly resultOutputFile?: string;
    /**
     * list of version retention execution records.
     */
    readonly retentionExecutionLists: outputs.Tcr.GetTagRetentionExecutionsRetentionExecutionList[];
    /**
     * retention id.
     */
    readonly retentionId: number;
}

export function getTagRetentionExecutionsOutput(args: GetTagRetentionExecutionsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetTagRetentionExecutionsResult> {
    return pulumi.output(args).apply(a => getTagRetentionExecutions(a, opts))
}

/**
 * A collection of arguments for invoking getTagRetentionExecutions.
 */
export interface GetTagRetentionExecutionsOutputArgs {
    /**
     * instance id.
     */
    registryId: pulumi.Input<string>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
    /**
     * retention id.
     */
    retentionId: pulumi.Input<number>;
}