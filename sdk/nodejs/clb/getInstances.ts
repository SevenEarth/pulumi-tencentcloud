// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of CLB
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const foo = pulumi.output(tencentcloud.Clb.getInstances({
 *     clbId: "lb-k2zjp9lv",
 *     clbName: "myclb",
 *     networkType: "OPEN",
 *     projectId: 0,
 *     resultOutputFile: "mytestpath",
 * }));
 * ```
 */
export function getInstances(args?: GetInstancesArgs, opts?: pulumi.InvokeOptions): Promise<GetInstancesResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:Clb/getInstances:getInstances", {
        "clbId": args.clbId,
        "clbName": args.clbName,
        "masterZone": args.masterZone,
        "networkType": args.networkType,
        "projectId": args.projectId,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getInstances.
 */
export interface GetInstancesArgs {
    /**
     * ID of the CLB to be queried.
     */
    clbId?: string;
    /**
     * Name of the CLB to be queried.
     */
    clbName?: string;
    /**
     * Master available zone id.
     */
    masterZone?: string;
    /**
     * Type of CLB instance, and available values include `OPEN` and `INTERNAL`.
     */
    networkType?: string;
    /**
     * Project ID of the CLB.
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
     * ID of CLB.
     */
    readonly clbId?: string;
    /**
     * A list of cloud load balancers. Each element contains the following attributes:
     */
    readonly clbLists: outputs.Clb.GetInstancesClbList[];
    /**
     * Name of CLB.
     */
    readonly clbName?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly masterZone?: string;
    /**
     * Types of CLB.
     */
    readonly networkType?: string;
    /**
     * ID of the project.
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
     * ID of the CLB to be queried.
     */
    clbId?: pulumi.Input<string>;
    /**
     * Name of the CLB to be queried.
     */
    clbName?: pulumi.Input<string>;
    /**
     * Master available zone id.
     */
    masterZone?: pulumi.Input<string>;
    /**
     * Type of CLB instance, and available values include `OPEN` and `INTERNAL`.
     */
    networkType?: pulumi.Input<string>;
    /**
     * Project ID of the CLB.
     */
    projectId?: pulumi.Input<number>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
}
