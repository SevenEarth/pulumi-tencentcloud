// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a wedata integrationOfflineTask
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const example = new tencentcloud.Wedata.IntegrationOfflineTask("example", {
 *     cycleStep: 1,
 *     delayTime: 0,
 *     endTime: "2099-12-31 00:00:00",
 *     notes: "terraform example demo.",
 *     projectId: "1612982498218618880",
 *     startTime: "2023-12-31 00:00:00",
 *     taskAction: "2",
 *     taskInfo: {
 *         configs: [
 *             {
 *                 name: "Args",
 *                 value: "args",
 *             },
 *             {
 *                 name: "dirtyDataThreshold",
 *                 value: "0",
 *             },
 *             {
 *                 name: "concurrency",
 *                 value: "1",
 *             },
 *             {
 *                 name: "syncRateLimitUnit",
 *                 value: "0",
 *             },
 *         ],
 *         executorId: "20230313175748567418",
 *         extConfigs: [{
 *             name: "TaskAlarmRegularList",
 *             value: "73",
 *         }],
 *         incharge: "demo",
 *         offlineTaskAddEntity: {
 *             crontabExpression: "0 0 1 * * ?",
 *             cycleType: 3,
 *             retriable: 1,
 *             retryWait: 5,
 *             selfDepend: 1,
 *             tryLimit: 5,
 *         },
 *     },
 *     taskMode: "1",
 *     taskName: "tf_example",
 * });
 * ```
 *
 * ## Import
 *
 * wedata integration_offline_task can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:Wedata/integrationOfflineTask:IntegrationOfflineTask example 1612982498218618880#20231102200955095
 * ```
 */
export class IntegrationOfflineTask extends pulumi.CustomResource {
    /**
     * Get an existing IntegrationOfflineTask resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IntegrationOfflineTaskState, opts?: pulumi.CustomResourceOptions): IntegrationOfflineTask {
        return new IntegrationOfflineTask(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Wedata/integrationOfflineTask:IntegrationOfflineTask';

    /**
     * Returns true if the given object is an instance of IntegrationOfflineTask.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IntegrationOfflineTask {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IntegrationOfflineTask.__pulumiType;
    }

    /**
     * Interval time of scheduling, the minimum value: 1.
     */
    public readonly cycleStep!: pulumi.Output<number>;
    /**
     * Execution time, unit is minutes, only available for day/week/month/year scheduling. For example, daily scheduling is executed once every day at 02:00, and the delayTime is 120 minutes.
     */
    public readonly delayTime!: pulumi.Output<number>;
    /**
     * Effective end time, the format is yyyy-MM-dd HH:mm:ss.
     */
    public readonly endTime!: pulumi.Output<string>;
    /**
     * Description information.
     */
    public readonly notes!: pulumi.Output<string>;
    /**
     * Project ID.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * Effective start time, the format is yyyy-MM-dd HH:mm:ss.
     */
    public readonly startTime!: pulumi.Output<string>;
    /**
     * Scheduling configuration: flexible period configuration, only available for hourly/weekly/monthly/yearly scheduling. If the hourly task is specified to run at 0:00, 3:00 and 4:00 every day, it is 0,3,4.
     */
    public readonly taskAction!: pulumi.Output<string>;
    /**
     * Task ID.
     */
    public /*out*/ readonly taskId!: pulumi.Output<string>;
    /**
     * Task Information.
     */
    public readonly taskInfo!: pulumi.Output<outputs.Wedata.IntegrationOfflineTaskTaskInfo>;
    /**
     * Task display mode, 0: canvas mode, 1: form mode.
     */
    public readonly taskMode!: pulumi.Output<string>;
    /**
     * Task name.
     */
    public readonly taskName!: pulumi.Output<string>;

    /**
     * Create a IntegrationOfflineTask resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IntegrationOfflineTaskArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IntegrationOfflineTaskArgs | IntegrationOfflineTaskState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as IntegrationOfflineTaskState | undefined;
            resourceInputs["cycleStep"] = state ? state.cycleStep : undefined;
            resourceInputs["delayTime"] = state ? state.delayTime : undefined;
            resourceInputs["endTime"] = state ? state.endTime : undefined;
            resourceInputs["notes"] = state ? state.notes : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["startTime"] = state ? state.startTime : undefined;
            resourceInputs["taskAction"] = state ? state.taskAction : undefined;
            resourceInputs["taskId"] = state ? state.taskId : undefined;
            resourceInputs["taskInfo"] = state ? state.taskInfo : undefined;
            resourceInputs["taskMode"] = state ? state.taskMode : undefined;
            resourceInputs["taskName"] = state ? state.taskName : undefined;
        } else {
            const args = argsOrState as IntegrationOfflineTaskArgs | undefined;
            if ((!args || args.cycleStep === undefined) && !opts.urn) {
                throw new Error("Missing required property 'cycleStep'");
            }
            if ((!args || args.delayTime === undefined) && !opts.urn) {
                throw new Error("Missing required property 'delayTime'");
            }
            if ((!args || args.endTime === undefined) && !opts.urn) {
                throw new Error("Missing required property 'endTime'");
            }
            if ((!args || args.notes === undefined) && !opts.urn) {
                throw new Error("Missing required property 'notes'");
            }
            if ((!args || args.projectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectId'");
            }
            if ((!args || args.startTime === undefined) && !opts.urn) {
                throw new Error("Missing required property 'startTime'");
            }
            if ((!args || args.taskAction === undefined) && !opts.urn) {
                throw new Error("Missing required property 'taskAction'");
            }
            if ((!args || args.taskInfo === undefined) && !opts.urn) {
                throw new Error("Missing required property 'taskInfo'");
            }
            if ((!args || args.taskMode === undefined) && !opts.urn) {
                throw new Error("Missing required property 'taskMode'");
            }
            if ((!args || args.taskName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'taskName'");
            }
            resourceInputs["cycleStep"] = args ? args.cycleStep : undefined;
            resourceInputs["delayTime"] = args ? args.delayTime : undefined;
            resourceInputs["endTime"] = args ? args.endTime : undefined;
            resourceInputs["notes"] = args ? args.notes : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["startTime"] = args ? args.startTime : undefined;
            resourceInputs["taskAction"] = args ? args.taskAction : undefined;
            resourceInputs["taskInfo"] = args ? args.taskInfo : undefined;
            resourceInputs["taskMode"] = args ? args.taskMode : undefined;
            resourceInputs["taskName"] = args ? args.taskName : undefined;
            resourceInputs["taskId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(IntegrationOfflineTask.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IntegrationOfflineTask resources.
 */
export interface IntegrationOfflineTaskState {
    /**
     * Interval time of scheduling, the minimum value: 1.
     */
    cycleStep?: pulumi.Input<number>;
    /**
     * Execution time, unit is minutes, only available for day/week/month/year scheduling. For example, daily scheduling is executed once every day at 02:00, and the delayTime is 120 minutes.
     */
    delayTime?: pulumi.Input<number>;
    /**
     * Effective end time, the format is yyyy-MM-dd HH:mm:ss.
     */
    endTime?: pulumi.Input<string>;
    /**
     * Description information.
     */
    notes?: pulumi.Input<string>;
    /**
     * Project ID.
     */
    projectId?: pulumi.Input<string>;
    /**
     * Effective start time, the format is yyyy-MM-dd HH:mm:ss.
     */
    startTime?: pulumi.Input<string>;
    /**
     * Scheduling configuration: flexible period configuration, only available for hourly/weekly/monthly/yearly scheduling. If the hourly task is specified to run at 0:00, 3:00 and 4:00 every day, it is 0,3,4.
     */
    taskAction?: pulumi.Input<string>;
    /**
     * Task ID.
     */
    taskId?: pulumi.Input<string>;
    /**
     * Task Information.
     */
    taskInfo?: pulumi.Input<inputs.Wedata.IntegrationOfflineTaskTaskInfo>;
    /**
     * Task display mode, 0: canvas mode, 1: form mode.
     */
    taskMode?: pulumi.Input<string>;
    /**
     * Task name.
     */
    taskName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a IntegrationOfflineTask resource.
 */
export interface IntegrationOfflineTaskArgs {
    /**
     * Interval time of scheduling, the minimum value: 1.
     */
    cycleStep: pulumi.Input<number>;
    /**
     * Execution time, unit is minutes, only available for day/week/month/year scheduling. For example, daily scheduling is executed once every day at 02:00, and the delayTime is 120 minutes.
     */
    delayTime: pulumi.Input<number>;
    /**
     * Effective end time, the format is yyyy-MM-dd HH:mm:ss.
     */
    endTime: pulumi.Input<string>;
    /**
     * Description information.
     */
    notes: pulumi.Input<string>;
    /**
     * Project ID.
     */
    projectId: pulumi.Input<string>;
    /**
     * Effective start time, the format is yyyy-MM-dd HH:mm:ss.
     */
    startTime: pulumi.Input<string>;
    /**
     * Scheduling configuration: flexible period configuration, only available for hourly/weekly/monthly/yearly scheduling. If the hourly task is specified to run at 0:00, 3:00 and 4:00 every day, it is 0,3,4.
     */
    taskAction: pulumi.Input<string>;
    /**
     * Task Information.
     */
    taskInfo: pulumi.Input<inputs.Wedata.IntegrationOfflineTaskTaskInfo>;
    /**
     * Task display mode, 0: canvas mode, 1: form mode.
     */
    taskMode: pulumi.Input<string>;
    /**
     * Task name.
     */
    taskName: pulumi.Input<string>;
}
