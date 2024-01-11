// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of bi userProject
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const userProject = pulumi.output(tencentcloud.Bi.getUserProject({
 *     allPage: true,
 *     projectId: 123,
 * }));
 * ```
 */
export function getUserProject(args?: GetUserProjectArgs, opts?: pulumi.InvokeOptions): Promise<GetUserProjectResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:Bi/getUserProject:getUserProject", {
        "allPage": args.allPage,
        "projectId": args.projectId,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getUserProject.
 */
export interface GetUserProjectArgs {
    /**
     * Whether to display all, if true, ignore paging.
     */
    allPage?: boolean;
    /**
     * Project id.
     */
    projectId?: number;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getUserProject.
 */
export interface GetUserProjectResult {
    readonly allPage?: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Array(Note: This field may return null, indicating that no valid value can be obtained).
     */
    readonly lists: outputs.Bi.GetUserProjectList[];
    readonly projectId?: number;
    readonly resultOutputFile?: string;
}

export function getUserProjectOutput(args?: GetUserProjectOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetUserProjectResult> {
    return pulumi.output(args).apply(a => getUserProject(a, opts))
}

/**
 * A collection of arguments for invoking getUserProject.
 */
export interface GetUserProjectOutputArgs {
    /**
     * Whether to display all, if true, ignore paging.
     */
    allPage?: pulumi.Input<boolean>;
    /**
     * Project id.
     */
    projectId?: pulumi.Input<number>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
}