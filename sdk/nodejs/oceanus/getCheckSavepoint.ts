// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of oceanus checkSavepoint
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const example = tencentcloud.Oceanus.getCheckSavepoint({
 *     jobId: "cql-314rw6w0",
 *     recordType: 1,
 *     savepointPath: "cosn://52xkpymp-12345/12345/10000/cql-12345/2/flink-savepoints/savepoint-000000-12334",
 *     serialId: "svp-52xkpymp",
 *     workSpaceId: "space-2idq8wbr",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getCheckSavepoint(args: GetCheckSavepointArgs, opts?: pulumi.InvokeOptions): Promise<GetCheckSavepointResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Oceanus/getCheckSavepoint:getCheckSavepoint", {
        "jobId": args.jobId,
        "recordType": args.recordType,
        "resultOutputFile": args.resultOutputFile,
        "savepointPath": args.savepointPath,
        "serialId": args.serialId,
        "workSpaceId": args.workSpaceId,
    }, opts);
}

/**
 * A collection of arguments for invoking getCheckSavepoint.
 */
export interface GetCheckSavepointArgs {
    /**
     * Job id.
     */
    jobId: string;
    /**
     * Snapshot type. 1:savepoint; 2:checkpoint; 3:cancelWithSavepoint.
     */
    recordType: number;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
    /**
     * Snapshot path, currently only supports COS path.
     */
    savepointPath: string;
    /**
     * Snapshot resource ID.
     */
    serialId: string;
    /**
     * Workspace ID.
     */
    workSpaceId: string;
}

/**
 * A collection of values returned by getCheckSavepoint.
 */
export interface GetCheckSavepointResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly jobId: string;
    readonly recordType: number;
    readonly resultOutputFile?: string;
    readonly savepointPath: string;
    /**
     * 1=available, 2=unavailable.
     */
    readonly savepointStatus: number;
    readonly serialId: string;
    readonly workSpaceId: string;
}
/**
 * Use this data source to query detailed information of oceanus checkSavepoint
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const example = tencentcloud.Oceanus.getCheckSavepoint({
 *     jobId: "cql-314rw6w0",
 *     recordType: 1,
 *     savepointPath: "cosn://52xkpymp-12345/12345/10000/cql-12345/2/flink-savepoints/savepoint-000000-12334",
 *     serialId: "svp-52xkpymp",
 *     workSpaceId: "space-2idq8wbr",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getCheckSavepointOutput(args: GetCheckSavepointOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetCheckSavepointResult> {
    return pulumi.output(args).apply((a: any) => getCheckSavepoint(a, opts))
}

/**
 * A collection of arguments for invoking getCheckSavepoint.
 */
export interface GetCheckSavepointOutputArgs {
    /**
     * Job id.
     */
    jobId: pulumi.Input<string>;
    /**
     * Snapshot type. 1:savepoint; 2:checkpoint; 3:cancelWithSavepoint.
     */
    recordType: pulumi.Input<number>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
    /**
     * Snapshot path, currently only supports COS path.
     */
    savepointPath: pulumi.Input<string>;
    /**
     * Snapshot resource ID.
     */
    serialId: pulumi.Input<string>;
    /**
     * Workspace ID.
     */
    workSpaceId: pulumi.Input<string>;
}
