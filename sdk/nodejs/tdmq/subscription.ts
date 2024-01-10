// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a tdmq subscription
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@tencentcloud_iac/pulumi";
 *
 * const exampleInstance = new tencentcloud.tdmq.Instance("exampleInstance", {
 *     clusterName: "tf_example",
 *     remark: "remark.",
 *     tags: {
 *         createdBy: "terraform",
 *     },
 * });
 * const exampleNamespace = new tencentcloud.tdmq.Namespace("exampleNamespace", {
 *     environName: "tf_example",
 *     msgTtl: 300,
 *     clusterId: exampleInstance.id,
 *     retentionPolicy: {
 *         timeInMinutes: 60,
 *         sizeInMb: 10,
 *     },
 *     remark: "remark.",
 * });
 * const exampleTopic = new tencentcloud.tdmq.Topic("exampleTopic", {
 *     clusterId: exampleInstance.id,
 *     environId: exampleNamespace.environName,
 *     topicName: "tf-example-topic",
 *     partitions: 1,
 *     pulsarTopicType: 3,
 *     remark: "remark.",
 * });
 * const exampleSubscription = new tencentcloud.tdmq.Subscription("exampleSubscription", {
 *     clusterId: exampleInstance.id,
 *     environmentId: exampleNamespace.environName,
 *     topicName: exampleTopic.topicName,
 *     subscriptionName: "tf-example-subscription",
 *     remark: "remark.",
 *     autoCreatePolicyTopic: true,
 *     autoDeletePolicyTopic: true,
 * });
 * ```
 *
 * ## Import
 *
 * tdmq subscription can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:Tdmq/subscription:Subscription example pulsar-q4k5898krpqj#tf_example#tf-example-topic#tf-example-subscription#true
 * ```
 */
export class Subscription extends pulumi.CustomResource {
    /**
     * Get an existing Subscription resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SubscriptionState, opts?: pulumi.CustomResourceOptions): Subscription {
        return new Subscription(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Tdmq/subscription:Subscription';

    /**
     * Returns true if the given object is an instance of Subscription.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Subscription {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Subscription.__pulumiType;
    }

    /**
     * Whether to automatically create a dead letter topic and a retry letter topic. true: yes; false: no(default value).
     */
    public readonly autoCreatePolicyTopic!: pulumi.Output<boolean | undefined>;
    /**
     * Whether to automatically delete a dead letter topic and a retry letter topic. Setting is only allowed when `autoCreatePolicyTopic` is true. Default is false.
     */
    public readonly autoDeletePolicyTopic!: pulumi.Output<boolean | undefined>;
    /**
     * Pulsar cluster ID.
     */
    public readonly clusterId!: pulumi.Output<string>;
    /**
     * Environment (namespace) name.
     */
    public readonly environmentId!: pulumi.Output<string>;
    /**
     * Remarks (up to 128 characters).
     */
    public readonly remark!: pulumi.Output<string | undefined>;
    /**
     * Subscriber name, which can contain up to 128 characters.
     */
    public readonly subscriptionName!: pulumi.Output<string>;
    /**
     * Topic name.
     */
    public readonly topicName!: pulumi.Output<string>;

    /**
     * Create a Subscription resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SubscriptionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SubscriptionArgs | SubscriptionState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SubscriptionState | undefined;
            resourceInputs["autoCreatePolicyTopic"] = state ? state.autoCreatePolicyTopic : undefined;
            resourceInputs["autoDeletePolicyTopic"] = state ? state.autoDeletePolicyTopic : undefined;
            resourceInputs["clusterId"] = state ? state.clusterId : undefined;
            resourceInputs["environmentId"] = state ? state.environmentId : undefined;
            resourceInputs["remark"] = state ? state.remark : undefined;
            resourceInputs["subscriptionName"] = state ? state.subscriptionName : undefined;
            resourceInputs["topicName"] = state ? state.topicName : undefined;
        } else {
            const args = argsOrState as SubscriptionArgs | undefined;
            if ((!args || args.clusterId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterId'");
            }
            if ((!args || args.environmentId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'environmentId'");
            }
            if ((!args || args.subscriptionName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'subscriptionName'");
            }
            if ((!args || args.topicName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'topicName'");
            }
            resourceInputs["autoCreatePolicyTopic"] = args ? args.autoCreatePolicyTopic : undefined;
            resourceInputs["autoDeletePolicyTopic"] = args ? args.autoDeletePolicyTopic : undefined;
            resourceInputs["clusterId"] = args ? args.clusterId : undefined;
            resourceInputs["environmentId"] = args ? args.environmentId : undefined;
            resourceInputs["remark"] = args ? args.remark : undefined;
            resourceInputs["subscriptionName"] = args ? args.subscriptionName : undefined;
            resourceInputs["topicName"] = args ? args.topicName : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Subscription.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Subscription resources.
 */
export interface SubscriptionState {
    /**
     * Whether to automatically create a dead letter topic and a retry letter topic. true: yes; false: no(default value).
     */
    autoCreatePolicyTopic?: pulumi.Input<boolean>;
    /**
     * Whether to automatically delete a dead letter topic and a retry letter topic. Setting is only allowed when `autoCreatePolicyTopic` is true. Default is false.
     */
    autoDeletePolicyTopic?: pulumi.Input<boolean>;
    /**
     * Pulsar cluster ID.
     */
    clusterId?: pulumi.Input<string>;
    /**
     * Environment (namespace) name.
     */
    environmentId?: pulumi.Input<string>;
    /**
     * Remarks (up to 128 characters).
     */
    remark?: pulumi.Input<string>;
    /**
     * Subscriber name, which can contain up to 128 characters.
     */
    subscriptionName?: pulumi.Input<string>;
    /**
     * Topic name.
     */
    topicName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Subscription resource.
 */
export interface SubscriptionArgs {
    /**
     * Whether to automatically create a dead letter topic and a retry letter topic. true: yes; false: no(default value).
     */
    autoCreatePolicyTopic?: pulumi.Input<boolean>;
    /**
     * Whether to automatically delete a dead letter topic and a retry letter topic. Setting is only allowed when `autoCreatePolicyTopic` is true. Default is false.
     */
    autoDeletePolicyTopic?: pulumi.Input<boolean>;
    /**
     * Pulsar cluster ID.
     */
    clusterId: pulumi.Input<string>;
    /**
     * Environment (namespace) name.
     */
    environmentId: pulumi.Input<string>;
    /**
     * Remarks (up to 128 characters).
     */
    remark?: pulumi.Input<string>;
    /**
     * Subscriber name, which can contain up to 128 characters.
     */
    subscriptionName: pulumi.Input<string>;
    /**
     * Topic name.
     */
    topicName: pulumi.Input<string>;
}
