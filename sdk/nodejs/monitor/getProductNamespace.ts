// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query product namespace in monitor)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const instances = pulumi.output(tencentcloud.Monitor.getProductNamespace({
 *     name: "Redis",
 * }));
 * ```
 */
export function getProductNamespace(args?: GetProductNamespaceArgs, opts?: pulumi.InvokeOptions): Promise<GetProductNamespaceResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:Monitor/getProductNamespace:getProductNamespace", {
        "name": args.name,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getProductNamespace.
 */
export interface GetProductNamespaceArgs {
    /**
     * Name for filter, eg:`Load Banlancer`.
     */
    name?: string;
    /**
     * Used to store results.
     */
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getProductNamespace.
 */
export interface GetProductNamespaceResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A list product namespaces. Each element contains the following attributes:
     */
    readonly lists: outputs.Monitor.GetProductNamespaceList[];
    readonly name?: string;
    readonly resultOutputFile?: string;
}

export function getProductNamespaceOutput(args?: GetProductNamespaceOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetProductNamespaceResult> {
    return pulumi.output(args).apply(a => getProductNamespace(a, opts))
}

/**
 * A collection of arguments for invoking getProductNamespace.
 */
export interface GetProductNamespaceOutputArgs {
    /**
     * Name for filter, eg:`Load Banlancer`.
     */
    name?: pulumi.Input<string>;
    /**
     * Used to store results.
     */
    resultOutputFile?: pulumi.Input<string>;
}
