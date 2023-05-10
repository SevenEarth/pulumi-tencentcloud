// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of tdcpg instances.
 *
 * > **NOTE:** This data source is still in internal testing. To experience its functions, you need to apply for a whitelist from Tencent Cloud.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const instances = pulumi.output(tencentcloud.Tdcpg.getInstances({
 *     clusterId: "",
 *     instanceId: "",
 *     instanceName: "",
 *     instanceType: "",
 *     status: "",
 * }));
 * ```
 */
export function getInstances(args: GetInstancesArgs, opts?: pulumi.InvokeOptions): Promise<GetInstancesResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:Tdcpg/getInstances:getInstances", {
        "clusterId": args.clusterId,
        "instanceId": args.instanceId,
        "instanceName": args.instanceName,
        "instanceType": args.instanceType,
        "resultOutputFile": args.resultOutputFile,
        "status": args.status,
    }, opts);
}

/**
 * A collection of arguments for invoking getInstances.
 */
export interface GetInstancesArgs {
    /**
     * instance id.
     */
    clusterId: string;
    /**
     * instance id.
     */
    instanceId?: string;
    /**
     * instance name.
     */
    instanceName?: string;
    /**
     * instance type.
     */
    instanceType?: string;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
    /**
     * instance status.
     */
    status?: string;
}

/**
 * A collection of values returned by getInstances.
 */
export interface GetInstancesResult {
    /**
     * cluster id.
     */
    readonly clusterId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * instance id.
     */
    readonly instanceId?: string;
    /**
     * instance name.
     */
    readonly instanceName?: string;
    /**
     * instance type.
     */
    readonly instanceType?: string;
    /**
     * instance list.
     */
    readonly lists: outputs.Tdcpg.GetInstancesList[];
    readonly resultOutputFile?: string;
    /**
     * status.
     */
    readonly status?: string;
}

export function getInstancesOutput(args: GetInstancesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetInstancesResult> {
    return pulumi.output(args).apply(a => getInstances(a, opts))
}

/**
 * A collection of arguments for invoking getInstances.
 */
export interface GetInstancesOutputArgs {
    /**
     * instance id.
     */
    clusterId: pulumi.Input<string>;
    /**
     * instance id.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * instance name.
     */
    instanceName?: pulumi.Input<string>;
    /**
     * instance type.
     */
    instanceType?: pulumi.Input<string>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
    /**
     * instance status.
     */
    status?: pulumi.Input<string>;
}
