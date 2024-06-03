// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query dayu DDoS policies
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const idTest = tencentcloud.Dayu.getDdosPolicies({
 *     resourceType: tencentcloud_dayu_ddos_policy.test_policy.resource_type,
 *     policyId: tencentcloud_dayu_ddos_policy.test_policy.policy_id,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getDdosPolicies(args: GetDdosPoliciesArgs, opts?: pulumi.InvokeOptions): Promise<GetDdosPoliciesResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Dayu/getDdosPolicies:getDdosPolicies", {
        "policyId": args.policyId,
        "resourceType": args.resourceType,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getDdosPolicies.
 */
export interface GetDdosPoliciesArgs {
    /**
     * ID of the DDoS policy to be query.
     */
    policyId?: string;
    /**
     * Type of the resource that the DDoS policy works for, valid values are `bgpip`, `bgp`, `bgp-multip` and `net`.
     */
    resourceType: string;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getDdosPolicies.
 */
export interface GetDdosPoliciesResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A list of DDoS policies. Each element contains the following attributes:
     */
    readonly lists: outputs.Dayu.GetDdosPoliciesList[];
    /**
     * Id of policy.
     */
    readonly policyId?: string;
    readonly resourceType: string;
    readonly resultOutputFile?: string;
}
/**
 * Use this data source to query dayu DDoS policies
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const idTest = tencentcloud.Dayu.getDdosPolicies({
 *     resourceType: tencentcloud_dayu_ddos_policy.test_policy.resource_type,
 *     policyId: tencentcloud_dayu_ddos_policy.test_policy.policy_id,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getDdosPoliciesOutput(args: GetDdosPoliciesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDdosPoliciesResult> {
    return pulumi.output(args).apply((a: any) => getDdosPolicies(a, opts))
}

/**
 * A collection of arguments for invoking getDdosPolicies.
 */
export interface GetDdosPoliciesOutputArgs {
    /**
     * ID of the DDoS policy to be query.
     */
    policyId?: pulumi.Input<string>;
    /**
     * Type of the resource that the DDoS policy works for, valid values are `bgpip`, `bgp`, `bgp-multip` and `net`.
     */
    resourceType: pulumi.Input<string>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
}
