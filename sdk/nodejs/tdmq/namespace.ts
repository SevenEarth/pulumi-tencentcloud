// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Provide a resource to create a tdmq namespace.
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
 * ```
 *
 * ## Import
 *
 * Tdmq namespace can be imported, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:Tdmq/namespace:Namespace test namespace_id
 * ```
 */
export class Namespace extends pulumi.CustomResource {
    /**
     * Get an existing Namespace resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NamespaceState, opts?: pulumi.CustomResourceOptions): Namespace {
        return new Namespace(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Tdmq/namespace:Namespace';

    /**
     * Returns true if the given object is an instance of Namespace.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Namespace {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Namespace.__pulumiType;
    }

    /**
     * The Dedicated Cluster Id.
     */
    public readonly clusterId!: pulumi.Output<string>;
    /**
     * The name of namespace to be created.
     */
    public readonly environName!: pulumi.Output<string>;
    /**
     * The expiration time of unconsumed message.
     */
    public readonly msgTtl!: pulumi.Output<number>;
    /**
     * Description of the namespace.
     */
    public readonly remark!: pulumi.Output<string | undefined>;
    /**
     * The Policy of message to retain. Format like: `{time_in_minutes: Int, size_in_mb: Int}`. `timeInMinutes`: the time of message to retain; `sizeInMb`: the size of message to retain.
     */
    public readonly retentionPolicy!: pulumi.Output<outputs.Tdmq.NamespaceRetentionPolicy>;

    /**
     * Create a Namespace resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NamespaceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NamespaceArgs | NamespaceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NamespaceState | undefined;
            resourceInputs["clusterId"] = state ? state.clusterId : undefined;
            resourceInputs["environName"] = state ? state.environName : undefined;
            resourceInputs["msgTtl"] = state ? state.msgTtl : undefined;
            resourceInputs["remark"] = state ? state.remark : undefined;
            resourceInputs["retentionPolicy"] = state ? state.retentionPolicy : undefined;
        } else {
            const args = argsOrState as NamespaceArgs | undefined;
            if ((!args || args.clusterId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterId'");
            }
            if ((!args || args.environName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'environName'");
            }
            if ((!args || args.msgTtl === undefined) && !opts.urn) {
                throw new Error("Missing required property 'msgTtl'");
            }
            resourceInputs["clusterId"] = args ? args.clusterId : undefined;
            resourceInputs["environName"] = args ? args.environName : undefined;
            resourceInputs["msgTtl"] = args ? args.msgTtl : undefined;
            resourceInputs["remark"] = args ? args.remark : undefined;
            resourceInputs["retentionPolicy"] = args ? args.retentionPolicy : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Namespace.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Namespace resources.
 */
export interface NamespaceState {
    /**
     * The Dedicated Cluster Id.
     */
    clusterId?: pulumi.Input<string>;
    /**
     * The name of namespace to be created.
     */
    environName?: pulumi.Input<string>;
    /**
     * The expiration time of unconsumed message.
     */
    msgTtl?: pulumi.Input<number>;
    /**
     * Description of the namespace.
     */
    remark?: pulumi.Input<string>;
    /**
     * The Policy of message to retain. Format like: `{time_in_minutes: Int, size_in_mb: Int}`. `timeInMinutes`: the time of message to retain; `sizeInMb`: the size of message to retain.
     */
    retentionPolicy?: pulumi.Input<inputs.Tdmq.NamespaceRetentionPolicy>;
}

/**
 * The set of arguments for constructing a Namespace resource.
 */
export interface NamespaceArgs {
    /**
     * The Dedicated Cluster Id.
     */
    clusterId: pulumi.Input<string>;
    /**
     * The name of namespace to be created.
     */
    environName: pulumi.Input<string>;
    /**
     * The expiration time of unconsumed message.
     */
    msgTtl: pulumi.Input<number>;
    /**
     * Description of the namespace.
     */
    remark?: pulumi.Input<string>;
    /**
     * The Policy of message to retain. Format like: `{time_in_minutes: Int, size_in_mb: Int}`. `timeInMinutes`: the time of message to retain; `sizeInMb`: the size of message to retain.
     */
    retentionPolicy?: pulumi.Input<inputs.Tdmq.NamespaceRetentionPolicy>;
}
