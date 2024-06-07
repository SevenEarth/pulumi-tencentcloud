// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a mysql rollbackStop
 *
 * ## Example Usage
 *
 * ### Revoke the ongoing rollback task of the instance
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@tencentcloud_iac/pulumi";
 *
 * const example = new tencentcloud.mysql.RollbackStop("example", {instanceId: "cdb-fitq5t9h"});
 * ```
 * <!--End PulumiCodeChooser -->
 */
export class RollbackStop extends pulumi.CustomResource {
    /**
     * Get an existing RollbackStop resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RollbackStopState, opts?: pulumi.CustomResourceOptions): RollbackStop {
        return new RollbackStop(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Mysql/rollbackStop:RollbackStop';

    /**
     * Returns true if the given object is an instance of RollbackStop.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RollbackStop {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RollbackStop.__pulumiType;
    }

    /**
     * Cloud database instance ID.
     */
    public readonly instanceId!: pulumi.Output<string>;

    /**
     * Create a RollbackStop resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RollbackStopArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RollbackStopArgs | RollbackStopState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RollbackStopState | undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
        } else {
            const args = argsOrState as RollbackStopArgs | undefined;
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RollbackStop.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RollbackStop resources.
 */
export interface RollbackStopState {
    /**
     * Cloud database instance ID.
     */
    instanceId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RollbackStop resource.
 */
export interface RollbackStopArgs {
    /**
     * Cloud database instance ID.
     */
    instanceId: pulumi.Input<string>;
}
