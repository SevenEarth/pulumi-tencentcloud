// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of chdfs fileSystems
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const fileSystems = pulumi.output(tencentcloud.Chdfs.getFileSystems());
 * ```
 */
export function getFileSystems(args?: GetFileSystemsArgs, opts?: pulumi.InvokeOptions): Promise<GetFileSystemsResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:Chdfs/getFileSystems:getFileSystems", {
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getFileSystems.
 */
export interface GetFileSystemsArgs {
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getFileSystems.
 */
export interface GetFileSystemsResult {
    /**
     * file system list.
     */
    readonly fileSystems: outputs.Chdfs.GetFileSystemsFileSystem[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly resultOutputFile?: string;
}

export function getFileSystemsOutput(args?: GetFileSystemsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFileSystemsResult> {
    return pulumi.output(args).apply(a => getFileSystems(a, opts))
}

/**
 * A collection of arguments for invoking getFileSystems.
 */
export interface GetFileSystemsOutputArgs {
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
}