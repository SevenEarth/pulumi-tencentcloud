// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a tdmqRocketmq environmentRole
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@tencentcloud_iac/pulumi";
 *
 * const exampleRocketmqCluster = new tencentcloud.tdmq.RocketmqCluster("exampleRocketmqCluster", {
 *     clusterName: "tf_example",
 *     remark: "remark.",
 * });
 * const exampleRocketmqRole = new tencentcloud.tdmq.RocketmqRole("exampleRocketmqRole", {
 *     roleName: "tf_example_role",
 *     remark: "remark.",
 *     clusterId: exampleRocketmqCluster.clusterId,
 * });
 * const exampleRocketmqNamespace = new tencentcloud.tdmq.RocketmqNamespace("exampleRocketmqNamespace", {
 *     clusterId: exampleRocketmqCluster.clusterId,
 *     namespaceName: "tf_example_namespace",
 *     remark: "remark.",
 * });
 * const exampleRocketmqEnvironmentRole = new tencentcloud.tdmq.RocketmqEnvironmentRole("exampleRocketmqEnvironmentRole", {
 *     environmentName: exampleRocketmqNamespace.namespaceName,
 *     roleName: exampleRocketmqRole.roleName,
 *     permissions: [
 *         "produce",
 *         "consume",
 *     ],
 *     clusterId: exampleRocketmqCluster.clusterId,
 * });
 * ```
 *
 * ## Import
 *
 * tdmqRocketmq environment_role can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:Tdmq/rocketmqEnvironmentRole:RocketmqEnvironmentRole environment_role environmentRole_id
 * ```
 */
export class RocketmqEnvironmentRole extends pulumi.CustomResource {
    /**
     * Get an existing RocketmqEnvironmentRole resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RocketmqEnvironmentRoleState, opts?: pulumi.CustomResourceOptions): RocketmqEnvironmentRole {
        return new RocketmqEnvironmentRole(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Tdmq/rocketmqEnvironmentRole:RocketmqEnvironmentRole';

    /**
     * Returns true if the given object is an instance of RocketmqEnvironmentRole.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RocketmqEnvironmentRole {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RocketmqEnvironmentRole.__pulumiType;
    }

    /**
     * Cluster ID (required).
     */
    public readonly clusterId!: pulumi.Output<string>;
    /**
     * Environment (namespace) name.
     */
    public readonly environmentName!: pulumi.Output<string>;
    /**
     * Permissions, which is a non-empty string array of `produce` and `consume` at the most.
     */
    public readonly permissions!: pulumi.Output<string[]>;
    /**
     * Role Name.
     */
    public readonly roleName!: pulumi.Output<string>;

    /**
     * Create a RocketmqEnvironmentRole resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RocketmqEnvironmentRoleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RocketmqEnvironmentRoleArgs | RocketmqEnvironmentRoleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RocketmqEnvironmentRoleState | undefined;
            resourceInputs["clusterId"] = state ? state.clusterId : undefined;
            resourceInputs["environmentName"] = state ? state.environmentName : undefined;
            resourceInputs["permissions"] = state ? state.permissions : undefined;
            resourceInputs["roleName"] = state ? state.roleName : undefined;
        } else {
            const args = argsOrState as RocketmqEnvironmentRoleArgs | undefined;
            if ((!args || args.clusterId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterId'");
            }
            if ((!args || args.environmentName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'environmentName'");
            }
            if ((!args || args.permissions === undefined) && !opts.urn) {
                throw new Error("Missing required property 'permissions'");
            }
            if ((!args || args.roleName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'roleName'");
            }
            resourceInputs["clusterId"] = args ? args.clusterId : undefined;
            resourceInputs["environmentName"] = args ? args.environmentName : undefined;
            resourceInputs["permissions"] = args ? args.permissions : undefined;
            resourceInputs["roleName"] = args ? args.roleName : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RocketmqEnvironmentRole.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RocketmqEnvironmentRole resources.
 */
export interface RocketmqEnvironmentRoleState {
    /**
     * Cluster ID (required).
     */
    clusterId?: pulumi.Input<string>;
    /**
     * Environment (namespace) name.
     */
    environmentName?: pulumi.Input<string>;
    /**
     * Permissions, which is a non-empty string array of `produce` and `consume` at the most.
     */
    permissions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Role Name.
     */
    roleName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RocketmqEnvironmentRole resource.
 */
export interface RocketmqEnvironmentRoleArgs {
    /**
     * Cluster ID (required).
     */
    clusterId: pulumi.Input<string>;
    /**
     * Environment (namespace) name.
     */
    environmentName: pulumi.Input<string>;
    /**
     * Permissions, which is a non-empty string array of `produce` and `consume` at the most.
     */
    permissions: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Role Name.
     */
    roleName: pulumi.Input<string>;
}
