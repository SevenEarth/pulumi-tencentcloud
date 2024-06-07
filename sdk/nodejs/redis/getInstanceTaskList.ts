// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of redis instanceTaskList
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const instanceTaskList = tencentcloud.Redis.getInstanceTaskList({
 *     beginTime: "2021-12-30 00:00:00",
 *     endTime: "2021-12-30 00:00:00",
 *     instanceId: "crs-c1nl9rpv",
 *     instanceName: "",
 *     operateUins: [""],
 *     projectIds: [""],
 *     results: [""],
 *     taskStatuses: [""],
 *     taskTypes: [""],
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getInstanceTaskList(args?: GetInstanceTaskListArgs, opts?: pulumi.InvokeOptions): Promise<GetInstanceTaskListResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Redis/getInstanceTaskList:getInstanceTaskList", {
        "beginTime": args.beginTime,
        "endTime": args.endTime,
        "instanceId": args.instanceId,
        "instanceName": args.instanceName,
        "operateUins": args.operateUins,
        "projectIds": args.projectIds,
        "resultOutputFile": args.resultOutputFile,
        "results": args.results,
        "taskStatuses": args.taskStatuses,
        "taskTypes": args.taskTypes,
    }, opts);
}

/**
 * A collection of arguments for invoking getInstanceTaskList.
 */
export interface GetInstanceTaskListArgs {
    /**
     * Start time.
     */
    beginTime?: string;
    /**
     * Termination time.
     */
    endTime?: string;
    /**
     * The ID of instance.
     */
    instanceId?: string;
    /**
     * Instance name.
     */
    instanceName?: string;
    /**
     * Operator Uin.
     */
    operateUins?: string[];
    /**
     * Project Id.
     */
    projectIds?: number[];
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
    /**
     * Task status.
     */
    results?: number[];
    /**
     * Task status.
     */
    taskStatuses?: number[];
    /**
     * Task type.
     */
    taskTypes?: string[];
}

/**
 * A collection of values returned by getInstanceTaskList.
 */
export interface GetInstanceTaskListResult {
    readonly beginTime?: string;
    /**
     * The end time.
     */
    readonly endTime?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The ID of instance.
     */
    readonly instanceId?: string;
    /**
     * The name of instance.
     */
    readonly instanceName?: string;
    readonly operateUins?: string[];
    readonly projectIds?: number[];
    readonly resultOutputFile?: string;
    /**
     * Task status.
     */
    readonly results?: number[];
    readonly taskStatuses?: number[];
    readonly taskTypes?: string[];
    /**
     * Task details.
     */
    readonly tasks: outputs.Redis.GetInstanceTaskListTask[];
}
/**
 * Use this data source to query detailed information of redis instanceTaskList
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const instanceTaskList = tencentcloud.Redis.getInstanceTaskList({
 *     beginTime: "2021-12-30 00:00:00",
 *     endTime: "2021-12-30 00:00:00",
 *     instanceId: "crs-c1nl9rpv",
 *     instanceName: "",
 *     operateUins: [""],
 *     projectIds: [""],
 *     results: [""],
 *     taskStatuses: [""],
 *     taskTypes: [""],
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getInstanceTaskListOutput(args?: GetInstanceTaskListOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetInstanceTaskListResult> {
    return pulumi.output(args).apply((a: any) => getInstanceTaskList(a, opts))
}

/**
 * A collection of arguments for invoking getInstanceTaskList.
 */
export interface GetInstanceTaskListOutputArgs {
    /**
     * Start time.
     */
    beginTime?: pulumi.Input<string>;
    /**
     * Termination time.
     */
    endTime?: pulumi.Input<string>;
    /**
     * The ID of instance.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * Instance name.
     */
    instanceName?: pulumi.Input<string>;
    /**
     * Operator Uin.
     */
    operateUins?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Project Id.
     */
    projectIds?: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
    /**
     * Task status.
     */
    results?: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * Task status.
     */
    taskStatuses?: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * Task type.
     */
    taskTypes?: pulumi.Input<pulumi.Input<string>[]>;
}
