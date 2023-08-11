// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a redis readOnly
 *
 * ## Example Usage
 * ### Set instance input mode
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@tencentcloud_iac/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const zone = tencentcloud.Redis.getZoneConfig({
 *     typeId: 7,
 * });
 * const vpc = new tencentcloud.vpc.Instance("vpc", {cidrBlock: "10.0.0.0/16"});
 * const subnet = new tencentcloud.subnet.Instance("subnet", {
 *     vpcId: vpc.id,
 *     availabilityZone: zone.then(zone => zone.lists?[1]?.zone),
 *     cidrBlock: "10.0.1.0/24",
 * });
 * const fooGroup = new tencentcloud.security.Group("fooGroup", {});
 * const fooGroupLiteRule = new tencentcloud.security.GroupLiteRule("fooGroupLiteRule", {
 *     securityGroupId: fooGroup.id,
 *     ingresses: [
 *         "ACCEPT#192.168.1.0/24#80#TCP",
 *         "DROP#8.8.8.8#80,90#UDP",
 *         "DROP#0.0.0.0/0#80-90#TCP",
 *     ],
 *     egresses: [
 *         "ACCEPT#192.168.0.0/16#ALL#TCP",
 *         "ACCEPT#10.0.0.0/8#ALL#ICMP",
 *         "DROP#0.0.0.0/0#ALL#ALL",
 *     ],
 * });
 * const fooInstance = new tencentcloud.redis.Instance("fooInstance", {
 *     availabilityZone: zone.then(zone => zone.lists?[0]?.zone),
 *     typeId: zone.then(zone => zone.lists?[0]?.typeId),
 *     password: "test12345789",
 *     memSize: 8192,
 *     redisShardNum: zone.then(zone => zone.lists?[0]?.redisShardNums?[0]),
 *     redisReplicasNum: zone.then(zone => zone.lists?[0]?.redisReplicasNums?[0]),
 *     port: 6379,
 *     vpcId: vpc.id,
 *     subnetId: subnet.id,
 *     securityGroups: [fooGroup.id],
 * });
 * const readOnly = new tencentcloud.redis.ReadOnly("readOnly", {
 *     instanceId: fooInstance.id,
 *     inputMode: "0",
 * });
 * ```
 *
 * ## Import
 *
 * redis read_only can be imported using the instanceId, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:Redis/readOnly:ReadOnly read_only crs-c1nl9rpv
 * ```
 */
export class ReadOnly extends pulumi.CustomResource {
    /**
     * Get an existing ReadOnly resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ReadOnlyState, opts?: pulumi.CustomResourceOptions): ReadOnly {
        return new ReadOnly(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Redis/readOnly:ReadOnly';

    /**
     * Returns true if the given object is an instance of ReadOnly.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ReadOnly {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ReadOnly.__pulumiType;
    }

    /**
     * Instance input mode: `0`: read-write; `1`: read-only.
     */
    public readonly inputMode!: pulumi.Output<string>;
    /**
     * The ID of instance.
     */
    public readonly instanceId!: pulumi.Output<string>;

    /**
     * Create a ReadOnly resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ReadOnlyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ReadOnlyArgs | ReadOnlyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ReadOnlyState | undefined;
            resourceInputs["inputMode"] = state ? state.inputMode : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
        } else {
            const args = argsOrState as ReadOnlyArgs | undefined;
            if ((!args || args.inputMode === undefined) && !opts.urn) {
                throw new Error("Missing required property 'inputMode'");
            }
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            resourceInputs["inputMode"] = args ? args.inputMode : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ReadOnly.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ReadOnly resources.
 */
export interface ReadOnlyState {
    /**
     * Instance input mode: `0`: read-write; `1`: read-only.
     */
    inputMode?: pulumi.Input<string>;
    /**
     * The ID of instance.
     */
    instanceId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ReadOnly resource.
 */
export interface ReadOnlyArgs {
    /**
     * Instance input mode: `0`: read-write; `1`: read-only.
     */
    inputMode: pulumi.Input<string>;
    /**
     * The ID of instance.
     */
    instanceId: pulumi.Input<string>;
}
