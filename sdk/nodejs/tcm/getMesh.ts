// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of tcm mesh
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const mesh = pulumi.output(tencentcloud.Tcm.getMesh({
 *     meshClusters: ["cls-xxxx"],
 *     meshIds: ["mesh-xxxxxx"],
 *     meshNames: ["KEEP_MASH"],
 *     tags: ["key"],
 * }));
 * ```
 */
export function getMesh(args?: GetMeshArgs, opts?: pulumi.InvokeOptions): Promise<GetMeshResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:Tcm/getMesh:getMesh", {
        "meshClusters": args.meshClusters,
        "meshIds": args.meshIds,
        "meshNames": args.meshNames,
        "resultOutputFile": args.resultOutputFile,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getMesh.
 */
export interface GetMeshArgs {
    /**
     * Mesh name.
     */
    meshClusters?: string[];
    /**
     * Mesh instance Id.
     */
    meshIds?: string[];
    /**
     * Display name.
     */
    meshNames?: string[];
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
    /**
     * tag key.
     */
    tags?: string[];
}

/**
 * A collection of values returned by getMesh.
 */
export interface GetMeshResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly meshClusters?: string[];
    /**
     * Mesh instance Id.
     */
    readonly meshIds?: string[];
    /**
     * The mesh information is queriedNote: This field may return null, indicating that a valid value is not available.
     */
    readonly meshLists: outputs.Tcm.GetMeshMeshList[];
    readonly meshNames?: string[];
    readonly resultOutputFile?: string;
    readonly tags?: string[];
}

export function getMeshOutput(args?: GetMeshOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetMeshResult> {
    return pulumi.output(args).apply(a => getMesh(a, opts))
}

/**
 * A collection of arguments for invoking getMesh.
 */
export interface GetMeshOutputArgs {
    /**
     * Mesh name.
     */
    meshClusters?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Mesh instance Id.
     */
    meshIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Display name.
     */
    meshNames?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
    /**
     * tag key.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
}
