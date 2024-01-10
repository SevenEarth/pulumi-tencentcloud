// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a pts cronJob
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const cronJob = new tencentcloud.Pts.CronJob("cron_job", {
 *     cronExpression: "* 1 * * *",
 *     frequencyType: 2,
 *     jobOwner: "userName",
 *     note: "desc",
 *     // end_time = ""
 *     noticeId: "notice-vp6i38jt",
 *     projectId: "project-7qkzxhea",
 *     scenarioId: "scenario-c22lqb1w",
 *     scenarioName: "pts-js(2022-11-10 21:53:53)",
 * });
 * ```
 *
 * ## Import
 *
 * pts cron_job can be imported using the projectId#cronJobId, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:Pts/cronJob:CronJob cron_job project-7qkzxhea#scenario-c22lqb1w
 * ```
 */
export class CronJob extends pulumi.CustomResource {
    /**
     * Get an existing CronJob resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CronJobState, opts?: pulumi.CustomResourceOptions): CronJob {
        return new CronJob(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Pts/cronJob:CronJob';

    /**
     * Returns true if the given object is an instance of CronJob.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CronJob {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CronJob.__pulumiType;
    }

    /**
     * Reason for suspension.
     */
    public /*out*/ readonly abortReason!: pulumi.Output<number>;
    /**
     * App ID.
     */
    public /*out*/ readonly appId!: pulumi.Output<number>;
    /**
     * Creation time; type: Timestamp ISO8601.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * Cron expression, When setting cronExpression at that time, frequencyType must be greater than 1.
     */
    public readonly cronExpression!: pulumi.Output<string>;
    /**
     * Cron job ID.
     */
    public /*out*/ readonly cronJobId!: pulumi.Output<string>;
    /**
     * End Time; type: Timestamp ISO8601.
     */
    public readonly endTime!: pulumi.Output<string | undefined>;
    /**
     * Execution frequency type, `1`: execute only once; `2`: daily granularity; `3`: weekly granularity; `4`: advanced.
     */
    public readonly frequencyType!: pulumi.Output<number>;
    /**
     * Job Owner.
     */
    public readonly jobOwner!: pulumi.Output<string>;
    /**
     * Cron Job Name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Note.
     */
    public readonly note!: pulumi.Output<string | undefined>;
    /**
     * Notice ID.
     */
    public readonly noticeId!: pulumi.Output<string | undefined>;
    /**
     * Project Id.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * Scenario Id.
     */
    public readonly scenarioId!: pulumi.Output<string>;
    /**
     * Scenario Name.
     */
    public readonly scenarioName!: pulumi.Output<string>;
    /**
     * Scheduled task status.
     */
    public /*out*/ readonly status!: pulumi.Output<number>;
    /**
     * Sub-user ID.
     */
    public /*out*/ readonly subAccountUin!: pulumi.Output<string>;
    /**
     * User ID.
     */
    public /*out*/ readonly uin!: pulumi.Output<string>;
    /**
     * Update time; type: Timestamp ISO8601.
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;

    /**
     * Create a CronJob resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CronJobArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CronJobArgs | CronJobState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CronJobState | undefined;
            resourceInputs["abortReason"] = state ? state.abortReason : undefined;
            resourceInputs["appId"] = state ? state.appId : undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["cronExpression"] = state ? state.cronExpression : undefined;
            resourceInputs["cronJobId"] = state ? state.cronJobId : undefined;
            resourceInputs["endTime"] = state ? state.endTime : undefined;
            resourceInputs["frequencyType"] = state ? state.frequencyType : undefined;
            resourceInputs["jobOwner"] = state ? state.jobOwner : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["note"] = state ? state.note : undefined;
            resourceInputs["noticeId"] = state ? state.noticeId : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["scenarioId"] = state ? state.scenarioId : undefined;
            resourceInputs["scenarioName"] = state ? state.scenarioName : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["subAccountUin"] = state ? state.subAccountUin : undefined;
            resourceInputs["uin"] = state ? state.uin : undefined;
            resourceInputs["updatedAt"] = state ? state.updatedAt : undefined;
        } else {
            const args = argsOrState as CronJobArgs | undefined;
            if ((!args || args.cronExpression === undefined) && !opts.urn) {
                throw new Error("Missing required property 'cronExpression'");
            }
            if ((!args || args.frequencyType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'frequencyType'");
            }
            if ((!args || args.jobOwner === undefined) && !opts.urn) {
                throw new Error("Missing required property 'jobOwner'");
            }
            if ((!args || args.projectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectId'");
            }
            if ((!args || args.scenarioId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'scenarioId'");
            }
            if ((!args || args.scenarioName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'scenarioName'");
            }
            resourceInputs["cronExpression"] = args ? args.cronExpression : undefined;
            resourceInputs["endTime"] = args ? args.endTime : undefined;
            resourceInputs["frequencyType"] = args ? args.frequencyType : undefined;
            resourceInputs["jobOwner"] = args ? args.jobOwner : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["note"] = args ? args.note : undefined;
            resourceInputs["noticeId"] = args ? args.noticeId : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["scenarioId"] = args ? args.scenarioId : undefined;
            resourceInputs["scenarioName"] = args ? args.scenarioName : undefined;
            resourceInputs["abortReason"] = undefined /*out*/;
            resourceInputs["appId"] = undefined /*out*/;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["cronJobId"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["subAccountUin"] = undefined /*out*/;
            resourceInputs["uin"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(CronJob.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CronJob resources.
 */
export interface CronJobState {
    /**
     * Reason for suspension.
     */
    abortReason?: pulumi.Input<number>;
    /**
     * App ID.
     */
    appId?: pulumi.Input<number>;
    /**
     * Creation time; type: Timestamp ISO8601.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * Cron expression, When setting cronExpression at that time, frequencyType must be greater than 1.
     */
    cronExpression?: pulumi.Input<string>;
    /**
     * Cron job ID.
     */
    cronJobId?: pulumi.Input<string>;
    /**
     * End Time; type: Timestamp ISO8601.
     */
    endTime?: pulumi.Input<string>;
    /**
     * Execution frequency type, `1`: execute only once; `2`: daily granularity; `3`: weekly granularity; `4`: advanced.
     */
    frequencyType?: pulumi.Input<number>;
    /**
     * Job Owner.
     */
    jobOwner?: pulumi.Input<string>;
    /**
     * Cron Job Name.
     */
    name?: pulumi.Input<string>;
    /**
     * Note.
     */
    note?: pulumi.Input<string>;
    /**
     * Notice ID.
     */
    noticeId?: pulumi.Input<string>;
    /**
     * Project Id.
     */
    projectId?: pulumi.Input<string>;
    /**
     * Scenario Id.
     */
    scenarioId?: pulumi.Input<string>;
    /**
     * Scenario Name.
     */
    scenarioName?: pulumi.Input<string>;
    /**
     * Scheduled task status.
     */
    status?: pulumi.Input<number>;
    /**
     * Sub-user ID.
     */
    subAccountUin?: pulumi.Input<string>;
    /**
     * User ID.
     */
    uin?: pulumi.Input<string>;
    /**
     * Update time; type: Timestamp ISO8601.
     */
    updatedAt?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a CronJob resource.
 */
export interface CronJobArgs {
    /**
     * Cron expression, When setting cronExpression at that time, frequencyType must be greater than 1.
     */
    cronExpression: pulumi.Input<string>;
    /**
     * End Time; type: Timestamp ISO8601.
     */
    endTime?: pulumi.Input<string>;
    /**
     * Execution frequency type, `1`: execute only once; `2`: daily granularity; `3`: weekly granularity; `4`: advanced.
     */
    frequencyType: pulumi.Input<number>;
    /**
     * Job Owner.
     */
    jobOwner: pulumi.Input<string>;
    /**
     * Cron Job Name.
     */
    name?: pulumi.Input<string>;
    /**
     * Note.
     */
    note?: pulumi.Input<string>;
    /**
     * Notice ID.
     */
    noticeId?: pulumi.Input<string>;
    /**
     * Project Id.
     */
    projectId: pulumi.Input<string>;
    /**
     * Scenario Id.
     */
    scenarioId: pulumi.Input<string>;
    /**
     * Scenario Name.
     */
    scenarioName: pulumi.Input<string>;
}
