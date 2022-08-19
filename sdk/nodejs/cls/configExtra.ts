// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a cls config extra
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const extra = new tencentcloud.cls.ConfigExtra("extra", {
 *     topicId: tencentcloud_cls_topic.topic.id,
 *     type: "container_file",
 *     logType: "json_log",
 *     configFlag: "label_k8s",
 *     logsetId: tencentcloud_cls_logset.logset.id,
 *     logsetName: tencentcloud_cls_logset.logset.logset_name,
 *     topicName: tencentcloud_cls_topic.topic.topic_name,
 *     containerFile: {
 *         container: "nginx",
 *         filePattern: "log",
 *         logPath: "/nginx",
 *         namespace: "default",
 *         workload: {
 *             container: "nginx",
 *             kind: "deployment",
 *             name: "nginx",
 *             namespace: "default",
 *         },
 *     },
 *     groupId: "27752a9b-9918-440a-8ee7-9c84a14a47ed",
 * });
 * ```
 */
export class ConfigExtra extends pulumi.CustomResource {
    /**
     * Get an existing ConfigExtra resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ConfigExtraState, opts?: pulumi.CustomResourceOptions): ConfigExtra {
        return new ConfigExtra(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Cls/configExtra:ConfigExtra';

    /**
     * Returns true if the given object is an instance of ConfigExtra.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ConfigExtra {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ConfigExtra.__pulumiType;
    }

    /**
     * Collection configuration flag.
     */
    public readonly configFlag!: pulumi.Output<string>;
    /**
     * Container file path info.
     */
    public readonly containerFile!: pulumi.Output<outputs.Cls.ConfigExtraContainerFile | undefined>;
    /**
     * Container stdout info.
     */
    public readonly containerStdout!: pulumi.Output<outputs.Cls.ConfigExtraContainerStdout | undefined>;
    /**
     * Collection path blocklist.
     */
    public readonly excludePaths!: pulumi.Output<outputs.Cls.ConfigExtraExcludePath[] | undefined>;
    /**
     * Extraction rule. If ExtractRule is set, LogType must be set.
     */
    public readonly extractRule!: pulumi.Output<outputs.Cls.ConfigExtraExtractRule | undefined>;
    /**
     * Binding group id.
     */
    public readonly groupId!: pulumi.Output<string | undefined>;
    /**
     * Binding group ids.
     */
    public readonly groupIds!: pulumi.Output<string[] | undefined>;
    /**
     * Node file config info.
     */
    public readonly hostFile!: pulumi.Output<outputs.Cls.ConfigExtraHostFile | undefined>;
    /**
     * Type of the log to be collected. Valid values: json_log: log in JSON format; delimiter_log: log in delimited format; minimalist_log: minimalist log; multiline_log: log in multi-line format; fullregex_log: log in full regex format. Default value: minimalist_log.
     */
    public readonly logType!: pulumi.Output<string>;
    /**
     * Logset Id.
     */
    public readonly logsetId!: pulumi.Output<string>;
    /**
     * Logset Name.
     */
    public readonly logsetName!: pulumi.Output<string>;
    /**
     * Collection configuration name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Log topic ID (TopicId) of collection configuration.
     */
    public readonly topicId!: pulumi.Output<string>;
    /**
     * Topic Name.
     */
    public readonly topicName!: pulumi.Output<string>;
    /**
     * Type. Valid values: container_stdout; container_file; host_file.
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * Custom collection rule, which is a serialized JSON string.
     */
    public readonly userDefineRule!: pulumi.Output<string | undefined>;

    /**
     * Create a ConfigExtra resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ConfigExtraArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ConfigExtraArgs | ConfigExtraState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ConfigExtraState | undefined;
            resourceInputs["configFlag"] = state ? state.configFlag : undefined;
            resourceInputs["containerFile"] = state ? state.containerFile : undefined;
            resourceInputs["containerStdout"] = state ? state.containerStdout : undefined;
            resourceInputs["excludePaths"] = state ? state.excludePaths : undefined;
            resourceInputs["extractRule"] = state ? state.extractRule : undefined;
            resourceInputs["groupId"] = state ? state.groupId : undefined;
            resourceInputs["groupIds"] = state ? state.groupIds : undefined;
            resourceInputs["hostFile"] = state ? state.hostFile : undefined;
            resourceInputs["logType"] = state ? state.logType : undefined;
            resourceInputs["logsetId"] = state ? state.logsetId : undefined;
            resourceInputs["logsetName"] = state ? state.logsetName : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["topicId"] = state ? state.topicId : undefined;
            resourceInputs["topicName"] = state ? state.topicName : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["userDefineRule"] = state ? state.userDefineRule : undefined;
        } else {
            const args = argsOrState as ConfigExtraArgs | undefined;
            if ((!args || args.configFlag === undefined) && !opts.urn) {
                throw new Error("Missing required property 'configFlag'");
            }
            if ((!args || args.logType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'logType'");
            }
            if ((!args || args.logsetId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'logsetId'");
            }
            if ((!args || args.logsetName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'logsetName'");
            }
            if ((!args || args.topicId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'topicId'");
            }
            if ((!args || args.topicName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'topicName'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["configFlag"] = args ? args.configFlag : undefined;
            resourceInputs["containerFile"] = args ? args.containerFile : undefined;
            resourceInputs["containerStdout"] = args ? args.containerStdout : undefined;
            resourceInputs["excludePaths"] = args ? args.excludePaths : undefined;
            resourceInputs["extractRule"] = args ? args.extractRule : undefined;
            resourceInputs["groupId"] = args ? args.groupId : undefined;
            resourceInputs["groupIds"] = args ? args.groupIds : undefined;
            resourceInputs["hostFile"] = args ? args.hostFile : undefined;
            resourceInputs["logType"] = args ? args.logType : undefined;
            resourceInputs["logsetId"] = args ? args.logsetId : undefined;
            resourceInputs["logsetName"] = args ? args.logsetName : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["topicId"] = args ? args.topicId : undefined;
            resourceInputs["topicName"] = args ? args.topicName : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["userDefineRule"] = args ? args.userDefineRule : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ConfigExtra.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ConfigExtra resources.
 */
export interface ConfigExtraState {
    /**
     * Collection configuration flag.
     */
    configFlag?: pulumi.Input<string>;
    /**
     * Container file path info.
     */
    containerFile?: pulumi.Input<inputs.Cls.ConfigExtraContainerFile>;
    /**
     * Container stdout info.
     */
    containerStdout?: pulumi.Input<inputs.Cls.ConfigExtraContainerStdout>;
    /**
     * Collection path blocklist.
     */
    excludePaths?: pulumi.Input<pulumi.Input<inputs.Cls.ConfigExtraExcludePath>[]>;
    /**
     * Extraction rule. If ExtractRule is set, LogType must be set.
     */
    extractRule?: pulumi.Input<inputs.Cls.ConfigExtraExtractRule>;
    /**
     * Binding group id.
     */
    groupId?: pulumi.Input<string>;
    /**
     * Binding group ids.
     */
    groupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Node file config info.
     */
    hostFile?: pulumi.Input<inputs.Cls.ConfigExtraHostFile>;
    /**
     * Type of the log to be collected. Valid values: json_log: log in JSON format; delimiter_log: log in delimited format; minimalist_log: minimalist log; multiline_log: log in multi-line format; fullregex_log: log in full regex format. Default value: minimalist_log.
     */
    logType?: pulumi.Input<string>;
    /**
     * Logset Id.
     */
    logsetId?: pulumi.Input<string>;
    /**
     * Logset Name.
     */
    logsetName?: pulumi.Input<string>;
    /**
     * Collection configuration name.
     */
    name?: pulumi.Input<string>;
    /**
     * Log topic ID (TopicId) of collection configuration.
     */
    topicId?: pulumi.Input<string>;
    /**
     * Topic Name.
     */
    topicName?: pulumi.Input<string>;
    /**
     * Type. Valid values: container_stdout; container_file; host_file.
     */
    type?: pulumi.Input<string>;
    /**
     * Custom collection rule, which is a serialized JSON string.
     */
    userDefineRule?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ConfigExtra resource.
 */
export interface ConfigExtraArgs {
    /**
     * Collection configuration flag.
     */
    configFlag: pulumi.Input<string>;
    /**
     * Container file path info.
     */
    containerFile?: pulumi.Input<inputs.Cls.ConfigExtraContainerFile>;
    /**
     * Container stdout info.
     */
    containerStdout?: pulumi.Input<inputs.Cls.ConfigExtraContainerStdout>;
    /**
     * Collection path blocklist.
     */
    excludePaths?: pulumi.Input<pulumi.Input<inputs.Cls.ConfigExtraExcludePath>[]>;
    /**
     * Extraction rule. If ExtractRule is set, LogType must be set.
     */
    extractRule?: pulumi.Input<inputs.Cls.ConfigExtraExtractRule>;
    /**
     * Binding group id.
     */
    groupId?: pulumi.Input<string>;
    /**
     * Binding group ids.
     */
    groupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Node file config info.
     */
    hostFile?: pulumi.Input<inputs.Cls.ConfigExtraHostFile>;
    /**
     * Type of the log to be collected. Valid values: json_log: log in JSON format; delimiter_log: log in delimited format; minimalist_log: minimalist log; multiline_log: log in multi-line format; fullregex_log: log in full regex format. Default value: minimalist_log.
     */
    logType: pulumi.Input<string>;
    /**
     * Logset Id.
     */
    logsetId: pulumi.Input<string>;
    /**
     * Logset Name.
     */
    logsetName: pulumi.Input<string>;
    /**
     * Collection configuration name.
     */
    name?: pulumi.Input<string>;
    /**
     * Log topic ID (TopicId) of collection configuration.
     */
    topicId: pulumi.Input<string>;
    /**
     * Topic Name.
     */
    topicName: pulumi.Input<string>;
    /**
     * Type. Valid values: container_stdout; container_file; host_file.
     */
    type: pulumi.Input<string>;
    /**
     * Custom collection rule, which is a serialized JSON string.
     */
    userDefineRule?: pulumi.Input<string>;
}
