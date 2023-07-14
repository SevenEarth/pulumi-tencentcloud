// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a cynosdb wan
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const wan = new tencentcloud.Cynosdb.Wan("wan", {
 *     clusterId: "cynosdbmysql-bws8h88b",
 *     instanceGrpId: "cynosdbmysql-grp-lxav0p9z",
 * });
 * ```
 *
 * ## Import
 *
 * cynosdb wan can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:Cynosdb/wan:Wan wan cynosdbmysql-bws8h88b#cynosdbmysql-grp-lxav0p9z
 * ```
 */
export class Wan extends pulumi.CustomResource {
    /**
     * Get an existing Wan resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WanState, opts?: pulumi.CustomResourceOptions): Wan {
        return new Wan(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Cynosdb/wan:Wan';

    /**
     * Returns true if the given object is an instance of Wan.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Wan {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Wan.__pulumiType;
    }

    /**
     * Cluster ID.
     */
    public readonly clusterId!: pulumi.Output<string>;
    /**
     * Instance Group ID.
     */
    public readonly instanceGrpId!: pulumi.Output<string>;
    /**
     * Domain name.
     */
    public /*out*/ readonly wanDomain!: pulumi.Output<string>;
    /**
     * Network ip.
     */
    public /*out*/ readonly wanIp!: pulumi.Output<string>;
    /**
     * Internet port.
     */
    public /*out*/ readonly wanPort!: pulumi.Output<number>;
    /**
     * Internet status.
     */
    public /*out*/ readonly wanStatus!: pulumi.Output<string>;

    /**
     * Create a Wan resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WanArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WanArgs | WanState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as WanState | undefined;
            resourceInputs["clusterId"] = state ? state.clusterId : undefined;
            resourceInputs["instanceGrpId"] = state ? state.instanceGrpId : undefined;
            resourceInputs["wanDomain"] = state ? state.wanDomain : undefined;
            resourceInputs["wanIp"] = state ? state.wanIp : undefined;
            resourceInputs["wanPort"] = state ? state.wanPort : undefined;
            resourceInputs["wanStatus"] = state ? state.wanStatus : undefined;
        } else {
            const args = argsOrState as WanArgs | undefined;
            if ((!args || args.clusterId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterId'");
            }
            if ((!args || args.instanceGrpId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceGrpId'");
            }
            resourceInputs["clusterId"] = args ? args.clusterId : undefined;
            resourceInputs["instanceGrpId"] = args ? args.instanceGrpId : undefined;
            resourceInputs["wanDomain"] = undefined /*out*/;
            resourceInputs["wanIp"] = undefined /*out*/;
            resourceInputs["wanPort"] = undefined /*out*/;
            resourceInputs["wanStatus"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Wan.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Wan resources.
 */
export interface WanState {
    /**
     * Cluster ID.
     */
    clusterId?: pulumi.Input<string>;
    /**
     * Instance Group ID.
     */
    instanceGrpId?: pulumi.Input<string>;
    /**
     * Domain name.
     */
    wanDomain?: pulumi.Input<string>;
    /**
     * Network ip.
     */
    wanIp?: pulumi.Input<string>;
    /**
     * Internet port.
     */
    wanPort?: pulumi.Input<number>;
    /**
     * Internet status.
     */
    wanStatus?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Wan resource.
 */
export interface WanArgs {
    /**
     * Cluster ID.
     */
    clusterId: pulumi.Input<string>;
    /**
     * Instance Group ID.
     */
    instanceGrpId: pulumi.Input<string>;
}
