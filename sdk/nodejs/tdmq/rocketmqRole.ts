// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a tdmqRocketmq role
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@tencentcloud_iac/pulumi";
 *
 * const cluster = new tencentcloud.tdmq.RocketmqCluster("cluster", {
 *     clusterName: "test_rocketmq",
 *     remark: "test recket mq",
 * });
 * const role = new tencentcloud.tdmq.RocketmqRole("role", {
 *     roleName: "test_rocketmq_role",
 *     remark: "test rocketmq role",
 *     clusterId: cluster.clusterId,
 * });
 * ```
 *
 * ## Import
 *
 * tdmqRocketmq role can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:Tdmq/rocketmqRole:RocketmqRole role role_id
 * ```
 */
export class RocketmqRole extends pulumi.CustomResource {
    /**
     * Get an existing RocketmqRole resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RocketmqRoleState, opts?: pulumi.CustomResourceOptions): RocketmqRole {
        return new RocketmqRole(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Tdmq/rocketmqRole:RocketmqRole';

    /**
     * Returns true if the given object is an instance of RocketmqRole.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RocketmqRole {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RocketmqRole.__pulumiType;
    }

    /**
     * Cluster ID (required).
     */
    public readonly clusterId!: pulumi.Output<string>;
    /**
     * Creation time.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Remarks (up to 128 characters).
     */
    public readonly remark!: pulumi.Output<string | undefined>;
    /**
     * Role name, which can contain up to 32 letters, digits, hyphens, and underscores.
     */
    public readonly roleName!: pulumi.Output<string>;
    /**
     * Value of the role token.
     */
    public /*out*/ readonly token!: pulumi.Output<string>;
    /**
     * Update time.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a RocketmqRole resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RocketmqRoleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RocketmqRoleArgs | RocketmqRoleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RocketmqRoleState | undefined;
            resourceInputs["clusterId"] = state ? state.clusterId : undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["remark"] = state ? state.remark : undefined;
            resourceInputs["roleName"] = state ? state.roleName : undefined;
            resourceInputs["token"] = state ? state.token : undefined;
            resourceInputs["updateTime"] = state ? state.updateTime : undefined;
        } else {
            const args = argsOrState as RocketmqRoleArgs | undefined;
            if ((!args || args.clusterId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterId'");
            }
            if ((!args || args.roleName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'roleName'");
            }
            resourceInputs["clusterId"] = args ? args.clusterId : undefined;
            resourceInputs["remark"] = args ? args.remark : undefined;
            resourceInputs["roleName"] = args ? args.roleName : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["token"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RocketmqRole.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RocketmqRole resources.
 */
export interface RocketmqRoleState {
    /**
     * Cluster ID (required).
     */
    clusterId?: pulumi.Input<string>;
    /**
     * Creation time.
     */
    createTime?: pulumi.Input<string>;
    /**
     * Remarks (up to 128 characters).
     */
    remark?: pulumi.Input<string>;
    /**
     * Role name, which can contain up to 32 letters, digits, hyphens, and underscores.
     */
    roleName?: pulumi.Input<string>;
    /**
     * Value of the role token.
     */
    token?: pulumi.Input<string>;
    /**
     * Update time.
     */
    updateTime?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RocketmqRole resource.
 */
export interface RocketmqRoleArgs {
    /**
     * Cluster ID (required).
     */
    clusterId: pulumi.Input<string>;
    /**
     * Remarks (up to 128 characters).
     */
    remark?: pulumi.Input<string>;
    /**
     * Role name, which can contain up to 32 letters, digits, hyphens, and underscores.
     */
    roleName: pulumi.Input<string>;
}
