// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of scaling policy.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const asScalingPolicies = pulumi.output(tencentcloud.As.getScalingPolicies({
 *     resultOutputFile: "mytestpath",
 *     scalingPolicyId: "asg-mvyghxu7",
 * }));
 * ```
 */
export function getScalingPolicies(args?: GetScalingPoliciesArgs, opts?: pulumi.InvokeOptions): Promise<GetScalingPoliciesResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:As/getScalingPolicies:getScalingPolicies", {
        "policyName": args.policyName,
        "resultOutputFile": args.resultOutputFile,
        "scalingGroupId": args.scalingGroupId,
        "scalingPolicyId": args.scalingPolicyId,
    }, opts);
}

/**
 * A collection of arguments for invoking getScalingPolicies.
 */
export interface GetScalingPoliciesArgs {
    /**
     * Scaling policy name.
     */
    policyName?: string;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
    /**
     * Scaling group ID.
     */
    scalingGroupId?: string;
    /**
     * Scaling policy ID.
     */
    scalingPolicyId?: string;
}

/**
 * A collection of values returned by getScalingPolicies.
 */
export interface GetScalingPoliciesResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Scaling policy name.
     */
    readonly policyName?: string;
    readonly resultOutputFile?: string;
    /**
     * Scaling policy ID.
     */
    readonly scalingGroupId?: string;
    readonly scalingPolicyId?: string;
    /**
     * A list of scaling policy. Each element contains the following attributes:
     */
    readonly scalingPolicyLists: outputs.As.GetScalingPoliciesScalingPolicyList[];
}

export function getScalingPoliciesOutput(args?: GetScalingPoliciesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetScalingPoliciesResult> {
    return pulumi.output(args).apply(a => getScalingPolicies(a, opts))
}

/**
 * A collection of arguments for invoking getScalingPolicies.
 */
export interface GetScalingPoliciesOutputArgs {
    /**
     * Scaling policy name.
     */
    policyName?: pulumi.Input<string>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
    /**
     * Scaling group ID.
     */
    scalingGroupId?: pulumi.Input<string>;
    /**
     * Scaling policy ID.
     */
    scalingPolicyId?: pulumi.Input<string>;
}
