// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of vpc classicLinkInstances
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const classicLinkInstances = tencentcloud.Vpc.getClassicLinkInstances({
 *     filters: [{
 *         name: "vpc-id",
 *         values: ["vpc-lh4nqig9"],
 *     }],
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getClassicLinkInstances(args?: GetClassicLinkInstancesArgs, opts?: pulumi.InvokeOptions): Promise<GetClassicLinkInstancesResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Vpc/getClassicLinkInstances:getClassicLinkInstances", {
        "filters": args.filters,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getClassicLinkInstances.
 */
export interface GetClassicLinkInstancesArgs {
    /**
     * Filter conditions.`vpc-id` - String - (Filter condition) The VPC instance ID. `vm-ip` - String - (Filter condition) The IP address of the CVM on the basic network.
     */
    filters?: inputs.Vpc.GetClassicLinkInstancesFilter[];
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getClassicLinkInstances.
 */
export interface GetClassicLinkInstancesResult {
    /**
     * Classiclink instance.
     */
    readonly classicLinkInstanceSets: outputs.Vpc.GetClassicLinkInstancesClassicLinkInstanceSet[];
    readonly filters?: outputs.Vpc.GetClassicLinkInstancesFilter[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly resultOutputFile?: string;
}
/**
 * Use this data source to query detailed information of vpc classicLinkInstances
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const classicLinkInstances = tencentcloud.Vpc.getClassicLinkInstances({
 *     filters: [{
 *         name: "vpc-id",
 *         values: ["vpc-lh4nqig9"],
 *     }],
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getClassicLinkInstancesOutput(args?: GetClassicLinkInstancesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetClassicLinkInstancesResult> {
    return pulumi.output(args).apply((a: any) => getClassicLinkInstances(a, opts))
}

/**
 * A collection of arguments for invoking getClassicLinkInstances.
 */
export interface GetClassicLinkInstancesOutputArgs {
    /**
     * Filter conditions.`vpc-id` - String - (Filter condition) The VPC instance ID. `vm-ip` - String - (Filter condition) The IP address of the CVM on the basic network.
     */
    filters?: pulumi.Input<pulumi.Input<inputs.Vpc.GetClassicLinkInstancesFilterArgs>[]>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
}
