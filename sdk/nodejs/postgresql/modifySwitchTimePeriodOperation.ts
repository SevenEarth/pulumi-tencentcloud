// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a postgresql modifySwitchTimePeriodOperation
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@tencentcloud_iac/pulumi";
 *
 * const modifySwitchTimePeriodOperation = new tencentcloud.postgresql.ModifySwitchTimePeriodOperation("modifySwitchTimePeriodOperation", {
 *     dbInstanceId: local.pgsql_id,
 *     switchTag: 0,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export class ModifySwitchTimePeriodOperation extends pulumi.CustomResource {
    /**
     * Get an existing ModifySwitchTimePeriodOperation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ModifySwitchTimePeriodOperationState, opts?: pulumi.CustomResourceOptions): ModifySwitchTimePeriodOperation {
        return new ModifySwitchTimePeriodOperation(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Postgresql/modifySwitchTimePeriodOperation:ModifySwitchTimePeriodOperation';

    /**
     * Returns true if the given object is an instance of ModifySwitchTimePeriodOperation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ModifySwitchTimePeriodOperation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ModifySwitchTimePeriodOperation.__pulumiType;
    }

    /**
     * The ID of the instance waiting for a switch.
     */
    public readonly dbInstanceId!: pulumi.Output<string>;
    /**
     * Valid value: `0` (switch immediately).
     */
    public readonly switchTag!: pulumi.Output<number>;

    /**
     * Create a ModifySwitchTimePeriodOperation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ModifySwitchTimePeriodOperationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ModifySwitchTimePeriodOperationArgs | ModifySwitchTimePeriodOperationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ModifySwitchTimePeriodOperationState | undefined;
            resourceInputs["dbInstanceId"] = state ? state.dbInstanceId : undefined;
            resourceInputs["switchTag"] = state ? state.switchTag : undefined;
        } else {
            const args = argsOrState as ModifySwitchTimePeriodOperationArgs | undefined;
            if ((!args || args.dbInstanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dbInstanceId'");
            }
            if ((!args || args.switchTag === undefined) && !opts.urn) {
                throw new Error("Missing required property 'switchTag'");
            }
            resourceInputs["dbInstanceId"] = args ? args.dbInstanceId : undefined;
            resourceInputs["switchTag"] = args ? args.switchTag : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ModifySwitchTimePeriodOperation.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ModifySwitchTimePeriodOperation resources.
 */
export interface ModifySwitchTimePeriodOperationState {
    /**
     * The ID of the instance waiting for a switch.
     */
    dbInstanceId?: pulumi.Input<string>;
    /**
     * Valid value: `0` (switch immediately).
     */
    switchTag?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a ModifySwitchTimePeriodOperation resource.
 */
export interface ModifySwitchTimePeriodOperationArgs {
    /**
     * The ID of the instance waiting for a switch.
     */
    dbInstanceId: pulumi.Input<string>;
    /**
     * Valid value: `0` (switch immediately).
     */
    switchTag: pulumi.Input<number>;
}
