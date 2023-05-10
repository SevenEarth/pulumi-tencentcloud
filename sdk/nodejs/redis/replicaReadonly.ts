// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a redis replicaReadonly
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const replicaReadonly = new tencentcloud.Redis.ReplicaReadonly("replica_readonly", {
 *     instanceId: "crs-c1nl9rpv",
 *     operate: "enable",
 *     readonlyPolicies: ["master"],
 * });
 * ```
 */
export class ReplicaReadonly extends pulumi.CustomResource {
    /**
     * Get an existing ReplicaReadonly resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ReplicaReadonlyState, opts?: pulumi.CustomResourceOptions): ReplicaReadonly {
        return new ReplicaReadonly(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Redis/replicaReadonly:ReplicaReadonly';

    /**
     * Returns true if the given object is an instance of ReplicaReadonly.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ReplicaReadonly {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ReplicaReadonly.__pulumiType;
    }

    /**
     * The ID of instance.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * The replica is read-only, `enable` - enable read-write splitting, `disable`- disable read-write splitting.
     */
    public readonly operate!: pulumi.Output<string>;
    /**
     * Routing policy: Enter `master` or `replication`, which indicates the master node or slave node.
     */
    public readonly readonlyPolicies!: pulumi.Output<string[] | undefined>;

    /**
     * Create a ReplicaReadonly resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ReplicaReadonlyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ReplicaReadonlyArgs | ReplicaReadonlyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ReplicaReadonlyState | undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["operate"] = state ? state.operate : undefined;
            resourceInputs["readonlyPolicies"] = state ? state.readonlyPolicies : undefined;
        } else {
            const args = argsOrState as ReplicaReadonlyArgs | undefined;
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            if ((!args || args.operate === undefined) && !opts.urn) {
                throw new Error("Missing required property 'operate'");
            }
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["operate"] = args ? args.operate : undefined;
            resourceInputs["readonlyPolicies"] = args ? args.readonlyPolicies : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ReplicaReadonly.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ReplicaReadonly resources.
 */
export interface ReplicaReadonlyState {
    /**
     * The ID of instance.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * The replica is read-only, `enable` - enable read-write splitting, `disable`- disable read-write splitting.
     */
    operate?: pulumi.Input<string>;
    /**
     * Routing policy: Enter `master` or `replication`, which indicates the master node or slave node.
     */
    readonlyPolicies?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a ReplicaReadonly resource.
 */
export interface ReplicaReadonlyArgs {
    /**
     * The ID of instance.
     */
    instanceId: pulumi.Input<string>;
    /**
     * The replica is read-only, `enable` - enable read-write splitting, `disable`- disable read-write splitting.
     */
    operate: pulumi.Input<string>;
    /**
     * Routing policy: Enter `master` or `replication`, which indicates the master node or slave node.
     */
    readonlyPolicies?: pulumi.Input<pulumi.Input<string>[]>;
}
