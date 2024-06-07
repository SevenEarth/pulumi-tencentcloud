// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a css pullStreamTask
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@tencentcloud_iac/pulumi";
 *
 * const pullStreamTask = new tencentcloud.css.PullStreamTask("pullStreamTask", {
 *     appName: "app_name",
 *     comment: "comment.",
 *     domainName: "domain_name",
 *     endTime: "2022-11-16T22:09:28Z",
 *     operator: "admin",
 *     sourceType: "source_type",
 *     sourceUrls: ["source_urls"],
 *     startTime: "2022-11-16T22:09:28Z",
 *     streamName: "stream_name",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * css pull_stream_task can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import tencentcloud:Css/pullStreamTask:PullStreamTask pull_stream_task pullStreamTask_id
 * ```
 */
export class PullStreamTask extends pulumi.CustomResource {
    /**
     * Get an existing PullStreamTask resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PullStreamTaskState, opts?: pulumi.CustomResourceOptions): PullStreamTask {
        return new PullStreamTask(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Css/pullStreamTask:PullStreamTask';

    /**
     * Returns true if the given object is an instance of PullStreamTask.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PullStreamTask {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PullStreamTask.__pulumiType;
    }

    /**
     * push app name.
     */
    public readonly appName!: pulumi.Output<string>;
    /**
     * backup pull source type.
     */
    public readonly backupSourceType!: pulumi.Output<string | undefined>;
    /**
     * backup pull source.
     */
    public readonly backupSourceUrl!: pulumi.Output<string | undefined>;
    /**
     * defind the callback event you need, null for all. TaskStart, TaskExit, VodSourceFileStart, VodSourceFileFinish, ResetTaskConfig, PullFileUnstable, PushStreamUnstable, PullFileFailed, PushStreamFailed, FileEndEarly.
     */
    public readonly callbackEvents!: pulumi.Output<string[]>;
    /**
     * task event callback url.
     */
    public readonly callbackUrl!: pulumi.Output<string | undefined>;
    /**
     * desc for pull task.
     */
    public readonly comment!: pulumi.Output<string | undefined>;
    /**
     * desc who create the task.
     */
    public /*out*/ readonly createBy!: pulumi.Output<string>;
    /**
     * create time.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * push domain name.
     */
    public readonly domainName!: pulumi.Output<string>;
    /**
     * task end time.
     */
    public readonly endTime!: pulumi.Output<string>;
    /**
     * ignore_region for ignore the input region and reblance inside the server.
     */
    public readonly extraCmd!: pulumi.Output<string | undefined>;
    /**
     * task enable or disable.
     */
    public readonly fileIndex!: pulumi.Output<number | undefined>;
    /**
     * task enable or disable.
     */
    public readonly offsetTime!: pulumi.Output<number | undefined>;
    /**
     * desc operator user name.
     */
    public readonly operator!: pulumi.Output<string | undefined>;
    /**
     * other pushing args.
     */
    public readonly pushArgs!: pulumi.Output<string | undefined>;
    /**
     * task run region.
     */
    public /*out*/ readonly region!: pulumi.Output<string>;
    /**
     * `PullLivePushLive`: SourceUrls live type, `PullVodPushLive`: SourceUrls vod type.
     */
    public readonly sourceType!: pulumi.Output<string>;
    /**
     * Pull Source media, SourceType=PullLivePushLive only 1 value, SourceType=PullLivePushLive can input multi values.
     */
    public readonly sourceUrls!: pulumi.Output<string[]>;
    /**
     * task begin time.
     */
    public readonly startTime!: pulumi.Output<string>;
    /**
     * task enable or disable.
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * push stream name.
     */
    public readonly streamName!: pulumi.Output<string>;
    /**
     * full target push url, DomainName, AppName, StreamName field must be empty.
     */
    public readonly toUrl!: pulumi.Output<string | undefined>;
    /**
     * desc who update the task.
     */
    public /*out*/ readonly updateBy!: pulumi.Output<string>;
    /**
     * update time.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;
    /**
     * loop time for vod.
     */
    public readonly vodLoopTimes!: pulumi.Output<number>;
    /**
     * vod refresh method. `ImmediateNewSource`: switch to new source at once, `ContinueBreakPoint`: switch to new source while old source finish.
     */
    public readonly vodRefreshType!: pulumi.Output<string>;
    /**
     * watermark list, max 4 setting.
     */
    public readonly watermarkLists!: pulumi.Output<outputs.Css.PullStreamTaskWatermarkList[] | undefined>;

    /**
     * Create a PullStreamTask resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PullStreamTaskArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PullStreamTaskArgs | PullStreamTaskState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PullStreamTaskState | undefined;
            resourceInputs["appName"] = state ? state.appName : undefined;
            resourceInputs["backupSourceType"] = state ? state.backupSourceType : undefined;
            resourceInputs["backupSourceUrl"] = state ? state.backupSourceUrl : undefined;
            resourceInputs["callbackEvents"] = state ? state.callbackEvents : undefined;
            resourceInputs["callbackUrl"] = state ? state.callbackUrl : undefined;
            resourceInputs["comment"] = state ? state.comment : undefined;
            resourceInputs["createBy"] = state ? state.createBy : undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["domainName"] = state ? state.domainName : undefined;
            resourceInputs["endTime"] = state ? state.endTime : undefined;
            resourceInputs["extraCmd"] = state ? state.extraCmd : undefined;
            resourceInputs["fileIndex"] = state ? state.fileIndex : undefined;
            resourceInputs["offsetTime"] = state ? state.offsetTime : undefined;
            resourceInputs["operator"] = state ? state.operator : undefined;
            resourceInputs["pushArgs"] = state ? state.pushArgs : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["sourceType"] = state ? state.sourceType : undefined;
            resourceInputs["sourceUrls"] = state ? state.sourceUrls : undefined;
            resourceInputs["startTime"] = state ? state.startTime : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["streamName"] = state ? state.streamName : undefined;
            resourceInputs["toUrl"] = state ? state.toUrl : undefined;
            resourceInputs["updateBy"] = state ? state.updateBy : undefined;
            resourceInputs["updateTime"] = state ? state.updateTime : undefined;
            resourceInputs["vodLoopTimes"] = state ? state.vodLoopTimes : undefined;
            resourceInputs["vodRefreshType"] = state ? state.vodRefreshType : undefined;
            resourceInputs["watermarkLists"] = state ? state.watermarkLists : undefined;
        } else {
            const args = argsOrState as PullStreamTaskArgs | undefined;
            if ((!args || args.appName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'appName'");
            }
            if ((!args || args.domainName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'domainName'");
            }
            if ((!args || args.endTime === undefined) && !opts.urn) {
                throw new Error("Missing required property 'endTime'");
            }
            if ((!args || args.sourceType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sourceType'");
            }
            if ((!args || args.sourceUrls === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sourceUrls'");
            }
            if ((!args || args.startTime === undefined) && !opts.urn) {
                throw new Error("Missing required property 'startTime'");
            }
            if ((!args || args.streamName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'streamName'");
            }
            resourceInputs["appName"] = args ? args.appName : undefined;
            resourceInputs["backupSourceType"] = args ? args.backupSourceType : undefined;
            resourceInputs["backupSourceUrl"] = args ? args.backupSourceUrl : undefined;
            resourceInputs["callbackEvents"] = args ? args.callbackEvents : undefined;
            resourceInputs["callbackUrl"] = args ? args.callbackUrl : undefined;
            resourceInputs["comment"] = args ? args.comment : undefined;
            resourceInputs["domainName"] = args ? args.domainName : undefined;
            resourceInputs["endTime"] = args ? args.endTime : undefined;
            resourceInputs["extraCmd"] = args ? args.extraCmd : undefined;
            resourceInputs["fileIndex"] = args ? args.fileIndex : undefined;
            resourceInputs["offsetTime"] = args ? args.offsetTime : undefined;
            resourceInputs["operator"] = args ? args.operator : undefined;
            resourceInputs["pushArgs"] = args ? args.pushArgs : undefined;
            resourceInputs["sourceType"] = args ? args.sourceType : undefined;
            resourceInputs["sourceUrls"] = args ? args.sourceUrls : undefined;
            resourceInputs["startTime"] = args ? args.startTime : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["streamName"] = args ? args.streamName : undefined;
            resourceInputs["toUrl"] = args ? args.toUrl : undefined;
            resourceInputs["vodLoopTimes"] = args ? args.vodLoopTimes : undefined;
            resourceInputs["vodRefreshType"] = args ? args.vodRefreshType : undefined;
            resourceInputs["watermarkLists"] = args ? args.watermarkLists : undefined;
            resourceInputs["createBy"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["region"] = undefined /*out*/;
            resourceInputs["updateBy"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(PullStreamTask.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PullStreamTask resources.
 */
export interface PullStreamTaskState {
    /**
     * push app name.
     */
    appName?: pulumi.Input<string>;
    /**
     * backup pull source type.
     */
    backupSourceType?: pulumi.Input<string>;
    /**
     * backup pull source.
     */
    backupSourceUrl?: pulumi.Input<string>;
    /**
     * defind the callback event you need, null for all. TaskStart, TaskExit, VodSourceFileStart, VodSourceFileFinish, ResetTaskConfig, PullFileUnstable, PushStreamUnstable, PullFileFailed, PushStreamFailed, FileEndEarly.
     */
    callbackEvents?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * task event callback url.
     */
    callbackUrl?: pulumi.Input<string>;
    /**
     * desc for pull task.
     */
    comment?: pulumi.Input<string>;
    /**
     * desc who create the task.
     */
    createBy?: pulumi.Input<string>;
    /**
     * create time.
     */
    createTime?: pulumi.Input<string>;
    /**
     * push domain name.
     */
    domainName?: pulumi.Input<string>;
    /**
     * task end time.
     */
    endTime?: pulumi.Input<string>;
    /**
     * ignore_region for ignore the input region and reblance inside the server.
     */
    extraCmd?: pulumi.Input<string>;
    /**
     * task enable or disable.
     */
    fileIndex?: pulumi.Input<number>;
    /**
     * task enable or disable.
     */
    offsetTime?: pulumi.Input<number>;
    /**
     * desc operator user name.
     */
    operator?: pulumi.Input<string>;
    /**
     * other pushing args.
     */
    pushArgs?: pulumi.Input<string>;
    /**
     * task run region.
     */
    region?: pulumi.Input<string>;
    /**
     * `PullLivePushLive`: SourceUrls live type, `PullVodPushLive`: SourceUrls vod type.
     */
    sourceType?: pulumi.Input<string>;
    /**
     * Pull Source media, SourceType=PullLivePushLive only 1 value, SourceType=PullLivePushLive can input multi values.
     */
    sourceUrls?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * task begin time.
     */
    startTime?: pulumi.Input<string>;
    /**
     * task enable or disable.
     */
    status?: pulumi.Input<string>;
    /**
     * push stream name.
     */
    streamName?: pulumi.Input<string>;
    /**
     * full target push url, DomainName, AppName, StreamName field must be empty.
     */
    toUrl?: pulumi.Input<string>;
    /**
     * desc who update the task.
     */
    updateBy?: pulumi.Input<string>;
    /**
     * update time.
     */
    updateTime?: pulumi.Input<string>;
    /**
     * loop time for vod.
     */
    vodLoopTimes?: pulumi.Input<number>;
    /**
     * vod refresh method. `ImmediateNewSource`: switch to new source at once, `ContinueBreakPoint`: switch to new source while old source finish.
     */
    vodRefreshType?: pulumi.Input<string>;
    /**
     * watermark list, max 4 setting.
     */
    watermarkLists?: pulumi.Input<pulumi.Input<inputs.Css.PullStreamTaskWatermarkList>[]>;
}

/**
 * The set of arguments for constructing a PullStreamTask resource.
 */
export interface PullStreamTaskArgs {
    /**
     * push app name.
     */
    appName: pulumi.Input<string>;
    /**
     * backup pull source type.
     */
    backupSourceType?: pulumi.Input<string>;
    /**
     * backup pull source.
     */
    backupSourceUrl?: pulumi.Input<string>;
    /**
     * defind the callback event you need, null for all. TaskStart, TaskExit, VodSourceFileStart, VodSourceFileFinish, ResetTaskConfig, PullFileUnstable, PushStreamUnstable, PullFileFailed, PushStreamFailed, FileEndEarly.
     */
    callbackEvents?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * task event callback url.
     */
    callbackUrl?: pulumi.Input<string>;
    /**
     * desc for pull task.
     */
    comment?: pulumi.Input<string>;
    /**
     * push domain name.
     */
    domainName: pulumi.Input<string>;
    /**
     * task end time.
     */
    endTime: pulumi.Input<string>;
    /**
     * ignore_region for ignore the input region and reblance inside the server.
     */
    extraCmd?: pulumi.Input<string>;
    /**
     * task enable or disable.
     */
    fileIndex?: pulumi.Input<number>;
    /**
     * task enable or disable.
     */
    offsetTime?: pulumi.Input<number>;
    /**
     * desc operator user name.
     */
    operator?: pulumi.Input<string>;
    /**
     * other pushing args.
     */
    pushArgs?: pulumi.Input<string>;
    /**
     * `PullLivePushLive`: SourceUrls live type, `PullVodPushLive`: SourceUrls vod type.
     */
    sourceType: pulumi.Input<string>;
    /**
     * Pull Source media, SourceType=PullLivePushLive only 1 value, SourceType=PullLivePushLive can input multi values.
     */
    sourceUrls: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * task begin time.
     */
    startTime: pulumi.Input<string>;
    /**
     * task enable or disable.
     */
    status?: pulumi.Input<string>;
    /**
     * push stream name.
     */
    streamName: pulumi.Input<string>;
    /**
     * full target push url, DomainName, AppName, StreamName field must be empty.
     */
    toUrl?: pulumi.Input<string>;
    /**
     * loop time for vod.
     */
    vodLoopTimes?: pulumi.Input<number>;
    /**
     * vod refresh method. `ImmediateNewSource`: switch to new source at once, `ContinueBreakPoint`: switch to new source while old source finish.
     */
    vodRefreshType?: pulumi.Input<string>;
    /**
     * watermark list, max 4 setting.
     */
    watermarkLists?: pulumi.Input<pulumi.Input<inputs.Css.PullStreamTaskWatermarkList>[]>;
}
