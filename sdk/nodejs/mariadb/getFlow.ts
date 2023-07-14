// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of mariadb flow
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const flow = pulumi.output(tencentcloud.Mariadb.getFlow({
 *     flowId: 1307,
 * }));
 * ```
 */
export function getFlow(args: GetFlowArgs, opts?: pulumi.InvokeOptions): Promise<GetFlowResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:Mariadb/getFlow:getFlow", {
        "flowId": args.flowId,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getFlow.
 */
export interface GetFlowArgs {
    /**
     * Flow ID returned by async request API.
     */
    flowId: number;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getFlow.
 */
export interface GetFlowResult {
    readonly flowId: number;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly resultOutputFile?: string;
    /**
     * Flow status. 0: succeeded, 1: failed, 2: running.
     */
    readonly status: number;
}

export function getFlowOutput(args: GetFlowOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFlowResult> {
    return pulumi.output(args).apply(a => getFlow(a, opts))
}

/**
 * A collection of arguments for invoking getFlow.
 */
export interface GetFlowOutputArgs {
    /**
     * Flow ID returned by async request API.
     */
    flowId: pulumi.Input<number>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
}
