// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of tat agent
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const agent = pulumi.output(tencentcloud.Tat.getAgent({
 *     // instance_ids = ["ins-f9jr4bd2"]
 *     filters: [{
 *         name: "environment",
 *         values: ["Linux"],
 *     }],
 * }));
 * ```
 */
export function getAgent(args?: GetAgentArgs, opts?: pulumi.InvokeOptions): Promise<GetAgentResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:Tat/getAgent:getAgent", {
        "filters": args.filters,
        "instanceIds": args.instanceIds,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getAgent.
 */
export interface GetAgentArgs {
    /**
     * Filter conditions. agent-status - String - Required: No - (Filter condition) Filter by agent status. Valid values: Online, Offline. environment - String - Required: No - (Filter condition) Filter by the agent environment. Valid value: Linux. instance-id - String - Required: No - (Filter condition) Filter by the instance ID. Up to 10 Filters allowed in one request. For each filter, five Filter.Values can be specified. InstanceIds and Filters cannot be specified at the same time.
     */
    filters?: inputs.Tat.GetAgentFilter[];
    /**
     * List of instance IDs for the query.
     */
    instanceIds?: string[];
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getAgent.
 */
export interface GetAgentResult {
    /**
     * List of agent message.
     */
    readonly automationAgentSets: outputs.Tat.GetAgentAutomationAgentSet[];
    readonly filters?: outputs.Tat.GetAgentFilter[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceIds?: string[];
    readonly resultOutputFile?: string;
}

export function getAgentOutput(args?: GetAgentOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAgentResult> {
    return pulumi.output(args).apply(a => getAgent(a, opts))
}

/**
 * A collection of arguments for invoking getAgent.
 */
export interface GetAgentOutputArgs {
    /**
     * Filter conditions. agent-status - String - Required: No - (Filter condition) Filter by agent status. Valid values: Online, Offline. environment - String - Required: No - (Filter condition) Filter by the agent environment. Valid value: Linux. instance-id - String - Required: No - (Filter condition) Filter by the instance ID. Up to 10 Filters allowed in one request. For each filter, five Filter.Values can be specified. InstanceIds and Filters cannot be specified at the same time.
     */
    filters?: pulumi.Input<pulumi.Input<inputs.Tat.GetAgentFilterArgs>[]>;
    /**
     * List of instance IDs for the query.
     */
    instanceIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
}
