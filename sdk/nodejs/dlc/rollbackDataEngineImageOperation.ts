// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a dlc rollbackDataEngineImage
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@tencentcloud_iac/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const checkDataEngineImageCanBeRollback = tencentcloud.Dlc.getCheckDataEngineImageCanBeRollback({
 *     dataEngineId: "DataEngine-cgkvbas6",
 * });
 * const rollbackDataEngineImage = new tencentcloud.dlc.RollbackDataEngineImageOperation("rollbackDataEngineImage", {
 *     dataEngineId: "DataEngine-cgkvbas6",
 *     fromRecordId: checkDataEngineImageCanBeRollback.then(checkDataEngineImageCanBeRollback => checkDataEngineImageCanBeRollback.fromRecordId),
 *     toRecordId: checkDataEngineImageCanBeRollback.then(checkDataEngineImageCanBeRollback => checkDataEngineImageCanBeRollback.toRecordId),
 * });
 * ```
 *
 * ## Import
 *
 * dlc rollback_data_engine_image can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:Dlc/rollbackDataEngineImageOperation:RollbackDataEngineImageOperation rollback_data_engine_image rollback_data_engine_image_id
 * ```
 */
export class RollbackDataEngineImageOperation extends pulumi.CustomResource {
    /**
     * Get an existing RollbackDataEngineImageOperation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RollbackDataEngineImageOperationState, opts?: pulumi.CustomResourceOptions): RollbackDataEngineImageOperation {
        return new RollbackDataEngineImageOperation(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Dlc/rollbackDataEngineImageOperation:RollbackDataEngineImageOperation';

    /**
     * Returns true if the given object is an instance of RollbackDataEngineImageOperation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RollbackDataEngineImageOperation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RollbackDataEngineImageOperation.__pulumiType;
    }

    /**
     * Engine unique id.
     */
    public readonly dataEngineId!: pulumi.Output<string>;
    /**
     * Log record id before rollback.
     */
    public readonly fromRecordId!: pulumi.Output<string | undefined>;
    /**
     * Log record id after rollback.
     */
    public readonly toRecordId!: pulumi.Output<string | undefined>;

    /**
     * Create a RollbackDataEngineImageOperation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RollbackDataEngineImageOperationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RollbackDataEngineImageOperationArgs | RollbackDataEngineImageOperationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RollbackDataEngineImageOperationState | undefined;
            resourceInputs["dataEngineId"] = state ? state.dataEngineId : undefined;
            resourceInputs["fromRecordId"] = state ? state.fromRecordId : undefined;
            resourceInputs["toRecordId"] = state ? state.toRecordId : undefined;
        } else {
            const args = argsOrState as RollbackDataEngineImageOperationArgs | undefined;
            if ((!args || args.dataEngineId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dataEngineId'");
            }
            resourceInputs["dataEngineId"] = args ? args.dataEngineId : undefined;
            resourceInputs["fromRecordId"] = args ? args.fromRecordId : undefined;
            resourceInputs["toRecordId"] = args ? args.toRecordId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RollbackDataEngineImageOperation.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RollbackDataEngineImageOperation resources.
 */
export interface RollbackDataEngineImageOperationState {
    /**
     * Engine unique id.
     */
    dataEngineId?: pulumi.Input<string>;
    /**
     * Log record id before rollback.
     */
    fromRecordId?: pulumi.Input<string>;
    /**
     * Log record id after rollback.
     */
    toRecordId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RollbackDataEngineImageOperation resource.
 */
export interface RollbackDataEngineImageOperationArgs {
    /**
     * Engine unique id.
     */
    dataEngineId: pulumi.Input<string>;
    /**
     * Log record id before rollback.
     */
    fromRecordId?: pulumi.Input<string>;
    /**
     * Log record id after rollback.
     */
    toRecordId?: pulumi.Input<string>;
}
