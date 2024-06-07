// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of oceanus savepointList
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const example = tencentcloud.Oceanus.getSavepointList({
 *     jobId: "cql-314rw6w0",
 *     workSpaceId: "space-2idq8wbr",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getSavepointList(args: GetSavepointListArgs, opts?: pulumi.InvokeOptions): Promise<GetSavepointListResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Oceanus/getSavepointList:getSavepointList", {
        "jobId": args.jobId,
        "resultOutputFile": args.resultOutputFile,
        "workSpaceId": args.workSpaceId,
    }, opts);
}

/**
 * A collection of arguments for invoking getSavepointList.
 */
export interface GetSavepointListArgs {
    /**
     * Job SerialId.
     */
    jobId: string;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
    /**
     * Workspace SerialId.
     */
    workSpaceId?: string;
}

/**
 * A collection of values returned by getSavepointList.
 */
export interface GetSavepointListResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly jobId: string;
    readonly resultOutputFile?: string;
    /**
     * Snapshot listNote: This field may return null, indicating that no valid value was found.
     */
    readonly savepoints: outputs.Oceanus.GetSavepointListSavepoint[];
    readonly workSpaceId?: string;
}
/**
 * Use this data source to query detailed information of oceanus savepointList
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const example = tencentcloud.Oceanus.getSavepointList({
 *     jobId: "cql-314rw6w0",
 *     workSpaceId: "space-2idq8wbr",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getSavepointListOutput(args: GetSavepointListOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSavepointListResult> {
    return pulumi.output(args).apply((a: any) => getSavepointList(a, opts))
}

/**
 * A collection of arguments for invoking getSavepointList.
 */
export interface GetSavepointListOutputArgs {
    /**
     * Job SerialId.
     */
    jobId: pulumi.Input<string>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
    /**
     * Workspace SerialId.
     */
    workSpaceId?: pulumi.Input<string>;
}
