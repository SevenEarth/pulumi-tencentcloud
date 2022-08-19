// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query policy group binding objects.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const name = tencentcloud.Monitor.getPolicyGroups({
 *     name: "test",
 * });
 * const objects = name.then(name => tencentcloud.Monitor.getBindingObjects({
 *     groupId: name.lists?[0]?.groupId,
 * }));
 * ```
 */
export function getBindingObjects(args: GetBindingObjectsArgs, opts?: pulumi.InvokeOptions): Promise<GetBindingObjectsResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:Monitor/getBindingObjects:getBindingObjects", {
        "groupId": args.groupId,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getBindingObjects.
 */
export interface GetBindingObjectsArgs {
    /**
     * Policy group ID for query.
     */
    groupId: number;
    /**
     * Used to store results.
     */
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getBindingObjects.
 */
export interface GetBindingObjectsResult {
    readonly groupId: number;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A list objects. Each element contains the following attributes:
     */
    readonly lists: outputs.Monitor.GetBindingObjectsList[];
    readonly resultOutputFile?: string;
}

export function getBindingObjectsOutput(args: GetBindingObjectsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetBindingObjectsResult> {
    return pulumi.output(args).apply(a => getBindingObjects(a, opts))
}

/**
 * A collection of arguments for invoking getBindingObjects.
 */
export interface GetBindingObjectsOutputArgs {
    /**
     * Policy group ID for query.
     */
    groupId: pulumi.Input<number>;
    /**
     * Used to store results.
     */
    resultOutputFile?: pulumi.Input<string>;
}
