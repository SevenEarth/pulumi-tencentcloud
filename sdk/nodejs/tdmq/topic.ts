// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provide a resource to create a TDMQ topic.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const foo = new tencentcloud.Tdmq.Instance("foo", {
 *     clusterName: "example",
 *     remark: "this is description.",
 * });
 * const barNamespace = new tencentcloud.Tdmq.Namespace("bar", {
 *     clusterId: foo.id,
 *     environName: "example",
 *     msgTtl: 300,
 *     remark: "this is description.",
 * });
 * const barTopic = new tencentcloud.Tdmq.Topic("bar", {
 *     clusterId: foo.id,
 *     environId: barNamespace.id,
 *     partitions: 6,
 *     remark: "this is description.",
 *     topicName: "example",
 *     topicType: 0,
 * });
 * ```
 *
 * ## Import
 *
 * Tdmq Topic can be imported, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:Tdmq/topic:Topic test topic_id
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
    public static readonly __pulumiType = 'tencentcloud:Tdmq/topic:Topic';

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
     * The Dedicated Cluster Id.
     */
    public readonly clusterId!: pulumi.Output<string>;
    /**
     * Creation time of resource.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * The name of tdmq namespace.
     */
    public readonly environId!: pulumi.Output<string>;
    /**
     * The partitions of topic.
     */
    public readonly partitions!: pulumi.Output<number>;
    /**
     * Pulsar Topic Type 0: Non-persistent non-partitioned 1: Non-persistent partitioned 2: Persistent non-partitioned 3: Persistent partitioned.
     */
    public readonly pulsarTopicType!: pulumi.Output<number>;
    /**
     * Description of the namespace.
     */
    public readonly remark!: pulumi.Output<string | undefined>;
    /**
     * The name of topic to be created.
     */
    public readonly topicName!: pulumi.Output<string>;
    /**
     * This input will be gradually discarded and can be switched to PulsarTopicType parameter 0: Normal message; 1: Global sequential messages; 2: Local sequential messages; 3: Retrying queue; 4: Dead letter queue. The type of topic.
     *
     * @deprecated This input will be gradually discarded and can be switched to PulsarTopicType parameter 0: Normal message; 1: Global sequential messages; 2: Local sequential messages; 3: Retrying queue; 4: Dead letter queue.
     */
    public readonly topicType!: pulumi.Output<number>;

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
            resourceInputs["clusterId"] = state ? state.clusterId : undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["environId"] = state ? state.environId : undefined;
            resourceInputs["partitions"] = state ? state.partitions : undefined;
            resourceInputs["pulsarTopicType"] = state ? state.pulsarTopicType : undefined;
            resourceInputs["remark"] = state ? state.remark : undefined;
            resourceInputs["topicName"] = state ? state.topicName : undefined;
            resourceInputs["topicType"] = state ? state.topicType : undefined;
        } else {
            const args = argsOrState as TopicArgs | undefined;
            if ((!args || args.clusterId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterId'");
            }
            if ((!args || args.environId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'environId'");
            }
            if ((!args || args.partitions === undefined) && !opts.urn) {
                throw new Error("Missing required property 'partitions'");
            }
            if ((!args || args.topicName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'topicName'");
            }
            resourceInputs["clusterId"] = args ? args.clusterId : undefined;
            resourceInputs["environId"] = args ? args.environId : undefined;
            resourceInputs["partitions"] = args ? args.partitions : undefined;
            resourceInputs["pulsarTopicType"] = args ? args.pulsarTopicType : undefined;
            resourceInputs["remark"] = args ? args.remark : undefined;
            resourceInputs["topicName"] = args ? args.topicName : undefined;
            resourceInputs["topicType"] = args ? args.topicType : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
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
     * The Dedicated Cluster Id.
     */
    clusterId?: pulumi.Input<string>;
    /**
     * Creation time of resource.
     */
    createTime?: pulumi.Input<string>;
    /**
     * The name of tdmq namespace.
     */
    environId?: pulumi.Input<string>;
    /**
     * The partitions of topic.
     */
    partitions?: pulumi.Input<number>;
    /**
     * Pulsar Topic Type 0: Non-persistent non-partitioned 1: Non-persistent partitioned 2: Persistent non-partitioned 3: Persistent partitioned.
     */
    pulsarTopicType?: pulumi.Input<number>;
    /**
     * Description of the namespace.
     */
    remark?: pulumi.Input<string>;
    /**
     * The name of topic to be created.
     */
    topicName?: pulumi.Input<string>;
    /**
     * This input will be gradually discarded and can be switched to PulsarTopicType parameter 0: Normal message; 1: Global sequential messages; 2: Local sequential messages; 3: Retrying queue; 4: Dead letter queue. The type of topic.
     *
     * @deprecated This input will be gradually discarded and can be switched to PulsarTopicType parameter 0: Normal message; 1: Global sequential messages; 2: Local sequential messages; 3: Retrying queue; 4: Dead letter queue.
     */
    topicType?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a Topic resource.
 */
export interface TopicArgs {
    /**
     * The Dedicated Cluster Id.
     */
    clusterId: pulumi.Input<string>;
    /**
     * The name of tdmq namespace.
     */
    environId: pulumi.Input<string>;
    /**
     * The partitions of topic.
     */
    partitions: pulumi.Input<number>;
    /**
     * Pulsar Topic Type 0: Non-persistent non-partitioned 1: Non-persistent partitioned 2: Persistent non-partitioned 3: Persistent partitioned.
     */
    pulsarTopicType?: pulumi.Input<number>;
    /**
     * Description of the namespace.
     */
    remark?: pulumi.Input<string>;
    /**
     * The name of topic to be created.
     */
    topicName: pulumi.Input<string>;
    /**
     * This input will be gradually discarded and can be switched to PulsarTopicType parameter 0: Normal message; 1: Global sequential messages; 2: Local sequential messages; 3: Retrying queue; 4: Dead letter queue. The type of topic.
     *
     * @deprecated This input will be gradually discarded and can be switched to PulsarTopicType parameter 0: Normal message; 1: Global sequential messages; 2: Local sequential messages; 3: Retrying queue; 4: Dead letter queue.
     */
    topicType?: pulumi.Input<number>;
}
