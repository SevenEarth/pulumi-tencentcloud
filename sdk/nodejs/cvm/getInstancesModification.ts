// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query cvm instances modification.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const foo = tencentcloud.Cvm.getInstancesModification({
 *     instanceIds: ["ins-xxxxxxx"],
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getInstancesModification(args?: GetInstancesModificationArgs, opts?: pulumi.InvokeOptions): Promise<GetInstancesModificationResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Cvm/getInstancesModification:getInstancesModification", {
        "filters": args.filters,
        "instanceIds": args.instanceIds,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getInstancesModification.
 */
export interface GetInstancesModificationArgs {
    /**
     * The upper limit of Filters for each request is 10 and the upper limit for Filter.Values is 2.
     */
    filters?: inputs.Cvm.GetInstancesModificationFilter[];
    /**
     * One or more instance ID to be queried. It can be obtained from the InstanceId in the returned value of API DescribeInstances. The maximum number of instances in batch for each request is 20.
     */
    instanceIds?: string[];
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getInstancesModification.
 */
export interface GetInstancesModificationResult {
    readonly filters?: outputs.Cvm.GetInstancesModificationFilter[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceIds?: string[];
    /**
     * The list of model configurations that can be adjusted by the instance.
     */
    readonly instanceTypeConfigStatusLists: outputs.Cvm.GetInstancesModificationInstanceTypeConfigStatusList[];
    readonly resultOutputFile?: string;
}
/**
 * Use this data source to query cvm instances modification.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const foo = tencentcloud.Cvm.getInstancesModification({
 *     instanceIds: ["ins-xxxxxxx"],
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getInstancesModificationOutput(args?: GetInstancesModificationOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetInstancesModificationResult> {
    return pulumi.output(args).apply((a: any) => getInstancesModification(a, opts))
}

/**
 * A collection of arguments for invoking getInstancesModification.
 */
export interface GetInstancesModificationOutputArgs {
    /**
     * The upper limit of Filters for each request is 10 and the upper limit for Filter.Values is 2.
     */
    filters?: pulumi.Input<pulumi.Input<inputs.Cvm.GetInstancesModificationFilterArgs>[]>;
    /**
     * One or more instance ID to be queried. It can be obtained from the InstanceId in the returned value of API DescribeInstances. The maximum number of instances in batch for each request is 20.
     */
    instanceIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
}
