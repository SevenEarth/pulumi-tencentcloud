// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of rum groupLog
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const groupLog = tencentcloud.Rum.getGroupLog({
 *     endTime: "1625454840000",
 *     groupField: "level",
 *     orderBy: "desc",
 *     projectId: 1,
 *     query: "id:123 AND type:\"log\"",
 *     startTime: "1625444040000",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getGroupLog(args: GetGroupLogArgs, opts?: pulumi.InvokeOptions): Promise<GetGroupLogResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Rum/getGroupLog:getGroupLog", {
        "endTime": args.endTime,
        "groupField": args.groupField,
        "orderBy": args.orderBy,
        "projectId": args.projectId,
        "query": args.query,
        "resultOutputFile": args.resultOutputFile,
        "startTime": args.startTime,
    }, opts);
}

/**
 * A collection of arguments for invoking getGroupLog.
 */
export interface GetGroupLogArgs {
    /**
     * End time but is represented using a timestamp in milliseconds.
     */
    endTime: string;
    /**
     * The field used for group.
     */
    groupField: string;
    /**
     * Sorting method. `desc`:Descending order; `asc`: Ascending order.
     */
    orderBy: string;
    /**
     * Project ID.
     */
    projectId: number;
    /**
     * Log Query syntax statement.
     */
    query: string;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
    /**
     * Start time but is represented using a timestamp in milliseconds.
     */
    startTime: string;
}

/**
 * A collection of values returned by getGroupLog.
 */
export interface GetGroupLogResult {
    readonly endTime: string;
    readonly groupField: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly orderBy: string;
    readonly projectId: number;
    readonly query: string;
    /**
     * Return value.
     */
    readonly result: string;
    readonly resultOutputFile?: string;
    readonly startTime: string;
}
/**
 * Use this data source to query detailed information of rum groupLog
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const groupLog = tencentcloud.Rum.getGroupLog({
 *     endTime: "1625454840000",
 *     groupField: "level",
 *     orderBy: "desc",
 *     projectId: 1,
 *     query: "id:123 AND type:\"log\"",
 *     startTime: "1625444040000",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getGroupLogOutput(args: GetGroupLogOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetGroupLogResult> {
    return pulumi.output(args).apply((a: any) => getGroupLog(a, opts))
}

/**
 * A collection of arguments for invoking getGroupLog.
 */
export interface GetGroupLogOutputArgs {
    /**
     * End time but is represented using a timestamp in milliseconds.
     */
    endTime: pulumi.Input<string>;
    /**
     * The field used for group.
     */
    groupField: pulumi.Input<string>;
    /**
     * Sorting method. `desc`:Descending order; `asc`: Ascending order.
     */
    orderBy: pulumi.Input<string>;
    /**
     * Project ID.
     */
    projectId: pulumi.Input<number>;
    /**
     * Log Query syntax statement.
     */
    query: pulumi.Input<string>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
    /**
     * Start time but is represented using a timestamp in milliseconds.
     */
    startTime: pulumi.Input<string>;
}
