// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a tdmq rocketmqVipInstance
 *
 * > **NOTE:** The instance cannot be downgraded, Include parameters `nodeCount`, `spec`, `storageSize`.
 */
export class RocketmqVipInstance extends pulumi.CustomResource {
    /**
     * Get an existing RocketmqVipInstance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RocketmqVipInstanceState, opts?: pulumi.CustomResourceOptions): RocketmqVipInstance {
        return new RocketmqVipInstance(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Tdmq/rocketmqVipInstance:RocketmqVipInstance';

    /**
     * Returns true if the given object is an instance of RocketmqVipInstance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RocketmqVipInstance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RocketmqVipInstance.__pulumiType;
    }

    /**
     * Instance name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Number of nodes, minimum 2, maximum 20.
     */
    public readonly nodeCount!: pulumi.Output<number>;
    /**
     * Instance specification: Basic type: `rocket-vip-basic-1`, Standard type: `rocket-vip-basic-2`, Advanced Type I: `rocket-vip-basic-3`, Advanced Type II: `rocket-vip-basic-4`.
     */
    public readonly spec!: pulumi.Output<string>;
    /**
     * Single node storage space, in GB, minimum 200GB.
     */
    public readonly storageSize!: pulumi.Output<number>;
    /**
     * Purchase period, in months.
     */
    public readonly timeSpan!: pulumi.Output<number>;
    /**
     * VPC information.
     */
    public readonly vpcInfo!: pulumi.Output<outputs.Tdmq.RocketmqVipInstanceVpcInfo>;
    /**
     * The Zone ID list for node deployment, such as Guangzhou Zone 1, is 100001. For details, please refer to the official website of Tencent Cloud.
     */
    public readonly zoneIds!: pulumi.Output<string[]>;

    /**
     * Create a RocketmqVipInstance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RocketmqVipInstanceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RocketmqVipInstanceArgs | RocketmqVipInstanceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RocketmqVipInstanceState | undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["nodeCount"] = state ? state.nodeCount : undefined;
            resourceInputs["spec"] = state ? state.spec : undefined;
            resourceInputs["storageSize"] = state ? state.storageSize : undefined;
            resourceInputs["timeSpan"] = state ? state.timeSpan : undefined;
            resourceInputs["vpcInfo"] = state ? state.vpcInfo : undefined;
            resourceInputs["zoneIds"] = state ? state.zoneIds : undefined;
        } else {
            const args = argsOrState as RocketmqVipInstanceArgs | undefined;
            if ((!args || args.nodeCount === undefined) && !opts.urn) {
                throw new Error("Missing required property 'nodeCount'");
            }
            if ((!args || args.spec === undefined) && !opts.urn) {
                throw new Error("Missing required property 'spec'");
            }
            if ((!args || args.storageSize === undefined) && !opts.urn) {
                throw new Error("Missing required property 'storageSize'");
            }
            if ((!args || args.timeSpan === undefined) && !opts.urn) {
                throw new Error("Missing required property 'timeSpan'");
            }
            if ((!args || args.vpcInfo === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vpcInfo'");
            }
            if ((!args || args.zoneIds === undefined) && !opts.urn) {
                throw new Error("Missing required property 'zoneIds'");
            }
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["nodeCount"] = args ? args.nodeCount : undefined;
            resourceInputs["spec"] = args ? args.spec : undefined;
            resourceInputs["storageSize"] = args ? args.storageSize : undefined;
            resourceInputs["timeSpan"] = args ? args.timeSpan : undefined;
            resourceInputs["vpcInfo"] = args ? args.vpcInfo : undefined;
            resourceInputs["zoneIds"] = args ? args.zoneIds : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RocketmqVipInstance.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RocketmqVipInstance resources.
 */
export interface RocketmqVipInstanceState {
    /**
     * Instance name.
     */
    name?: pulumi.Input<string>;
    /**
     * Number of nodes, minimum 2, maximum 20.
     */
    nodeCount?: pulumi.Input<number>;
    /**
     * Instance specification: Basic type: `rocket-vip-basic-1`, Standard type: `rocket-vip-basic-2`, Advanced Type I: `rocket-vip-basic-3`, Advanced Type II: `rocket-vip-basic-4`.
     */
    spec?: pulumi.Input<string>;
    /**
     * Single node storage space, in GB, minimum 200GB.
     */
    storageSize?: pulumi.Input<number>;
    /**
     * Purchase period, in months.
     */
    timeSpan?: pulumi.Input<number>;
    /**
     * VPC information.
     */
    vpcInfo?: pulumi.Input<inputs.Tdmq.RocketmqVipInstanceVpcInfo>;
    /**
     * The Zone ID list for node deployment, such as Guangzhou Zone 1, is 100001. For details, please refer to the official website of Tencent Cloud.
     */
    zoneIds?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a RocketmqVipInstance resource.
 */
export interface RocketmqVipInstanceArgs {
    /**
     * Instance name.
     */
    name?: pulumi.Input<string>;
    /**
     * Number of nodes, minimum 2, maximum 20.
     */
    nodeCount: pulumi.Input<number>;
    /**
     * Instance specification: Basic type: `rocket-vip-basic-1`, Standard type: `rocket-vip-basic-2`, Advanced Type I: `rocket-vip-basic-3`, Advanced Type II: `rocket-vip-basic-4`.
     */
    spec: pulumi.Input<string>;
    /**
     * Single node storage space, in GB, minimum 200GB.
     */
    storageSize: pulumi.Input<number>;
    /**
     * Purchase period, in months.
     */
    timeSpan: pulumi.Input<number>;
    /**
     * VPC information.
     */
    vpcInfo: pulumi.Input<inputs.Tdmq.RocketmqVipInstanceVpcInfo>;
    /**
     * The Zone ID list for node deployment, such as Guangzhou Zone 1, is 100001. For details, please refer to the official website of Tencent Cloud.
     */
    zoneIds: pulumi.Input<pulumi.Input<string>[]>;
}
