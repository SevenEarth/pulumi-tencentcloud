// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a sqlserver renewDbInstance
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const renewDbInstance = new tencentcloud.Sqlserver.RenewDbInstance("renew_db_instance", {
 *     instanceId: "mssql-i1z41iwd",
 *     period: 1,
 * });
 * ```
 *
 * ## Import
 *
 * sqlserver renew_db_instance can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:Sqlserver/renewDbInstance:RenewDbInstance renew_db_instance renew_db_instance_id
 * ```
 */
export class RenewDbInstance extends pulumi.CustomResource {
    /**
     * Get an existing RenewDbInstance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RenewDbInstanceState, opts?: pulumi.CustomResourceOptions): RenewDbInstance {
        return new RenewDbInstance(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Sqlserver/renewDbInstance:RenewDbInstance';

    /**
     * Returns true if the given object is an instance of RenewDbInstance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RenewDbInstance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RenewDbInstance.__pulumiType;
    }

    /**
     * Instance ID.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * How many months to renew, the value range is 1-48, the default is 1.
     */
    public readonly period!: pulumi.Output<number | undefined>;

    /**
     * Create a RenewDbInstance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RenewDbInstanceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RenewDbInstanceArgs | RenewDbInstanceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RenewDbInstanceState | undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["period"] = state ? state.period : undefined;
        } else {
            const args = argsOrState as RenewDbInstanceArgs | undefined;
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["period"] = args ? args.period : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RenewDbInstance.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RenewDbInstance resources.
 */
export interface RenewDbInstanceState {
    /**
     * Instance ID.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * How many months to renew, the value range is 1-48, the default is 1.
     */
    period?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a RenewDbInstance resource.
 */
export interface RenewDbInstanceArgs {
    /**
     * Instance ID.
     */
    instanceId: pulumi.Input<string>;
    /**
     * How many months to renew, the value range is 1-48, the default is 1.
     */
    period?: pulumi.Input<number>;
}
