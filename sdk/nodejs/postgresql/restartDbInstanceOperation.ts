// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a postgresql restartDbInstanceOperation
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@tencentcloud_iac/pulumi";
 *
 * const restartDbInstanceOperation = new tencentcloud.postgresql.RestartDbInstanceOperation("restartDbInstanceOperation", {dbInstanceId: local.pgsql_id});
 * ```
 */
export class RestartDbInstanceOperation extends pulumi.CustomResource {
    /**
     * Get an existing RestartDbInstanceOperation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RestartDbInstanceOperationState, opts?: pulumi.CustomResourceOptions): RestartDbInstanceOperation {
        return new RestartDbInstanceOperation(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Postgresql/restartDbInstanceOperation:RestartDbInstanceOperation';

    /**
     * Returns true if the given object is an instance of RestartDbInstanceOperation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RestartDbInstanceOperation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RestartDbInstanceOperation.__pulumiType;
    }

    /**
     * dbInstance ID.
     */
    public readonly dbInstanceId!: pulumi.Output<string>;

    /**
     * Create a RestartDbInstanceOperation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RestartDbInstanceOperationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RestartDbInstanceOperationArgs | RestartDbInstanceOperationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RestartDbInstanceOperationState | undefined;
            resourceInputs["dbInstanceId"] = state ? state.dbInstanceId : undefined;
        } else {
            const args = argsOrState as RestartDbInstanceOperationArgs | undefined;
            if ((!args || args.dbInstanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dbInstanceId'");
            }
            resourceInputs["dbInstanceId"] = args ? args.dbInstanceId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RestartDbInstanceOperation.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RestartDbInstanceOperation resources.
 */
export interface RestartDbInstanceOperationState {
    /**
     * dbInstance ID.
     */
    dbInstanceId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RestartDbInstanceOperation resource.
 */
export interface RestartDbInstanceOperationArgs {
    /**
     * dbInstance ID.
     */
    dbInstanceId: pulumi.Input<string>;
}
