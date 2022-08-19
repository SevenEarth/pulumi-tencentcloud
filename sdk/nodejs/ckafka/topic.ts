// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this resource to create ckafka topic.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const foo = new tencentcloud.Ckafka.Topic("foo", {
 *     cleanUpPolicy: "delete",
 *     enableWhiteList: true,
 *     instanceId: "ckafka-f9ife4zz",
 *     ipWhiteLists: [
 *         "ip1",
 *         "ip2",
 *     ],
 *     maxMessageBytes: 0,
 *     note: "topic note",
 *     partitionNum: 1,
 *     replicaNum: 2,
 *     retention: 60000,
 *     segment: 3600000,
 *     syncReplicaMinNum: 1,
 *     topicName: "example",
 *     uncleanLeaderElectionEnable: false,
 * });
 * ```
 *
 * ## Import
 *
 * ckafka topic can be imported using the instance_id#topic_name, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:Ckafka/topic:Topic foo ckafka-f9ife4zz#example
 * ```
 */
export class Topic extends pulumi.CustomResource {
    /**
     * Get an existing Topic resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TopicState, opts?: pulumi.CustomResourceOptions): Topic {
        return new Topic(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Ckafka/topic:Topic';

    /**
     * Returns true if the given object is an instance of Topic.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Topic {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Topic.__pulumiType;
    }

    /**
     * Clear log policy, log clear mode, default is `delete`. `delete`: logs are deleted according to the storage time. `compact`: logs are compressed according to the key. `compact, delete`: logs are compressed according to the key and will be deleted according to the storage time.
     */
    public readonly cleanUpPolicy!: pulumi.Output<string | undefined>;
    /**
     * Create time of the CKafka topic.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Whether to open the ip whitelist, `true`: open, `false`: close.
     */
    public readonly enableWhiteList!: pulumi.Output<boolean | undefined>;
    /**
     * Data backup cos bucket: the bucket address that is dumped to cos.
     */
    public /*out*/ readonly forwardCosBucket!: pulumi.Output<string>;
    /**
     * Periodic frequency of data backup to cos.
     */
    public /*out*/ readonly forwardInterval!: pulumi.Output<number>;
    /**
     * Data backup cos status. Valid values: `0`, `1`. `1`: do not open data backup, `0`: open data backup.
     */
    public /*out*/ readonly forwardStatus!: pulumi.Output<number>;
    /**
     * Ckafka instance ID.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * Ip whitelist, quota limit, required when enableWhileList=true.
     */
    public readonly ipWhiteLists!: pulumi.Output<string[] | undefined>;
    /**
     * Max message bytes.
     */
    public readonly maxMessageBytes!: pulumi.Output<number | undefined>;
    /**
     * Message storage location.
     */
    public /*out*/ readonly messageStorageLocation!: pulumi.Output<string>;
    /**
     * The subject note. It must start with a letter, and the remaining part can contain letters, numbers and dashes (-).
     */
    public readonly note!: pulumi.Output<string | undefined>;
    /**
     * The number of partition.
     */
    public readonly partitionNum!: pulumi.Output<number>;
    /**
     * The number of replica.
     */
    public readonly replicaNum!: pulumi.Output<number>;
    /**
     * Message can be selected. Retention time, unit is ms, the current minimum value is 60000ms.
     */
    public readonly retention!: pulumi.Output<number | undefined>;
    /**
     * Segment scrolling time, in ms, the current minimum is 3600000ms.
     */
    public readonly segment!: pulumi.Output<number | undefined>;
    /**
     * Number of bytes rolled by shard.
     */
    public /*out*/ readonly segmentBytes!: pulumi.Output<number>;
    /**
     * Min number of sync replicas, Default is `1`.
     */
    public readonly syncReplicaMinNum!: pulumi.Output<number | undefined>;
    /**
     * Name of the CKafka topic. It must start with a letter, the rest can contain letters, numbers and dashes(-).
     */
    public readonly topicName!: pulumi.Output<string>;
    /**
     * Whether to allow unsynchronized replicas to be selected as leader, default is `false`, `true: `allowed, `false`: not allowed.
     */
    public readonly uncleanLeaderElectionEnable!: pulumi.Output<boolean | undefined>;

    /**
     * Create a Topic resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TopicArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TopicArgs | TopicState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TopicState | undefined;
            resourceInputs["cleanUpPolicy"] = state ? state.cleanUpPolicy : undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["enableWhiteList"] = state ? state.enableWhiteList : undefined;
            resourceInputs["forwardCosBucket"] = state ? state.forwardCosBucket : undefined;
            resourceInputs["forwardInterval"] = state ? state.forwardInterval : undefined;
            resourceInputs["forwardStatus"] = state ? state.forwardStatus : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["ipWhiteLists"] = state ? state.ipWhiteLists : undefined;
            resourceInputs["maxMessageBytes"] = state ? state.maxMessageBytes : undefined;
            resourceInputs["messageStorageLocation"] = state ? state.messageStorageLocation : undefined;
            resourceInputs["note"] = state ? state.note : undefined;
            resourceInputs["partitionNum"] = state ? state.partitionNum : undefined;
            resourceInputs["replicaNum"] = state ? state.replicaNum : undefined;
            resourceInputs["retention"] = state ? state.retention : undefined;
            resourceInputs["segment"] = state ? state.segment : undefined;
            resourceInputs["segmentBytes"] = state ? state.segmentBytes : undefined;
            resourceInputs["syncReplicaMinNum"] = state ? state.syncReplicaMinNum : undefined;
            resourceInputs["topicName"] = state ? state.topicName : undefined;
            resourceInputs["uncleanLeaderElectionEnable"] = state ? state.uncleanLeaderElectionEnable : undefined;
        } else {
            const args = argsOrState as TopicArgs | undefined;
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            if ((!args || args.partitionNum === undefined) && !opts.urn) {
                throw new Error("Missing required property 'partitionNum'");
            }
            if ((!args || args.replicaNum === undefined) && !opts.urn) {
                throw new Error("Missing required property 'replicaNum'");
            }
            if ((!args || args.topicName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'topicName'");
            }
            resourceInputs["cleanUpPolicy"] = args ? args.cleanUpPolicy : undefined;
            resourceInputs["enableWhiteList"] = args ? args.enableWhiteList : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["ipWhiteLists"] = args ? args.ipWhiteLists : undefined;
            resourceInputs["maxMessageBytes"] = args ? args.maxMessageBytes : undefined;
            resourceInputs["note"] = args ? args.note : undefined;
            resourceInputs["partitionNum"] = args ? args.partitionNum : undefined;
            resourceInputs["replicaNum"] = args ? args.replicaNum : undefined;
            resourceInputs["retention"] = args ? args.retention : undefined;
            resourceInputs["segment"] = args ? args.segment : undefined;
            resourceInputs["syncReplicaMinNum"] = args ? args.syncReplicaMinNum : undefined;
            resourceInputs["topicName"] = args ? args.topicName : undefined;
            resourceInputs["uncleanLeaderElectionEnable"] = args ? args.uncleanLeaderElectionEnable : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["forwardCosBucket"] = undefined /*out*/;
            resourceInputs["forwardInterval"] = undefined /*out*/;
            resourceInputs["forwardStatus"] = undefined /*out*/;
            resourceInputs["messageStorageLocation"] = undefined /*out*/;
            resourceInputs["segmentBytes"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Topic.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Topic resources.
 */
export interface TopicState {
    /**
     * Clear log policy, log clear mode, default is `delete`. `delete`: logs are deleted according to the storage time. `compact`: logs are compressed according to the key. `compact, delete`: logs are compressed according to the key and will be deleted according to the storage time.
     */
    cleanUpPolicy?: pulumi.Input<string>;
    /**
     * Create time of the CKafka topic.
     */
    createTime?: pulumi.Input<string>;
    /**
     * Whether to open the ip whitelist, `true`: open, `false`: close.
     */
    enableWhiteList?: pulumi.Input<boolean>;
    /**
     * Data backup cos bucket: the bucket address that is dumped to cos.
     */
    forwardCosBucket?: pulumi.Input<string>;
    /**
     * Periodic frequency of data backup to cos.
     */
    forwardInterval?: pulumi.Input<number>;
    /**
     * Data backup cos status. Valid values: `0`, `1`. `1`: do not open data backup, `0`: open data backup.
     */
    forwardStatus?: pulumi.Input<number>;
    /**
     * Ckafka instance ID.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * Ip whitelist, quota limit, required when enableWhileList=true.
     */
    ipWhiteLists?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Max message bytes.
     */
    maxMessageBytes?: pulumi.Input<number>;
    /**
     * Message storage location.
     */
    messageStorageLocation?: pulumi.Input<string>;
    /**
     * The subject note. It must start with a letter, and the remaining part can contain letters, numbers and dashes (-).
     */
    note?: pulumi.Input<string>;
    /**
     * The number of partition.
     */
    partitionNum?: pulumi.Input<number>;
    /**
     * The number of replica.
     */
    replicaNum?: pulumi.Input<number>;
    /**
     * Message can be selected. Retention time, unit is ms, the current minimum value is 60000ms.
     */
    retention?: pulumi.Input<number>;
    /**
     * Segment scrolling time, in ms, the current minimum is 3600000ms.
     */
    segment?: pulumi.Input<number>;
    /**
     * Number of bytes rolled by shard.
     */
    segmentBytes?: pulumi.Input<number>;
    /**
     * Min number of sync replicas, Default is `1`.
     */
    syncReplicaMinNum?: pulumi.Input<number>;
    /**
     * Name of the CKafka topic. It must start with a letter, the rest can contain letters, numbers and dashes(-).
     */
    topicName?: pulumi.Input<string>;
    /**
     * Whether to allow unsynchronized replicas to be selected as leader, default is `false`, `true: `allowed, `false`: not allowed.
     */
    uncleanLeaderElectionEnable?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a Topic resource.
 */
export interface TopicArgs {
    /**
     * Clear log policy, log clear mode, default is `delete`. `delete`: logs are deleted according to the storage time. `compact`: logs are compressed according to the key. `compact, delete`: logs are compressed according to the key and will be deleted according to the storage time.
     */
    cleanUpPolicy?: pulumi.Input<string>;
    /**
     * Whether to open the ip whitelist, `true`: open, `false`: close.
     */
    enableWhiteList?: pulumi.Input<boolean>;
    /**
     * Ckafka instance ID.
     */
    instanceId: pulumi.Input<string>;
    /**
     * Ip whitelist, quota limit, required when enableWhileList=true.
     */
    ipWhiteLists?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Max message bytes.
     */
    maxMessageBytes?: pulumi.Input<number>;
    /**
     * The subject note. It must start with a letter, and the remaining part can contain letters, numbers and dashes (-).
     */
    note?: pulumi.Input<string>;
    /**
     * The number of partition.
     */
    partitionNum: pulumi.Input<number>;
    /**
     * The number of replica.
     */
    replicaNum: pulumi.Input<number>;
    /**
     * Message can be selected. Retention time, unit is ms, the current minimum value is 60000ms.
     */
    retention?: pulumi.Input<number>;
    /**
     * Segment scrolling time, in ms, the current minimum is 3600000ms.
     */
    segment?: pulumi.Input<number>;
    /**
     * Min number of sync replicas, Default is `1`.
     */
    syncReplicaMinNum?: pulumi.Input<number>;
    /**
     * Name of the CKafka topic. It must start with a letter, the rest can contain letters, numbers and dashes(-).
     */
    topicName: pulumi.Input<string>;
    /**
     * Whether to allow unsynchronized replicas to be selected as leader, default is `false`, `true: `allowed, `false`: not allowed.
     */
    uncleanLeaderElectionEnable?: pulumi.Input<boolean>;
}
