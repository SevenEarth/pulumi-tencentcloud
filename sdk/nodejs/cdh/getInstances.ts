// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query CDH instances.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const list = pulumi.output(tencentcloud.Cdh.getInstances({
 *     availabilityZone: "ap-guangzhou-3",
 *     hostId: "host-d6s7i5q4",
 *     hostName: "test",
 *     hostState: "RUNNING",
 *     projectId: 1154137,
 * }));
 * ```
 */
export function getInstances(args?: GetInstancesArgs, opts?: pulumi.InvokeOptions): Promise<GetInstancesResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:Cdh/getInstances:getInstances", {
        "availabilityZone": args.availabilityZone,
        "hostId": args.hostId,
        "hostName": args.hostName,
        "hostState": args.hostState,
        "projectId": args.projectId,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getInstances.
 */
export interface GetInstancesArgs {
    /**
     * The available zone that the CDH instance locates at.
     */
    availabilityZone?: string;
    /**
     * ID of the CDH instances to be queried.
     */
    hostId?: string;
    /**
     * Name of the CDH instances to be queried.
     */
    hostName?: string;
    /**
     * State of the CDH instances to be queried. Valid values: `PENDING`, `LAUNCH_FAILURE`, `RUNNING`, `EXPIRED`.
     */
    hostState?: string;
    /**
     * The project CDH belongs to.
     */
    projectId?: number;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getInstances.
 */
export interface GetInstancesResult {
    /**
     * The available zone that the CDH instance locates at.
     */
    readonly availabilityZone?: string;
    /**
     * An information list of cdh instance. Each element contains the following attributes:
     */
    readonly cdhInstanceLists: outputs.Cdh.GetInstancesCdhInstanceList[];
    /**
     * ID of the CDH instance.
     */
    readonly hostId?: string;
    /**
     * Name of the CDH instance.
     */
    readonly hostName?: string;
    /**
     * State of the CDH instance.
     */
    readonly hostState?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The project CDH belongs to.
     */
    readonly projectId?: number;
    readonly resultOutputFile?: string;
}

export function getInstancesOutput(args?: GetInstancesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetInstancesResult> {
    return pulumi.output(args).apply(a => getInstances(a, opts))
}

/**
 * A collection of arguments for invoking getInstances.
 */
export interface GetInstancesOutputArgs {
    /**
     * The available zone that the CDH instance locates at.
     */
    availabilityZone?: pulumi.Input<string>;
    /**
     * ID of the CDH instances to be queried.
     */
    hostId?: pulumi.Input<string>;
    /**
     * Name of the CDH instances to be queried.
     */
    hostName?: pulumi.Input<string>;
    /**
     * State of the CDH instances to be queried. Valid values: `PENDING`, `LAUNCH_FAILURE`, `RUNNING`, `EXPIRED`.
     */
    hostState?: pulumi.Input<string>;
    /**
     * The project CDH belongs to.
     */
    projectId?: pulumi.Input<number>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
}
