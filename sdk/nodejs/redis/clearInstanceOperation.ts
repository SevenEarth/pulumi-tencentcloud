// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a redis clearInstanceOperation
 *
 * ## Example Usage
 * ### Clear the instance data of the Redis instance
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@tencentcloud_iac/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const config = new pulumi.Config();
 * const password = config.get("password") || "test12345789";
 * const zone = tencentcloud.Redis.getZoneConfig({
 *     typeId: 7,
 * });
 * const vpc = new tencentcloud.vpc.Instance("vpc", {cidrBlock: "10.0.0.0/16"});
 * const subnet = new tencentcloud.subnet.Instance("subnet", {
 *     vpcId: vpc.id,
 *     availabilityZone: zone.then(zone => zone.lists?[1]?.zone),
 *     cidrBlock: "10.0.1.0/24",
 * });
 * const foo = new tencentcloud.redis.Instance("foo", {
 *     availabilityZone: zone.then(zone => zone.lists?[1]?.zone),
 *     typeId: zone.then(zone => zone.lists?[1]?.typeId),
 *     password: password,
 *     memSize: 8192,
 *     redisShardNum: zone.then(zone => zone.lists?[1]?.redisShardNums?[0]),
 *     redisReplicasNum: zone.then(zone => zone.lists?[1]?.redisReplicasNums?[0]),
 *     port: 6379,
 *     vpcId: vpc.id,
 *     subnetId: subnet.id,
 * });
 * const clearInstanceOperation = new tencentcloud.redis.ClearInstanceOperation("clearInstanceOperation", {
 *     instanceId: foo.id,
 *     password: password,
 * });
 * ```
 */
export class ClearInstanceOperation extends pulumi.CustomResource {
    /**
     * Get an existing ClearInstanceOperation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ClearInstanceOperationState, opts?: pulumi.CustomResourceOptions): ClearInstanceOperation {
        return new ClearInstanceOperation(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Redis/clearInstanceOperation:ClearInstanceOperation';

    /**
     * Returns true if the given object is an instance of ClearInstanceOperation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ClearInstanceOperation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ClearInstanceOperation.__pulumiType;
    }

    /**
     * The ID of instance.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * Redis instance password (password-free instances do not need to pass passwords, non-password-free instances must be transmitted).
     */
    public readonly password!: pulumi.Output<string | undefined>;

    /**
     * Create a ClearInstanceOperation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ClearInstanceOperationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ClearInstanceOperationArgs | ClearInstanceOperationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ClearInstanceOperationState | undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
        } else {
            const args = argsOrState as ClearInstanceOperationArgs | undefined;
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["password"] = args ? args.password : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ClearInstanceOperation.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ClearInstanceOperation resources.
 */
export interface ClearInstanceOperationState {
    /**
     * The ID of instance.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * Redis instance password (password-free instances do not need to pass passwords, non-password-free instances must be transmitted).
     */
    password?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ClearInstanceOperation resource.
 */
export interface ClearInstanceOperationArgs {
    /**
     * The ID of instance.
     */
    instanceId: pulumi.Input<string>;
    /**
     * Redis instance password (password-free instances do not need to pass passwords, non-password-free instances must be transmitted).
     */
    password?: pulumi.Input<string>;
}
