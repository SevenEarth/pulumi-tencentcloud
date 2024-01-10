// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of oceanus treeResources
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const example = pulumi.output(tencentcloud.Oceanus.getTreeResources({
 *     workSpaceId: "space-2idq8wbr",
 * }));
 * ```
 */
export function getTreeResources(args: GetTreeResourcesArgs, opts?: pulumi.InvokeOptions): Promise<GetTreeResourcesResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:Oceanus/getTreeResources:getTreeResources", {
        "resultOutputFile": args.resultOutputFile,
        "workSpaceId": args.workSpaceId,
    }, opts);
}

/**
 * A collection of arguments for invoking getTreeResources.
 */
export interface GetTreeResourcesArgs {
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
    /**
     * Workspace SerialId.
     */
    workSpaceId: string;
}

/**
 * A collection of values returned by getTreeResources.
 */
export interface GetTreeResourcesResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly resultOutputFile?: string;
    /**
     * Tree structure information.
     */
    readonly treeInfos: outputs.Oceanus.GetTreeResourcesTreeInfo[];
    readonly workSpaceId: string;
}

export function getTreeResourcesOutput(args: GetTreeResourcesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetTreeResourcesResult> {
    return pulumi.output(args).apply(a => getTreeResources(a, opts))
}

/**
 * A collection of arguments for invoking getTreeResources.
 */
export interface GetTreeResourcesOutputArgs {
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
    /**
     * Workspace SerialId.
     */
    workSpaceId: pulumi.Input<string>;
}
