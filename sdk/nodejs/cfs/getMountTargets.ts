// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of cfs mountTargets
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const mountTargets = pulumi.output(tencentcloud.Cfs.getMountTargets({
 *     fileSystemId: "cfs-iobiaxtj",
 * }));
 * ```
 */
export function getMountTargets(args: GetMountTargetsArgs, opts?: pulumi.InvokeOptions): Promise<GetMountTargetsResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:Cfs/getMountTargets:getMountTargets", {
        "fileSystemId": args.fileSystemId,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getMountTargets.
 */
export interface GetMountTargetsArgs {
    /**
     * File system ID.
     */
    fileSystemId: string;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getMountTargets.
 */
export interface GetMountTargetsResult {
    /**
     * File system ID.
     */
    readonly fileSystemId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Mount target details.
     */
    readonly mountTargets: outputs.Cfs.GetMountTargetsMountTarget[];
    readonly resultOutputFile?: string;
}

export function getMountTargetsOutput(args: GetMountTargetsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetMountTargetsResult> {
    return pulumi.output(args).apply(a => getMountTargets(a, opts))
}

/**
 * A collection of arguments for invoking getMountTargets.
 */
export interface GetMountTargetsOutputArgs {
    /**
     * File system ID.
     */
    fileSystemId: pulumi.Input<string>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
}
