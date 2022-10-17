// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query API gateway usage plans.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@tencentcloud_iac/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const plan = new tencentcloud.apigateway.UsagePlan("plan", {
 *     usagePlanName: "my_plan",
 *     usagePlanDesc: "nice plan",
 *     maxRequestNum: 100,
 *     maxRequestNumPreSec: 10,
 * });
 * const name = tencentcloud.ApiGateway.getUsagePlansOutput({
 *     usagePlanName: plan.usagePlanName,
 * });
 * const id = tencentcloud.ApiGateway.getUsagePlansOutput({
 *     usagePlanId: plan.id,
 * });
 * ```
 */
export function getUsagePlans(args?: GetUsagePlansArgs, opts?: pulumi.InvokeOptions): Promise<GetUsagePlansResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:ApiGateway/getUsagePlans:getUsagePlans", {
        "resultOutputFile": args.resultOutputFile,
        "usagePlanId": args.usagePlanId,
        "usagePlanName": args.usagePlanName,
    }, opts);
}

/**
 * A collection of arguments for invoking getUsagePlans.
 */
export interface GetUsagePlansArgs {
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
    /**
     * ID of the usage plan.
     */
    usagePlanId?: string;
    /**
     * Name of the usage plan.
     */
    usagePlanName?: string;
}

/**
 * A collection of values returned by getUsagePlans.
 */
export interface GetUsagePlansResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A list of usage plans.
     */
    readonly lists: outputs.ApiGateway.GetUsagePlansList[];
    readonly resultOutputFile?: string;
    /**
     * ID of the usage plan.
     */
    readonly usagePlanId?: string;
    /**
     * Name of the usage plan.
     */
    readonly usagePlanName?: string;
}

export function getUsagePlansOutput(args?: GetUsagePlansOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetUsagePlansResult> {
    return pulumi.output(args).apply(a => getUsagePlans(a, opts))
}

/**
 * A collection of arguments for invoking getUsagePlans.
 */
export interface GetUsagePlansOutputArgs {
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
    /**
     * ID of the usage plan.
     */
    usagePlanId?: pulumi.Input<string>;
    /**
     * Name of the usage plan.
     */
    usagePlanName?: pulumi.Input<string>;
}
